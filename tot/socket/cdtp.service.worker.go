package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/service/worker"
)

/*
ServiceWorkerProtocol provides a namespace for the Chrome ServiceWorker protocol
methods.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/
EXPERIMENTAL.
*/
type ServiceWorkerProtocol struct {
	Socket Socketer
}

/*
DeliverPushMessage is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) DeliverPushMessage(
	params *worker.DeliverPushMessageParams,
) <-chan *worker.DeliverPushMessageResult {
	resultChan := make(chan *worker.DeliverPushMessageResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.deliverPushMessage", params)
	result := &worker.DeliverPushMessageResult{}

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
Disable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Disable() <-chan *worker.DisableResult {
	resultChan := make(chan *worker.DisableResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.disable", nil)
	result := &worker.DisableResult{}

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
DispatchSyncEvent is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) DispatchSyncEvent(
	params *worker.DispatchSyncEventParams,
) <-chan *worker.DispatchSyncEventResult {
	resultChan := make(chan *worker.DispatchSyncEventResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.dispatchSyncEvent", params)
	result := &worker.DispatchSyncEventResult{}

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
Enable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Enable() <-chan *worker.EnableResult {
	resultChan := make(chan *worker.EnableResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.enable", nil)
	result := &worker.EnableResult{}

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
InspectWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) InspectWorker(
	params *worker.InspectWorkerParams,
) <-chan *worker.InspectWorkerResult {
	resultChan := make(chan *worker.InspectWorkerResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.inspectWorker", params)
	result := &worker.InspectWorkerResult{}

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
SetForceUpdateOnPageLoad is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) SetForceUpdateOnPageLoad(
	params *worker.SetForceUpdateOnPageLoadParams,
) <-chan *worker.SetForceUpdateOnPageLoadResult {
	resultChan := make(chan *worker.SetForceUpdateOnPageLoadResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.setForceUpdateOnPageLoad", params)
	result := &worker.SetForceUpdateOnPageLoadResult{}

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
SkipWaiting is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) SkipWaiting(
	params *worker.SkipWaitingParams,
) <-chan *worker.SkipWaitingResult {
	resultChan := make(chan *worker.SkipWaitingResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.skipWaiting", params)
	result := &worker.SkipWaitingResult{}

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
StartWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StartWorker(
	params *worker.StartWorkerParams,
) <-chan *worker.StartWorkerResult {
	resultChan := make(chan *worker.StartWorkerResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.startWorker", params)
	result := &worker.StartWorkerResult{}

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
StopAllWorkers is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StopAllWorkers() <-chan *worker.StopAllWorkersResult {
	resultChan := make(chan *worker.StopAllWorkersResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.stopAllWorkers", nil)
	result := &worker.StopAllWorkersResult{}

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
StopWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StopWorker(
	params *worker.StopWorkerParams,
) <-chan *worker.StopWorkerResult {
	resultChan := make(chan *worker.StopWorkerResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.stopWorker", params)
	result := &worker.StopWorkerResult{}

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
Unregister is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Unregister(
	params *worker.UnregisterParams,
) <-chan *worker.UnregisterResult {
	resultChan := make(chan *worker.UnregisterResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.unregister", params)
	result := &worker.UnregisterResult{}

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
UpdateRegistration is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) UpdateRegistration(
	params *worker.UpdateRegistrationParams,
) <-chan *worker.UpdateRegistrationResult {
	resultChan := make(chan *worker.UpdateRegistrationResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.updateRegistration", params)
	result := &worker.UpdateRegistrationResult{}

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
OnWorkerErrorReported is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) OnWorkerErrorReported(
	callback func(event *worker.ErrorReportedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerErrorReported",
		func(response *Response) {
			event := &worker.ErrorReportedEvent{}
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
OnWorkerRegistrationUpdated is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) OnWorkerRegistrationUpdated(
	callback func(event *worker.RegistrationUpdatedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerRegistrationUpdated",
		func(response *Response) {
			event := &worker.RegistrationUpdatedEvent{}
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
OnWorkerVersionUpdated is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) OnWorkerVersionUpdated(
	callback func(event *worker.VersionUpdatedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerVersionUpdated",
		func(response *Response) {
			event := &worker.VersionUpdatedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
