package socket

/*
NewEventHandler returns a pointer to an event handler.
*/
func NewEventHandler(
	name string,
	callback func(response *Response),
) EventHandler {
	return &handler{
		callback: callback,
		name:     name,
	}
}

/*
handler implements EventHandler.
*/
type handler struct {
	callback func(response *Response)
	name     string
}

/*
Name implements EventHandler.
*/
func (handler *handler) Name() string {
	return handler.name
}

/*
Handle implements EventHandler.
*/
func (handler *handler) Handle(
	response *Response,
) {
	handler.callback(response)
}
