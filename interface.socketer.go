package chrome

import (
	"github.com/mkenney/go-chrome/socket"
)

/*
Socketer defins a simple socket interface for the tab struct to implement
*/
type Socketer interface {
	// AddEventHandler adds an event handler to the stack of listeners for an
	// event.
	AddEventHandler(handler socket.EventHandler)

	// RemoveEventHandler removes a handler from the stack of listeners for an
	// event.
	RemoveEventHandler(handler socket.EventHandler)

	// SendCommand delivers a command payload to the websocket connection.
	SendCommand(command socket.Commander) *socket.Payload
}
