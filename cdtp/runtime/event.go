package runtime

/*
ConsoleAPICalledEvent represents Runtime.consoleAPICalled event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-consoleAPICalled
*/
type ConsoleAPICalledEvent struct {
	// Type of the call.
	//
	// Allowed values:
	//	- log
	//	- debug
	//	- info
	//	- error
	//	- warning
	//	- dir
	//	- dirxml
	//	- table
	//	- trace
	//	- clear
	//	- startGroup
	//	- startGroupCollapsed
	//	- endGroup
	//	- assert
	//	- profile
	//	- profileEnd
	//	- count
	//	- timeEnd
	Type string `json:"type"`

	// Call arguments.
	Args []*RemoteObject `json:"args"`

	// Identifier of the context where the call was made.
	ExecutionContextID *ExecutionContextID `json:"executionContextId"`

	// Call timestamp.
	Timestamp Timestamp `json:"timestamp"`

	// Optional. Stack trace captured when the call was made.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`

	// Optional. Console context descriptor for calls on non-default console
	// context (not console.*): 'anonymous#unique-logger-id' for call on unnamed
	// context, 'name#unique-logger-id' for call on named context. EXPERIMENTAL.
	Context string `json:"context,omitempty"`
}

/*
ExceptionRevokedEvent represents Runtime.exceptionRevoked event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionRevoked
*/
type ExceptionRevokedEvent struct {
	// Reason describing why exception was revoked.
	Reason string `json:"reason"`

	// The ID of revoked exception, as reported in exceptionThrown.
	ExceptionID int `json:"exceptionId"`
}

/*
ExceptionThrownEvent represents Runtime.exceptionThrown event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionThrown
*/
type ExceptionThrownEvent struct {
	// Timestamp of the exception.
	Timestamp Timestamp `json:"timestamp"`

	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`
}

/*
ExecutionContextCreatedEvent represents Runtime.executionContextCreated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextCreated
*/
type ExecutionContextCreatedEvent struct {
	// A newly created execution context.
	Context *ExecutionContextDescription `json:"context"`
}

/*
ExecutionContextDestroyedEvent represents Runtime.executionContextDestroyed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextDestroyed
*/
type ExecutionContextDestroyedEvent struct {
	// ID of the destroyed context.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`
}

/*
ExecutionContextsClearedEvent represents Runtime.executionContextsCleared event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextsCleared
*/
type ExecutionContextsClearedEvent struct{}

/*
InspectRequestedEvent represents Runtime.inspectRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-inspectRequested
*/
type InspectRequestedEvent struct {
	// Remote object.
	Object *RemoteObject `json:"object"`

	// Hints.
	Hints map[string]string `json:"hints"`
}
