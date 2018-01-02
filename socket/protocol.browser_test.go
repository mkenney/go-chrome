package socket

import (
	"net/url"
	"testing"

	browser "github.com/mkenney/go-chrome/protocol/browser"
	target "github.com/mkenney/go-chrome/protocol/target"
)

func TestBrowserClose(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"Browser.close",
		nil,
	)
	go func() {
		err := mockSocket.Browser().Close()
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
	}()
}

func TestBrowserGetVersion(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockResult := &browser.GetVersionResult{
		ProtocolVersion: "1.2",
		Product:         "product",
		Revision:        "revision",
		UserAgent:       "user-agent",
		JSVersion:       "V8 version",
	}
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"Browser.getVersion",
		mockResult,
	)
	go func() {
		result, err := mockSocket.Browser().GetVersion()
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.ProtocolVersion != mockResult.ProtocolVersion {
			t.Errorf(
				"Expected %s, received %s",
				mockResult.ProtocolVersion,
				result.ProtocolVersion,
			)
		}
	}()
}

func TestBrowserGetWindowBounds(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockResult := &browser.GetWindowBoundsResult{
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	}
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"Browser.getWindowBounds",
		mockResult,
	)
	go func() {
		result, err := mockSocket.Browser().GetWindowBounds(&browser.GetWindowBoundsParams{
			WindowID: browser.WindowID(1),
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.Bounds.Height != mockResult.Bounds.Height {
			t.Errorf(
				"Expected %d, received %d",
				mockResult.Bounds.Height,
				result.Bounds.Height,
			)
		}
	}()
}

func TestBrowserGetWindowForTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"Browser.getWindowForTarget",
		mockResult,
	)
	go func() {
		result, err := mockSocket.Browser().GetWindowForTarget(&browser.GetWindowForTargetParams{
			TargetID: target.ID("target-id"),
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
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
	}()
}

func TestBrowserSetWindowBounds(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"Browser.setWindowBounds",
		mockResult,
	)
	go func() {
		result, err := mockSocket.Browser().SetWindowBounds(&browser.SetWindowBoundsParams{
			WindowID: browser.WindowID(1),
			Bounds: &browser.Bounds{
				Height:      100,
				Left:        0,
				Top:         0,
				Width:       100,
				WindowState: browser.WindowState("window-state"),
			},
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
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
	}()
}
