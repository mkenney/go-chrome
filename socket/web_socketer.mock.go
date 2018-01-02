package socket

import (
	"encoding/json"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
MockWebSocket implements WebSocketer for mocking
*/
type MockWebSocket struct {
	mockResponses []*Response
}

/*
NewMockWebsocket returns a mock websocket connection
*/
func NewMockWebsocket(socketURL *url.URL) (WebSocketer, error) {
	log.Infof("Mock websocket connection to %s established", socketURL.String())
	return &MockWebSocket{
		mockResponses: make([]*Response, 0),
	}, nil
}

/*
Close implements WebSocketer
*/
func (socket *MockWebSocket) Close() error {
	return nil
}

/*
ReadJSON implements WebSocketer

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
func (socket *MockWebSocket) ReadJSON(v interface{}) error {
	var data interface{}

	if len(socket.mockResponses) > 0 {
		time.Sleep(time.Millisecond * 100)
		data = socket.mockResponses[0]
		socket.mockResponses = socket.mockResponses[1:]

	} else {
		time.Sleep(time.Second * 1)
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
WriteJSON implements WebSocketer
*/
func (socket *MockWebSocket) WriteJSON(v interface{}) error {
	return nil
}

/*
AddMockData implements WebSocketer
*/
func (socket *MockWebSocket) AddMockData(id int, err *Error, method string, data ...interface{}) {
	response := &Response{
		Error:  err,
		ID:     id,
		Method: method,
	}
	if len(data) > 0 {
		response.Result, _ = json.Marshal(data[0])
	}
	if len(data) > 1 {
		response.Params, _ = json.Marshal(data[1])
	}
	socket.mockResponses = append(socket.mockResponses, response)
}
