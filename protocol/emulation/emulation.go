/*
Package emulation provides type definitions for use with the Chrome Emulation protocol

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
*/
package emulation

import (
	"github.com/mkenney/go-chrome/protocol/runtime"
)

/*
CanEmulateResult represents the result of calls to Emulation.canEmulate.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
*/
type CanEmulateResult struct {
	// True if emulation is supported.
	Result bool `json:"result"`
}

/*
SetCPUThrottlingRateParams represents Emulation.setCPUThrottlingRate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
*/
type SetCPUThrottlingRateParams struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate int `json:"rate"`
}

/*
SetDefaultBackgroundColorOverrideParams represents Emulation.setDefaultBackgroundColorOverride
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
*/
type SetDefaultBackgroundColorOverrideParams struct {
	// Optional. RGBA of the default background color. If not specified, any existing override will be
	// cleared.
	//
	// This expects an instance of DOM.RGBA, but that doesn't omitempty correctly so it must be
	// added manually.
	Color interface{} `json:"color,omitempty"`
}

/*
SetDeviceMetricsOverrideParams represents Emulation.setDeviceMetricsOverride parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
*/
type SetDeviceMetricsOverrideParams struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`

	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`

	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor int `json:"deviceScaleFactor"`

	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text
	// autosizing and more.
	Mobile bool `json:"mobile"`

	// Optional. Scale to apply to resulting view image. EXPERIMENTAL
	Scale int `json:"scale,omitempty"`

	// Optional. Overriding screen width value in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	ScreenWidth int `json:"screenWidth,omitempty"`

	// Optional. Overriding screen height value in pixels (minimum 0, maximum 10000000).
	// EXPERIMENTAL
	ScreenHeight int `json:"screenHeight,omitempty"`

	// Optional. Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	// EXPERIMENTAL
	PositionX int `json:"positionX,omitempty"`

	// Optional. Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	// EXPERIMENTAL
	PositionY int `json:"positionY,omitempty"`

	// Optional. Do not set visible view size, rely upon explicit setVisibleSize call. EXPERIMENTAL
	DontSetVisibleSize bool `json:"dontSetVisibleSize,omitempty"`

	// Optional. Screen orientation override.
	//
	// This is an instance of ScreenOrientation, but that doesn't omitempty correctly so it must be
	// added manually.
	ScreenOrientation interface{} `json:"screenOrientation,omitempty"`

	// Optional. If set, the visible area of the page will be overridden to this viewport. This
	// viewport change is not observed by the page, e.g. viewport-relative elements do not change
	// positions. EXPERIMENTAL
	//
	// This is an instance of Page.Viewport, but that doesn't omitempty correctly so it must be
	// added manually.
	Viewport interface{} `json:"viewport,omitempty"`
}

/*
SetEmitTouchEventsForMouseParams represents Emulation.setEmitTouchEventsForMouse parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
*/
type SetEmitTouchEventsForMouseParams struct {
	// Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`

	// Optional. Touch/gesture events configuration. Default: current platform. Allowed values:
	// mobile, desktop.
	Configuration string `json:"configuration,omitempty"`
}

/*
SetEmulatedMediaParams represents Emulation.setEmulatedMedia parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
*/
type SetEmulatedMediaParams struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media"`
}

/*
SetGeolocationOverrideParams represents Emulation.setGeolocationOverride parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
*/
type SetGeolocationOverrideParams struct {
	// Optional. Mock latitude.
	Latitude int `json:"latitude,omitempty"`

	// Optional. Mock longitude.
	Longitude int `json:"longitude,omitempty"`

	// Optional. Mock accuracy.
	Accuracy int `json:"accuracy,omitempty"`
}

/*
SetNavigatorOverridesParams represents Emulation.setNavigatorOverrides parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
*/
type SetNavigatorOverridesParams struct {
	// The platform navigator.platform should return.
	Platform string `json:"platform"`
}

/*
SetPageScaleFactorParams represents Emulation.setPageScaleFactor parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
*/
type SetPageScaleFactorParams struct {
	// Page scale factor.
	PageScaleFactor int `json:"pageScaleFactor"`
}

/*
SetScriptExecutionDisabledParams represents Emulation.setScriptExecutionDisabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
*/
type SetScriptExecutionDisabledParams struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

/*
SetTouchEmulationEnabledParams represents Emulation.setTouchEmulationEnabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
*/
type SetTouchEmulationEnabledParams struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`

	// Optional. Maximum touch points supported. Defaults to one.
	MaxTouchPoints int `json:"maxTouchPoints,omitempty"`
}

/*
SetVirtualTimePolicyParams represents Emulation.setVirtualTimePolicy parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
*/
type SetVirtualTimePolicyParams struct {
	// desc.
	Policy VirtualTimePolicy `json:"policy"`

	// Optional. If set, after this many virtual milliseconds have elapsed virtual time will be
	// paused and a virtualTimeBudgetExpired event is sent.
	Budget int `json:"budget,omitempty"`

	// Optional. If set this specifies the maximum number of tasks that can be run before virtual is
	// forced forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount int `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
}

/*
SetVirtualTimePolicyResult represents the result of calls to Emulation.setVirtualTimePolicy.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
*/
type SetVirtualTimePolicyResult struct {
	// Absolute timestamp at which virtual time was first enabled (milliseconds since epoch).
	VirtualTimeBase runtime.Timestamp `json:"virtualTimeBase"`
}

/*
SetVisibleSizeParams represents Emulation.setVisibleSize parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
*/
type SetVisibleSizeParams struct {
	// Frame width (DIP).
	Width int `json:"width"`

	// Frame height (DIP).
	Height int `json:"height"`
}

/*
VirtualTimeAdvancedEvent represents Emulation.virtualTimeAdvanced event data.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeAdvanced
*/
type VirtualTimeAdvancedEvent struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

/*
VirtualTimeBudgetExpiredEvent represents Emulation.virtualTimeBudgetExpired event data.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeBudgetExpired
*/
type VirtualTimeBudgetExpiredEvent struct{}

/*
VirtualTimePausedEvent represents Emulation.virtualTimePaused event data.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimePaused
*/
type VirtualTimePausedEvent struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

/*
ScreenOrientation represents a screen orientation.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-ScreenOrientation
*/
type ScreenOrientation struct {
	// Orientation type. Allowed values: portraitPrimary, portraitSecondary, landscapePrimary,
	// landscapeSecondary.
	Type string `json:"type"`

	// Orientation angle.
	Angle int `json:"angle"`
}

/*
VirtualTimePolicy is the time policy
	- advance: If the scheduler runs out of immediate work, the virtual time base may fast forward
	  to allow the next delayed task (if any) to run
	- pause: The virtual time base may not advance
	- pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending
	  resource fetches.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-VirtualTimePolicy
*/
type VirtualTimePolicy string
