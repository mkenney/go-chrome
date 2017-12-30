package socket

import (
	"encoding/json"

	animation "github.com/mkenney/go-chrome/protocol/animation"
	log "github.com/sirupsen/logrus"
)

/*
AnimationProtocol provides a namespace for the Chrome Animation protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/ EXPERIMENTAL.
*/
type AnimationProtocol struct {
	Socket Socketer
}

/*
Disable animation domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-disable
*/
func (protocol *AnimationProtocol) Disable() error {
	command := NewCommand("Animation.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable animation domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-enable
*/
func (protocol *AnimationProtocol) Enable() error {
	command := NewCommand("Animation.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetCurrentTime returns the current time of the an animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
*/
func (protocol *AnimationProtocol) GetCurrentTime(
	params *animation.GetCurrentTimeParams,
) (*animation.GetCurrentTimeResult, error) {
	command := NewCommand("Animation.getCurrentTime", params)
	result := &animation.GetCurrentTimeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetPlaybackRate gets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
*/
func (protocol *AnimationProtocol) GetPlaybackRate() (*animation.GetPlaybackRateResult, error) {
	command := NewCommand("Animation.getPlaybackRate", nil)
	result := &animation.GetPlaybackRateResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
ReleaseAnimations releases a set of animations to no longer be manipulated.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
*/
func (protocol *AnimationProtocol) ReleaseAnimations(
	params *animation.ReleaseAnimationsParams,
) error {
	command := NewCommand("Animation.releaseAnimations", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ResolveAnimation gets the remote object of the Animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
*/
func (protocol *AnimationProtocol) ResolveAnimation(
	params *animation.ResolveAnimationParams,
) (*animation.ResolveAnimationResult, error) {
	command := NewCommand("Animation.resolveAnimation", params)
	result := &animation.ResolveAnimationResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SeekAnimations seeks a set of animations to a particular time within each animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
*/
func (protocol *AnimationProtocol) SeekAnimations(
	params *animation.SeekAnimationsParams,
) error {
	command := NewCommand("Animation.seekAnimations", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetPaused sets the paused state of a set of animations.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
*/
func (protocol *AnimationProtocol) SetPaused(
	params *animation.SetPausedParams,
) error {
	command := NewCommand("Animation.setPaused", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetPlaybackRate sets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
*/
func (protocol *AnimationProtocol) SetPlaybackRate(
	params *animation.SetPlaybackRateParams,
) error {
	command := NewCommand("Animation.setPlaybackRate", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetTiming sets the timing of an animation node.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
*/
func (protocol *AnimationProtocol) SetTiming(
	params *animation.SetTimingParams,
) error {
	command := NewCommand("Animation.setTiming", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnAnimationCanceled adds a handler to the Animation.Canceled event. Animation.Canceled fires when
when an animation has been cancelled.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
*/
func (protocol *AnimationProtocol) OnAnimationCanceled(
	callback func(event *animation.CanceledEvent),
) {
	handler := NewEventHandler(
		"Animation.animationCanceled",
		func(response *Response) {
			event := &animation.CanceledEvent{}
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
OnAnimationCreated adds a handler to the Animation.Created event. Animation.Created fires for each
animation that has been created.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
*/
func (protocol *AnimationProtocol) OnAnimationCreated(
	callback func(event *animation.CreatedEvent),
) {
	handler := NewEventHandler(
		"Animation.animationCreated",
		func(response *Response) {
			event := &animation.CreatedEvent{}
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
OnAnimationStarted adds a handler to the Animation.Started event. Animation.Started fires for each
animation that has been started.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
*/
func (protocol *AnimationProtocol) OnAnimationStarted(
	callback func(event *animation.StartedEvent),
) {
	handler := NewEventHandler(
		"Animation.animationStarted",
		func(response *Response) {
			event := &animation.StartedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
