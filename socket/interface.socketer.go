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

	// Connect creates a new websocket connection.
	Connect() error

	// Disconnect closes the current socket connection.
	Disconnect() error

	// HandleCommand receives the responses to requests sent to the websocket
	// connection.
	HandleCommand(response *Response)

	// HandlEvent receives all events and associated data read from the
	// websocket connection.
	HandleEvent(response *Response)

	// Listen starts the socket read loop and delivers messages to
	// HandleCommand() and HandleEvent() as appropriate
	Listen() error

	// RemoveEventHandler removes a handler from the stack of listeners for an
	// event.
	RemoveEventHandler(handler EventHandler)

	// SendCommand delivers a command payload to the websocket connection.
	SendCommand(command Commander) *Payload

	// Stop signals the socket read loop to stop listening for data and close
	// the websocket connection.
	Stop()

	// URL returns the URL of the websocket connection
	URL() *url.URL
}
