package socket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

/*
NewWebsocket returns a connected socket connection that implements the
WebSocketer interface.
*/
func NewWebsocket(socketURL *url.URL) (*ChromeWebSocket, error) {
	dialer := &websocket.Dialer{EnableCompression: false}
	header := http.Header{"Origin": []string{socketURL.String()}}

	websocket, response, err := dialer.Dial(socketURL.String(), header)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"%s websocket connection failed",
			socketURL.String(),
		))
	}
	log.Infof("Websocket connection to %s established: %s", socketURL.String(), response.Status)

	return &ChromeWebSocket{conn: websocket}, nil
}

/*
ChromeWebSocket provides an WebSocketer interface for managing a websocket
connection.

ChromeWebSocket represents a WebSocketer interface
*/
type ChromeWebSocket struct {
	conn          *websocket.Conn
	mockResponses []*Response
}

/*
Close closes the current websocket connection.

Close is a WebSocketer implementation.
*/
func (socket *ChromeWebSocket) Close() error {
	if nil != socket.conn {
		return socket.conn.Close()
	}
	return nil
}

/*
ReadJSON listens for the next websocket message and unmarshalls it into the
provided variable.

ReadJSON is a WebSocketer implementation.

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
	return errors.Wrap(err, fmt.Sprintf("could not unmarshal %s", jsonBytes))
}

/*
WriteJSON marshalls the provided data as JSON and writes it to the websocket.

WriteJSON is a WebSocketer implementation.
*/
func (socket *ChromeWebSocket) WriteJSON(v interface{}) error {
	if nil != socket.conn {
		return socket.conn.WriteJSON(v)
	}
	return nil
}

/*
AddMockData is a WebSocketer implementation.

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
