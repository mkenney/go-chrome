package socket

import (
	"encoding/json"

	schema "github.com/mkenney/go-chrome/tot/cdtp/schema"
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
func (protocol *SchemaProtocol) GetDomains() chan *schema.GetDomainsResult {
	resultChan := make(chan *schema.GetDomainsResult)
	command := NewCommand(protocol.Socket, "Schema.getDomains", nil)
	result := &schema.GetDomainsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}
