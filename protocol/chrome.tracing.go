package protocol

import (
	"encoding/json"

	tracing "github.com/mkenney/go-chrome/protocol/tracing"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Tracing is a struct that provides a namespace for the Chrome Tracing protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/
*/
type Tracing struct{}

/*
End stops trace events collection.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-end
*/
func (Tracing) End(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Tracing.end", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetCategories gets supported tracing categories.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-getCategories
*/
func (Tracing) GetCategories(
	socket sock.Socketer,
) (tracing.GetCategoriesResult, error) {
	command := sock.NewCommand("Tracing.getCategories", nil)
	result := tracing.GetCategoriesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
RecordClockSyncMarker records a clock sync marker in the trace.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-recordClockSyncMarker
*/
func (Tracing) RecordClockSyncMarker(
	socket sock.Socketer,
	params *tracing.RecordClockSyncMarkerParams,
) error {
	command := sock.NewCommand("Tracing.recordClockSyncMarker", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RequestMemoryDump requests a global memory dump.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-requestMemoryDump
*/
func (Tracing) RequestMemoryDump(
	socket sock.Socketer,
) (tracing.GetCategoriesResult, error) {
	command := sock.NewCommand("Tracing.requestMemoryDump", nil)
	result := tracing.GetCategoriesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Start starts trace events collection.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-start
*/
func (Tracing) Start(
	socket sock.Socketer,
	params *tracing.StartParams,
) error {
	command := sock.NewCommand("Tracing.start", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnBufferUsage adds a handler to the Tracing.bufferUsage event. Tracing.bufferUsage fires when a
buffer is used.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-bufferUsage
*/
func (Tracing) OnBufferUsage(
	socket sock.Socketer,
	callback func(event *tracing.BufferUsageEvent),
) {
	handler := sock.NewEventHandler(
		"Tracing.bufferUsage",
		func(response *sock.Response) {
			event := &tracing.BufferUsageEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnDataCollected adds a handler to the Tracing.dataCollected event. Tracing.dataCollected fires when
tracing is stopped, collected events will be sent as a sequence of dataCollected events followed by
tracingComplete event. Contains an bucket of collected trace events.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-dataCollected
*/
func (Tracing) OnDataCollected(
	socket sock.Socketer,
	callback func(event *tracing.DataCollectedEvent),
) {
	handler := sock.NewEventHandler(
		"Tracing.dataCollected",
		func(response *sock.Response) {
			event := &tracing.DataCollectedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnTracingComplete adds a handler to the Tracing.Complete event. Tracing.Complete fires when tracing
is stopped and there is no trace buffers pending flush, all data were delivered via DataCollected
events.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-tracingComplete
*/
func (Tracing) OnTracingComplete(
	socket sock.Socketer,
	callback func(event *tracing.CompleteEvent),
) {
	handler := sock.NewEventHandler(
		"Tracing.tracingComplete",
		func(response *sock.Response) {
			event := &tracing.CompleteEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
