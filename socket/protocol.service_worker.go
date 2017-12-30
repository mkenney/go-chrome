package socket

import (
	"encoding/json"

	serviceWorker "github.com/mkenney/go-chrome/protocol/service_worker"
	log "github.com/sirupsen/logrus"
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
) error {
	command := NewCommand("ServiceWorker.deliverPushMessage", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Disable() error {
	command := NewCommand("ServiceWorker.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DispatchSyncEvent is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) DispatchSyncEvent(
	params *serviceWorker.DispatchSyncEventParams,
) error {
	command := NewCommand("ServiceWorker.dispatchSyncEvent", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Enable() error {
	command := NewCommand("ServiceWorker.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
InspectWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) InspectWorker(
	params *serviceWorker.InspectWorkerParams,
) error {
	command := NewCommand("ServiceWorker.inspectWorker", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetForceUpdateOnPageLoad is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) SetForceUpdateOnPageLoad(
	params *serviceWorker.SetForceUpdateOnPageLoadParams,
) error {
	command := NewCommand("ServiceWorker.setForceUpdateOnPageLoad", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SkipWaiting is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) SkipWaiting(
	params *serviceWorker.SkipWaitingParams,
) error {
	command := NewCommand("ServiceWorker.skipWaiting", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StartWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StartWorker(
	params *serviceWorker.StartWorkerParams,
) error {
	command := NewCommand("ServiceWorker.startWorker", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopAllWorkers is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StopAllWorkers() error {
	command := NewCommand("ServiceWorker.stopAllWorkers", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) StopWorker(
	params *serviceWorker.StopWorkerParams,
) error {
	command := NewCommand("ServiceWorker.stopWorker", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Unregister is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) Unregister(
	params *serviceWorker.UnregisterParams,
) error {
	command := NewCommand("ServiceWorker.unregister", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
UpdateRegistration is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) UpdateRegistration(
	params *serviceWorker.UpdateRegistrationParams,
) error {
	command := NewCommand("ServiceWorker.updateRegistration", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnWorkerErrorReported is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) OnWorkerErrorReported(
	callback func(event *serviceWorker.WorkerErrorReportedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerErrorReported",
		func(response *Response) {
			event := &serviceWorker.WorkerErrorReportedEvent{}
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
OnWorkerRegistrationUpdated is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) OnWorkerRegistrationUpdated(
	callback func(event *serviceWorker.WorkerRegistrationUpdatedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerRegistrationUpdated",
		func(response *Response) {
			event := &serviceWorker.WorkerRegistrationUpdatedEvent{}
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
OnWorkerVersionUpdated is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
EXPERIMENTAL.
*/
func (protocol *ServiceWorkerProtocol) OnWorkerVersionUpdated(
	callback func(event *serviceWorker.WorkerVersionUpdatedEvent),
) {
	handler := NewEventHandler(
		"ServiceWorker.workerVersionUpdated",
		func(response *Response) {
			event := &serviceWorker.WorkerVersionUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
