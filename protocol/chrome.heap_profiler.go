package protocol

import (
	"encoding/json"

	heapProfiler "github.com/mkenney/go-chrome/protocol/heap_profiler"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
HeapProfiler is a struct that provides a namespace for the Chrome HeapProfiler protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/
*/
type HeapProfiler struct{}

/*
AddInspectedHeapObject enables console to refer to the node with given id via $x (see Command Line
API for more details $x functions).

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
*/
func (HeapProfiler) AddInspectedHeapObject(
	socket *sock.Socket,
	params *heapProfiler.AddInspectedHeapObjectParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *heapProfiler.GetHeapObjectIDParams,
) (heapProfiler.GetHeapObjectIDResult, error) {
	command := &sock.Command{
		Method: "HeapProfiler.getHeapObjectID",
		Params: params,
	}
	result := heapProfiler.GetHeapObjectIDResult{}
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
	socket *sock.Socket,
	params *heapProfiler.GetObjectByHeapObjectIDParams,
) (heapProfiler.GetObjectByHeapObjectIDResult, error) {
	command := &sock.Command{
		Method: "HeapProfiler.getObjectByHeapObjectId",
		Params: params,
	}
	result := heapProfiler.GetObjectByHeapObjectIDResult{}
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
	socket *sock.Socket,
	params *heapProfiler.GetSamplingProfileParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *heapProfiler.StartSamplingParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *heapProfiler.StartTrackingHeapObjectsParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *heapProfiler.StopSamplingParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *heapProfiler.StopTrackingHeapObjectsParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *heapProfiler.TakeHeapSnapshotParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	callback func(event *heapProfiler.AddHeapSnapshotChunkEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.addHeapSnapshotChunk",
		func(name string, params []byte) {
			event := &heapProfiler.AddHeapSnapshotChunkEvent{}
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
	socket *sock.Socket,
	callback func(event *heapProfiler.HeapStatsUpdateEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.heapStatsUpdate",
		func(name string, params []byte) {
			event := &heapProfiler.HeapStatsUpdateEvent{}
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
	socket *sock.Socket,
	callback func(event *heapProfiler.LastSeenObjectIDEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.lastSeenObjectID",
		func(name string, params []byte) {
			event := &heapProfiler.LastSeenObjectIDEvent{}
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
	socket *sock.Socket,
	callback func(event *heapProfiler.ReportHeapSnapshotProgressEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.reportHeapSnapshotProgress",
		func(name string, params []byte) {
			event := &heapProfiler.ReportHeapSnapshotProgressEvent{}
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
	socket *sock.Socket,
	callback func(event *heapProfiler.ResetProfilesEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.resetProfiles",
		func(name string, params []byte) {
			event := &heapProfiler.ResetProfilesEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
