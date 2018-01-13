package socket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

/*
ChromeWebSocket represents a WebSocketer interface
*/
type ChromeWebSocket struct {
	conn          *websocket.Conn
	mockResponses []*Response
}

/*
NewWebsocket returns a connected socket connection that implements the
WebSocketer interface.
*/
func NewWebsocket(socketURL *url.URL) (*ChromeWebSocket, error) {
	dialer := &websocket.Dialer{EnableCompression: false}
	header := http.Header{"Origin": []string{socketURL.String()}}

	websocket, response, err := dialer.Dial(socketURL.String(), header)
	if err != nil {
		return nil, fmt.Errorf(
			"Could not create websocket connection. %s responded with '%s'",
			socketURL.String(),
			err.Error(),
		)
	}
	log.Infof("Websocket connection to %s established: %s", socketURL.String(), response.Status)

	return &ChromeWebSocket{conn: websocket}, nil
}

/*
Close implements WebSocketer.
*/
func (socket *ChromeWebSocket) Close() error {
	if nil != socket.conn {
		return socket.conn.Close()
	}
	return nil
}

/*
ReadJSON implements WebSocketer.

To support tests, when a live web socket connection does not exist this method
reads from a stack of manually populated responses in an attempt to emulate the
Chromium DevProtocol behavior. To populate the mock response stack, add a
Response{} pointer with the AddMockData() method.
*/
func (socket *ChromeWebSocket) ReadJSON(v interface{}) error {
	if nil != socket.conn {
		return socket.conn.ReadJSON(&v)
	}

	var data interface{}

	time.Sleep(time.Millisecond * 10)

	if len(socket.mockResponses) > 0 {
		data = socket.mockResponses[0]
		socket.mockResponses = socket.mockResponses[1:]
	} else {
		data = &Response{
			Error:  &Error{},
			ID:     0,
			Method: "Unknown.event",
		}
	}

	jsonBytes, err := json.Marshal(data)
	log.Debugf("ReadJSON(): returning mock data %s", jsonBytes)
	err = json.Unmarshal(jsonBytes, &v)
	return err
}

/*
WriteJSON implements WebSocketer.
*/
func (socket *ChromeWebSocket) WriteJSON(v interface{}) error {
	if nil != socket.conn {
		return socket.conn.WriteJSON(v)
	}
	return nil
}

/*
AddMockData implements WebSocketer.

This method is only indended for mocking data for unit tests and will panic if
used on a live websocket connection.
*/
func (socket *ChromeWebSocket) AddMockData(response *Response) {
	if nil == socket.conn {
		socket.mockResponses = append(socket.mockResponses, response)
	} else {
		panic("AddMockData cannot be called on a live websocket connection and is intended for mocking data for unit testing only.")
	}
}
