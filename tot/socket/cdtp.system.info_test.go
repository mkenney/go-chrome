package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	systemInfo "github.com/mkenney/go-chrome/tot/cdtp/system/info"
)

func TestSystemInfoGetInfo(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.SystemInfo().GetInfo()
	mockResult := &systemInfo.GetInfoResult{
		GPU: &systemInfo.GPUInfo{
			Devices: []*systemInfo.GPUDevice{{
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.GPU.Devices[0].VendorID != result.GPU.Devices[0].VendorID {
		t.Errorf("Expected %d, got %d", mockResult.GPU.Devices[0].VendorID, result.GPU.Devices[0].VendorID)
	}

	resultChan = mockSocket.SystemInfo().GetInfo()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
