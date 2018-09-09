package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/tethering"
)

func TestTetheringBind(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tethering.BindParams{
		Port: 1,
	}
	mockResult := &tethering.BindResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Tethering().Bind(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Tethering().Bind(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTetheringUnbind(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tethering.UnbindParams{
		Port: 1,
	}
	mockResult := &tethering.UnbindResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Tethering().Unbind(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Tethering().Unbind(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTetheringOnAccepted(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *tethering.AcceptedEvent)
	soc.Tethering().OnAccepted(func(eventData *tethering.AcceptedEvent) {
		resultChan <- eventData
	})

	mockResult := &tethering.AcceptedEvent{
		Port:         1,
		ConnectionID: "ConnectionID",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Tethering.accepted",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Port != result.Port {
		t.Errorf("Expected %d, got %d", mockResult.Port, result.Port)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Tethering.accepted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
