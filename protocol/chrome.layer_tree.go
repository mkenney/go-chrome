package protocol

import (
	"encoding/json"

	layerTree "github.com/mkenney/go-chrome/protocol/layer_tree"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
LayerTree provides a namespace for the Chrome LayerTree protocol methods. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/
*/
var LayerTree = LayerTreeProtocol{}

/*
LayerTreeProtocol defines the LayerTree protocol methods.
*/
type LayerTreeProtocol struct{}

/*
CompositingReasons provides the reasons why the given layer was composited.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
*/
func (LayerTreeProtocol) CompositingReasons(
	socket sock.Socketer,
	params *layerTree.CompositingReasonsParams,
) (*layerTree.CompositingReasonsResult, error) {
	command := sock.NewCommand("LayerTree.compositingReasons", params)
	result := &layerTree.CompositingReasonsResult{}
	socket.SendCommand(command)

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
func (LayerTreeProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("LayerTree.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables compositing tree inspection.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-enable
*/
func (LayerTreeProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("LayerTree.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
LoadSnapshot returns the snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
func (LayerTreeProtocol) LoadSnapshot(
	socket sock.Socketer,
	params *layerTree.LoadSnapshotParams,
) (*layerTree.LoadSnapshotResult, error) {
	command := sock.NewCommand("LayerTree.loadSnapshot", params)
	result := &layerTree.LoadSnapshotResult{}
	socket.SendCommand(command)

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
func (LayerTreeProtocol) MakeSnapshot(
	socket sock.Socketer,
	params *layerTree.MakeSnapshotParams,
) (*layerTree.MakeSnapshotResult, error) {
	command := sock.NewCommand("LayerTree.makeSnapshot", params)
	result := &layerTree.MakeSnapshotResult{}
	socket.SendCommand(command)

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
func (LayerTreeProtocol) ProfileSnapshot(
	socket sock.Socketer,
	params *layerTree.ProfileSnapshotParams,
) (*layerTree.ProfileSnapshotResult, error) {
	command := sock.NewCommand("LayerTree.profileSnapshot", params)
	result := &layerTree.ProfileSnapshotResult{}
	socket.SendCommand(command)

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
func (LayerTreeProtocol) ReleaseSnapshot(
	socket sock.Socketer,
	params *layerTree.ReleaseSnapshotParams,
) error {
	command := sock.NewCommand("LayerTree.releaseSnapshot", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
ReplaySnapshot replays the layer snapshot and returns the resulting bitmap.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
func (LayerTreeProtocol) ReplaySnapshot(
	socket sock.Socketer,
	params *layerTree.ReplaySnapshotParams,
) (*layerTree.ReplaySnapshotResult, error) {
	command := sock.NewCommand("LayerTree.replaySnapshot", params)
	result := &layerTree.ReplaySnapshotResult{}
	socket.SendCommand(command)

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
func (LayerTreeProtocol) SnapshotCommandLog(
	socket sock.Socketer,
	params *layerTree.SnapshotCommandLogParams,
) (*layerTree.SnapshotCommandLogResult, error) {
	command := sock.NewCommand("LayerTree.snapshotCommandLog", params)
	result := &layerTree.SnapshotCommandLogResult{}
	socket.SendCommand(command)

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
func (LayerTreeProtocol) OnLayerPainted(
	socket sock.Socketer,
	callback func(event *layerTree.LayerPaintedEvent),
) {
	handler := sock.NewEventHandler(
		"LayerTree.layerPainted",
		func(response *sock.Response) {
			event := &layerTree.LayerPaintedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLayerTreeDidChange adds a handler to the LayerTree.DidChange event. LayerTree.DidChange fires when
the layer tree changes.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerTreeDidChange
*/
func (LayerTreeProtocol) OnLayerTreeDidChange(
	socket sock.Socketer,
	callback func(event *layerTree.DidChangeEvent),
) {
	handler := sock.NewEventHandler(
		"LayerTree.layerTreeDidChange",
		func(response *sock.Response) {
			event := &layerTree.DidChangeEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
