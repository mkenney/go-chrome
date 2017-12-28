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
var SystemInfo = _systemInfo{}

type _systemInfo struct{}

/*
GetInfo returns information about the system.

https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
*/
func (_systemInfo) GetInfo(
	socket sock.Socketer,
) (systemInfo.GetInfoResult, error) {
	command := sock.NewCommand("SystemInfo.getInfo", nil)
	result := systemInfo.GetInfoResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}
