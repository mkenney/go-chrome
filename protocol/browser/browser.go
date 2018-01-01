/*
Package browser provides type definitions for use with the Chrome Browser protocol

https://chromedevtools.github.io/devtools-protocol/tot/Browser/
*/
package browser

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
	// Optional. The offset from the left edge of the screen to the window in
	// pixels.
	Left int `json:"left,omitempty"`

	// Optional. The offset from the top edge of the screen to the window in
	// pixels.
	Top int `json:"top,omitempty"`

	// Optional. The window width in pixels.
	Width int `json:"width,omitempty"`

	// Optional. The window height in pixels.
	Height int `json:"height,omitempty"`

	// Optional. The window state. Default to normal.
	WindowState WindowState `json:"windowState,omitempty"`
}
