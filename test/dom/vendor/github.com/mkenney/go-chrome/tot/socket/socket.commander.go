package socket

/*
NewCommand creates and returns a pointer to a struct that implements the
Commander interface.
*/
func NewCommand(socket Socketer, method string, params interface{}) *Command {
	return &Command{
		id:       socket.NextCommandID(),
		method:   method,
		params:   params,
		response: make(chan *Response),
		socket:   socket,
	}
}

/*
Command provides a Commander interface for sending commands to a websocket.
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
	response chan *Response

	// socket contains the Socketer instance
	socket Socketer
}

/*
Error returns the most recent error, if any.

Error is a Commander implementation.
*/
func (cmd *Command) Error() error {
	return cmd.err
}

/*
ID returns the command ID.

ID is a Commander implementation.
*/
func (cmd *Command) ID() int {
	return cmd.id
}

/*
Method returns the name of the Chrom DevTools Protocol method to be called.

Method is a Commander implementation.
*/
func (cmd *Command) Method() string {
	return cmd.method
}

/*
Params returns the command parameters.

Params is a Commander implementation.
*/
func (cmd *Command) Params() interface{} {
	return cmd.params
}

/*
Respond sends a response to the command response channel.

Respond is a Commander implementation.
*/
func (cmd *Command) Respond(response *Response) {
	cmd.response <- response
}

/*
Response returns the command response channel.

Response is a Commander implementation.
*/
func (cmd *Command) Response() chan *Response {
	return cmd.response
}

/*
SetError sets the error value

SetError is a Commander implementation.
*/
func (cmd *Command) SetError(err error) {
	cmd.err = err
}

/*
SetID sets the ID value

SetID is a Commander implementation.
*/
func (cmd *Command) SetID(id int) {
	cmd.id = id
}
