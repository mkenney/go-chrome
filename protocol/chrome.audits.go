package protocol

import (
	audits "github.com/mkenney/go-chrome/protocol/audits"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Audits provides a namespace for the Chrome Audits protocol methods. The Audits protocol allows
investigation of page violations and possible improvements.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/ EXPERIMENTAL.
*/
var Audits = AuditsProtocol{}

/*
AuditsProtocol defines the Audits protocol methods.
*/
type AuditsProtocol struct{}

/*
GetEncodedResponse returns the response body and size if it were re-encoded with the specified
settings. Only applies to images.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
*/
func (AuditsProtocol) GetEncodedResponse(
	socket sock.Socketer,
	params *audits.GetEncodedResponseParams,
) (*audits.GetEncodedResponseResult, error) {
	command := sock.NewCommand("Audits.getEncodedResponse", params)
	result := &audits.GetEncodedResponseResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
