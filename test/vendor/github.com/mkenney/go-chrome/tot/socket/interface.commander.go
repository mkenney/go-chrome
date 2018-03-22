package socket

/*
Commander defines the interface for websocket commands.
*/
type Commander interface {
	// Error returns the most recent error, if any.
	Error() error

	// ID returns the command ID.
	ID() int

	// Method returns the name of the Chrom DevTools Protocol method to be
	// called.
	Method() string

	// Params returns the command parameters.
	Params() interface{}

	// Respond sends a response to the command response channel.
	Respond(response *Response)

	// Response returns the command response channel.
	Response() chan *Response

	// SetError sets the error value
	SetError(error)

	// SetID sets the ID value
	SetID(int)
}
