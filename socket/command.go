package socket

import (
	"encoding/json"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

/*
NewCommand creates and returns a pointer to a Command.
*/
func NewCommand(method string, params interface{}) Commander {
	return &command{
		method: method,
		params: &params,
		sync:   &sync.WaitGroup{},
	}
}

/*
NewCommandMap creates and returns a pointer to a CommandMap.
*/
func NewCommandMap() CommandMapper {
	return &commandMap{
		stack: make(map[int]Commander),
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

	// method is the Chrome protocol method being executed
	method string

	// Optional. params holds the parameter struct for the command being executed
	params interface{}

	// Optional. result holds the result struct for the command being executed
	result interface{}

	// sync handles command synchronization
	sync *sync.WaitGroup
}

/*
Done implements Commander.
*/
func (cmd *command) Done(result []byte, err error) {
	if err == nil {
		err = json.Unmarshal(result, cmd.Result())
	}
	cmd.err = err
	cmd.Sync().Done()
}

/*
Error implements Commander.
*/
func (cmd *command) Error() error {
	return cmd.err
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
	return &cmd.params
}

/*
Result implements Commander.
*/
func (cmd *command) Result() interface{} {
	return &cmd.result
}

/*
Sync implements Commander.
*/
func (cmd *command) Sync() *sync.WaitGroup {
	return cmd.sync
}

//////////////////////////////////////////////////
// Command map
//////////////////////////////////////////////////

/*
commandMap implements CommandMapper.
*/
type commandMap struct {
	stack map[int]Commander
	mux   sync.Mutex
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
	5. When the command has been executed and the socket responds, socket.HandleCmd() is triggered
	from the command instance to generate the response and the command unlocks itself.
*/
func (socket *socket) SendCommand(command Commander) *commandPayload {

	// Safely add a command to the internal stack
	socket.commands.Lock()
	payload := &commandPayload{
		ID:     socket.GenerateCommandID(),
		Method: command.Method(),
		Params: command.Params(),
	}
	socket.commands.Set(payload.ID, command)
	socket.commands.Unlock() // Do not defer, HandleCommand() also locks

	log.Debugf("Sending command '%s'", payload)
	command.Sync().Add(1)
	if err := socket.conn.WriteJSON(payload); err != nil {
		command.Done(nil, err)
	}
	command.Sync().Wait()

	return payload
}

/*
HandleCommand implements Socketer.
*/
func (socket *socket) HandleCommand(response *Response) {
	socket.commands.Lock()
	defer socket.commands.Unlock()

	if command, err := socket.commands.Get(response.ID); nil != err {
		log.Errorf("Command %d not found: result=%s err=%s", response.ID, response.Result, response.Error.Message)

	} else {
		socket.commands.Delete(response.ID)
		if "" != response.Error.Message {
			err = fmt.Errorf("%d %s: - %s", response.Error.Code, response.Error.Message, response.Error.Data)
		}
		command.Done(response.Result, err)
		log.Debugf("Command '%s' complete", command.Method())
	}
}
