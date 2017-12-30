package socket

import (
	"encoding/json"

	schema "github.com/mkenney/go-chrome/protocol/schema"
)

/*
SchemaProtocol provides a namespace for the Chrome Schema protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/
DEPRECATED.
*/
type SchemaProtocol struct {
	Socket Socketer
}

/*
GetDomains returns supported domains.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/#method-getDomains
*/
func (protocol *SchemaProtocol) GetDomains() (*schema.GetDomainsResult, error) {
	command := NewCommand("Schema.getDomains", nil)
	result := &schema.GetDomainsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}
