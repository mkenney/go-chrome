package socket

import (
	"encoding/json"

	heapProfiler "github.com/mkenney/go-chrome/cdtp/heap_profiler"
	log "github.com/sirupsen/logrus"
)

/*
HeapProfilerProtocol provides a namespace for the Chrome HeapProfiler protocol
methods.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/ EXPERIMENTAL.
*/
type HeapProfilerProtocol struct {
	Socket Socketer
}

/*
AddInspectedHeapObject enables console to refer to the node with given id via $x
(see Command Line API for more details $x functions).

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
*/
func (protocol *HeapProfilerProtocol) AddInspectedHeapObject(
	params *heapProfiler.AddInspectedHeapObjectParams,
) error {
	command := NewCommand("HeapProfiler.addInspectedHeapObject", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
CollectGarbage is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) CollectGarbage() error {
	command := NewCommand("HeapProfiler.collectGarbage", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
*/
func (protocol *HeapProfilerProtocol) Disable() error {
	command := NewCommand("HeapProfiler.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
*/
func (protocol *HeapProfilerProtocol) Enable() error {
	command := NewCommand("HeapProfiler.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetHeapObjectID is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetHeapObjectID(
	params *heapProfiler.GetHeapObjectIDParams,
) (*heapProfiler.GetHeapObjectIDResult, error) {
	command := NewCommand("HeapProfiler.getHeapObjectID", params)
	result := &heapProfiler.GetHeapObjectIDResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetObjectByHeapObjectID is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetObjectByHeapObjectID(
	params *heapProfiler.GetObjectByHeapObjectIDParams,
) (*heapProfiler.GetObjectByHeapObjectIDResult, error) {
	command := NewCommand("HeapProfiler.getObjectByHeapObjectId", params)
	result := &heapProfiler.GetObjectByHeapObjectIDResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetSamplingProfile is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetSamplingProfile(
	params *heapProfiler.GetSamplingProfileParams,
) error {
	command := NewCommand("HeapProfiler.getSamplingProfile", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StartSampling is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StartSampling(
	params *heapProfiler.StartSamplingParams,
) error {
	command := NewCommand("HeapProfiler.startSampling", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StartTrackingHeapObjects is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StartTrackingHeapObjects(
	params *heapProfiler.StartTrackingHeapObjectsParams,
) error {
	command := NewCommand("HeapProfiler.startTrackingHeapObjects", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopSampling is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StopSampling(
	params *heapProfiler.StopSamplingParams,
) error {
	command := NewCommand("HeapProfiler.stopSampling", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopTrackingHeapObjects is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StopTrackingHeapObjects(
	params *heapProfiler.StopTrackingHeapObjectsParams,
) error {
	command := NewCommand("HeapProfiler.stopTrackingHeapObjects", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
TakeHeapSnapshot is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) TakeHeapSnapshot(
	params *heapProfiler.TakeHeapSnapshotParams,
) error {
	command := NewCommand("HeapProfiler.takeHeapSnapshot", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnAddHeapSnapshotChunk adds a handler to the HeapProfiler.AddHeapSnapshotChunk
event.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-addHeapSnapshotChunk
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) OnAddHeapSnapshotChunk(
	callback func(event *heapProfiler.AddHeapSnapshotChunkEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.addHeapSnapshotChunk",
		func(response *Response) {
			event := &heapProfiler.AddHeapSnapshotChunkEvent{}
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
OnHeapStatsUpdate adds a handler to the DOM.heapStatsUpdate event. DOM.heapStatsUpdate
fires if heap objects tracking has been started then backend may send update for
one or more fragments.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-heapStatsUpdate
*/
func (protocol *HeapProfilerProtocol) OnHeapStatsUpdate(
	callback func(event *heapProfiler.HeapStatsUpdateEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.heapStatsUpdate",
		func(response *Response) {
			event := &heapProfiler.HeapStatsUpdateEvent{}
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
OnLastSeenObjectID adds a handler to the DOM.LastSeenObjectID event. DOM.LastSeenObjectID
fires if heap objects tracking has been started then backend regularly sends a
current value for last seen object id and corresponding timestamp. If the were
changes in the heap since last event then one or more heapStatsUpdate events
will be sent before a new lastSeenObjectId event.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-lastSeenObjectId
*/
func (protocol *HeapProfilerProtocol) OnLastSeenObjectID(
	callback func(event *heapProfiler.LastSeenObjectIDEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.lastSeenObjectID",
		func(response *Response) {
			event := &heapProfiler.LastSeenObjectIDEvent{}
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
OnReportHeapSnapshotProgress adds a handler to the DOM.ReportHeapSnapshotProgress
event.


https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-reportHeapSnapshotProgress
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) OnReportHeapSnapshotProgress(
	callback func(event *heapProfiler.ReportHeapSnapshotProgressEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.reportHeapSnapshotProgress",
		func(response *Response) {
			event := &heapProfiler.ReportHeapSnapshotProgressEvent{}
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
OnResetProfiles adds a handler to the HeapProfiler.ResetProfiles event.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-resetProfiles
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) OnResetProfiles(
	callback func(event *heapProfiler.ResetProfilesEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.resetProfiles",
		func(response *Response) {
			event := &heapProfiler.ResetProfilesEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
