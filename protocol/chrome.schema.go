package protocol

import (
	"encoding/json"

	schema "github.com/mkenney/go-chrome/protocol/schema"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Schema is a struct that provides a namespace for the Chrome Schema protocol methods.

DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Schema/
*/
type Schema struct{}

/*
GetDomains returns supported domains.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/#method-getDomains
*/
func (Schema) GetDomains(
	socket sock.Socketer,
) (schema.GetDomainsResult, error) {
	command := sock.NewCommand("Schema.getDomains", nil)
	result := schema.GetDomainsResult{}
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
