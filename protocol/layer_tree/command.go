package layerTree

import (
	"github.com/mkenney/go-chrome/protocol/dom"
)

/*
CompositingReasonsParams represents LayerTree.compositingReasons parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
*/
type CompositingReasonsParams struct {
	// The ID of the layer for which we want to get the reasons it was
	// composited.
	LayerID LayerID `json:"layerId"`
}

/*
CompositingReasonsResult represents the result of calls to LayerTree.compositingReasons.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
*/
type CompositingReasonsResult struct {
	// A list of strings specifying reasons for the given layer to become
	// composited.
	CompositingReasons []string `json:"compositingReasons"`
}

/*
LoadSnapshotParams represents LayerTree.loadSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
type LoadSnapshotParams struct {
	// An array of tiles composing the snapshot.
	Tiles []PictureTile `json:"tiles"`
}

/*
LoadSnapshotResult represents the result of calls to LayerTree.loadSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
type LoadSnapshotResult struct {
	// The ID of the snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
MakeSnapshotParams represents LayerTree.makeSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
*/
type MakeSnapshotParams struct {
	// The ID of the layer.
	LayerID LayerID `json:"layerId"`
}

/*
MakeSnapshotResult represents the result of calls to LayerTree.makeSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
*/
type MakeSnapshotResult struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
ProfileSnapshotParams represents LayerTree.profileSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
*/
type ProfileSnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`

	// Optional. The maximum number of times to replay the snapshot (1, if not
	// specified).
	MinRepeatCount int `json:"minRepeatCount,omitempty"`

	// Optional. The minimum duration (in seconds) to replay the snapshot.
	MinDuration float64 `json:"minDuration,omitempty"`

	// Optional. The clip rectangle to apply when replaying the snapshot.
	ClipRect *dom.Rect `json:"clipRect,omitempty"`
}

/*
ProfileSnapshotResult represents the result of calls to LayerTree.profileSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
*/
type ProfileSnapshotResult struct {
	// The array of paint profiles, one per run.
	Timings []PaintProfile `json:"timings"`
}

/*
ReleaseSnapshotParams represents LayerTree.releaseSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-releaseSnapshot
*/
type ReleaseSnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
ReplaySnapshotParams represents LayerTree.replaySnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
type ReplaySnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`

	// Optional. The first step to replay from (replay from the very start if
	// not specified).
	FromStep int `json:"fromStep,omitempty"`

	// Optional. The last step to replay to (replay till the end if not specified).
	ToStep int `json:"toStep,omitempty"`

	// Optional. The scale to apply while replaying (defaults to 1).
	Scale float64 `json:"scale,omitempty"`
}

/*
ReplaySnapshotResult represents the result of calls to LayerTree.replaySnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
type ReplaySnapshotResult struct {
	// A data: URL for resulting image.
	DataURL string `json:"dataURL"`
}

/*
SnapshotCommandLogParams represents LayerTree.snapshotCommandLog parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
*/
type SnapshotCommandLogParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
SnapshotCommandLogResult represents the result of calls to LayerTree.snapshotCommandLog.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
*/
type SnapshotCommandLogResult struct {
	// The array of canvas function calls.
	CommandLog []map[string]string `json:"commandLog"`
}
