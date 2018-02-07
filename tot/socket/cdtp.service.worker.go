package socket

import (
	"encoding/json"

	serviceWorker "github.com/mkenney/go-chrome/tot/cdtp/service/worker"
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
	params *serviceWorker.DeliverPushMessageParams,
) chan *serviceWorker.DeliverPushMessageResult {
	resultChan := make(chan *serviceWorker.DeliverPushMessageResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.deliverPushMessage", params)
	result := &serviceWorker.DeliverPushMessageResult{}

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
Disable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Disable() chan *serviceWorker.DisableResult {
	resultChan := make(chan *serviceWorker.DisableResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.disable", nil)
	result := &serviceWorker.DisableResult{}

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
DispatchSyncEvent is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) DispatchSyncEvent(
	params *serviceWorker.DispatchSyncEventParams,
) chan *serviceWorker.DispatchSyncEventResult {
	resultChan := make(chan *serviceWorker.DispatchSyncEventResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.dispatchSyncEvent", params)
	result := &serviceWorker.DispatchSyncEventResult{}

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
Enable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Enable() chan *serviceWorker.EnableResult {
	resultChan := make(chan *serviceWorker.EnableResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.enable", nil)
	result := &serviceWorker.EnableResult{}

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
InspectWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) InspectWorker(
	params *serviceWorker.InspectWorkerParams,
) chan *serviceWorker.InspectWorkerResult {
	resultChan := make(chan *serviceWorker.InspectWorkerResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.inspectWorker", params)
	result := &serviceWorker.InspectWorkerResult{}

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
SetForceUpdateOnPageLoad is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) SetForceUpdateOnPageLoad(
	params *serviceWorker.SetForceUpdateOnPageLoadParams,
) chan *serviceWorker.SetForceUpdateOnPageLoadResult {
	resultChan := make(chan *serviceWorker.SetForceUpdateOnPageLoadResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.setForceUpdateOnPageLoad", params)
	result := &serviceWorker.SetForceUpdateOnPageLoadResult{}

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
SkipWaiting is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) SkipWaiting(
	params *serviceWorker.SkipWaitingParams,
) chan *serviceWorker.SkipWaitingResult {
	resultChan := make(chan *serviceWorker.SkipWaitingResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.skipWaiting", params)
	result := &serviceWorker.SkipWaitingResult{}

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
StartWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StartWorker(
	params *serviceWorker.StartWorkerParams,
) chan *serviceWorker.StartWorkerResult {
	resultChan := make(chan *serviceWorker.StartWorkerResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.startWorker", params)
	result := &serviceWorker.StartWorkerResult{}

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
StopAllWorkers is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StopAllWorkers() chan *serviceWorker.StopAllWorkersResult {
	resultChan := make(chan *serviceWorker.StopAllWorkersResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.stopAllWorkers", nil)
	result := &serviceWorker.StopAllWorkersResult{}

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
StopWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StopWorker(
	params *serviceWorker.StopWorkerParams,
) chan *serviceWorker.StopWorkerResult {
	resultChan := make(chan *serviceWorker.StopWorkerResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.stopWorker", params)
	result := &serviceWorker.StopWorkerResult{}

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
Unregister is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Unregister(
	params *serviceWorker.UnregisterParams,
) chan *serviceWorker.UnregisterResult {
	resultChan := make(chan *serviceWorker.UnregisterResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.unregister", params)
	result := &serviceWorker.UnregisterResult{}

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
UpdateRegistration is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) UpdateRegistration(
	params *serviceWorker.UpdateRegistrationParams,
) chan *serviceWorker.UpdateRegistrationResult {
	resultChan := make(chan *serviceWorker.UpdateRegistrationResult)
	command := NewCommand(protocol.Socket, "ServiceWorker.updateRegistration", params)
	result := &serviceWorker.UpdateRegistrationResult{}

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
OnWorkerErrorReported is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) OnWorkerErrorReported(
	callback func(event *serviceWorker.ErrorReportedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerErrorReported",
		func(response *Response) {
			event := &serviceWorker.ErrorReportedEvent{}
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
	callback func(event *serviceWorker.RegistrationUpdatedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerRegistrationUpdated",
		func(response *Response) {
			event := &serviceWorker.RegistrationUpdatedEvent{}
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
	callback func(event *serviceWorker.VersionUpdatedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerVersionUpdated",
		func(response *Response) {
			event := &serviceWorker.VersionUpdatedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
