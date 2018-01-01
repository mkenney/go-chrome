package socket

import "sync"

/*
NewCommand creates and returns a pointer to a Commander instance.
*/
func NewCommand(method string, params interface{}) *Command {
	return &Command{
		method: method,
		params: params,
		wg:     &sync.WaitGroup{},
		mux:    &sync.Mutex{},
	}
}

/*
Command implements Commander.
*/
type Command struct {
	// err contains any error resulting from executing the command.
	err error

	// id contains the command ID.
	id int

	// method is the Chrome protocol method being executed.
	method string

	// Optional. params holds the parameter struct for the command being
	// executed.
	params interface{}

	// Optional. result holds the JSON results returned by the socket when the
	// command is complete.
	result []byte

	// sync handles command synchronization.
	wg *sync.WaitGroup

	mux *sync.Mutex
}

/*
Done implements Commander.
*/
func (cmd *Command) Done(result []byte, err error) {
	cmd.result = result
	cmd.err = err
	cmd.WaitGroup().Done()
}

/*
Error implements Commander.
*/
func (cmd *Command) Error() error {
	return cmd.err
}

/*
ID implements Commander.
*/
func (cmd *Command) ID() int {
	if 0 == cmd.id {
		cmd.id = nextCommandID()
	}
	return cmd.id
}

/*
Method implements Commander.
*/
func (cmd *Command) Method() string {
	return cmd.method
}

/*
nextCommandID generates and returns a unique command ID.
*/
func nextCommandID() int {
	_commandIDMux.Lock()
	_commandID++
	id := _commandID
	_commandIDMux.Unlock()
	return id
}

var _commandIDMux = &sync.Mutex{}
var _commandID = 0

/*
Params implements Commander.
*/
func (cmd *Command) Params() interface{} {
	return cmd.params
}

/*
Result implements Commander.
*/
func (cmd *Command) Result() []byte {
	return cmd.result
}

/*
WaitGroup implements Commander.
*/
func (cmd *Command) WaitGroup() *sync.WaitGroup {
	return cmd.wg
}
