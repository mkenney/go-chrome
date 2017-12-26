package protocol

import (
	"encoding/json"

	domSnapshot "github.com/mkenney/go-chrome/protocol/dom_snapshot"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
DOMSnapshot facilitates obtaining document snapshots with DOM, layout, and style information.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/
*/
type DOMSnapshot struct{}

/*
Get returns a document snapshot, including the full DOM tree of the root node (including iframes,
template contents, and imported documents) in a flattened array, as well as layout and white-listed
computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
*/
func (DOMSnapshot) Get(
	socket *sock.Socket,
	params *domSnapshot.GetParams,
) (domSnapshot.GetResult, error) {
	command := &sock.Command{
		Method: "DOMSnapshot.getSnapshot",
		Params: params,
	}
	result := domSnapshot.GetResult{}
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
