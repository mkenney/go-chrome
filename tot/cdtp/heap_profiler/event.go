package heapProfiler

/*
AddHeapSnapshotChunkEvent represents DOM.addHeapSnapshotChunk event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-addHeapSnapshotChunk
*/
type AddHeapSnapshotChunkEvent struct {
	Chunk string `json:"chunk"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
HeapStatsUpdateEvent represents DOM.heapStatsUpdate event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-heapStatsUpdate
*/
type HeapStatsUpdateEvent struct {
	// An array of triplets. Each triplet describes a fragment. The first
	// integer is the fragment index, the second integer is a total count of
	// objects for the fragment, the third integer is a total size of the
	// objects for the fragment.
	StatsUpdate []int `json:"statsUpdate"`

	// Error information related to this event
	Err error `json:"-"`
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

	// Error information related to this event
	Err error `json:"-"`
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

	// Error information related to this event
	Err error `json:"-"`
}

/*
ResetProfilesEvent represents DOM.resetProfiles event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-resetProfiles
*/
type ResetProfilesEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}
