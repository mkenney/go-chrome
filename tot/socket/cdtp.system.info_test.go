package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/system/info"
)

func TestSystemInfoGetInfo(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &info.GetInfoResult{
		GPU: &info.GPUInfo{
			Devices: []*info.GPUDevice{{
				VendorID:     1,
				DeviceID:     1,
				VendorString: "VendorString",
				DeviceString: "DeviceString",
			}},
			AuxAttributes:        map[string]string{"AuxAttributes": "value"},
			FeatureStatus:        map[string]string{"FeatureStatus": "value"},
			DriverBugWorkarounds: []string{"DriverBugWorkarounds1", "DriverBugWorkarounds2"},
		},
		ModelName:    "ModelName",
		ModelVersion: "ModelVersion",
		CommandLine:  "CommandLine",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.SystemInfo().GetInfo()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.GPU.Devices[0].VendorID != result.GPU.Devices[0].VendorID {
		t.Errorf("Expected %d, got %d", mockResult.GPU.Devices[0].VendorID, result.GPU.Devices[0].VendorID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.SystemInfo().GetInfo()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
