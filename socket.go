/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

/*
Command represents a command to send to Chrome
*/
type Command interface {
	Done(result []byte, err error)
	Name() string
	Params() interface{}
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
	commandMutex    sync.Mutex
	eventMutex      sync.Mutex
	events          map[string][]EventInterface
	nextCommandID   int
	pendingCommands map[int]Command // key is id.
	socket          *websocket.Conn
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
	Error  ErrorJSON       `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}

type simpleEvent struct {
	cb func(name string, params []byte)
}

/*
Return a new websocket connection
*/
func newSocket(url string) (*Socket, error) {

	dialer := &websocket.Dialer{
		EnableCompression: false,
	}
	header := http.Header{
		"Origin": []string{url},
	}

	webSocket, response, err := dialer.Dial(url, header)
	if err != nil {
		log.Warningf("Could not create socket connection. %s responded with '%s'", url, response.Status)
		return nil, err
	}

	socket := &Socket{
		socket:          webSocket,
		pendingCommands: make(map[int]Command),
		events:          make(map[string][]EventInterface),
	}
	go socket.listen()

	log.Infof("Opened new socket connection listening to %s", url)
	return socket, nil
}

/*
Close closes the current socket connection
*/
func (socket *Socket) Close() error {
	return socket.socket.Close()
}

/*
Send sends a command to a connected socket
*/
func (socket *Socket) Send(command Command) {
	socket.commandMutex.Lock()
	defer socket.commandMutex.Unlock()

	socket.nextCommandID++
	cj := &CommandJSON{
		ID:     socket.nextCommandID,
		Method: command.Name(),
		Params: command.Params(),
	}
	log.Infof("Send %#v", cj)
	if err := socket.socket.WriteJSON(cj); err != nil {
		command.Done(nil, err)
		return
	}
	socket.pendingCommands[socket.nextCommandID] = command
	// TODO: Implement timeout.
}

/*
AddEvent adds an event to the event stack
*/
func (socket *Socket) AddEvent(name string, event EventInterface) {
	socket.eventMutex.Lock()
	defer socket.eventMutex.Unlock()
	events := socket.events[name]
	for _, s := range events {
		if s == event {
			return
		}
	}
	socket.events[name] = append(events, event)
}

/*
RemoveEvent removes an event from the stack
*/
func (socket *Socket) RemoveEvent(name string, event EventInterface) {
	socket.eventMutex.Lock()
	defer socket.eventMutex.Unlock()
	events := socket.events[name]
	for i, s := range events {
		if s == event {
			l := len(events)
			events[i] = events[l-1]
			socket.events[name] = events[:l-1]
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

func (s *simpleEvent) OnEvent(name string, params []byte) {
	s.cb(name, params)
}

func (socket *Socket) handleResponse(mj *MessageJSON) {
	id := mj.ID
	errStr := mj.Error.Message
	result := []byte(mj.Result)

	log.Infof("handleResponse %d %s %s", id, string(result), errStr)
	socket.commandMutex.Lock()
	defer socket.commandMutex.Unlock()

	if cmd, ok := socket.pendingCommands[id]; !ok {
		log.Infof("Unknown command %d: result=%s err=%s", id, string(result), errStr)
	} else {
		delete(socket.pendingCommands, id)
		var err error
		if errStr != "" {
			err = errors.New(errStr)
		}
		go cmd.Done(result, err)
	}
}

func (socket *Socket) handleEvent(mj *MessageJSON) {
	name := mj.Method
	params := []byte(mj.Params)

	log.Infof("handleEvent %s %s", name, string(params))
	if name == "Inspector.targetCrashed" {
		log.Fatalf("Chrome has crashed!")
	}
	socket.eventMutex.Lock()
	defer socket.eventMutex.Unlock()

	events := socket.events[name]
	for _, event := range events {
		go event.OnEvent(name, params)
	}
}

func (socket *Socket) listen() {
	for {
		mj := &MessageJSON{}
		if err := socket.socket.ReadJSON(mj); err != nil {
			if err == io.EOF || websocket.IsCloseError(err, 1006) ||
				strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			log.Errorf("%v", err)
		} else if mj.ID > 0 {
			socket.handleResponse(mj)
		} else {
			socket.handleEvent(mj)
		}
	}
}
