/*
Package heapProfiler provides type definitions for use with the Chrome HeapProfiler protocol

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/
*/
package heapProfiler

import (
	"github.com/mkenney/go-chrome/protocol/runtime"
)

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
	CallFrame *runtime.CallFrame `json:"callFrame"`

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
	Head *SamplingHeapProfileNode `json:"head"`
}
