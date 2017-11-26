package HeapProfiler

import (
	Runtime "app/chrome/protocol/runtime"
)

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
