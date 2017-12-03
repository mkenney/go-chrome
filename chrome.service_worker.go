package chrome

import (
	"app/chrome/protocol"
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
func (ServiceWorker) DeliverPushMessage(socket *Socket, params *service_worker.DeliverPushMessageParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.deliverPushMessage",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable EXPERIMENTAL
*/
func (ServiceWorker) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "ServiceWorker.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DispatchSyncEvent EXPERIMENTAL
*/
func (ServiceWorker) DispatchSyncEvent(socket *Socket, params *service_worker.DispatchSyncEventParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.dispatchSyncEvent",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable EXPERIMENTAL
*/
func (ServiceWorker) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "ServiceWorker.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
InspectWorker EXPERIMENTAL
*/
func (ServiceWorker) InspectWorker(socket *Socket, params *service_worker.InspectWorkerParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.inspectWorker",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetForceUpdateOnPageLoad EXPERIMENTAL
*/
func (ServiceWorker) SetForceUpdateOnPageLoad(socket *Socket, params *service_worker.SetForceUpdateOnPageLoadParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.setForceUpdateOnPageLoad",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SkipWaiting EXPERIMENTAL
*/
func (ServiceWorker) SkipWaiting(socket *Socket, params *service_worker.SkipWaitingParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.skipWaiting",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartWorker EXPERIMENTAL
*/
func (ServiceWorker) StartWorker(socket *Socket, params *service_worker.StartWorkerParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.startWorker",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopAllWorkers EXPERIMENTAL
*/
func (ServiceWorker) StopAllWorkers(socket *Socket) error {
	command := &protocol.Command{
		method: "ServiceWorker.stopAllWorkers",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopWorker EXPERIMENTAL
*/
func (ServiceWorker) StopWorker(socket *Socket, params *service_worker.StopWorkerParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.stopWorker",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Unregister EXPERIMENTAL
*/
func (ServiceWorker) Unregister(socket *Socket, params *service_worker.UnregisterParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.unregister",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
UpdateRegistration EXPERIMENTAL
*/
func (ServiceWorker) UpdateRegistration(socket *Socket, params *service_worker.UpdateRegistrationParams) error {
	command := &protocol.Command{
		method: "ServiceWorker.updateRegistration",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnWorkerErrorReported EXPERIMENTAL
*/
func (ServiceWorker) OnWorkerErrorReported(socket *Socket, callback func(event *service_worker.WorkerErrorReportedEvent)) error {
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
	return command.Err
}

/*
OnWorkerRegistrationUpdated EXPERIMENTAL
*/
func (ServiceWorker) OnWorkerRegistrationUpdated(socket *Socket, callback func(event *service_worker.WorkerRegistrationUpdatedEvent)) error {
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
	return command.Err
}

/*
OnWorkerVersionUpdated EXPERIMENTAL
*/
func (ServiceWorker) OnWorkerVersionUpdated(socket *Socket, callback func(event *service_worker.WorkerVersionUpdatedEvent)) error {
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
	return command.Err
}
