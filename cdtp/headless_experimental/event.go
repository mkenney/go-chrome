package headlessExperimental

/*
MainFrameReadyForScreenshotsEvent represents HeadlessExperimental.mainFrameReadyForScreenshots event
data.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-mainFrameReadyForScreenshots
*/
type MainFrameReadyForScreenshotsEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
NeedsBeginFramesChangedEvent represents HeadlessExperimental.needsBeginFramesChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-needsBeginFramesChanged
*/
type NeedsBeginFramesChangedEvent struct {
	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames bool `json:"needsBeginFrames"`

	// Error information related to this event
	Err error `json:"-"`
}
