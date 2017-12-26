package protocol

import (
	"encoding/json"

	debugger "github.com/mkenney/go-chrome/protocol/debugger"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Debugger is a struct that provides a namespace for the Chrome Debugger protocol methods.

The Debugger protocol exposes JavaScript debugging capabilities. It allows setting and removing
breakpoints, stepping through execution, exploring stack traces, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/
*/
type Debugger struct{}

/*
ContinueToLocation continues execution until specific location is reached.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
*/
func (Debugger) ContinueToLocation(
	socket *sock.Socket,
	params *debugger.ContinueToLocationParams,
) error {
	command := &sock.Command{
		Method: "Debugger.continueToLocation",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables debugger for given page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-disable
*/
func (Debugger) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Debugger.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables debugger for the given page. Clients should not assume that the debugging has been
enabled until the result for this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-enable
*/
func (Debugger) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Debugger.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
EvaluateOnCallFrame evaluates expression on a given call frame.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
*/
func (Debugger) EvaluateOnCallFrame(
	socket *sock.Socket,
	params *debugger.EvaluateOnCallFrameParams,
) (debugger.EvaluateOnCallFrameResult, error) {
	command := &sock.Command{
		Method: "Debugger.evaluateOnCallFrame",
		Params: params,
	}
	result := debugger.EvaluateOnCallFrameResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetPossibleBreakpoints returns possible locations for breakpoint. scriptId in start and end range
locations should be the same.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
*/
func (Debugger) GetPossibleBreakpoints(
	socket *sock.Socket,
	params *debugger.GetPossibleBreakpointsParams,
) (debugger.GetPossibleBreakpointsResult, error) {
	command := &sock.Command{
		Method: "Debugger.getPossibleBreakpoints",
		Params: params,
	}
	result := debugger.GetPossibleBreakpointsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetScriptSource returns source for the script with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
*/
func (Debugger) GetScriptSource(
	socket *sock.Socket,
	params *debugger.GetScriptSourceParams,
) (debugger.GetScriptSourceResult, error) {
	command := &sock.Command{
		Method: "Debugger.getScriptSource",
		Params: params,
	}
	result := debugger.GetScriptSourceResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetStackTrace returns stack trace with given stackTraceId. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace
*/
func (Debugger) GetStackTrace(
	socket *sock.Socket,
	params *debugger.GetStackTraceParams,
) (debugger.GetStackTraceResult, error) {
	command := &sock.Command{
		Method: "Debugger.getStackTrace",
		Params: params,
	}
	result := debugger.GetStackTraceResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
Pause stops on the next JavaScript statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pause
*/
func (Debugger) Pause(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Debugger.pause",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
PauseOnAsyncCall EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall
*/
func (Debugger) PauseOnAsyncCall(
	socket *sock.Socket,
	params *debugger.PauseOnAsyncCallParams,
) error {
	command := &sock.Command{
		Method: "Debugger.pauseOnAsyncCall",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveBreakpoint removes JavaScript breakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
*/
func (Debugger) RemoveBreakpoint(
	socket *sock.Socket,
	params *debugger.RemoveBreakpointParams,
) error {
	command := &sock.Command{
		Method: "Debugger.removeBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RestartFrame restarts particular call frame from the beginning.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
*/
func (Debugger) RestartFrame(
	socket *sock.Socket,
	params *debugger.RestartFrameParams,
) (debugger.RestartFrameResult, error) {
	command := &sock.Command{
		Method: "Debugger.restartFrame",
		Params: params,
	}
	result := debugger.RestartFrameResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
Resume resumes JavaScript execution.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-resume
*/
func (Debugger) Resume(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Debugger.resume",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ScheduleStepIntoAsync is deprecated - use Debugger.stepInto with breakOnAsyncCall and
Debugger.pauseOnAsyncTask instead. Steps into next scheduled async task if any is scheduled before
next pause. Returns success when async task is actually scheduled, returns error if no task were
scheduled or another scheduleStepIntoAsync was called. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-scheduleStepIntoAsync
*/
func (Debugger) ScheduleStepIntoAsync(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Debugger.scheduleStepIntoAsync",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SearchInContent searches for given string in script content.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
*/
func (Debugger) SearchInContent(
	socket *sock.Socket,
	params *debugger.SearchInContentParams,
) (debugger.SearchInContentResult, error) {
	command := &sock.Command{
		Method: "Debugger.searchInContent",
		Params: params,
	}
	result := debugger.SearchInContentResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetAsyncCallStackDepth enables or disables async call stacks tracking.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
*/
func (Debugger) SetAsyncCallStackDepth(
	socket *sock.Socket,
	params *debugger.SetAsyncCallStackDepthParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setAsyncCallStackDepth",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetBlackboxPatterns replaces previous blackbox patterns with passed ones. Forces backend to skip
stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed
script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxPatterns
*/
func (Debugger) SetBlackboxPatterns(
	socket *sock.Socket,
	params *debugger.SetBlackboxPatternsParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setBlackboxPatterns",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetBlackboxedRanges makes backend skip steps in the script in blackboxed ranges. VM will try leave
blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if
unsuccessful. Positions array contains positions where blackbox state is changed. First interval
isn't blackboxed. Array should be sorted.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxedRanges
*/
func (Debugger) SetBlackboxedRanges(
	socket *sock.Socket,
	params *debugger.SetBlackboxedRangesParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setBlackboxedRanges",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetBreakpoint sets JavaScript breakpoint at a given location.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
*/
func (Debugger) SetBreakpoint(
	socket *sock.Socket,
	params *debugger.SetBreakpointParams,
) (debugger.SetBreakpointResult, error) {
	command := &sock.Command{
		Method: "Debugger.setBreakpoint",
		Params: params,
	}
	result := debugger.SetBreakpointResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetBreakpointByURL sets JavaScript breakpoint at given location specified either by URL or URL
regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and
returned in locations property. Further matching script parsing will result in subsequent
breakpointResolved events issued. This logical breakpoint will survive page reloads.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointByUrl
*/
func (Debugger) SetBreakpointByURL(
	socket *sock.Socket,
	params *debugger.SetBreakpointByURLParams,
) (debugger.SetBreakpointByURLResult, error) {
	command := &sock.Command{
		Method: "Debugger.setBreakpointByUrl",
		Params: params,
	}
	result := debugger.SetBreakpointByURLResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetBreakpointsActive activates / deactivates all breakpoints on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
*/
func (Debugger) SetBreakpointsActive(
	socket *sock.Socket,
	params *debugger.SetBreakpointsActiveParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setBreakpointsActive",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPauseOnExceptions defines the pause on exceptions state. Can be set to stop on all exceptions,
uncaught exceptions or no exceptions. Initial pause on exceptions state is none.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
*/
func (Debugger) SetPauseOnExceptions(
	socket *sock.Socket,
	params *debugger.SetPauseOnExceptionsParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setPauseOnExceptions",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetReturnValue changes return value in top frame. Available only at return break position.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue
*/
func (Debugger) SetReturnValue(
	socket *sock.Socket,
	params *debugger.SetReturnValueParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setReturnValue",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetScriptSource edits JavaScript source live.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
*/
func (Debugger) SetScriptSource(
	socket *sock.Socket,
	params *debugger.SetScriptSourceParams,
) (debugger.SetScriptSourceResult, error) {
	command := &sock.Command{
		Method: "Debugger.setScriptSource",
		Params: params,
	}
	result := debugger.SetScriptSourceResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetSkipAllPauses makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
*/
func (Debugger) SetSkipAllPauses(
	socket *sock.Socket,
	params *debugger.SetSkipAllPausesParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setSkipAllPauses",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetVariableValue changes value of variable in a callframe. Object-based scopes are not supported and
must be mutated manually.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
*/
func (Debugger) SetVariableValue(
	socket *sock.Socket,
	params *debugger.SetVariableValueParams,
) error {
	command := &sock.Command{
		Method: "Debugger.setVariableValue",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StepInto steps into the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
*/
func (Debugger) StepInto(
	socket *sock.Socket,
	params *debugger.StepIntoParams,
) error {
	command := &sock.Command{
		Method: "Debugger.stepInto",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StepOut steps out of the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOut
*/
func (Debugger) StepOut(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Debugger.stepOut",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StepOver steps over the statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOver
*/
func (Debugger) StepOver(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Debugger.stepOver",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnBreakpointResolved adds a handler to the Debugger.breakpointResolved event.
Debugger.breakpointResolved fires when breakpoint is resolved to an actual script and location.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-breakpointResolved
*/
func (Debugger) OnBreakpointResolved(
	socket *sock.Socket,
	callback func(event *debugger.BreakpointResolvedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.breakpointResolved",
		func(name string, params []byte) {
			event := &debugger.BreakpointResolvedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnPaused adds a handler to the Debugger.paused event. Debugger.paused fires when the virtual machine
stopped on breakpoint or exception or any other stop criteria.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-paused
*/
func (Debugger) OnPaused(
	socket *sock.Socket,
	callback func(event *debugger.PausedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.paused",
		func(name string, params []byte) {
			event := &debugger.PausedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnResumed adds a handler to the Debugger.resumed event. Debugger.resumed fires when the virtual
machine resumes execution.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-resumed
*/
func (Debugger) OnResumed(
	socket *sock.Socket,
	callback func(event *debugger.ResumedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.resumed",
		func(name string, params []byte) {
			event := &debugger.ResumedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScriptFailedToParse adds a handler to the Debugger.scriptFailedToParse event.
Debugger.scriptFailedToParse fires when the virtual machine fails to parse the script.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptFailedToParse
*/
func (Debugger) OnScriptFailedToParse(
	socket *sock.Socket,
	callback func(event *debugger.ScriptFailedToParseEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.scriptFailedToParse",
		func(name string, params []byte) {
			event := &debugger.ScriptFailedToParseEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScriptParsed adds a handler to the Debugger.ScriptParsed event. Debugger.ScriptParsed fires when
virtual machine parses script. This event is also fired for all known and uncollected scripts upon
enabling debugger.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptParsed
*/
func (Debugger) OnScriptParsed(
	socket *sock.Socket,
	callback func(event *debugger.ScriptParsedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.scriptParsed",
		func(name string, params []byte) {
			event := &debugger.ScriptParsedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
