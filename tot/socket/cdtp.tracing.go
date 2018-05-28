package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/tracing"
)

/*
TracingProtocol provides a namespace for the Chrome Tracing protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/ EXPERIMENTAL.
*/
type TracingProtocol struct {
	Socket Socketer
}

/*
End stops trace events collection.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-end
*/
func (protocol *TracingProtocol) End() <-chan *tracing.EndResult {
	resultChan := make(chan *tracing.EndResult)
	command := NewCommand(protocol.Socket, "Tracing.end", nil)
	result := &tracing.EndResult{}

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
GetCategories gets supported tracing categories.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-getCategories
*/
func (protocol *TracingProtocol) GetCategories() <-chan *tracing.GetCategoriesResult {
	resultChan := make(chan *tracing.GetCategoriesResult)
	command := NewCommand(protocol.Socket, "Tracing.getCategories", nil)
	result := &tracing.GetCategoriesResult{}

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
RecordClockSyncMarker records a clock sync marker in the trace.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-recordClockSyncMarker
*/
func (protocol *TracingProtocol) RecordClockSyncMarker(
	params *tracing.RecordClockSyncMarkerParams,
) <-chan *tracing.RecordClockSyncMarkerResult {
	resultChan := make(chan *tracing.RecordClockSyncMarkerResult)
	command := NewCommand(protocol.Socket, "Tracing.recordClockSyncMarker", params)
	result := &tracing.RecordClockSyncMarkerResult{}

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
RequestMemoryDump requests a global memory dump.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-requestMemoryDump
*/
func (protocol *TracingProtocol) RequestMemoryDump() <-chan *tracing.RequestMemoryDumpResult {
	resultChan := make(chan *tracing.RequestMemoryDumpResult)
	command := NewCommand(protocol.Socket, "Tracing.requestMemoryDump", nil)
	result := &tracing.RequestMemoryDumpResult{}

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
Start starts trace events collection.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-start
*/
func (protocol *TracingProtocol) Start(
	params *tracing.StartParams,
) <-chan *tracing.StartResult {
	resultChan := make(chan *tracing.StartResult)
	command := NewCommand(protocol.Socket, "Tracing.start", params)
	result := &tracing.StartResult{}

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
OnBufferUsage adds a handler to the Tracing.bufferUsage event. Tracing.bufferUsage
fires when a buffer is used.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-bufferUsage
*/
func (protocol *TracingProtocol) OnBufferUsage(
	callback func(event *tracing.BufferUsageEvent),
) {
	handler := NewEventHandler(
		"Tracing.bufferUsage",
		func(response *Response) {
			event := &tracing.BufferUsageEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnDataCollected adds a handler to the Tracing.dataCollected event. Tracing.dataCollected
fires when tracing is stopped, collected events will be sent as a sequence of
dataCollected events followed by tracingComplete event. Contains an bucket of
collected trace events.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-dataCollected
*/
func (protocol *TracingProtocol) OnDataCollected(
	callback func(event *tracing.DataCollectedEvent),
) {
	handler := NewEventHandler(
		"Tracing.dataCollected",
		func(response *Response) {
			event := &tracing.DataCollectedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnTracingComplete adds a handler to the Tracing.Complete event. Tracing.Complete
fires when tracing is stopped and there is no trace buffers pending flush, all
data were delivered via DataCollected events.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-tracingComplete
*/
func (protocol *TracingProtocol) OnTracingComplete(
	callback func(event *tracing.CompleteEvent),
) {
	handler := NewEventHandler(
		"Tracing.tracingComplete",
		func(response *Response) {
			event := &tracing.CompleteEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
