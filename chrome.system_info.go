package chrome

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/protocol"
	system_info "github.com/mkenney/go-chrome/system_info"
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
	result := system_info.GetInfoResult{}
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
