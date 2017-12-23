package chrome

import (
	"encoding/json"

	layerTree "github.com/mkenney/go-chrome/layer_tree"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/Sirupsen/logrus"
)

/*
LayerTree - https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/
EXPERIMENTAL
*/
type LayerTree struct{}

/*
CompositingReasons provides the reasons why the given layer was composited.
*/
func (LayerTree) CompositingReasons(
	socket *Socket,
	params *layerTree.CompositingReasonsParams,
) (layerTree.CompositingReasonsResult, error) {
	command := &protocol.Command{
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
*/
func (LayerTree) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "LayerTree.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables compositing tree inspection.
*/
func (LayerTree) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "LayerTree.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
LoadSnapshot returns the snapshot identifier.
*/
func (LayerTree) LoadSnapshot(
	socket *Socket,
	params *layerTree.LoadSnapshotParams,
) (layerTree.LoadSnapshotResult, error) {
	command := &protocol.Command{
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
*/
func (LayerTree) MakeSnapshot(
	socket *Socket,
	params *layerTree.MakeSnapshotParams,
) (layerTree.MakeSnapshotResult, error) {
	command := &protocol.Command{
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
*/
func (LayerTree) ProfileSnapshot(
	socket *Socket,
	params *layerTree.ProfileSnapshotParams,
) (layerTree.ProfileSnapshotResult, error) {
	command := &protocol.Command{
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
*/
func (LayerTree) ReleaseSnapshot(
	socket *Socket,
	params *layerTree.ReleaseSnapshotParams,
) error {
	command := &protocol.Command{
		Method: "LayerTree.releaseSnapshot",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ReplaySnapshot replays the layer snapshot and returns the resulting bitmap.
*/
func (LayerTree) ReplaySnapshot(
	socket *Socket,
	params *layerTree.ReplaySnapshotParams,
) (layerTree.ReplaySnapshotResult, error) {
	command := &protocol.Command{
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
*/
func (LayerTree) SnapshotCommandLog(
	socket *Socket,
	params *layerTree.SnapshotCommandLogParams,
) (layerTree.SnapshotCommandLogResult, error) {
	command := &protocol.Command{
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
*/
func (LayerTree) OnLayerPainted(
	socket *Socket,
	callback func(event *layerTree.LayerPaintedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (LayerTree) OnLayerTreeDidChange(
	socket *Socket,
	callback func(event *layerTree.DidChangeEvent),
) {
	handler := protocol.NewEventHandler(
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
