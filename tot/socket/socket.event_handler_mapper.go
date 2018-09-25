package socket

import (
	"fmt"
	"sync"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
	"github.com/mkenney/go-chrome/codes"
)

/*
NewEventHandlerMap creates and returns a pointer to an EventHandlerMapper.
*/
func NewEventHandlerMap() *EventHandlerMap {
	return &EventHandlerMap{
		stack: make(map[string][]EventHandler),
		mux:   &sync.Mutex{},
	}
}

/*
EventHandlerMap provides an EventHandlerMapper interface for handling the event
handler stack.
*/
type EventHandlerMap struct {
	stack map[string][]EventHandler
	mux   *sync.Mutex
}

/*
Add adds a handler to the stack of handlers for an event.

Add is an EventHandlerMapper implementation.
*/
func (stack *EventHandlerMap) Add(
	handler EventHandler,
) error {
	stack.Lock()
	defer stack.Unlock()

	handlers, err := stack.Get(handler.Name())
	if nil != err {
		handlers = make([]EventHandler, 0)
	}
	for _, hndl := range handlers {
		if hndl == handler {
			return errs.New(codes.SocketDuplicateEventHandler, fmt.Sprintf("Attempted to add a duplicate handler for event '%s'", handler.Name()))
		}
	}

	log.WithFields(log.Fields{"event": handler.Name()}).
		Debug("Adding event handler")
	handlers = append(handlers, handler)
	stack.Set(handler.Name(), handlers)
	return nil
}

/*
Delete removes the entire stack of handlers for an event.

Delete is an EventHandlerMapper implementation.
*/
func (stack *EventHandlerMap) Delete(
	name string,
) {
	delete(stack.stack, name)
}

/*
Get retrieves the entire stack of handlers for an event.

Get is an EventHandlerMapper implementation.
*/
func (stack *EventHandlerMap) Get(
	name string,
) ([]EventHandler, error) {
	if handlers, ok := stack.stack[name]; ok {
		return handlers, nil
	}
	return nil, errs.New(codes.SocketDuplicateEventHandler, fmt.Sprintf("No event listeners found for %s", name))
}

/*
Lock locks the sync mutex.

Lock is an EventHandlerMapper implementation.
*/
func (stack *EventHandlerMap) Lock() {
	stack.mux.Lock()
}

/*
Remove removes a handler from the stack of handlers for an event.

Remove is an EventHandlerMapper implementation.
*/
func (stack *EventHandlerMap) Remove(
	handler EventHandler,
) error {
	stack.Lock()
	defer stack.Unlock()

	for k, hndl := range stack.stack[handler.Name()] {
		if hndl == handler {
			stack.stack[handler.Name()] = append(stack.stack[handler.Name()][:k], stack.stack[handler.Name()][k+1:]...)
			return nil
		}
	}
	return errs.New(codes.SocketEventHandlerNotFound, fmt.Sprintf("Could not remove handler for '%s': not found", handler.Name()))
}

/*
Set sets the entire stack of handlers for an event.

Set is an EventHandlerMapper implementation.
*/
func (stack *EventHandlerMap) Set(
	eventName string,
	handlers []EventHandler,
) {
	stack.stack[eventName] = handlers
}

/*
Unlock unlocks the sync mutex.

Unlock is an EventHandlerMapper implementation.
*/
func (stack *EventHandlerMap) Unlock() {
	stack.mux.Unlock()
}
