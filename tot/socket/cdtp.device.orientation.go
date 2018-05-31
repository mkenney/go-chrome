package socket

import (
	"github.com/mkenney/go-chrome/tot/device/orientation"
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
func (protocol *DeviceOrientationProtocol) ClearOverride() <-chan *orientation.ClearOverrideResult {
	resultChan := make(chan *orientation.ClearOverrideResult)
	command := NewCommand(protocol.Socket, "DeviceOrientation.clearDeviceOrientationOverride", nil)
	result := &orientation.ClearOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetOverride overrides the Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
func (protocol *DeviceOrientationProtocol) SetOverride(
	params *orientation.SetOverrideParams,
) <-chan *orientation.SetOverrideResult {
	resultChan := make(chan *orientation.SetOverrideResult)
	command := NewCommand(protocol.Socket, "DeviceOrientation.setDeviceOrientationOverride", params)
	result := &orientation.SetOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}
