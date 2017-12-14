package chrome

import (
	"app/chrome/protocol"
	runtime "app/chrome/runtime"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Runtime - https://chromedevtools.github.io/devtools-protocol/tot/Runtime/
Exposes JavaScript runtime by means of remote evaluation and mirror objects. Evaluation results are
returned as mirror object that expose object type, string representation and unique identifier that
can be used for further object reference. Original objects are maintained in memory unless they are
either explicitly released or are released along with the other objects in their object group.
*/
type Runtime struct{}

/*
AwaitPromise adds handler to promise with given promise object ID.
*/
func (Runtime) AwaitPromise(
	socket *Socket,
	params *runtime.AwaitPromiseParams,
) (runtime.AwaitPromiseResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) CallFunctionOn(
	socket *Socket,
	params *runtime.CallFunctionOnParams,
) (runtime.CallFunctionOnResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) CompileScript(
	socket *Socket,
	params *runtime.CompileScriptParams,
) (runtime.CompileScriptResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Runtime.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DiscardConsoleEntries discards collected exceptions and console API calls.
*/
func (Runtime) DiscardConsoleEntries(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Runtime.discardConsoleEntries",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables reporting of execution contexts creation by means of executionContextCreated event.
When the reporting gets enabled the event will be sent immediately for each existing execution
context.
*/
func (Runtime) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Runtime.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Evaluate evaluates expression on global object.
*/
func (Runtime) Evaluate(
	socket *Socket,
	params *runtime.EvaluateParams,
) (runtime.EvaluateResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) GetProperties(
	socket *Socket,
	params *runtime.GetPropertiesParams,
) (runtime.GetPropertiesResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) GlobalLexicalScopeNames(
	socket *Socket,
	params *runtime.GlobalLexicalScopeNamesParams,
) (runtime.GlobalLexicalScopeNamesResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) QueryObjects(
	socket *Socket,
	params *runtime.QueryObjectsParams,
) (runtime.QueryObjectsResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) ReleaseObject(
	socket *Socket,
	params *runtime.ReleaseObjectParams,
) error {
	command := &protocol.Command{
		Method: "Runtime.releaseObject",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ReleaseObjectGroup releases all remote objects that belong to a given group.
*/
func (Runtime) ReleaseObjectGroup(
	socket *Socket,
	params *runtime.ReleaseObjectGroupParams,
) error {
	command := &protocol.Command{
		Method: "Runtime.releaseObjectGroup",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RunIfWaitingForDebugger tells inspected instance to run if it was waiting for debugger to attach.
*/
func (Runtime) RunIfWaitingForDebugger(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Runtime.runIfWaitingForDebugger",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RunScript runs the script with given ID in a given context.
*/
func (Runtime) RunScript(
	socket *Socket,
	params *runtime.RunScriptParams,
) (runtime.RunScriptResult, error) {
	command := &protocol.Command{
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
*/
func (Runtime) SetCustomObjectFormatterEnabled(
	socket *Socket,
	params *runtime.SetCustomObjectFormatterEnabledParams,
) error {
	command := &protocol.Command{
		Method: "Runtime.setCustomObjectFormatterEnabled",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnConsoleAPICalled adds a handler to the Runtime.consoleAPICalled event. Runtime.consoleAPICalled
fires when the console API is called.
*/
func (Runtime) OnConsoleAPICalled(
	socket *Socket,
	callback func(event *runtime.ConsoleAPICalledEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Runtime) OnExceptionRevoked(
	socket *Socket,
	callback func(event *runtime.ExceptionRevokedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Runtime) OnExceptionThrown(
	socket *Socket,
	callback func(event *runtime.ExceptionThrownEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Runtime) OnExecutionContextCreated(
	socket *Socket,
	callback func(event *runtime.ExecutionContextCreatedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Runtime) OnExecutionContextDestroyed(
	socket *Socket,
	callback func(event *runtime.ExecutionContextDestroyedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Runtime) OnExecutionContextsCleared(
	socket *Socket,
	callback func(event *runtime.ExecutionContextsClearedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Runtime) OnInspectRequested(
	socket *Socket,
	callback func(event *runtime.InspectRequestedEvent),
) {
	handler := protocol.NewEventHandler(
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
