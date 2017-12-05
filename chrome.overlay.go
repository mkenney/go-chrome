package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Overlay - https://chromedevtools.github.io/devtools-protocol/tot/Overlay/
Provides various functionality related to drawing atop the inspected page. EXPERIMENTAL
*/
type Overlay struct{}

/*
Disable disables domain notifications.
*/
func (Overlay) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Overlay.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables domain notifications.
*/
func (Overlay) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Overlay.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetHighlightObjectForTest is for testing.
*/
func (Overlay) GetHighlightObjectForTest(socket *Socket, params *overlay.GetHighlightObjectForTestParams) error {
	command := &protocol.Command{
		method: "Overlay.getHighlightObjectForTest",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HideHighlight hides any highlight.
*/
func (Overlay) HideHighlight(socket *Socket) error {
	command := &protocol.Command{
		method: "Overlay.hideHighlight",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightFrame highlights owner element of the frame with given ID.
*/
func (Overlay) HighlightFrame(socket *Socket, params *overlay.HighlightFrameParams) error {
	command := &protocol.Command{
		method: "Overlay.highlightFrame",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightNode highlights DOM node with given ID or with the given JavaScript object wrapper. Either
nodeID or objectID must be specified.
*/
func (Overlay) HighlightNode(socket *Socket, params *overlay.HighlightNodeParams) error {
	command := &protocol.Command{
		method: "Overlay.highlightNode",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightQuad highlights given quad. Coordinates are absolute with respect to the main frame
viewport.
*/
func (Overlay) HighlightQuad(socket *Socket, params *overlay.HighlightQuadParams) error {
	command := &protocol.Command{
		method: "Overlay.highlightQuad",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HighlightRect highlights given rectangle. Coordinates are absolute with respect to the main frame
viewport.
*/
func (Overlay) HighlightRect(socket *Socket, params *overlay.HighlightRectParams) error {
	command := &protocol.Command{
		method: "Overlay.highlightRect",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetInspectMode enters the 'inspect' mode. In this mode, elements that user is hovering over are
highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
*/
func (Overlay) SetInspectMode(socket *Socket, params *overlay.SetInspectModeParams) error {
	command := &protocol.Command{
		method: "Overlay.setInspectMode",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPausedInDebuggerMessage sets the paused message
*/
func (Overlay) SetPausedInDebuggerMessage(socket *Socket, params *overlay.SetPausedInDebuggerMessageParams) error {
	command := &protocol.Command{
		method: "Overlay.setPausedInDebuggerMessage",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowDebugBorders requests that backend shows debug borders on layers.
*/
func (Overlay) SetShowDebugBorders(socket *Socket, params *overlay.SetShowDebugBordersParams) error {
	command := &protocol.Command{
		method: "Overlay.setShowDebugBorders",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowFPSCounter requests that backend shows the FPS counter.
*/
func (Overlay) SetShowFPSCounter(socket *Socket, params *overlay.SetShowFPSCounterParams) error {
	command := &protocol.Command{
		method: "Overlay.setShowFPSCounter",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowScrollBottleneckRects requests that backend shows scroll bottleneck rects.
*/
func (Overlay) SetShowScrollBottleneckRects(socket *Socket, params *overlay.SetShowScrollBottleneckRectsParams) error {
	command := &protocol.Command{
		method: "Overlay.setShowScrollBottleneckRects",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetShowViewportSizeOnResize paints viewport size upon main frame resize.
*/
func (Overlay) SetShowViewportSizeOnResize(socket *Socket, params *overlay.SetShowViewportSizeOnResizeParams) error {
	command := &protocol.Command{
		method: "Overlay.setShowViewportSizeOnResize",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetSuspended sets the suspended state
*/
func (Overlay) SetSuspended(socket *Socket, params *overlay.SetSuspendedParams) error {
	command := &protocol.Command{
		method: "Overlay.setSuspended",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnInspectNodeRequested adds a handler to the Overlay.inspectNodeRequested event.
Overlay.inspectNodeRequested fires when the node should be inspected. This happens after call to
`setInspectMode` or when user manually inspects an element.
*/
func (Overlay) OnInspectNodeRequested(socket *Socket, callback func(event *overlay.InspectNodeRequestedEvent)) {
	handler := protocol.NewEventHandler(
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
*/
func (Overlay) OnNodeHighlightRequested(socket *Socket, callback func(event *overlay.NodeHighlightRequestedEvent)) {
	handler := protocol.NewEventHandler(
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
*/
func (Overlay) OnScreenshotRequested(socket *Socket, callback func(event *overlay.ScreenshotRequestedEvent)) {
	handler := protocol.NewEventHandler(
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
