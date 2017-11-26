package LayerTree

import (
	DOM "app/chrome/protocol/dom"
)

/*
LayerID is a unique Layer identifier.
*/
type LayerID string

/*
SnapshotID is a unique snapshot identifier.
*/
type SnapshotID string

/*
ScrollRect is a rectangle where scrolling happens on the main thread.
*/
type ScrollRect struct {
	// Rectangle itself.
	Rect *DOM.Rect `json:"rect"`

	// Reason for rectangle to force scrolling on the main thread Allowed values: RepaintsOnScroll, \
	// TouchEventHandler, WheelEventHandler.
	Type string `json:"type"`
}

/*
StickyPositionConstraint is sticky position constraints.
*/
type StickyPositionConstraint struct {
	// Layout rectangle of the sticky element before being shifted.
	StickyBoxRect *DOM.Rect `json:"stickyBoxRect"`

	// Layout rectangle of the containing block of the sticky element.
	ContainingBlockRect *DOM.Rect `json:"containingBlockRect"`

	// Optional. The nearest sticky layer that shifts the sticky box.
	NearestLayerShiftingStickyBox LayerID `json:"nearestLayerShiftingStickyBox,omitempty"`

	// Optional. The nearest sticky layer that shifts the containing block.
	NearestLayerShiftingContainingBlock LayerID `json:"nearestLayerShiftingContainingBlock,omitempty"`
}

/*
PictureTile is a serialized fragment of layer picture along with its offset within the layer.
*/
type PictureTile struct {
	// Offset from owning layer left boundary.
	X int `json:"x"`

	// Offset from owning layer top boundary.
	Y int `json:"y"`

	// Base64-encoded snapshot data.
	Picture string `json:"picture"`
}

/*
Layer is information about a compositing layer.
*/
type Layer struct {
	// The unique ID for this layer.
	LayerID LayerID `json:"layerId"`

	// Optional. The ID of parent (not present for root).
	ParentLayerID LayerID `json:"parentLayerId,omitempty"`

	// Optional. The backend ID for the node associated with this layer.
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId,omitempty"`

	// Offset from parent layer, X coordinate.
	OffsetX int `json:"offsetX"`

	// Offset from parent layer, Y coordinate.
	OffsetY int `json:"offsetY"`

	// Layer width.
	Width int `json:"width"`

	// Layer height.
	Height int `json:"height"`

	// Optional. Transformation matrix for layer, default is identity matrix.
	Transform []int `json:"transform,omitempty"`

	// Optional. Transform anchor point X, absent if no transform specified.
	AnchorX int `json:"anchorX,omitempty"`

	// Optional. Transform anchor point Y, absent if no transform specified.
	AnchorY int `json:"anchorY,omitempty"`

	// Optional. Transform anchor point Z, absent if no transform specified.
	AnchorZ int `json:"anchorZ,omitempty"`

	// Indicates how many time this layer has painted.
	PaintCount int `json:"paintCount"`

	// Indicates whether this layer hosts any content, rather than being used for
	// transform/scrolling purposes only.
	DrawsContent bool `json:"drawsContent"`

	// Optional. Set if layer is not visible.
	Invisible bool `json:"invisible,omitempty"`

	// Optional. Rectangles scrolling on main thread only.
	ScrollRects []*ScrollRect `json:"scrollRects,omitempty"`

	// Optional. Sticky position constraint information.
	StickyPositionConstraint *StickyPositionConstraint `json:"stickyPositionConstraint,omitempty"`
}

/*
PaintProfile is an array of timings, one per paint step.
*/
type PaintProfile []interface{}
