package socket

import (
	"encoding/json"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestSocketClose(t *testing.T) {
	mockSocket, _ := NewMock("https://www.example.com/")
	if err := mockSocket.Close(); nil != err && "*errors.errorString" != reflect.TypeOf(err).String() {
		t.Errorf("Socketer.Close() must return an error or nil, %s found", reflect.TypeOf(err).String())
	}
}

func TestSocketConn(t *testing.T) {
	mockSocket, _ := NewMock("https://www.example.com/")
	conn := mockSocket.Conn()
	if _, ok := interface{}(conn).(Conner); !ok {
		t.Errorf("Socketer.Conn() must return a Conner interface")
	}
}

func TestGenerateCommandID(t *testing.T) {
	for a := 1; a <= 10; a++ {
		if b := GenerateCommandID(); a != b {
			t.Errorf("GenerateCommandID() failed to generate predictable IDs: %d expected, %d received", a, b)
		}
	}
}

func TestListenCommand(t *testing.T) {
	MockJSONData = []byte(`"Mock Command Result"`)
	MockJSONRead = false
	MockJSONType = "command"
	MockJSONError = false
	MockJSONThrowError = false

	command := NewCommand("Some.method", nil)
	log.Infof("TestListenCommand: %s", command.ID())
	MockJSONCommandID = command.ID()

	mockSocket, _ := NewMock("https://www.example.com/command")
	go mockSocket.Listen()
	mockSocket.SendCommand(command)
	mockSocket.Stop()

	if "Mock Command Result" != command.Result() {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", command.Result())
	}
}

func TestListenCommandError(t *testing.T) {
	MockJSONData = []byte(`"Mock Command Result"`)
	MockJSONRead = false
	MockJSONType = "command"
	MockJSONError = true
	MockJSONThrowError = false

	command := NewCommand("Some.methodError", nil)
	log.Infof("TestListenError: %s", command.ID())
	MockJSONCommandID = command.ID()

	mockSocket, _ := NewMock("https://www.example.com/error")
	go mockSocket.Listen()
	mockSocket.SendCommand(command)
	mockSocket.Stop()

	if "Mock Command Result" != command.Result() {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", command.Result())
	}
}

func TestListenEvent(t *testing.T) {
	MockJSONData = []byte(`"Mock Event Result"`)
	MockJSONRead = false
	MockJSONType = "event"
	MockJSONError = false
	MockJSONThrowError = false

	eventResponse1 := make(chan *Response)
	handler1 := NewEventHandler("Some.event", func(response *Response) {
		eventResponse1 <- response
	})

	eventResponse2 := make(chan []byte)
	handler2 := NewEventHandler("Some.event", func(response *Response) {
		tmp, _ := json.Marshal(response)
		eventResponse2 <- tmp
	})

	mockSocket, _ := NewMock("https://www.example.com/event")
	mockSocket.AddEventHandler(handler1)
	mockSocket.AddEventHandler(handler2)
	go mockSocket.Listen()
	mockSocket.Stop()

	response1 := <-eventResponse1
	response2 := <-eventResponse2

	if nil == response1.Result {
		t.Errorf("Invalid result: expected 'Mock Event Result', received nil")
	}

	if `"Mock Event Result"` != string(response1.Result) {
		t.Errorf("Invalid result: expected 'Mock Event Result', received '%s'", response1.Result)
	}

	response1JSON, _ := json.Marshal(response1)
	if string(response2) != string(response1JSON) {
		t.Errorf("Event handlers returned mismatched data")
	}
}

func TestReadJSONError(t *testing.T) {
	MockJSONData = []byte(`"Mock Read Error"`)
	MockJSONRead = false
	MockJSONType = "command"
	MockJSONError = true
	MockJSONThrowError = true

	mockSocket, _ := NewMock("https://www.example.com/error")
	err := mockSocket.Listen()

	if nil == err {
		t.Errorf("Expected an error, received nil")
	}

	if "Mock Read Error" != err.Error() {
		t.Errorf("Expected error message 'Mock Read Error', received '%s'", err.Error())
	}
}

func TestURL(t *testing.T) {
	url := "https://www.example.com/"
	mockSocket, _ := NewMock(url)
	if mockSocket.URL() != url {
		t.Errorf("Socketer.URL() failed to return the correct URL")
	}
}

func TestErrorType(t *testing.T) {
	var v Error

	if "int" != reflect.TypeOf(v.Code).String() {
		t.Errorf("Error.Code is expected to be of type int, %s found", reflect.TypeOf(v.Code).String())
	}
	if 0 != v.Code {
		t.Errorf("Error.Code should have a default value of 0, %v found", v.Code)
	}

	if "json.RawMessage" != reflect.TypeOf(v.Data).String() {
		t.Errorf("Error.Data is expected to be of type json.RawMessage, %s found", reflect.TypeOf(v.Data).String())
	}
	if nil != v.Data {
		t.Errorf("Error.Data should have a default value of nil, %v found", v.Data)
	}

	if "string" != reflect.TypeOf(v.Message).String() {
		t.Errorf("Error.Message is expected to be of type int, %s found", reflect.TypeOf(v.Message).String())
	}
	if "" != v.Message {
		t.Errorf("Error.Message should have a default value of empty string, %v found", v.Message)
	}
}

func TestResponseType(t *testing.T) {
	var v Response

	if "*socket.Error" != reflect.TypeOf(v.Error).String() {
		t.Errorf("Error.Code is expected to be of type socket.Error, %s found", reflect.TypeOf(v.Error).String())
	}
	if nil != v.Error {
		t.Errorf("Response.Code should have a default value of nil, %v found", v.Error)
	}

	if "int" != reflect.TypeOf(v.ID).String() {
		t.Errorf("Response.ID is expected to be of type int, %s found", reflect.TypeOf(v.ID).String())
	}
	if 0 != v.ID {
		t.Errorf("Response.ID should have a default value of 0, %v found", v.ID)
	}

	if "string" != reflect.TypeOf(v.Method).String() {
		t.Errorf("Response.Method is expected to be of type int, %s found", reflect.TypeOf(v.Method).String())
	}
	if "" != v.Method {
		t.Errorf("Response.Method should have a default value of empty string, %v found", v.Method)
	}

	if "json.RawMessage" != reflect.TypeOf(v.Params).String() {
		t.Errorf("Response.Params is expected to be of type json.RawMessage, %s found", reflect.TypeOf(v.Params).String())
	}
	if nil != v.Params {
		t.Errorf("Response.Params should have a default value of nil, %v found", v.Params)
	}

	if "json.RawMessage" != reflect.TypeOf(v.Result).String() {
		t.Errorf("Response.Result is expected to be of type json.RawMessage, %s found", reflect.TypeOf(v.Result).String())
	}
	if nil != v.Result {
		t.Errorf("Response.Result should have a default value of nil, %v found", v.Result)
	}
}
