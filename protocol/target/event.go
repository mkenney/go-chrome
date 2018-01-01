package target

/*
AttachedToTargetEvent represents Target.attachedToTarget event data.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-attachedToTarget
*/
type AttachedToTargetEvent struct {
	// Identifier assigned to the session used to send/receive messages.
	SessionID SessionID `json:"sessionId"`

	// Target info.
	Info Info `json:"targetInfo"`

	// Waiting for debugger.
	WaitingForDebugger bool `json:"waitingForDebugger"`
}

/*
DetachedFromTargetEvent represents Target.detachedFromTarget event data.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-detachedFromTarget
*/
type DetachedFromTargetEvent struct {
	// Detached session identifier.
	SessionID SessionID `json:"sessionId"`

	// DEPRECATED
	ID ID `json:"targetId"`
}

/*
ReceivedMessageFromTargetEvent represents Target.receivedMessageFromTarget event data.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-receivedMessageFromTarget
*/
type ReceivedMessageFromTargetEvent struct {
	// Identifier of a session which sends a message.
	SessionID SessionID `json:"sessionId"`

	// Message.
	Message string `json:"message"`

	// Optional. Deprecated. DEPRECATED.
	ID ID `json:"targetId,omitempty"`
}

/*
CreatedEvent represents Target.targetCreated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCreated
*/
type CreatedEvent struct {
	// Target info.
	Info *Info `json:"targetInfo"`
}

/*
DestroyedEvent represents Target.targetDestroyed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetDestroyed
*/
type DestroyedEvent struct {
	// Target ID.
	ID ID `json:"targetId"`
}

/*
InfoChangedEvent represents Target.targetInfoChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetInfoChanged
*/
type InfoChangedEvent struct {
	// Target info.
	Info *Info `json:"targetInfo"`
}
