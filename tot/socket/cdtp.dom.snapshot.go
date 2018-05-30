package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/dom/snapshot"
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
Disable disables the DOM snapshot functionality for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-disable
*/
func (protocol *DOMSnapshotProtocol) Disable() <-chan *snapshot.DisableResult {
	resultChan := make(chan *snapshot.DisableResult)
	command := NewCommand(protocol.Socket, "DOMSnapshot.disable", nil)
	result := &snapshot.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables the DOM snapshot functionality for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-enable
*/
func (protocol *DOMSnapshotProtocol) Enable() <-chan *snapshot.EnableResult {
	resultChan := make(chan *snapshot.EnableResult)
	command := NewCommand(protocol.Socket, "DOMSnapshot.enable", nil)
	result := &snapshot.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Get returns a document snapshot, including the full DOM tree of the root node
(including iframes, template contents, and imported documents) in a flattened
array, as well as layout and white-listed computed style information for the
nodes. Shadow DOM in the returned DOM tree is flattened.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
*/
func (protocol *DOMSnapshotProtocol) Get(
	params *snapshot.GetParams,
) <-chan *snapshot.GetResult {
	resultChan := make(chan *snapshot.GetResult)
	command := NewCommand(protocol.Socket, "DOMSnapshot.getSnapshot", params)
	result := &snapshot.GetResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}
