package socket

import (
	"encoding/json"

	tracing "github.com/mkenney/go-chrome/protocol/tracing"
	log "github.com/sirupsen/logrus"
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
func (protocol *TracingProtocol) End() error {
	command := NewCommand("Tracing.end", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetCategories gets supported tracing categories.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-getCategories
*/
func (protocol *TracingProtocol) GetCategories() (*tracing.GetCategoriesResult, error) {
	command := NewCommand("Tracing.getCategories", nil)
	result := &tracing.GetCategoriesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
RecordClockSyncMarker records a clock sync marker in the trace.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-recordClockSyncMarker
*/
func (protocol *TracingProtocol) RecordClockSyncMarker(
	params *tracing.RecordClockSyncMarkerParams,
) error {
	command := NewCommand("Tracing.recordClockSyncMarker", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RequestMemoryDump requests a global memory dump.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-requestMemoryDump
*/
func (protocol *TracingProtocol) RequestMemoryDump() (*tracing.GetCategoriesResult, error) {
	command := NewCommand("Tracing.requestMemoryDump", nil)
	result := &tracing.GetCategoriesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Start starts trace events collection.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-start
*/
func (protocol *TracingProtocol) Start(
	params *tracing.StartParams,
) error {
	command := NewCommand("Tracing.start", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
