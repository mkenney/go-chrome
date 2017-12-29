package protocol

import (
	"encoding/json"

	serviceWorker "github.com/mkenney/go-chrome/protocol/service_worker"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
ServiceWorker provides a namespace for the Chrome ServiceWorker protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/ EXPERIMENTAL.
*/
var ServiceWorker = ServiceWorkerProtocol{}

/*
ServiceWorkerProtocol defines the ServiceWorker protocol methods.
*/
type ServiceWorkerProtocol struct{}

/*
DeliverPushMessage is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) DeliverPushMessage(
	socket sock.Socketer,
	params *serviceWorker.DeliverPushMessageParams,
) error {
	command := sock.NewCommand("ServiceWorker.deliverPushMessage", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ServiceWorker.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
DispatchSyncEvent is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) DispatchSyncEvent(
	socket sock.Socketer,
	params *serviceWorker.DispatchSyncEventParams,
) error {
	command := sock.NewCommand("ServiceWorker.dispatchSyncEvent", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ServiceWorker.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
InspectWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) InspectWorker(
	socket sock.Socketer,
	params *serviceWorker.InspectWorkerParams,
) error {
	command := sock.NewCommand("ServiceWorker.inspectWorker", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetForceUpdateOnPageLoad is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) SetForceUpdateOnPageLoad(
	socket sock.Socketer,
	params *serviceWorker.SetForceUpdateOnPageLoadParams,
) error {
	command := sock.NewCommand("ServiceWorker.setForceUpdateOnPageLoad", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SkipWaiting is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) SkipWaiting(
	socket sock.Socketer,
	params *serviceWorker.SkipWaitingParams,
) error {
	command := sock.NewCommand("ServiceWorker.skipWaiting", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StartWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) StartWorker(
	socket sock.Socketer,
	params *serviceWorker.StartWorkerParams,
) error {
	command := sock.NewCommand("ServiceWorker.startWorker", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopAllWorkers is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) StopAllWorkers(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ServiceWorker.stopAllWorkers", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopWorker is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) StopWorker(
	socket sock.Socketer,
	params *serviceWorker.StopWorkerParams,
) error {
	command := sock.NewCommand("ServiceWorker.stopWorker", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Unregister is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) Unregister(
	socket sock.Socketer,
	params *serviceWorker.UnregisterParams,
) error {
	command := sock.NewCommand("ServiceWorker.unregister", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
UpdateRegistration is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) UpdateRegistration(
	socket sock.Socketer,
	params *serviceWorker.UpdateRegistrationParams,
) error {
	command := sock.NewCommand("ServiceWorker.updateRegistration", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnWorkerErrorReported is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) OnWorkerErrorReported(
	socket sock.Socketer,
	callback func(event *serviceWorker.WorkerErrorReportedEvent),
) {
	handler := sock.NewEventHandler(
		"ServiceWorker.workerErrorReported",
		func(response *sock.Response) {
			event := &serviceWorker.WorkerErrorReportedEvent{}
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
OnWorkerRegistrationUpdated is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) OnWorkerRegistrationUpdated(
	socket sock.Socketer,
	callback func(event *serviceWorker.WorkerRegistrationUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ServiceWorker.workerRegistrationUpdated",
		func(response *sock.Response) {
			event := &serviceWorker.WorkerRegistrationUpdatedEvent{}
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
OnWorkerVersionUpdated is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
EXPERIMENTAL.
*/
func (ServiceWorkerProtocol) OnWorkerVersionUpdated(
	socket sock.Socketer,
	callback func(event *serviceWorker.WorkerVersionUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ServiceWorker.workerVersionUpdated",
		func(response *sock.Response) {
			event := &serviceWorker.WorkerVersionUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
