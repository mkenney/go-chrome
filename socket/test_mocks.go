package socket

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
newMock returns a mock websocket for unit testing
*/
func newMock(socketURL string) (Socketer, error) {
	return &socket{
		conn:     &mockConn{},
		commands: NewCommandMap(),
		handlers: NewEventHandlerMap(),
		url:      socketURL,
	}, nil
}

/*
mockConn implements Conner for mocking
*/
type mockConn struct{}

/*
Close implements Conner
*/
func (mockConn) Close() error {
	return nil
}

/*
ReadJSON implements Conner
*/
func (mockConn) ReadJSON(v interface{}) error {
	var err Error
	var data []byte

	log.Infof("mockJSONType: %s", mockJSONType)
	log.Infof("mockJSONCommandID: %d", mockJSONCommandID)
	log.Infof("mockJSONRead: %v", mockJSONRead)
	log.Infof("mockJSONError: %v", mockJSONError)
	log.Infof("mockJSONThrowError: %v", mockJSONThrowError)

	time.Sleep(time.Millisecond * 100)
	if !mockJSONRead {
		if mockJSONThrowError {
			return fmt.Errorf("Mock Read Error")
		}

		mockJSONRead = true
		if mockJSONError {
			err.Code = 1
			err.Data = []byte(`{"data": "Error data"}`)
			err.Message = "Mock Error"
		}

		if "command" == mockJSONType {
			data, _ = json.Marshal(struct {
				Error  Error
				ID     int
				Method string
				Params json.RawMessage
				Result json.RawMessage
			}{
				Error:  err,
				ID:     mockJSONCommandID,
				Method: "Some.method",
				Params: nil,
				Result: []byte(`"Mock Command Result"`),
			})
		} else if "event" == mockJSONType {
			data, _ = json.Marshal(struct {
				Error  Error
				ID     int
				Method string
				Params json.RawMessage
				Result json.RawMessage
			}{
				Error:  err,
				ID:     0,
				Method: "Some.event",
				Params: nil,
				Result: []byte(`"Mock Event Result"`),
			})
		}
		json.Unmarshal(data, &v)
	}
	return nil
}

var mockJSONRead = false
var mockJSONType = "command"
var mockJSONError = true
var mockJSONCommandID = 1
var mockJSONThrowError = false

/*
WriteJSON implements Conner
*/
func (mockConn) WriteJSON(v interface{}) error {
	return nil
}
