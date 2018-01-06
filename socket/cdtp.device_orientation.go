package socket

import (
	deviceOrientation "github.com/mkenney/go-chrome/cdtp/device_orientation"
)

/*
DeviceOrientationProtocol provides a namespace for the Chrome DeviceOrientation
protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/
EXPERIMENTAL.
*/
type DeviceOrientationProtocol struct {
	Socket Socketer
}

/*
ClearOverride clears the overridden Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
*/
func (protocol *DeviceOrientationProtocol) ClearOverride() chan *deviceOrientation.ClearOverrideResult {
	resultChan := make(chan *deviceOrientation.ClearOverrideResult)
	command := NewCommand(protocol.Socket, "DeviceOrientation.clearDeviceOrientationOverride", nil)
	result := &deviceOrientation.ClearOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetOverride overrides the Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
func (protocol *DeviceOrientationProtocol) SetOverride(
	params *deviceOrientation.SetOverrideParams,
) chan *deviceOrientation.SetOverrideResult {
	resultChan := make(chan *deviceOrientation.SetOverrideResult)
	command := NewCommand(protocol.Socket, "DeviceOrientation.setDeviceOrientationOverride", params)
	result := &deviceOrientation.SetOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}
