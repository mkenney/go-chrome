package socket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
MockWebSocket implements WebSocketer for mocking
*/
type MockWebSocket struct{}

/*
NewMockWebsocket returns a mock websocket connection
*/
func NewMockWebsocket(socketURL *url.URL) (WebSocketer, error) {
	log.Infof("Mock websocket connection to %s established", socketURL.String())
	return &MockWebSocket{}, nil
}

/*
Close implements WebSocketer
*/
func (sock *MockWebSocket) Close() error {
	return nil
}

/*
ReadJSON implements WebSocketer
*/
func (sock *MockWebSocket) ReadJSON(v interface{}) error {
	var err Error
	var data []byte

	//log.Infof("MockJSONData: %s", MockJSONData)
	//log.Infof("MockJSONType: %s", MockJSONType)
	//log.Infof("MockJSONRead: %v", MockJSONRead)
	//log.Infof("MockJSONError: %v", MockJSONError)
	//log.Infof("MockJSONThrowError: %v", MockJSONThrowError)

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
			//log.Infof("Mocking Command")
			data, _ = json.Marshal(struct {
				Error  Error
				ID     int
				Method string
				Params json.RawMessage
				Result json.RawMessage
			}{
				Error:  err,
				ID:     _commandID,
				Method: "Some.method",
				Params: nil,
				Result: MockJSONData,
			})
		} else if "event" == MockJSONType {
			//log.Infof("Mocking Event")
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

// MockJSONData flag for mocking ReadJSON()
var MockJSONData []byte

// MockJSONRead flag for mocking ReadJSON()
var MockJSONRead = false

// MockJSONType flag for mocking ReadJSON()
var MockJSONType = "command"

// MockJSONError flag for mocking ReadJSON()
var MockJSONError = true

// MockJSONThrowError flag for mocking ReadJSON()
var MockJSONThrowError = false

/*
WriteJSON implements WebSocketer
*/
func (sock *MockWebSocket) WriteJSON(v interface{}) error {
	return nil
}
