package socket

import (
	"encoding/json"

	systemInfo "github.com/mkenney/go-chrome/cdtp/system_info"
)

/*
SystemInfoProtocol provides a namespace for the Chrome SystemInfo protocol
methods. The SystemInfo protocol defines methods and events for querying
low-level system information.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/ EXPERIMENTAL.
*/
type SystemInfoProtocol struct {
	Socket Socketer
}

/*
GetInfo returns information about the system.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
*/
func (protocol *SystemInfoProtocol) GetInfo() chan *systemInfo.GetInfoResult {
	resultChan := make(chan *systemInfo.GetInfoResult)
	command := NewCommand(protocol.Socket, "SystemInfo.getInfo", nil)
	result := &systemInfo.GetInfoResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}
