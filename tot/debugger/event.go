package debugger

import "github.com/mkenney/go-chrome/tot/runtime"

/*
BreakpointResolvedEvent represents Debugger.breakpointResolved event data.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-breakpointResolved
*/
type BreakpointResolvedEvent struct {
	// Breakpoint unique identifier.
	BreakpointID BreakpointID `json:"breakpointId"`

	// Actual breakpoint location.
	Location *Location `json:"location"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
PausedEvent represents Debugger.paused event data.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-paused
*/
type PausedEvent struct {
	// Call stack the virtual machine stopped on.
	CallFrames []*CallFrame `json:"callFrames"`

	// Pause reason. Allowed values: XHR, DOM, EventListener, exception, assert,
	// debugCommand, promiseRejection, OOM, other, ambiguous.
	Reason string `json:"reason"`

	// Optional. Object containing break-specific auxiliary properties.
	Data map[string]string `json:"data,omitempty"`

	// Optional. Hit breakpoints IDs.
	HitBreakpoints []string `json:"hitBreakpoints,omitempty"`

	// Optional. Async stack trace, if any.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`

	// Optional. Async stack trace, if any. EXPERIMENTAL
	AsyncStackTraceID *runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`

	// Optional. Just scheduled async call will have this stack trace as parent
	// stack during async execution. This field is available only after
	// Debugger.stepInto call with breakOnAsynCall flag. EXPERIMENTAL.
	AsyncCallStackTraceID *runtime.StackTraceID `json:"asyncCallStackTraceId,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ResumedEvent represents Debugger.resumed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-resumed
*/
type ResumedEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
ScriptFailedToParseEvent represents Debugger.scriptFailedToParse event data.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptFailedToParse
*/
type ScriptFailedToParseEvent struct {
	// Identifier of the script parsed.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// URL or name of the script parsed (if any).
	URL string `json:"url"`

	// Line offset of the script within the resource with given URL (for script
	// tags).
	StartLine int64 `json:"startLine"`

	// Column offset of the script within the resource with given URL.
	StartColumn int64 `json:"startColumn"`

	// Last line of the script.
	EndLine int64 `json:"endLine"`

	// Length of the last line of the script.
	EndColumn int64 `json:"endColumn"`

	// Specifies script creation context.
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId"`

	// Content hash of the script.
	Hash string `json:"hash"`

	// Optional. Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]string `json:"executionContextAuxData,omitempty"`

	// Optional. URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Optional. True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Optional. True, if this script is ES6 module.
	IsModule bool `json:"isModule,omitempty"`

	// Optional. This script length.
	Length int64 `json:"length,omitempty"`

	// Optional. JavaScript top stack frame of where the script parsed event was
	// triggered if available. EXPERIMENTAL.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ScriptParsedEvent represents Debugger.scriptParsed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptParsed
*/
type ScriptParsedEvent struct {
	// Identifier of the script parsed.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// URL or name of the script parsed (if any).
	URL string `json:"url"`

	// Line offset of the script within the resource with given URL (for script
	// tags).
	StartLine int `json:"startLine"`

	// Column offset of the script within the resource with given URL.
	StartColumn int `json:"startColumn"`

	// Last line of the script.
	EndLine int `json:"endLine"`

	// Length of the last line of the script.
	EndColumn int `json:"endColumn"`

	// Specifies script creation context.
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId"`

	// Content hash of the script.
	Hash string `json:"hash"`

	// Optional. Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]string `json:"executionContextAuxData,omitempty"`

	// Optional. True, if this script is generated as a result of the live edit
	// operation. EXPERIMENTAL.
	IsLiveEdit bool `json:"isLiveEdit,omitempty"`

	// Optional. URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Optional. True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Optional. True, if this script is ES6 module.
	IsModule bool `json:"isModule,omitempty"`

	// Optional. This script length.
	Length int `json:"length,omitempty"`

	// Optional. JavaScript top stack frame of where the script parsed event was
	// triggered if available. EXPERIMENTAL.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}
