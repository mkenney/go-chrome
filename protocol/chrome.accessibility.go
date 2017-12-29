package protocol

import (
	accessibility "github.com/mkenney/go-chrome/protocol/accessibility"
	sock "github.com/mkenney/go-chrome/socket"
	log "github.com/sirupsen/logrus"
)

/*
Accessibility provides a namespace for the Chrome Accessibility protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/ EXPERIMENTAL.
*/
var Accessibility = AccessibilityProtocol{}

/*
AccessibilityProtocol defines the Accessibility protocol methods.
*/
type AccessibilityProtocol struct{}

/*
GetPartialAXTree fetches the accessibility node and partial accessibility tree for this DOM node, if
it exists.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getPartialAXTree
*/
func (AccessibilityProtocol) GetPartialAXTree(
	socket sock.Socketer,
	params *accessibility.PartialAXTreeParams,
) (*accessibility.PartialAXTreeResult, error) {
	command := sock.NewCommand("Accessibility.getPartialAXTree", params)
	log.Infof("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ command id %d", command.ID())
	result := &accessibility.PartialAXTreeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return nil, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
