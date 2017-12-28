package protocol

import (
	"encoding/json"

	audits "github.com/mkenney/go-chrome/protocol/audits"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Audits is a struct that provides a namespace for the Chrome Audits protocol methods.

The Audits protocol allows investigation of page violations and possible improvements.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Audits/
*/
var Audits = _audits{}

type _audits struct{}

/*
GetEncodedResponse returns the response body and size if it were re-encoded with the specified
settings. Only applies to images.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
*/
func (_audits) GetEncodedResponse(
	socket sock.Socketer,
	params *audits.GetEncodedResponseParams,
) (audits.GetEncodedResponseResult, error) {
	command := sock.NewCommand("Audits.getEncodedResponse", params)
	result := audits.GetEncodedResponseResult{}
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
