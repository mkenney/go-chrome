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
type Runtime struct{}

/*
AwaitPromise adds handler to promise with given promise object ID.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
*/
func (Runtime) AwaitPromise(
	socket *sock.Socket,
	params *runtime.AwaitPromiseParams,
) (runtime.AwaitPromiseResult, error) {
	command := &sock.Command{
		Method: "Runtime.awaitPromise",
		Params: params,
	}
	result := runtime.AwaitPromiseResult{}
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
CallFunctionOn calls a function with given declaration on the given object. Object group of the
result is inherited from the target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
func (Runtime) CallFunctionOn(
	socket *sock.Socket,
	params *runtime.CallFunctionOnParams,
) (runtime.CallFunctionOnResult, error) {
	command := &sock.Command{
		Method: "Runtime.callFunctionOn",
		Params: params,
	}
	result := runtime.CallFunctionOnResult{}
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
CompileScript compiles an expression.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
func (Runtime) CompileScript(
	socket *sock.Socket,
	params *runtime.CompileScriptParams,
) (runtime.CompileScriptResult, error) {
	command := &sock.Command{
		Method: "Runtime.compileScript",
		Params: params,
	}
	result := runtime.CompileScriptResult{}
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
Disable disables reporting of execution contexts creation.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-disable
*/
func (Runtime) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Runtime.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DiscardConsoleEntries discards collected exceptions and console API calls.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-discardConsoleEntries
*/
func (Runtime) DiscardConsoleEntries(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Runtime.discardConsoleEntries",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables reporting of execution contexts creation by means of executionContextCreated event.
When the reporting gets enabled the event will be sent immediately for each existing execution
context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-enable
*/
func (Runtime) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Runtime.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Evaluate evaluates expression on global object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
func (Runtime) Evaluate(
	socket *sock.Socket,
	params *runtime.EvaluateParams,
) (runtime.EvaluateResult, error) {
	command := &sock.Command{
		Method: "Runtime.evaluate",
		Params: params,
	}
	result := runtime.EvaluateResult{}
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
GetProperties returns properties of a given object. Object group of the result is inherited from the
target object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
func (Runtime) GetProperties(
	socket *sock.Socket,
	params *runtime.GetPropertiesParams,
) (runtime.GetPropertiesResult, error) {
	command := &sock.Command{
		Method: "Runtime.getProperties",
		Params: params,
	}
	result := runtime.GetPropertiesResult{}
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
GlobalLexicalScopeNames returns all let, const and class variables from global scope.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
func (Runtime) GlobalLexicalScopeNames(
	socket *sock.Socket,
	params *runtime.GlobalLexicalScopeNamesParams,
) (runtime.GlobalLexicalScopeNamesResult, error) {
	command := &sock.Command{
		Method: "Runtime.globalLexicalScopeNames",
		Params: params,
	}
	result := runtime.GlobalLexicalScopeNamesResult{}
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
QueryObjects returns objects for a given prototype ID.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
func (Runtime) QueryObjects(
	socket *sock.Socket,
	params *runtime.QueryObjectsParams,
) (runtime.QueryObjectsResult, error) {
	command := &sock.Command{
		Method: "Runtime.queryObjects",
		Params: params,
	}
	result := runtime.QueryObjectsResult{}
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
ReleaseObject releases remote object with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
*/
func (Runtime) ReleaseObject(
	socket *sock.Socket,
	params *runtime.ReleaseObjectParams,
) error {
	command := &sock.Command{
		Method: "Runtime.releaseObject",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ReleaseObjectGroup releases all remote objects that belong to a given group.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
*/
func (Runtime) ReleaseObjectGroup(
	socket *sock.Socket,
	params *runtime.ReleaseObjectGroupParams,
) error {
	command := &sock.Command{
		Method: "Runtime.releaseObjectGroup",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RunIfWaitingForDebugger tells inspected instance to run if it was waiting for debugger to attach.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runIfWaitingForDebugger
*/
func (Runtime) RunIfWaitingForDebugger(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Runtime.runIfWaitingForDebugger",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RunScript runs the script with given ID in a given context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
func (Runtime) RunScript(
	socket *sock.Socket,
	params *runtime.RunScriptParams,
) (runtime.RunScriptResult, error) {
	command := &sock.Command{
		Method: "Runtime.runScript",
		Params: params,
	}
	result := runtime.RunScriptResult{}
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
SetCustomObjectFormatterEnabled EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
*/
func (Runtime) SetCustomObjectFormatterEnabled(
	socket *sock.Socket,
	params *runtime.SetCustomObjectFormatterEnabledParams,
) error {
	command := &sock.Command{
		Method: "Runtime.setCustomObjectFormatterEnabled",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnConsoleAPICalled adds a handler to the Runtime.consoleAPICalled event. Runtime.consoleAPICalled
fires when the console API is called.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-consoleAPICalled
*/
func (Runtime) OnConsoleAPICalled(
	socket *sock.Socket,
	callback func(event *runtime.ConsoleAPICalledEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.consoleAPICalled",
		func(name string, params []byte) {
			event := &runtime.ConsoleAPICalledEvent{}
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
OnExceptionRevoked adds a handler to the Runtime.exceptionRevoked event. Runtime.exceptionRevoked
fires when an unhandled exception is revoked.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionRevoked
*/
func (Runtime) OnExceptionRevoked(
	socket *sock.Socket,
	callback func(event *runtime.ExceptionRevokedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.exceptionRevoked",
		func(name string, params []byte) {
			event := &runtime.ExceptionRevokedEvent{}
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
OnExceptionThrown adds a handler to the Runtime.exceptionThrown event. Runtime.exceptionThrown fires
when an exception is thrown and is unhandled.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionThrown
*/
func (Runtime) OnExceptionThrown(
	socket *sock.Socket,
	callback func(event *runtime.ExceptionThrownEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.exceptionThrown",
		func(name string, params []byte) {
			event := &runtime.ExceptionThrownEvent{}
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
OnExecutionContextCreated adds a handler to the Runtime.executionContextCreated event.
Runtime.executionContextCreated fires when a new execution context is created.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextCreated
*/
func (Runtime) OnExecutionContextCreated(
	socket *sock.Socket,
	callback func(event *runtime.ExecutionContextCreatedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.executionContextCreated",
		func(name string, params []byte) {
			event := &runtime.ExecutionContextCreatedEvent{}
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
OnExecutionContextDestroyed adds a handler to the Runtime.executionContextDestroyed event.
Runtime.executionContextDestroyed fires when execution context is destroyed.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextDestroyed
*/
func (Runtime) OnExecutionContextDestroyed(
	socket *sock.Socket,
	callback func(event *runtime.ExecutionContextDestroyedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.executionContextDestroyed",
		func(name string, params []byte) {
			event := &runtime.ExecutionContextDestroyedEvent{}
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
OnExecutionContextsCleared adds a handler to the Runtime.executionContextsCleared event.
Runtime.executionContextsCleared fires when all executionContexts were cleared in browser.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextsCleared
*/
func (Runtime) OnExecutionContextsCleared(
	socket *sock.Socket,
	callback func(event *runtime.ExecutionContextsClearedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.executionContextsCleared",
		func(name string, params []byte) {
			event := &runtime.ExecutionContextsClearedEvent{}
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
OnInspectRequested adds a handler to the Runtime.inspectRequested event. Runtime.inspectRequested
fires when an object should be inspected (for example, as a result of inspect() command line API
call).

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-inspectRequested
*/
func (Runtime) OnInspectRequested(
	socket *sock.Socket,
	callback func(event *runtime.InspectRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Runtime.inspectRequested",
		func(name string, params []byte) {
			event := &runtime.InspectRequestedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
