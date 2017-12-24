package chrome

import (
	"encoding/json"

	animation "github.com/mkenney/go-chrome/animation"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/sirupsen/logrus"
)

/*
Animation - https://chromedevtools.github.io/devtools-protocol/tot/Animation/
EXPERIMENTAL
*/
type Animation struct{}

/*
Disable animation domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-disable
*/
func (Animation) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Animation.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable animation domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-enable
*/
func (Animation) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Animation.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetCurrentTime returns the current time of the an animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
*/
func (Animation) GetCurrentTime(
	socket *Socket,
	params *animation.GetCurrentTimeParams,
) (animation.GetCurrentTimeResult, error) {
	command := &protocol.Command{
		Method: "Animation.getCurrentTime",
		Params: params,
	}
	result := animation.GetCurrentTimeResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetPlaybackRate gets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
*/
func (Animation) GetPlaybackRate(
	socket *Socket,
) (animation.GetPlaybackRateResult, error) {
	command := &protocol.Command{
		Method: "Animation.getPlaybackRate",
	}
	result := animation.GetPlaybackRateResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
ReleaseAnimations releases a set of animations to no longer be manipulated.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
*/
func (Animation) ReleaseAnimations(
	socket *Socket,
	params *animation.ReleaseAnimationsParams,
) error {
	command := &protocol.Command{
		Method: "Animation.releaseAnimations",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ResolveAnimation gets the remote object of the Animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
*/
func (Animation) ResolveAnimation(
	socket *Socket,
	params *animation.ResolveAnimationParams,
) (animation.ResolveAnimationResult, error) {
	command := &protocol.Command{
		Method: "Animation.resolveAnimation",
		Params: params,
	}
	result := animation.ResolveAnimationResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SeekAnimations seeks a set of animations to a particular time within each animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
*/
func (Animation) SeekAnimations(
	socket *Socket,
	params *animation.SeekAnimationsParams,
) error {
	command := &protocol.Command{
		Method: "Animation.seekAnimations",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPaused sets the paused state of a set of animations.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
*/
func (Animation) SetPaused(
	socket *Socket,
	params *animation.SetPausedParams,
) error {
	command := &protocol.Command{
		Method: "Animation.setPaused",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPlaybackRate sets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
*/
func (Animation) SetPlaybackRate(
	socket *Socket,
	params *animation.SetPlaybackRateParams,
) error {
	command := &protocol.Command{
		Method: "Animation.setPlaybackRate",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetTiming sets the timing of an animation node.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
*/
func (Animation) SetTiming(
	socket *Socket,
	params *animation.SetTimingParams,
) error {
	command := &protocol.Command{
		Method: "Animation.setTiming",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAnimationCanceled adds a handler to the Animation.Canceled event. Animation.Canceled fires when
when an animation has been cancelled.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
*/
func (Animation) OnAnimationCanceled(
	socket *Socket,
	callback func(event *animation.CanceledEvent),
) {
	handler := protocol.NewEventHandler(
		"Animation.animationCanceled",
		func(name string, params []byte) {
			event := &animation.CanceledEvent{}
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
OnAnimationCreated adds a handler to the Animation.Created event. Animation.Created fires for each
animation that has been created.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
*/
func (Animation) OnAnimationCreated(
	socket *Socket,
	callback func(event *animation.CreatedEvent),
) {
	handler := protocol.NewEventHandler(
		"Animation.animationCreated",
		func(name string, params []byte) {
			event := &animation.CreatedEvent{}
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
OnAnimationStarted adds a handler to the Animation.Started event. Animation.Started fires for each
animation that has been started.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
*/
func (Animation) OnAnimationStarted(
	socket *Socket,
	callback func(event *animation.StartedEvent),
) {
	handler := protocol.NewEventHandler(
		"Animation.animationStarted",
		func(name string, params []byte) {
			event := &animation.StartedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
