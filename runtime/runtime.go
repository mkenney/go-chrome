package Runtime

import (
	"fmt"
)

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
	// only. EXPERIMENTAL.
	Preview ObjectPreview `json:"preview,omitempty"`

	// Optional. EXPERIMENTAL.
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
	// EXPERIMENTAL.
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
