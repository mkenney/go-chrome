package socket

import (
	"encoding/json"

	target "github.com/mkenney/go-chrome/cdtp/target"
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
) chan *target.ActivateTargetResult {
	resultChan := make(chan *target.ActivateTargetResult)
	command := NewCommand(protocol.Socket, "Target.activateTarget", params)
	result := &target.ActivateTargetResult{}

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
AttachToTarget attaches to the target with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
*/
func (protocol *TargetProtocol) AttachToTarget(
	params *target.AttachToTargetParams,
) chan *target.AttachToTargetResult {
	resultChan := make(chan *target.AttachToTargetResult)
	command := NewCommand(protocol.Socket, "Target.attachToTarget", params)
	result := &target.AttachToTargetResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
CloseTarget closes the target. If the target is a page that gets closed too.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
*/
func (protocol *TargetProtocol) CloseTarget(
	params *target.CloseTargetParams,
) chan *target.CloseTargetResult {
	resultChan := make(chan *target.CloseTargetResult)
	command := NewCommand(protocol.Socket, "Target.closeTarget", params)
	result := &target.CloseTargetResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
CreateBrowserContext creates a new empty BrowserContext. Similar to an incognito
profile but you can have more than one.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createBrowserContext
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) CreateBrowserContext() chan *target.CreateBrowserContextResult {
	resultChan := make(chan *target.CreateBrowserContextResult)
	command := NewCommand(protocol.Socket, "Target.createBrowserContext", nil)
	result := &target.CreateBrowserContextResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
CreateTarget creates a new page.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
*/
func (protocol *TargetProtocol) CreateTarget(
	params *target.CreateTargetParams,
) chan *target.CreateTargetResult {
	resultChan := make(chan *target.CreateTargetResult)
	command := NewCommand(protocol.Socket, "Target.createTarget", params)
	result := &target.CreateTargetResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DetachFromTarget detaches session with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-detachFromTarget
*/
func (protocol *TargetProtocol) DetachFromTarget(
	params *target.DetachFromTargetParams,
) chan *target.DetachFromTargetResult {
	resultChan := make(chan *target.DetachFromTargetResult)
	command := NewCommand(protocol.Socket, "Target.detachFromTarget", params)
	result := &target.DetachFromTargetResult{}

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
DisposeBrowserContext deletes a BrowserContext, will fail of any open page uses
it.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) DisposeBrowserContext(
	params *target.DisposeBrowserContextParams,
) chan *target.DisposeBrowserContextResult {
	resultChan := make(chan *target.DisposeBrowserContextResult)
	command := NewCommand(protocol.Socket, "Target.disposeBrowserContext", params)
	result := &target.DisposeBrowserContextResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetTargetInfo returns information about a target.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) GetTargetInfo(
	params *target.GetTargetInfoParams,
) chan *target.GetTargetInfoResult {
	resultChan := make(chan *target.GetTargetInfoResult)
	command := NewCommand(protocol.Socket, "Target.getTargetInfo", params)
	result := &target.GetTargetInfoResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetTargets retrieves a list of available targets.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
*/
func (protocol *TargetProtocol) GetTargets(
	params *target.GetTargetsParams,
) chan *target.GetTargetsResult {
	resultChan := make(chan *target.GetTargetsResult)
	command := NewCommand(protocol.Socket, "Target.getTargets", params)
	result := &target.GetTargetsResult{}

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
SendMessageToTarget sends protocol message over session with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-sendMessageToTarget
*/
func (protocol *TargetProtocol) SendMessageToTarget(
	params *target.SendMessageToTargetParams,
) chan *target.SendMessageToTargetResult {
	resultChan := make(chan *target.SendMessageToTargetResult)
	command := NewCommand(protocol.Socket, "Target.sendMessageToTarget", params)
	result := &target.SendMessageToTargetResult{}

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
SetAttachToFrames is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAttachToFrames
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) SetAttachToFrames(
	params *target.SetAttachToFramesParams,
) chan *target.SetAttachToFramesResult {
	resultChan := make(chan *target.SetAttachToFramesResult)
	command := NewCommand(protocol.Socket, "Target.setAttachToFrames", params)
	result := &target.SetAttachToFramesResult{}

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
SetAutoAttach controls whether to automatically attach to new targets which are
considered to be related to this one. When turned on, attaches to all existing
related targets as well. When turned off, automatically detaches from all
currently attached targets.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAutoAttach EXPERIMENTAL.
*/
func (protocol *TargetProtocol) SetAutoAttach(
	params *target.SetAutoAttachParams,
) chan *target.SetAutoAttachResult {
	resultChan := make(chan *target.SetAutoAttachResult)
	command := NewCommand(protocol.Socket, "Target.setAutoAttach", params)
	result := &target.SetAutoAttachResult{}

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
SetDiscoverTargets controls whether to discover available targets and notify via
`targetCreated`, `targetInfoChanged` and `targetDestroyed` events.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
*/
func (protocol *TargetProtocol) SetDiscoverTargets(
	params *target.SetDiscoverTargetsParams,
) chan *target.SetDiscoverTargetsResult {
	resultChan := make(chan *target.SetDiscoverTargetsResult)
	command := NewCommand(protocol.Socket, "Target.setDiscoverTargets", params)
	result := &target.SetDiscoverTargetsResult{}

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
SetRemoteLocations enables target discovery for the specified locations, when
SetDiscoverTargets was set to `true`.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setRemoteLocations
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) SetRemoteLocations(
	params *target.SetRemoteLocationsParams,
) chan *target.SetRemoteLocationsResult {
	resultChan := make(chan *target.SetRemoteLocationsResult)
	command := NewCommand(protocol.Socket, "Target.setRemoteLocations", params)
	result := &target.SetRemoteLocationsResult{}

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
OnAttachedToTarget adds a handler to the Target.attachedToTarget event.
Target.attachedToTarget fires when attached to target because of auto-attach or
AttachToTarget command.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-attachedToTarget EXPERIMENTAL.
*/
func (protocol *TargetProtocol) OnAttachedToTarget(
	callback func(event *target.AttachedToTargetEvent),
) {
	handler := NewEventHandler(
		"Target.attachedToTarget",
		func(response *Response) {
			event := &target.AttachedToTargetEvent{}
			if err := json.Unmarshal([]byte(response.Result), event); err != nil {
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
Target.detachedFromTarget fires when detached from target for any reason
(including `DetachFromTarget` command). Can be issued multiple times per target
if multiple sessions have been attached to it.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-detachedFromTarget
EXPERIMENTAL.
*/
func (protocol *TargetProtocol) OnDetachedFromTarget(
	callback func(event *target.DetachedFromTargetEvent),
) {
	handler := NewEventHandler(
		"Target.detachedFromTarget",
		func(response *Response) {
			event := &target.DetachedFromTargetEvent{}
			if err := json.Unmarshal([]byte(response.Result), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnReceivedMessageFromTarget adds a handler to the Target.receivedMessageFromTarget
event. Target.receivedMessageFromTarget fires when a new protocol message
received from the session (as reported in `attachedToTarget` event).

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-receivedMessageFromTarget
*/
func (protocol *TargetProtocol) OnReceivedMessageFromTarget(
	callback func(event *target.ReceivedMessageFromTargetEvent),
) {
	handler := NewEventHandler(
		"Target.receivedMessageFromTarget",
		func(response *Response) {
			event := &target.ReceivedMessageFromTargetEvent{}
			if err := json.Unmarshal([]byte(response.Result), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnTargetCreated adds a handler to the Target.Created event. Target.Created fires
when a possible inspection target is created.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCreated
*/
func (protocol *TargetProtocol) OnTargetCreated(
	callback func(event *target.CreatedEvent),
) {
	handler := NewEventHandler(
		"Target.targetCreated",
		func(response *Response) {
			event := &target.CreatedEvent{}
			if err := json.Unmarshal([]byte(response.Result), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnTargetDestroyed adds a handler to the Target.Destroyed event. Target.Destroyed
fires when a target is destroyed.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetDestroyed
*/
func (protocol *TargetProtocol) OnTargetDestroyed(
	callback func(event *target.DestroyedEvent),
) {
	handler := NewEventHandler(
		"Target.targetDestroyed",
		func(response *Response) {
			event := &target.DestroyedEvent{}
			if err := json.Unmarshal([]byte(response.Result), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnTargetInfoChanged adds a handler to the Target.InfoChanged event. Target.InfoChanged
fires when some information about a target has changed. This only happens
between `targetCreated` and `targetDestroyed` events.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetInfoChanged
*/
func (protocol *TargetProtocol) OnTargetInfoChanged(
	callback func(event *target.InfoChangedEvent),
) {
	handler := NewEventHandler(
		"Target.targetInfoChanged",
		func(response *Response) {
			event := &target.InfoChangedEvent{}
			if err := json.Unmarshal([]byte(response.Result), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
