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
}

/*
NodeHighlightRequestedEvent represents Overlay.nodeHighlightRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-nodeHighlightRequested
*/
type NodeHighlightRequestedEvent struct {
	// ID of the node to highlight.
	NodeID dom.NodeID `json:"nodeId"`
}

/*
ScreenshotRequestedEvent represents Overlay.screenshotRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-screenshotRequested
*/
type ScreenshotRequestedEvent struct {
	// Viewport to capture, in CSS.
	Viewport *page.Viewport `json:"viewport"`
}
