package socket

import (
	"encoding/json"

	domSnapshot "github.com/mkenney/go-chrome/protocol/dom_snapshot"
)

/*
DOMSnapshotProtocol provides a namespace for the Chrome DOMSnapshot protocol
methods. The DOMSnapshot protocol facilitates obtaining document snapshots with
DOM, layout, and style information.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/
*/
type DOMSnapshotProtocol struct {
	Socket Socketer
}

/*
Get returns a document snapshot, including the full DOM tree of the root node
(including iframes, template contents, and imported documents) in a flattened
array, as well as layout and white-listed computed style information for the
nodes. Shadow DOM in the returned DOM tree is flattened.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
*/
func (protocol *DOMSnapshotProtocol) Get(
	params *domSnapshot.GetParams,
) (*domSnapshot.GetResult, error) {
	command := NewCommand("DOMSnapshot.getSnapshot", params)
	result := &domSnapshot.GetResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}
