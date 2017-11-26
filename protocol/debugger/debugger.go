package Debugger

import (
	Runtime "app/chrome/protocol/runtime"
)

/*
BreakpointID is a breakpoint identifier.
*/
type BreakpointID string

/*
CallFrameID is a call frame identifier.
*/
type CallFrameID string

/*
Location is a location in the source code.
*/
type Location struct {
	// Script identifier as reported in the Debugger.scriptParsed.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Optional. Column number in the script (0-based).
	ColumnNumber int `json:"columnNumber,omitempty"`
}

/*
ScriptPosition is a location in the source code. EXPERIMENTAL
*/
type ScriptPosition struct {
	// Line number
	LineNumber int `json:"lineNumber"`

	// Column number
	ColumnNumber int `json:"columnNumber"`
}

/*
CallFrame is a JavaScript call frame. Array of call frames form the call stack.
*/
type CallFrame struct {
	// Call frame identifier. This identifier is only valid while the virtual machine is paused.
	CallFrameID CallFrameID `json:"callFrameId"`

	// Name of the JavaScript function called on this call frame.
	FunctionName string `json:"functionName"`

	// Optional. Location in the source code.
	FunctionLocation Location `json:"functionLocation,omitempty"`

	// Location in the source code.
	Location Location `json:"location"`

	// JavaScript script name or url.
	URL string `json:"url"`

	// Scope chain for this call frame.
	ScopeChain []*Scope `json:"scopeChain"`

	// this object for this call frame.
	This Runtime.RemoteObject `json:"this"`

	// Optional. The value being returned, if the function is at return point.
	ReturnValue Runtime.RemoteObject `json:"returnValue,omitempty"`
}

/*
Scope represents a scope description
*/
type Scope struct {
	// Scope type. Allowed values: global, local, with, closure, catch, block, script, eval, module.
	Type string `json:"type"`

	// Object representing the scope. For global and with scopes it represents the actual object;
	// for the rest of the scopes, it is artificial transient object enumerating scope variables as
	// its properties.
	Object Runtime.RemoteObject `json:"object"`

	// Optional. The scope name.
	Name string `json:"name,omitempty"`

	// Optional. Location in the source code where scope starts.
	StartLocation Location `json:"startLocation,omitempty"`

	// Optional. Location in the source code where scope ends.
	EndLocation Location `json:"endLocation,omitempty"`
}

/*
SearchMatch is a search match for a resource.
*/
type SearchMatch struct {
	// Line number in resource content.
	LineNumber int `json:"lineNumber"`

	// Line with match content.
	LineContent string `json:"lineContent"`
}

/*
BreakLocation is a break location
*/
type BreakLocation struct {
	// Script identifier as reported in the Debugger.scriptParsed.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Optional. Column number in the script (0-based).
	ColumnNumber int `json:"columnNumber,omitempty"`

	// Optional. Allowed values: debuggerStatement, call, return.
	Type string `json:"type,omitempty"`
}

/*
SetBreakpointsActiveParams represents Debugger.setBreakpointsActive parameters.
*/
type SetBreakpointsActiveParams struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

/*
SetSkipAllPausesParams represents Debugger.setSkipAllPauses parameters.
*/
type SetSkipAllPausesParams struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
}

/*
SetBreakpointByURLParams represents Debugger.setBreakpointByUrl parameters.
*/
type SetBreakpointByURLParams struct {
	// Line number to set breakpoint at.
	LineNumber int `json:"lineNumber"`

	// URL of the resources to set breakpoint on.
	URL string `json:"url"`

	// Regex pattern for the URLs of the resources to set breakpoints on. Either url or urlRegex
	// must be specified.
	URLRegex string `json:"urlRegex"`

	// Script hash of the resources to set breakpoint on.
	ScriptHash string `json:"scriptHash"`

	// Offset in the line to set breakpoint at.
	ColumnNumber int `json:"columnNumber"`

	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition string `json:"condition"`
}

/*
SetBreakpointParams represents Debugger.setBreakpoint parameters.
*/
type SetBreakpointParams struct {
	// Location to set breakpoint in.
	Location Location `json:"location"`

	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition string `json:"condition"`
}

/*
RemoveBreakpointParams represents Debugger.removeBreakpoint parameters.
*/
type RemoveBreakpointParams struct {
	BreakpointID BreakpointID `json:"breakpointId"`
}

/*
GetPossibleBreakpointsParams represents Debugger.getPossibleBreakpoints parameters.
*/
type GetPossibleBreakpointsParams struct {
	// Start of range to search possible breakpoint locations in.
	Start Location `json:"start"`

	// End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
	End Location `json:"end"`

	// Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction bool `json:"restrictToFunction"`
}

/*
ContinueToLocationParams represents Debugger.continueToLocation parameters.
*/
type ContinueToLocationParams struct {
	// Location to continue to.
	Location Location `json:"location"`

	// Allowed values: any, current.
	TargetCallFrames string `json:"targetCallFrames"`
}

/*
PauseOnAsyncCallParams represents Debugger.pauseOnAsyncCall parameters.
*/
type PauseOnAsyncCallParams struct {
	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceID Runtime.StackTraceID `json:"parentStackTraceId"`
}

/*
StepIntoParams represents Debugger.stepInto parameters.
*/
type StepIntoParams struct {
	// Debugger will issue additional Debugger.paused notification if any async task is scheduled
	// before next pause. EXPERIMENTAL
	BreakOnAsyncCall bool `json:"breakOnAsyncCall"`
}

/*
GetStackTraceParams represents Debugger.getStackTrace parameters.
*/
type GetStackTraceParams struct {
	StackTraceID Runtime.StackTraceID `json:"stackTraceId"`
}

/*
SearchInContentParams represents Debugger.searchInContent parameters.
*/
type SearchInContentParams struct {
	// ID of the script to search in.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// String to search for.
	Query string `json:"query"`

	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive"`

	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex"`
}

/*
SetScriptSourceParams represents Debugger.setScriptSource parameters.
*/
type SetScriptSourceParams struct {
	// ID of the script to edit.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// New content of the script.
	ScriptSource string `json:"scriptSource"`

	// If true the change will not actually be applied. Dry run may be used to get result
	// description without actually modifying the code.
	DryRun bool `json:"dryRun"`
}

/*
RestartFrameParams represents Debugger.restartFrame parameters.
*/
type RestartFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameID CallFrameID `json:"callFrameId"`
}

/*
GetScriptSourceParams represents Debugger.getScriptSource parameters.
*/
type GetScriptSourceParams struct {
	// ID of the script to get source for.
	ScriptID Runtime.ScriptID `json:"scriptId"`
}

/*
SetPauseOnExceptionsParams represents Debugger.setPauseOnExceptions parameters.
*/
type SetPauseOnExceptionsParams struct {
	// Pause on exceptions mode. Allowed values: none, uncaught, all.
	State string `json:"state"`
}

/*
EvaluateOnCallFrameParams represents Debugger.evaluateOnCallFrame parameters.
*/
type EvaluateOnCallFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameID CallFrameID `json:"callFrameId"`

	// Expression to evaluate.
	Expression string `json:"expression"`

	// String object group name to put result into (allows rapid releasing resulting object handles
	// using releaseObjectGroup).
	ObjectGroup string `json:"objectGroup"`

	// Specifies whether command line API should be available to the evaluated expression, defaults
	// to false.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI"`

	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides setPauseOnException state.
	Silent bool `json:"silent"`

	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue"`

	// Whether preview should be generated for the result. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview"`

	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect bool `json:"throwOnSideEffect"`
}

/*
SetVariableValueParams represents Debugger.setVariableValue parameters.
*/
type SetVariableValueParams struct {
	// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
	ScopeNumber int `json:"scopeNumber"`

	// Variable name.
	VariableName string `json:"variableName"`

	// New variable value.
	NewValue Runtime.CallArgument `json:"newValue"`

	// ID of callframe that holds variable.
	CallFrameID CallFrameID `json:"callFrameId"`
}

/*
SetReturnValueParams represents Debugger.setReturnValue parameters.
*/
type SetReturnValueParams struct {
	// New return value.
	NewValue Runtime.CallArgument `json:"newValue"`
}

/*
SetAsyncCallStackDepthParams represents Debugger.setAsyncCallStackDepth parameters.
*/
type SetAsyncCallStackDepthParams struct {
	// Maximum depth of async call stacks. Setting to 0 will effectively disable collecting async
	// call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

/*
SetBlackboxPatternsParams represents Debugger.setBlackboxPatterns parameters.
*/
type SetBlackboxPatternsParams struct {
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns []string `json:"patterns"`
}

/*
SetBlackboxedRangesParams represents Debugger.setBlackboxedRanges parameters.
*/
type SetBlackboxedRangesParams struct {
	// ID of the script.
	ScriptID  Runtime.ScriptID  `json:"scriptId"`
	Positions []*ScriptPosition `json:"positions"`
}

/*
ScriptParsedEvent represents Debugger.scriptParsed event data.
*/
type ScriptParsedEvent struct {
	// Identifier of the script parsed.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// URL or name of the script parsed (if any).
	URL string `json:"url"`

	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int `json:"startLine"`

	// Column offset of the script within the resource with given URL.
	StartColumn int `json:"startColumn"`

	// Last line of the script.
	EndLine int `json:"endLine"`

	// Length of the last line of the script.
	EndColumn int `json:"endColumn"`

	// Specifies script creation context.
	ExecutionContextID Runtime.ExecutionContextID `json:"executionContextId"`

	// Content hash of the script.
	Hash string `json:"hash"`

	// Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]string `json:"executionContextAuxData"`

	// True, if this script is generated as a result of the live edit operation. EXPERIMENTAL.
	IsLiveEdit bool `json:"isLiveEdit"`

	// URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL"`

	// True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL"`

	// True, if this script is ES6 module.
	IsModule bool `json:"isModule"`

	// This script length.
	Length int `json:"length"`

	// JavaScript top stack frame of where the script parsed event was triggered if available. EXPERIMENTAL.
	StackTrace Runtime.StackTrace `json:"stackTrace"`
}

/*
ScriptFailedToParseEvent represents Debugger.scriptFailedToParse event data.
*/
type ScriptFailedToParseEvent struct {
	// Identifier of the script parsed.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// URL or name of the script parsed (if any).
	URL string `json:"url"`

	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int `json:"startLine"`

	// Column offset of the script within the resource with given URL.
	StartColumn int `json:"startColumn"`

	// Last line of the script.
	EndLine int `json:"endLine"`

	// Length of the last line of the script.
	EndColumn int `json:"endColumn"`

	// Specifies script creation context.
	ExecutionContextID Runtime.ExecutionContextID `json:"executionContextId"`

	// Content hash of the script.
	Hash string `json:"hash"`

	// Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]string `json:"executionContextAuxData"`

	// URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL"`

	// True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL"`

	// True, if this script is ES6 module.
	IsModule bool `json:"isModule"`

	// This script length.
	Length int `json:"length"`

	// JavaScript top stack frame of where the script parsed event was triggered if available.
	// EXPERIMENTAL
	StackTrace Runtime.StackTrace `json:"stackTrace"`
}

/*
BreakpointResolvedEvent represents Debugger.breakpointResolved event data.
*/
type BreakpointResolvedEvent struct {
	// Breakpoint unique identifier.
	BreakpointID BreakpointID `json:"breakpointId"`

	// Actual breakpoint location.
	Location Location `json:"location"`
}

/*
PausedEvent represents Debugger.paused event data.
*/
type PausedEvent struct {
	// Call stack the virtual machine stopped on.
	CallFrames []*CallFrame `json:"callFrames"`

	// Pause reason. Allowed values: XHR, DOM, EventListener, exception, assert, debugCommand,
	// promiseRejection, OOM, other, ambiguous.
	Reason string `json:"reason"`

	// Object containing break-specific auxiliary properties.
	Data map[string]string `json:"data"`

	// Hit breakpoints IDs.
	HitBreakpoints []string `json:"hitBreakpoints"`

	// Async stack trace, if any.
	AsyncStackTrace Runtime.StackTrace `json:"asyncStackTrace"`

	// Async stack trace, if any. EXPERIMENTAL
	AsyncStackTraceID Runtime.StackTraceID `json:"asyncStackTraceId"`

	// Just scheduled async call will have this stack trace as parent stack during async execution.
	// This field is available only after Debugger.stepInto call with breakOnAsynCall flag.
	// EXPERIMENTAL
	AsyncCallStackTraceID Runtime.StackTraceID `json:"asyncCallStackTraceId"`
}

/*
ResumedEvent represents Debugger.resumed event data.
*/
type ResumedEvent struct{}
