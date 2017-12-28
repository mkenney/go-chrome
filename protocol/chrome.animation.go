package protocol

import (
	"encoding/json"

	animation "github.com/mkenney/go-chrome/protocol/animation"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Animation is a struct that provides a namespace for the Chrome Animation protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Animation/
*/
var Animation = _animation{}

type _animation struct{}

/*
Disable animation domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-disable
*/
func (_animation) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Animation.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable animation domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-enable
*/
func (_animation) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Animation.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetCurrentTime returns the current time of the an animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
*/
func (_animation) GetCurrentTime(
	socket sock.Socketer,
	params *animation.GetCurrentTimeParams,
) (animation.GetCurrentTimeResult, error) {
	command := sock.NewCommand("Animation.getCurrentTime", params)
	result := animation.GetCurrentTimeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetPlaybackRate gets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
*/
func (_animation) GetPlaybackRate(
	socket sock.Socketer,
) (animation.GetPlaybackRateResult, error) {
	command := sock.NewCommand("Animation.getPlaybackRate", nil)
	result := animation.GetPlaybackRateResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
ReleaseAnimations releases a set of animations to no longer be manipulated.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
*/
func (_animation) ReleaseAnimations(
	socket sock.Socketer,
	params *animation.ReleaseAnimationsParams,
) error {
	command := sock.NewCommand("Animation.releaseAnimations", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
ResolveAnimation gets the remote object of the Animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
*/
func (_animation) ResolveAnimation(
	socket sock.Socketer,
	params *animation.ResolveAnimationParams,
) (animation.ResolveAnimationResult, error) {
	command := sock.NewCommand("Animation.resolveAnimation", params)
	result := animation.ResolveAnimationResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SeekAnimations seeks a set of animations to a particular time within each animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
*/
func (_animation) SeekAnimations(
	socket sock.Socketer,
	params *animation.SeekAnimationsParams,
) error {
	command := sock.NewCommand("Animation.seekAnimations", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetPaused sets the paused state of a set of animations.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
*/
func (_animation) SetPaused(
	socket sock.Socketer,
	params *animation.SetPausedParams,
) error {
	command := sock.NewCommand("Animation.setPaused", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetPlaybackRate sets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
*/
func (_animation) SetPlaybackRate(
	socket sock.Socketer,
	params *animation.SetPlaybackRateParams,
) error {
	command := sock.NewCommand("Animation.setPlaybackRate", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetTiming sets the timing of an animation node.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
*/
func (_animation) SetTiming(
	socket sock.Socketer,
	params *animation.SetTimingParams,
) error {
	command := sock.NewCommand("Animation.setTiming", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnAnimationCanceled adds a handler to the Animation.Canceled event. Animation.Canceled fires when
when an animation has been cancelled.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
*/
func (_animation) OnAnimationCanceled(
	socket sock.Socketer,
	callback func(event *animation.CanceledEvent),
) {
	handler := sock.NewEventHandler(
		"Animation.animationCanceled",
		func(response *sock.Response) {
			event := &animation.CanceledEvent{}
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
OnAnimationCreated adds a handler to the Animation.Created event. Animation.Created fires for each
animation that has been created.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
*/
func (_animation) OnAnimationCreated(
	socket sock.Socketer,
	callback func(event *animation.CreatedEvent),
) {
	handler := sock.NewEventHandler(
		"Animation.animationCreated",
		func(response *sock.Response) {
			event := &animation.CreatedEvent{}
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
OnAnimationStarted adds a handler to the Animation.Started event. Animation.Started fires for each
animation that has been started.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
*/
func (_animation) OnAnimationStarted(
	socket sock.Socketer,
	callback func(event *animation.StartedEvent),
) {
	handler := sock.NewEventHandler(
		"Animation.animationStarted",
		func(response *sock.Response) {
			event := &animation.StartedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
