package socket

import (
	"fmt"
	"sync"

	"github.com/mkenney/go-chrome/errors"
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
	_commandID++
	return _commandID
}

var _commandID = 0

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
	err error

	// id contains the command ID
	id int

	// method is the Chrome protocol method being executed
	method string

	// Optional. params holds the parameter struct for the command being
	// executed
	params interface{}

	// Optional. result holds the JSON returned by the socket when the command
	// is complete
	result []byte

	// sync handles command synchronization
	wg *sync.WaitGroup
}

/*
Done implements Commander.
*/
func (cmd *command) Done(result []byte, err error) {
	cmd.result = result
	cmd.err = err
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
func (cmd *command) Result() []byte {
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
func (stack *commandMap) Set(cmd Commander) {
	stack.stack[cmd.ID()] = cmd
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
	// Add the command to the internal stack and create the payload
	socket.commands.Set(command)
	payload := &commandPayload{
		ID:     command.ID(),
		Method: command.Method(),
		Params: command.Params(),
	}

	command.WaitGroup().Add(1)
	if err := socket.Conn().WriteJSON(payload); err != nil {
		command.Done(nil, errors.SocketWriteFailed{Type: errors.Type{
			Caller: errors.GetCaller(),
			Err:    err,
			Msg:    "Failed to send command payload to socket connection",
		}})
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
		log.Errorf(
			"%s: result=%s err=%s",
			err.Error(),
			response.Result,
			response.Error.Message,
		)

	} else {
		socket.commands.Delete(command.ID())
		if nil != response.Error && "" != response.Error.Message {
			log.Errorf(
				"Socket command responded with error: result=%s err=%s",
				response.Result,
				response.Error.Message,
			)
			err = errors.SocketErrorResponse{Type: errors.Type{
				Caller: errors.GetCaller(),
				Err: fmt.Errorf(
					"%d %s: %s",
					response.Error.Code,
					response.Error.Message,
					response.Error.Data,
				),
				Msg: "The socket command failed",
			}}
		}

		command.Done(response.Result, err)
		log.Debugf(
			"Command #%d complete: - %s {%s}",
			command.ID(),
			socket.URL(),
			command.Method(),
		)
	}
}
