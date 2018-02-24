package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/heap/profiler"
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
	params *profiler.AddInspectedHeapObjectParams,
) <-chan *profiler.AddInspectedHeapObjectResult {
	resultChan := make(chan *profiler.AddInspectedHeapObjectResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.addInspectedHeapObject", params)
	result := &profiler.AddInspectedHeapObjectResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
CollectGarbage is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) CollectGarbage() <-chan *profiler.CollectGarbageResult {
	resultChan := make(chan *profiler.CollectGarbageResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.collectGarbage", nil)
	result := &profiler.CollectGarbageResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
*/
func (protocol *HeapProfilerProtocol) Disable() <-chan *profiler.DisableResult {
	resultChan := make(chan *profiler.DisableResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.disable", nil)
	result := &profiler.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
*/
func (protocol *HeapProfilerProtocol) Enable() <-chan *profiler.EnableResult {
	resultChan := make(chan *profiler.EnableResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.enable", nil)
	result := &profiler.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetHeapObjectID is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetHeapObjectID(
	params *profiler.GetHeapObjectIDParams,
) <-chan *profiler.GetHeapObjectIDResult {
	resultChan := make(chan *profiler.GetHeapObjectIDResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.getHeapObjectID", params)
	result := &profiler.GetHeapObjectIDResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetObjectByHeapObjectID is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetObjectByHeapObjectID(
	params *profiler.GetObjectByHeapObjectIDParams,
) <-chan *profiler.GetObjectByHeapObjectIDResult {
	resultChan := make(chan *profiler.GetObjectByHeapObjectIDResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.getObjectByHeapObjectId", params)
	result := &profiler.GetObjectByHeapObjectIDResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetSamplingProfile is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) GetSamplingProfile(
	params *profiler.GetSamplingProfileParams,
) <-chan *profiler.GetSamplingProfileResult {
	resultChan := make(chan *profiler.GetSamplingProfileResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.getSamplingProfile", params)
	result := &profiler.GetSamplingProfileResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StartSampling is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StartSampling(
	params *profiler.StartSamplingParams,
) <-chan *profiler.StartSamplingResult {
	resultChan := make(chan *profiler.StartSamplingResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.startSampling", params)
	result := &profiler.StartSamplingResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StartTrackingHeapObjects is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StartTrackingHeapObjects(
	params *profiler.StartTrackingHeapObjectsParams,
) <-chan *profiler.StartTrackingHeapObjectsResult {
	resultChan := make(chan *profiler.StartTrackingHeapObjectsResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.startTrackingHeapObjects", params)
	result := &profiler.StartTrackingHeapObjectsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StopSampling is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StopSampling(
	params *profiler.StopSamplingParams,
) <-chan *profiler.StopSamplingResult {
	resultChan := make(chan *profiler.StopSamplingResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.stopSampling", params)
	result := &profiler.StopSamplingResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StopTrackingHeapObjects is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) StopTrackingHeapObjects(
	params *profiler.StopTrackingHeapObjectsParams,
) <-chan *profiler.StopTrackingHeapObjectsResult {
	resultChan := make(chan *profiler.StopTrackingHeapObjectsResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.stopTrackingHeapObjects", params)
	result := &profiler.StopTrackingHeapObjectsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
TakeHeapSnapshot is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
EXPERIMENTAL.
*/
func (protocol *HeapProfilerProtocol) TakeHeapSnapshot(
	params *profiler.TakeHeapSnapshotParams,
) <-chan *profiler.TakeHeapSnapshotResult {
	resultChan := make(chan *profiler.TakeHeapSnapshotResult)
	command := NewCommand(protocol.Socket, "HeapProfiler.takeHeapSnapshot", params)
	result := &profiler.TakeHeapSnapshotResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
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
	callback func(event *profiler.AddHeapSnapshotChunkEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.addHeapSnapshotChunk",
		func(response *Response) {
			event := &profiler.AddHeapSnapshotChunkEvent{}
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
	callback func(event *profiler.HeapStatsUpdateEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.heapStatsUpdate",
		func(response *Response) {
			event := &profiler.HeapStatsUpdateEvent{}
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
	callback func(event *profiler.LastSeenObjectIDEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.lastSeenObjectID",
		func(response *Response) {
			event := &profiler.LastSeenObjectIDEvent{}
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
	callback func(event *profiler.ReportHeapSnapshotProgressEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.reportHeapSnapshotProgress",
		func(response *Response) {
			event := &profiler.ReportHeapSnapshotProgressEvent{}
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
	callback func(event *profiler.ResetProfilesEvent),
) {
	handler := NewEventHandler(
		"HeapProfiler.resetProfiles",
		func(response *Response) {
			event := &profiler.ResetProfilesEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
