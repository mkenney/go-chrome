package socket

import (
	"encoding/json"

	layerTree "github.com/mkenney/go-chrome/tot/cdtp/layer/tree"
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
	params *layerTree.CompositingReasonsParams,
) <-chan *layerTree.CompositingReasonsResult {
	resultChan := make(chan *layerTree.CompositingReasonsResult)
	command := NewCommand(protocol.Socket, "LayerTree.compositingReasons", params)
	result := &layerTree.CompositingReasonsResult{}

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
func (protocol *LayerTreeProtocol) Disable() <-chan *layerTree.DisableResult {
	resultChan := make(chan *layerTree.DisableResult)
	command := NewCommand(protocol.Socket, "LayerTree.disable", nil)
	result := &layerTree.DisableResult{}

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
func (protocol *LayerTreeProtocol) Enable() <-chan *layerTree.EnableResult {
	resultChan := make(chan *layerTree.EnableResult)
	command := NewCommand(protocol.Socket, "LayerTree.enable", nil)
	result := &layerTree.EnableResult{}

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
	params *layerTree.LoadSnapshotParams,
) <-chan *layerTree.LoadSnapshotResult {
	resultChan := make(chan *layerTree.LoadSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.loadSnapshot", params)
	result := &layerTree.LoadSnapshotResult{}

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
	params *layerTree.MakeSnapshotParams,
) <-chan *layerTree.MakeSnapshotResult {
	resultChan := make(chan *layerTree.MakeSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.makeSnapshot", params)
	result := &layerTree.MakeSnapshotResult{}

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
	params *layerTree.ProfileSnapshotParams,
) <-chan *layerTree.ProfileSnapshotResult {
	resultChan := make(chan *layerTree.ProfileSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.profileSnapshot", params)
	result := &layerTree.ProfileSnapshotResult{}

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
	params *layerTree.ReleaseSnapshotParams,
) <-chan *layerTree.ReleaseSnapshotResult {
	resultChan := make(chan *layerTree.ReleaseSnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.releaseSnapshot", params)
	result := &layerTree.ReleaseSnapshotResult{}

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
	params *layerTree.ReplaySnapshotParams,
) <-chan *layerTree.ReplaySnapshotResult {
	resultChan := make(chan *layerTree.ReplaySnapshotResult)
	command := NewCommand(protocol.Socket, "LayerTree.replaySnapshot", params)
	result := &layerTree.ReplaySnapshotResult{}

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
	params *layerTree.SnapshotCommandLogParams,
) <-chan *layerTree.SnapshotCommandLogResult {
	resultChan := make(chan *layerTree.SnapshotCommandLogResult)
	command := NewCommand(protocol.Socket, "LayerTree.snapshotCommandLog", params)
	result := &layerTree.SnapshotCommandLogResult{}

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
	callback func(event *layerTree.LayerPaintedEvent),
) {
	handler := NewEventHandler(
		"LayerTree.layerPainted",
		func(response *Response) {
			event := &layerTree.LayerPaintedEvent{}
			json.Unmarshal([]byte(response.Result), event)
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
	callback func(event *layerTree.DidChangeEvent),
) {
	handler := NewEventHandler(
		"LayerTree.layerTreeDidChange",
		func(response *Response) {
			event := &layerTree.DidChangeEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
