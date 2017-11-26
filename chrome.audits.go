package chrome

import (
	"app/chrome/protocol"
	audits "app/chrome/protocol/audits"
)

type Audits struct{}

/*
GetEncodedResponse returns the response body and size if it were re-encoded with the specified
settings. Only applies to images.
*/
func (Audits) GetEncodedResponse(socket *Socket, params *audits.GetEncodedResponseParams) error {
	command := new(protocol.Command)
	command.method = "Audits.getEncodedResponse"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
