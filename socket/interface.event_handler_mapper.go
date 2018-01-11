package socket

/*
EventHandlerMapper defines a management interface for the stack of event
handlers.
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

	// Remove removes a handler from the stack of handlers for an event.
	Remove(handler EventHandler)

	// Set sets the entire stack of handlers for an event.
	Set(eventName string, handlers []EventHandler)

	// Unlock unlocks the sync mutex.
	Unlock()
}
