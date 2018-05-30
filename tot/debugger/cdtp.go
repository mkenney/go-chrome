/*
Package debugger provides type definitions for use with the Chrome Debugger protocol

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/
*/
package debugger

import (
	"github.com/mkenney/go-chrome/tot/runtime"
)

/*
BreakpointID is a breakpoint identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-BreakpointId
*/
type BreakpointID string

/*
CallFrameID is a call frame identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-CallFrameId
*/
type CallFrameID string

/*
Location is a location in the source code.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-Location
*/
type Location struct {
	// Script identifier as reported in the Debugger.scriptParsed.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int64 `json:"lineNumber"`

	// Optional. Column number in the script (0-based).
	ColumnNumber int64 `json:"columnNumber,omitempty"`
}

/*
ScriptPosition is a location in the source code. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-ScriptPosition
*/
type ScriptPosition struct {
	// Line number
	LineNumber int64 `json:"lineNumber"`

	// Column number
	ColumnNumber int64 `json:"columnNumber"`
}

/*
CallFrame is a JavaScript call frame. Array of call frames form the call stack.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-CallFrame
*/
type CallFrame struct {
	// Call frame identifier. This identifier is only valid while the virtual
	// machine is paused.
	CallFrameID CallFrameID `json:"callFrameId"`

	// Name of the JavaScript function called on this call frame.
	FunctionName string `json:"functionName"`

	// Optional. Location in the source code.
	FunctionLocation *Location `json:"functionLocation,omitempty"`

	// Location in the source code.
	Location *Location `json:"location"`

	// JavaScript script name or url.
	URL string `json:"url"`

	// Scope chain for this call frame.
	ScopeChain []*Scope `json:"scopeChain"`

	// `this` object for this call frame.
	This *runtime.RemoteObject `json:"this"`

	// Optional. The value being returned, if the function is at return point.
	ReturnValue *runtime.RemoteObject `json:"returnValue,omitempty"`
}

/*
Scope represents a scope description

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-Scope
*/
type Scope struct {
	// Scope type. Allowed values:
	//	- ScopeType.Global
	//	- ScopeType.Local
	//	- ScopeType.With
	//	- ScopeType.Closure
	//	- ScopeType.Catch
	//	- ScopeType.Block
	//	- ScopeType.Script
	//	- ScopeType.Eval
	//	- ScopeType.Module.
	Type ScopeTypeEnum `json:"type"`

	// Object representing the scope. For global and with scopes it represents
	// the actual object; for the rest of the scopes, it is artificial transient
	// object enumerating scope variables as its properties.
	Object *runtime.RemoteObject `json:"object"`

	// Optional. The scope name.
	Name string `json:"name,omitempty"`

	// Optional. Location in the source code where scope starts.
	StartLocation *Location `json:"startLocation,omitempty"`

	// Optional. Location in the source code where scope ends.
	EndLocation *Location `json:"endLocation,omitempty"`
}

/*
SearchMatch is a search match for a resource.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-SearchMatch
*/
type SearchMatch struct {
	// Line number in resource content.
	LineNumber int64 `json:"lineNumber"`

	// Line with match content.
	LineContent string `json:"lineContent"`
}

/*
BreakLocation is a break location

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-BreakLocation
*/
type BreakLocation struct {
	// Script identifier as reported in the Debugger.scriptParsed.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Optional. Column number in the script (0-based).
	ColumnNumber int `json:"columnNumber,omitempty"`

	// Optional. Allowed values:
	//	- BreakLocationType.DebuggerStatement
	//	- BreakLocationType.Call
	//	- BreakLocationType.Return.
	Type BreakLocationTypeEnum `json:"type,omitempty"`
}
