package runtime

/*
AwaitPromiseParams represents Runtime.awaitPromise parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
*/
type AwaitPromiseParams struct {
	// Identifier of the promise.
	PromiseObjectID RemoteObjectID `json:"promiseObjectId"`

	// Optional. Whether the result is expected to be a JSON object that should
	// be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

/*
AwaitPromiseResult represents the result of calls to Runtime.awaitPromise.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
*/
type AwaitPromiseResult struct {
	// Promise result. Will contain rejected value if promise was rejected.
	Result *RemoteObject `json:"result"`

	// Exception details if stack strace is available.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CallFunctionOnParams represents Runtime.callFunctionOn parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
type CallFunctionOnParams struct {
	// Declaration of the function to call.
	FunctionDeclaration string `json:"functionDeclaration"`

	// Optional. Identifier of the object to call function on. Either objectID
	// or executionContextID should be specified.
	ObjectID RemoteObjectID `json:"objectId,omitempty"`

	// Optional. Call arguments. All call arguments must belong to the same
	// JavaScript world as the target object.
	Arguments []*CallArgument `json:"arguments,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not
	// reported and do not pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Whether the result is expected to be a JSON object which should be sent
	// by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	// EXPERIMENTAL.
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should be treated as initiated by user in the
	// UI.
	UserGesture bool `json:"userGesture,omitempty"`

	// Optional. Whether execution should await for resulting value and return
	// once awaited promise is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`

	// Optional. Specifies execution context which global object will be used to
	// call function on. Either executionContextID or objectID should be
	// specified.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple
	// objects. If objectGroup is not specified and objectID is, objectGroup
	// will be inherited from object.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
CallFunctionOnResult represents the result of calls to Runtime.callFunctionOn.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
type CallFunctionOnResult struct {
	// Call result.
	Result *RemoteObject `json:"result"`

	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CompileScriptParams represents Runtime.compileScript parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
type CompileScriptParams struct {
	// Expression to compile.
	Expression string `json:"expression"`

	// Source url to be set for the script.
	SourceURL string `json:"sourceURL"`

	// Specifies whether the compiled script should be persisted.
	PersistScript bool `json:"persistScript"`

	// Optional. Specifies in which execution context to perform script run. If
	// the parameter is omitted the evaluation will be performed in the context
	// of the inspected page.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
CompileScriptResult represents the result of calls to Runtime.compileScript.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
type CompileScriptResult struct {
	// ID of the script.
	ScriptID ScriptID `json:"scriptId"`

	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to Runtime.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DiscardConsoleEntriesResult represents the result of calls to Runtime.discardConsoleEntries.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-discardConsoleEntries
*/
type DiscardConsoleEntriesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Runtime.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EvaluateParams represents Runtime.evaluate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
type EvaluateParams struct {
	// Expression to evaluate.
	Expression string `json:"expression"`

	// Optional. Symbolic group name that can be used to release multiple
	// objects.
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. Determines whether Command Line API should be available during
	// the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not
	// reported and do not pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Specifies in which execution context to perform evaluation. If
	// the parameter is omitted the evaluation will be performed in the context
	// of the inspected page.
	ContextID ExecutionContextID `json:"contextId,omitempty"`

	// Optional. Whether the result is expected to be a JSON object that should
	// be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	// EXPERIMENTAL.
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should be treated as initiated by user in the
	// UI.
	UserGesture bool `json:"userGesture,omitempty"`

	// Optional. Whether execution should await for resulting value and return
	// once awaited promise is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

/*
EvaluateResult represents the result of calls to Runtime.evaluate.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
type EvaluateResult struct {
	// Evaluation result.
	Result *RemoteObject `json:"result"`

	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetPropertiesParams represents Runtime.getProperties parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
type GetPropertiesParams struct {
	// Identifier of the object to return properties for.
	ObjectID RemoteObjectID `json:"objectId"`

	// Optional. If true, returns properties belonging only to the element
	// itself, not to its prototype chain.
	OwnProperties bool `json:"ownProperties,omitempty"`

	// Optional. If true, returns accessor properties (with getter/setter) only;
	// internal properties are not returned either. EXPERIMENTAL.
	AccessorPropertiesOnly bool `json:"accessorPropertiesOnly,omitempty"`

	// Optional. Whether preview should be generated for the results.
	// EXPERIMENTAL.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

/*
GetPropertiesResult represents the result of calls to Runtime.getProperties.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
type GetPropertiesResult struct {
	// Object properties.
	Result []*PropertyDescriptor `json:"result"`

	// Internal object properties (only of the element itself).
	InternalProperties []*InternalPropertyDescriptor `json:"internalProperties"`

	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GlobalLexicalScopeNamesParams represents Runtime.globalLexicalScopeNames parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
type GlobalLexicalScopeNamesParams struct {
	// Optional. Specifies in which execution context to lookup global scope variables.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
GlobalLexicalScopeNamesResult represents the result of calls to Runtime.globalLexicalScopeNames.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
type GlobalLexicalScopeNamesResult struct {
	Names []string `json:"names"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
QueryObjectsParams represents Runtime.queryObjects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
type QueryObjectsParams struct {
	// Identifier of the prototype to return objects for.
	PrototypeObjectID RemoteObjectID `json:"prototypeObjectId"`
}

/*
QueryObjectsResult represents the result of calls to Runtime.queryObjects.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
type QueryObjectsResult struct {
	// Identifier of the object to release.
	ObjectID RemoteObjectID `json:"objectId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ReleaseObjectParams represents Runtime.releaseObject parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
*/
type ReleaseObjectParams struct {
	// Identifier of the object to release.
	ObjectID RemoteObjectID `json:"objectId"`
}

/*
ReleaseObjectResult represents the result of calls to Runtime.releaseObject.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
*/
type ReleaseObjectResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ReleaseObjectGroupParams represents Runtime.releaseObjectGroup parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
*/
type ReleaseObjectGroupParams struct {
	// Symbolic object group name.
	ObjectGroup string `json:"objectGroup"`
}

/*
ReleaseObjectGroupResult represents the result of calls to Runtime.releaseObjectGroup.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
*/
type ReleaseObjectGroupResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RunIfWaitingForDebuggerResult represents the result of calls to Runtime.runIfWaitingForDebugger.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runIfWaitingForDebugger
*/
type RunIfWaitingForDebuggerResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RunScriptParams represents Runtime.runScript parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
type RunScriptParams struct {
	// ID of the script to run.
	ScriptID ScriptID `json:"scriptId"`

	// Optional. Specifies in which execution context to perform script run. If
	// the parameter is omitted the evaluation will be performed in the context
	// of the inspected page.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple
	// objects.
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not
	// reported and do not pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Determines whether Command Line API should be available during
	// the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. Whether the result is expected to be a JSON object which should
	// be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should await for resulting value and return
	// once awaited promise is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

/*
RunScriptResult represents the result of calls to Runtime.runScript.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
type RunScriptResult struct {
	// Identifier of the object to release.
	ObjectID RemoteObjectID `json:"objectId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetCustomObjectFormatterEnabledParams represents Runtime.setCustomObjectFormatterEnabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
*/
type SetCustomObjectFormatterEnabledParams struct {
	// Run result.
	Result *RemoteObject `json:"result"`

	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`
}

/*
SetCustomObjectFormatterEnabledResult represents the result of calls to
Runtime.setCustomObjectFormatterEnabled.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
*/
type SetCustomObjectFormatterEnabledResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
