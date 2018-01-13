package socket

import (
	"encoding/json"

	heapProfiler "github.com/mkenney/go-chrome/tot/cdtp/heap_profiler"
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
) chan *heapProfiler.AddInspectedHeapObjectResult {
	resultChan := make(chan *heapProfiler.AddInspectedHeapObjectResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.addInspectedHeapObject", params)
	result := &heapProfiler.AddInspectedHeapObjectResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
CollectGarbage is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) CollectGarbage() chan *heapProfiler.CollectGarbageResult {
	resultChan := make(chan *heapProfiler.CollectGarbageResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.collectGarbage", nil)
	result := &heapProfiler.CollectGarbageResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Disable disables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
*/
func (protocol *HeapProfilerProtocol) Disable() chan *heapProfiler.DisableResult {
	resultChan := make(chan *heapProfiler.DisableResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.disable", nil)
	result := &heapProfiler.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
*/
func (protocol *HeapProfilerProtocol) Enable() chan *heapProfiler.EnableResult {
	resultChan := make(chan *heapProfiler.EnableResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.enable", nil)
	result := &heapProfiler.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetHeapObjectID is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetHeapObjectID(
	params *heapProfiler.GetHeapObjectIDParams,
) chan *heapProfiler.GetHeapObjectIDResult {
	resultChan := make(chan *heapProfiler.GetHeapObjectIDResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.getHeapObjectID", params)
	result := &heapProfiler.GetHeapObjectIDResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetObjectByHeapObjectID is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetObjectByHeapObjectID(
	params *heapProfiler.GetObjectByHeapObjectIDParams,
) chan *heapProfiler.GetObjectByHeapObjectIDResult {
	resultChan := make(chan *heapProfiler.GetObjectByHeapObjectIDResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.getObjectByHeapObjectId", params)
	result := &heapProfiler.GetObjectByHeapObjectIDResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetSamplingProfile is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetSamplingProfile(
	params *heapProfiler.GetSamplingProfileParams,
) chan *heapProfiler.GetSamplingProfileResult {
	resultChan := make(chan *heapProfiler.GetSamplingProfileResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.getSamplingProfile", params)
	result := &heapProfiler.GetSamplingProfileResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StartSampling is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StartSampling(
	params *heapProfiler.StartSamplingParams,
) chan *heapProfiler.StartSamplingResult {
	resultChan := make(chan *heapProfiler.StartSamplingResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.startSampling", params)
	result := &heapProfiler.StartSamplingResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StartTrackingHeapObjects is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StartTrackingHeapObjects(
	params *heapProfiler.StartTrackingHeapObjectsParams,
) chan *heapProfiler.StartTrackingHeapObjectsResult {
	resultChan := make(chan *heapProfiler.StartTrackingHeapObjectsResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.startTrackingHeapObjects", params)
	result := &heapProfiler.StartTrackingHeapObjectsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StopSampling is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StopSampling(
	params *heapProfiler.StopSamplingParams,
) chan *heapProfiler.StopSamplingResult {
	resultChan := make(chan *heapProfiler.StopSamplingResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.stopSampling", params)
	result := &heapProfiler.StopSamplingResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StopTrackingHeapObjects is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StopTrackingHeapObjects(
	params *heapProfiler.StopTrackingHeapObjectsParams,
) chan *heapProfiler.StopTrackingHeapObjectsResult {
	resultChan := make(chan *heapProfiler.StopTrackingHeapObjectsResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.stopTrackingHeapObjects", params)
	result := &heapProfiler.StopTrackingHeapObjectsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
TakeHeapSnapshot is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) TakeHeapSnapshot(
	params *heapProfiler.TakeHeapSnapshotParams,
) chan *heapProfiler.TakeHeapSnapshotResult {
	resultChan := make(chan *heapProfiler.TakeHeapSnapshotResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.takeHeapSnapshot", params)
	result := &heapProfiler.TakeHeapSnapshotResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
