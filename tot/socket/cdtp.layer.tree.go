package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/layer/tree"
)

/*
LayerTreeProtocol provides a namespace for the Chrome LayerTree protocol
methods.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/ EXPERIMENTAL.
*/
type LayerTreeProtocol struct {
	Socket Socketer
}

/*
CompositingReasons provides the reasons why the given layer was composited.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
*/
func (protocol *LayerTreeProtocol) CompositingReasons(
	params *tree.CompositingReasonsParams,
) <-chan *tree.CompositingReasonsResult {
	resultChan := make(chan *tree.CompositingReasonsResult)
	command := NewCommand(protocol.Socket, "LayerTree.compositingReasons", params)
	result := &tree.CompositingReasonsResult{}

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

/*
Disable disables compositing tree inspection.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-disable
*/
func (protocol *LayerTreeProtocol) Disable() <-chan *tree.DisableResult {
	resultChan := make(chan *tree.DisableResult)
	command := NewCommand(protocol.Socket, "LayerTree.disable", nil)
	result := &tree.DisableResult{}

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
Enable enables compositing tree inspection.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-enable
*/
func (protocol *LayerTreeProtocol) Enable() <-chan *tree.EnableResult {
	resultChan := make(chan *tree.EnableResult)
	command := NewCommand(protocol.Socket, "LayerTree.enable", nil)
	result := &tree.EnableResult{}

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
LoadSnapshot returns the snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
func (protocol *LayerTreeProtocol) LoadSnapshot(
	params *tree.LoadSnapshotParams,
) <-chan *tree.LoadSnapshotResult {
	resultChan := make(chan *tree.LoadSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.loadSnapshot", params)
	result := &tree.LoadSnapshotResult{}

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

/*
MakeSnapshot returns the layer snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
*/
func (protocol *LayerTreeProtocol) MakeSnapshot(
	params *tree.MakeSnapshotParams,
) <-chan *tree.MakeSnapshotResult {
	resultChan := make(chan *tree.MakeSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.makeSnapshot", params)
	result := &tree.MakeSnapshotResult{}

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

/*
ProfileSnapshot profiles a snapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
*/
func (protocol *LayerTreeProtocol) ProfileSnapshot(
	params *tree.ProfileSnapshotParams,
) <-chan *tree.ProfileSnapshotResult {
	resultChan := make(chan *tree.ProfileSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.profileSnapshot", params)
	result := &tree.ProfileSnapshotResult{}

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

/*
ReleaseSnapshot releases layer snapshot captured by the back-end.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-releaseSnapshot
*/
func (protocol *LayerTreeProtocol) ReleaseSnapshot(
	params *tree.ReleaseSnapshotParams,
) <-chan *tree.ReleaseSnapshotResult {
	resultChan := make(chan *tree.ReleaseSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.releaseSnapshot", params)
	result := &tree.ReleaseSnapshotResult{}

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
ReplaySnapshot replays the layer snapshot and returns the resulting bitmap.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
func (protocol *LayerTreeProtocol) ReplaySnapshot(
	params *tree.ReplaySnapshotParams,
) <-chan *tree.ReplaySnapshotResult {
	resultChan := make(chan *tree.ReplaySnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.replaySnapshot", params)
	result := &tree.ReplaySnapshotResult{}

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

/*
SnapshotCommandLog replays the layer snapshot and returns canvas log.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
*/
func (protocol *LayerTreeProtocol) SnapshotCommandLog(
	params *tree.SnapshotCommandLogParams,
) <-chan *tree.SnapshotCommandLogResult {
	resultChan := make(chan *tree.SnapshotCommandLogResult)
	command := NewCommand(protocol.Socket, "LayerTree.snapshotCommandLog", params)
	result := &tree.SnapshotCommandLogResult{}

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

/*
OnLayerPainted adds a handler to the LayerTree.layerPainted event. LayerTree.layerPainted
fires when the layer is painted.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerPainted
*/
func (protocol *LayerTreeProtocol) OnLayerPainted(
	callback func(event *tree.LayerPaintedEvent),
) {
	handler := NewEventHandler(
		"LayerTree.layerPainted",
		func(response *Response) {
			event := &tree.LayerPaintedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLayerTreeDidChange adds a handler to the LayerTree.DidChange event.
LayerTree.DidChange fires when the layer tree changes.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerTreeDidChange
*/
func (protocol *LayerTreeProtocol) OnLayerTreeDidChange(
	callback func(event *tree.DidChangeEvent),
) {
	handler := NewEventHandler(
		"LayerTree.layerTreeDidChange",
		func(response *Response) {
			event := &tree.DidChangeEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
