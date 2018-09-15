package socket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
	"github.com/gorilla/websocket"
)

/*
NewWebsocket returns a connected socket connection that implements the
WebSocketer interface.
*/
func NewWebsocket(socketURL *url.URL) (WebSocketer, error) {
	dialer := &websocket.Dialer{
		EnableCompression: true,
		// See: https://github.com/gorilla/websocket/issues/245
		// Chrome does not support socket fragmentation: https://chromium.googlesource.com/chromium/src/+/master/net/server/web_socket_encoder.cc#85
		// Chrome does not support payloads larger than 1MB: https://chromium.googlesource.com/chromium/src/+/master/net/server/http_connection.h#33
		WriteBufferSize: 1 * 1024 * 1024,
	}

	websocket, response, err := dialer.Dial(socketURL.String(), nil)
	if err != nil {
		return nil, errs.Wrap(err, 0, fmt.Sprintf(
			"%s websocket connection failed",
			socketURL.String(),
		))
	}
	log.WithFields(log.Fields{
		"status": response.Status,
		"url":    socketURL.String(),
	}).Info("Websocket connection established")

	return &ChromeWebSocket{conn: websocket, mux: &sync.Mutex{}}, nil
}

/*
ChromeWebSocket provides an WebSocketer interface for managing a websocket
connection.

ChromeWebSocket represents a WebSocketer interface
*/
type ChromeWebSocket struct {
	conn          *websocket.Conn
	mockResponses []*Response
	mux           *sync.Mutex
}

/*
Close closes the current websocket connection.

Close is a WebSocketer implementation.
*/
func (websocket *ChromeWebSocket) Close() error {
	websocket.mux.Lock()
	defer websocket.mux.Unlock()
	if nil != websocket.conn {
		return websocket.conn.Close()
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
func (websocket *ChromeWebSocket) ReadJSON(v interface{}) error {
	websocket.mux.Lock()
	if nil == websocket.conn {
		websocket.mux.Unlock()
		return errs.New(0, "not connected")
	}
	return websocket.conn.ReadJSON(&v)
}

/*
WriteJSON marshalls the provided data as JSON and writes it to the websocket.

WriteJSON is a WebSocketer implementation.
*/
func (websocket *ChromeWebSocket) WriteJSON(v interface{}) error {
	websocket.mux.Lock()
	defer websocket.mux.Unlock()
	if nil == websocket.conn {
		return errs.New(0, "not connected")
	}
	tmp, _ := json.Marshal(v)
	if len(tmp) > 1*1024*1024 {
		return fmt.Errorf("payload too large. chrome supports a maximum payload size of 1MB. See https://github.com/gorilla/websocket/issues/245")
	}
	return websocket.conn.WriteJSON(v)
}
