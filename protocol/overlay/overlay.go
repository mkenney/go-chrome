/*
Package overlay provides type definitions for use with the Chrome Overlay protocol

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/
*/
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
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	ContentColor interface{} `json:"contentColor,omitempty"`

	// Optional. The content box highlight outline color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	ContentOutlineColor interface{} `json:"contentOutlineColor,omitempty"`
}

/*
HighlightNodeParams represents Overlay.highlightNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
*/
type HighlightNodeParams struct {
	// A descriptor for the highlight appearance.
	HighlightConfig HighlightConfig `json:"highlightConfig"`

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
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	Color interface{} `json:"color,omitempty"`

	// Optional. The highlight outline color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	OutlineColor interface{} `json:"outlineColor,omitempty"`
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
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	Color interface{} `json:"color,omitempty"`

	// Optional. The highlight outline color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	OutlineColor interface{} `json:"outlineColor,omitempty"`
}

/*
SetInspectModeParams represents Overlay.setInspectMode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
type SetInspectModeParams struct {
	// Set an inspection mode.
	Mode InspectMode `json:"mode"`

	// Optional. A descriptor for the highlight appearance of hovered-over nodes. May be omitted if
	// `enabled == false`.
	//
	// This is an instance of HighlightConfig, but that doesn't omitempty correctly so it must be
	// added manually.
	HighlightConfig interface{} `json:"highlightConfig,omitempty"`
}

/*
SetPausedInDebuggerMessageParams represents Overlay.setPausedInDebuggerMessage parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
type SetPausedInDebuggerMessageParams struct {
	// Optional. The message to display, also triggers resume and step over controls.
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
	// Whether overlay should be suspended and not consume any resources until resumed.
	Suspended bool `json:"suspended"`
}

/*
InspectNodeRequestedEvent represents Overlay.inspectNodeRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-inspectNodeRequested
*/
type InspectNodeRequestedEvent struct {
	// ID of the node to inspect.
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId"`
}

/*
NodeHighlightRequestedEvent represents Overlay.nodeHighlightRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-nodeHighlightRequested
*/
type NodeHighlightRequestedEvent struct {
	// ID of the node to highlight.
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
ScreenshotRequestedEvent represents Overlay.screenshotRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-screenshotRequested
*/
type ScreenshotRequestedEvent struct {
	// Viewport to capture, in CSS.
	Viewport page.Viewport `json:"viewport"`
}

/*
HighlightConfig is configuration data for the highlighting of page elements.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-HighlightConfig
*/
type HighlightConfig struct {
	// Optional. Whether the node info tooltip should be shown (default: false).
	ShowInfo bool `json:"showInfo,omitempty"`

	// Optional. Whether the rulers should be shown (default: false).
	ShowRulers bool `json:"showRulers,omitempty"`

	// Optional. Whether the extension lines from node to the rulers should be shown (default:
	// false).
	ShowExtensionLines bool `json:"showExtensionLines,omitempty"`

	// Optional. Display as material.
	DisplayAsMaterial bool `json:"displayAsMaterial,omitempty"`

	// Optional. The content box highlight fill color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	ContentColor interface{} `json:"contentColor,omitempty"`

	// Optional. The padding highlight fill color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	PaddingColor interface{} `json:"paddingColor,omitempty"`

	// Optional. The border highlight fill color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	BorderColor interface{} `json:"borderColor,omitempty"`

	// Optional. The margin highlight fill color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	MarginColor interface{} `json:"marginColor,omitempty"`

	// Optional. The event target element highlight fill color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	EventTargetColor interface{} `json:"eventTargetColor,omitempty"`

	// Optional. The shape outside fill color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	ShapeColor interface{} `json:"shapeColor,omitempty"`

	// Optional. The shape margin fill color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	ShapeMarginColor interface{} `json:"shapeMarginColor,omitempty"`

	// Optional. Selectors to highlight relevant nodes.
	SelectorList string `json:"selectorList,omitempty"`

	// Optional. The grid layout color (default: transparent).
	//
	// This is an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be added
	// manually.
	CSSGridColor interface{} `json:"cssGridColor,omitempty"`
}

/*
InspectMode is the inspect mode

ALLOWED VALUES
	- searchForNode
	- searchForUAShadowDOM
	- none

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-InspectMode
*/
type InspectMode string
