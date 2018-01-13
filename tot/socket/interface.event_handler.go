package socket

/*
EventHandler defines the event handler interface required by the event handler
process.
*/
type EventHandler interface {
	// Handle executes the event handler callback.
	Handle(response *Response)

	// Name returns the name of the event the handler is assigned to.
	Name() string
}
