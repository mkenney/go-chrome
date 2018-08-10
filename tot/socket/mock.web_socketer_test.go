package socket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
)

/*
NewMockWebsocket returns a WebSocketer mock for tests that have socket
connection dependencies.
*/
func NewMockWebsocket(socketURL *url.URL) (WebSocketer, error) {
	log.Infof("Mock websocket connection to %s established", socketURL.String())
	return &MockChromeWebSocket{
		mockResponses: make([]*Response, 0),
	}, nil
}

type MockChromeWebSocket struct {
	mockResponses []*Response
	sleep         time.Duration
}

func (socket *MockChromeWebSocket) Close() error {
	socket.mockResponses = []*Response{{}, {}, {}, {}, {}}
	return nil
}

/*
This method populates a queue of mock data that will be delivered to the
websocket API for testing.
*/
func (socket *MockChromeWebSocket) AddMockData(response *Response) {
	socket.mockResponses = append(socket.mockResponses, response)
}

/*
ReadJSON reads from a stack of manually populated responses in an
attempt to emulate the Chromium DevProtocol behavior. To populate the
mock response stack, add a Response{} pointer with the AddMockData()
method.

ReadJSON is a WebSocketer implementation.
*/
func (socket *MockChromeWebSocket) ReadJSON(v interface{}) error {
	var data interface{}
	time.Sleep(time.Millisecond * 10)

	if socket.sleep > 0 {
		time.Sleep(socket.sleep)
		socket.sleep = 0
	}

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

	jsonBytes, _ := json.Marshal(data)
	log.Debugf("Mock ReadJSON(): returning mock data %s", jsonBytes)
	err = json.Unmarshal(jsonBytes, &v)
	if nil != err {
		return errs.Wrap(err, 0, fmt.Sprintf("could not unmarshal %s", jsonBytes))
	}
	return nil
}

/*
Sleep sets the sleep duration for the next ReadJSON call to replicate
timeouts and delays.
*/
func (socket *MockChromeWebSocket) Sleep(duration time.Duration) {
	socket.sleep = duration
}

/*
WriteJSON marshalls the provided data as JSON and writes it to the websocket.

WriteJSON is a WebSocketer implementation.
*/
func (socket *MockChromeWebSocket) WriteJSON(v interface{}) error {
	return nil
}
