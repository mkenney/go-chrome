package socket

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestNewSocket(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestNewSocket")
	socket := New(socketURL)
	if err := socket.Disconnect(); nil != err && "*errors.errorString" != reflect.TypeOf(err).String() {
		t.Errorf("Socketer.Disconnect() must return an error or nil, %s found", reflect.TypeOf(err).String())
	}
}

func TestCommandNotFound(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestCommandNotFound")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     999,
		Error:  &Error{},
		Method: "Some.methodError",
		Result: []byte(`"Mock Command Result"`),
	})
	time.Sleep(1 * time.Second)
}

func TestSocketStop(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestSocketStop")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	time.Sleep(1 * time.Second)
	mockSocket.Stop()
}

func TestSocketDisconnect(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestSocketDisconnect")
	mockSocket := NewMock(socketURL)
	if err := mockSocket.Disconnect(); nil != err && "*errors.errorString" != reflect.TypeOf(err).String() && "errors.Err" != reflect.TypeOf(err).String() {
		t.Errorf("Socketer.Disconnect() must return an error or nil, %s found", reflect.TypeOf(err).String())
	}

	// Test the disconnect timeout
	mockSocket = NewMock(socketURL)
	mockSocket.Listen()
	mockSocket.Conn().(*MockChromeWebSocket).Sleep(10 * time.Second)
	start := time.Now()
	if err := mockSocket.Disconnect(); nil != err && "*errors.errorString" != reflect.TypeOf(err).String() && "errors.Err" != reflect.TypeOf(err).String() {
		t.Errorf("Socketer.Disconnect() must return an error or nil, %s found", reflect.TypeOf(err).String())
	}
	elapsed := time.Since(start)
	if elapsed < 1*time.Second {
		t.Errorf("Expected disconnect timeout in 1 seconds, %s elapsed", elapsed)
	}
}

func TestListenCommand(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestListenCommand")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	command := NewCommand(mockSocket, "Some.method", nil)
	resultChan := mockSocket.SendCommand(command)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     command.ID(),
		Error:  &Error{},
		Method: "Some.method",
		Result: []byte(`"Mock Command Result"`),
	})
	result := <-resultChan
	if `"Mock Command Result"` != string(result.Result) {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", result.Result)
	}
}

func TestListenCommandError(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestListenCommandError")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	command := NewCommand(mockSocket, "Some.methodError", nil)
	resultChan := mockSocket.SendCommand(command)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     command.ID(),
		Error:  &Error{},
		Method: "Some.methodError",
		Result: []byte(`"Mock Command Result"`),
	})
	result := <-resultChan
	if `"Mock Command Result"` != string(result.Result) {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", result.Result)
	}
}

func TestListenCommandUnknown(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestListenCommandUnknown")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	command := NewCommand(mockSocket, "", nil)
	command.SetID(0)
	resultChan := mockSocket.SendCommand(command)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  nil,
		Method: "",
	})
	result := <-resultChan
	if nil == result.Error {
		t.Errorf("Expected error, received nil")
	}
}

func TestRemoveEventHandler(t *testing.T) {
	var err error
	socketURL, _ := url.Parse("https://test:9222/TestRemoveEventHandler")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	handler1 := NewEventHandler(
		"Test.event",
		func(response *Response) {},
	)
	handler2 := NewEventHandler(
		"Test.event",
		func(response *Response) {},
	)

	// Remove before added
	err = mockSocket.RemoveEventHandler(handler1)
	if nil == err {
		t.Errorf("Expected error, received nil")
	}

	// Remove added handler
	mockSocket.AddEventHandler(handler1)
	err = mockSocket.RemoveEventHandler(handler1)
	if nil != err {
		t.Errorf("Expected nil, received error: %s", err.Error())
	}

	// Removed but never added
	err = mockSocket.RemoveEventHandler(handler2)
	if nil != err {
		t.Errorf("Expected nil, received error: %s", err.Error())
	}
}

//func TestReadJSONError(t *testing.T) {
//	socketURL, _ := url.Parse("https://test:9222/TestReadJSONError")
//	mockSocket := NewMock(socketURL)
//	mockSocket.Listen()
//	defer mockSocket.Stop()
//
//	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(
//		0,
//		&Error{},
//		"Some.event",
//		"Mock Event Result",
//	)
//
//	mockSocket := NewMock(socketURL)
//	err := mockSocket.Listen()
//
//	if nil == err {
//		t.Errorf("Expected an error, received nil")
//	}
//
//	if "Mock Read Error" != err.Error() {
//		t.Errorf("Expected error message 'Mock Read Error', received '%s'", err.Error())
//	}
//}

func TestURL(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestURL")
	mockSocket := NewMock(socketURL)
	if "https://test:9222/TestURL" != mockSocket.URL().String() {
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
		t.Errorf("Error.Code is expected to be of type Error, %s found", reflect.TypeOf(v.Error).String())
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
