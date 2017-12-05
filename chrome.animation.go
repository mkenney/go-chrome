package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Animation - https://chromedevtools.github.io/devtools-protocol/tot/Animation/
EXPERIMENTAL
*/
type Animation struct{}

/*
Disable animation domain notifications.
*/
func (Animation) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Animation.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable animation domain notifications.
*/
func (Animation) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Animation.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetCurrentTime returns the current time of the an animation.
*/
func (Animation) GetCurrentTime(
	socket *Socket,
	params *animation.GetCurrentTimeParams,
) (animation.GetCurrentTimeResult, error) {
	command := &protocol.Command{
		method: "Animation.getCurrentTime",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(protocol.GetCurrentTimeResult), command.Err
}

/*
GetPlaybackRate gets the playback rate of the document timeline.
*/
func (Animation) GetPlaybackRate(
	socket *Socket,
) (animation.GetPlaybackRateResult, error) {
	command := &protocol.Command{
		method: "Animation.getPlaybackRate",
	}
	socket.SendCommand(command)
	return command.Result.(protocol.GetPlaybackRateResult), command.Err
}

/*
ReleaseAnimations releases a set of animations to no longer be manipulated.
*/
func (Animation) ReleaseAnimations(
	socket *Socket,
	params *animation.ReleaseAnimationsParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Animation.releaseAnimations",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
ResolveAnimation gets the remote object of the Animation.
*/
func (Animation) ResolveAnimation(
	socket *Socket,
	params *animation.ResolveAnimationParams,
) (animation.ResolveAnimationResult, error) {
	command := &protocol.Command{
		method: "Animation.resolveAnimation",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(protocol.ResolveAnimationResult), command.Err
}

/*
SeekAnimations seeks a set of animations to a particular time within each animation.
*/
func (Animation) SeekAnimations(
	socket *Socket,
	params *animation.SeekAnimationsParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Animation.seekAnimations",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetPaused sets the paused state of a set of animations.
*/
func (Animation) SetPaused(
	socket *Socket,
	params *animation.SetPausedParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Animation.setPaused",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetPlaybackRate sets the playback rate of the document timeline.
*/
func (Animation) SetPlaybackRate(
	socket *Socket,
	params *animation.SetPlaybackRateParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Animation.setPlaybackRate",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetTiming sets the timing of an animation node.
*/
func (Animation) SetTiming(
	socket *Socket,
	params *animation.SetTimingParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Animation.setTiming",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
OnAnimationCanceled adds a handler to the Animation.animationCanceled event.
Animation.animationCanceled fires when when an animation has been cancelled.
*/
func (Animation) OnAnimationCanceled(
	socket *Socket,
	callback func(event *animation.AnimationCanceledEvent),
) {
	handler := protocol.NewEventHandler(
		"Animation.animationCanceled",
		func(name string, params []byte) {
			event := &animation.AnimationCanceledEvent{}
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
OnAnimationCreated adds a handler to the Animation.animationCreated event.
Animation.animationCreated fires for each animation that has been created.
*/
func (Animation) OnAnimationCreated(
	socket *Socket,
	callback func(event *animation.AnimationCreatedEvent),
) {
	handler := protocol.NewEventHandler(
		"Animation.animationCreated",
		func(name string, params []byte) {
			event := &animation.AnimationCreatedEvent{}
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
OnAnimationStarted adds a handler to the Animation.animationStarted event.
Animation.animationStarted fires for each animation that has been started.
*/
func (Animation) OnAnimationStarted(
	socket *Socket,
	callback func(event *animation.AnimationStartedEvent),
) {
	handler := protocol.NewEventHandler(
		"Animation.animationStarted",
		func(name string, params []byte) {
			event := &animation.AnimationStartedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
