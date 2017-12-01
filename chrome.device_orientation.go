package chrome

import (
	do "app/chrome/device_orientation"
	"app/chrome/protocol"
)

/*
DeviceOrientation EXPERIMENTAL
*/
type DeviceOrientation struct{}

/*
ClearDeviceOrientationOverride clears the overridden Device Orientation.
*/
func (DeviceOrientation) ClearDeviceOrientationOverride(socket *Socket, params *do.ClearDeviceOrientationOverrideParams) error {
	command := new(protocol.Command)
	command.method = "DeviceOrientation.clearDeviceOrientationOverride"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetDeviceOrientationOverride overrides the Device Orientation.
*/
func (DeviceOrientation) SetDeviceOrientationOverride(socket *Socket, params *do.SetDeviceOrientationOverrideParams) error {
	command := new(protocol.Command)
	command.method = "DeviceOrientation.setDeviceOrientationOverride"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
