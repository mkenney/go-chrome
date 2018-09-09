package socket

import (
	"testing"

	console "github.com/mkenney/go-chrome/tot/console"
)

func TestConsoleClearMessages(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &console.ClearMessagesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Console().ClearMessages()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Console().ClearMessages()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestConsoleDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &console.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Console().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Console().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestConsoleEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &console.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Console().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Console().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestConsoleOnMessageAdded(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &console.MessageAddedEvent{
		Message: &console.Message{
			Source: console.MessageSource.XML,
			Level:  console.MessageLevelLog,
			Text:   "message text",
			URL:    "message URL",
			Line:   10,
			Column: 10,
		},
	}

	resultChan := make(chan *console.MessageAddedEvent)
	soc.Console().OnMessageAdded(func(eventData *console.MessageAddedEvent) {
		resultChan <- eventData
	})

	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Console.messageAdded",
	})
	result := <-resultChan
	if result.Message.Source != mockResult.Message.Source {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockResult.Message.Source,
			result.Message.Source,
		)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Console.messageAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
