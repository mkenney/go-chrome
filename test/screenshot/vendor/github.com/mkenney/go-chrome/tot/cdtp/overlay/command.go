package overlay

import (
	"github.com/mkenney/go-chrome/tot/cdtp/dom"
	"github.com/mkenney/go-chrome/tot/cdtp/page"
	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

/*
DisableResult represents the result of calls to Overlay.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Overlay.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetHighlightObjectForTestParams represents Overlay.getHighlightObjectForTest parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
*/
type GetHighlightObjectForTestParams struct {
	// ID of the node to get highlight object for.
	NodeID dom.NodeID `json:"nodeId"`
}

/*
GetHighlightObjectForTestResult represents the result of calls to Overlay.getHighlightObjectForTest.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
*/
type GetHighlightObjectForTestResult struct {
	// Highlight data for the node.
	Highlight map[string]string `json:"highlight"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
HideHighlightResult represents the result of calls to Overlay.hideHighlight.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-hideHighlight
*/
type HideHighlightResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
HighlightFrameParams represents Overlay.highlightFrame parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
*/
type HighlightFrameParams struct {
	// Identifier of the frame to highlight.
	FrameID page.FrameID `json:"frameId"`

	// Optional. The content box highlight fill color (default: transparent).
	ContentColor *dom.RGBA `json:"contentColor,omitempty"`

	// Optional. The content box highlight outline color (default: transparent).
	ContentOutlineColor *dom.RGBA `json:"contentOutlineColor,omitempty"`
}

/*
HighlightFrameResult represents the result of calls to Overlay.highlightFrame.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
*/
type HighlightFrameResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
HighlightNodeParams represents Overlay.highlightNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
*/
type HighlightNodeParams struct {
	// A descriptor for the highlight appearance.
	HighlightConfig *HighlightConfig `json:"highlightConfig"`

	// Optional. Identifier of the node to highlight.
	NodeID dom.NodeID `json:"nodeId,omitempty"`

	// Optional. Identifier of the backend node to highlight.
	BackendNodeID dom.BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node to be highlighted.
	ObjectID runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
HighlightNodeResult represents the result of calls to Overlay.highlightNode.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
*/
type HighlightNodeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
HighlightQuadParams represents Overlay.highlightQuad parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
*/
type HighlightQuadParams struct {
	// Quad to highlight.
	Quad dom.Quad `json:"quad"`

	// Optional. The highlight fill color (default: transparent).
	Color *dom.RGBA `json:"color,omitempty"`

	// Optional. The highlight outline color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"`
}

/*
HighlightQuadResult represents the result of calls to Overlay.highlightQuad.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
*/
type HighlightQuadResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
HighlightRectParams represents Overlay.highlightRect parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
*/
type HighlightRectParams struct {
	// X coordinate.
	X int `json:"x"`

	// Y coordinate.
	Y int `json:"y"`

	// Rectangle width.
	Width int `json:"width"`

	// Rectangle height.
	Height int `json:"height"`

	// Optional. The highlight fill color (default: transparent).
	Color *dom.RGBA `json:"color,omitempty"`

	// Optional. The highlight outline color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"`
}

/*
HighlightRectResult represents the result of calls to Overlay.highlightRect.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
*/
type HighlightRectResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetInspectModeParams represents Overlay.setInspectMode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
type SetInspectModeParams struct {
	// Set an inspection mode. Allowed values:
	//	- InspectMode.SearchForNode
	//	- InspectMode.SearchForUAShadowDOM
	//	- InspectMode.None
	Mode InspectModeEnum `json:"mode"`

	// Optional. A descriptor for the highlight appearance of hovered-over nodes.
	// May be omitted if `enabled == false`.
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"`
}

/*
SetInspectModeResult represents the result of calls to Overlay.setInspectMode.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
type SetInspectModeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetPausedInDebuggerMessageParams represents Overlay.setPausedInDebuggerMessage
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
type SetPausedInDebuggerMessageParams struct {
	// Optional. The message to display, also triggers resume and step over
	// controls.
	Message string `json:"message,omitempty"`
}

/*
SetPausedInDebuggerMessageResult represents the result of calls to
Overlay.setPausedInDebuggerMessage.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
type SetPausedInDebuggerMessageResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetShowDebugBordersParams represents Overlay.setShowDebugBorders parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
*/
type SetShowDebugBordersParams struct {
	// True for showing debug borders.
	Show bool `json:"show"`
}

/*
SetShowDebugBordersResult represents the result of calls to Overlay.setShowDebugBorders.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
*/
type SetShowDebugBordersResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetShowFPSCounterParams represents Overlay.setShowFPSCounter parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
*/
type SetShowFPSCounterParams struct {
	// True for showing paint rectangles.
	Show bool `json:"show"`
}

/*
SetShowFPSCounterResult represents the result of calls to Overlay.setShowFPSCounter.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
*/
type SetShowFPSCounterResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetShowPaintRectsParams represents Overlay.setShowPaintRects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
*/
type SetShowPaintRectsParams struct {
	// True for showing paint rectangles.
	Result bool `json:"result"`
}

/*
SetShowPaintRectsResult represents the result of calls to Overlay.setShowPaintRects.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
*/
type SetShowPaintRectsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetShowScrollBottleneckRectsParams represents Overlay.setShowScrollBottleneckRects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
*/
type SetShowScrollBottleneckRectsParams struct {
	// True for showing scroll bottleneck rects.
	Show bool `json:"show"`
}

/*
SetShowScrollBottleneckRectsResult represents the result of calls to
Overlay.SetShowScrollBottleneckRects.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-SetShowScrollBottleneckRects
*/
type SetShowScrollBottleneckRectsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetShowViewportSizeOnResizeParams represents Overlay.setShowViewportSizeOnResize parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowViewportSizeOnResize
*/
type SetShowViewportSizeOnResizeParams struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

/*
SetShowViewportSizeOnResizeResult represents the result of calls to
Overlay.SetShowViewportSizeOnResize.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-SetShowViewportSizeOnResize
*/
type SetShowViewportSizeOnResizeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetSuspendedParams represents Overlay.setSuspended parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setSuspended
*/
type SetSuspendedParams struct {
	// Whether overlay should be suspended and not consume any resources until
	// resumed.
	Suspended bool `json:"suspended"`
}

/*
SetSuspendedResult represents the result of calls to Overlay.setSuspended.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setSuspended
*/
type SetSuspendedResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
