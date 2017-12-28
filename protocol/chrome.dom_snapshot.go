package protocol

import (
	"encoding/json"

	domSnapshot "github.com/mkenney/go-chrome/protocol/dom_snapshot"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
DOMSnapshot is a struct that provides a namespace for the Chrome DOMSnapshot protocol methods.

The DOMSnapshot protocol facilitates obtaining document snapshots with DOM, layout, and style
information.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/
*/
var DOMSnapshot = _domSnapshot{}

type _domSnapshot struct{}

/*
Get returns a document snapshot, including the full DOM tree of the root node (including iframes,
template contents, and imported documents) in a flattened array, as well as layout and white-listed
computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
*/
func (_domSnapshot) Get(
	socket sock.Socketer,
	params *domSnapshot.GetParams,
) (domSnapshot.GetResult, error) {
	command := sock.NewCommand("DOMSnapshot.getSnapshot", params)
	result := domSnapshot.GetResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}
