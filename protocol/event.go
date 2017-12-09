package protocol

import (
	"encoding/json"
)

/*
EventHandlerInterface is the interface definition for a socket event
*/
type EventHandlerInterface interface {
	OnEvent(name string, params []byte)
	Name() string
}

/*
EventHandler is a generic EventHandlerInterface type
*/
type EventHandler struct {
	callback func(name string, params []byte)
	name     string
}

/*
NewEventHandler returns a pointer to a generic event handler
*/
func NewEventHandler(name string, callback func(name string, params []byte)) *EventHandler {
	return &EventHandler{
		name:     name,
		callback: callback,
	}
}

/*
OnEvent is an EventHandlerInterface implementation
*/
func (e *EventHandler) OnEvent(name string, params []byte) {
	e.callback(name, params)
}

/*
Name is an EventHandlerInterface implementation
*/
func (e *EventHandler) Name() string {
	return e.name
}

/*
EventHandlerResult represents the result of the socket event
*/
type EventHandlerResult struct {
	Data string `json:"data"`
}

/*
EventHandlerPayload is a representation of a WebSocket JSON payload
*/
type EventHandlerPayload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

/*
NewEventHandlerPayload generates a new EventHandlerPayload pointer
*/
func NewEventHandlerPayload(id int, method string, params interface{}) *EventHandlerPayload {
	return &EventHandlerPayload{
		ID:     id,
		Method: method,
		Params: params,
	}
}

/*
SocketError represents an error
*/
type SocketError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
EventHandlerResponse represents a socket message
*/
type EventHandlerResponse struct {
	Error  SocketError     `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}
