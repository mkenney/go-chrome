package chrome

import "app/chrome/protocol"

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
func (DOMSnapshot) GetSnapshot(socket *Socket, params *dom_snapshot.GetSnapshotParams) error {
	command := &protocol.Command{
		method: "DOMSnapshot.getSnapshot",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
