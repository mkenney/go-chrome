package socket

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

/*
New returns a new websocket connection
*/
func New(socketURL string) (*Socket, error) {

	dialer := &websocket.Dialer{
		EnableCompression: false,
	}
	header := http.Header{
		"Origin": []string{socketURL},
	}

	webSocket, response, err := dialer.Dial(socketURL, header)
	if err != nil {
		log.Warningf("Could not create socket connection. %s responded with '%s'", socketURL, response.Status)
		return nil, err
	}

	socket := &Socket{
		conn:     webSocket,
		Commands: &commandMap{},
		Events:   &eventMap{},
	}
	socket.Commands.Map = make(map[int]*Command)
	socket.Events.Map = make(map[string][]*EventHandler)
	go socket.Listen()

	log.Infof("New socket connection listening on %s", socketURL)
	socket.URL = socketURL
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
	Map map[int]*Command
	Mux sync.Mutex
}

type eventMap struct {
	Map map[string][]*EventHandler
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
SocketError represents an error
*/
type SocketError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
