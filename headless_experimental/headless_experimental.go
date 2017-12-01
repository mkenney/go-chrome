package HeadlessExperimental

/*
ScreenshotParams represents encoding options for a screenshot.
*/
type ScreenshotParams struct {
	// Optional. Image compression format (defaults to png). Allowed values: jpeg, png.
	Format string `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`
}

/*
BeginFrameParams represents HeadlessExperimental.beginFrame parameters.
*/
type BeginFrameParams struct {
	// Optional. Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current
	// time will be used.
	FrameTime Runtime.Timestamp `json:"frameTime,omitempty"`

	// Optional. Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline
	// will be calculated from the frameTime and interval.
	Deadline Runtime.Timestamp `json:"deadline,omitempty"`

	// Optional. The interval between BeginFrames that is reported to the compositor, in
	// milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval float64 `json:"interval,omitempty"`

	// Optional. If set, a screenshot of the frame will be captured and returned in the response.
	// Otherwise, no screenshot will be captured.
	Screenshot ScreenshotParams `json:"screenshot,omitempty"`
}

/*
NeedsBeginFramesChangedEvent represents HeadlessExperimental.needsBeginFramesChanged event data.
*/
type NeedsBeginFramesChangedEvent struct {
	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames bool `json:"needsBeginFrames"`
}

/*
MainFrameReadyForScreenshotsEvent represents HeadlessExperimental.mainFrameReadyForScreenshots event
data.
*/
type MainFrameReadyForScreenshotsEvent struct{}
