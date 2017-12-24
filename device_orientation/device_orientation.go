/*
Package DeviceOrientation provides type definitions for use with the Chrome DeviceOrientation
protocol

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/
*/
package DeviceOrientation

/*
SetOverrideParams represents DeviceOrientation.setDeviceOrientationOverride
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
type SetOverrideParams struct {
	// Mock alpha.
	Alpha float64 `json:"alpha"`

	// Mock beta.
	Beta float64 `json:"beta"`

	// Mock gamma.
	Gamma float64 `json:"gamma"`
}
