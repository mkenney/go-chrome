package protocol

import (
	"encoding/json"

	layerTree "github.com/mkenney/go-chrome/protocol/layer_tree"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
LayerTree is a struct that provides a namespace for the Chrome LayerTree protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/
*/
type LayerTree struct{}

/*
CompositingReasons provides the reasons why the given layer was composited.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
*/
func (LayerTree) CompositingReasons(
	socket *sock.Socket,
	params *layerTree.CompositingReasonsParams,
) (layerTree.CompositingReasonsResult, error) {
	command := &sock.Command{
		Method: "LayerTree.compositingReasons",
		Params: params,
	}
	result := layerTree.CompositingReasonsResult{}
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

/*
Disable disables compositing tree inspection.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-disable
*/
func (LayerTree) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "LayerTree.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables compositing tree inspection.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-enable
*/
func (LayerTree) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "LayerTree.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
LoadSnapshot returns the snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
func (LayerTree) LoadSnapshot(
	socket *sock.Socket,
	params *layerTree.LoadSnapshotParams,
) (layerTree.LoadSnapshotResult, error) {
	command := &sock.Command{
		Method: "LayerTree.loadSnapshot",
		Params: params,
	}
	result := layerTree.LoadSnapshotResult{}
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

/*
MakeSnapshot returns the layer snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
*/
func (LayerTree) MakeSnapshot(
	socket *sock.Socket,
	params *layerTree.MakeSnapshotParams,
) (layerTree.MakeSnapshotResult, error) {
	command := &sock.Command{
		Method: "LayerTree.makeSnapshot",
		Params: params,
	}
	result := layerTree.MakeSnapshotResult{}
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

/*
ProfileSnapshot profiles a snapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
*/
func (LayerTree) ProfileSnapshot(
	socket *sock.Socket,
	params *layerTree.ProfileSnapshotParams,
) (layerTree.ProfileSnapshotResult, error) {
	command := &sock.Command{
		Method: "LayerTree.profileSnapshot",
		Params: params,
	}
	result := layerTree.ProfileSnapshotResult{}
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

/*
ReleaseSnapshot releases layer snapshot captured by the back-end.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-releaseSnapshot
*/
func (LayerTree) ReleaseSnapshot(
	socket *sock.Socket,
	params *layerTree.ReleaseSnapshotParams,
) error {
	command := &sock.Command{
		Method: "LayerTree.releaseSnapshot",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ReplaySnapshot replays the layer snapshot and returns the resulting bitmap.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
func (LayerTree) ReplaySnapshot(
	socket *sock.Socket,
	params *layerTree.ReplaySnapshotParams,
) (layerTree.ReplaySnapshotResult, error) {
	command := &sock.Command{
		Method: "LayerTree.replaySnapshot",
		Params: params,
	}
	result := layerTree.ReplaySnapshotResult{}
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

/*
SnapshotCommandLog replays the layer snapshot and returns canvas log.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
*/
func (LayerTree) SnapshotCommandLog(
	socket *sock.Socket,
	params *layerTree.SnapshotCommandLogParams,
) (layerTree.SnapshotCommandLogResult, error) {
	command := &sock.Command{
		Method: "LayerTree.snapshotCommandLog",
		Params: params,
	}
	result := layerTree.SnapshotCommandLogResult{}
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

/*
OnLayerPainted adds a handler to the LayerTree.layerPainted event. LayerTree.layerPainted fires when the layer
is painted.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerPainted
*/
func (LayerTree) OnLayerPainted(
	socket *sock.Socket,
	callback func(event *layerTree.LayerPaintedEvent),
) {
	handler := sock.NewEventHandler(
		"LayerTree.layerPainted",
		func(name string, params []byte) {
			event := &layerTree.LayerPaintedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (LayerTree) OnLayerTreeDidChange(
	socket *sock.Socket,
	callback func(event *layerTree.DidChangeEvent),
) {
	handler := sock.NewEventHandler(
		"LayerTree.layerTreeDidChange",
		func(name string, params []byte) {
			event := &layerTree.DidChangeEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
