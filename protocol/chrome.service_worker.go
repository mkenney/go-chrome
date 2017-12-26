package protocol

import (
	"encoding/json"

	serviceWorker "github.com/mkenney/go-chrome/protocol/service_worker"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
ServiceWorker is a struct that provides a namespace for the Chrome ServiceWorker protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/
*/
type ServiceWorker struct{}

/*
DeliverPushMessage EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
*/
func (ServiceWorker) DeliverPushMessage(
	socket *sock.Socket,
	params *serviceWorker.DeliverPushMessageParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.deliverPushMessage",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
*/
func (ServiceWorker) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DispatchSyncEvent EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
*/
func (ServiceWorker) DispatchSyncEvent(
	socket *sock.Socket,
	params *serviceWorker.DispatchSyncEventParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.dispatchSyncEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
*/
func (ServiceWorker) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
InspectWorker EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
*/
func (ServiceWorker) InspectWorker(
	socket *sock.Socket,
	params *serviceWorker.InspectWorkerParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.inspectWorker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetForceUpdateOnPageLoad EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
*/
func (ServiceWorker) SetForceUpdateOnPageLoad(
	socket *sock.Socket,
	params *serviceWorker.SetForceUpdateOnPageLoadParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.setForceUpdateOnPageLoad",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SkipWaiting EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
*/
func (ServiceWorker) SkipWaiting(
	socket *sock.Socket,
	params *serviceWorker.SkipWaitingParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.skipWaiting",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartWorker EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
*/
func (ServiceWorker) StartWorker(
	socket *sock.Socket,
	params *serviceWorker.StartWorkerParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.startWorker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopAllWorkers EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
*/
func (ServiceWorker) StopAllWorkers(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.stopAllWorkers",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopWorker EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
*/
func (ServiceWorker) StopWorker(
	socket *sock.Socket,
	params *serviceWorker.StopWorkerParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.stopWorker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Unregister EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
*/
func (ServiceWorker) Unregister(
	socket *sock.Socket,
	params *serviceWorker.UnregisterParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.unregister",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
UpdateRegistration EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
*/
func (ServiceWorker) UpdateRegistration(
	socket *sock.Socket,
	params *serviceWorker.UpdateRegistrationParams,
) error {
	command := &sock.Command{
		Method: "ServiceWorker.updateRegistration",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnWorkerErrorReported EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
*/
func (ServiceWorker) OnWorkerErrorReported(
	socket *sock.Socket,
	callback func(event *serviceWorker.WorkerErrorReportedEvent),
) {
	handler := sock.NewEventHandler(
		"ServiceWorker.workerErrorReported",
		func(name string, params []byte) {
			event := &serviceWorker.WorkerErrorReportedEvent{}
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
OnWorkerRegistrationUpdated EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
*/
func (ServiceWorker) OnWorkerRegistrationUpdated(
	socket *sock.Socket,
	callback func(event *serviceWorker.WorkerRegistrationUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ServiceWorker.workerRegistrationUpdated",
		func(name string, params []byte) {
			event := &serviceWorker.WorkerRegistrationUpdatedEvent{}
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
OnWorkerVersionUpdated EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
*/
func (ServiceWorker) OnWorkerVersionUpdated(
	socket *sock.Socket,
	callback func(event *serviceWorker.WorkerVersionUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ServiceWorker.workerVersionUpdated",
		func(name string, params []byte) {
			event := &serviceWorker.WorkerVersionUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
