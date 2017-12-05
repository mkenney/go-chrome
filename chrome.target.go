package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Target - https://chromedevtools.github.io/devtools-protocol/tot/Target/
Supports additional targets discovery and allows to attach to them.
*/
type Target struct{}

/*
ActivateTarget activates (focuses) the target.
*/
func (Target) ActivateTarget(socket *Socket, params *target.ActivateTargetParams) error {
	command := &protocol.Command{
		method: "Target.activateTarget",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
AttachToTarget attaches to the target with given id.
*/
func (Target) AttachToTarget(socket *Socket, params *target.AttachToTargetParams) error {
	command := &protocol.Command{
		method: "Target.attachToTarget",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CloseTarget closes the target. If the target is a page that gets closed too.
*/
func (Target) CloseTarget(socket *Socket, params *target.CloseTargetParams) error {
	command := &protocol.Command{
		method: "Target.closeTarget",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CreateBrowserContext creates a new empty BrowserContext. Similar to an incognito profile but you can
have more than one. EXPERIMENTAL
*/
func (Target) CreateBrowserContext(socket *Socket, params *target.CreateBrowserContextParams) error {
	command := &protocol.Command{
		method: "Target.createBrowserContext",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CreateTarget creates a new page.
*/
func (Target) CreateTarget(socket *Socket, params *target.CreateTargetParams) error {
	command := &protocol.Command{
		method: "Target.createTarget",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DetachFromTarget detaches session with given id.
*/
func (Target) DetachFromTarget(socket *Socket, params *target.DetachFromTargetParams) error {
	command := &protocol.Command{
		method: "Target.detachFromTarget",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DisposeBrowserContext deletes a BrowserContext, will fail of any open page uses it. EXPERIMENTAL
*/
func (Target) DisposeBrowserContext(socket *Socket, params *target.DisposeBrowserContextParams) error {
	command := &protocol.Command{
		method: "Target.disposeBrowserContext",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetTargetInfo returns information about a target. EXPERIMENTAL
*/
func (Target) GetTargetInfo(socket *Socket, params *target.GetTargetInfoParams) error {
	command := &protocol.Command{
		method: "Target.getTargetInfo",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetTargets retrieves a list of available targets.
*/
func (Target) GetTargets(socket *Socket, params *target.GetTargetsParams) error {
	command := &protocol.Command{
		method: "Target.getTargets",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SendMessageToTarget sends protocol message over session with given id.
*/
func (Target) SendMessageToTarget(socket *Socket, params *target.SendMessageToTargetParams) error {
	command := &protocol.Command{
		method: "Target.sendMessageToTarget",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAttachToFrames EXPERIMENTAL
*/
func (Target) SetAttachToFrames(socket *Socket, params *target.SetAttachToFramesParams) error {
	command := &protocol.Command{
		method: "Target.setAttachToFrames",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAutoAttach controls whether to automatically attach to new targets which are considered to be
related to this one. When turned on, attaches to all existing related targets as well. When turned
off, automatically detaches from all currently attached targets. EXPERIMENTAL
*/
func (Target) SetAutoAttach(socket *Socket, params *target.SetAutoAttachParams) error {
	command := &protocol.Command{
		method: "Target.setAutoAttach",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDiscoverTargets controls whether to discover available targets and notify via
`targetCreated/targetInfoChanged/targetDestroyed` events.
*/
func (Target) SetDiscoverTargets(socket *Socket, params *target.SetDiscoverTargetsParams) error {
	command := &protocol.Command{
		method: "Target.setDiscoverTargets",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetRemoteLocations enables target discovery for the specified locations, when `setDiscoverTargets`
was set to `true`. EXPERIMENTAL
*/
func (Target) SetRemoteLocations(socket *Socket, params *target.SetRemoteLocationsParams) error {
	command := &protocol.Command{
		method: "Target.setRemoteLocations",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnDetachedFromTarget adds a handler to the Target.detachedFromTarget event.
Target.detachedFromTarget fires when detached from target for any reason (including
`detachFromTarget` command). Can be issued multiple times per target if multiple sessions have been
attached to it. EXPERIMENTAL
*/
func (Target) OnDetachedFromTarget(socket *Socket, callback func(event *target.DetachedFromTargetEvent)) {
	handler := protocol.NewEventHandler(
		"Target.detachedFromTarget",
		func(name string, params []byte) {
			event := &target.DetachedFromTargetEvent{}
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
OnReceivedMessageFromTarget adds a handler to the Target.receivedMessageFromTarget event.
Target.receivedMessageFromTarget fires when a new protocol message received from the session (as
reported in `attachedToTarget` event).
*/
func (Target) OnReceivedMessageFromTarget(socket *Socket, callback func(event *target.ReceivedMessageFromTargetEvent)) {
	handler := protocol.NewEventHandler(
		"Target.receivedMessageFromTarget",
		func(name string, params []byte) {
			event := &target.ReceivedMessageFromTargetEvent{}
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
OnTargetCreated adds a handler to the Target.targetCreated event. Target.targetCreated fires when a
possible inspection target is created.
*/
func (Target) OnTargetCreated(socket *Socket, callback func(event *target.TargetCreatedEvent)) {
	handler := protocol.NewEventHandler(
		"Target.targetCreated",
		func(name string, params []byte) {
			event := &target.TargetCreatedEvent{}
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
OnTargetDestroyed adds a handler to the Target.targetDestroyed event. Target.targetDestroyed fires
when a a target is destroyed.
*/
func (Target) OnTargetDestroyed(socket *Socket, callback func(event *target.TargetDestroyedEvent)) {
	handler := protocol.NewEventHandler(
		"Target.targetDestroyed",
		func(name string, params []byte) {
			event := &target.TargetDestroyedEvent{}
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
OnTargetInfoChanged adds a handler to the Target.targetInfoChanged event. Target.targetInfoChanged
fires when some information about a target has changed. This only happens between `targetCreated`
and `targetDestroyed`.
*/
func (Target) OnTargetInfoChanged(socket *Socket, callback func(event *target.TargetInfoChangedEvent)) {
	handler := protocol.NewEventHandler(
		"Target.targetInfoChanged",
		func(name string, params []byte) {
			event := &target.TargetInfoChangedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
