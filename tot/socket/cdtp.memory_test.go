package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	memory "github.com/mkenney/go-chrome/tot/cdtp/memory"
)

func TestMemoryGetDOMCounters(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestMemoryGetDOMCounters")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &memory.GetDOMCountersParams{
		Documents:        1,
		Nodes:            1,
		JsEventListeners: 1,
	}
	resultChan := mockSocket.Memory().GetDOMCounters(params)
	mockResult := &memory.GetDOMCountersResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Memory().GetDOMCounters(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestMemoryPrepareForLeakDetection(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestMemoryPrepareForLeakDetection")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Memory().PrepareForLeakDetection()
	mockResult := &memory.PrepareForLeakDetectionResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Memory().PrepareForLeakDetection()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestMemorySetPressureNotificationsSuppressed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestMemorySetPressureNotificationsSuppressed")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &memory.SetPressureNotificationsSuppressedParams{
		Suppressed: true,
	}
	resultChan := mockSocket.Memory().SetPressureNotificationsSuppressed(params)
	mockResult := &memory.SetPressureNotificationsSuppressedResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Memory().SetPressureNotificationsSuppressed(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestMemorySimulatePressureNotification(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestMemorySimulatePressureNotification")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &memory.SimulatePressureNotificationParams{
		Level: memory.PressureLevel("pressure-level"),
	}
	resultChan := mockSocket.Memory().SimulatePressureNotification(params)
	mockResult := &memory.SimulatePressureNotificationResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Memory().SimulatePressureNotification(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
