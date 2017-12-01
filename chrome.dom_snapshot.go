package chrome

import (
	domsnapshot "app/chrome/domsnapshot"
	"app/chrome/protocol"
)

/*
DOMSnapshot facilitates obtaining document snapshots with DOM, layout, and style information.
*/
type DOMSnapshot struct{}

/*
GetSnapshot returns a document snapshot, including the full DOM tree of the root node (including
iframes, template contents, and imported documents) in a flattened array, as well as layout and
white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
flattened.
*/
func (DOMSnapshot) GetSnapshot(socket *Socket, params *domsnapshot.GetSnapshotParams) error {
	command := new(protocol.Command)
	command.method = "DOMSnapshot.getSnapshot"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
