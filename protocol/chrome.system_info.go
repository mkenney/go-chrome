package protocol

import (
	systemInfo "github.com/mkenney/go-chrome/protocol/system_info"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
SystemInfo provides a namespace for the Chrome SystemInfo protocol methods. The SystemInfo protocol
defines methods and events for querying low-level system information. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/
*/
var SystemInfo = SystemInfoProtocol{}

/*
SystemInfoProtocol defines the SystemInfo protocol methods.
*/
type SystemInfoProtocol struct{}

/*
GetInfo returns information about the system.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
*/
func (SystemInfoProtocol) GetInfo(
	socket sock.Socketer,
) (*systemInfo.GetInfoResult, error) {
	command := sock.NewCommand("SystemInfo.getInfo", nil)
	result := &systemInfo.GetInfoResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
