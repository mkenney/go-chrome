package socket

import (
	"sync"
)

/*
Commander defines an interface for websocket commands.
*/
type Commander interface {
	// Done manages the result of a command and signals completion.
	Done(result []byte, err error)

	// Error returns the current error, if any.
	Error() error

	// ID returns the command ID
	ID() int

	// Method returns the name of the protocol method to be executed.
	Method() string

	// Params returns the method parameters.
	Params() interface{}

	// Result returns the result of the method call.
	Result() interface{}

	// WaitGroup returns the sync.WaitGroup instance.
	WaitGroup() *sync.WaitGroup
}

/*
CommandMapper defines the interface for the event handler stack.
*/
type CommandMapper interface {
	// Delete removes the entire stack of handlers for an event.
	Delete(commandID int)

	// Get retrieves the entire stack of handlers for an event.
	Get(commandID int) (Commander, error)

	// Lock locks the sync mutex.
	Lock()

	// Unlock unlocks the sync mutex.
	Unlock()

	// Set sets the entire stack of handlers for an event.
	Set(commandID int, command Commander)
}

/*
Conner defines the websocket.Conn methods used by the API

Primarily used for mocking
*/
type Conner interface {
	Close() error
	ReadJSON(v interface{}) error
	WriteJSON(v interface{}) error
}

/*
EventHandler is the interface definition for a socket event.
*/
type EventHandler interface {
	// Name returns the name of the event the handler is assigned to.
	Name() string

	// Handle executes the event handler callback.
	Handle(response *Response)
}

/*
EventHandlerMapper defines the interface for the event handler stack.
*/
type EventHandlerMapper interface {
	// Add adds a handler to the stack of handlers for an event.
	Add(handler EventHandler)

	// Delete removes the entire stack of handlers for an event.
	Delete(eventName string)

	// Get retrieves the entire stack of handlers for an event.
	Get(eventName string) ([]EventHandler, error)

	// Lock locks the sync mutex.
	Lock()

	// Unlock unlocks the sync mutex.
	Unlock()

	// Remove removes a handler from the stack of handlers for an event.
	Remove(handler EventHandler)

	// Set sets the entire stack of handlers for an event.
	Set(eventName string, handlers []EventHandler)
}

/*
Socketer defines a websocket connection interface for Chrome instances.
*/
type Socketer interface {
	// AddEventHandler adds an event handler to the stack of listeners for an event.
	AddEventHandler(handler EventHandler)

	// Close closes the current socket connection.
	Close() error

	// Conn returns the websocket connection
	Conn() Conner

	// HandleCommand handles all responses to commands sent to the websocket connection.
	HandleCommand(response *Response)

	// HandlEvent handles all events delivered to the websocket listener.
	HandleEvent(response *Response)

	// Listen starts the socket read loop.
	Listen() error

	// RemoveEventHandler removes a handler from the stack of listeners for an event.
	RemoveEventHandler(handler EventHandler)

	// SendCommand sends a command to the websocket connection.
	SendCommand(command Commander) *commandPayload

	// Stop tells the socket connection to stop listening for data
	Stop()

	// URL returns the URL of the websocket connection
	URL() string
}
