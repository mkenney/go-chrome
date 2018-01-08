package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	browser "github.com/mkenney/go-chrome/cdtp/browser"
	target "github.com/mkenney/go-chrome/cdtp/target"
)

func TestBrowserClose(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Browser().Close()
	mockSocket.Conn().AddMockData(&Response{
		ID:    mockSocket.CurCommandID(),
		Error: &Error{},
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Browser().Close()
	mockSocket.Conn().AddMockData(&Response{
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

func TestBrowserGetVersion(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Browser().GetVersion()
	mockResult := &browser.GetVersionResult{
		ProtocolVersion: "1.2",
		Product:         "product",
		Revision:        "revision",
		UserAgent:       "user-agent",
		JSVersion:       "V8 version",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.ProtocolVersion != mockResult.ProtocolVersion {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.ProtocolVersion,
			result.ProtocolVersion,
		)
	}

	resultChan = mockSocket.Browser().GetVersion()
	mockSocket.Conn().AddMockData(&Response{
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

func TestBrowserGetWindowBounds(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Browser().GetWindowBounds(&browser.GetWindowBoundsParams{
		WindowID: browser.WindowID(1),
	})
	mockResult := &browser.GetWindowBoundsResult{
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.Bounds.Height != mockResult.Bounds.Height {
		t.Errorf(
			"Expected %d, received %d",
			mockResult.Bounds.Height,
			result.Bounds.Height,
		)
	}

	resultChan = mockSocket.Browser().GetWindowBounds(&browser.GetWindowBoundsParams{
		WindowID: browser.WindowID(1),
	})
	mockSocket.Conn().AddMockData(&Response{
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

func TestBrowserGetWindowForTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Browser().GetWindowForTarget(&browser.GetWindowForTargetParams{
		TargetID: target.ID("target-id"),
	})
	mockResult := &browser.GetWindowForTargetResult{
		WindowID: browser.WindowID(1),
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.WindowID != mockResult.WindowID {
		t.Errorf(
			"Expected %d, received %d",
			mockResult.WindowID,
			result.WindowID,
		)
	}
	if result.Bounds.Height != mockResult.Bounds.Height {
		t.Errorf(
			"Expected %d, received %d",
			mockResult.Bounds.Height,
			result.Bounds.Height,
		)
	}

	resultChan = mockSocket.Browser().GetWindowForTarget(&browser.GetWindowForTargetParams{
		TargetID: target.ID("target-id"),
	})
	mockSocket.Conn().AddMockData(&Response{
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

func TestBrowserSetWindowBounds(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Browser().SetWindowBounds(&browser.SetWindowBoundsParams{
		WindowID: browser.WindowID(1),
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	})
	mockResult := &browser.SetWindowBoundsResult{
		WindowID: browser.WindowID(1),
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.WindowID != mockResult.WindowID {
		t.Errorf(
			"Expected %d, received %d",
			mockResult.WindowID,
			result.WindowID,
		)
	}
	if result.Bounds.Height != mockResult.Bounds.Height {
		t.Errorf(
			"Expected %d, received %d",
			mockResult.Bounds.Height,
			result.Bounds.Height,
		)
	}

	resultChan = mockSocket.Browser().SetWindowBounds(&browser.SetWindowBoundsParams{
		WindowID: browser.WindowID(1),
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	})
	mockSocket.Conn().AddMockData(&Response{
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
