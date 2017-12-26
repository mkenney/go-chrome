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
	socket *sock.Socket,
) (schema.GetDomainsResult, error) {
	command := &sock.Command{
		Method: "Schema.getDomains",
	}
	result := schema.GetDomainsResult{}
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
