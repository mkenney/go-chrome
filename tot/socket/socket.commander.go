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
	response chan *Response

	// socket contains the Socketer instance
	socket Socketer
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
	return cmd.id
}

/*
Method implements Commander.
*/
func (cmd *Command) Method() string {
	return cmd.method
}

/*
Params implements Commander.
*/
func (cmd *Command) Params() interface{} {
	return cmd.params
}

/*
Respond implements Commander.
*/
func (cmd *Command) Respond(response *Response) {
	cmd.response <- response
}

/*
Response implements Commander.
*/
func (cmd *Command) Response() chan *Response {
	return cmd.response
}
