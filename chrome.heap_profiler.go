package chrome

import (
	"encoding/json"

	heap_profiler "github.com/mkenney/go-chrome/heap_profiler"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/Sirupsen/logrus"
)

/*
HeapProfiler - https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/
EXPERIMENTAL
*/
type HeapProfiler struct{}

/*
AddInspectedHeapObject enables console to refer to the node with given id via $x (see Command Line
API for more details $x functions).

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
*/
func (HeapProfiler) AddInspectedHeapObject(
	socket *Socket,
	params *heap_profiler.AddInspectedHeapObjectParams,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.addInspectedHeapObject",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CollectGarbage EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
*/
func (HeapProfiler) CollectGarbage(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.collectGarbage",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
*/
func (HeapProfiler) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
*/
func (HeapProfiler) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetHeapObjectID EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
*/
func (HeapProfiler) GetHeapObjectID(
	socket *Socket,
	params *heap_profiler.GetHeapObjectIDParams,
) (heap_profiler.GetHeapObjectIDResult, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.getHeapObjectID",
		Params: params,
	}
	result := heap_profiler.GetHeapObjectIDResult{}
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
GetObjectByHeapObjectID EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
*/
func (HeapProfiler) GetObjectByHeapObjectID(
	socket *Socket,
	params *heap_profiler.GetObjectByHeapObjectIDParams,
) (heap_profiler.GetObjectByHeapObjectIDResult, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.getObjectByHeapObjectId",
		Params: params,
	}
	result := heap_profiler.GetObjectByHeapObjectIDResult{}
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
GetSamplingProfile EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
*/
func (HeapProfiler) GetSamplingProfile(
	socket *Socket,
	params *heap_profiler.GetSamplingProfileParams,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.getSamplingProfile",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartSampling EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
*/
func (HeapProfiler) StartSampling(
	socket *Socket,
	params *heap_profiler.StartSamplingParams,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.startSampling",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartTrackingHeapObjects EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
*/
func (HeapProfiler) StartTrackingHeapObjects(
	socket *Socket,
	params *heap_profiler.StartTrackingHeapObjectsParams,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.startTrackingHeapObjects",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopSampling EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
*/
func (HeapProfiler) StopSampling(
	socket *Socket,
	params *heap_profiler.StopSamplingParams,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.stopSampling",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopTrackingHeapObjects EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
*/
func (HeapProfiler) StopTrackingHeapObjects(
	socket *Socket,
	params *heap_profiler.StopTrackingHeapObjectsParams,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.stopTrackingHeapObjects",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
TakeHeapSnapshot EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
*/
func (HeapProfiler) TakeHeapSnapshot(
	socket *Socket,
	params *heap_profiler.TakeHeapSnapshotParams,
) error {
	command := &protocol.Command{
		Method: "HeapProfiler.takeHeapSnapshot",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAddHeapSnapshotChunk adds a handler to the HeapProfiler.AddHeapSnapshotChunk event.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-addHeapSnapshotChunk
*/
func (HeapProfiler) OnAddHeapSnapshotChunk(
	socket *Socket,
	callback func(event *heap_profiler.AddHeapSnapshotChunkEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.addHeapSnapshotChunk",
		func(name string, params []byte) {
			event := &heap_profiler.AddHeapSnapshotChunkEvent{}
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
OnHeapStatsUpdate adds a handler to the DOM.heapStatsUpdate event. DOM.heapStatsUpdate fires if heap
objects tracking has been started then backend may send update for one or more fragments.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-heapStatsUpdate
*/
func (HeapProfiler) OnHeapStatsUpdate(
	socket *Socket,
	callback func(event *heap_profiler.HeapStatsUpdateEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.heapStatsUpdate",
		func(name string, params []byte) {
			event := &heap_profiler.HeapStatsUpdateEvent{}
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
OnLastSeenObjectID adds a handler to the DOM.LastSeenObjectID event. DOM.LastSeenObjectID fires if
heap objects tracking has been started then backend regularly sends a current value for last seen
object id and corresponding timestamp. If the were changes in the heap since last event then one or
more heapStatsUpdate events will be sent before a new lastSeenObjectId event.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-lastSeenObjectId
*/
func (HeapProfiler) OnLastSeenObjectID(
	socket *Socket,
	callback func(event *heap_profiler.LastSeenObjectIDEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.lastSeenObjectID",
		func(name string, params []byte) {
			event := &heap_profiler.LastSeenObjectIDEvent{}
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
OnReportHeapSnapshotProgress adds a handler to the DOM.ReportHeapSnapshotProgress event.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-reportHeapSnapshotProgress
*/
func (HeapProfiler) OnReportHeapSnapshotProgress(
	socket *Socket,
	callback func(event *heap_profiler.ReportHeapSnapshotProgressEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.reportHeapSnapshotProgress",
		func(name string, params []byte) {
			event := &heap_profiler.ReportHeapSnapshotProgressEvent{}
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
OnResetProfiles adds a handler to the HeapProfiler.ResetProfiles event. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-resetProfiles
*/
func (HeapProfiler) OnResetProfiles(
	socket *Socket,
	callback func(event *heap_profiler.ResetProfilesEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.resetProfiles",
		func(name string, params []byte) {
			event := &heap_profiler.ResetProfilesEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
