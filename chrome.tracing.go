package chrome

import (
	"app/chrome/protocol"
	tracing "app/chrome/tracing"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Tracing - https://chromedevtools.github.io/devtools-protocol/tot/Tracing/
EXPERIMENTAL
*/
type Tracing struct{}

/*
End stops trace events collection.
*/
func (Tracing) End(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Tracing.end",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetCategories gets supported tracing categories.
*/
func (Tracing) GetCategories(
	socket *Socket,
) (tracing.GetCategoriesResult, error) {
	command := &protocol.Command{
		Method: "Tracing.getCategories",
	}
	result := tracing.GetCategoriesResult{}
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
RecordClockSyncMarker records a clock sync marker in the trace.
*/
func (Tracing) RecordClockSyncMarker(
	socket *Socket,
	params *tracing.RecordClockSyncMarkerParams,
) error {
	command := &protocol.Command{
		Method: "Tracing.recordClockSyncMarker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestMemoryDump requests a global memory dump.
*/
func (Tracing) RequestMemoryDump(
	socket *Socket,
) (tracing.GetCategoriesResult, error) {
	command := &protocol.Command{
		Method: "Tracing.requestMemoryDump",
	}
	result := tracing.GetCategoriesResult{}
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
Start starts trace events collection.
*/
func (Tracing) Start(
	socket *Socket,
	params *tracing.StartParams,
) error {
	command := &protocol.Command{
		Method: "Tracing.start",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnBufferUsage adds a handler to the Tracing.bufferUsage event. Tracing.bufferUsage fires when a
buffer is used.
*/
func (Tracing) OnBufferUsage(
	socket *Socket,
	callback func(event *tracing.BufferUsageEvent),
) {
	handler := protocol.NewEventHandler(
		"Tracing.bufferUsage",
		func(name string, params []byte) {
			event := &tracing.BufferUsageEvent{}
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
OnDataCollected adds a handler to the Tracing.dataCollected event. Tracing.dataCollected fires when
tracing is stopped, collected events will be sent as a sequence of dataCollected events followed by
tracingComplete event. Contains an bucket of collected trace events.
*/
func (Tracing) OnDataCollected(
	socket *Socket,
	callback func(event *tracing.DataCollectedEvent),
) {
	handler := protocol.NewEventHandler(
		"Tracing.dataCollected",
		func(name string, params []byte) {
			event := &tracing.DataCollectedEvent{}
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
OnTracingComplete adds a handler to the Tracing.tracingComplete event. Tracing.tracingComplete fires
tracing is stopped and there is no trace buffers pending flush, all data were delivered via
dataCollected events.
*/
func (Tracing) OnTracingComplete(
	socket *Socket,
	callback func(event *tracing.TracingCompleteEvent),
) {
	handler := protocol.NewEventHandler(
		"Tracing.tracingComplete",
		func(name string, params []byte) {
			event := &tracing.TracingCompleteEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
