package chrome

import (
	heap "app/chrome/heap_profiler"
	"app/chrome/protocol"
)

/*
HeapProfiler EXPERIMENTAL
*/
type HeapProfiler struct{}

/*
Enable enables the HeapProfiler.
*/
func (HeapProfiler) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables the HeapProfiler.
*/
func (HeapProfiler) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
StartTrackingHeapObjects EXPERIMENTAL
*/
func (HeapProfiler) StartTrackingHeapObjects(socket *Socket, params *heap.StartTrackingHeapObjectsParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.startTrackingHeapObjects"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
StopTrackingHeapObjects EXPERIMENTAL
*/
func (HeapProfiler) StopTrackingHeapObjects(socket *Socket, params *heap.StopTrackingHeapObjectsParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.stopTrackingHeapObjects"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
TakeHeapSnapshot EXPERIMENTAL
*/
func (HeapProfiler) TakeHeapSnapshot(socket *Socket, params *heap.TakeHeapSnapshotParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.takeHeapSnapshot"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
CollectGarbage EXPERIMENTAL
*/
func (HeapProfiler) CollectGarbage(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.collectGarbage"
	socket.SendCommand(command)
	return command.Err
}

/*
GetObjectByHeapObjectID EXPERIMENTAL
*/
func (HeapProfiler) GetObjectByHeapObjectID(socket *Socket, params *heap.GetObjectByHeapObjectIDParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.getObjectByHeapObjectId"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
AddInspectedHeapObject enables console to refer to the node with given id via $x (see Command Line
API for more details $x functions).
*/
func (HeapProfiler) AddInspectedHeapObject(socket *Socket, params *heap.AddInspectedHeapObjectParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.addInspectedHeapObject"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
GetHeapObjectID EXPERIMENTAL
*/
func (HeapProfiler) GetHeapObjectID(socket *Socket, params *heap.GetHeapObjectIDParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.getHeapObjectID"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
StartSampling EXPERIMENTAL
*/
func (HeapProfiler) StartSampling(socket *Socket, params *heap.StartSamplingParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.startSampling"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
StopSampling EXPERIMENTAL
*/
func (HeapProfiler) StopSampling(socket *Socket, params *heap.StopSamplingParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.stopSampling"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
GetSamplingProfile EXPERIMENTAL
*/
func (HeapProfiler) GetSamplingProfile(socket *Socket, params *heap.GetSamplingProfileParams) error {
	command := new(protocol.Command)
	command.method = "HeapProfiler.getSamplingProfile"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
OnOnAddHeapSnapshotChunk adds a handler to the HeapProfiler.OnAddHeapSnapshotChunk event.
EXPERIMENTAL
*/
func (HeapProfiler) OnAddHeapSnapshotChunk(socket *Socket, callback func(event *heap.AddHeapSnapshotChunkEvent)) error {
	handler := protocol.NewEventHandler(
		"HeapProfiler.addHeapSnapshotChunk",
		func(name string, params []byte) {
			event := new(heap.AddHeapSnapshotChunkEvent)
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
OnResetProfiles adds a handler to the HeapProfiler.ResetProfiles event. EXPERIMENTAL
*/
func (HeapProfiler) OnResetProfiles(socket *Socket, callback func(event *heap.ResetProfilesEvent)) error {
	handler := protocol.NewEventHandler(
		"HeapProfiler.resetProfiles",
		func(name string, params []byte) {
			event := new(heap.ResetProfilesEvent)
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
OnReportHeapSnapshotProgress adds a handler to the DOM.ReportHeapSnapshotProgress event.
EXPERIMENTAL
*/
func (HeapProfiler) OnReportHeapSnapshotProgress(socket *Socket, callback func(event *heap.ReportHeapSnapshotProgressEvent)) error {
	handler := protocol.NewEventHandler(
		"HeapProfiler.reportHeapSnapshotProgress",
		func(name string, params []byte) {
			event := new(heap.ReportHeapSnapshotProgressEvent)
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
OnLastSeenObjectID adds a handler to the DOM.LastSeenObjectID event. DOM.LastSeenObjectID fires if
heap objects tracking has been started then backend regularly sends a current value for last seen
object id and corresponding timestamp. If the were changes in the heap since last event then one or
more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
*/
func (HeapProfiler) OnLastSeenObjectID(socket *Socket, callback func(event *heap.LastSeenObjectIDEvent)) error {
	handler := protocol.NewEventHandler(
		"HeapProfiler.lastSeenObjectID",
		func(name string, params []byte) {
			event := new(heap.LastSeenObjectIDEvent)
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
OnHeapStatsUpdate adds a handler to the DOM.heapStatsUpdate event. DOM.heapStatsUpdate fires if heap
objects tracking has been started then backend may send update for one or more fragments.
*/
func (HeapProfiler) OnHeapStatsUpdate(socket *Socket, callback func(event *heap.HeapStatsUpdateEvent)) error {
	handler := protocol.NewEventHandler(
		"HeapProfiler.heapStatsUpdate",
		func(name string, params []byte) {
			event := new(heap.HeapStatsUpdateEvent)
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
