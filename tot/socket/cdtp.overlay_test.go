package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/overlay"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestOverlayDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &overlay.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &overlay.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayGetHighlightObjectForTest(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.GetHighlightObjectForTestParams{
		NodeID: dom.NodeID(1),
	}
	mockResult := &overlay.GetHighlightObjectForTestResult{
		Highlight: map[string]string{"key": "data"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().GetHighlightObjectForTest(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Highlight["key"] != result.Highlight["key"] {
		t.Errorf("Expected %v, got %v", mockResult.Highlight["key"], result.Highlight["key"])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().GetHighlightObjectForTest(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayHideHighlight(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &overlay.HideHighlightResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().HideHighlight()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().HideHighlight()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayHighlightFrame(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.HighlightFrameParams{
		FrameID:             page.FrameID("frame-id"),
		ContentColor:        &dom.RGBA{},
		ContentOutlineColor: &dom.RGBA{},
	}
	mockResult := &overlay.HighlightFrameResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().HighlightFrame(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().HighlightFrame(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayHighlightNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.HighlightNodeParams{
		HighlightConfig: &overlay.HighlightConfig{},
		NodeID:          dom.NodeID(1),
		BackendNodeID:   dom.BackendNodeID(1),
		ObjectID:        runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &overlay.HighlightNodeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().HighlightNode(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().HighlightNode(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayHighlightQuad(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.HighlightQuadParams{
		Quad:         dom.Quad{1, 2},
		Color:        &dom.RGBA{},
		OutlineColor: &dom.RGBA{},
	}
	mockResult := &overlay.HighlightQuadResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().HighlightQuad(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().HighlightQuad(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayHighlightRect(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.HighlightRectParams{
		X:            1,
		Y:            1,
		Width:        1,
		Height:       1,
		Color:        &dom.RGBA{},
		OutlineColor: &dom.RGBA{},
	}
	mockResult := &overlay.HighlightRectResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().HighlightRect(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().HighlightRect(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetInspectMode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetInspectModeParams{
		Mode: overlay.InspectMode.SearchForNode,
	}
	mockResult := &overlay.SetInspectModeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetInspectMode(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetInspectMode(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetPausedInDebuggerMessage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetPausedInDebuggerMessageParams{
		Message: "message",
	}
	mockResult := &overlay.SetPausedInDebuggerMessageResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetPausedInDebuggerMessage(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetPausedInDebuggerMessage(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetShowDebugBorders(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetShowDebugBordersParams{
		Show: true,
	}
	mockResult := &overlay.SetShowDebugBordersResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetShowDebugBorders(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetShowDebugBorders(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetShowFPSCounter(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetShowFPSCounterParams{
		Show: true,
	}
	mockResult := &overlay.SetShowFPSCounterResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetShowFPSCounter(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetShowFPSCounter(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetShowPaintRects(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetShowPaintRectsParams{
		Result: true,
	}
	mockResult := &overlay.SetShowPaintRectsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetShowPaintRects(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetShowPaintRects(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetShowScrollBottleneckRects(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetShowScrollBottleneckRectsParams{
		Show: true,
	}
	mockResult := &overlay.SetShowScrollBottleneckRectsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetShowScrollBottleneckRects(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetShowScrollBottleneckRects(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetShowViewportSizeOnResize(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetShowViewportSizeOnResizeParams{
		Show: true,
	}
	mockResult := &overlay.SetShowViewportSizeOnResizeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetShowViewportSizeOnResize(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetShowViewportSizeOnResize(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlaySetSuspended(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &overlay.SetSuspendedParams{
		Suspended: true,
	}
	mockResult := &overlay.SetSuspendedResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Overlay().SetSuspended(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Overlay().SetSuspended(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayOnInspectNodeRequested(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *overlay.InspectNodeRequestedEvent)
	soc.Overlay().OnInspectNodeRequested(func(eventData *overlay.InspectNodeRequestedEvent) {
		resultChan <- eventData
	})

	mockResult := &overlay.InspectNodeRequestedEvent{
		BackendNodeID: dom.BackendNodeID(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Overlay.inspectNodeRequested",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.BackendNodeID != result.BackendNodeID {
		t.Errorf("Expected %d, got %d", mockResult.BackendNodeID, result.BackendNodeID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Overlay.inspectNodeRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayOnNodeHighlightRequested(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *overlay.NodeHighlightRequestedEvent)
	soc.Overlay().OnNodeHighlightRequested(func(eventData *overlay.NodeHighlightRequestedEvent) {
		resultChan <- eventData
	})

	mockResult := &overlay.NodeHighlightRequestedEvent{
		NodeID: dom.NodeID(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Overlay.nodeHighlightRequested",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Overlay.nodeHighlightRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayOnScreenshotRequested(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *overlay.ScreenshotRequestedEvent)
	soc.Overlay().OnScreenshotRequested(func(eventData *overlay.ScreenshotRequestedEvent) {
		resultChan <- eventData
	})

	mockResult := &overlay.ScreenshotRequestedEvent{
		Viewport: &page.Viewport{
			X:      1,
			Y:      2,
			Width:  3,
			Height: 4,
			Scale:  5,
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Overlay.screenshotRequested",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Viewport.X != result.Viewport.X {
		t.Errorf("Expected %d, got %d", mockResult.Viewport.X, result.Viewport.X)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Overlay.screenshotRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
