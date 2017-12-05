package Target

/*
ActivateTargetParams represents Target.activateTarget parameters.
*/
type ActivateTargetParams struct {
	// Target ID.
	TargetID TargetID `json:"targetId"`
}

/*
AttachToTargetParams represents Target.attachToTarget parameters.
*/
type AttachToTargetParams struct {
	// Target ID.
	TargetID TargetID `json:"targetId"`
}

/*
CloseTargetParams represents Target.closeTarget parameters.
*/
type CloseTargetParams struct {
	// Target ID.
	TargetID TargetID `json:"targetId"`
}

/*
CreateBrowserContextParams represents Target.createBrowserContext parameters.
*/
type CreateBrowserContextParams struct {
	// The ID of the context created.
	BrowserContextID BrowserContextID `json:"browserContextId"`
}

/*
CreateTargetParams represents Target.createTarget parameters.
*/
type CreateTargetParams struct {
	// The initial URL the page will be navigated to.
	URL string `json:"url"`

	// Optional. Frame width in DIP (headless chrome only).
	Width int `json:"width,omitempty"`

	// Optional. Frame height in DIP (headless chrome only).
	Height int `json:"height,omitempty"`

	// Optional. The browser context to create the page in (headless chrome only).
	BrowserContextID BrowserContextID `json:"browserContextId,omitempty"`

	// Optional. Whether BeginFrames for this target will be controlled via DevTools (headless
	// chrome only, not supported on MacOS yet, false by default). EXPERIMENTAL
	EnableBeginFrameControl bool `json:"enableBeginFrameControl,omitempty"`
}

/*
DetachFromTargetParams represents Target.detachFromTarget parameters.
*/
type DetachFromTargetParams struct {
	// Optional. Session to detach.
	SessionID SessionID `json:"sessionId,omitempty"`

	// Optional. DEPRECATED
	TargetID TargetID `json:"targetId,omitempty"`
}

/*
DisposeBrowserContextParams represents Target.disposeBrowserContext parameters.
*/
type DisposeBrowserContextParams struct {
	// Target ID.
	TargetID TargetID `json:"targetId"`
}

/*
GetTargetInfoParams represents Target.getTargetInfo parameters.
*/
type GetTargetInfoParams struct {
	// Target ID.
	TargetID TargetID `json:"targetId"`
}

/*
GetTargetsParams represents Target.getTargets parameters.
*/
type GetTargetsParams struct {
	// The list of targets.
	TargetInfos []TargetInfo `json:"targetInfos"`
}

/*
SendMessageToTargetParams represents Target.sendMessageToTarget parameters.
*/
type SendMessageToTargetParams struct {
	// Message.
	Message string `json:"message"`

	// Optional. Identifier of the session.
	SessionID SessionID `json:"sessionId,omitempty"`

	// Optional. DEPRECATED
	TargetID TargetID `json:"targetId,omitempty"`
}

/*
SetAttachToFramesParams represents Target.setAttachToFrames parameters.
*/
type SetAttachToFramesParams struct {
	// Whether to attach to frames.
	Value bool `json:"value"`
}

/*
SetAutoAttachParams represents Target.setAutoAttach parameters.
*/
type SetAutoAttachParams struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`

	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger` to
	// run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

/*
SetDiscoverTargetsParams represents Target.setDiscoverTargets parameters.
*/
type SetDiscoverTargetsParams struct {
	// Whether to discover available targets.
	Discover bool `json:"discover"`
}

/*
SetRemoteLocationsParams represents Target.setRemoteLocations parameters.
*/
type SetRemoteLocationsParams struct {
	// List of remote locations.
	Locations []RemoteLocation `json:"locations"`
}

/*
AttachedToTargetEvent represents Target.attachedToTarget event data.
*/
type AttachedToTargetEvent struct {
	// Identifier assigned to the session used to send/receive messages.
	SessionID SessionID `json:"sessionId"`

	// Target info.
	TargetInfo TargetInfo `json:"targetInfo"`

	// Waiting for debugger.
	WaitingForDebugger bool `json:"waitingForDebugger"`
}

/*
DetachedFromTargetEvent represents Target.detachedFromTarget event data.
*/
type DetachedFromTargetEvent struct {
	// Detached session identifier.
	SessionID SessionID `json:"sessionId"`

	// DEPRECATED
	TargetID TargetID `json:"targetId"`
}

/*
ReceivedMessageFromTargetEvent represents Target.receivedMessageFromTarget event data.
*/
type ReceivedMessageFromTargetEvent struct {
	// Identifier of a session which sends a message.
	SessionID SessionID `json:"sessionId"`

	// Message.
	Message string `json:"message"`

	// Optional. Deprecated. DEPRECATED.
	TargetID TargetID `json:"targetId,omitempty"`
}

/*
TargetCreatedEvent represents Target.targetCreated event data.
*/
type TargetCreatedEvent struct {
	// Target info.
	TargetInfo TargetInfo `json:"targetInfo"`
}

/*
TargetDestroyedEvent represents Target.targetDestroyed event data.
*/
type TargetDestroyedEvent struct {
	// Target ID.
	TargetID TargetID `json:"targetId"`
}

/*
TargetInfoChangedEvent represents Target.targetInfoChanged event data.
*/
type TargetInfoChangedEvent struct {
	// Target info.
	TargetInfo TargetInfo `json:"targetInfo"`
}

///////////////////////

/*
TargetID is the target ID.
*/
type TargetID string

/*
SessionID is a unique identifier of attached debugging session.
*/
type SessionID string

/*
BrowserContextID is EXPERIMENTAL
*/
type BrowserContextID string

/*
TargetInfo holds the target info
*/
type TargetInfo struct {
	// desc.
	TargetID TargetID `json:"targetId"`

	// desc.
	Type string `json:"type"`

	// desc.
	Title string `json:"title"`

	// desc.
	URL string `json:"url"`

	// Whether the target has an attached client.
	Attached bool `json:"attached"`

	// Optional. Opener target Id.
	OpenerID TargetID `json:"openerId,omitempty"`
}

/*
RemoteLocation is EXPERIMENTAL
*/
type RemoteLocation struct {
	// desc.
	Host string `json:"host"`

	// desc.
	Port int `json:"port"`
}
