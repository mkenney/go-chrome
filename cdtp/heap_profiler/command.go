package heapProfiler

import "github.com/mkenney/go-chrome/cdtp/runtime"

/*
AddInspectedHeapObjectParams represents HeapProfiler.addInspectedHeapObject parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
*/
type AddInspectedHeapObjectParams struct {
	// Heap snapshot object ID to be accessible by means of $x command line API.
	HeapObjectID HeapSnapshotObjectID `json:"heapObjectId"`
}

/*
AddInspectedHeapObjectResult represents the result of calls to
HeapProfiler.AddInspectedHeapObject.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
*/
type AddInspectedHeapObjectResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CollectGarbageResult represents the result of calls to HeapProfiler.collectGarbage.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
*/
type CollectGarbageResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to HeapProfiler.disable.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to HeapProfiler.enable.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetHeapObjectIDParams represents HeapProfiler.getHeapObjectId parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
*/
type GetHeapObjectIDParams struct {
	// Identifier of the object to get heap object ID for.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
}

/*
GetHeapObjectIDResult represents the result of calls to HeapProfiler.getHeapObjectId.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
*/
type GetHeapObjectIDResult struct {
	// ID of the heap snapshot object corresponding to the passed remote object
	// id.
	HeapSnapshotObjectID HeapSnapshotObjectID `json:"heapSnapshotObjectId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetObjectByHeapObjectIDParams represents HeapProfiler.getObjectByHeapObjectId parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
*/
type GetObjectByHeapObjectIDParams struct {
	// desc.
	ObjectID HeapSnapshotObjectID `json:"objectId"`

	// Optional. Symbolic group name that can be used to release multiple
	// objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
GetObjectByHeapObjectIDResult represents the result of calls to
HeapProfiler.getObjectByHeapObjectId.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
*/
type GetObjectByHeapObjectIDResult struct {
	// Evaluation result.
	Result *runtime.RemoteObject `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetSamplingProfileParams represents HeapProfiler.getSamplingProfile parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
*/
type GetSamplingProfileParams struct {
	// Return the sampling profile being collected.
	Profile *SamplingHeapProfile `json:"profile"`
}

/*
GetSamplingProfileResult represents the result of calls to HeapProfiler.getSamplingProfile.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
*/
type GetSamplingProfileResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StartSamplingParams represents HeapProfiler.startSampling parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
*/
type StartSamplingParams struct {
	// Optional. Average sample interval in bytes. Poisson distribution is used
	// for the intervals. The default value is 32768 bytes.
	SamplingInterval int `json:"samplingInterval,omitempty"`
}

/*
StartSamplingResult represents the result of calls to HeapProfiler.startSampling.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
*/
type StartSamplingResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
StartTrackingHeapObjectsResult represents the result of calls to
HeapProfiler.startTrackingHeapObjects.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
*/
type StartTrackingHeapObjectsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopSamplingParams represents HeapProfiler.stopSampling parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
*/
type StopSamplingParams struct {
	// Recorded sampling heap profile.
	Profile *SamplingHeapProfile `json:"profile"`
}

/*
StopSamplingResult represents the result of calls to HeapProfiler.stopSampling.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
*/
type StopSamplingResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopTrackingHeapObjectsParams represents HeapProfiler.stopTrackingHeapObjects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
*/
type StopTrackingHeapObjectsParams struct {
	// Optional. If true 'reportHeapSnapshotProgress' events will be generated
	// while snapshot is being taken when the tracking is stopped.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

/*
StopTrackingHeapObjectsResult represents the result of calls to
HeapProfiler.stopTrackingHeapObjects.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
*/
type StopTrackingHeapObjectsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
TakeHeapSnapshotParams represents HeapProfiler.takeHeapSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
*/
type TakeHeapSnapshotParams struct {
	// Optional. If true 'reportHeapSnapshotProgress' events will be generated
	// while snapshot is being taken.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

/*
TakeHeapSnapshotResult represents the result of calls to HeapProfiler.takeHeapSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
*/
type TakeHeapSnapshotResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
