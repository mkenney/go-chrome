package chrome

import (
	"app/chrome/protocol"
)

/*
Schema - https://chromedevtools.github.io/devtools-protocol/tot/Schema/
DEPRECATED
*/
type Schema struct{}

/*
GetDomains returns supported domains.
*/
func (Schema) GetDomains(
	socket *Socket,
) (schema.GetDomainsResult, error) {
	command := &protocol.Command{
		method: "Schema.getDomains",
	}
	socket.SendCommand(command)
	return command.Result.(schema.GetDomainsResult), command.Err
}
