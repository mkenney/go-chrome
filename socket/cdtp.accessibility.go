package socket

import (
	"encoding/json"

	accessibility "github.com/mkenney/go-chrome/cdtp/accessibility"
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
) (*accessibility.PartialAXTreeResult, error) {
	command := NewCommand("Accessibility.getPartialAXTree", params)
	result := &accessibility.PartialAXTreeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return nil, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}
