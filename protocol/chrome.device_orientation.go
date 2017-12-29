package protocol

import (
	deviceOrientation "github.com/mkenney/go-chrome/protocol/device_orientation"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
DeviceOrientation provides a namespace for the Chrome DeviceOrientation protocol methods.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/
*/
var DeviceOrientation = DeviceOrientationProtocol{}

/*
DeviceOrientationProtocol defines the DeviceOrientation protocol methods.
*/
type DeviceOrientationProtocol struct{}

/*
ClearOverride clears the overridden Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
*/
func (DeviceOrientationProtocol) ClearOverride(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DeviceOrientation.clearDeviceOrientationOverride", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetOverride overrides the Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
func (DeviceOrientationProtocol) SetOverride(
	socket sock.Socketer,
	params *deviceOrientation.SetOverrideParams,
) error {
	command := sock.NewCommand("DeviceOrientation.setDeviceOrientationOverride", params)
	socket.SendCommand(command)
	return command.Error()
}
