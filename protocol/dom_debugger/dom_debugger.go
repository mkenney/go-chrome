package DOMDebugger

import (
	DOM "app/chrome/protocol/dom"
	Runtime "app/chrome/protocol/runtime"
	"fmt"
)

/*
DOMBreakpointType is the DOM breakpoint type.
*/
type DOMBreakpointType int

const (
	_subtreeModified DOMBreakpointType = iota
	_attributeModified
	_nodeRemoved
)

func (a DOMBreakpointType) String() string {
	if a == 0 {
		return "subtree-modified"
	}
	if a == 1 {
		return "attribute-modified"
	}
	if a == 2 {
		return "node-removed"
	}
	panic(fmt.Errorf("Invalid DOMBreakpointType %d", a))
}

/*
EventListener is the object event listener.
*/
type EventListener struct {
	// EventListener's type.
	Type string `json:"type"`

	// EventListener's useCapture.
	UseCapture bool `json:"useCapture"`

	// EventListener's passive flag.
	Passive bool `json:"passive"`

	// EventListener's once flag.
	Once bool `json:"once"`

	// Script ID of the handler code.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Column number in the script (0-based).
	ColumnNumber int `json:"columnNumber"`

	// Optional. Event handler function value.
	Handler Runtime.RemoteObject `json:"handler,omitempty"`

	// Optional. Event original handler function value.
	OriginalHandler Runtime.RemoteObject `json:"originalHandler,omitempty"`

	// Optional. Node the listener is added to (if any).
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId,omitempty"`
}
