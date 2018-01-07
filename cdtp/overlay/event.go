package overlay

import (
	"github.com/mkenney/go-chrome/cdtp/dom"
	"github.com/mkenney/go-chrome/cdtp/page"
)

/*
InspectNodeRequestedEvent represents Overlay.inspectNodeRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-inspectNodeRequested
*/
type InspectNodeRequestedEvent struct {
	// ID of the node to inspect.
	BackendNodeID dom.BackendNodeID `json:"backendNodeId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
NodeHighlightRequestedEvent represents Overlay.nodeHighlightRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-nodeHighlightRequested
*/
type NodeHighlightRequestedEvent struct {
	// ID of the node to highlight.
	NodeID dom.NodeID `json:"nodeId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ScreenshotRequestedEvent represents Overlay.screenshotRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-screenshotRequested
*/
type ScreenshotRequestedEvent struct {
	// Viewport to capture, in CSS.
	Viewport *page.Viewport `json:"viewport"`

	// Error information related to this event
	Err error `json:"-"`
}
