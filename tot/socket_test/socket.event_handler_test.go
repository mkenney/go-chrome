package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/socket"
)

func TestHandleEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://www.example.com/error")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	eventResponse1 := make(chan *socket.Response)
	handler1 := socket.NewEventHandler("Some.event", func(response *socket.Response) {
		eventResponse1 <- response
	})

	eventResponse2 := make(chan *socket.Response)
	handler2 := socket.NewEventHandler("Some.event", func(response *socket.Response) {
		eventResponse2 <- response
	})

	eventResponse3 := make(chan *socket.Response)
	handler3 := socket.NewEventHandler("Inspector.targetCrashed", func(response *socket.Response) {
		eventResponse3 <- response
	})

	mockSocket.AddEventHandler(handler1)
	mockSocket.AddEventHandler(handler2)
	mockSocket.AddEventHandler(handler3)

	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Some.event",
		Result: []byte(`"Mock Event Result"`),
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Inspector.targetCrashed",
		Result: []byte(`"Mock Target Crashed"`),
	})

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
