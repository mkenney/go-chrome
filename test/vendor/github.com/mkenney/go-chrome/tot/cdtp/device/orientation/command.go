package orientation

/*
ClearOverrideResult represents the result of calls to DeviceOrientation.clearOverride.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
*/
type ClearOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetOverrideParams represents DeviceOrientation.setDeviceOrientationOverride
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
type SetOverrideParams struct {
	// Mock alpha.
	Alpha int64 `json:"alpha"`

	// Mock beta.
	Beta int64 `json:"beta"`

	// Mock gamma.
	Gamma int64 `json:"gamma"`
}

/*
SetOverrideResult represents the result of calls to DeviceOrientation.setOverride.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
type SetOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
