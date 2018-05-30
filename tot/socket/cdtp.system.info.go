package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/system/info"
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
func (protocol *SystemInfoProtocol) GetInfo() <-chan *info.GetInfoResult {
	resultChan := make(chan *info.GetInfoResult)
	command := NewCommand(protocol.Socket, "SystemInfo.getInfo", nil)
	result := &info.GetInfoResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}
