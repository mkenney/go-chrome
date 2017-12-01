package chrome

import "app/chrome/protocol"

/*
Accessibility - https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/
EXPERIMENTAL
*/
type Accessibility struct{}

/*
GetPartialAXTree fetches the accessibility node and partial accessibility tree for this DOM node, if
it exists.
*/
func (Accessibility) GetPartialAXTree(socket *Socket, params *accessibility.PartialAXTreeParams) error {
	command := &protocol.Command{
		method: "Accessibility.getPartialAXTree",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
