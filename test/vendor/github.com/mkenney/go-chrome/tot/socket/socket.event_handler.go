package socket

/*
NewEventHandler returns a pointer to an event handler.
*/
func NewEventHandler(
	name string,
	callback func(response *Response),
) *Handler {
	return &Handler{
		callback: callback,
		name:     name,
	}
}

/*
Handler provides an EventHandler interface for managing an event handler.
*/
type Handler struct {
	callback func(response *Response)
	name     string
}

/*
Handle executes the event handler callback.

Handle is an EventHandler implementation.
*/
func (handler *Handler) Handle(
	response *Response,
) {
	handler.callback(response)
}

/*
Name returns the name of the event the handler is assigned to.

Name is an EventHandler implementation.
*/
func (handler *Handler) Name() string {
	return handler.name
}
