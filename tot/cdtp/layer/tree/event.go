package tree

import (
	"github.com/mkenney/go-chrome/tot/cdtp/dom"
)

/*
LayerPaintedEvent represents LayerTree.layerPainted event data.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerPainted
*/
type LayerPaintedEvent struct {
	// The ID of the painted layer.
	LayerID LayerID `json:"layerId"`

	// Clip rectangle.
	Clip *dom.Rect `json:"clip"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
DidChangeEvent represents LayerTree.layerTreeDidChange event data.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerTreeDidChange
*/
type DidChangeEvent struct {
	// Optional. Layer tree, absent if not in the comspositing mode.
	Layers []*Layer `json:"layers,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}
