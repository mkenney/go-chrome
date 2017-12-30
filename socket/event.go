package socket

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

/*
NewEventHandler returns a pointer to an event handler.
*/
func NewEventHandler(
	name string,
	callback func(response *Response),
) EventHandler {
	return &handler{
		callback: callback,
		name:     name,
	}
}

/*
NewEventHandlerMap creates and returns a pointer to an EventHandlerMapper.
*/
func NewEventHandlerMap() EventHandlerMapper {
	return &eventHandlerMap{
		stack: make(map[string][]EventHandler),
		mux:   &sync.Mutex{},
	}
}

//////////////////////////////////////////////////
// Event handler map
//////////////////////////////////////////////////

/*
eventHandlerMap defines the event handler stacks for all handled events.
*/
type eventHandlerMap struct {
	stack map[string][]EventHandler
	mux   *sync.Mutex
}

/*
Add implements EventHandlerMapper.
*/
func (stack *eventHandlerMap) Add(
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
			log.Warnf("Attempted to add a duplicate handler for event '%s'", handler.Name())
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
func (stack *eventHandlerMap) Delete(
	name string,
) {
	delete(stack.stack, name)
}

/*
Get implements EventHandlerMapper.
*/
func (stack *eventHandlerMap) Get(
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
func (stack *eventHandlerMap) Lock() {
	stack.mux.Lock()
}

/*
Remove implements EventHandlerMapper.
*/
func (stack *eventHandlerMap) Remove(
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
func (stack *eventHandlerMap) Set(
	eventName string,
	handlers []EventHandler,
) {
	stack.stack[eventName] = handlers
}

/*
Unlock implements EventHandlerMapper.
*/
func (stack *eventHandlerMap) Unlock() {
	stack.mux.Unlock()
}

//////////////////////////////////////////////////
// Event handler
//////////////////////////////////////////////////

/*
handler implements EventHandler.
*/
type handler struct {
	callback func(response *Response)
	name     string
}

/*
Name implements EventHandler.
*/
func (handler *handler) Name() string {
	return handler.name
}

/*
Handle implements EventHandler.
*/
func (handler *handler) Handle(
	response *Response,
) {
	handler.callback(response)
}

//////////////////////////////////////////////////
// Socketer
//////////////////////////////////////////////////

/*
AddEventHandler implements Socketer.
*/
func (socket *socket) AddEventHandler(
	handler EventHandler,
) {
	socket.handlers.Lock()
	defer socket.handlers.Unlock()

	handlers, err := socket.handlers.Get(handler.Name())
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
	socket.handlers.Set(handler.Name(), handlers)
}

/*
HandleEvent implements Socketer.
*/
func (socket *socket) HandleEvent(
	response *Response,
) {
	log.Debugf("%s event received from %s", response.Method, socket.URL())
	if response.Method == "Inspector.targetCrashed" {
		log.Fatalf("Chrome has crashed!")
	}
	socket.handlers.Lock()
	defer socket.handlers.Unlock()

	if handlers, err := socket.handlers.Get(response.Method); nil != err {
		log.Debugf("Could not handle event %s: %s", response.Method, err.Error())

	} else {
		for a, event := range handlers {
			log.Infof("Executing handler #%d for event %s", a, response.Method)
			go event.Handle(response)
		}
	}
}

/*
RemoveEventHandler implements Socketer.
*/
func (socket *socket) RemoveEventHandler(
	handler EventHandler,
) {
	socket.handlers.Lock()
	defer socket.handlers.Unlock()

	if handlers, err := socket.handlers.Get(handler.Name()); nil != err {
		log.Warnf("Could not remove handler: %s", err.Error())
	} else {
		for i, hndlr := range handlers {
			if hndlr == handler {
				handlers = append(handlers[:i], handlers[i+1:]...)
				socket.handlers.Set(handler.Name(), handlers)
				return
			}
		}
	}
}
