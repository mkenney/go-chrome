package socket

import (
	"encoding/json"

	debugger "github.com/mkenney/go-chrome/tot/cdtp/debugger"
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
) chan *debugger.ContinueToLocationResult {
	resultChan := make(chan *debugger.ContinueToLocationResult)
	command := NewCommand(protocol.Socket, "Debugger.continueToLocation", params)
	result := &debugger.ContinueToLocationResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables debugger for given page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-disable
*/
func (protocol *DebuggerProtocol) Disable() chan *debugger.DisableResult {
	resultChan := make(chan *debugger.DisableResult)
	command := NewCommand(protocol.Socket, "Debugger.disable", nil)
	result := &debugger.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables debugger for the given page. Clients should not assume that the
debugging has been enabled until the result for this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-enable
*/
func (protocol *DebuggerProtocol) Enable() chan *debugger.EnableResult {
	resultChan := make(chan *debugger.EnableResult)
	command := NewCommand(protocol.Socket, "Debugger.enable", nil)
	result := &debugger.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
EvaluateOnCallFrame evaluates expression on a given call frame.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
*/
func (protocol *DebuggerProtocol) EvaluateOnCallFrame(
	params *debugger.EvaluateOnCallFrameParams,
) chan *debugger.EvaluateOnCallFrameResult {
	resultChan := make(chan *debugger.EvaluateOnCallFrameResult)
	command := NewCommand(protocol.Socket, "Debugger.evaluateOnCallFrame", params)
	result := &debugger.EvaluateOnCallFrameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetPossibleBreakpoints returns possible locations for breakpoint. scriptId in
start and end range locations should be the same.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
*/
func (protocol *DebuggerProtocol) GetPossibleBreakpoints(
	params *debugger.GetPossibleBreakpointsParams,
) chan *debugger.GetPossibleBreakpointsResult {
	resultChan := make(chan *debugger.GetPossibleBreakpointsResult)
	command := NewCommand(protocol.Socket, "Debugger.getPossibleBreakpoints", params)
	result := &debugger.GetPossibleBreakpointsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetScriptSource returns source for the script with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
*/
func (protocol *DebuggerProtocol) GetScriptSource(
	params *debugger.GetScriptSourceParams,
) chan *debugger.GetScriptSourceResult {
	resultChan := make(chan *debugger.GetScriptSourceResult)
	command := NewCommand(protocol.Socket, "Debugger.getScriptSource", params)
	result := &debugger.GetScriptSourceResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetStackTrace returns stack trace with given stackTraceId.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace EXPERIMENTAL.
*/
func (protocol *DebuggerProtocol) GetStackTrace(
	params *debugger.GetStackTraceParams,
) chan *debugger.GetStackTraceResult {
	resultChan := make(chan *debugger.GetStackTraceResult)
	command := NewCommand(protocol.Socket, "Debugger.getStackTrace", params)
	result := &debugger.GetStackTraceResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Pause stops on the next JavaScript statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pause
*/
func (protocol *DebuggerProtocol) Pause() chan *debugger.PauseResult {
	resultChan := make(chan *debugger.PauseResult)
	command := NewCommand(protocol.Socket, "Debugger.pause", nil)
	result := &debugger.PauseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
PauseOnAsyncCall is experimental

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall EXPERIMENTAL.
*/
func (protocol *DebuggerProtocol) PauseOnAsyncCall(
	params *debugger.PauseOnAsyncCallParams,
) chan *debugger.PauseOnAsyncCallResult {
	resultChan := make(chan *debugger.PauseOnAsyncCallResult)
	command := NewCommand(protocol.Socket, "Debugger.pauseOnAsyncCall", params)
	result := &debugger.PauseOnAsyncCallResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveBreakpoint removes JavaScript breakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
*/
func (protocol *DebuggerProtocol) RemoveBreakpoint(
	params *debugger.RemoveBreakpointParams,
) chan *debugger.RemoveBreakpointResult {
	resultChan := make(chan *debugger.RemoveBreakpointResult)
	command := NewCommand(protocol.Socket, "Debugger.removeBreakpoint", params)
	result := &debugger.RemoveBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RestartFrame restarts particular call frame from the beginning.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
*/
func (protocol *DebuggerProtocol) RestartFrame(
	params *debugger.RestartFrameParams,
) chan *debugger.RestartFrameResult {
	resultChan := make(chan *debugger.RestartFrameResult)
	command := NewCommand(protocol.Socket, "Debugger.restartFrame", params)
	result := &debugger.RestartFrameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Resume resumes JavaScript execution.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-resume
*/
func (protocol *DebuggerProtocol) Resume() chan *debugger.ResumeResult {
	resultChan := make(chan *debugger.ResumeResult)
	command := NewCommand(protocol.Socket, "Debugger.resume", nil)
	result := &debugger.ResumeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
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
func (protocol *DebuggerProtocol) ScheduleStepIntoAsync() chan *debugger.ScheduleStepIntoAsyncResult {
	resultChan := make(chan *debugger.ScheduleStepIntoAsyncResult)
	command := NewCommand(protocol.Socket, "Debugger.scheduleStepIntoAsync", nil)
	result := &debugger.ScheduleStepIntoAsyncResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SearchInContent searches for given string in script content.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
*/
func (protocol *DebuggerProtocol) SearchInContent(
	params *debugger.SearchInContentParams,
) chan *debugger.SearchInContentResult {
	resultChan := make(chan *debugger.SearchInContentResult)
	command := NewCommand(protocol.Socket, "Debugger.searchInContent", params)
	result := &debugger.SearchInContentResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetAsyncCallStackDepth enables or disables async call stacks tracking.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
*/
func (protocol *DebuggerProtocol) SetAsyncCallStackDepth(
	params *debugger.SetAsyncCallStackDepthParams,
) chan *debugger.SetAsyncCallStackDepthResult {
	resultChan := make(chan *debugger.SetAsyncCallStackDepthResult)
	command := NewCommand(protocol.Socket, "Debugger.setAsyncCallStackDepth", params)
	result := &debugger.SetAsyncCallStackDepthResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
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
) chan *debugger.SetBlackboxPatternsResult {
	resultChan := make(chan *debugger.SetBlackboxPatternsResult)
	command := NewCommand(protocol.Socket, "Debugger.setBlackboxPatterns", params)
	result := &debugger.SetBlackboxPatternsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
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
) chan *debugger.SetBlackboxedRangesResult {
	resultChan := make(chan *debugger.SetBlackboxedRangesResult)
	command := NewCommand(protocol.Socket, "Debugger.setBlackboxedRanges", params)
	result := &debugger.SetBlackboxedRangesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetBreakpoint sets JavaScript breakpoint at a given location.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
*/
func (protocol *DebuggerProtocol) SetBreakpoint(
	params *debugger.SetBreakpointParams,
) chan *debugger.SetBreakpointResult {
	resultChan := make(chan *debugger.SetBreakpointResult)
	command := NewCommand(protocol.Socket, "Debugger.setBreakpoint", params)
	result := &debugger.SetBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
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
) chan *debugger.SetBreakpointByURLResult {
	resultChan := make(chan *debugger.SetBreakpointByURLResult)
	command := NewCommand(protocol.Socket, "Debugger.setBreakpointByUrl", params)
	result := &debugger.SetBreakpointByURLResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetBreakpointsActive activates / deactivates all breakpoints on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
*/
func (protocol *DebuggerProtocol) SetBreakpointsActive(
	params *debugger.SetBreakpointsActiveParams,
) chan *debugger.SetBreakpointsActiveResult {
	resultChan := make(chan *debugger.SetBreakpointsActiveResult)
	command := NewCommand(protocol.Socket, "Debugger.setBreakpointsActive", params)
	result := &debugger.SetBreakpointsActiveResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetPauseOnExceptions defines the pause on exceptions state. Can be set to stop on all exceptions,
uncaught exceptions or no exceptions. Initial pause on exceptions state is none.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
*/
func (protocol *DebuggerProtocol) SetPauseOnExceptions(
	params *debugger.SetPauseOnExceptionsParams,
) chan *debugger.SetPauseOnExceptionsResult {
	resultChan := make(chan *debugger.SetPauseOnExceptionsResult)
	command := NewCommand(protocol.Socket, "Debugger.setPauseOnExceptions", params)
	result := &debugger.SetPauseOnExceptionsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetReturnValue changes return value in top frame. Available only at return break position.


https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue EXPERIMENTAL.
*/
func (protocol *DebuggerProtocol) SetReturnValue(
	params *debugger.SetReturnValueParams,
) chan *debugger.SetReturnValueResult {
	resultChan := make(chan *debugger.SetReturnValueResult)
	command := NewCommand(protocol.Socket, "Debugger.setReturnValue", params)
	result := &debugger.SetReturnValueResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetScriptSource edits JavaScript source live.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
*/
func (protocol *DebuggerProtocol) SetScriptSource(
	params *debugger.SetScriptSourceParams,
) chan *debugger.SetScriptSourceResult {
	resultChan := make(chan *debugger.SetScriptSourceResult)
	command := NewCommand(protocol.Socket, "Debugger.setScriptSource", params)
	result := &debugger.SetScriptSourceResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetSkipAllPauses makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
*/
func (protocol *DebuggerProtocol) SetSkipAllPauses(
	params *debugger.SetSkipAllPausesParams,
) chan *debugger.SetSkipAllPausesResult {
	resultChan := make(chan *debugger.SetSkipAllPausesResult)
	command := NewCommand(protocol.Socket, "Debugger.setSkipAllPauses", params)
	result := &debugger.SetSkipAllPausesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetVariableValue changes value of variable in a callframe. Object-based scopes are not supported and
must be mutated manually.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
*/
func (protocol *DebuggerProtocol) SetVariableValue(
	params *debugger.SetVariableValueParams,
) chan *debugger.SetVariableValueResult {
	resultChan := make(chan *debugger.SetVariableValueResult)
	command := NewCommand(protocol.Socket, "Debugger.setVariableValue", params)
	result := &debugger.SetVariableValueResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StepInto steps into the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
*/
func (protocol *DebuggerProtocol) StepInto(
	params *debugger.StepIntoParams,
) chan *debugger.StepIntoResult {
	resultChan := make(chan *debugger.StepIntoResult)
	command := NewCommand(protocol.Socket, "Debugger.stepInto", params)
	result := &debugger.StepIntoResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StepOut steps out of the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOut
*/
func (protocol *DebuggerProtocol) StepOut() chan *debugger.StepOutResult {
	resultChan := make(chan *debugger.StepOutResult)
	command := NewCommand(protocol.Socket, "Debugger.stepOut", nil)
	result := &debugger.StepOutResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StepOver steps over the statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOver
*/
func (protocol *DebuggerProtocol) StepOver() chan *debugger.StepOverResult {
	resultChan := make(chan *debugger.StepOverResult)
	command := NewCommand(protocol.Socket, "Debugger.stepOver", nil)
	result := &debugger.StepOverResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
