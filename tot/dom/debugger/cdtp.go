/*
Package debugger provides type definitions for use with the Chrome DOMDebugger protocol

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/
*/
package debugger

import (
	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/runtime"
)

/*
DOMBreakpointType is the DOM breakpoint type.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#type-DOMBreakpointType
*/
type DOMBreakpointType string

/*
EventListener is the object event listener.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#type-EventListener
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
	ScriptID runtime.ScriptID `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int64 `json:"lineNumber"`

	// Column number in the script (0-based).
	ColumnNumber int64 `json:"columnNumber"`

	// Optional. Event handler function value.
	Handler *runtime.RemoteObject `json:"handler,omitempty"`

	// Optional. Event original handler function value.
	OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"`

	// Optional. Node the listener is added to (if any).
	BackendNodeID dom.BackendNodeID `json:"backendNodeId,omitempty"`
}
