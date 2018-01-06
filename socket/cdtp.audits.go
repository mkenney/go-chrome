package socket

import (
	"encoding/json"

	audits "github.com/mkenney/go-chrome/cdtp/audits"
)

/*
AuditsProtocol provides a namespace for the Chrome Audits protocol methods. The
Audits protocol allows investigation of page violations and possible
improvements.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/ EXPERIMENTAL.
*/
type AuditsProtocol struct {
	Socket Socketer
}

/*
GetEncodedResponse returns the response body and size if it were re-encoded with
the specified settings. Only applies to images.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
*/
func (protocol *AuditsProtocol) GetEncodedResponse(
	params *audits.GetEncodedResponseParams,
) chan *audits.GetEncodedResponseResult {
	resultChan := make(chan *audits.GetEncodedResponseResult)
	command := NewCommand(protocol.Socket, "Audits.getEncodedResponse", params)
	result := &audits.GetEncodedResponseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}
