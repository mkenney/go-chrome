package socket

import (
	systemInfo "github.com/mkenney/go-chrome/protocol/system_info"
)

/*
SystemInfoProtocol provides a namespace for the Chrome SystemInfo protocol methods. The SystemInfo
protocol defines methods and events for querying low-level system information.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/ EXPERIMENTAL.
*/
type SystemInfoProtocol struct {
	Socket Socketer
}

/*
GetInfo returns information about the system.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
*/
func (protocol *SystemInfoProtocol) GetInfo() (*systemInfo.GetInfoResult, error) {
	command := NewCommand("SystemInfo.getInfo", nil)
	result := &systemInfo.GetInfoResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
