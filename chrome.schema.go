package chrome

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/protocol"
	schema "github.com/mkenney/go-chrome/schema"
)

/*
Schema - https://chromedevtools.github.io/devtools-protocol/tot/Schema/
DEPRECATED
*/
type Schema struct{}

/*
GetDomains returns supported domains.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/#method-getDomains
*/
func (Schema) GetDomains(
	socket *Socket,
) (schema.GetDomainsResult, error) {
	command := &protocol.Command{
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
