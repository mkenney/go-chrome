package chrome

import (
	accessibility "app/chrome/accessibility"
	"app/chrome/protocol"
)

/*
Accessibility - https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/
EXPERIMENTAL
*/
type Accessibility struct{}

/*
GetPartialAXTree fetches the accessibility node and partial accessibility tree for this DOM node, if
it exists.
*/
func (Accessibility) GetPartialAXTree(
	socket *Socket,
	params *accessibility.PartialAXTreeParams,
) (accessibility.PartialAXTreeResult, error) {
	command := &protocol.Command{
		Method: "Accessibility.getPartialAXTree",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(accessibility.PartialAXTreeResult), command.Err
}
