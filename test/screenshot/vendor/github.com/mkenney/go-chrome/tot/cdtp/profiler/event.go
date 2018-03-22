package profiler

import (
	"github.com/mkenney/go-chrome/tot/cdtp/debugger"
)

/*
ConsoleProfileFinishedEvent represents Overlay.consoleProfileFinished event data.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileFinished
*/
type ConsoleProfileFinishedEvent struct {
	// Profile ID.
	ID string `json:"id"`

	// Location of console.profileEnd().
	Location *debugger.Location `json:"location"`

	// Profile data.
	Profile *Profile `json:"profile"`

	// Profile title passed as an argument to console.profile().
	Title string `json:"title"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ConsoleProfileStartedEvent represents Overlay.consoleProfileStarted event data.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileStarted
*/
type ConsoleProfileStartedEvent struct {
	// Profile ID.
	ID string `json:"id"`

	// Location of console.profile().
	Location *debugger.Location `json:"location"`

	// Profile title passed as an argument to console.profile().
	Title string `json:"title"`

	// Error information related to this event
	Err error `json:"-"`
}
