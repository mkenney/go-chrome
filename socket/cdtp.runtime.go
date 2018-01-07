package socket

import (
	"encoding/json"

	runtime "github.com/mkenney/go-chrome/cdtp/runtime"
)

/*
RuntimeProtocol provides a namespace for the Chrome Runtime protocol methods.
The Runtime protocol exposes JavaScript runtime by means of remote evaluation
and mirror objects. Evaluation results are returned as mirror object that expose
object type, string representation and unique identifier that can be used for
further object reference. Original objects are maintained in memory unless they
are either explicitly released or are released along with the other objects in
their object group.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/
*/
type RuntimeProtocol struct {
	Socket Socketer
}

/*
AwaitPromise adds handler to promise with given promise object ID.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
*/
func (protocol *RuntimeProtocol) AwaitPromise(
	params *runtime.AwaitPromiseParams,
) chan *runtime.AwaitPromiseResult {
	resultChan := make(chan *runtime.AwaitPromiseResult)
	command := NewCommand(protocol.Socket, "Runtime.awaitPromise", params)
	result := &runtime.AwaitPromiseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
CallFunctionOn calls a function with given declaration on the given object.
Object group of the result is inherited from the target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
func (protocol *RuntimeProtocol) CallFunctionOn(
	params *runtime.CallFunctionOnParams,
) chan *runtime.CallFunctionOnResult {
	resultChan := make(chan *runtime.CallFunctionOnResult)
	command := NewCommand(protocol.Socket, "Runtime.callFunctionOn", params)
	result := &runtime.CallFunctionOnResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
CompileScript compiles an expression.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
func (protocol *RuntimeProtocol) CompileScript(
	params *runtime.CompileScriptParams,
) chan *runtime.CompileScriptResult {
	resultChan := make(chan *runtime.CompileScriptResult)
	command := NewCommand(protocol.Socket, "Runtime.compileScript", params)
	result := &runtime.CompileScriptResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Disable disables reporting of execution contexts creation.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-disable
*/
func (protocol *RuntimeProtocol) Disable() chan *runtime.DisableResult {
	resultChan := make(chan *runtime.DisableResult)
	command := NewCommand(protocol.Socket, "Runtime.disable", nil)
	result := &runtime.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DiscardConsoleEntries discards collected exceptions and console API calls.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-discardConsoleEntries
*/
func (protocol *RuntimeProtocol) DiscardConsoleEntries() chan *runtime.DiscardConsoleEntriesResult {
	resultChan := make(chan *runtime.DiscardConsoleEntriesResult)
	command := NewCommand(protocol.Socket, "Runtime.discardConsoleEntries", nil)
	result := &runtime.DiscardConsoleEntriesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables reporting of execution contexts creation by means of the
Runtime.executionContextCreated event. When the reporting gets enabled the event
will be sent immediately for each existing execution context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-enable
*/
func (protocol *RuntimeProtocol) Enable() chan *runtime.EnableResult {
	resultChan := make(chan *runtime.EnableResult)
	command := NewCommand(protocol.Socket, "Runtime.enable", nil)
	result := &runtime.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Evaluate evaluates expression on global object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
func (protocol *RuntimeProtocol) Evaluate(
	params *runtime.EvaluateParams,
) chan *runtime.EvaluateResult {
	resultChan := make(chan *runtime.EvaluateResult)
	command := NewCommand(protocol.Socket, "Runtime.evaluate", params)
	result := &runtime.EvaluateResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetProperties returns properties of a given object. Object group of the result
is inherited from the target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
func (protocol *RuntimeProtocol) GetProperties(
	params *runtime.GetPropertiesParams,
) chan *runtime.GetPropertiesResult {
	resultChan := make(chan *runtime.GetPropertiesResult)
	command := NewCommand(protocol.Socket, "Runtime.getProperties", params)
	result := &runtime.GetPropertiesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GlobalLexicalScopeNames returns all let, const and class variables from global
scope.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
func (protocol *RuntimeProtocol) GlobalLexicalScopeNames(
	params *runtime.GlobalLexicalScopeNamesParams,
) chan *runtime.GlobalLexicalScopeNamesResult {
	resultChan := make(chan *runtime.GlobalLexicalScopeNamesResult)
	command := NewCommand(protocol.Socket, "Runtime.globalLexicalScopeNames", params)
	result := &runtime.GlobalLexicalScopeNamesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
QueryObjects returns objects for a given prototype ID.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
func (protocol *RuntimeProtocol) QueryObjects(
	params *runtime.QueryObjectsParams,
) chan *runtime.QueryObjectsResult {
	resultChan := make(chan *runtime.QueryObjectsResult)
	command := NewCommand(protocol.Socket, "Runtime.queryObjects", params)
	result := &runtime.QueryObjectsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
ReleaseObject releases remote object with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
*/
func (protocol *RuntimeProtocol) ReleaseObject(
	params *runtime.ReleaseObjectParams,
) chan *runtime.ReleaseObjectResult {
	resultChan := make(chan *runtime.ReleaseObjectResult)
	command := NewCommand(protocol.Socket, "Runtime.releaseObject", params)
	result := &runtime.ReleaseObjectResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
ReleaseObjectGroup releases all remote objects that belong to a given group.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
*/
func (protocol *RuntimeProtocol) ReleaseObjectGroup(
	params *runtime.ReleaseObjectGroupParams,
) chan *runtime.ReleaseObjectGroupResult {
	resultChan := make(chan *runtime.ReleaseObjectGroupResult)
	command := NewCommand(protocol.Socket, "Runtime.releaseObjectGroup", params)
	result := &runtime.ReleaseObjectGroupResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RunIfWaitingForDebugger tells inspected instance to run if it was waiting for
debugger to attach.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runIfWaitingForDebugger
*/
func (protocol *RuntimeProtocol) RunIfWaitingForDebugger() chan *runtime.RunIfWaitingForDebuggerResult {
	resultChan := make(chan *runtime.RunIfWaitingForDebuggerResult)
	command := NewCommand(protocol.Socket, "Runtime.runIfWaitingForDebugger", nil)
	result := &runtime.RunIfWaitingForDebuggerResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RunScript runs the script with given ID in a given context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
func (protocol *RuntimeProtocol) RunScript(
	params *runtime.RunScriptParams,
) chan *runtime.RunScriptResult {
	resultChan := make(chan *runtime.RunScriptResult)
	command := NewCommand(protocol.Socket, "Runtime.runScript", params)
	result := &runtime.RunScriptResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetCustomObjectFormatterEnabled is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
EXPERIMENTAL.
*/
func (protocol *RuntimeProtocol) SetCustomObjectFormatterEnabled(
	params *runtime.SetCustomObjectFormatterEnabledParams,
) chan *runtime.SetCustomObjectFormatterEnabledResult {
	resultChan := make(chan *runtime.SetCustomObjectFormatterEnabledResult)
	command := NewCommand(protocol.Socket, "Runtime.setCustomObjectFormatterEnabled", params)
	result := &runtime.SetCustomObjectFormatterEnabledResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
OnConsoleAPICalled adds a handler to the Runtime.consoleAPICalled event.
Runtime.consoleAPICalled fires when the console API is called.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-consoleAPICalled
*/
func (protocol *RuntimeProtocol) OnConsoleAPICalled(
	callback func(event *runtime.ConsoleAPICalledEvent),
) {
	handler := NewEventHandler(
		"Runtime.consoleAPICalled",
		func(response *Response) {
			event := &runtime.ConsoleAPICalledEvent{}
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
OnExceptionRevoked adds a handler to the Runtime.exceptionRevoked event.
Runtime.exceptionRevoked fires when an unhandled exception is revoked.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionRevoked
*/
func (protocol *RuntimeProtocol) OnExceptionRevoked(
	callback func(event *runtime.ExceptionRevokedEvent),
) {
	handler := NewEventHandler(
		"Runtime.exceptionRevoked",
		func(response *Response) {
			event := &runtime.ExceptionRevokedEvent{}
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
OnExceptionThrown adds a handler to the Runtime.exceptionThrown event.
Runtime.exceptionThrown fires when an exception is thrown and is unhandled.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionThrown
*/
func (protocol *RuntimeProtocol) OnExceptionThrown(
	callback func(event *runtime.ExceptionThrownEvent),
) {
	handler := NewEventHandler(
		"Runtime.exceptionThrown",
		func(response *Response) {
			event := &runtime.ExceptionThrownEvent{}
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
OnExecutionContextCreated adds a handler to the Runtime.executionContextCreated
event. Runtime.executionContextCreated fires when a new execution context is
created.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextCreated
*/
func (protocol *RuntimeProtocol) OnExecutionContextCreated(
	callback func(event *runtime.ExecutionContextCreatedEvent),
) {
	handler := NewEventHandler(
		"Runtime.executionContextCreated",
		func(response *Response) {
			event := &runtime.ExecutionContextCreatedEvent{}
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
OnExecutionContextDestroyed adds a handler to the Runtime.executionContextDestroyed
event. Runtime.executionContextDestroyed fires when execution context is
destroyed.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextDestroyed
*/
func (protocol *RuntimeProtocol) OnExecutionContextDestroyed(
	callback func(event *runtime.ExecutionContextDestroyedEvent),
) {
	handler := NewEventHandler(
		"Runtime.executionContextDestroyed",
		func(response *Response) {
			event := &runtime.ExecutionContextDestroyedEvent{}
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
OnExecutionContextsCleared adds a handler to the Runtime.executionContextsCleared
event. Runtime.executionContextsCleared fires when all executionContexts were
cleared in browser.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextsCleared
*/
func (protocol *RuntimeProtocol) OnExecutionContextsCleared(
	callback func(event *runtime.ExecutionContextsClearedEvent),
) {
	handler := NewEventHandler(
		"Runtime.executionContextsCleared",
		func(response *Response) {
			event := &runtime.ExecutionContextsClearedEvent{}
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
OnInspectRequested adds a handler to the Runtime.inspectRequested event.
Runtime.inspectRequested fires when an object should be inspected (for example,
as a result of inspect() command line API call).

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-inspectRequested
*/
func (protocol *RuntimeProtocol) OnInspectRequested(
	callback func(event *runtime.InspectRequestedEvent),
) {
	handler := NewEventHandler(
		"Runtime.inspectRequested",
		func(response *Response) {
			event := &runtime.InspectRequestedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
