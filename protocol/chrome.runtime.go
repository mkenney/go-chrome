package protocol

import (
	"encoding/json"

	runtime "github.com/mkenney/go-chrome/protocol/runtime"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Runtime is a struct that provides a namespace for the Chrome Runtime protocol methods.

The Runtime protocol exposes JavaScript runtime by means of remote evaluation and mirror objects.
Evaluation results are returned as mirror object that expose object type, string representation and
unique identifier that can be used for further object reference. Original objects are maintained in
memory unless they are either explicitly released or are released along with the other objects in
their object group.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/
*/
var Runtime = _runtime{}

type _runtime struct{}

/*
AwaitPromise adds handler to promise with given promise object ID.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
*/
func (_runtime) AwaitPromise(
	socket sock.Socketer,
	params *runtime.AwaitPromiseParams,
) (runtime.AwaitPromiseResult, error) {
	command := sock.NewCommand("Runtime.awaitPromise", params)
	result := runtime.AwaitPromiseResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
CallFunctionOn calls a function with given declaration on the given object. Object group of the
result is inherited from the target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
func (_runtime) CallFunctionOn(
	socket sock.Socketer,
	params *runtime.CallFunctionOnParams,
) (runtime.CallFunctionOnResult, error) {
	command := sock.NewCommand("Runtime.callFunctionOn", params)
	result := runtime.CallFunctionOnResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
CompileScript compiles an expression.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
func (_runtime) CompileScript(
	socket sock.Socketer,
	params *runtime.CompileScriptParams,
) (runtime.CompileScriptResult, error) {
	command := sock.NewCommand("Runtime.compileScript", params)
	result := runtime.CompileScriptResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Disable disables reporting of execution contexts creation.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-disable
*/
func (_runtime) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Runtime.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
DiscardConsoleEntries discards collected exceptions and console API calls.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-discardConsoleEntries
*/
func (_runtime) DiscardConsoleEntries(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Runtime.discardConsoleEntries", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables reporting of execution contexts creation by means of executionContextCreated event.
When the reporting gets enabled the event will be sent immediately for each existing execution
context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-enable
*/
func (_runtime) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Runtime.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Evaluate evaluates expression on global object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
func (_runtime) Evaluate(
	socket sock.Socketer,
	params *runtime.EvaluateParams,
) (runtime.EvaluateResult, error) {
	command := sock.NewCommand("Runtime.evaluate", params)
	result := runtime.EvaluateResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetProperties returns properties of a given object. Object group of the result is inherited from the
target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
func (_runtime) GetProperties(
	socket sock.Socketer,
	params *runtime.GetPropertiesParams,
) (runtime.GetPropertiesResult, error) {
	command := sock.NewCommand("Runtime.getProperties", params)
	result := runtime.GetPropertiesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GlobalLexicalScopeNames returns all let, const and class variables from global scope.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
func (_runtime) GlobalLexicalScopeNames(
	socket sock.Socketer,
	params *runtime.GlobalLexicalScopeNamesParams,
) (runtime.GlobalLexicalScopeNamesResult, error) {
	command := sock.NewCommand("Runtime.globalLexicalScopeNames", params)
	result := runtime.GlobalLexicalScopeNamesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
QueryObjects returns objects for a given prototype ID.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
func (_runtime) QueryObjects(
	socket sock.Socketer,
	params *runtime.QueryObjectsParams,
) (runtime.QueryObjectsResult, error) {
	command := sock.NewCommand("Runtime.queryObjects", params)
	result := runtime.QueryObjectsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
ReleaseObject releases remote object with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
*/
func (_runtime) ReleaseObject(
	socket sock.Socketer,
	params *runtime.ReleaseObjectParams,
) error {
	command := sock.NewCommand("Runtime.releaseObject", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
ReleaseObjectGroup releases all remote objects that belong to a given group.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
*/
func (_runtime) ReleaseObjectGroup(
	socket sock.Socketer,
	params *runtime.ReleaseObjectGroupParams,
) error {
	command := sock.NewCommand("Runtime.releaseObjectGroup", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RunIfWaitingForDebugger tells inspected instance to run if it was waiting for debugger to attach.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runIfWaitingForDebugger
*/
func (_runtime) RunIfWaitingForDebugger(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Runtime.runIfWaitingForDebugger", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
RunScript runs the script with given ID in a given context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
func (_runtime) RunScript(
	socket sock.Socketer,
	params *runtime.RunScriptParams,
) (runtime.RunScriptResult, error) {
	command := sock.NewCommand("Runtime.runScript", params)
	result := runtime.RunScriptResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetCustomObjectFormatterEnabled EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
*/
func (_runtime) SetCustomObjectFormatterEnabled(
	socket sock.Socketer,
	params *runtime.SetCustomObjectFormatterEnabledParams,
) error {
	command := sock.NewCommand("Runtime.setCustomObjectFormatterEnabled", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnConsoleAPICalled adds a handler to the Runtime.consoleAPICalled event. Runtime.consoleAPICalled
fires when the console API is called.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-consoleAPICalled
*/
func (_runtime) OnConsoleAPICalled(
	socket sock.Socketer,
	callback func(event *runtime.ConsoleAPICalledEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.consoleAPICalled",
		func(response *sock.Response) {
			event := &runtime.ConsoleAPICalledEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnExceptionRevoked adds a handler to the Runtime.exceptionRevoked event. Runtime.exceptionRevoked
fires when an unhandled exception is revoked.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionRevoked
*/
func (_runtime) OnExceptionRevoked(
	socket sock.Socketer,
	callback func(event *runtime.ExceptionRevokedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.exceptionRevoked",
		func(response *sock.Response) {
			event := &runtime.ExceptionRevokedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnExceptionThrown adds a handler to the Runtime.exceptionThrown event. Runtime.exceptionThrown fires
when an exception is thrown and is unhandled.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionThrown
*/
func (_runtime) OnExceptionThrown(
	socket sock.Socketer,
	callback func(event *runtime.ExceptionThrownEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.exceptionThrown",
		func(response *sock.Response) {
			event := &runtime.ExceptionThrownEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnExecutionContextCreated adds a handler to the Runtime.executionContextCreated event.
Runtime.executionContextCreated fires when a new execution context is created.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextCreated
*/
func (_runtime) OnExecutionContextCreated(
	socket sock.Socketer,
	callback func(event *runtime.ExecutionContextCreatedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.executionContextCreated",
		func(response *sock.Response) {
			event := &runtime.ExecutionContextCreatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnExecutionContextDestroyed adds a handler to the Runtime.executionContextDestroyed event.
Runtime.executionContextDestroyed fires when execution context is destroyed.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextDestroyed
*/
func (_runtime) OnExecutionContextDestroyed(
	socket sock.Socketer,
	callback func(event *runtime.ExecutionContextDestroyedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.executionContextDestroyed",
		func(response *sock.Response) {
			event := &runtime.ExecutionContextDestroyedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnExecutionContextsCleared adds a handler to the Runtime.executionContextsCleared event.
Runtime.executionContextsCleared fires when all executionContexts were cleared in browser.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextsCleared
*/
func (_runtime) OnExecutionContextsCleared(
	socket sock.Socketer,
	callback func(event *runtime.ExecutionContextsClearedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.executionContextsCleared",
		func(response *sock.Response) {
			event := &runtime.ExecutionContextsClearedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnInspectRequested adds a handler to the Runtime.inspectRequested event. Runtime.inspectRequested
fires when an object should be inspected (for example, as a result of inspect() command line API
call).

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-inspectRequested
*/
func (_runtime) OnInspectRequested(
	socket sock.Socketer,
	callback func(event *runtime.InspectRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.inspectRequested",
		func(response *sock.Response) {
			event := &runtime.InspectRequestedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
