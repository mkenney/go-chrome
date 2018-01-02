package socket

import (
	"encoding/json"

	overlay "github.com/mkenney/go-chrome/cdtp/overlay"
	log "github.com/sirupsen/logrus"
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
func (protocol *OverlayProtocol) Disable() error {
	command := NewCommand("Overlay.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-enable
*/
func (protocol *OverlayProtocol) Enable() error {
	command := NewCommand("Overlay.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetHighlightObjectForTest is for testing.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
*/
func (protocol *OverlayProtocol) GetHighlightObjectForTest(
	params *overlay.GetHighlightObjectForTestParams,
) (*overlay.GetHighlightObjectForTestResult, error) {
	command := NewCommand("Overlay.getHighlightObjectForTest", params)
	result := &overlay.GetHighlightObjectForTestResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
HideHighlight hides any highlight.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-hideHighlight
*/
func (protocol *OverlayProtocol) HideHighlight() error {
	command := NewCommand("Overlay.hideHighlight", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
HighlightFrame highlights owner element of the frame with given ID.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
*/
func (protocol *OverlayProtocol) HighlightFrame(
	params *overlay.HighlightFrameParams,
) error {
	command := NewCommand("Overlay.highlightFrame", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
HighlightNode highlights DOM node with given ID or with the given JavaScript
object wrapper. Either nodeID or objectID must be specified.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
*/
func (protocol *OverlayProtocol) HighlightNode(
	params *overlay.HighlightNodeParams,
) error {
	command := NewCommand("Overlay.highlightNode", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
HighlightQuad highlights given quad. Coordinates are absolute with respect to
the main frame viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
*/
func (protocol *OverlayProtocol) HighlightQuad(
	params *overlay.HighlightQuadParams,
) error {
	command := NewCommand("Overlay.highlightQuad", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
HighlightRect highlights given rectangle. Coordinates are absolute with respect
to the main frame viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
*/
func (protocol *OverlayProtocol) HighlightRect(
	params *overlay.HighlightRectParams,
) error {
	command := NewCommand("Overlay.highlightRect", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetInspectMode enters the 'inspect' mode. In this mode, elements that user is
hovering over are highlighted. Backend then generates 'inspectNodeRequested'
event upon element selection.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
func (protocol *OverlayProtocol) SetInspectMode(
	params *overlay.SetInspectModeParams,
) error {
	command := NewCommand("Overlay.setInspectMode", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetPausedInDebuggerMessage sets the paused message

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
func (protocol *OverlayProtocol) SetPausedInDebuggerMessage(
	params *overlay.SetPausedInDebuggerMessageParams,
) error {
	command := NewCommand("Overlay.setPausedInDebuggerMessage", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetShowDebugBorders requests that backend shows debug borders on layers.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
*/
func (protocol *OverlayProtocol) SetShowDebugBorders(
	params *overlay.SetShowDebugBordersParams,
) error {
	command := NewCommand("Overlay.setShowDebugBorders", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetShowFPSCounter requests that backend shows the FPS counter.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
*/
func (protocol *OverlayProtocol) SetShowFPSCounter(
	params *overlay.SetShowFPSCounterParams,
) error {
	command := NewCommand("Overlay.setShowFPSCounter", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetShowPaintRects that backend shows paint rectangles.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
*/
func (protocol *OverlayProtocol) SetShowPaintRects(
	params *overlay.SetShowPaintRectsParams,
) error {
	command := NewCommand("Overlay.setShowPaintRects", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetShowScrollBottleneckRects requests that backend shows scroll bottleneck rects.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
*/
func (protocol *OverlayProtocol) SetShowScrollBottleneckRects(
	params *overlay.SetShowScrollBottleneckRectsParams,
) error {
	command := NewCommand("Overlay.setShowScrollBottleneckRects", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetShowViewportSizeOnResize paints viewport size upon main frame resize.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowViewportSizeOnResize
*/
func (protocol *OverlayProtocol) SetShowViewportSizeOnResize(
	params *overlay.SetShowViewportSizeOnResizeParams,
) error {
	command := NewCommand("Overlay.setShowViewportSizeOnResize", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetSuspended sets the suspended state

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setSuspended
*/
func (protocol *OverlayProtocol) SetSuspended(
	params *overlay.SetSuspendedParams,
) error {
	command := NewCommand("Overlay.setSuspended", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
