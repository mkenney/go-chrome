package socket

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

/*
EventHandlerInterface is the interface definition for a socket event
*/
type EventHandlerInterface interface {
	OnEvent(name string, params []byte)
}

/*
EventHandler is a generic EventHandlerInterface type
*/
type EventHandler struct {
	Callback func(name string, params []byte)
	Name     string
}

/*
NewEventHandler returns a pointer to a generic event handler
*/
func NewEventHandler(name string, callback func(name string, params []byte)) *EventHandler {
	return &EventHandler{
		Callback: callback,
		Name:     name,
	}
}

/*
OnEvent is an EventHandlerInterface implementation
*/
func (e *EventHandler) OnEvent(name string, params []byte) {
	e.Callback(name, params)
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
EventHandlerResponse represents a socket message
*/
type EventHandlerResponse struct {
	Error  Error           `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}

/*
AddEventHandler adds an event handler to the socket
*/
func (socket *Socket) AddEventHandler(handler *EventHandler) {
	socket.Events.Mux.Lock()
	defer socket.Events.Mux.Unlock()

	for _, hndl := range socket.Events.Map[handler.Name] {
		if hndl == handler {
			log.Warnf("Attempted to add a duplicate handler for event '%s'", handler.Name)
			return
		}
	}
	log.Debugf("Adding handler for event '%s'", handler.Name)
	socket.Events.Map[handler.Name] = append(socket.Events.Map[handler.Name], handler)
}

/*
RemoveEventHandler removes an event handler from the socket
*/
func (socket *Socket) RemoveEventHandler(event *EventHandler) {
	socket.Events.Mux.Lock()
	defer socket.Events.Mux.Unlock()

	events := socket.Events.Map[event.Name]
	for i, evt := range events {
		if evt == event {
			evtCount := len(events)
			events[i] = events[evtCount-1]
			socket.Events.Map[event.Name] = events[:evtCount-1]
			return
		}
	}
}

func (socket *Socket) handleEvent(response *Response) {
	log.Debugf("%s: %s event received", socket.URL, response.Method)
	if response.Method == "Inspector.targetCrashed" {
		log.Fatalf("Chrome has crashed!")
	}
	socket.Events.Mux.Lock()
	defer socket.Events.Mux.Unlock()

	for _, event := range socket.Events.Map[response.Method] {
		log.Infof("Handling event '%s'", response.Method)
		go event.OnEvent(response.Method, []byte(response.Params))
	}
}
