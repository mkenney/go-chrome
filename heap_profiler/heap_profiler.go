package HeapProfiler

import (
	"github.com/mkenney/go-chrome/runtime"
)

/*
AddInspectedHeapObjectParams represents HeapProfiler.addInspectedHeapObject parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
*/
type AddInspectedHeapObjectParams struct {
	// Heap snapshot object ID to be accessible by means of $x command line API.
	HeapObjectID HeapSnapshotObjectID `json:"heapObjectId"`
}

/*
GetHeapObjectIDParams represents HeapProfiler.getHeapObjectId parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
*/
type GetHeapObjectIDParams struct {
	// Identifier of the object to get heap object ID for.
	ObjectID Runtime.RemoteObjectID `json:"objectId"`
}

/*
GetHeapObjectIDResult represents the result of calls to HeapProfiler.getHeapObjectId.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
*/
type GetHeapObjectIDResult struct {
	// ID of the heap snapshot object corresponding to the passed remote object id.
	HeapSnapshotObjectID HeapSnapshotObjectID `json:"heapSnapshotObjectId"`
}

/*
GetObjectByHeapObjectIDParams represents HeapProfiler.getObjectByHeapObjectId parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
*/
type GetObjectByHeapObjectIDParams struct {
	// desc.
	ObjectID HeapSnapshotObjectID `json:"objectId"`

	// Optional. Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
GetObjectByHeapObjectIDResult represents the result of calls to
HeapProfiler.getObjectByHeapObjectId.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
*/
type GetObjectByHeapObjectIDResult struct {
	// Evaluation result.
	Result Runtime.RemoteObject `json:"result"`
}

/*
GetSamplingProfileParams represents HeapProfiler.getSamplingProfile parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
*/
type GetSamplingProfileParams struct {
	// Return the sampling profile being collected.
	Profile SamplingHeapProfile `json:"profile"`
}

/*
StartSamplingParams represents HeapProfiler.startSampling parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
*/
type StartSamplingParams struct {
	// Optional. Average sample interval in bytes. Poisson distribution is used for the intervals.
	// The default value is 32768 bytes.
	SamplingInterval int `json:"samplingInterval,omitempty"`
}

/*
StartTrackingHeapObjectsParams represents HeapProfiler.startTrackingHeapObjects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
*/
type StartTrackingHeapObjectsParams struct {
	// Optional.
	TrackAllocations bool `json:"trackAllocations,omitempty"`
}

/*
StopSamplingParams represents HeapProfiler.stopSampling parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
*/
type StopSamplingParams struct {
	// Recorded sampling heap profile.
	Profile SamplingHeapProfile `json:"profile"`
}

/*
StopTrackingHeapObjectsParams represents HeapProfiler.stopTrackingHeapObjects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
*/
type StopTrackingHeapObjectsParams struct {
	// Optional. If true 'reportHeapSnapshotProgress' events will be generated while snapshot is
	// being taken when the tracking is stopped.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

/*
TakeHeapSnapshotParams represents HeapProfiler.takeHeapSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
*/
type TakeHeapSnapshotParams struct {
	// Optional. If true 'reportHeapSnapshotProgress' events will be generated while snapshot is
	// being taken.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

/*
AddHeapSnapshotChunkEvent represents DOM.addHeapSnapshotChunk event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-addHeapSnapshotChunk
*/
type AddHeapSnapshotChunkEvent struct {
	Chunk string `json:"chunk"`
}

/*
HeapStatsUpdateEvent represents DOM.heapStatsUpdate event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-heapStatsUpdate
*/
type HeapStatsUpdateEvent struct {
	// An array of triplets. Each triplet describes a fragment. The first integer is the fragment
	// index, the second integer is a total count of objects for the fragment, the third integer is
	// a total size of the objects for the fragment.
	StatsUpdate []int `json:"statsUpdate"`
}

/*
LastSeenObjectIDEvent represents DOM.lastSeenObjectId event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-lastSeenObjectId
*/
type LastSeenObjectIDEvent struct {
	//
	LastSeenObjectID int `json:"lastSeenObjectId"`

	//
	Timestamp int `json:"timestamp"`
}

/*
ReportHeapSnapshotProgressEvent represents DOM.reportHeapSnapshotProgress event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-reportHeapSnapshotProgress
*/
type ReportHeapSnapshotProgressEvent struct {
	//
	Done int `json:"done"`

	//
	Total int `json:"total"`

	// Optional.
	Finished bool `json:"finished,omitempty"`
}

/*
ResetProfilesEvent represents DOM.resetProfiles event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-resetProfiles
*/
type ResetProfilesEvent struct{}

/*
HeapSnapshotObjectID is the heap snapshot object id.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#type-HeapSnapshotObjectId
*/
type HeapSnapshotObjectID string

/*
SamplingHeapProfileNode is the sampling Heap Profile node. Holds callsite information, allocation
statistics and child nodes.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#type-SamplingHeapProfileNode
*/
type SamplingHeapProfileNode struct {
	// Function location.
	CallFrame Runtime.CallFrame `json:"callFrame"`

	// Allocations size in bytes for the node excluding children.
	SelfSize int `json:"selfSize"`

	// Child nodes.
	Children []*SamplingHeapProfileNode `json:"children"`
}

/*
SamplingHeapProfile represents a heap sample profile

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#type-SamplingHeapProfile
*/
type SamplingHeapProfile struct {
	Head SamplingHeapProfileNode `json:"head"`
}
