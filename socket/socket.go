package socket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	chrome_error "github.com/mkenney/go-chrome/error"
	log "github.com/sirupsen/logrus"
)

/*
New returns a new Socketer websocket connection listening to the specified URL.
*/
func New(socketURL string) (Socketer, *chrome_error.Error) {

	dialer := &websocket.Dialer{EnableCompression: false}
	header := http.Header{"Origin": []string{socketURL}}

	webSocket, response, err := dialer.Dial(socketURL, header)
	if err != nil {
		return nil, chrome_error.New(
			fmt.Sprintf("Could not create websocket connection. %s responded with '%s'", socketURL, response.Status),
			chrome_error.LevelWarn,
			chrome_error.CodeCouldNotConnectToSocket,
			err,
		)
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
	commandID     int
	commands      CommandMapper
	handlers      EventHandlerMapper
	url           string
	conn          Conner
	stopListening bool
}

/*
Close implements Socketer.
*/
func (socket *socket) Close() error {
	return socket.conn.Close()
}

/*
Conn implements Socketer.
*/
func (socket *socket) Conn() Conner {
	return socket.conn
}

/*
Listen implements Socketer.
*/
func (socket *socket) Listen() (err error) {
	socket.stopListening = false
	for {
		if socket.stopListening {
			break
		}

		response := &Response{}
		err = socket.conn.ReadJSON(&response)
		if nil != err {
			log.Error(err)
			socket.Stop() // This will end the loop after handling the current response (if any)
		}

		if response.ID > 0 {
			socket.HandleCommand(response)

		} else {
			socket.HandleEvent(response)
		}
	}
	log.Infof("Socket shutting down")

	return
}

/*
Stop implements Socketer.
*/
func (socket *socket) Stop() {
	socket.stopListening = true
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
	Error  *Error          `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}
