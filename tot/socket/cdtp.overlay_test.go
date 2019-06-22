package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/overlay"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestOverlayDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayDisable")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Overlay().Disable()
	mockResult := &overlay.DisableResult{}
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

	resultChan = mockSocket.Overlay().Disable()
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

func TestOverlayEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayEnable")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Overlay().Enable()
	mockResult := &overlay.EnableResult{}
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

	resultChan = mockSocket.Overlay().Enable()
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

func TestOverlayGetHighlightObjectForTest(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayGetHighlightObjectForTest")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.GetHighlightObjectForTestParams{
		NodeID: dom.NodeID(1),
	}
	resultChan := mockSocket.Overlay().GetHighlightObjectForTest(params)
	mockResult := &overlay.GetHighlightObjectForTestResult{
		Highlight: map[string]string{"key": "data"},
	}
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
	if mockResult.Highlight["key"] != result.Highlight["key"] {
		t.Errorf("Expected %v, got %v", mockResult.Highlight["key"], result.Highlight["key"])
	}

	resultChan = mockSocket.Overlay().GetHighlightObjectForTest(params)
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

func TestOverlayHideHighlight(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayHideHighlight")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Overlay().HideHighlight()
	mockResult := &overlay.HideHighlightResult{}
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

	resultChan = mockSocket.Overlay().HideHighlight()
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

func TestOverlayHighlightFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayHighlightFrame")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.HighlightFrameParams{
		FrameID:             page.FrameID("frame-id"),
		ContentColor:        &dom.RGBA{},
		ContentOutlineColor: &dom.RGBA{},
	}
	resultChan := mockSocket.Overlay().HighlightFrame(params)
	mockResult := &overlay.HighlightFrameResult{}
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

	resultChan = mockSocket.Overlay().HighlightFrame(params)
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

func TestOverlayHighlightNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayHighlightNode")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.HighlightNodeParams{
		HighlightConfig: &overlay.HighlightConfig{},
		NodeID:          dom.NodeID(1),
		BackendNodeID:   dom.BackendNodeID(1),
		ObjectID:        runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.Overlay().HighlightNode(params)
	mockResult := &overlay.HighlightNodeResult{}
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

	resultChan = mockSocket.Overlay().HighlightNode(params)
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

func TestOverlayHighlightQuad(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayHighlightQuad")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.HighlightQuadParams{
		Quad:         dom.Quad{1, 2},
		Color:        &dom.RGBA{},
		OutlineColor: &dom.RGBA{},
	}
	resultChan := mockSocket.Overlay().HighlightQuad(params)
	mockResult := &overlay.HighlightQuadResult{}
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

	resultChan = mockSocket.Overlay().HighlightQuad(params)
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

func TestOverlayHighlightRect(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayHighlightRect")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.HighlightRectParams{
		X:            1,
		Y:            1,
		Width:        1,
		Height:       1,
		Color:        &dom.RGBA{},
		OutlineColor: &dom.RGBA{},
	}
	resultChan := mockSocket.Overlay().HighlightRect(params)
	mockResult := &overlay.HighlightRectResult{}
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

	resultChan = mockSocket.Overlay().HighlightRect(params)
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

func TestOverlaySetInspectMode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetInspectMode")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetInspectModeParams{
		Mode: overlay.InspectMode.SearchForNode,
	}
	resultChan := mockSocket.Overlay().SetInspectMode(params)
	mockResult := &overlay.SetInspectModeResult{}
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

	resultChan = mockSocket.Overlay().SetInspectMode(params)
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

func TestOverlaySetPausedInDebuggerMessage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetPausedInDebuggerMessage")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetPausedInDebuggerMessageParams{
		Message: "message",
	}
	resultChan := mockSocket.Overlay().SetPausedInDebuggerMessage(params)
	mockResult := &overlay.SetPausedInDebuggerMessageResult{}
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

	resultChan = mockSocket.Overlay().SetPausedInDebuggerMessage(params)
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

func TestOverlaySetShowDebugBorders(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetShowDebugBorders")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetShowDebugBordersParams{
		Show: true,
	}
	resultChan := mockSocket.Overlay().SetShowDebugBorders(params)
	mockResult := &overlay.SetShowDebugBordersResult{}
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

	resultChan = mockSocket.Overlay().SetShowDebugBorders(params)
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

func TestOverlaySetShowFPSCounter(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetShowFPSCounter")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetShowFPSCounterParams{
		Show: true,
	}
	resultChan := mockSocket.Overlay().SetShowFPSCounter(params)
	mockResult := &overlay.SetShowFPSCounterResult{}
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

	resultChan = mockSocket.Overlay().SetShowFPSCounter(params)
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

func TestOverlaySetShowPaintRects(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetShowPaintRects")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetShowPaintRectsParams{
		Result: true,
	}
	resultChan := mockSocket.Overlay().SetShowPaintRects(params)
	mockResult := &overlay.SetShowPaintRectsResult{}
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

	resultChan = mockSocket.Overlay().SetShowPaintRects(params)
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

func TestOverlaySetShowScrollBottleneckRects(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetShowScrollBottleneckRects")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetShowScrollBottleneckRectsParams{
		Show: true,
	}
	resultChan := mockSocket.Overlay().SetShowScrollBottleneckRects(params)
	mockResult := &overlay.SetShowScrollBottleneckRectsResult{}
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

	resultChan = mockSocket.Overlay().SetShowScrollBottleneckRects(params)
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

func TestOverlaySetShowViewportSizeOnResize(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetShowViewportSizeOnResize")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetShowViewportSizeOnResizeParams{
		Show: true,
	}
	resultChan := mockSocket.Overlay().SetShowViewportSizeOnResize(params)
	mockResult := &overlay.SetShowViewportSizeOnResizeResult{}
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

	resultChan = mockSocket.Overlay().SetShowViewportSizeOnResize(params)
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

func TestOverlaySetSuspended(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlaySetSuspended")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &overlay.SetSuspendedParams{
		Suspended: true,
	}
	resultChan := mockSocket.Overlay().SetSuspended(params)
	mockResult := &overlay.SetSuspendedResult{}
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

	resultChan = mockSocket.Overlay().SetSuspended(params)
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

func TestOverlayOnInspectNodeRequested(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayOnInspectNodeRequested")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *overlay.InspectNodeRequestedEvent)
	mockSocket.Overlay().OnInspectNodeRequested(func(eventData *overlay.InspectNodeRequestedEvent) {
		resultChan <- eventData
	})
	mockResult := &overlay.InspectNodeRequestedEvent{
		BackendNodeID: dom.BackendNodeID(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Overlay.inspectNodeRequested",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.BackendNodeID != result.BackendNodeID {
		t.Errorf("Expected %d, got %d", mockResult.BackendNodeID, result.BackendNodeID)
	}

	resultChan = make(chan *overlay.InspectNodeRequestedEvent)
	mockSocket.Overlay().OnInspectNodeRequested(func(eventData *overlay.InspectNodeRequestedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Overlay.inspectNodeRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayOnNodeHighlightRequested(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayOnNodeHighlightRequested")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *overlay.NodeHighlightRequestedEvent)
	mockSocket.Overlay().OnNodeHighlightRequested(func(eventData *overlay.NodeHighlightRequestedEvent) {
		resultChan <- eventData
	})
	mockResult := &overlay.NodeHighlightRequestedEvent{
		NodeID: dom.NodeID(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Overlay.nodeHighlightRequested",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = make(chan *overlay.NodeHighlightRequestedEvent)
	mockSocket.Overlay().OnNodeHighlightRequested(func(eventData *overlay.NodeHighlightRequestedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Overlay.nodeHighlightRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestOverlayOnScreenshotRequested(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestOverlayOnScreenshotRequested")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *overlay.ScreenshotRequestedEvent)
	mockSocket.Overlay().OnScreenshotRequested(func(eventData *overlay.ScreenshotRequestedEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Overlay.screenshotRequested",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Viewport.X != result.Viewport.X {
		t.Errorf("Expected %d, got %d", mockResult.Viewport.X, result.Viewport.X)
	}

	resultChan = make(chan *overlay.ScreenshotRequestedEvent)
	mockSocket.Overlay().OnScreenshotRequested(func(eventData *overlay.ScreenshotRequestedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Overlay.screenshotRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
