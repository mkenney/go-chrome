package Overlay

import (
	DOM "app/chrome/dom"
	Page "app/chrome/page"
	Runtime "app/chrome/runtime"
	"fmt"
)

/*
GetHighlightObjectForTestParams represents Overlay.getHighlightObjectForTest parameters.
*/
type GetHighlightObjectForTestParams struct {
	// ID of the node to get highlight object for.
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetHighlightObjectForTestResult represents the result of calls to Overlay.getHighlightObjectForTest.
*/
type GetHighlightObjectForTestResult struct {
	// Highlight data for the node.
	Highlight map[string]string `json:"highlight"`
}

/*
HighlightFrameParams represents Overlay.highlightFrame parameters.
*/
type HighlightFrameParams struct {
	// Identifier of the frame to highlight.
	FrameID Page.FrameID `json:"frameId"`

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
*/
type HighlightNodeParams struct {
	// A descriptor for the highlight appearance.
	HighlightConfig HighlightConfig `json:"highlightConfig"`

	// Optional. Identifier of the node to highlight.
	NodeID DOM.NodeID `json:"nodeId,omitempty"`

	// Optional. Identifier of the backend node to highlight.
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node to be highlighted.
	ObjectID Runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
HighlightQuadParams represents Overlay.highlightQuad parameters.
*/
type HighlightQuadParams struct {
	// Quad to highlight.
	Quad DOM.Quad `json:"quad"`

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
*/
type SetPausedInDebuggerMessageParams struct {
	// Optional. The message to display, also triggers resume and step over controls.
	Message string `json:"message,omitempty"`
}

/*
SetShowDebugBordersParams represents Overlay.setShowDebugBorders parameters.
*/
type SetShowDebugBordersParams struct {
	// True for showing debug borders.
	Show bool `json:"show"`
}

/*
SetShowFPSCounterParams represents Overlay.setShowFPSCounter parameters.
*/
type SetShowFPSCounterParams struct {
	// True for showing paint rectangles.
	Result bool `json:"result"`
}

/*
SetShowScrollBottleneckRectsParams represents Overlay.setShowScrollBottleneckRects parameters.
*/
type SetShowScrollBottleneckRectsParams struct {
	// True for showing scroll bottleneck rects.
	Show bool `json:"show"`
}

/*
SetShowViewportSizeOnResizeParams represents Overlay.setShowViewportSizeOnResize parameters.
*/
type SetShowViewportSizeOnResizeParams struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

/*
SetSuspendedParams represents Overlay.setSuspended parameters.
*/
type SetSuspendedParams struct {
	// Whether overlay should be suspended and not consume any resources until resumed.
	Suspended bool `json:"suspended"`
}

/*
InspectNodeRequestedEvent represents Overlay.inspectNodeRequested event data.
*/
type InspectNodeRequestedEvent struct {
	// ID of the node to inspect.
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId"`
}

/*
NodeHighlightRequestedEvent represents Overlay.nodeHighlightRequested event data.
*/
type NodeHighlightRequestedEvent struct {
	// ID of the node to highlight.
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
ScreenshotRequestedEvent represents Overlay.screenshotRequested event data.
*/
type ScreenshotRequestedEvent struct {
	// Viewport to capture, in CSS.
	Viewport Page.Viewport `json:"viewport"`
}

/*
HighlightConfig is configuration data for the highlighting of page elements.
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
*/
type InspectMode int

const (
	_searchForNode InspectMode = iota
	_searchForUAShadowDOM
	_none
)

func (a InspectMode) String() string {
	if a == 0 {
		return "searchForNode"
	}
	if a == 1 {
		return "searchForUAShadowDOM"
	}
	if a == 2 {
		return "none"
	}
	panic(fmt.Errorf("Invalid InspectMode %d", a))
}
