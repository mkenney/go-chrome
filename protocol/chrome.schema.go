package protocol

import (
	schema "github.com/mkenney/go-chrome/protocol/schema"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Schema provides a namespace for the Chrome Schema protocol methods. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/
*/
var Schema = SchemaProtocol{}

/*
SchemaProtocol defines the Schema protocol methods.
*/
type SchemaProtocol struct{}

/*
GetDomains returns supported domains.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/#method-getDomains
*/
func (SchemaProtocol) GetDomains(
	socket sock.Socketer,
) (*schema.GetDomainsResult, error) {
	command := sock.NewCommand("Schema.getDomains", nil)
	result := &schema.GetDomainsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
