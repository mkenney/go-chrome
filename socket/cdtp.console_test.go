package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	console "github.com/mkenney/go-chrome/cdtp/console"
)

func TestConsoleClearMessages(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Console().ClearMessages()
	mockResult := &console.ClearMessagesResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Console.clearMessages",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Console().ClearMessages()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Console.clearMessages",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestConsoleDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Console().Disable()
	mockResult := &console.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Console.disable",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Console().Disable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Console.disable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestConsoleEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Console().Enable()
	mockResult := &console.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Console.enable",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Console().Enable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Console.enable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestConsoleOnMessageAdded(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *console.MessageAddedEvent)
	mockSocket.Console().OnMessageAdded(func(eventData *console.MessageAddedEvent) {
		resultChan <- eventData
	})
	mockResult := &console.MessageAddedEvent{
		Message: &console.Message{
			Source: "message source",
			Level:  "log level",
			Text:   "message text",
			URL:    "message URL",
			Line:   10,
			Column: 10,
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Console.messageAdded",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if result.Message.Source != mockResult.Message.Source {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockResult.Message.Source,
			result.Message.Source,
		)
	}
}
