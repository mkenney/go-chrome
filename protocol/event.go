package protocol

import (
	"encoding/json"
)

/*
EventInterface is the interface definition for a socket event
*/
type EventInterface interface {
	OnEvent(name string, params []byte)
	Name() string
}

/*
Event is a generic EventInterface type
*/
type Event struct {
	callback func(name string, params []byte)
	name     func() string
}

func NewEvent(name string, callback func(name string, params []byte)) *Event {
	event = new(Event)
	event.name = name
	event.callback = callback
	return event
}

/*
OnEvent is an EventInterface implementation
*/
func (e *Event) OnEvent(name string, params []byte) {
	e.Callback(name, params)
}

/*
Name is an EventInterface implementation
*/
func (e *Event) Name() string {
	return e.name
}

/*
EventResult represents the result of the socket event
*/
type EventResult struct {
	Data string `json:"data"`
}

/*
EventPayload is a representation of a WebSocket JSON payload
*/
type EventPayload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

/*
NewEventPayload generates a new EventPayload pointer
*/
func NewEventPayload(id int, method string, params interface{}) *EventPayload {
	payload := new(EventPayload)
	payload.ID = id
	payload.Method = method
	payload.Params = params
	return payload
}

/*
EventResponse represents a socket message
*/
type EventResponse struct {
	Error  SocketError     `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}
