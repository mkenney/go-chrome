package protocol

import (
	accessibility "github.com/mkenney/go-chrome/protocol/accessibility"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Accessibility is a struct that provides a namespace for the Chrome Accessibility protocol methods.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/
*/
var Accessibility = _accessibility{}

type _accessibility struct{}

/*
GetPartialAXTree fetches the accessibility node and partial accessibility tree for this DOM node, if
it exists.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getPartialAXTree
*/
func (_accessibility) GetPartialAXTree(
	socket sock.Socketer,
	params *accessibility.PartialAXTreeParams,
) (*accessibility.PartialAXTreeResult, error) {
	command := sock.NewCommand("Accessibility.getPartialAXTree", params)
	result := &accessibility.PartialAXTreeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return nil, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
