package socket

import (
	"encoding/json"
	"fmt"
	"sync"

	chrome_error "github.com/mkenney/go-chrome/error"
	log "github.com/sirupsen/logrus"
)

/*
NewCommand creates and returns a pointer to a Command.
*/
func NewCommand(method string, params interface{}) Commander {
	return &command{
		id:     GenerateCommandID(),
		method: method,
		params: &params,
		wg:     &sync.WaitGroup{},
	}
}

/*
GenerateCommandID generates and returns a unique command ID.
*/
func GenerateCommandID() int {
	commandID++
	return commandID
}

var commandID = 0

/*
NewCommandMap creates and returns a pointer to a CommandMapper.
*/
func NewCommandMap() CommandMapper {
	return &commandMap{
		stack: make(map[int]Commander),
		mux:   &sync.Mutex{},
	}
}

//////////////////////////////////////////////////
// Command
//////////////////////////////////////////////////

/*
command implements Commander.
*/
type command struct {
	// err contains any error resulting from executing the command
	err *chrome_error.Error

	// id contains the command ID
	id int

	// method is the Chrome protocol method being executed
	method string

	// Optional. params holds the parameter struct for the command being
	// executed
	params interface{}

	// Optional. result holds the result struct for the command being executed
	result interface{}

	// sync handles command synchronization
	wg *sync.WaitGroup
}

/*
Done implements Commander.
*/
func (cmd *command) Done(result []byte, err error) {
	if err != nil {
		log.Error(err)
		cmd.err = chrome_error.NewFromErr(err)
	}

	err = json.Unmarshal(result, &cmd.result)
	if err != nil {
		log.Error(err)
		cmd.err = chrome_error.New(
			fmt.Sprintf("%s. In addition, a JSON error occurred while decoding the data", cmd.err.Error()),
			chrome_error.LevelError,
			chrome_error.CodeJSONError,
			err,
		)
	}

	cmd.WaitGroup().Done()
}

/*
Error implements Commander.
*/
func (cmd *command) Error() error {
	return cmd.err
}

/*
ID implements Commander.
*/
func (cmd *command) ID() int {
	return cmd.id
}

/*
Method implements Commander.
*/
func (cmd *command) Method() string {
	return cmd.method
}

/*
Params implements Commander.
*/
func (cmd *command) Params() interface{} {
	return cmd.params
}

/*
Result implements Commander.
*/
func (cmd *command) Result() interface{} {
	return cmd.result
}

/*
Sync implements Commander.
*/
func (cmd *command) WaitGroup() *sync.WaitGroup {
	return cmd.wg
}

//////////////////////////////////////////////////
// Command map
//////////////////////////////////////////////////

/*
commandMap implements CommandMapper.
*/
type commandMap struct {
	stack map[int]Commander
	mux   *sync.Mutex
}

/*
Delete implements CommandMapper.
*/
func (stack *commandMap) Delete(id int) {
	delete(stack.stack, id)
}

/*
Get implements CommandMapper.
*/
func (stack *commandMap) Get(id int) (Commander, error) {
	command, ok := stack.stack[id]
	if !ok {
		return nil, fmt.Errorf("Command %d not found", id)
	}
	return command, nil
}

/*
Lock implements CommandMapper.
*/
func (stack *commandMap) Lock() {
	stack.mux.Lock()
}

/*
Set implements CommandMapper.
*/
func (stack *commandMap) Set(id int, cmd Commander) {
	stack.stack[id] = cmd
}

/*
Unlock implements CommandMapper.
*/
func (stack *commandMap) Unlock() {
	stack.mux.Unlock()
}

//////////////////////////////////////////////////
// Command payload
//////////////////////////////////////////////////

/*
commandPayload represents a WebSocket JSON payload for a command.
*/
type commandPayload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

//////////////////////////////////////////////////
// Socketer
//////////////////////////////////////////////////

/*
SendCommand implements Socketer.

Workflow:
	1. The socket's command mutex is locked.
	2. The command counter is incremented.
	3. The payload is sent to the socket connection and the mutex is unlocked.
	4. The command is stored using the generated ID.
	5. When the command has been executed and the socket responds,
	socket.HandleCmd() is triggered from the command instance to generate the
	response and the command unlocks itself.
*/
func (socket *socket) SendCommand(command Commander) *commandPayload {

	// Safely add a command to the internal stack
	socket.commands.Lock()
	payload := &commandPayload{
		ID:     command.ID(),
		Method: command.Method(),
		Params: command.Params(),
	}
	socket.commands.Set(payload.ID, command)
	socket.commands.Unlock() // Do not defer, HandleCommand() also locks

	log.Debugf("Sending command '%s'", payload)
	command.WaitGroup().Add(1)
	if err := socket.Conn().WriteJSON(payload); err != nil {
		command.Done(nil, err)
	}
	command.WaitGroup().Wait()

	return payload
}

/*
HandleCommand implements Socketer.
*/
func (socket *socket) HandleCommand(response *Response) {
	socket.commands.Lock()
	defer socket.commands.Unlock()

	if command, err := socket.commands.Get(response.ID); nil != err {
		log.Errorf("%s: result=%s err=%s", err.Error(), response.Result, response.Error.Message)

	} else {
		socket.commands.Delete(command.ID())
		if nil != response.Error && "" != response.Error.Message {
			err = fmt.Errorf(
				"%d %s: - %s",
				response.Error.Code,
				response.Error.Message,
				response.Error.Data,
			)
		}
		command.Done(response.Result, err)
		log.Infof("Command '%s' complete", command.Method())
	}
}
