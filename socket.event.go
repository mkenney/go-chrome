package chrome

import (
	"app/chrome/protocol"

	log "github.com/Sirupsen/logrus"
)

/*
AddEventHandler adds an event handler to the socket
*/
func (socket *Socket) AddEventHandler(event protocol.EventInterface) {
	socket.eventMutex.Lock()
	defer socket.eventMutex.Unlock()

	for _, evt := range socket.events[event.Name()] {
		if evt == event {
			return
		}
	}
	socket.events[event.Name()] = append(socket.events[event.Name()], event)
}

/*
RemoveEventHandler removes an event handler from the socket
*/
func (socket *Socket) RemoveEventHandler(event protocol.EventInterface) {
	socket.eventMutex.Lock()
	defer socket.eventMutex.Unlock()

	events := socket.events[event.Name()]
	for i, evt := range events {
		if evt == sink {
			evtCount := len(events)
			events[i] = events[evtCount-1]
			socket.events[event.Name()] = events[:evtCount-1]
			return
		}
	}
}

func (socket *Socket) handleEvent(response *SocketResponse) {
	if response.Method == "Inspector.targetCrashed" {
		log.Fatalf("Chrome has crashed!")
	}
	socket.eventMutex.Lock()
	defer socket.eventMutex.Unlock()

	log.Debugf("Event received: '%v'", response)
	for _, event := range socket.events[response.Method] {
		log.Debugf("%s handler(s) found", response.Method)
		go event.OnEvent(response.Method, []byte(response.Params))
	}
}
