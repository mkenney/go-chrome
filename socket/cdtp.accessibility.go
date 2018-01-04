package socket

import (
	"encoding/json"

	accessibility "github.com/mkenney/go-chrome/cdtp/accessibility"
	log "github.com/sirupsen/logrus"
)

/*
AccessibilityProtocol provides a namespace for the Chrome Accessibility protocol
methods.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/
EXPERIMENTAL.
*/
type AccessibilityProtocol struct {
	Socket Socketer
}

/*
GetPartialAXTree fetches the accessibility node and partial accessibility tree
for this DOM node, if it exists.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getPartialAXTree
*/
func (protocol *AccessibilityProtocol) GetPartialAXTree(
	params *accessibility.PartialAXTreeParams,
) chan *accessibility.PartialAXTreeResult {
	resultChan := make(chan *accessibility.PartialAXTreeResult)
	command := NewCommand(protocol.Socket, "Accessibility.getPartialAXTree", params)
	result := &accessibility.PartialAXTreeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		log.Debugf("GOT HERE")
		resultChan <- result
	}()

	return resultChan
}
