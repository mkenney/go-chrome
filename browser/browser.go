package Browser

import (
	Target "app/chrome/target"
	"fmt"
)

/*
GetWindowForTargetParams represents Browser.getWindowForTarget parameters
*/
type GetWindowForTargetParams struct {
	// Devtools agent host id.
	TargetID Target.TargetID `json:"targetId"`
}

/*
SetWindowBoundsParams represents Browser.setWindowBounds parameters
*/
type SetWindowBoundsParams struct {
	// Browser window id.
	WindowID WindowID `json:"windowId"`

	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined
	// with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds Bounds `json:"bounds"`
}

/*
GetWindowBoundsParams represents Browser.getWindowBounds parameters
*/
type GetWindowBoundsParams struct {
	// Browser window id.
	WindowID WindowID `json:"windowId"`
}

/*
WindowID - EXPERIMENTAL
*/
type WindowID int

/*
WindowState holds the state of the browser window. EXPERIMENTAL
*/
type WindowState int

const (
	_normal WindowState = iota
	_minimized
	_maximized
	_fullscreen
)

func (a WindowState) String() string {
	if a == 0 {
		return "normal"
	}
	if a == 1 {
		return "minimized"
	}
	if a == 2 {
		return "maximized"
	}
	if a == 3 {
		return "fullscreen"
	}
	panic(fmt.Errorf("Invalid WindowState %d", a))
}

/*
Bounds holds the browser window bounds information. EXPERIMENTAL
*/
type Bounds struct {
	// Optional. The offset from the left edge of the screen to the window in pixels.
	Left int `json:"left,omitempty"`

	// Optional. The offset from the top edge of the screen to the window in pixels.
	Top int `json:"top,omitempty"`

	// Optional. The window width in pixels.
	Width int `json:"width,omitempty"`

	// Optional. The window height in pixels.
	Height int `json:"height,omitempty"`

	// Optional. The window state. Default to normal.
	WindowState WindowState `json:"windowState,omitempty"`
}
