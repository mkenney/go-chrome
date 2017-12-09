package chrome

import (
	"app/chrome/protocol"
	system_info "app/chrome/system_info"
)

/*
SystemInfo - https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/
Defines methods and events for querying low-level system information. EXPERIMENTAL
*/
type SystemInfo struct{}

/*
GetInfo returns information about the system.
*/
func (SystemInfo) GetInfo(
	socket *Socket,
) (system_info.GetInfoResult, error) {
	command := &protocol.Command{
		Method: "SystemInfo.getInfo",
	}
	socket.SendCommand(command)
	return command.Result.(system_info.GetInfoResult), command.Err
}
