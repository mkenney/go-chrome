package protocol

import (
	"encoding/json"

	systemInfo "github.com/mkenney/go-chrome/protocol/system_info"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
SystemInfo is a struct that provides a namespace for the Chrome SystemInfo protocol methods.

The SystemInfo protocol defines methods and events for querying low-level system information.
EXPERIMENTAL

- https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/
*/
type SystemInfo struct{}

/*
GetInfo returns information about the system.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
*/
func (SystemInfo) GetInfo(
	socket *sock.Socket,
) (systemInfo.GetInfoResult, error) {
	command := &sock.Command{
		Method: "SystemInfo.getInfo",
	}
	result := systemInfo.GetInfoResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}
