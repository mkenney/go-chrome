package socket

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

/*
New returns a new Socketer websocket connection listening to the specified URL.
*/
func New(socketURL string) (Socketer, error) {

	dialer := &websocket.Dialer{EnableCompression: false}
	header := http.Header{"Origin": []string{socketURL}}

	webSocket, response, err := dialer.Dial(socketURL, header)
	if err != nil {
		log.Warningf("Could not create websocket connection. %s responded with '%s'", socketURL, response.Status)
		return nil, err
	}

	socket := &socket{
		conn:     webSocket,
		commands: NewCommandMap(),
		handlers: NewEventHandlerMap(),
		url:      socketURL,
	}
	go socket.Listen()

	log.Infof("New socket connection listening on %s: %s", socket.url, response.Status)
	return socket, nil
}

//////////////////////////////////////////////////
// Socketer
//////////////////////////////////////////////////

/*
socket implements Socketer.
*/
type socket struct {
	commandID int
	commands  CommandMapper
	handlers  EventHandlerMapper
	mux       sync.Mutex
	url       string
	conn      *websocket.Conn
}

/*
Close implements Socketer.
*/
func (socket *socket) Close() error {
	return socket.conn.Close()
}

/*
GenerateCommandID generates and returns a unique command ID.
*/
func (socket *socket) GenerateCommandID() int {
	socket.commandID++
	return socket.commandID
}

/*
Listen implements Socketer.
*/
func (socket *socket) Listen() {
	for {
		response := &Response{}
		err := socket.conn.ReadJSON(&response)

		if nil != err {
			log.Error(err)
			if err == io.EOF ||
				websocket.IsCloseError(err, 1006) ||
				strings.Contains(err.Error(), "closed network connection") {
				break
			}

		} else if response.ID > 0 {
			socket.HandleCommand(response)

		} else {
			socket.HandleEvent(response)
		}
	}
}

/*
Lock implements Socketer.
*/
func (socket *socket) Lock() {
	socket.mux.Lock()
}

/*
Unlock implements Socketer.
*/
func (socket *socket) Unlock() {
	socket.mux.Unlock()
}

/*
URL implements Socketer.
*/
func (socket *socket) URL() string {
	return socket.url
}

//////////////////////////////////////////////////
// Socket response data
//////////////////////////////////////////////////

/*
Error represents a socket response error.
*/
type Error struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

/*
Response represents a socket message.
*/
type Response struct {
	Error  Error           `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}
