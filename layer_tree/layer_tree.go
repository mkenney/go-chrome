package LayerTree

import (
	DOM "app/chrome/dom"
)

/*
CompositingReasonsParams represents LayerTree.compositingReasons parameters.
*/
type CompositingReasonsParams struct {
	// The ID of the layer for which we want to get the reasons it was composited.
	LayerID LayerID `json:"layerId"`
}

/*
LoadSnapshotParams represents LayerTree.loadSnapshot parameters.
*/
type LoadSnapshotParams struct {
	// An array of tiles composing the snapshot.
	Tiles []PictureTile `json:"tiles"`
}

/*
MakeSnapshotParams represents LayerTree.makeSnapshot parameters.
*/
type MakeSnapshotParams struct {
	// The ID of the layer.
	LayerID LayerID `json:"layerId"`
}

/*
ProfileSnapshotParams represents LayerTree.profileSnapshot parameters.
*/
type ProfileSnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`

	// Optional. The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount int `json:"minRepeatCount,omitempty"`

	// Optional. The minimum duration (in seconds) to replay the snapshot.
	MinDuration float64 `json:"minDuration,omitempty"`

	// Optional. The clip rectangle to apply when replaying the snapshot.
	ClipRect DOM.Rect `json:"clipRect,omitempty"`
}

/*
ReleaseSnapshotParams represents LayerTree.releaseSnapshot parameters.
*/
type ReleaseSnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
ReplaySnapshotParams represents LayerTree.replaySnapshot parameters.
*/
type ReplaySnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`

	// Optional. The first step to replay from (replay from the very start if not specified).
	FromStep int `json:"fromStep,omitempty"`

	// Optional. The last step to replay to (replay till the end if not specified).
	ToStep int `json:"toStep,omitempty"`

	// Optional. The scale to apply while replaying (defaults to 1).
	Scale float64 `json:"scale,omitempty"`
}

/*
ReplaySnapshotParams represents LayerTree.replaySnapshot parameters.
*/
type ReplaySnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
LayerPaintedEvent represents DOM.layerPainted event data.
*/
type LayerPaintedEvent struct {
	// The ID of the painted layer.
	LayerID LayerID `json:"layerId"`

	// Clip rectangle.
	Clip DOM.Rect `json:"clip"`
}

/*
LayerPaintedEvent represents DOM.layerPainted event data.
*/
type LayerPaintedEvent struct {
	// Layer tree, absent if not in the comspositing mode.
	Layers []Layer `json:"layers"`
}

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
	X float64 `json:"x"`

	// Offset from owning layer top boundary.
	Y float64 `json:"y"`

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
	OffsetX float64 `json:"offsetX"`

	// Offset from parent layer, Y coordinate.
	OffsetY float64 `json:"offsetY"`

	// Layer width.
	Width float64 `json:"width"`

	// Layer height.
	Height float64 `json:"height"`

	// Optional. Transformation matrix for layer, default is identity matrix.
	Transform []float64 `json:"transform,omitempty"`

	// Optional. Transform anchor point X, absent if no transform specified.
	AnchorX float64 `json:"anchorX,omitempty"`

	// Optional. Transform anchor point Y, absent if no transform specified.
	AnchorY float64 `json:"anchorY,omitempty"`

	// Optional. Transform anchor point Z, absent if no transform specified.
	AnchorZ float64 `json:"anchorZ,omitempty"`

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
