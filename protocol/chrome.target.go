package protocol

import (
	"encoding/json"

	target "github.com/mkenney/go-chrome/protocol/target"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Target provides a namespace for the Chrome Target protocol methods. The Target protocol supports
additional targets discovery and allows to attach to them.

https://chromedevtools.github.io/devtools-protocol/tot/Target/
*/
var Target = TargetProtocol{}

/*
TargetProtocol defines the Target protocol methods.
*/
type TargetProtocol struct{}

/*
ActivateTarget activates (focuses) the target.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
*/
func (TargetProtocol) ActivateTarget(
	socket sock.Socketer,
	params *target.ActivateTargetParams,
) error {
	command := sock.NewCommand("Target.activateTarget", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
AttachToTarget attaches to the target with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
*/
func (TargetProtocol) AttachToTarget(
	socket sock.Socketer,
	params *target.AttachToTargetParams,
) error {
	command := sock.NewCommand("Target.attachToTarget", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
CloseTarget closes the target. If the target is a page that gets closed too.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
*/
func (TargetProtocol) CloseTarget(
	socket sock.Socketer,
	params *target.CloseTargetParams,
) (*target.CloseTargetResult, error) {
	command := sock.NewCommand("Target.closeTarget", params)
	result := &target.CloseTargetResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
CreateBrowserContext creates a new empty BrowserContext. Similar to an incognito profile but you can
have more than one. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createBrowserContext
*/
func (TargetProtocol) CreateBrowserContext(
	socket sock.Socketer,
) (*target.CreateBrowserContextResult, error) {
	command := sock.NewCommand("Target.createBrowserContext", nil)
	result := &target.CreateBrowserContextResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
CreateTarget creates a new page.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
*/
func (TargetProtocol) CreateTarget(
	socket sock.Socketer,
	params *target.CreateTargetParams,
) (*target.CreateTargetResult, error) {
	command := sock.NewCommand("Target.createTarget", params)
	result := &target.CreateTargetResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
DetachFromTarget detaches session with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-detachFromTarget
*/
func (TargetProtocol) DetachFromTarget(
	socket sock.Socketer,
	params *target.DetachFromTargetParams,
) error {
	command := sock.NewCommand("Target.detachFromTarget", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DisposeBrowserContext deletes a BrowserContext, will fail of any open page uses it. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
*/
func (TargetProtocol) DisposeBrowserContext(
	socket sock.Socketer,
	params *target.DisposeBrowserContextParams,
) error {
	command := sock.NewCommand("Target.disposeBrowserContext", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetTargetInfo returns information about a target. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo
*/
func (TargetProtocol) GetTargetInfo(
	socket sock.Socketer,
	params *target.GetTargetInfoParams,
) (*target.GetTargetInfoResult, error) {
	command := sock.NewCommand("Target.getTargetInfo", params)
	result := &target.GetTargetInfoResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetTargets retrieves a list of available targets.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
*/
func (TargetProtocol) GetTargets(
	socket sock.Socketer,
	params *target.GetTargetsParams,
) error {
	command := sock.NewCommand("Target.getTargets", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SendMessageToTarget sends protocol message over session with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-sendMessageToTarget
*/
func (TargetProtocol) SendMessageToTarget(
	socket sock.Socketer,
	params *target.SendMessageToTargetParams,
) error {
	command := sock.NewCommand("Target.sendMessageToTarget", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetAttachToFrames EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAttachToFrames
*/
func (TargetProtocol) SetAttachToFrames(
	socket sock.Socketer,
	params *target.SetAttachToFramesParams,
) error {
	command := sock.NewCommand("Target.setAttachToFrames", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetAutoAttach controls whether to automatically attach to new targets which are considered to be
related to this one. When turned on, attaches to all existing related targets as well. When turned
off, automatically detaches from all currently attached targets. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAutoAttach
*/
func (TargetProtocol) SetAutoAttach(
	socket sock.Socketer,
	params *target.SetAutoAttachParams,
) error {
	command := sock.NewCommand("Target.setAutoAttach", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetDiscoverTargets controls whether to discover available targets and notify via
`targetCreated/targetInfoChanged/targetDestroyed` events.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
*/
func (TargetProtocol) SetDiscoverTargets(
	socket sock.Socketer,
	params *target.SetDiscoverTargetsParams,
) error {
	command := sock.NewCommand("Target.setDiscoverTargets", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetRemoteLocations enables target discovery for the specified locations, when `setDiscoverTargets`
was set to `true`. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setRemoteLocations
*/
func (TargetProtocol) SetRemoteLocations(
	socket sock.Socketer,
	params *target.SetRemoteLocationsParams,
) error {
	command := sock.NewCommand("Target.setRemoteLocations", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnAttachedToTarget adds a handler to the Target.attachedToTarget event.
Target.attachedToTarget fires when attached to target because of auto-attach or `attachToTarget`
command. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-attachedToTarget
*/
func (TargetProtocol) OnAttachedToTarget(
	socket sock.Socketer,
	callback func(event *target.AttachedToTargetEvent),
) {
	handler := sock.NewEventHandler(
		"Target.attachedToTarget",
		func(response *sock.Response) {
			event := &target.AttachedToTargetEvent{}
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
OnDetachedFromTarget adds a handler to the Target.detachedFromTarget event.
Target.detachedFromTarget fires when detached from target for any reason (including
`detachFromTarget` command). Can be issued multiple times per target if multiple sessions have been
attached to it. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-detachedFromTarget
*/
func (TargetProtocol) OnDetachedFromTarget(
	socket sock.Socketer,
	callback func(event *target.DetachedFromTargetEvent),
) {
	handler := sock.NewEventHandler(
		"Target.detachedFromTarget",
		func(response *sock.Response) {
			event := &target.DetachedFromTargetEvent{}
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
OnReceivedMessageFromTarget adds a handler to the Target.receivedMessageFromTarget event.
Target.receivedMessageFromTarget fires when a new protocol message received from the session (as
reported in `attachedToTarget` event).

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-receivedMessageFromTarget
*/
func (TargetProtocol) OnReceivedMessageFromTarget(
	socket sock.Socketer,
	callback func(event *target.ReceivedMessageFromTargetEvent),
) {
	handler := sock.NewEventHandler(
		"Target.receivedMessageFromTarget",
		func(response *sock.Response) {
			event := &target.ReceivedMessageFromTargetEvent{}
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
OnTargetCreated adds a handler to the Target.Created event. Target.Created fires when a possible
inspection target is created.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCreated
*/
func (TargetProtocol) OnTargetCreated(
	socket sock.Socketer,
	callback func(event *target.CreatedEvent),
) {
	handler := sock.NewEventHandler(
		"Target.targetCreated",
		func(response *sock.Response) {
			event := &target.CreatedEvent{}
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
OnTargetDestroyed adds a handler to the Target.Destroyed event. Target.Destroyed fires when a target
is destroyed.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetDestroyed
*/
func (TargetProtocol) OnTargetDestroyed(
	socket sock.Socketer,
	callback func(event *target.DestroyedEvent),
) {
	handler := sock.NewEventHandler(
		"Target.targetDestroyed",
		func(response *sock.Response) {
			event := &target.DestroyedEvent{}
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
OnTargetInfoChanged adds a handler to the Target.InfoChanged event. Target.InfoChanged fires when
some information about a target has changed. This only happens between `targetCreated` and
`targetDestroyed`.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetInfoChanged
*/
func (TargetProtocol) OnTargetInfoChanged(
	socket sock.Socketer,
	callback func(event *target.InfoChangedEvent),
) {
	handler := sock.NewEventHandler(
		"Target.targetInfoChanged",
		func(response *sock.Response) {
			event := &target.InfoChangedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
