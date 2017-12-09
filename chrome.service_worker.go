package chrome

import (
	"app/chrome/protocol"
	service_worker "app/chrome/service_worker"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
ServiceWorker - https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/
EXPERIMENTAL
*/
type ServiceWorker struct{}

/*
DeliverPushMessage EXPERIMENTAL
*/
func (ServiceWorker) DeliverPushMessage(
	socket *Socket,
	params *service_worker.DeliverPushMessageParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.deliverPushMessage",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable EXPERIMENTAL
*/
func (ServiceWorker) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DispatchSyncEvent EXPERIMENTAL
*/
func (ServiceWorker) DispatchSyncEvent(
	socket *Socket,
	params *service_worker.DispatchSyncEventParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.dispatchSyncEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable EXPERIMENTAL
*/
func (ServiceWorker) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
InspectWorker EXPERIMENTAL
*/
func (ServiceWorker) InspectWorker(
	socket *Socket,
	params *service_worker.InspectWorkerParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.inspectWorker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetForceUpdateOnPageLoad EXPERIMENTAL
*/
func (ServiceWorker) SetForceUpdateOnPageLoad(
	socket *Socket,
	params *service_worker.SetForceUpdateOnPageLoadParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.setForceUpdateOnPageLoad",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SkipWaiting EXPERIMENTAL
*/
func (ServiceWorker) SkipWaiting(
	socket *Socket,
	params *service_worker.SkipWaitingParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.skipWaiting",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartWorker EXPERIMENTAL
*/
func (ServiceWorker) StartWorker(
	socket *Socket,
	params *service_worker.StartWorkerParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.startWorker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopAllWorkers EXPERIMENTAL
*/
func (ServiceWorker) StopAllWorkers(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.stopAllWorkers",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopWorker EXPERIMENTAL
*/
func (ServiceWorker) StopWorker(
	socket *Socket,
	params *service_worker.StopWorkerParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.stopWorker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Unregister EXPERIMENTAL
*/
func (ServiceWorker) Unregister(
	socket *Socket,
	params *service_worker.UnregisterParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.unregister",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
UpdateRegistration EXPERIMENTAL
*/
func (ServiceWorker) UpdateRegistration(
	socket *Socket,
	params *service_worker.UpdateRegistrationParams,
) error {
	command := &protocol.Command{
		Method: "ServiceWorker.updateRegistration",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnWorkerErrorReported EXPERIMENTAL
*/
func (ServiceWorker) OnWorkerErrorReported(
	socket *Socket,
	callback func(event *service_worker.WorkerErrorReportedEvent),
) {
	handler := protocol.NewEventHandler(
		"ServiceWorker.workerErrorReported",
		func(name string, params []byte) {
			event := &service_worker.WorkerErrorReportedEvent{}
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
*/
func (ServiceWorker) OnWorkerRegistrationUpdated(
	socket *Socket,
	callback func(event *service_worker.WorkerRegistrationUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
		"ServiceWorker.workerRegistrationUpdated",
		func(name string, params []byte) {
			event := &service_worker.WorkerRegistrationUpdatedEvent{}
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
*/
func (ServiceWorker) OnWorkerVersionUpdated(
	socket *Socket,
	callback func(event *service_worker.WorkerVersionUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
		"ServiceWorker.workerVersionUpdated",
		func(name string, params []byte) {
			event := &service_worker.WorkerVersionUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
