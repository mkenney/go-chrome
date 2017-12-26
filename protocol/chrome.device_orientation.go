package protocol

import (
	deviceOrientation "github.com/mkenney/go-chrome/protocol/device_orientation"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
DeviceOrientation is a struct that provides a namespace for the Chrome DeviceOrientation protocol
methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/
*/
type DeviceOrientation struct{}

/*
ClearOverride clears the overridden Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
*/
func (DeviceOrientation) ClearOverride(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "DeviceOrientation.clearDeviceOrientationOverride",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetOverride overrides the Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
func (DeviceOrientation) SetOverride(
	socket *sock.Socket,
	params *deviceOrientation.SetOverrideParams,
) error {
	command := &sock.Command{
		Method: "DeviceOrientation.setDeviceOrientationOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
