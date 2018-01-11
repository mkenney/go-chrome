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
NewWebsocket returns a connected socket connection
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

This uses a stack of responses to attempt to emulate Chromium behavior for
testsing. To use, add a response to the stack with addMockWebsocketResponse().

There is a potential timing issue when emulating commands. "Commands" are
structs that implement the Commander interface and are a type of event handler
that makes a request to the socket and only handles responses to that request.

Due to the nature of the socket read loop, which is a stream, Commander objects
use `sync.WaitGroup` to wait for the socket response to the request that was
submitted. Because of this, your mock data must be added to the response stack
BEFORE the Commander handler is executed or the test routine will lock forever,
preventing you from adding your mock data.

That means that it's possible for the socket read loop to receive your mock
response before the Commander handler command is registered with the socket read
loop. If that happens and a handler isn't present to receive it then the
response is discarded, and then when the Commander handler is executed the
routine will lock forever waiting for a response that will never come.

As a workaround, each response is set to take 1 second to return, giving command
registration some time to complete. Alternatively, you could launch the command
registration in a goroutine and then add the mock response data to the stack.
I'll refactor it that way when the tests start taking too long.

This is only a problem with mocking the socket stream data for unit tests. It
does impact interacting Chrome in any way.
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
*/
func (socket *ChromeWebSocket) AddMockData(response *Response) {
	if nil == socket.conn {
		socket.mockResponses = append(socket.mockResponses, response)
	}
}
