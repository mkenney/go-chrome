package socket

import (
	"testing"

	browser "github.com/mkenney/go-chrome/tot/browser"
	target "github.com/mkenney/go-chrome/tot/target"
)

func TestBrowserClose(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Browser().Close()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Browser().Close()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestBrowserGetVersion(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &browser.GetVersionResult{
		ProtocolVersion: "1.2",
		Product:         "product",
		Revision:        "revision",
		UserAgent:       "user-agent",
		JSVersion:       "V8 version",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Browser().GetVersion()
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Browser().GetVersion()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestBrowserGetWindowBounds(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &browser.GetWindowBoundsResult{
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Browser().GetWindowBounds(&browser.GetWindowBoundsParams{
		WindowID: browser.WindowID(1),
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Browser().GetWindowBounds(&browser.GetWindowBoundsParams{
		WindowID: browser.WindowID(1),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestBrowserGetWindowForTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Browser().GetWindowForTarget(&browser.GetWindowForTargetParams{
		TargetID: target.ID("target-id"),
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Browser().GetWindowForTarget(&browser.GetWindowForTargetParams{
		TargetID: target.ID("target-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestBrowserSetWindowBounds(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Browser().SetWindowBounds(&browser.SetWindowBoundsParams{
		WindowID: browser.WindowID(1),
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Browser().SetWindowBounds(&browser.SetWindowBoundsParams{
		WindowID: browser.WindowID(1),
		Bounds: &browser.Bounds{
			Height:      100,
			Left:        0,
			Top:         0,
			Width:       100,
			WindowState: browser.WindowState("window-state"),
		},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
