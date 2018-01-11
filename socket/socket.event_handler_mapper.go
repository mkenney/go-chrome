package socket

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
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
eventHandlerMap defines the event handler stacks for all handled events.
*/
type EventHandlerMap struct {
	stack map[string][]EventHandler
	mux   *sync.Mutex
}

/*
Add implements EventHandlerMapper.
*/
func (stack *EventHandlerMap) Add(
	handler EventHandler,
) {
	stack.Lock()
	defer stack.Unlock()

	handlers, err := stack.Get(handler.Name())
	if nil != err {
		handlers = make([]EventHandler, 0)
	}
	for _, hndl := range handlers {
		if hndl == handler {
			log.Warnf("Attempted to add a duplicate handler for event '%s', skipping...", handler.Name())
			return
		}
	}

	log.Debugf("Adding handler for event '%s'", handler.Name())
	handlers = append(handlers, handler)
	stack.Set(handler.Name(), handlers)
}

/*
Delete implements EventHandlerMapper.
*/
func (stack *EventHandlerMap) Delete(
	name string,
) {
	delete(stack.stack, name)
}

/*
Get implements EventHandlerMapper.
*/
func (stack *EventHandlerMap) Get(
	name string,
) ([]EventHandler, error) {
	if handlers, ok := stack.stack[name]; ok {
		return handlers, nil
	}
	return nil, fmt.Errorf("No event listeners found for %s", name)
}

/*
Lock implements EventHandlerMapper.
*/
func (stack *EventHandlerMap) Lock() {
	stack.mux.Lock()
}

/*
Remove implements EventHandlerMapper.
*/
func (stack *EventHandlerMap) Remove(
	handler EventHandler,
) {
	stack.Lock()
	defer stack.Unlock()

	for k, hndl := range stack.stack[handler.Name()] {
		if hndl == handler {
			stack.stack[handler.Name()] = append(stack.stack[handler.Name()][:k], stack.stack[handler.Name()][k+1:]...)
			return
		}
	}
	log.Warnf("Could not remove handler for '%s': not found", handler.Name())
}

/*
Set implements EventHandlerMapper.
*/
func (stack *EventHandlerMap) Set(
	eventName string,
	handlers []EventHandler,
) {
	stack.stack[eventName] = handlers
}

/*
Unlock implements EventHandlerMapper.
*/
func (stack *EventHandlerMap) Unlock() {
	stack.mux.Unlock()
}
