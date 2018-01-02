package socket

import (
	"encoding/json"

	runtime "github.com/mkenney/go-chrome/cdtp/runtime"
	log "github.com/sirupsen/logrus"
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
) (*runtime.AwaitPromiseResult, error) {
	command := NewCommand("Runtime.awaitPromise", params)
	result := &runtime.AwaitPromiseResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CallFunctionOn calls a function with given declaration on the given object.
Object group of the result is inherited from the target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
func (protocol *RuntimeProtocol) CallFunctionOn(
	params *runtime.CallFunctionOnParams,
) (*runtime.CallFunctionOnResult, error) {
	command := NewCommand("Runtime.callFunctionOn", params)
	result := &runtime.CallFunctionOnResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CompileScript compiles an expression.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
func (protocol *RuntimeProtocol) CompileScript(
	params *runtime.CompileScriptParams,
) (*runtime.CompileScriptResult, error) {
	command := NewCommand("Runtime.compileScript", params)
	result := &runtime.CompileScriptResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Disable disables reporting of execution contexts creation.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-disable
*/
func (protocol *RuntimeProtocol) Disable() error {
	command := NewCommand("Runtime.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DiscardConsoleEntries discards collected exceptions and console API calls.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-discardConsoleEntries
*/
func (protocol *RuntimeProtocol) DiscardConsoleEntries() error {
	command := NewCommand("Runtime.discardConsoleEntries", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables reporting of execution contexts creation by means of the
Runtime.executionContextCreated event. When the reporting gets enabled the event
will be sent immediately for each existing execution context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-enable
*/
func (protocol *RuntimeProtocol) Enable() error {
	command := NewCommand("Runtime.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Evaluate evaluates expression on global object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
func (protocol *RuntimeProtocol) Evaluate(
	params *runtime.EvaluateParams,
) (*runtime.EvaluateResult, error) {
	command := NewCommand("Runtime.evaluate", params)
	result := &runtime.EvaluateResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetProperties returns properties of a given object. Object group of the result
is inherited from the target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
func (protocol *RuntimeProtocol) GetProperties(
	params *runtime.GetPropertiesParams,
) (*runtime.GetPropertiesResult, error) {
	command := NewCommand("Runtime.getProperties", params)
	result := &runtime.GetPropertiesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GlobalLexicalScopeNames returns all let, const and class variables from global
scope.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
func (protocol *RuntimeProtocol) GlobalLexicalScopeNames(
	params *runtime.GlobalLexicalScopeNamesParams,
) (*runtime.GlobalLexicalScopeNamesResult, error) {
	command := NewCommand("Runtime.globalLexicalScopeNames", params)
	result := &runtime.GlobalLexicalScopeNamesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
QueryObjects returns objects for a given prototype ID.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
func (protocol *RuntimeProtocol) QueryObjects(
	params *runtime.QueryObjectsParams,
) (*runtime.QueryObjectsResult, error) {
	command := NewCommand("Runtime.queryObjects", params)
	result := &runtime.QueryObjectsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
ReleaseObject releases remote object with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
*/
func (protocol *RuntimeProtocol) ReleaseObject(
	params *runtime.ReleaseObjectParams,
) error {
	command := NewCommand("Runtime.releaseObject", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ReleaseObjectGroup releases all remote objects that belong to a given group.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
*/
func (protocol *RuntimeProtocol) ReleaseObjectGroup(
	params *runtime.ReleaseObjectGroupParams,
) error {
	command := NewCommand("Runtime.releaseObjectGroup", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RunIfWaitingForDebugger tells inspected instance to run if it was waiting for
debugger to attach.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runIfWaitingForDebugger
*/
func (protocol *RuntimeProtocol) RunIfWaitingForDebugger() error {
	command := NewCommand("Runtime.runIfWaitingForDebugger", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RunScript runs the script with given ID in a given context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
func (protocol *RuntimeProtocol) RunScript(
	params *runtime.RunScriptParams,
) (*runtime.RunScriptResult, error) {
	command := NewCommand("Runtime.runScript", params)
	result := &runtime.RunScriptResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetCustomObjectFormatterEnabled is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
EXPERIMENTAL.
*/
func (protocol *RuntimeProtocol) SetCustomObjectFormatterEnabled(
	params *runtime.SetCustomObjectFormatterEnabledParams,
) error {
	command := NewCommand("Runtime.setCustomObjectFormatterEnabled", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
