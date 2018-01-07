package socket

import (
	"encoding/json"

	overlay "github.com/mkenney/go-chrome/cdtp/overlay"
)

/*
OverlayProtocol provides a namespace for the Chrome Overlay protocol methods.
The Overlay protocol provides various functionality related to drawing atop the
inspected page.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/ EXPERIMENTAL.
*/
type OverlayProtocol struct {
	Socket Socketer
}

/*
Disable disables domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-disable
*/
func (protocol *OverlayProtocol) Disable() chan *overlay.DisableResult {
	resultChan := make(chan *overlay.DisableResult)
	command := NewCommand(protocol.Socket, "Overlay.disable", nil)
	result := &overlay.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-enable
*/
func (protocol *OverlayProtocol) Enable() chan *overlay.EnableResult {
	resultChan := make(chan *overlay.EnableResult)
	command := NewCommand(protocol.Socket, "Overlay.enable", nil)
	result := &overlay.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetHighlightObjectForTest is for testing.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
*/
func (protocol *OverlayProtocol) GetHighlightObjectForTest(
	params *overlay.GetHighlightObjectForTestParams,
) chan *overlay.GetHighlightObjectForTestResult {
	resultChan := make(chan *overlay.GetHighlightObjectForTestResult)
	command := NewCommand(protocol.Socket, "Overlay.getHighlightObjectForTest", params)
	result := &overlay.GetHighlightObjectForTestResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
HideHighlight hides any highlight.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-hideHighlight
*/
func (protocol *OverlayProtocol) HideHighlight() chan *overlay.HideHighlightResult {
	resultChan := make(chan *overlay.HideHighlightResult)
	command := NewCommand(protocol.Socket, "Overlay.hideHighlight", nil)
	result := &overlay.HideHighlightResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
HighlightFrame highlights owner element of the frame with given ID.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
*/
func (protocol *OverlayProtocol) HighlightFrame(
	params *overlay.HighlightFrameParams,
) chan *overlay.HighlightFrameResult {
	resultChan := make(chan *overlay.HighlightFrameResult)
	command := NewCommand(protocol.Socket, "Overlay.highlightFrame", params)
	result := &overlay.HighlightFrameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
HighlightNode highlights DOM node with given ID or with the given JavaScript
object wrapper. Either nodeID or objectID must be specified.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
*/
func (protocol *OverlayProtocol) HighlightNode(
	params *overlay.HighlightNodeParams,
) chan *overlay.HighlightNodeResult {
	resultChan := make(chan *overlay.HighlightNodeResult)
	command := NewCommand(protocol.Socket, "Overlay.highlightNode", params)
	result := &overlay.HighlightNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
HighlightQuad highlights given quad. Coordinates are absolute with respect to
the main frame viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
*/
func (protocol *OverlayProtocol) HighlightQuad(
	params *overlay.HighlightQuadParams,
) chan *overlay.HighlightQuadResult {
	resultChan := make(chan *overlay.HighlightQuadResult)
	command := NewCommand(protocol.Socket, "Overlay.highlightQuad", params)
	result := &overlay.HighlightQuadResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
HighlightRect highlights given rectangle. Coordinates are absolute with respect
to the main frame viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
*/
func (protocol *OverlayProtocol) HighlightRect(
	params *overlay.HighlightRectParams,
) chan *overlay.HighlightRectResult {
	resultChan := make(chan *overlay.HighlightRectResult)
	command := NewCommand(protocol.Socket, "Overlay.highlightRect", params)
	result := &overlay.HighlightRectResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetInspectMode enters the 'inspect' mode. In this mode, elements that user is
hovering over are highlighted. Backend then generates 'inspectNodeRequested'
event upon element selection.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
func (protocol *OverlayProtocol) SetInspectMode(
	params *overlay.SetInspectModeParams,
) chan *overlay.SetInspectModeResult {
	resultChan := make(chan *overlay.SetInspectModeResult)
	command := NewCommand(protocol.Socket, "Overlay.setInspectMode", params)
	result := &overlay.SetInspectModeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetPausedInDebuggerMessage sets the paused message

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
func (protocol *OverlayProtocol) SetPausedInDebuggerMessage(
	params *overlay.SetPausedInDebuggerMessageParams,
) chan *overlay.SetPausedInDebuggerMessageResult {
	resultChan := make(chan *overlay.SetPausedInDebuggerMessageResult)
	command := NewCommand(protocol.Socket, "Overlay.setPausedInDebuggerMessage", params)
	result := &overlay.SetPausedInDebuggerMessageResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetShowDebugBorders requests that backend shows debug borders on layers.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
*/
func (protocol *OverlayProtocol) SetShowDebugBorders(
	params *overlay.SetShowDebugBordersParams,
) chan *overlay.SetShowDebugBordersResult {
	resultChan := make(chan *overlay.SetShowDebugBordersResult)
	command := NewCommand(protocol.Socket, "Overlay.setShowDebugBorders", params)
	result := &overlay.SetShowDebugBordersResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetShowFPSCounter requests that backend shows the FPS counter.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
*/
func (protocol *OverlayProtocol) SetShowFPSCounter(
	params *overlay.SetShowFPSCounterParams,
) chan *overlay.SetShowFPSCounterResult {
	resultChan := make(chan *overlay.SetShowFPSCounterResult)
	command := NewCommand(protocol.Socket, "Overlay.setShowFPSCounter", params)
	result := &overlay.SetShowFPSCounterResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetShowPaintRects that backend shows paint rectangles.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
*/
func (protocol *OverlayProtocol) SetShowPaintRects(
	params *overlay.SetShowPaintRectsParams,
) chan *overlay.SetShowPaintRectsResult {
	resultChan := make(chan *overlay.SetShowPaintRectsResult)
	command := NewCommand(protocol.Socket, "Overlay.setShowPaintRects", params)
	result := &overlay.SetShowPaintRectsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetShowScrollBottleneckRects requests that backend shows scroll bottleneck rects.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
*/
func (protocol *OverlayProtocol) SetShowScrollBottleneckRects(
	params *overlay.SetShowScrollBottleneckRectsParams,
) chan *overlay.SetShowScrollBottleneckRectsResult {
	resultChan := make(chan *overlay.SetShowScrollBottleneckRectsResult)
	command := NewCommand(protocol.Socket, "Overlay.setShowScrollBottleneckRects", params)
	result := &overlay.SetShowScrollBottleneckRectsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetShowViewportSizeOnResize paints viewport size upon main frame resize.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowViewportSizeOnResize
*/
func (protocol *OverlayProtocol) SetShowViewportSizeOnResize(
	params *overlay.SetShowViewportSizeOnResizeParams,
) chan *overlay.SetShowViewportSizeOnResizeResult {
	resultChan := make(chan *overlay.SetShowViewportSizeOnResizeResult)
	command := NewCommand(protocol.Socket, "Overlay.setShowViewportSizeOnResize", params)
	result := &overlay.SetShowViewportSizeOnResizeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetSuspended sets the suspended state

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setSuspended
*/
func (protocol *OverlayProtocol) SetSuspended(
	params *overlay.SetSuspendedParams,
) chan *overlay.SetSuspendedResult {
	resultChan := make(chan *overlay.SetSuspendedResult)
	command := NewCommand(protocol.Socket, "Overlay.setSuspended", params)
	result := &overlay.SetSuspendedResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
OnInspectNodeRequested adds a handler to the Overlay.inspectNodeRequested event.
Overlay.inspectNodeRequested fires when the node should be inspected. This
happens after call to `setInspectMode` or when user manually inspects an
element.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-inspectNodeRequested
*/
func (protocol *OverlayProtocol) OnInspectNodeRequested(
	callback func(event *overlay.InspectNodeRequestedEvent),
) {
	handler := NewEventHandler(
		"Overlay.inspectNodeRequested",
		func(response *Response) {
			event := &overlay.InspectNodeRequestedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnNodeHighlightRequested adds a handler to the Overlay.nodeHighlightRequested
event. Overlay.nodeHighlightRequested fires when the node should be highlighted.
This happens after call to `setInspectMode`.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-nodeHighlightRequested
*/
func (protocol *OverlayProtocol) OnNodeHighlightRequested(
	callback func(event *overlay.NodeHighlightRequestedEvent),
) {
	handler := NewEventHandler(
		"Overlay.nodeHighlightRequested",
		func(response *Response) {
			event := &overlay.NodeHighlightRequestedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnScreenshotRequested adds a handler to the Overlay.screenshotRequested event.
Overlay.screenshotRequested fires when user asks to capture screenshot of some
area on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-screenshotRequested
*/
func (protocol *OverlayProtocol) OnScreenshotRequested(
	callback func(event *overlay.ScreenshotRequestedEvent),
) {
	handler := NewEventHandler(
		"Overlay.screenshotRequested",
		func(response *Response) {
			event := &overlay.ScreenshotRequestedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
