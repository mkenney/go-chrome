package emulation

import (
	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/runtime"
)

/*
CanEmulateResult represents the result of calls to Emulation.canEmulate.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
*/
type CanEmulateResult struct {
	// True if emulation is supported.
	Result bool `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ClearDeviceMetricsOverrideResult represents the result of calls to
Emulation.clearDeviceMetricsOverride.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearDeviceMetricsOverride
*/
type ClearDeviceMetricsOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ClearGeolocationOverrideResult represents the result of calls to Emulation.clearGeolocationOverride.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearGeolocationOverride
*/
type ClearGeolocationOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ResetPageScaleFactorResult represents the result of calls to Emulation.resetPageScaleFactor.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
*/
type ResetPageScaleFactorResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetCPUThrottlingRateParams represents Emulation.setCPUThrottlingRate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
*/
type SetCPUThrottlingRateParams struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate int64 `json:"rate"`
}

/*
SetCPUThrottlingRateResult represents the result of calls to Emulation.setCPUThrottlingRate.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
*/
type SetCPUThrottlingRateResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetDefaultBackgroundColorOverrideParams represents Emulation.setDefaultBackgroundColorOverride
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
*/
type SetDefaultBackgroundColorOverrideParams struct {
	// Optional. RGBA of the default background color. If not specified, any
	// existing override will be cleared.
	Color *dom.RGBA `json:"color,omitempty"`
}

/*
SetDefaultBackgroundColorOverrideResult represents the result of calls to
Emulation.setDefaultBackgroundColorOverride.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
*/
type SetDefaultBackgroundColorOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetDeviceMetricsOverrideParams represents Emulation.setDeviceMetricsOverride parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
*/
type SetDeviceMetricsOverrideParams struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0
	// disables the override.
	Width int `json:"width"`

	// Overriding height value in pixels (minimum 0, maximum 10000000). 0
	// disables the override.
	Height int `json:"height"`

	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"`

	// Whether to emulate mobile device. This includes viewport meta tag,
	// overlay scrollbars, text autosizing and more.
	Mobile bool `json:"mobile"`

	// Optional. Scale to apply to resulting view image. EXPERIMENTAL
	Scale int `json:"scale,omitempty"`

	// Optional. Overriding screen width value in pixels (minimum 0, maximum
	// 10000000). EXPERIMENTAL.
	ScreenWidth int `json:"screenWidth,omitempty"`

	// Optional. Overriding screen height value in pixels (minimum 0, maximum
	// 10000000). EXPERIMENTAL.
	ScreenHeight int `json:"screenHeight,omitempty"`

	// Optional. Overriding view X position on screen in pixels (minimum 0,
	// maximum 10000000). EXPERIMENTAL.
	PositionX int `json:"positionX,omitempty"`

	// Optional. Overriding view Y position on screen in pixels (minimum 0,
	// maximum 10000000). EXPERIMENTAL.
	PositionY int `json:"positionY,omitempty"`

	// Optional. Do not set visible view size, rely upon explicit setVisibleSize
	// call. EXPERIMENTAL.
	DontSetVisibleSize bool `json:"dontSetVisibleSize,omitempty"`

	// Optional. Screen orientation override.
	ScreenOrientation *ScreenOrientation `json:"screenOrientation,omitempty"`

	// Optional. If set, the visible area of the page will be overridden to this
	// viewport. This viewport change is not observed by the page, e.g.
	// viewport-relative elements do not change positions. EXPERIMENTAL.
	Viewport *page.Viewport `json:"viewport,omitempty"`
}

/*
SetDeviceMetricsOverrideResult represents the result of calls to
Emulation.setDeviceMetricsOverride.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
*/
type SetDeviceMetricsOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetEmitTouchEventsForMouseParams represents Emulation.setEmitTouchEventsForMouse parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
*/
type SetEmitTouchEventsForMouseParams struct {
	// Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`

	// Optional. Touch/gesture events configuration. Default: current platform.
	// Allowed values:
	//	- mobile
	//	- desktop
	Configuration ConfigurationEnum `json:"configuration,omitempty"`
}

/*
SetEmitTouchEventsForMouseResult represents the result of calls to
Emulation.setEmitTouchEventsForMouse.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
*/
type SetEmitTouchEventsForMouseResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
SetEmulatedMediaResult represents the result of calls to Emulation.setEmulatedMedia.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
*/
type SetEmulatedMediaResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
SetGeolocationOverrideResult represents the result of calls to Emulation.setGeolocationOverride.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
*/
type SetGeolocationOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
SetNavigatorOverridesResult represents the result of calls to Emulation.setNavigatorOverrides.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
*/
type SetNavigatorOverridesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
SetPageScaleFactorResult represents the result of calls to Emulation.setPageScaleFactor.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
*/
type SetPageScaleFactorResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
SetScriptExecutionDisabledResult represents the result of calls to
Emulation.setScriptExecutionDisabled.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
*/
type SetScriptExecutionDisabledResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
SetTouchEmulationEnabledResult represents the result of calls to
Emulation.setTouchEmulationEnabled.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
*/
type SetTouchEmulationEnabledResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetVirtualTimePolicyParams represents Emulation.setVirtualTimePolicy parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
*/
type SetVirtualTimePolicyParams struct {
	// desc.
	Policy VirtualTimePolicy `json:"policy"`

	// Optional. If set, after this many virtual milliseconds have elapsed
	// virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	Budget int `json:"budget,omitempty"`

	// Optional. If set this specifies the maximum number of tasks that can be
	// run before virtual is forced forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount int `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
}

/*
SetVirtualTimePolicyResult represents the result of calls to Emulation.setVirtualTimePolicy.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
*/
type SetVirtualTimePolicyResult struct {
	// Absolute timestamp at which virtual time was first enabled (milliseconds
	// since epoch).
	VirtualTimeBase runtime.Timestamp `json:"virtualTimeBase"`

	// Error information related to executing this method
	Err error `json:"-"`
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
SetVisibleSizeResult represents the result of calls to Emulation.setVisibleSize.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
*/
type SetVisibleSizeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
