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

	// HandleCommand receives the responses to requests sent to the websocket
	// connection.
	HandleCommand(response *Response)

	// HandleEvent receives all events and associated data read from the
	// websocket connection.
	HandleEvent(response *Response)

	// HandleUnknown receives all other socket responses.
	HandleUnknown(response *Response)

	// Listen starts the socket read loop and delivers messages to
	// HandleCommand() and HandleEvent() as appropriate.
	Listen() error

	// NextCommandID generates and returns the next command ID.
	NextCommandID() int

	// RemoveEventHandler removes a handler from the stack of listeners for an
	// event.
	RemoveEventHandler(handler EventHandler)

	// SendCommand delivers a command payload to the websocket connection.
	SendCommand(command Commander) chan *Response

	// Stop signals the socket read loop to stop listening for data and close
	// the websocket connection.
	Stop()

	// URL returns the URL of the websocket connection.
	URL() *url.URL
}
