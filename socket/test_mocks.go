package socket

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
NewMock returns a mock websocket for unit testing
*/
func NewMock(socketURL string) (Socketer, error) {
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

	log.Infof("MockJSONData: %s", MockJSONData)
	log.Infof("MockJSONType: %s", MockJSONType)
	log.Infof("MockJSONCommandID: %d", MockJSONCommandID)
	log.Infof("MockJSONRead: %v", MockJSONRead)
	log.Infof("MockJSONError: %v", MockJSONError)
	log.Infof("MockJSONThrowError: %v", MockJSONThrowError)

	time.Sleep(time.Millisecond * 100)
	if !MockJSONRead {
		if MockJSONThrowError {
			return fmt.Errorf("Mock Read Error")
		}

		MockJSONRead = true
		if MockJSONError {
			err.Code = 1
			err.Data = []byte(`{"data": "Error data"}`)
			err.Message = "Mock Error"
		}

		if "command" == MockJSONType {
			data, _ = json.Marshal(struct {
				Error  Error
				ID     int
				Method string
				Params json.RawMessage
				Result json.RawMessage
			}{
				Error:  err,
				ID:     MockJSONCommandID,
				Method: "Some.method",
				Params: nil,
				Result: MockJSONData,
			})
		} else if "event" == MockJSONType {
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
				Result: MockJSONData,
			})
		}
		json.Unmarshal(data, &v)
	}
	return nil
}

var MockJSONData []byte
var MockJSONRead = false
var MockJSONType = "command"
var MockJSONError = true
var MockJSONCommandID = 1
var MockJSONThrowError = false

/*
WriteJSON implements Conner
*/
func (mockConn) WriteJSON(v interface{}) error {
	return nil
}
