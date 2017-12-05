package chrome

import "app/chrome/protocol"

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
		method: "Audits.getEncodedResponse",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(audits.GetEncodedResponseResult), command.Err
}
