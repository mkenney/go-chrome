package protocol

import (
	"encoding/json"

	overlay "github.com/mkenney/go-chrome/protocol/overlay"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Overlay is a struct that provides a namespace for the Chrome Overlay protocol methods.

The Overlay protocol provides various functionality related to drawing atop the inspected page.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/
*/
type Overlay struct{}

/*
Disable disables domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-disable
*/
func (Overlay) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Overlay.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-enable
*/
func (Overlay) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Overlay.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetHighlightObjectForTest is for testing.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
*/
func (Overlay) GetHighlightObjectForTest(
	socket *sock.Socket,
	params *overlay.GetHighlightObjectForTestParams,
) (overlay.GetHighlightObjectForTestResult, error) {
	command := &sock.Command{
		Method: "Overlay.getHighlightObjectForTest",
		Params: params,
	}
	result := overlay.GetHighlightObjectForTestResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
HideHighlight hides any highlight.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-hideHighlight
*/
func (Overlay) HideHighlight(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Overlay.hideHighlight",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightFrame highlights owner element of the frame with given ID.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
*/
func (Overlay) HighlightFrame(
	socket *sock.Socket,
	params *overlay.HighlightFrameParams,
) error {
	command := &sock.Command{
		Method: "Overlay.highlightFrame",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightNode highlights DOM node with given ID or with the given JavaScript object wrapper. Either
nodeID or objectID must be specified.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
*/
func (Overlay) HighlightNode(
	socket *sock.Socket,
	params *overlay.HighlightNodeParams,
) error {
	command := &sock.Command{
		Method: "Overlay.highlightNode",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightQuad highlights given quad. Coordinates are absolute with respect to the main frame
viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
*/
func (Overlay) HighlightQuad(
	socket *sock.Socket,
	params *overlay.HighlightQuadParams,
) error {
	command := &sock.Command{
		Method: "Overlay.highlightQuad",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightRect highlights given rectangle. Coordinates are absolute with respect to the main frame
viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
*/
func (Overlay) HighlightRect(
	socket *sock.Socket,
	params *overlay.HighlightRectParams,
) error {
	command := &sock.Command{
		Method: "Overlay.highlightRect",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetInspectMode enters the 'inspect' mode. In this mode, elements that user is hovering over are
highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
func (Overlay) SetInspectMode(
	socket *sock.Socket,
	params *overlay.SetInspectModeParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setInspectMode",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPausedInDebuggerMessage sets the paused message

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
func (Overlay) SetPausedInDebuggerMessage(
	socket *sock.Socket,
	params *overlay.SetPausedInDebuggerMessageParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setPausedInDebuggerMessage",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowDebugBorders requests that backend shows debug borders on layers.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
*/
func (Overlay) SetShowDebugBorders(
	socket *sock.Socket,
	params *overlay.SetShowDebugBordersParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setShowDebugBorders",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowFPSCounter requests that backend shows the FPS counter.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
*/
func (Overlay) SetShowFPSCounter(
	socket *sock.Socket,
	params *overlay.SetShowFPSCounterParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setShowFPSCounter",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowPaintRects that backend shows paint rectangles.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
*/
func (Overlay) SetShowPaintRects(
	socket *sock.Socket,
	params *overlay.SetShowPaintRectsParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setShowPaintRects",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowScrollBottleneckRects requests that backend shows scroll bottleneck rects.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
*/
func (Overlay) SetShowScrollBottleneckRects(
	socket *sock.Socket,
	params *overlay.SetShowScrollBottleneckRectsParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setShowScrollBottleneckRects",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowViewportSizeOnResize paints viewport size upon main frame resize.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowViewportSizeOnResize
*/
func (Overlay) SetShowViewportSizeOnResize(
	socket *sock.Socket,
	params *overlay.SetShowViewportSizeOnResizeParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setShowViewportSizeOnResize",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetSuspended sets the suspended state

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setSuspended
*/
func (Overlay) SetSuspended(
	socket *sock.Socket,
	params *overlay.SetSuspendedParams,
) error {
	command := &sock.Command{
		Method: "Overlay.setSuspended",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnInspectNodeRequested adds a handler to the Overlay.inspectNodeRequested event.
Overlay.inspectNodeRequested fires when the node should be inspected. This happens after call to
`setInspectMode` or when user manually inspects an element.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-inspectNodeRequested
*/
func (Overlay) OnInspectNodeRequested(
	socket *sock.Socket,
	callback func(event *overlay.InspectNodeRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Overlay.inspectNodeRequested",
		func(name string, params []byte) {
			event := &overlay.InspectNodeRequestedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnNodeHighlightRequested adds a handler to the Overlay.nodeHighlightRequested event.
Overlay.nodeHighlightRequested fires when the node should be highlighted. This happens after call to
`setInspectMode`.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-nodeHighlightRequested
*/
func (Overlay) OnNodeHighlightRequested(
	socket *sock.Socket,
	callback func(event *overlay.NodeHighlightRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Overlay.nodeHighlightRequested",
		func(name string, params []byte) {
			event := &overlay.NodeHighlightRequestedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScreenshotRequested adds a handler to the Overlay.screenshotRequested event.
Overlay.screenshotRequested fires when user asks to capture screenshot of some area on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-screenshotRequested
*/
func (Overlay) OnScreenshotRequested(
	socket *sock.Socket,
	callback func(event *overlay.ScreenshotRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Overlay.screenshotRequested",
		func(name string, params []byte) {
			event := &overlay.ScreenshotRequestedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
