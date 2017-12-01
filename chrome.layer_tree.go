package chrome

import (
	lt "app/chrome/layer_tree"
	"app/chrome/protocol"
)

/*
LayerTree domain
*/
type LayerTree struct{}

/*
CompositingReasons provides the reasons why the given layer was composited.
*/
func (LayerTree) CompositingReasons(socket *Socket, params *lt.CompositingReasonsParams) error {
	command := new(protocol.Command)
	command.method = "LayerTree.compositingReasons"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables compositing tree inspection.
*/
func (LayerTree) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "LayerTree.disable"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables compositing tree inspection.
*/
func (LayerTree) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "LayerTree.enable"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
LoadSnapshot returns the snapshot identifier.
*/
func (LayerTree) LoadSnapshot(socket *Socket, params *lt.LoadSnapshotParams) error {
	command := new(protocol.Command)
	command.method = "LayerTree.loadSnapshot"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
MakeSnapshot returns the layer snapshot identifier.
*/
func (LayerTree) MakeSnapshot(socket *Socket, params *lt.MakeSnapshotParams) error {
	command := new(protocol.Command)
	command.method = "LayerTree.makeSnapshot"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ProfileSnapshot profiles a snapshot.
*/
func (LayerTree) ProfileSnapshot(socket *Socket, params *lt.ProfileSnapshotParams) error {
	command := new(protocol.Command)
	command.method = "LayerTree.profileSnapshot"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ReleaseSnapshot releases layer snapshot captured by the back-end.
*/
func (LayerTree) ReleaseSnapshot(socket *Socket, params *lt.ReleaseSnapshotParams) error {
	command := new(protocol.Command)
	command.method = "LayerTree.releaseSnapshot"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ReplaySnapshot replays the layer snapshot and returns the resulting bitmap.
*/
func (LayerTree) ReplaySnapshot(socket *Socket, params *lt.ReplaySnapshotParams) error {
	command := new(protocol.Command)
	command.method = "LayerTree.replaySnapshot"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SnapshotCommandLog replays the layer snapshot and returns canvas log.
*/
func (LayerTree) SnapshotCommandLog(socket *Socket, params *lt.SnapshotCommandLogParams) error {
	command := new(protocol.Command)
	command.method = "LayerTree.snapshotCommandLog"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
OnLayerPainted adds a handler to the DOM.layerPainted event. DOM.layerPainted fires when the layer
is painted.
*/
func (LayerTree) OnLayerPainted(socket *Socket, callback func(event *lt.LayerPaintedEvent)) error {
	handler := protocol.NewEventHandler(
		"LayerTree.layerPainted",
		func(name string, params []byte) {
			event := new(heap.LayerPaintedEvent)
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		}
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnLayerTreeDidChange adds a handler to the DOM.layerTreeDidChange event. DOM.layerTreeDidChange
fires when the layer tree changes.
*/
func (LayerTree) OnLayerTreeDidChange(socket *Socket, callback func(event *lt.LayerTreeDidChangeEvent)) error {
	handler := protocol.NewEventHandler(
		"LayerTree.layerTreeDidChange",
		func(name string, params []byte) {
			event := new(heap.LayerTreeDidChangeEvent)
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		}
	)
	socket.AddEventHandler(handler)
	return command.Err
}
