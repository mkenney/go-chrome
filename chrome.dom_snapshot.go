package chrome

import (
	"encoding/json"

	domSnapshot "github.com/mkenney/go-chrome/dom_snapshot"
	"github.com/mkenney/go-chrome/protocol"
)

/*
DOMSnapshot facilitates obtaining document snapshots with DOM, layout, and style information.
*/
type DOMSnapshot struct{}

/*
GetSnapshot - https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/
Returns a document snapshot, including the full DOM tree of the root node (including iframes,
template contents, and imported documents) in a flattened array, as well as layout and white-listed
computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
*/
func (DOMSnapshot) GetSnapshot(
	socket *Socket,
	params *domSnapshot.GetSnapshotParams,
) (domSnapshot.GetSnapshotResult, error) {
	command := &protocol.Command{
		Method: "DOMSnapshot.getSnapshot",
		Params: params,
	}
	result := domSnapshot.GetSnapshotResult{}
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
