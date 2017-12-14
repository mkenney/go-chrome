package chrome

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"app/chrome/protocol"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

/*
NewSocket returns a new websocket connection
*/
func NewSocket(tab *Tab) (*Socket, error) {

	dialer := &websocket.Dialer{
		EnableCompression: false,
	}
	header := http.Header{
		"Origin": []string{tab.WebSocketDebuggerURL},
	}

	webSocket, response, err := dialer.Dial(tab.WebSocketDebuggerURL, header)
	if err != nil {
		log.Warningf("Could not create socket connection. %s responded with '%s'", tab.WebSocketDebuggerURL, response.Status)
		return nil, err
	}

	socket := &Socket{
		conn:     webSocket,
		Commands: &commandMap{},
		Events:   &eventMap{},
	}
	socket.Commands.Map = make(map[int]*protocol.Command)
	socket.Events.Map = make(map[string][]*protocol.EventHandler)
	go socket.Listen()

	log.Infof("New socket connection listening on %s", tab.WebSocketDebuggerURL)
	socket.URL = tab.WebSocketDebuggerURL
	return socket, nil
}

/*
Socket represents a websocket connection to the Browser instance
*/
type Socket struct {
	CommandID int
	Commands  *commandMap
	Events    *eventMap
	Mux       sync.Mutex
	URL       string
	conn      *websocket.Conn
}

type commandMap struct {
	Map map[int]*protocol.Command
	Mux sync.Mutex
}

type eventMap struct {
	Map map[string][]*protocol.EventHandler
	Mux sync.Mutex
}

/*
SocketResponse represents a socket message
*/
type SocketResponse struct {
	Error  SocketError     `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}

/*
Close closes the current socket connection
*/
func (socket *Socket) Close() error {
	return socket.conn.Close()
}

/*
Listen starts the socket read loop
*/
func (socket *Socket) Listen() {
	for {
		response := &SocketResponse{}
		err := socket.conn.ReadJSON(&response)

		if nil != err {
			log.Error(err)
			if err == io.EOF ||
				websocket.IsCloseError(err, 1006) ||
				strings.Contains(err.Error(), "use of closed network connection") {
				log.Error(err)
				break
			}

		} else if response.ID > 0 {
			socket.handleCommand(response)

		} else {
			socket.handleEvent(response)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

/*
SendCommand sends a command to a connected socket.

The socket's command mutex is locked, the command counter is incremented, the payload is sent to the
socket connection and the mutex is unlocked. The command is stored using the counter value as it's
Id. When the command is executed and the socket responds, handleCmd() is executed to generate a
response.
*/
func (socket *Socket) SendCommand(command *protocol.Command) int {

	// Safely get generate a command ID
	socket.Mux.Lock()
	socket.CommandID++
	commandID := socket.CommandID
	socket.Mux.Unlock()

	// Safely add a command to the internal stack
	socket.Commands.Mux.Lock()
	payload := &protocol.CommandPayload{
		ID:     commandID,
		Method: command.Method,
		Params: command.Params,
	}
	socket.Commands.Map[payload.ID] = command
	socket.Commands.Mux.Unlock() // Do not defer, handlCommand also locks

	tmp, _ := json.Marshal(payload)
	log.Debugf("Sending command '%s'", string(tmp))

	command.WG.Add(1)
	if err := socket.conn.WriteJSON(payload); err != nil {
		command.Done(nil, err)
	}
	command.WG.Wait()

	return payload.ID
}

func (socket *Socket) handleCommand(response *SocketResponse) {
	var err error

	socket.Commands.Mux.Lock()
	defer socket.Commands.Mux.Unlock()

	if command, ok := socket.Commands.Map[response.ID]; !ok {
		log.Errorf("Command %d not found: result=%s err=%s", response.ID, response.Result, response.Error.Message)

	} else {
		delete(socket.Commands.Map, response.ID)
		if "" != response.Error.Message {
			err = errors.New(response.Error.Message)
		}
		command.Done(response.Result, err)
	}
}
