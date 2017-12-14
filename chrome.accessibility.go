package chrome

import (
	accessibility "app/chrome/accessibility"
	"app/chrome/protocol"
	"encoding/json"
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
	result := accessibility.PartialAXTreeResult{}
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
