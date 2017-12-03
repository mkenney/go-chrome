package SystemInfo

/*
GetInfoParams represents Overlay.getInfo parameters.
*/
type GetInfoParams struct {
	// Information about the GPUs on the system.
	Gpu GPUInfo `json:"gpu"`

	// A platform-dependent description of the model of the machine. On Mac OS, this is, for
	// example, 'MacBookPro'. Will be the empty string if not supported.
	ModelName string `json:"modelName"`

	// A platform-dependent description of the version of the machine. On Mac OS, this is, for
	// example, '10.1'. Will be the empty string if not supported.
	ModelVersion string `json:"modelVersion"`

	// The command line string used to launch the browser. Will be the empty string if not
	// supported.
	CommandLine string `json:"commandLine"`
}

/////////////////

/*
GPUDevice describes a single graphics processor (GPU).
*/
type GPUDevice struct {
	// PCI ID of the GPU vendor, if available; 0 otherwise.
	VendorID int `json:"vendorId"`

	// PCI ID of the GPU device, if available; 0 otherwise.
	DeviceID int `json:"deviceId"`

	// String description of the GPU vendor, if the PCI ID is not available.
	VendorString string `json:"vendorString"`

	// String description of the GPU device, if the PCI ID is not available.
	DeviceString string `json:"deviceString"`
}

/*
GPUInfo provides information about the GPU(s) on the system.
*/
type GPUInfo struct {
	// The graphics devices on the system. Element 0 is the primary GPU.
	Devices []*GPUDevice `json:"devices"`

	// Optional. An optional dictionary of additional GPU related attributes.
	AuxAttributes map[string]string `json:"auxAttributes,omitempty"`

	// Optional. An optional dictionary of graphics features and their status.
	FeatureStatus map[string]string `json:"featureStatus,omitempty"`

	// An optional array of GPU driver bug workarounds.
	DriverBugWorkarounds []string `json:"driverBugWorkarounds"`
}
