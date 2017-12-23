package Browser

import (
	"fmt"

	"github.com/mkenney/go-chrome/target"
)

/*
GetVersionResult represents the result of calls to Browser.getVersion.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion
*/
type GetVersionResult struct {
	// Protocol version.
	ProtocolVersion string `json:"protocolVersion"`

	// Product name.
	Product string `json:"product"`

	// Product revision.
	Revision string `json:"revision"`

	// User-Agent.
	UserAgent string `json:"userAgent"`

	// V8 version.
	JSVersion string `json:"jsVersion"`
}

/*
GetWindowBoundsParams represents Browser.getWindowBounds parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds
*/
type GetWindowBoundsParams struct {
	// Browser window id.
	WindowID WindowID `json:"windowId"`
}

/*
GetWindowBoundsResult represents the result of calls to Browser.getWindowBounds.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds
*/
type GetWindowBoundsResult struct {
	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds Bounds `json:"bounds"`
}

/*
GetWindowForTargetParams represents Browser.getWindowForTarget parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
*/
type GetWindowForTargetParams struct {
	// Devtools agent host id.
	TargetID Target.ID `json:"targetId"`
}

/*
GetWindowForTargetResult represents the result of calls to Browser.getWindowForTarget.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
*/
type GetWindowForTargetResult struct {
	// Browser window ID.
	WindowID WindowID `json:"windowId"`

	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds Bounds `json:"bounds"`
}

/*
SetWindowBoundsParams represents Browser.setWindowBounds parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds
*/
type SetWindowBoundsParams struct {
	// Browser window id.
	WindowID WindowID `json:"windowId"`

	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined
	// with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds Bounds `json:"bounds"`
}

/*
SetWindowBoundsResult represents the result of calls to Browser.setWindowBounds.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds
*/
type SetWindowBoundsResult struct {
	// Browser window ID.
	WindowID WindowID `json:"windowId"`

	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined
	// with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds Bounds `json:"bounds"`
}

/*
WindowID - EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-WindowID
*/
type WindowID int

/*
WindowState holds the state of the browser window. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-WindowState
*/
type WindowState string

/*
Bounds holds the browser window bounds information. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-Bounds
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
	//
	// This expects an instance of WindowState, but that doesn't omitempty correctly so it must be
	// added manually.
	WindowState interface{} `json:"windowState,omitempty"`
}
