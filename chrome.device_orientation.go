package chrome

import (
	device_orientation "github.com/mkenney/go-chrome/device_orientation"
	"github.com/mkenney/go-chrome/protocol"
)

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
) error {
	command := &protocol.Command{
		Method: "DeviceOrientation.clearDeviceOrientationOverride",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDeviceOrientationOverride overrides the Device Orientation.
*/
func (DeviceOrientation) SetDeviceOrientationOverride(
	socket *Socket,
	params *device_orientation.SetDeviceOrientationOverrideParams,
) error {
	command := &protocol.Command{
		Method: "DeviceOrientation.setDeviceOrientationOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
