package socket

import (
	"encoding/json"

	audits "github.com/mkenney/go-chrome/protocol/audits"
)

/*
AuditsProtocol provides a namespace for the Chrome Audits protocol methods. The Audits protocol
allows investigation of page violations and possible improvements.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/ EXPERIMENTAL.
*/
type AuditsProtocol struct {
	Socket Socketer
}

/*
GetEncodedResponse returns the response body and size if it were re-encoded with the specified
settings. Only applies to images.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
*/
func (protocol *AuditsProtocol) GetEncodedResponse(
	params *audits.GetEncodedResponseParams,
) (*audits.GetEncodedResponseResult, error) {
	command := NewCommand("Audits.getEncodedResponse", params)
	result := &audits.GetEncodedResponseResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}
