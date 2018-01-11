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
Handler implements EventHandler.
*/
type Handler struct {
	callback func(response *Response)
	name     string
}

/*
Name implements EventHandler.
*/
func (handler *Handler) Name() string {
	return handler.name
}

/*
Handle implements EventHandler.
*/
func (handler *Handler) Handle(
	response *Response,
) {
	handler.callback(response)
}
