package socket

import (
	"encoding/json"

	target "github.com/mkenney/go-chrome/protocol/target"
	log "github.com/sirupsen/logrus"
)

/*
TargetProtocol provides a namespace for the Chrome Target protocol methods. The Target protocol
supports additional targets discovery and allows to attach to them.

https://chromedevtools.github.io/devtools-protocol/tot/Target/
*/
type TargetProtocol struct {
	Socket Socketer
}

/*
ActivateTarget activates (focuses) the target.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
*/
func (protocol *TargetProtocol) ActivateTarget(
	params *target.ActivateTargetParams,
) error {
	command := NewCommand("Target.activateTarget", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
AttachToTarget attaches to the target with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
*/
func (protocol *TargetProtocol) AttachToTarget(
	params *target.AttachToTargetParams,
) error {
	command := NewCommand("Target.attachToTarget", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
CloseTarget closes the target. If the target is a page that gets closed too.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
*/
func (protocol *TargetProtocol) CloseTarget(
	params *target.CloseTargetParams,
) (*target.CloseTargetResult, error) {
	command := NewCommand("Target.closeTarget", params)
	result := &target.CloseTargetResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
CreateBrowserContext creates a new empty BrowserContext. Similar to an incognito profile but you can
have more than one.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createBrowserContext
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) CreateBrowserContext() (*target.CreateBrowserContextResult, error) {
	command := NewCommand("Target.createBrowserContext", nil)
	result := &target.CreateBrowserContextResult{}
	protocol.Socket.SendCommand(command)

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
func (protocol *TargetProtocol) CreateTarget(
	params *target.CreateTargetParams,
) (*target.CreateTargetResult, error) {
	command := NewCommand("Target.createTarget", params)
	result := &target.CreateTargetResult{}
	protocol.Socket.SendCommand(command)

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
func (protocol *TargetProtocol) DetachFromTarget(
	params *target.DetachFromTargetParams,
) error {
	command := NewCommand("Target.detachFromTarget", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DisposeBrowserContext deletes a BrowserContext, will fail of any open page uses it.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) DisposeBrowserContext(
	params *target.DisposeBrowserContextParams,
) error {
	command := NewCommand("Target.disposeBrowserContext", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetTargetInfo returns information about a target.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo EXPERIMENTAL.
*/
func (protocol *TargetProtocol) GetTargetInfo(
	params *target.GetTargetInfoParams,
) (*target.GetTargetInfoResult, error) {
	command := NewCommand("Target.getTargetInfo", params)
	result := &target.GetTargetInfoResult{}
	protocol.Socket.SendCommand(command)

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
func (protocol *TargetProtocol) GetTargets(
	params *target.GetTargetsParams,
) error {
	command := NewCommand("Target.getTargets", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SendMessageToTarget sends protocol message over session with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-sendMessageToTarget
*/
func (protocol *TargetProtocol) SendMessageToTarget(
	params *target.SendMessageToTargetParams,
) error {
	command := NewCommand("Target.sendMessageToTarget", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetAttachToFrames is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAttachToFrames EXPERIMENTAL.
*/
func (protocol *TargetProtocol) SetAttachToFrames(
	params *target.SetAttachToFramesParams,
) error {
	command := NewCommand("Target.setAttachToFrames", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetAutoAttach controls whether to automatically attach to new targets which are considered to be
related to this one. When turned on, attaches to all existing related targets as well. When turned
off, automatically detaches from all currently attached targets.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAutoAttach EXPERIMENTAL.
*/
func (protocol *TargetProtocol) SetAutoAttach(
	params *target.SetAutoAttachParams,
) error {
	command := NewCommand("Target.setAutoAttach", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetDiscoverTargets controls whether to discover available targets and notify via
`targetCreated/targetInfoChanged/targetDestroyed` events.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
*/
func (protocol *TargetProtocol) SetDiscoverTargets(
	params *target.SetDiscoverTargetsParams,
) error {
	command := NewCommand("Target.setDiscoverTargets", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetRemoteLocations enables target discovery for the specified locations, when `setDiscoverTargets`
was set to `true`.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setRemoteLocations
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) SetRemoteLocations(
	params *target.SetRemoteLocationsParams,
) error {
	command := NewCommand("Target.setRemoteLocations", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnAttachedToTarget adds a handler to the Target.attachedToTarget event.
Target.attachedToTarget fires when attached to target because of auto-attach or `attachToTarget`
command.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-attachedToTarget EXPERIMENTAL.
*/
func (protocol *TargetProtocol) OnAttachedToTarget(
	callback func(event *target.AttachedToTargetEvent),
) {
	handler := NewEventHandler(
		"Target.attachedToTarget",
		func(response *Response) {
			event := &target.AttachedToTargetEvent{}
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
OnDetachedFromTarget adds a handler to the Target.detachedFromTarget event.
Target.detachedFromTarget fires when detached from target for any reason (including
`detachFromTarget` command). Can be issued multiple times per target if multiple sessions have been
attached to it.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-detachedFromTarget EXPERIMENTAL.
*/
func (protocol *TargetProtocol) OnDetachedFromTarget(
	callback func(event *target.DetachedFromTargetEvent),
) {
	handler := NewEventHandler(
		"Target.detachedFromTarget",
		func(response *Response) {
			event := &target.DetachedFromTargetEvent{}
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
OnReceivedMessageFromTarget adds a handler to the Target.receivedMessageFromTarget event.
Target.receivedMessageFromTarget fires when a new protocol message received from the session (as
reported in `attachedToTarget` event).

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-receivedMessageFromTarget
*/
func (protocol *TargetProtocol) OnReceivedMessageFromTarget(
	callback func(event *target.ReceivedMessageFromTargetEvent),
) {
	handler := NewEventHandler(
		"Target.receivedMessageFromTarget",
		func(response *Response) {
			event := &target.ReceivedMessageFromTargetEvent{}
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
OnTargetCreated adds a handler to the Target.Created event. Target.Created fires when a possible
inspection target is created.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCreated
*/
func (protocol *TargetProtocol) OnTargetCreated(
	callback func(event *target.CreatedEvent),
) {
	handler := NewEventHandler(
		"Target.targetCreated",
		func(response *Response) {
			event := &target.CreatedEvent{}
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
OnTargetDestroyed adds a handler to the Target.Destroyed event. Target.Destroyed fires when a target
is destroyed.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetDestroyed
*/
func (protocol *TargetProtocol) OnTargetDestroyed(
	callback func(event *target.DestroyedEvent),
) {
	handler := NewEventHandler(
		"Target.targetDestroyed",
		func(response *Response) {
			event := &target.DestroyedEvent{}
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
OnTargetInfoChanged adds a handler to the Target.InfoChanged event. Target.InfoChanged fires when
some information about a target has changed. This only happens between `targetCreated` and
`targetDestroyed`.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetInfoChanged
*/
func (protocol *TargetProtocol) OnTargetInfoChanged(
	callback func(event *target.InfoChangedEvent),
) {
	handler := NewEventHandler(
		"Target.targetInfoChanged",
		func(response *Response) {
			event := &target.InfoChangedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
