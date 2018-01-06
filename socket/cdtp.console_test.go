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
}

func TestConsoleOnMessageAdded(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	results := make(chan *console.MessageAddedEvent)
	mockSocket.Console().OnMessageAdded(func(eventData *console.MessageAddedEvent) {
		results <- eventData
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
	result := <-results
	if result.Message.Source != mockResult.Message.Source {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockResult.Message.Source,
			result.Message.Source,
		)
	}
}
