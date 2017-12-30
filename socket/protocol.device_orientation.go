package socket

import (
	deviceOrientation "github.com/mkenney/go-chrome/protocol/device_orientation"
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
func (protocol *DeviceOrientationProtocol) ClearOverride() error {
	command := NewCommand("DeviceOrientation.clearDeviceOrientationOverride", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetOverride overrides the Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
func (protocol *DeviceOrientationProtocol) SetOverride(
	params *deviceOrientation.SetOverrideParams,
) error {
	command := NewCommand("DeviceOrientation.setDeviceOrientationOverride", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}
