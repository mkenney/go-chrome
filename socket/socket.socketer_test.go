package socket

import (
	"net/url"
	"reflect"
	"testing"
)

func TestSocketDisconnect(t *testing.T) {
	socketURL, _ := url.Parse("https://www.example.com/")
	socket := NewMock(socketURL)
	if err := socket.Disconnect(); nil != err && "*errors.errorString" != reflect.TypeOf(err).String() {
		t.Errorf("Socketer.Disconnect() must return an error or nil, %s found", reflect.TypeOf(err).String())
	}
}

func TestListenCommand(t *testing.T) {
	socketURL, _ := url.Parse("https://www.example.com/command")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	addMockWebsocketResponse(
		_commandID+1,
		&Error{},
		"Some.method",
		"Mock Command Result",
	)
	command := NewCommand("Some.method", nil)
	mockSocket.SendCommand(command)

	if `"Mock Command Result"` != string(command.Result()) {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", command.Result())
	}
}

func TestListenCommandError(t *testing.T) {
	socketURL, _ := url.Parse("https://www.example.com/error")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	addMockWebsocketResponse(
		_commandID+1,
		&Error{},
		"Some.methodError",
		"Mock Command Result",
	)
	command := NewCommand("Some.methodError", nil)
	mockSocket.SendCommand(command)
	if `"Mock Command Result"` != string(command.Result()) {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", command.Result())
	}
}

func TestReadJSONError(t *testing.T) {
	//	socketURL, _ := url.Parse("https://www.example.com/error")
	//	mockSocket := NewMock(socketURL)
	//	go mockSocket.Listen()
	//	defer mockSocket.Stop()
	//
	//	addMockWebsocketResponse(
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
}

func TestURL(t *testing.T) {
	socketURL, _ := url.Parse("https://www.example.com/")
	mockSocket := NewMock(socketURL)
	if "https://www.example.com/" != mockSocket.URL().String() {
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
