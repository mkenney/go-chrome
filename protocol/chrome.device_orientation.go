package protocol

import (
	deviceOrientation "github.com/mkenney/go-chrome/protocol/device_orientation"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
DeviceOrientation is a struct that provides a namespace for the Chrome DeviceOrientation protocol
methods. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/
*/
var DeviceOrientation = _deviceOrientation{}

type _deviceOrientation struct{}

/*
ClearOverride clears the overridden Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
*/
func (_deviceOrientation) ClearOverride(
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
func (_deviceOrientation) SetOverride(
	socket sock.Socketer,
	params *deviceOrientation.SetOverrideParams,
) error {
	command := sock.NewCommand("DeviceOrientation.setDeviceOrientationOverride", params)
	socket.SendCommand(command)
	return command.Error()
}
