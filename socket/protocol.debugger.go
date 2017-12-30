package socket

import (
	"encoding/json"

	debugger "github.com/mkenney/go-chrome/protocol/debugger"
	log "github.com/sirupsen/logrus"
)

/*
DebuggerProtocol provides a namespace for the Chrome Debugger protocol methods.
The Debugger protocol exposes JavaScript debugging capabilities. It allows
setting and removing breakpoints, stepping through execution, exploring stack
traces, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/
*/
type DebuggerProtocol struct {
	Socket Socketer
}

/*
ContinueToLocation continues execution until specific location is reached.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
*/
func (protocol *DebuggerProtocol) ContinueToLocation(
	params *debugger.ContinueToLocationParams,
) error {
	command := NewCommand("Debugger.continueToLocation", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables debugger for given page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-disable
*/
func (protocol *DebuggerProtocol) Disable() error {
	command := NewCommand("Debugger.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables debugger for the given page. Clients should not assume that the
debugging has been enabled until the result for this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-enable
*/
func (protocol *DebuggerProtocol) Enable() error {
	command := NewCommand("Debugger.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
EvaluateOnCallFrame evaluates expression on a given call frame.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
*/
func (protocol *DebuggerProtocol) EvaluateOnCallFrame(
	params *debugger.EvaluateOnCallFrameParams,
) (*debugger.EvaluateOnCallFrameResult, error) {
	command := NewCommand("Debugger.evaluateOnCallFrame", params)
	result := &debugger.EvaluateOnCallFrameResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetPossibleBreakpoints returns possible locations for breakpoint. scriptId in
start and end range locations should be the same.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
*/
func (protocol *DebuggerProtocol) GetPossibleBreakpoints(
	params *debugger.GetPossibleBreakpointsParams,
) (*debugger.GetPossibleBreakpointsResult, error) {
	command := NewCommand("Debugger.getPossibleBreakpoints", params)
	result := &debugger.GetPossibleBreakpointsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetScriptSource returns source for the script with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
*/
func (protocol *DebuggerProtocol) GetScriptSource(
	params *debugger.GetScriptSourceParams,
) (*debugger.GetScriptSourceResult, error) {
	command := NewCommand("Debugger.getScriptSource", params)
	result := &debugger.GetScriptSourceResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetStackTrace returns stack trace with given stackTraceId.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace EXPERIMENTAL.
*/
func (protocol *DebuggerProtocol) GetStackTrace(
	params *debugger.GetStackTraceParams,
) (*debugger.GetStackTraceResult, error) {
	command := NewCommand("Debugger.getStackTrace", params)
	result := &debugger.GetStackTraceResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Pause stops on the next JavaScript statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pause
*/
func (protocol *DebuggerProtocol) Pause() error {
	command := NewCommand("Debugger.pause", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
PauseOnAsyncCall is experimental

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall EXPERIMENTAL.
*/
func (protocol *DebuggerProtocol) PauseOnAsyncCall(
	params *debugger.PauseOnAsyncCallParams,
) error {
	command := NewCommand("Debugger.pauseOnAsyncCall", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveBreakpoint removes JavaScript breakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
*/
func (protocol *DebuggerProtocol) RemoveBreakpoint(
	params *debugger.RemoveBreakpointParams,
) error {
	command := NewCommand("Debugger.removeBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RestartFrame restarts particular call frame from the beginning.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
*/
func (protocol *DebuggerProtocol) RestartFrame(
	params *debugger.RestartFrameParams,
) (*debugger.RestartFrameResult, error) {
	command := NewCommand("Debugger.restartFrame", params)
	result := &debugger.RestartFrameResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Resume resumes JavaScript execution.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-resume
*/
func (protocol *DebuggerProtocol) Resume() error {
	command := NewCommand("Debugger.resume", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ScheduleStepIntoAsync is deprecated - use Debugger.stepInto with
breakOnAsyncCall and Debugger.pauseOnAsyncTask instead. Steps into next
scheduled async task if any is scheduled before next pause. Returns success when
async task is actually scheduled, returns error if no task were scheduled or
another scheduleStepIntoAsync was called.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-scheduleStepIntoAsync
EXPERIMENTAL. DEPRECATED.
*/
func (protocol *DebuggerProtocol) ScheduleStepIntoAsync() error {
	command := NewCommand("Debugger.scheduleStepIntoAsync", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SearchInContent searches for given string in script content.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
*/
func (protocol *DebuggerProtocol) SearchInContent(
	params *debugger.SearchInContentParams,
) (*debugger.SearchInContentResult, error) {
	command := NewCommand("Debugger.searchInContent", params)
	result := &debugger.SearchInContentResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetAsyncCallStackDepth enables or disables async call stacks tracking.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
*/
func (protocol *DebuggerProtocol) SetAsyncCallStackDepth(
	params *debugger.SetAsyncCallStackDepthParams,
) error {
	command := NewCommand("Debugger.setAsyncCallStackDepth", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetBlackboxPatterns replaces previous blackbox patterns with passed ones. Forces
backend to skip stepping/pausing in scripts with url matching one of the
patterns. VM will try to leave blackboxed script by performing 'step in' several
times, finally resorting to 'step out' if unsuccessful.


https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxPatterns
EXPERIMENTAL.
*/
func (protocol *DebuggerProtocol) SetBlackboxPatterns(
	params *debugger.SetBlackboxPatternsParams,
) error {
	command := NewCommand("Debugger.setBlackboxPatterns", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetBlackboxedRanges makes backend skip steps in the script in blackboxed ranges. VM will try leave
blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if
unsuccessful. Positions array contains positions where blackbox state is changed. First interval
isn't blackboxed. Array should be sorted.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxedRanges
*/
func (protocol *DebuggerProtocol) SetBlackboxedRanges(
	params *debugger.SetBlackboxedRangesParams,
) error {
	command := NewCommand("Debugger.setBlackboxedRanges", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetBreakpoint sets JavaScript breakpoint at a given location.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
*/
func (protocol *DebuggerProtocol) SetBreakpoint(
	params *debugger.SetBreakpointParams,
) (*debugger.SetBreakpointResult, error) {
	command := NewCommand("Debugger.setBreakpoint", params)
	result := &debugger.SetBreakpointResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetBreakpointByURL sets JavaScript breakpoint at given location specified either by URL or URL
regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and
returned in locations property. Further matching script parsing will result in subsequent
breakpointResolved events issued. This logical breakpoint will survive page reloads.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointByUrl
*/
func (protocol *DebuggerProtocol) SetBreakpointByURL(
	params *debugger.SetBreakpointByURLParams,
) (*debugger.SetBreakpointByURLResult, error) {
	command := NewCommand("Debugger.setBreakpointByUrl", params)
	result := &debugger.SetBreakpointByURLResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetBreakpointsActive activates / deactivates all breakpoints on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
*/
func (protocol *DebuggerProtocol) SetBreakpointsActive(
	params *debugger.SetBreakpointsActiveParams,
) error {
	command := NewCommand("Debugger.setBreakpointsActive", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetPauseOnExceptions defines the pause on exceptions state. Can be set to stop on all exceptions,
uncaught exceptions or no exceptions. Initial pause on exceptions state is none.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
*/
func (protocol *DebuggerProtocol) SetPauseOnExceptions(
	params *debugger.SetPauseOnExceptionsParams,
) error {
	command := NewCommand("Debugger.setPauseOnExceptions", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetReturnValue changes return value in top frame. Available only at return break position.


https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue EXPERIMENTAL.
*/
func (protocol *DebuggerProtocol) SetReturnValue(
	params *debugger.SetReturnValueParams,
) error {
	command := NewCommand("Debugger.setReturnValue", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetScriptSource edits JavaScript source live.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
*/
func (protocol *DebuggerProtocol) SetScriptSource(
	params *debugger.SetScriptSourceParams,
) (*debugger.SetScriptSourceResult, error) {
	command := NewCommand("Debugger.setScriptSource", params)
	result := &debugger.SetScriptSourceResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetSkipAllPauses makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
*/
func (protocol *DebuggerProtocol) SetSkipAllPauses(
	params *debugger.SetSkipAllPausesParams,
) error {
	command := NewCommand("Debugger.setSkipAllPauses", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetVariableValue changes value of variable in a callframe. Object-based scopes are not supported and
must be mutated manually.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
*/
func (protocol *DebuggerProtocol) SetVariableValue(
	params *debugger.SetVariableValueParams,
) error {
	command := NewCommand("Debugger.setVariableValue", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StepInto steps into the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
*/
func (protocol *DebuggerProtocol) StepInto(
	params *debugger.StepIntoParams,
) error {
	command := NewCommand("Debugger.stepInto", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StepOut steps out of the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOut
*/
func (protocol *DebuggerProtocol) StepOut() error {
	command := NewCommand("Debugger.stepOut", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StepOver steps over the statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOver
*/
func (protocol *DebuggerProtocol) StepOver() error {
	command := NewCommand("Debugger.stepOver", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnBreakpointResolved adds a handler to the Debugger.breakpointResolved event.
Debugger.breakpointResolved fires when breakpoint is resolved to an actual script and location.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-breakpointResolved
*/
func (protocol *DebuggerProtocol) OnBreakpointResolved(
	callback func(event *debugger.BreakpointResolvedEvent),
) {
	handler := NewEventHandler(
		"Debugger.breakpointResolved",
		func(response *Response) {
			event := &debugger.BreakpointResolvedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnPaused adds a handler to the Debugger.paused event. Debugger.paused fires when the virtual machine
stopped on breakpoint or exception or any other stop criteria.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-paused
*/
func (protocol *DebuggerProtocol) OnPaused(
	callback func(event *debugger.PausedEvent),
) {
	handler := NewEventHandler(
		"Debugger.paused",
		func(response *Response) {
			event := &debugger.PausedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnResumed adds a handler to the Debugger.resumed event. Debugger.resumed fires when the virtual
machine resumes execution.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-resumed
*/
func (protocol *DebuggerProtocol) OnResumed(
	callback func(event *debugger.ResumedEvent),
) {
	handler := NewEventHandler(
		"Debugger.resumed",
		func(response *Response) {
			event := &debugger.ResumedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnScriptFailedToParse adds a handler to the Debugger.scriptFailedToParse event.
Debugger.scriptFailedToParse fires when the virtual machine fails to parse the script.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptFailedToParse
*/
func (protocol *DebuggerProtocol) OnScriptFailedToParse(
	callback func(event *debugger.ScriptFailedToParseEvent),
) {
	handler := NewEventHandler(
		"Debugger.scriptFailedToParse",
		func(response *Response) {
			event := &debugger.ScriptFailedToParseEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnScriptParsed adds a handler to the Debugger.ScriptParsed event. Debugger.ScriptParsed fires when
virtual machine parses script. This event is also fired for all known and uncollected scripts upon
enabling debugger.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptParsed
*/
func (protocol *DebuggerProtocol) OnScriptParsed(
	callback func(event *debugger.ScriptParsedEvent),
) {
	handler := NewEventHandler(
		"Debugger.scriptParsed",
		func(response *Response) {
			event := &debugger.ScriptParsedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
