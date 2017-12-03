package chrome

import (
	"app/chrome/protocol"
)

/*
SystemInfo - https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/
Defines methods and events for querying low-level system information. EXPERIMENTAL
*/
type SystemInfo struct{}

/*
GetInfo returns information about the system.
*/
func (SystemInfo) GetInfo(socket *Socket, params *system_info.GetInfoParams) error {
	command := &protocol.Command{
		method: "SystemInfo.getInfo",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
