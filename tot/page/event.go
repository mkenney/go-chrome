package page

import (
	"github.com/mkenney/go-chrome/tot/runtime"
)

/*
DOMContentEventFiredEvent represents Page.domContentEventFired event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-domContentEventFired
*/
type DOMContentEventFiredEvent struct {
	Timestamp MonotonicTime `json:"timestamp"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameAttachedEvent represents Page.frameAttached event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameAttached
*/
type FrameAttachedEvent struct {
	// ID of the frame that has been attached.
	FrameID FrameID `json:"frameId"`

	// Parent frame identifier.
	ParentFrameID FrameID `json:"parentFrameId"`

	// Optional. JavaScript stack trace of when frame was attached, only set if
	// frame initiated from script.
	Stack *runtime.StackTrace `json:"stack,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameClearedScheduledNavigationEvent represents Page.frameClearedScheduledNavigation event
data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameClearedScheduledNavigation
*/
type FrameClearedScheduledNavigationEvent struct {
	// ID of the frame that has cleared its scheduled navigation.
	FrameID FrameID `json:"frameId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameDetachedEvent represents Page.frameDetached event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameDetached
*/
type FrameDetachedEvent struct {
	// ID of the frame that has been detached.
	FrameID FrameID `json:"frameId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameNavigatedEvent represents Page.frameNavigated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameNavigated
*/
type FrameNavigatedEvent struct {
	// Frame object.
	Frame *Frame `json:"frame"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameResizedEvent represents Page.frameResized event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameResized
*/
type FrameResizedEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameScheduledNavigationEvent represents Page.frameScheduledNavigation event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameScheduledNavigation
*/
type FrameScheduledNavigationEvent struct {
	// ID of the frame that has scheduled a navigation.
	FrameID FrameID `json:"frameId"`

	// Delay (in seconds) until the navigation is scheduled to begin. The
	// navigation is not guaranteed to start.
	Delay float64 `json:"delay"`

	// The reason for the navigation. Allowed values:
	//	- formSubmissionGet
	//	- formSubmissionPost
	//	- httpHeaderRefresh
	//	- scriptInitiated
	//	- metaTagRefresh
	//	- pageBlockInterstitial
	//	- reload
	Reason ReasonEnum `json:"reason"`

	// The destination URL for the scheduled navigation.
	URL string `json:"url"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameStartedLoadingEvent represents Page.frameStartedLoading event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStartedLoading
*/
type FrameStartedLoadingEvent struct {
	// ID of the frame that has started loading.
	FrameID FrameID `json:"frameId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
FrameStoppedLoadingEvent represents Page.frameStoppedLoading event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStoppedLoading
*/
type FrameStoppedLoadingEvent struct {
	// ID of the frame that has stopped loading.
	FrameID FrameID `json:"frameId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
InterstitialHiddenEvent represents Page.interstitialHidden event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialHidden
*/
type InterstitialHiddenEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
InterstitialShownEvent represents Page.interstitialShown event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialShown
*/
type InterstitialShownEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
JavascriptDialogClosedEvent represents Page.javascriptDialogClosed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogClosed
*/
type JavascriptDialogClosedEvent struct {
	// Whether dialog was confirmed.
	Result bool `json:"result"`

	// User input in case of prompt.
	UserInput string `json:"userInput"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
JavascriptDialogOpeningEvent represents Page.javascriptDialogOpening event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogOpening
*/
type JavascriptDialogOpeningEvent struct {
	// Frame url.
	URL string `json:"url"`

	// Message that will be displayed by the dialog.
	Message string `json:"message"`

	// Dialog type. Allowed values:
	//	- DialogType.Alert
	//	- DialogType.Confirm
	//	- DialogType.Prompt
	//	- DialogType.Beforeunload
	Type DialogTypeEnum `json:"type"`

	// Optional. Default dialog prompt.
	DefaultPrompt string `json:"defaultPrompt,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
LifecycleEventEvent represents Page.lifecycleEvent event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-lifecycleEvent
*/
type LifecycleEventEvent struct {
	// ID of the frame.
	FrameID FrameID `json:"frameId"`

	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID LoaderID `json:"loaderId"`

	// name.
	Name string `json:"name"`

	// timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
LoadEventFiredEvent represents Page.loadEventFired event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-loadEventFired
*/
type LoadEventFiredEvent struct {
	// timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ScreencastFrameEvent represents Page.screencastFrame event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastFrame
*/
type ScreencastFrameEvent struct {
	// Base64-encoded compressed image.
	Data string `json:"data"`

	// Screencast frame metadata.
	Metadata *ScreencastFrameMetadata `json:"metadata"`

	// Frame number.
	SessionID int `json:"sessionId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ScreencastVisibilityChangedEvent represents Page.screencastVisibilityChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastVisibilityChanged
*/
type ScreencastVisibilityChangedEvent struct {
	// True if the page is visible.
	Visible bool `json:"visible"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WindowOpenEvent represents Page.windowOpen event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-windowOpen
*/
type WindowOpenEvent struct {
	// The URL for the new window.
	URL string `json:"url"`

	// Window name.
	WindowName string `json:"windowName"`

	// An array of enabled window features.
	WindowFeatures []string `json:"windowFeatures"`

	// Whether or not it was triggered by user gesture.
	UserGesture bool `json:"userGesture"`

	// Error information related to this event
	Err error `json:"-"`
}
