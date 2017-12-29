package socket

import (
	"encoding/json"

	layerTree "github.com/mkenney/go-chrome/protocol/layer_tree"
	log "github.com/sirupsen/logrus"
)

/*
LayerTreeProtocol provides a namespace for the Chrome LayerTree protocol methods.

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
) (*layerTree.CompositingReasonsResult, error) {
	command := NewCommand("LayerTree.compositingReasons", params)
	result := &layerTree.CompositingReasonsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
Disable disables compositing tree inspection.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-disable
*/
func (protocol *LayerTreeProtocol) Disable() error {
	command := NewCommand("LayerTree.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables compositing tree inspection.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-enable
*/
func (protocol *LayerTreeProtocol) Enable() error {
	command := NewCommand("LayerTree.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
LoadSnapshot returns the snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
func (protocol *LayerTreeProtocol) LoadSnapshot(
	params *layerTree.LoadSnapshotParams,
) (*layerTree.LoadSnapshotResult, error) {
	command := NewCommand("LayerTree.loadSnapshot", params)
	result := &layerTree.LoadSnapshotResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
MakeSnapshot returns the layer snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
*/
func (protocol *LayerTreeProtocol) MakeSnapshot(
	params *layerTree.MakeSnapshotParams,
) (*layerTree.MakeSnapshotResult, error) {
	command := NewCommand("LayerTree.makeSnapshot", params)
	result := &layerTree.MakeSnapshotResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
ProfileSnapshot profiles a snapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
*/
func (protocol *LayerTreeProtocol) ProfileSnapshot(
	params *layerTree.ProfileSnapshotParams,
) (*layerTree.ProfileSnapshotResult, error) {
	command := NewCommand("LayerTree.profileSnapshot", params)
	result := &layerTree.ProfileSnapshotResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
ReleaseSnapshot releases layer snapshot captured by the back-end.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-releaseSnapshot
*/
func (protocol *LayerTreeProtocol) ReleaseSnapshot(
	params *layerTree.ReleaseSnapshotParams,
) error {
	command := NewCommand("LayerTree.releaseSnapshot", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ReplaySnapshot replays the layer snapshot and returns the resulting bitmap.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
func (protocol *LayerTreeProtocol) ReplaySnapshot(
	params *layerTree.ReplaySnapshotParams,
) (*layerTree.ReplaySnapshotResult, error) {
	command := NewCommand("LayerTree.replaySnapshot", params)
	result := &layerTree.ReplaySnapshotResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
SnapshotCommandLog replays the layer snapshot and returns canvas log.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
*/
func (protocol *LayerTreeProtocol) SnapshotCommandLog(
	params *layerTree.SnapshotCommandLogParams,
) (*layerTree.SnapshotCommandLogResult, error) {
	command := NewCommand("LayerTree.snapshotCommandLog", params)
	result := &layerTree.SnapshotCommandLogResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
OnLayerPainted adds a handler to the LayerTree.layerPainted event. LayerTree.layerPainted fires when the layer
is painted.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerPainted
*/
func (protocol *LayerTreeProtocol) OnLayerPainted(
	callback func(event *layerTree.LayerPaintedEvent),
) {
	handler := NewEventHandler(
		"LayerTree.layerPainted",
		func(response *Response) {
			event := &layerTree.LayerPaintedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLayerTreeDidChange adds a handler to the LayerTree.DidChange event. LayerTree.DidChange fires when
the layer tree changes.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerTreeDidChange
*/
func (protocol *LayerTreeProtocol) OnLayerTreeDidChange(
	callback func(event *layerTree.DidChangeEvent),
) {
	handler := NewEventHandler(
		"LayerTree.layerTreeDidChange",
		func(response *Response) {
			event := &layerTree.DidChangeEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
