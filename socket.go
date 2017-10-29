/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

/*
Command represents a command to send to Chrome
*/
type Command interface {
	Name() string
	Params() interface{}
	Done(result []byte, err error)
}

/*
CommandJSON is a JSON representation of a Command
*/
type CommandJSON struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

/*
Socket represents a websocket connection to the Browser instance
*/
type Socket struct {
	socket          *websocket.Conn
	commandMutex    sync.Mutex
	pendingCommands map[int]Command // key is id.
	nextCommandID   int
	eventMutex      sync.Mutex
	events          map[string][]EventInterface
}

/*
ErrorJSON represents an error
*/
type ErrorJSON struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
EventInterface is an interface
*/
type EventInterface interface {
	OnEvent(name string, params []byte)
}

/*
MessageJSON represents a message
*/
type MessageJSON struct {
	ID     int             `json:"id"`
	Error  ErrorJSON       `json:"error"`
	Result json.RawMessage `json:"result"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

/*
Return a new websocket connection
*/
func newSocket(url string) (*Socket, error) {
	log.Printf("Connecting to %s ...", url)

	dialer := &websocket.Dialer{
		EnableCompression: false,
	}
	header := http.Header{
		"Origin": []string{"http://localhost/"},
	}
	webSocket, _, err := dialer.Dial(url, header)
	if err != nil {
		return nil, err
	}

	socket := &Socket{
		socket:          webSocket,
		pendingCommands: make(map[int]Command),
		events:          make(map[string][]EventInterface),
	}
	go socket.listen()

	return socket, nil
}

/*
Close closes the current socket connection
*/
func (c *Socket) Close() error {
	return c.socket.Close()
}

/*
Send sends a command to a connected socket
*/
func (c *Socket) Send(command Command) {
	c.commandMutex.Lock()
	defer c.commandMutex.Unlock()

	c.nextCommandID++
	cj := &CommandJSON{
		ID:     c.nextCommandID,
		Method: command.Name(),
		Params: command.Params(),
	}
	log.Printf("Send %#v", cj)
	if err := c.socket.WriteJSON(cj); err != nil {
		command.Done(nil, err)
		return
	}
	c.pendingCommands[c.nextCommandID] = command
	// TODO: Implement timeout.
}

/*
AddEvent adds an event to the event stack
*/
func (c *Socket) AddEvent(name string, event EventInterface) {
	c.eventMutex.Lock()
	defer c.eventMutex.Unlock()
	events := c.events[name]
	for _, s := range events {
		if s == event {
			return
		}
	}
	c.events[name] = append(events, event)
}

/*
RemoveEvent removes an event from the stack
*/
func (c *Socket) RemoveEvent(name string, event EventInterface) {
	c.eventMutex.Lock()
	defer c.eventMutex.Unlock()
	events := c.events[name]
	for i, s := range events {
		if s == event {
			l := len(events)
			events[i] = events[l-1]
			c.events[name] = events[:l-1]
			return
		}
	}
}

/*
NewEvent creates a new event struct
*/
func NewEvent(cb func(name string, params []byte)) EventInterface {
	return &simpleEvent{cb}
}

type simpleEvent struct {
	cb func(name string, params []byte)
}

func (s *simpleEvent) OnEvent(name string, params []byte) {
	s.cb(name, params)
}

func (c *Socket) handleResponse(mj *MessageJSON) {
	id := mj.ID
	errStr := mj.Error.Message
	result := []byte(mj.Result)

	log.Printf("handleResponse %d %s %s", id, string(result), errStr)
	c.commandMutex.Lock()
	defer c.commandMutex.Unlock()

	if cmd, ok := c.pendingCommands[id]; !ok {
		log.Printf("Unknown command %d: result=%s err=%s", id, string(result), errStr)
	} else {
		delete(c.pendingCommands, id)
		var err error
		if errStr != "" {
			err = errors.New(errStr)
		}
		go cmd.Done(result, err)
	}
}

func (c *Socket) handleEvent(mj *MessageJSON) {
	name := mj.Method
	params := []byte(mj.Params)

	log.Printf("handleEvent %s %s", name, string(params))
	if name == "Inspector.targetCrashed" {
		log.Fatal("Chrome has crashed!")
	}
	c.eventMutex.Lock()
	defer c.eventMutex.Unlock()

	events := c.events[name]
	for _, event := range events {
		go event.OnEvent(name, params)
	}
}

func (c *Socket) listen() {
	for {
		mj := &MessageJSON{}
		if err := c.socket.ReadJSON(mj); err != nil {
			if err == io.EOF || websocket.IsCloseError(err, 1006) ||
				strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			log.Printf("%v", err)
		} else if mj.ID > 0 {
			c.handleResponse(mj)
		} else {
			c.handleEvent(mj)
		}
	}
}
