package socket

import (
	"reflect"
	"testing"
	"time"
)

func TestNewSocket(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()

	soc := New(chrome.URL)
	defer soc.Stop()

	if err := soc.Disconnect(); nil != err && "*errors.errorString" != reflect.TypeOf(err).String() {
		t.Errorf("Socketer.Disconnect() must return an error or nil, %s found", reflect.TypeOf(err).String())
	}
}

func TestCommandNotFound(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	chrome.AddData(MockData{
		ID:     999,
		Err:    &Error{},
		Method: "Some.methodError",
		Result: "Mock Command Result",
	})
	time.Sleep(1 * time.Second)
}

func TestSocketStop(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)

	time.Sleep(1 * time.Second)
	if err := soc.Stop(); nil != err {
		t.Errorf("Expected nil, got error: %T", err)
	}
}

func TestSocketDisconnect(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)

	if err := soc.Disconnect(); nil != err && "*errors.errorString" != reflect.TypeOf(err).String() && "errors.Err" != reflect.TypeOf(err).String() {
		t.Errorf("Socketer.Disconnect() must return an error or nil, %s found", reflect.TypeOf(err).String())
	}
}

func TestListenCommand(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	command := NewCommand(soc, "Some.method", nil)
	resultChan := soc.SendCommand(command)
	chrome.AddData(MockData{
		ID:     command.ID(),
		Err:    &Error{},
		Method: "Some.method",
		Result: "Mock Command Result",
	})
	result := <-resultChan
	if `"Mock Command Result"` != string(result.Result) {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", result.Result)
	}
}

func TestListenCommandError(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	command := NewCommand(soc, "Some.methodError", nil)
	resultChan := soc.SendCommand(command)
	chrome.AddData(MockData{
		ID:     command.ID(),
		Err:    &Error{},
		Method: "Some.methodError",
		Result: "Mock Command Result",
	})
	result := <-resultChan
	if `"Mock Command Result"` != string(result.Result) {
		t.Errorf("Invalid result: expected 'Mock Command Result', received '%s'", result.Result)
	}
}

func TestListenCommandUnknown(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	command := NewCommand(soc, "", nil)
	command.SetID(0)
	resultChan := soc.SendCommand(command)
	chrome.AddData(MockData{
		ID:  0,
		Err: nil,
	})
	result := <-resultChan
	if nil == result.Error {
		t.Errorf("Expected error, received nil")
	}
}

func TestRemoveEventHandler(t *testing.T) {
	var err error
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	handler1 := NewEventHandler(
		"Test.event",
		func(response *Response) {},
	)
	handler2 := NewEventHandler(
		"Test.event",
		func(response *Response) {},
	)

	// Remove before added
	err = soc.RemoveEventHandler(handler1)
	if nil == err {
		t.Errorf("Expected error, received nil")
	}

	// Remove added handler
	soc.AddEventHandler(handler1)
	err = soc.RemoveEventHandler(handler1)
	if nil != err {
		t.Errorf("Expected nil, received error: %s", err.Error())
	}

	// Removed but never added
	err = soc.RemoveEventHandler(handler2)
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
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	if chrome.URL.String() != soc.URL().String() {
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
