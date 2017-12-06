package chrome

import "app/chrome/protocol"

/*
DeviceOrientation - https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/
EXPERIMENTAL
*/
type DeviceOrientation struct{}

/*
ClearDeviceOrientationOverride clears the overridden Device Orientation.
*/
func (DeviceOrientation) ClearDeviceOrientationOverride(
	socket *Socket,
	params *device_orientation.ClearDeviceOrientationOverrideParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DeviceOrientation.clearDeviceOrientationOverride",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetDeviceOrientationOverride overrides the Device Orientation.
*/
func (DeviceOrientation) SetDeviceOrientationOverride(
	socket *Socket,
	params *device_orientation.SetDeviceOrientationOverrideParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DeviceOrientation.setDeviceOrientationOverride",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}
