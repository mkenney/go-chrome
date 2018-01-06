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
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Browser.close",
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
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
		Method: "Browser.getVersion",
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
		Method: "Browser.getWindowBounds",
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
		Method: "Browser.getWindowForTarget",
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
		Method: "Browser.setWindowBounds",
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
}
