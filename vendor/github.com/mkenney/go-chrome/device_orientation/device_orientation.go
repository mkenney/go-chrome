package DeviceOrientation

/*
SetDeviceOrientationOverrideParams represents DeviceOrientation.setDeviceOrientationOverride
parameters.
*/
type SetDeviceOrientationOverrideParams struct {
	// Mock alpha.
	Alpha float64 `json:"alpha"`

	// Mock beta.
	Beta float64 `json:"beta"`

	// Mock gamma.
	Gamma float64 `json:"gamma"`
}
