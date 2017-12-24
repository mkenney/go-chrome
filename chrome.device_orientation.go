package Chrome

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
ClearOverride clears the overridden Device Orientation.

https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
*/
func (DeviceOrientation) ClearOverride(
	socket *Socket,
) error {
	command := &protocol.Command{
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
	socket *Socket,
	params *deviceOrientation.SetOverrideParams,
) error {
	command := &protocol.Command{
		Method: "DeviceOrientation.setDeviceOrientationOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
