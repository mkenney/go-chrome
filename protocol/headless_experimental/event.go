package headlessExperimental

/*
MainFrameReadyForScreenshotsEvent represents HeadlessExperimental.mainFrameReadyForScreenshots event
data.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-mainFrameReadyForScreenshots
*/
type MainFrameReadyForScreenshotsEvent struct{}

/*
NeedsBeginFramesChangedEvent represents HeadlessExperimental.needsBeginFramesChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-needsBeginFramesChanged
*/
type NeedsBeginFramesChangedEvent struct {
	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames bool `json:"needsBeginFrames"`
}
