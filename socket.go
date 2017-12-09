package chrome

import (
	"encoding/json"
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
		commands: make(map[int]*protocol.Command),
		events:   make(map[string][]*protocol.EventHandler),
	}
	go socket.Listen(tab)

	log.Infof("New socket connection listening on %s", tab.WebSocketDebuggerURL)
	return socket, nil
}

/*
Socket represents a websocket connection to the Browser instance
*/
type Socket struct {
	cmdID      int
	cmdMutex   sync.Mutex
	eventMutex sync.Mutex
	events     map[string][]*protocol.EventHandler // key is event name.
	commands   map[int]*protocol.Command           // key is id.
	conn       *websocket.Conn
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
func (socket *Socket) Listen(tab *Tab) {
	for {
		response := &SocketResponse{}
		err := socket.conn.ReadJSON(response)
		if err != nil {
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
			params, _ := json.Marshal(response.Params)
			log.Infof("%s: %s", tab.ID, response.Method)
			log.Debugf("Event params: %s", params)
			socket.handleEvent(response)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

///*
//SendCommand sends a command payload to the socket
//*/
//func (socket *Socket) SendCommand(command protocol.CommandIface) int {
//	command.WG.Add(1)
//	socket.cmdMutex.Lock()
//	defer socket.cmdMutex.Unlock()
//
//	socket.cmdID++
//	payload := &protocol.CommandPayload{
//		socket.cmdID,
//		command.Method,
//		command.Params,
//	}
//	tmp, _ := json.Marshal(payload)
//	log.Debugf("Sending %#v", string(tmp))
//	if err := socket.conn.WriteJSON(payload); err != nil {
//		command.Done(nil, err)
//	}
//	socket.commands[payload.ID] = command
//
//	command.WG.Wait()
//	return payload.ID
//}
//
