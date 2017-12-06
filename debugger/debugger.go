package Debugger

import (
	Runtime "app/chrome/runtime"
)

/*
ContinueToLocationParams represents Debugger.continueToLocation parameters.
*/
type ContinueToLocationParams struct {
	// Location to continue to.
	Location Location `json:"location"`

	// Optional. Allowed values: any, current.
	TargetCallFrames string `json:"targetCallFrames,omitempty"`
}

/*
EnableResult represents the result of calls to Debugger.enable.
*/
type EnableResult struct {
	// Unique identifier of the debugger. EXPERIMENTAL
	DebuggerID Runtime.UniqueDebuggerID `json:"debuggerId"`
}

/*
EvaluateOnCallFrameParams represents Debugger.evaluateOnCallFrame parameters.
*/
type EvaluateOnCallFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameID CallFrameID `json:"callFrameId"`

	// Expression to evaluate.
	Expression string `json:"expression"`

	// Optional. String object group name to put result into (allows rapid releasing resulting
	// object handles using releaseObjectGroup).
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. Specifies whether command line API should be available to the evaluated expression,
	// defaults to false.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not reported and do not
	// pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
}

/*
EvaluateOnCallFrameResult represents the result of calls to Debugger.evaluateOnCallFrame.
*/
type EvaluateOnCallFrameResult struct {
	// Object wrapper for the evaluation result.
	Result Runtime.RemoteObject `json:"result"`

	// Optional. Exception details.
	ExceptionDetails Runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`
}

/*
GetPossibleBreakpointsParams represents Debugger.getPossibleBreakpoints parameters.
*/
type GetPossibleBreakpointsParams struct {
	// Start of range to search possible breakpoint locations in.
	Start Location `json:"start"`

	// Optional. End of range to search possible breakpoint locations in (excluding). When not
	// specified, end of scripts is used as end of range.
	End Location `json:"end,omitempty"`

	// Optional. Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction bool `json:"restrictToFunction,omitempty"`
}

/*
GetPossibleBreakpointsResult represents the result of calls to Debugger.getPossibleBreakpoints.
*/
type GetPossibleBreakpointsResult struct {
	// List of the possible breakpoint locations.
	Locations []BreakLocation `json:"locations"`
}

/*
GetScriptSourceParams represents Debugger.getScriptSource parameters.
*/
type GetScriptSourceParams struct {
	// ID of the script to get source for.
	ScriptID Runtime.ScriptID `json:"scriptId"`
}

/*
GetScriptSourceResult represents the result of calls to Debugger.getScriptSource.
*/
type GetScriptSourceResult struct {
	// Script source.
	ScriptSource string `json:"scriptSource"`
}

/*
GetStackTraceParams represents Debugger.getStackTrace parameters.
*/
type GetStackTraceParams struct {
	StackTraceID Runtime.StackTraceID `json:"stackTraceId"`
}

/*
GetStackTraceResult represents the result of calls to Debugger.getStackTrace.
*/
type GetStackTraceResult struct {
	StackTrace Runtime.StackTrace `json:"stackTrace"`
}

/*
PauseOnAsyncCallParams represents Debugger.pauseOnAsyncCall parameters.
*/
type PauseOnAsyncCallParams struct {
	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceID Runtime.StackTraceID `json:"parentStackTraceId"`
}

/*
RemoveBreakpointParams represents Debugger.removeBreakpoint parameters.
*/
type RemoveBreakpointParams struct {
	BreakpointID BreakpointID `json:"breakpointId"`
}

/*
RestartFrameParams represents Debugger.restartFrame parameters.
*/
type RestartFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameID CallFrameID `json:"callFrameId"`
}

/*
RestartFrameResult represents the result of calls to Debugger.restartFrame.
*/
type RestartFrameResult struct {
	// New stack trace.
	CallFrames []CallFrame `json:"callFrames"`

	// Optional. Async stack trace, if any.
	AsyncStackTrace Runtime.StackTrace `json:"asyncStackTrace,omitempty"`

	// Optional. Async stack trace, if any. EXPERIMENTAL
	AsyncStackTraceID Runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`
}

/*
SearchInContentParams represents Debugger.searchInContent parameters.
*/
type SearchInContentParams struct {
	// ID of the script to search in.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// String to search for.
	Query string `json:"query"`

	// Optional. If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// Optional. If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

/*
SearchInContentResult represents the result of calls to Debugger.searchInContent.
*/
type SearchInContentResult struct {
	// List of search matches.
	Result []SearchMatch `json:"result"`
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
SetBreakpointParams represents Debugger.setBreakpoint parameters.
*/
type SetBreakpointParams struct {
	// Location to set breakpoint in.
	Location Location `json:"location"`

	// Optional. Expression to use as a breakpoint condition. When specified, debugger will only
	// stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

/*
SetBreakpointResult represents the result of calls to Debugger.setBreakpoint.
*/
type SetBreakpointResult struct {
	// ID of the created breakpoint for further reference.
	BreakpointID BreakpointID `json:"breakpointId"`

	// Location this breakpoint resolved into.
	ActualLocation Location `json:"actualLocation"`
}

/*
SetBreakpointByURLParams represents Debugger.setBreakpointByUrl parameters.
*/
type SetBreakpointByURLParams struct {
	// Line number to set breakpoint at.
	LineNumber int `json:"lineNumber"`

	// Optional. URL of the resources to set breakpoint on.
	URL string `json:"url,omitempty"`

	// Optional. Regex pattern for the URLs of the resources to set breakpoints on. Either url or
	// urlRegex must be specified.
	URLRegex string `json:"urlRegex,omitempty"`

	// Optional. Script hash of the resources to set breakpoint on.
	ScriptHash string `json:"scriptHash,omitempty"`

	// Optional. Offset in the line to set breakpoint at.
	ColumnNumber int `json:"columnNumber,omitempty"`

	// Optional. Expression to use as a breakpoint condition. When specified, debugger will only
	// stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

/*
SetBreakpointByURLResult represents the result of calls to Debugger.setBreakpointByUrl.
*/
type SetBreakpointByURLResult struct {
	// ID of the created breakpoint for further reference.
	BreakpointID BreakpointID `json:"breakpointId"`

	// List of the locations this breakpoint resolved into upon addition.
	Locations []Location `json:"locations"`
}

/*
SetBreakpointsActiveParams represents Debugger.setBreakpointsActive parameters.
*/
type SetBreakpointsActiveParams struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

/*
SetPauseOnExceptionsParams represents Debugger.setPauseOnExceptions parameters.
*/
type SetPauseOnExceptionsParams struct {
	// Pause on exceptions mode. Allowed values: none, uncaught, all.
	State string `json:"state"`
}

/*
SetReturnValueParams represents Debugger.setReturnValue parameters.
*/
type SetReturnValueParams struct {
	// New return value.
	NewValue Runtime.CallArgument `json:"newValue"`
}

/*
SetScriptSourceParams represents Debugger.setScriptSource parameters.
*/
type SetScriptSourceParams struct {
	// ID of the script to edit.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// New content of the script.
	ScriptSource string `json:"scriptSource"`

	// Optional. If true the change will not actually be applied. Dry run may be used to get result
	// description without actually modifying the code.
	DryRun bool `json:"dryRun,omitempty"`
}

/*
SetScriptSourceResult represents the result of calls to Debugger.setScriptSource.
*/
type SetScriptSourceResult struct {
	// Optional. New stack trace in case editing has happened while VM was stopped.
	CallFrames []CallFrame `json:"callFrames,omitempty"`

	// Optional. Whether current call stack was modified after applying the changes.
	StackChanged bool `json:"stackChanged,omitempty"`

	// Optional. Async stack trace, if any.
	AsyncStackTrace Runtime.StackTrace `json:"asyncStackTrace,omitempty"`

	// Optional. Async stack trace, if any. EXPERIMENTAL
	AsyncStackTraceID Runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`

	// Optional. Exception details if any.
	ExceptionDetails Runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`
}

/*
SetSkipAllPausesParams represents Debugger.setSkipAllPauses parameters.
*/
type SetSkipAllPausesParams struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
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
StepIntoParams represents Debugger.stepInto parameters.
*/
type StepIntoParams struct {
	// Optional. Debugger will issue additional Debugger.paused notification if any async task is
	// scheduled before next pause. EXPERIMENTAL
	BreakOnAsyncCall bool `json:"breakOnAsyncCall,omitempty"`
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

	// Optional. Object containing break-specific auxiliary properties.
	Data map[string]string `json:"data,omitempty"`

	// Optional. Hit breakpoints IDs.
	HitBreakpoints []string `json:"hitBreakpoints,omitempty"`

	// Optional. Async stack trace, if any.
	AsyncStackTrace Runtime.StackTrace `json:"asyncStackTrace,omitempty"`

	// Optional. Async stack trace, if any. EXPERIMENTAL
	AsyncStackTraceID Runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`

	// Optional. Just scheduled async call will have this stack trace as parent stack during async
	// execution. This field is available only after Debugger.stepInto call with breakOnAsynCall
	// flag. EXPERIMENTAL
	AsyncCallStackTraceID Runtime.StackTraceID `json:"asyncCallStackTraceId,omitempty"`
}

/*
ResumedEvent represents Debugger.resumed event data.
*/
type ResumedEvent struct{}

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

	// Optional. Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]string `json:"executionContextAuxData,omitempty"`

	// Optional. URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Optional. True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Optional. True, if this script is ES6 module.
	IsModule bool `json:"isModule,omitempty"`

	// Optional. This script length.
	Length int `json:"length,omitempty"`

	// Optional. JavaScript top stack frame of where the script parsed event was triggered if
	// available. EXPERIMENTAL
	StackTrace Runtime.StackTrace `json:"stackTrace,omitempty"`
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

	// Optional. Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]string `json:"executionContextAuxData,omitempty"`

	// Optional. True, if this script is generated as a result of the live edit operation.
	// EXPERIMENTAL.
	IsLiveEdit bool `json:"isLiveEdit,omitempty"`

	// Optional. URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Optional. True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Optional. True, if this script is ES6 module.
	IsModule bool `json:"isModule,omitempty"`

	// Optional. This script length.
	Length int `json:"length,omitempty"`

	// Optional. JavaScript top stack frame of where the script parsed event was triggered if
	// available. EXPERIMENTAL.
	StackTrace Runtime.StackTrace `json:"stackTrace,omitempty"`
}

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
	LineNumber float64 `json:"lineNumber"`

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
