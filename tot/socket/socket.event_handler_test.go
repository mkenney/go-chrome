package socket

import (
	"encoding/json"
	"testing"
)

func TestHandleEvent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	eventResponse1 := make(chan *Response)
	handler1 := NewEventHandler("Some.event", func(response *Response) {
		eventResponse1 <- response
	})
	soc.AddEventHandler(handler1)

	eventResponse2 := make(chan *Response)
	handler2 := NewEventHandler("Some.event", func(response *Response) {
		eventResponse2 <- response
	})
	soc.AddEventHandler(handler2)

	eventResponse3 := make(chan *Response)
	handler3 := NewEventHandler("Inspector.targetCrashed", func(response *Response) {
		eventResponse3 <- response
	})
	soc.AddEventHandler(handler3)

	chrome.AddData(MockData{
		Err:    &Error{},
		Result: "Mock Event Result",
		Method: "Some.event",
	})
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: "Mock Target Crashed",
		Method: "Inspector.targetCrashed",
	})

	response1 := <-eventResponse1
	response2 := <-eventResponse2
	response3 := <-eventResponse3

	if nil == response1.Params {
		t.Errorf("Invalid result: expected 'Mock Event Result', received nil")
	}

	if `"Mock Event Result"` != string(response1.Params) {
		t.Errorf("Invalid result: expected 'Mock Event Result', received '%v'", string(response1.Result))
	}

	response1JSON, _ := json.Marshal(response1)
	response2JSON, _ := json.Marshal(response2)
	if string(response2JSON) != string(response1JSON) {
		t.Errorf("Event handlers returned mismatched data")
	}

	if nil == response3.Params {
		t.Errorf("Invalid result: expected 'Mock Target Crashed', received nil")
	}

	if `"Mock Target Crashed"` != string(response3.Params) {
		t.Errorf("Invalid result: expected 'Mock Target Crashed', received '%v'", string(response3.Result))
	}
}
