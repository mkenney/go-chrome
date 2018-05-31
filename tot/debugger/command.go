package debugger

import "github.com/mkenney/go-chrome/tot/runtime"

/*
ContinueToLocationParams represents Debugger.continueToLocation parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
*/
type ContinueToLocationParams struct {
	// Location to continue to.
	Location *Location `json:"location"`

	// Optional. Allowed values:
	//	- TargetCallFrames.Any
	//	- TargetCallFrames.Current
	TargetCallFrames TargetCallFramesEnum `json:"targetCallFrames,omitempty"`
}

/*
ContinueToLocationResult represents the result of calls to
Debugger.continueToLocation.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
*/
type ContinueToLocationResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to Debugger.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Debugger.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-enable
*/
type EnableResult struct {
	// Unique identifier of the debugger. EXPERIMENTAL
	ID runtime.UniqueDebuggerID `json:"debuggerId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EvaluateOnCallFrameParams represents Debugger.evaluateOnCallFrame parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
*/
type EvaluateOnCallFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameID CallFrameID `json:"callFrameId"`

	// Expression to evaluate.
	Expression string `json:"expression"`

	// Optional. String object group name to put result into (allows rapid
	// releasing resulting object handles using releaseObjectGroup).
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. Specifies whether command line API should be available to the
	// evaluated expression, defaults to false.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not
	// reported and do not pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Whether the result is expected to be a JSON object that should
	// be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	// EXPERIMENTAL.
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether to throw an exception if side effect cannot be ruled
	// out during evaluation.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
}

/*
EvaluateOnCallFrameResult represents the result of calls to Debugger.evaluateOnCallFrame.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
*/
type EvaluateOnCallFrameResult struct {
	// Object wrapper for the evaluation result.
	Result *runtime.RemoteObject `json:"result"`

	// Optional. Exception details.
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetPossibleBreakpointsParams represents Debugger.getPossibleBreakpoints parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
*/
type GetPossibleBreakpointsParams struct {
	// Start of range to search possible breakpoint locations in.
	Start *Location `json:"start"`

	// Optional. End of range to search possible breakpoint locations in
	// (excluding). When not specified, end of scripts is used as end of range.
	End *Location `json:"end,omitempty"`

	// Optional. Only consider locations which are in the same (non-nested)
	// function as start.
	RestrictToFunction bool `json:"restrictToFunction,omitempty"`
}

/*
GetPossibleBreakpointsResult represents the result of calls to Debugger.getPossibleBreakpoints.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
*/
type GetPossibleBreakpointsResult struct {
	// List of the possible breakpoint locations.
	Locations []*BreakLocation `json:"locations"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetScriptSourceParams represents Debugger.getScriptSource parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
*/
type GetScriptSourceParams struct {
	// ID of the script to get source for.
	ScriptID runtime.ScriptID `json:"scriptId"`
}

/*
GetScriptSourceResult represents the result of calls to Debugger.getScriptSource.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
*/
type GetScriptSourceResult struct {
	// Script source.
	ScriptSource string `json:"scriptSource"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetStackTraceParams represents Debugger.getStackTrace parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace
*/
type GetStackTraceParams struct {
	StackTraceID runtime.StackTraceID `json:"stackTraceId"`
}

/*
GetStackTraceResult represents the result of calls to Debugger.getStackTrace.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace
*/
type GetStackTraceResult struct {
	StackTrace *runtime.StackTrace `json:"stackTrace"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
PauseResult represents the result of calls to Debugger.pause.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pause
*/
type PauseResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
PauseOnAsyncCallParams represents Debugger.pauseOnAsyncCall parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall
*/
type PauseOnAsyncCallParams struct {
	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceID runtime.StackTraceID `json:"parentStackTraceId"`
}

/*
PauseOnAsyncCallResult represents the result of calls to Debugger.pauseOnAsyncCall.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall
*/
type PauseOnAsyncCallResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RemoveBreakpointParams represents Debugger.removeBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
*/
type RemoveBreakpointParams struct {
	BreakpointID BreakpointID `json:"breakpointId"`
}

/*
RemoveBreakpointResult represents the result of calls to Debugger.removeBreakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
*/
type RemoveBreakpointResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RestartFrameParams represents Debugger.restartFrame parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
*/
type RestartFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameID CallFrameID `json:"callFrameId"`
}

/*
RestartFrameResult represents the result of calls to Debugger.restartFrame.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
*/
type RestartFrameResult struct {
	// New stack trace.
	CallFrames []*CallFrame `json:"callFrames"`

	// Optional. Async stack trace, if any.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`

	// Optional. Async stack trace, if any. EXPERIMENTAL
	AsyncStackTraceID runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ResumeResult represents the result of calls to Debugger.resume.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-resume
*/
type ResumeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ScheduleStepIntoAsyncResult represents the result of calls to Debugger.scheduleStepIntoAsync.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-scheduleStepIntoAsync
*/
type ScheduleStepIntoAsyncResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SearchInContentParams represents Debugger.searchInContent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
*/
type SearchInContentParams struct {
	// ID of the script to search in.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// String to search for.
	Query string `json:"query"`

	// Optional. If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// Optional. If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

/*
SearchInContentResult represents the result of calls to Debugger.searchInContent.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
*/
type SearchInContentResult struct {
	// List of search matches.
	Result []*SearchMatch `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetAsyncCallStackDepthParams represents Debugger.setAsyncCallStackDepth parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
*/
type SetAsyncCallStackDepthParams struct {
	// Maximum depth of async call stacks. Setting to 0 will effectively disable
	// collecting async call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

/*
SetAsyncCallStackDepthResult represents the result of calls to Debugger.setAsyncCallStackDepth.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
*/
type SetAsyncCallStackDepthResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetBlackboxPatternsParams represents Debugger.setBlackboxPatterns parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxPatterns
*/
type SetBlackboxPatternsParams struct {
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns []string `json:"patterns"`
}

/*
SetBlackboxPatternsResult represents the result of calls to Debugger.setBlackboxPatterns.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxPatterns
*/
type SetBlackboxPatternsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetBlackboxedRangesParams represents Debugger.setBlackboxedRanges parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxedRanges
*/
type SetBlackboxedRangesParams struct {
	// ID of the script.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// Script positions.
	Positions []*ScriptPosition `json:"positions"`
}

/*
SetBlackboxedRangesResult represents the result of calls to Debugger.setBlackboxedRanges.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxedRanges
*/
type SetBlackboxedRangesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetBreakpointParams represents Debugger.setBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
*/
type SetBreakpointParams struct {
	// Location to set breakpoint in.
	Location *Location `json:"location"`

	// Optional. Expression to use as a breakpoint condition. When specified,
	// debugger will only stop on the breakpoint if this expression evaluates to
	// true.
	Condition string `json:"condition,omitempty"`
}

/*
SetBreakpointResult represents the result of calls to Debugger.setBreakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
*/
type SetBreakpointResult struct {
	// ID of the created breakpoint for further reference.
	BreakpointID BreakpointID `json:"breakpointId"`

	// Location this breakpoint resolved into.
	ActualLocation *Location `json:"actualLocation"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetBreakpointByURLParams represents Debugger.setBreakpointByUrl parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointByUrl
*/
type SetBreakpointByURLParams struct {
	// Line number to set breakpoint at.
	LineNumber int64 `json:"lineNumber"`

	// Optional. URL of the resources to set breakpoint on.
	URL string `json:"url,omitempty"`

	// Optional. Regex pattern for the URLs of the resources to set breakpoints
	// on. Either url or urlRegex must be specified.
	URLRegex string `json:"urlRegex,omitempty"`

	// Optional. Script hash of the resources to set breakpoint on.
	ScriptHash string `json:"scriptHash,omitempty"`

	// Optional. Offset in the line to set breakpoint at.
	ColumnNumber int64 `json:"columnNumber,omitempty"`

	// Optional. Expression to use as a breakpoint condition. When specified,
	// debugger will only stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

/*
SetBreakpointByURLResult represents the result of calls to Debugger.setBreakpointByUrl.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointByUrl
*/
type SetBreakpointByURLResult struct {
	// ID of the created breakpoint for further reference.
	BreakpointID BreakpointID `json:"breakpointId"`

	// List of the locations this breakpoint resolved into upon addition.
	Locations []*Location `json:"locations"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetBreakpointsActiveParams represents Debugger.setBreakpointsActive parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
*/
type SetBreakpointsActiveParams struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

/*
SetBreakpointsActiveResult represents the result of calls to Debugger.setBreakpointsActive.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
*/
type SetBreakpointsActiveResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetPauseOnExceptionsParams represents Debugger.setPauseOnExceptions parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
*/
type SetPauseOnExceptionsParams struct {
	// Pause on exceptions mode. Allowed values:
	//	- State.None
	//	- State.Uncaught
	//	- State.All
	State StateEnum `json:"state"`
}

/*
SetPauseOnExceptionsResult represents the result of calls to Debugger.setPauseOnExceptions.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
*/
type SetPauseOnExceptionsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetReturnValueParams represents Debugger.setReturnValue parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue
*/
type SetReturnValueParams struct {
	// New return value.
	NewValue *runtime.CallArgument `json:"newValue"`
}

/*
SetReturnValueResult represents the result of calls to Debugger.setReturnValue.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue
*/
type SetReturnValueResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetScriptSourceParams represents Debugger.setScriptSource parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
*/
type SetScriptSourceParams struct {
	// ID of the script to edit.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// New content of the script.
	ScriptSource string `json:"scriptSource"`

	// Optional. If true the change will not actually be applied. Dry run may be
	// used to get result description without actually modifying the code.
	DryRun bool `json:"dryRun,omitempty"`
}

/*
SetScriptSourceResult represents the result of calls to Debugger.setScriptSource.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
*/
type SetScriptSourceResult struct {
	// Optional. New stack trace in case editing has happened while VM was
	// stopped.
	CallFrames []*CallFrame `json:"callFrames,omitempty"`

	// Optional. Whether current call stack was modified after applying the
	// changes.
	StackChanged bool `json:"stackChanged,omitempty"`

	// Optional. Async stack trace, if any.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`

	// Optional. Async stack trace, if any. EXPERIMENTAL.
	AsyncStackTraceID runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`

	// Optional. Exception details if any.
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetSkipAllPausesParams represents Debugger.setSkipAllPauses parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
*/
type SetSkipAllPausesParams struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
}

/*
SetSkipAllPausesResult represents the result of calls to
Debugger.setSkipAllPauses.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
*/
type SetSkipAllPausesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetVariableValueParams represents Debugger.setVariableValue parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
*/
type SetVariableValueParams struct {
	// 0-based number of scope as was listed in scope chain. Only 'local',
	// 'closure' and 'catch' scope types are allowed. Other scopes could be
	// manipulated manually.
	ScopeNumber int64 `json:"scopeNumber"`

	// Variable name.
	VariableName string `json:"variableName"`

	// New variable value.
	NewValue *runtime.CallArgument `json:"newValue"`

	// ID of callframe that holds variable.
	CallFrameID CallFrameID `json:"callFrameId"`
}

/*
SetVariableValueResult represents the result of calls to Debugger.setVariableValue.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
*/
type SetVariableValueResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StepIntoParams represents Debugger.stepInto parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
*/
type StepIntoParams struct {
	// Optional. Debugger will issue additional Debugger.paused notification if
	// any async task is scheduled before next pause. EXPERIMENTAL.
	BreakOnAsyncCall bool `json:"breakOnAsyncCall,omitempty"`
}

/*
StepIntoResult represents the result of calls to Debugger.stepInto.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
*/
type StepIntoResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StepOutResult represents the result of calls to Debugger.stepOut.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOut
*/
type StepOutResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StepOverResult represents the result of calls to Debugger.stepOver.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOver
*/
type StepOverResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
