package socket

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestHandleEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://www.example.com/error")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	eventResponse1 := make(chan *Response)
	handler1 := NewEventHandler("Some.event", func(response *Response) {
		eventResponse1 <- response
	})

	eventResponse2 := make(chan *Response)
	handler2 := NewEventHandler("Some.event", func(response *Response) {
		eventResponse2 <- response
	})

	eventResponse3 := make(chan *Response)
	handler3 := NewEventHandler("Inspector.targetCrashed", func(response *Response) {
		eventResponse3 <- response
	})

	mockSocket.AddEventHandler(handler1)
	mockSocket.AddEventHandler(handler2)
	mockSocket.AddEventHandler(handler3)

	addMockWebsocketResponse(
		0,
		&Error{},
		"Some.event",
		"Mock Event Result",
	)
	addMockWebsocketResponse(
		0,
		&Error{},
		"Inspector.targetCrashed",
		"Mock Target Crashed",
	)

	response1 := <-eventResponse1
	response2 := <-eventResponse2
	response3 := <-eventResponse3

	if nil == response1.Result {
		t.Errorf("Invalid result: expected 'Mock Event Result', received nil")
	}

	if `"Mock Event Result"` != string(response1.Result) {
		t.Errorf("Invalid result: expected 'Mock Event Result', received '%s'", response1.Result)
	}

	response1JSON, _ := json.Marshal(response1)
	response2JSON, _ := json.Marshal(response2)
	if string(response2JSON) != string(response1JSON) {
		t.Errorf("Event handlers returned mismatched data")
	}

	if nil == response3.Result {
		t.Errorf("Invalid result: expected 'Mock Target Crashed', received nil")
	}

	if `"Mock Target Crashed"` != string(response3.Result) {
		t.Errorf("Invalid result: expected 'Mock Target Crashed', received '%s'", response3.Result)
	}

}
