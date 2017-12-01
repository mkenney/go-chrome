package HeapProfiler

import (
	Runtime "app/chrome/runtime"
)

/*
StartTrackingHeapObjectsParams represents HeapProfiler.startTrackingHeapObjects parameters.
*/
type StartTrackingHeapObjectsParams struct {
	// Optional.
	TrackAllocations bool `json:"trackAllocations,omitempty"`
}

/*
StopTrackingHeapObjectsParams represents HeapProfiler.stopTrackingHeapObjects parameters.
*/
type StopTrackingHeapObjectsParams struct {
	// Optional. If true 'reportHeapSnapshotProgress' events will be generated while snapshot is
	// being taken when the tracking is stopped.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

/*
TakeHeapSnapshotParams represents HeapProfiler.takeHeapSnapshot parameters.
*/
type TakeHeapSnapshotParams struct {
	// Optional. If true 'reportHeapSnapshotProgress' events will be generated while snapshot is
	// being taken.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

/*
GetObjectByHeapObjectIDParams represents HeapProfiler.getObjectByHeapObjectId parameters.
*/
type GetObjectByHeapObjectIDParams struct {
	// desc.
	ObjectID HeapSnapshotObjectID `json:"objectId"`

	// Optional. Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
AddInspectedHeapObjectParams represents HeapProfiler.addInspectedHeapObject parameters.
*/
type AddInspectedHeapObjectParams struct {
	// Heap snapshot object ID to be accessible by means of $x command line API.
	HeapObjectID HeapSnapshotObjectID `json:"heapObjectId"`
}

/*
GetHeapObjectIDParams represents HeapProfiler.getHeapObjectId parameters.
*/
type GetHeapObjectIDParams struct {
	// Identifier of the object to get heap object ID for.
	ObjectID Runtime.RemoteObjectID `json:"objectId"`
}

/*
StartSamplingParams represents HeapProfiler.startSampling parameters.
*/
type StartSamplingParams struct {
	// Optional. Average sample interval in bytes. Poisson distribution is used for the intervals.
	// The default value is 32768 bytes.
	SamplingInterval int `json:"samplingInterval,omitempty"`
}

/*
StopSamplingParams represents HeapProfiler.stopSampling parameters.
*/
type StopSamplingParams struct {
	// Recorded sampling heap profile.
	Profile SamplingHeapProfile `json:"profile"`
}

/*
GetSamplingProfileParams represents HeapProfiler.getSamplingProfile parameters.
*/
type GetSamplingProfileParams struct {
	// Return the sampling profile being collected.
	Profile SamplingHeapProfile `json:"profile"`
}

/*
AddHeapSnapshotChunkEvent represents DOM.addHeapSnapshotChunk event data.
*/
type AddHeapSnapshotChunkEvent struct {
	Chunk string `json:"chunk"`
}

/*
ResetProfilesEvent represents DOM.resetProfiles event data.
*/
type ResetProfilesEvent struct{}

/*
ReportHeapSnapshotProgressEvent represents DOM.reportHeapSnapshotProgress event data.
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
LastSeenObjectIDEvent represents DOM.lastSeenObjectId event data.
*/
type LastSeenObjectIDEvent struct {
	//
	LastSeenObjectID int `json:"lastSeenObjectId"`

	//
	Timestamp int `json:"timestamp"`
}

/*
HeapStatsUpdateEvent represents DOM.heapStatsUpdate event data.
*/
type HeapStatsUpdateEvent struct {
	// An array of triplets. Each triplet describes a fragment. The first integer is the fragment
	// index, the second integer is a total count of objects for the fragment, the third integer is
	// a total size of the objects for the fragment.
	StatsUpdate []int `json:"statsUpdate"`
}

/*
HeapSnapshotObjectID is the heap snapshot object id.
*/
type HeapSnapshotObjectID string

/*
SamplingHeapProfileNode is the sampling Heap Profile node. Holds callsite information, allocation
statistics and child nodes.
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
*/
type SamplingHeapProfile struct {
	Head SamplingHeapProfileNode `json:"head"`
}
