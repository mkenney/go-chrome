package overlay

import (
	"github.com/mkenney/go-chrome/protocol/dom"
	"github.com/mkenney/go-chrome/protocol/page"
	"github.com/mkenney/go-chrome/protocol/runtime"
)

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
SetInspectModeParams represents Overlay.setInspectMode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
type SetInspectModeParams struct {
	// Set an inspection mode.
	Mode InspectMode `json:"mode"`

	// Optional. A descriptor for the highlight appearance of hovered-over nodes.
	// May be omitted if `enabled == false`.
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"`
}

/*
SetPausedInDebuggerMessageParams represents Overlay.setPausedInDebuggerMessage parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
type SetPausedInDebuggerMessageParams struct {
	// Optional. The message to display, also triggers resume and step over
	// controls.
	Message string `json:"message,omitempty"`
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
SetShowFPSCounterParams represents Overlay.setShowFPSCounter parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
*/
type SetShowFPSCounterParams struct {
	// True for showing paint rectangles.
	Show bool `json:"show"`
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
SetShowScrollBottleneckRectsParams represents Overlay.setShowScrollBottleneckRects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
*/
type SetShowScrollBottleneckRectsParams struct {
	// True for showing scroll bottleneck rects.
	Show bool `json:"show"`
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
SetSuspendedParams represents Overlay.setSuspended parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setSuspended
*/
type SetSuspendedParams struct {
	// Whether overlay should be suspended and not consume any resources until
	// resumed.
	Suspended bool `json:"suspended"`
}
