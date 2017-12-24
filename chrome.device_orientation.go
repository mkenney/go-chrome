package chrome

import (
	deviceOrientation "github.com/mkenney/go-chrome/device_orientation"
	"github.com/mkenney/go-chrome/protocol"
)

/*
DeviceOrientation - https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/
EXPERIMENTAL
*/
type DeviceOrientation struct{}

/*
ClearDeviceOrientationOverride clears the overridden Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
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

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
*/
func (DeviceOrientation) SetDeviceOrientationOverride(
	socket *Socket,
	params *deviceOrientation.SetDeviceOrientationOverrideParams,
) error {
	command := &protocol.Command{
		Method: "DeviceOrientation.setDeviceOrientationOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
