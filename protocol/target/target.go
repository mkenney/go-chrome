/*
Package target provides type definitions for use with the Chrome Target protocol

https://chromedevtools.github.io/devtools-protocol/tot/Target/
*/
package target

/*
ActivateTargetParams represents Target.activateTarget parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
*/
type ActivateTargetParams struct {
	// Target ID.
	ID ID `json:"targetId"`
}

/*
AttachToTargetParams represents Target.attachToTarget parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
*/
type AttachToTargetParams struct {
	// Target ID.
	ID ID `json:"targetId"`
}

/*
AttachToTargetResult represents the result of calls to Target.attachToTarget.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
*/
type AttachToTargetResult struct {
	// Id assigned to the session.
	SessionID SessionID `json:"sessionId"`
}

/*
CloseTargetParams represents Target.closeTarget parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
*/
type CloseTargetParams struct {
	// Target ID.
	ID ID `json:"targetId"`
}

/*
CloseTargetResult represents the result of calls to Target.closeTarget.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
*/
type CloseTargetResult struct {
	Success bool `json:"success"`
}

/*
CreateBrowserContextResult represents the result of calls to Target.createBrowserContext.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createBrowserContext
*/
type CreateBrowserContextResult struct {
	// The ID of the context created.
	BrowserContextID BrowserContextID `json:"browserContextId"`
}

/*
CreateTargetParams represents Target.createTarget parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
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

	// Optional. Whether BeginFrames for this target will be controlled via
	// DevTools (headless chrome only, not supported on MacOS yet, false by
	// default). EXPERIMENTAL.
	EnableBeginFrameControl bool `json:"enableBeginFrameControl,omitempty"`
}

/*
CreateTargetResult represents the result of calls to Target.createTarget.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
*/
type CreateTargetResult struct {
	// The ID of the page opened.
	ID ID `json:"targetId"`
}

/*
DetachFromTargetParams represents Target.detachFromTarget parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-detachFromTarget
*/
type DetachFromTargetParams struct {
	// Optional. Session to detach.
	SessionID SessionID `json:"sessionId,omitempty"`

	// Optional. DEPRECATED
	ID ID `json:"targetId,omitempty"`
}

/*
DisposeBrowserContextParams represents Target.disposeBrowserContext parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
*/
type DisposeBrowserContextParams struct {
	// Target ID.
	ID ID `json:"targetId"`
}

/*
DisposeBrowserContextResult represents the result of calls to Target.disposeBrowserContext.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
*/
type DisposeBrowserContextResult struct {
	Success bool `json:"success"`
}

/*
GetTargetInfoParams represents Target.getTargetInfo parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo
*/
type GetTargetInfoParams struct {
	// Target ID.
	ID ID `json:"targetId"`
}

/*
GetTargetInfoResult represents the result of calls to Target.getTargetInfo.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo
*/
type GetTargetInfoResult struct {
	// The list of targets.
	Infos []*Info `json:"targetInfos"`
}

/*
GetTargetsParams represents Target.getTargets parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
*/
type GetTargetsParams struct {
	// The list of targets.
	Infos []*Info `json:"targetInfos"`
}

/*
SendMessageToTargetParams represents Target.sendMessageToTarget parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-sendMessageToTarget
*/
type SendMessageToTargetParams struct {
	// Message.
	Message string `json:"message"`

	// Optional. Identifier of the session.
	SessionID SessionID `json:"sessionId,omitempty"`

	// Optional. DEPRECATED
	ID ID `json:"targetId,omitempty"`
}

/*
SetAttachToFramesParams represents Target.setAttachToFrames parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAttachToFrames
*/
type SetAttachToFramesParams struct {
	// Whether to attach to frames.
	Value bool `json:"value"`
}

/*
SetAutoAttachParams represents Target.setAutoAttach parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAutoAttach
*/
type SetAutoAttachParams struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`

	// Whether to pause new targets when attaching to them. Use
	// `Runtime.runIfWaitingForDebugger` to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

/*
SetDiscoverTargetsParams represents Target.setDiscoverTargets parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
*/
type SetDiscoverTargetsParams struct {
	// Whether to discover available targets.
	Discover bool `json:"discover"`
}

/*
SetRemoteLocationsParams represents Target.setRemoteLocations parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setRemoteLocations
*/
type SetRemoteLocationsParams struct {
	// List of remote locations.
	Locations []*RemoteLocation `json:"locations"`
}

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

/*
ID is the target ID.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetID
*/
type ID string

/*
SessionID is a unique identifier of attached debugging session.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-SessionID
*/
type SessionID string

/*
BrowserContextID is EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-BrowserContextID
*/
type BrowserContextID string

/*
Info holds the target info.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetInfo
*/
type Info struct {
	// desc.
	ID ID `json:"targetId"`

	// desc.
	Type string `json:"type"`

	// desc.
	Title string `json:"title"`

	// desc.
	URL string `json:"url"`

	// Whether the target has an attached client.
	Attached bool `json:"attached"`

	// Optional. Opener target Id.
	OpenerID ID `json:"openerId,omitempty"`
}

/*
RemoteLocation is EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-RemoteLocation
*/
type RemoteLocation struct {
	// desc.
	Host string `json:"host"`

	// desc.
	Port int `json:"port"`
}
