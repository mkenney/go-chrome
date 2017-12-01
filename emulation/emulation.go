package Emulation

import (
	Page "app/chrome/page"
	"fmt"
)

/*
SetDeviceMetricsOverrideParams represents Emulation.setDeviceMetricsOverride parameters.
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

	// Scale to apply to resulting view image. EXPERIMENTAL
	Scale int `json:"scale"`

	// Overriding screen width value in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	ScreenWidth int `json:"screenWidth"`

	// Overriding screen height value in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	ScreenHeight int `json:"screenHeight"`

	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	PositionX int `json:"positionX"`

	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	PositionY int `json:"positionY"`

	// Do not set visible view size, rely upon explicit setVisibleSize call. EXPERIMENTAL
	DontSetVisibleSize bool `json:"dontSetVisibleSize"`

	// Screen orientation override.
	ScreenOrientation ScreenOrientation `json:"screenOrientation"`

	// If set, the visible area of the page will be overridden to this viewport. This viewport
	// change is not observed by the page, e.g. viewport-relative elements do not change positions.
	// EXPERIMENTAL
	Viewport Page.Viewport `json:"viewport"`
}

/*
SetDeviceMetricsOverrideParams represents Emulation.setDeviceMetricsOverride parameters.
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

	// Scale to apply to resulting view image. EXPERIMENTAL
	Scale int `json:"scale"`

	// Overriding screen width value in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	ScreenWidth int `json:"screenWidth"`

	// Overriding screen height value in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	ScreenHeight int `json:"screenHeight"`

	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	PositionX int `json:"positionX"`

	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). EXPERIMENTAL
	PositionY int `json:"positionY"`

	// Do not set visible view size, rely upon explicit setVisibleSize call. EXPERIMENTAL
	DontSetVisibleSize bool `json:"dontSetVisibleSize"`

	// Screen orientation override.
	ScreenOrientation ScreenOrientation `json:"screenOrientation"`

	// If set, the visible area of the page will be overridden to this viewport. This viewport
	// change is not observed by the page, e.g. viewport-relative elements do not change positions.
	// EXPERIMENTAL
	Viewport Page.Viewport `json:"viewport"`
}

/*
SetPageScaleFactorParams represents Emulation.setPageScaleFactor parameters.
*/
type SetPageScaleFactorParams struct {
	// Page scale factor.
	PageScaleFactor int `json:"pageScaleFactor"`
}

/*
SetVisibleSizeParams represents Emulation.setVisibleSize parameters.
*/
type SetVisibleSizeParams struct {
	// Frame width (DIP).
	Width int `json:"width"`

	// Frame height (DIP).
	Height int `json:"height"`
}

/*
SetScriptExecutionDisabledParams represents Emulation.setScriptExecutionDisabled parameters.
*/
type SetScriptExecutionDisabledParams struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

/*
SetGeolocationOverrideParams represents Emulation.setGeolocationOverride parameters.
*/
type SetGeolocationOverrideParams struct {
	// Mock latitude.
	Latitude int `json:"latitude"`

	// Mock longitude.
	Longitude int `json:"longitude"`

	// Mock accuracy.
	Accuracy int `json:"accuracy"`
}

/*
SetTouchEmulationEnabledParams represents Emulation.setTouchEmulationEnabled parameters.
*/
type SetTouchEmulationEnabledParams struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`

	// Maximum touch points supported. Defaults to one.
	MaxTouchPoints int `json:"maxTouchPoints"`
}

/*
SetEmitTouchEventsForMouseParams represents Emulation.setEmitTouchEventsForMouse parameters.
*/
type SetEmitTouchEventsForMouseParams struct {
	// Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`

	// Touch/gesture events configuration. Default: current platform. Allowed values: mobile,
	// desktop.
	Configuration string `json:"configuration"`
}

/*
SetEmulatedMediaParams represents Emulation.setEmulatedMedia parameters.
*/
type SetEmulatedMediaParams struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media"`
}

/*
SetCPUThrottlingRateParams represents Emulation.setCPUThrottlingRate parameters.
*/
type SetCPUThrottlingRateParams struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	rate int `json:"rate"`
}

/*
SetVirtualTimePolicyParams represents Emulation.setVirtualTimePolicy parameters.
*/
type SetVirtualTimePolicyParams struct {
	// desc.
	Policy VirtualTimePolicy `json:"policy"`

	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a
	// virtualTimeBudgetExpired event is sent.
	Budget int `json:"budget"`

	// If set this specifies the maximum number of tasks that can be run before virtual is forced
	// forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount int `json:"maxVirtualTimeTaskStarvationCount"`
}

/*
SetNavigatorOverridesParams represents Emulation.setNavigatorOverrides parameters.
*/
type SetNavigatorOverridesParams struct {
	// The platform navigator.platform should return.
	Platform string `json:"platform"`
}

/*
SetDefaultBackgroundColorOverrideParams represents Emulation.setDefaultBackgroundColorOverride
parameters.
*/
type SetDefaultBackgroundColorOverrideParams struct {
	// RGBA of the default background color. If not specified, any existing override will be
	// cleared.
	Color DOM.RGBA `json:"color"`
}

/*
VirtualTimeBudgetExpiredEvent represents Emulation.virtualTimeBudgetExpired event data.
*/
type VirtualTimeBudgetExpiredEvent struct{}

/*
VirtualTimeAdvancedEvent represents Emulation.virtualTimeAdvanced event data.
*/
type VirtualTimeAdvancedEvent struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

/*
VirtualTimePausedEvent represents Emulation.virtualTimePaused event data.
*/
type VirtualTimePausedEvent struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

/*
ScreenOrientation represents a screen orientation.
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
*/
type VirtualTimePolicy int

const (
	_advance VirtualTimePolicy = iota
	_pause
	_pauseIfNetworkFetchesPending
)

func (a VirtualTimePolicy) String() string {
	if a == 0 {
		return "advance"
	}
	if a == 1 {
		return "pause"
	}
	if a == 2 {
		return "pauseIfNetworkFetchesPending"
	}
	panic(fmt.Errorf("Invalid VirtualTimePolicy %d", a))
}
