package chrome

import (
	audits "app/chrome/audits"
	"app/chrome/protocol"
	"encoding/json"
)

/*
Audits - https://chromedevtools.github.io/devtools-protocol/tot/Audits/
Allows investigation of page violations and possible improvements. EXPERIMENTAL
*/
type Audits struct{}

/*
GetEncodedResponse returns the response body and size if it were re-encoded with the specified
settings. Only applies to images.
*/
func (Audits) GetEncodedResponse(
	socket *Socket,
	params *audits.GetEncodedResponseParams,
) (audits.GetEncodedResponseResult, error) {
	command := &protocol.Command{
		Method: "Audits.getEncodedResponse",
		Params: params,
	}
	result := audits.GetEncodedResponseResult{}
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
