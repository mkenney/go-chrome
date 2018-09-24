package socket

import (
	"net/url"
)

/*
Socketer defines the websocket connection interface for managing sockets.
*/
type Socketer interface {
	// AddEventHandler adds an event handler to the stack of listeners for an
	// event.
	AddEventHandler(handler EventHandler)

	// CurCommandID returns the latest command ID.
	CurCommandID() int

	// Errors returns a channel of errors
	Errors() chan error

	// Listen starts the socket read loop and delivers messages to
	// HandleCommand() and HandleEvent() as appropriate.
	Listen()

	// NextCommandID generates and returns the next command ID.
	NextCommandID() int

	// RemoveEventHandler removes a handler from the stack of listeners for an
	// event.
	RemoveEventHandler(handler EventHandler) error

	// SendCommand delivers a command payload to the websocket connection.
	SendCommand(command Commander) chan *Response

	// Stop signals the socket read loop to stop listening for data and close
	// the websocket connection.
	Stop()

	// URL returns the URL of the websocket connection.
	URL() *url.URL
}
