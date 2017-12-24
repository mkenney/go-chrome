/*
Package HeadlessExperimental provides type definitions for use with the Chrome HeadlessExperimental
protocol

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/
*/
package HeadlessExperimental

import (
	"github.com/mkenney/go-chrome/runtime"
)

/*
BeginFrameParams represents HeadlessExperimental.beginFrame parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
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
	//
	// This is an instance of ScreenshotParams, but that doesn't omitempty correctly so it must be
	// added manually.
	Screenshot interface{} `json:"screenshot,omitempty"`
}

/*
BeginFrameResult represents the result of calls to HeadlessExperimental.beginFrame.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
*/
type BeginFrameResult struct {
	// Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the
	// display.
	HasDamage bool `json:"hasDamage"`

	// Whether the main frame submitted a new display frame in response to this BeginFrame.
	MainFrameContentUpdated bool `json:"mainFrameContentUpdated"`

	// Base64-encoded image data of the screenshot, if one was requested and successfully taken.
	ScreenshotData string `json:"screenshotData"`
}

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

/*
ScreenshotParams represents encoding options for a screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#type-ScreenshotParams
*/
type ScreenshotParams struct {
	// Optional. Image compression format (defaults to png). Allowed values: jpeg, png.
	Format string `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`
}
