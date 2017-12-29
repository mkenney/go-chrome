package protocol

import (
	"encoding/json"

	serviceWorker "github.com/mkenney/go-chrome/protocol/service_worker"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
ServiceWorker provides a namespace for the Chrome ServiceWorker protocol methods. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/
*/
var ServiceWorker = ServiceWorkerProtocol{}

/*
ServiceWorkerProtocol defines the ServiceWorker protocol methods.
*/
type ServiceWorkerProtocol struct{}

/*
DeliverPushMessage EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
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
Disable EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
*/
func (ServiceWorkerProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ServiceWorker.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
DispatchSyncEvent EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
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
Enable EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
*/
func (ServiceWorkerProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ServiceWorker.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
InspectWorker EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
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
SetForceUpdateOnPageLoad EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
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
SkipWaiting EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
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
StartWorker EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
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
StopAllWorkers EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
*/
func (ServiceWorkerProtocol) StopAllWorkers(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ServiceWorker.stopAllWorkers", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopWorker EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
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
Unregister EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
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
UpdateRegistration EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
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
OnWorkerErrorReported EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
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
OnWorkerRegistrationUpdated EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
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
OnWorkerVersionUpdated EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
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
