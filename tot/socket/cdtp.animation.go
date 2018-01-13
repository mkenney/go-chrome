package socket

import (
	"encoding/json"

	animation "github.com/mkenney/go-chrome/tot/cdtp/animation"
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
func (protocol *AnimationProtocol) Disable() chan *animation.DisableResult {
	resultChan := make(chan *animation.DisableResult)
	command := NewCommand(protocol.Socket, "Animation.disable", nil)
	result := &animation.DisableResult{}

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
Enable animation domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-enable
*/
func (protocol *AnimationProtocol) Enable() chan *animation.EnableResult {
	resultChan := make(chan *animation.EnableResult)
	command := NewCommand(protocol.Socket, "Animation.enable", nil)
	result := &animation.EnableResult{}

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
GetCurrentTime returns the current time of the an animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
*/
func (protocol *AnimationProtocol) GetCurrentTime(
	params *animation.GetCurrentTimeParams,
) chan *animation.GetCurrentTimeResult {
	resultChan := make(chan *animation.GetCurrentTimeResult)
	command := NewCommand(protocol.Socket, "Animation.getCurrentTime", params)
	result := &animation.GetCurrentTimeResult{}

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
GetPlaybackRate gets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
*/
func (protocol *AnimationProtocol) GetPlaybackRate() chan *animation.GetPlaybackRateResult {
	resultChan := make(chan *animation.GetPlaybackRateResult)
	command := NewCommand(protocol.Socket, "Animation.getPlaybackRate", nil)
	result := &animation.GetPlaybackRateResult{}

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
ReleaseAnimations releases a set of animations to no longer be manipulated.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
*/
func (protocol *AnimationProtocol) ReleaseAnimations(
	params *animation.ReleaseAnimationsParams,
) chan *animation.ReleaseAnimationsResult {
	resultChan := make(chan *animation.ReleaseAnimationsResult)
	command := NewCommand(protocol.Socket, "Animation.releaseAnimations", params)
	result := &animation.ReleaseAnimationsResult{}

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
ResolveAnimation gets the remote object of the Animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
*/
func (protocol *AnimationProtocol) ResolveAnimation(
	params *animation.ResolveAnimationParams,
) chan *animation.ResolveAnimationResult {
	resultChan := make(chan *animation.ResolveAnimationResult)
	command := NewCommand(protocol.Socket, "Animation.resolveAnimation", params)
	result := &animation.ResolveAnimationResult{}

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
SeekAnimations seeks a set of animations to a particular time within each
animation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
*/
func (protocol *AnimationProtocol) SeekAnimations(
	params *animation.SeekAnimationsParams,
) chan *animation.SeekAnimationsResult {
	resultChan := make(chan *animation.SeekAnimationsResult)
	command := NewCommand(protocol.Socket, "Animation.seekAnimations", params)
	result := &animation.SeekAnimationsResult{}

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
SetPaused sets the paused state of a set of animations.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
*/
func (protocol *AnimationProtocol) SetPaused(
	params *animation.SetPausedParams,
) chan *animation.SetPausedResult {
	resultChan := make(chan *animation.SetPausedResult)
	command := NewCommand(protocol.Socket, "Animation.setPaused", params)
	result := &animation.SetPausedResult{}

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
SetPlaybackRate sets the playback rate of the document timeline.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
*/
func (protocol *AnimationProtocol) SetPlaybackRate(
	params *animation.SetPlaybackRateParams,
) chan *animation.SetPlaybackRateResult {
	resultChan := make(chan *animation.SetPlaybackRateResult)
	command := NewCommand(protocol.Socket, "Animation.setPlaybackRate", params)
	result := &animation.SetPlaybackRateResult{}

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
SetTiming sets the timing of an animation node.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
*/
func (protocol *AnimationProtocol) SetTiming(
	params *animation.SetTimingParams,
) chan *animation.SetTimingResult {
	resultChan := make(chan *animation.SetTimingResult)
	command := NewCommand(protocol.Socket, "Animation.setTiming", params)
	result := &animation.SetTimingResult{}

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
OnAnimationCanceled adds a handler to the Animation.Canceled event.
Animation.Canceled fires when when an animation has been cancelled.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
*/
func (protocol *AnimationProtocol) OnAnimationCanceled(
	callback func(event *animation.CanceledEvent),
) {
	handler := NewEventHandler(
		"Animation.animationCanceled",
		func(response *Response) {
			event := &animation.CanceledEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnAnimationCreated adds a handler to the Animation.Created event.
Animation.Created fires for each animation that has been created.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
*/
func (protocol *AnimationProtocol) OnAnimationCreated(
	callback func(event *animation.CreatedEvent),
) {
	handler := NewEventHandler(
		"Animation.animationCreated",
		func(response *Response) {
			event := &animation.CreatedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnAnimationStarted adds a handler to the Animation.Started event.
Animation.Started fires for each animation that has been started.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
*/
func (protocol *AnimationProtocol) OnAnimationStarted(
	callback func(event *animation.StartedEvent),
) {
	handler := NewEventHandler(
		"Animation.animationStarted",
		func(response *Response) {
			event := &animation.StartedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
