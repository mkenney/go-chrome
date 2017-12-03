package Runtime

import (
	"fmt"
)

/*
EvaluateParams represents Runtime.evaluate parameters.
*/
type EvaluateParams struct {
	// Expression to evaluate.
	Expression string `json:"expression"`

	// Optional. Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Specifies in which execution context to perform evaluation. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ContextID ExecutionContextID `json:"contextId,omitempty"`

	// Optional. Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`

	// Optional. Whether execution should await for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

/*
AwaitPromiseParams represents Runtime.awaitPromise parameters.
*/
type AwaitPromiseParams struct {
	// Identifier of the promise.
	PromiseObjectID RemoteObjectID `json:"promiseObjectId"`

	// Optional. Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

/*
CallFunctionOnParams represents Runtime.callFunctionOn parameters.
*/
type CallFunctionOnParams struct {
	// Declaration of the function to call.
	FunctionDeclaration string `json:"functionDeclaration"`

	// Optional. Identifier of the object to call function on. Either objectID or executionContextID
	// should be specified.
	ObjectID RemoteObjectID `json:"objectId,omitempty"`

	// Optional. Call arguments. All call arguments must belong to the same JavaScript world as the
	// target object.
	Arguments []CallArgument `json:"arguments,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not reported and do not
	// pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`

	// Optional. Whether execution should await for resulting value and return once awaited promise
	// is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`

	// Optional. Specifies execution context which global object will be used to call function on.
	// Either executionContextID or objectID should be specified.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple objects. If objectGroup is
	// not specified and objectID is, objectGroup will be inherited from object.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
GetPropertiesParams represents Runtime.getProperties parameters.
*/
type GetPropertiesParams struct {
	// Identifier of the object to return properties for.
	ObjectID RemoteObjectID `json:"objectId"`

	// Optional. If true, returns properties belonging only to the element itself, not to its
	// prototype chain.
	OwnProperties bool `json:"ownProperties,omitempty"`

	// Optional. If true, returns accessor properties (with getter/setter) only; internal properties
	// are not returned either. EXPERIMENTAL
	AccessorPropertiesOnly bool `json:"accessorPropertiesOnly,omitempty"`

	// Optional. Whether preview should be generated for the results. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

/*
ReleaseObjectParams represents Runtime.releaseObject parameters.
*/
type ReleaseObjectParams struct {
	// Identifier of the object to release.
	ObjectID RemoteObjectID `json:"objectId"`
}

/*
ReleaseObjectGroupParams represents Runtime.releaseObjectGroup parameters.
*/
type ReleaseObjectGroupParams struct {
	// Symbolic object group name.
	ObjectGroup string `json:"objectGroup"`
}

/*
CompileScriptParams represents Runtime.compileScript parameters.
*/
type CompileScriptParams struct {
	// Expression to compile.
	Expression string `json:"expression"`

	// Source url to be set for the script.
	SourceURL string `json:"sourceURL"`

	// Specifies whether the compiled script should be persisted.
	PersistScript bool `json:"persistScript"`

	// Optional. Specifies in which execution context to perform script run. If the parameter is
	// omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
RunScriptParams represents Runtime.runScript parameters.
*/
type RunScriptParams struct {
	// ID of the script to run.
	ScriptID ScriptID `json:"scriptId"`

	// Optional. Specifies in which execution context to perform script run. If the parameter is
	// omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not reported and do not
	// pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should await for resulting value and return once awaited promise
	// is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

/*
QueryObjectsParams represents Runtime.queryObjects parameters.
*/
type QueryObjectsParams struct {
	// Identifier of the prototype to return objects for.
	PrototypeObjectID RemoteObjectID `json:"prototypeObjectId"`
}

/*
GlobalLexicalScopeNamesParams represents Runtime.globalLexicalScopeNames parameters.
*/
type GlobalLexicalScopeNamesParams struct {
	// Optional. Specifies in which execution context to lookup global scope variables.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
ExecutionContextCreatedEvent represents Runtime.executionContextCreated event data.
*/
type ExecutionContextCreatedEvent struct {
	// A newly created execution context.
	Context ExecutionContextDescription `json:"context"`
}

/*
ExecutionContextDestroyedEvent represents Runtime.executionContextDestroyed event data.
*/
type ExecutionContextDestroyedEvent struct {
	// ID of the destroyed context.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`
}

/*
ExecutionContextsClearedEvent represents Runtime.executionContextsCleared event data.
*/
type ExecutionContextsClearedEvent struct{}

/*
ExceptionThrownEvent represents Runtime.exceptionThrown event data.
*/
type ExceptionThrownEvent struct {
	// Timestamp of the exception.
	Timestamp Timestamp `json:"timestamp"`

	// Exception details.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
ExceptionRevokedEvent represents Runtime.exceptionRevoked event data.
*/
type ExceptionRevokedEvent struct {
	// Reason describing why exception was revoked.
	Reason string `json:"reason"`

	// The ID of revoked exception, as reported in exceptionThrown.
	ExceptionID int `json:"exceptionId"`
}

/*
ConsoleAPICalledEvent represents Runtime.consoleAPICalled event data.
*/
type ConsoleAPICalledEvent struct {
	// Type of the call. Allowed values: log, debug, info, error, warning, dir, dirxml, table,
	// trace, clear, startGroup, startGroupCollapsed, endGroup, assert, profile, profileEnd, count,
	// timeEnd.
	Type string `json:"type"`

	// Call arguments.
	Args []RemoteObject `json:"args"`

	// Identifier of the context where the call was made.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`

	// Call timestamp.
	Timestamp Timestamp `json:"timestamp"`

	// Optional. Stack trace captured when the call was made.
	StackTrace StackTrace `json:"stackTrace,omitempty"`

	// Optional. Console context descriptor for calls on non-default console context (not
	// console.*): 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id'
	// for call on named context. EXPERIMENTAL
	Context string `json:"context,omitempty"`
}

/*
InspectRequestedEvent represents Runtime.inspectRequested event data.
*/
type InspectRequestedEvent struct {
	// Remote object.
	Object RemoteObject `json:"object"`

	// Hints.
	Hints map[string]string `json:"hints"`
}

/*
ScriptID is a unique script identifier.
*/
type ScriptID string

/*
RemoteObjectID is a unique object identifier.
*/
type RemoteObjectID string

/*
UnserializableValue is a primitive value which cannot be JSON-stringified.
*/
type UnserializableValue int

const (
	_infinity UnserializableValue = iota
	_nan
	_negInfinity
	_negNil
)

func (s UnserializableValue) String() string {
	str := string(s)
	if str == "Infinity" ||
		str == "NaN" ||
		str == "-Infinity" ||
		str == "-0" {
		return str
	}
	panic(fmt.Errorf("Invalid UnserializableValue '%s'", str))
}

/*
RemoteObject is a mirror object referencing the original JavaScript object.

*/
type RemoteObject struct {
	// Object type. Allowed values: object, function, undefined, string, number, boolean, symbol.
	Type string `json:"type"`

	// Optional. Object subtype hint. Specified for object type values only. Allowed values: array,
	// null, node, regexp, date, map, set, weakmap, weakset, iterator, generator, error, proxy,
	// promise, typedarray.
	Subtype string `json:"subtype,omitempty"`

	// Optional. Object class (constructor) name. Specified for object type values only.
	ClassName string `json:"className,omitempty"`

	// Optional. Remote object value in case of primitive values or JSON values (if it was
	// requested).
	Value interface{} `json:"value,omitempty"`

	// Optional. Primitive value which can not be JSON-stringified does not have value, but gets
	// this property.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"`

	// Optional. String representation of the object.
	Description string `json:"description,omitempty"`

	// Optional. Unique object identifier (for non-primitive values).
	ObjectID RemoteObjectID `json:"objectId,omitempty"`

	// Optional. Preview containing abbreviated property values. Specified for object type values
	// only. EXPERIMENTAL
	Preview ObjectPreview `json:"preview,omitempty"`

	// Optional. EXPERIMENTAL
	CustomPreview CustomPreview `json:"customPreview,omitempty"`
}

/*
CustomPreview is EXPERIMENTAL
*/
type CustomPreview struct {
	Header                     string         `json:"header"`
	HasBody                    bool           `json:"hasBody"`
	FormatterObjectID          RemoteObjectID `json:"formatterObjectId"`
	BindRemoteObjectFunctionID RemoteObjectID `json:"bindRemoteObjectFunctionId"`
	// Optional.
	ConfigObjectID RemoteObjectID `json:"configObjectId,omitempty,omitempty"`
}

/*
ObjectPreview is an object containing abbreviated remote object value. EXPERIMENTAL
*/
type ObjectPreview struct {
	// Object type. Allowed values: object, function, undefined, string, number, boolean, symbol.
	Type string `json:"type"`

	// Object subtype hint. Specified for object type values only. Allowed values: array, null,
	// node, regexp, date, map, set, weakmap, weakset, iterator, generator, error.
	Subtype string `json:"subtype"`

	// Optional. String representation of the object.
	Description string `json:"description,omitempty"`

	// True iff some of the properties or entries of the original object did not fit.
	Overflow bool `json:"overflow"`

	// List of the properties.
	Properties []*PropertyPreview `json:"properties"`

	// Optional. List of the entries. Specified for map and set subtype values only.
	Entries []*EntryPreview `json:"entries,omitempty"`
}

/*
PropertyPreview is EXPERIMENTAL
*/
type PropertyPreview struct {
	// Property name.
	Name string `json:"name"`

	// Object type. Accessor means that the property itself is an accessor property. Allowed values:
	// object, function, undefined, string, number, boolean, symbol, accessor.
	Type string `json:"type"`

	// Optional. User-friendly property value string.
	Value string `json:"value,omitempty"`

	// Optional. Nested value preview.
	ValuePreview ObjectPreview `json:"valuePreview,omitempty"`

	// Optional. Object subtype hint. Specified for object type values only. Allowed values: array,
	// null, node, regexp, date, map, set, weakmap, weakset, iterator, generator, error.
	Subtype string `json:"subtype,omitempty"`
}

/*
EntryPreview is EXPERIMENTAL
*/
type EntryPreview struct {
	// Optional. Preview of the key. Specified for map-like collection entries.
	Key ObjectPreview `json:"key,omitempty"`

	// Preview of the value.
	Value ObjectPreview `json:"value"`
}

/*
PropertyDescriptor is an object property descriptor.
*/
type PropertyDescriptor struct {
	// Property name or symbol description.
	Name string `json:"name"`

	// Optional. The value associated with the property.
	Value RemoteObject `json:"value,omitempty"`

	// Optional. True if the value associated with the property may be changed (data descriptors
	// only).
	Writable bool `json:"writable,omitempty"`

	// Optional. A function which serves as a getter for the property, or undefined if there is no
	// getter (accessor descriptors only).
	Get RemoteObject `json:"get,omitempty"`

	// Optional. A function which serves as a setter for the property, or undefined if there is no
	// setter (accessor descriptors only).
	Set RemoteObject `json:"set,omitempty"`

	// True if the type of this property descriptor may be changed and if the property may be
	// deleted from the corresponding object.
	Configurable bool `json:"configurable"`

	// True if this property shows up during enumeration of the properties on the corresponding
	// object.
	Enumerable bool `json:"enumerable"`

	// Optional. True if the result was thrown during the evaluation.
	WasThrown bool `json:"wasThrown,omitempty"`

	// Optional. True if the property is owned for the object.
	IsOwn bool `json:"isOwn,omitempty"`

	// Optional. Property symbol object, if the property is of the symbol type.
	Symbol RemoteObject `json:"symbol,omitempty"`
}

/*
InternalPropertyDescriptor is an object's internal property descriptor. This property isn't normally
visible in JavaScript code.
*/
type InternalPropertyDescriptor struct {
	// Conventional property name.
	Name string `json:"name"`

	// Optional. The value associated with the property.
	Value RemoteObject `json:"value,omitempty"`
}

/*
CallArgument represents a function call argument. Either remote object id objectId, primitive value,
unserializable primitive value or neither of (for undefined) them should be specified.
*/
type CallArgument struct {
	// Optional. Primitive value or serializable javascript object.
	Value interface{} `json:"value,omitempty"`

	// Optional. Primitive value which can not be JSON-stringified.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"`

	// Optional. Remote object handle.
	ObjectID RemoteObjectID `json:"objectId,omitempty"`
}

/*
ExecutionContextID is the ID of an execution context.
*/
type ExecutionContextID int

/*
ExecutionContextDescription is the description of an isolated world.
*/
type ExecutionContextDescription struct {
	// Unique ID of the execution context. It can be used to specify in which execution context
	// script evaluation should be performed.
	ID ExecutionContextID `json:"id"`

	// Execution context origin.
	Origin string `json:"origin"`

	// Human readable name describing given context.
	Name string `json:"name"`

	// Optional. Embedder-specific auxiliary data.
	AuxData map[string]string `json:"auxData,omitempty"`
}

/*
ExceptionDetails contains detailed information about exception (or error) that was thrown during
script compilation or execution.
*/
type ExceptionDetails struct {
	// Exception id.
	ExceptionID int `json:"exceptionId"`

	// Exception text, which should be used together with exception object when available.
	Text string `json:"text"`

	// Line number of the exception location (0-based).
	LineNumber int `json:"lineNumber"`

	// Column number of the exception location (0-based).
	ColumnNumber int `json:"columnNumber"`

	// Optional. Script ID of the exception location.
	ScriptID ScriptID `json:"scriptId,omitempty"`

	// Optional. URL of the exception location, to be used when the script was not reported.
	URL string `json:"url,omitempty"`

	// Optional. JavaScript stack trace if available.
	StackTrace StackTrace `json:"stackTrace,omitempty"`

	// Optional. Exception object if available.
	Exception RemoteObject `json:"exception,omitempty"`

	// Optional. Identifier of the context where exception happened.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
Timestamp is the number of milliseconds since epoch.
*/
type Timestamp int

/*
CallFrame is a stack entry for runtime errors and assertions.
*/
type CallFrame struct {
	// JavaScript function name.
	FunctionName string `json:"functionName"`

	// JavaScript script id.
	ScriptID ScriptID `json:"scriptId"`

	// JavaScript script name or url.
	URL string `json:"url"`

	// JavaScript script line number (0-based).
	LineNumber int `json:"lineNumber"`

	// JavaScript script column number (0-based).
	ColumnNumber int `json:"columnNumber"`
}

/*
StackTrace contains call frames for assertions or error messages.
*/
type StackTrace struct {
	// Optional. String label of this stack trace. For async traces this may be a name of the
	// function that initiated the async call.
	Description string `json:"description,omitempty"`

	// JavaScript function name.
	CallFrames []*CallFrame `json:"callFrames"`

	// Optional. Asynchronous JavaScript stack trace that preceded this stack, if available.
	Parent *StackTrace `json:"parent,omitempty"`

	// Optional. Asynchronous JavaScript stack trace that preceded this stack, if available.
	// EXPERIMENTAL
	ParentID StackTraceID `json:"parentId,omitempty"`
}

/*
UniqueDebuggerID is the unique identifier of the  current debugger. EXPERIMENTAL
*/
type UniqueDebuggerID string

/*
StackTraceID - If debuggerID is set stack trace comes from another debugger and can be resolved
there. This allows to track cross-debugger calls. See StackTrace and Debugger.paused for usages.
EXPERIMENTAL
*/
type StackTraceID struct {
	ID string `json:"id"`
	// Optional.
	DebuggerID UniqueDebuggerID `json:"debuggerId,omitempty"`
}
