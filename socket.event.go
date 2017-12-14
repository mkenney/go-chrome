package chrome

import (
	"app/chrome/protocol"

	log "github.com/Sirupsen/logrus"
)

/*
AddEventHandler adds an event handler to the socket
*/
func (socket *Socket) AddEventHandler(handler *protocol.EventHandler) {
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
func (socket *Socket) RemoveEventHandler(event *protocol.EventHandler) {
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

func (socket *Socket) handleEvent(response *SocketResponse) {
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
