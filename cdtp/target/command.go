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
