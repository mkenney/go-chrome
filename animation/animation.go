/*
Package Animation provides type definitions for use with the Chrome Animation protocol

https://chromedevtools.github.io/devtools-protocol/tot/Animation/
*/
package Animation

import (
	"github.com/mkenney/go-chrome/runtime"
)

/*
GetCurrentTimeParams represents Animation.getCurrentTime parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
*/
type GetCurrentTimeParams struct {
	// ID of animation.
	ID string `json:"id"`
}

/*
GetCurrentTimeResult represents the result of calls to Animation.getCurrentTime.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
*/
type GetCurrentTimeResult struct {
	// ID of animation.
	ID string `json:"id"`
}

/*
GetPlaybackRateResult represents the result of calls to Animation.getPlaybackRate.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
*/
type GetPlaybackRateResult struct {
	// Playback rate for animations on page.
	PlaybackRate float64 `json:"playbackRate"`
}

/*
ReleaseAnimationsParams represents Animation.releaseAnimations parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
*/
type ReleaseAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

/*
ResolveAnimationParams represents Animation.resolveAnimation parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
*/
type ResolveAnimationParams struct {
	// Animation ID.
	AnimationID string `json:"animationId"`
}

/*
ResolveAnimationResult represents the result of calls to Animation.resolveAnimation.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
*/
type ResolveAnimationResult struct {
	// Corresponding remote object.
	RemoteObject Runtime.RemoteObject `json:"remoteObject"`
}

/*
SeekAnimationsParams represents Animation.seekAnimations parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
*/
type SeekAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`

	// Set the current time of each animation.
	CurrentTime float64 `json:"currentTime"`
}

/*
SetPausedParams represents Animation.setPaused parameters

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
*/
type SetPausedParams struct {
	// Animations to set the pause state of.
	Animations []string `json:"animations"`

	// Paused state to set to.
	Paused bool `json:"paused"`
}

/*
SetPlaybackRateParams represents Animation.setPlaybackRate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
*/
type SetPlaybackRateParams struct {
	// Playback rate for animations on page.
	PlaybackRate float64 `json:"playbackRate"`
}

/*
SetTimingParams represents Animation.setTiming parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
*/
type SetTimingParams struct {
	// Animation ID.
	AnimationID string `json:"animationId"`

	// Duration of the animation.
	Duration float64 `json:"duration"`

	// Delay of the animation.
	Delay float64 `json:"delay"`
}

/*
CanceledEvent represents Animation.animationCanceled event data.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
*/
type CanceledEvent struct {
	// ID of the animation that was cancelled.
	ID string `json:"id"`
}

/*
CreatedEvent represents Animation.animationCreated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
*/
type CreatedEvent struct {
	// ID of the animation that was created.
	ID string `json:"id"`
}

/*
StartedEvent represents Animation.animationStarted event data.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
*/
type StartedEvent struct {
	// Animation that was started.
	Animation Animation `json:"animation"`
}

/*
Animation instance.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-Animation
*/
type Animation struct {
	// Animation's id.
	ID string `json:"id"`

	// Animation's name.
	Name string `json:"name"`

	// Animation's internal paused state.
	PausedState bool `json:"pausedState"`

	// Animation's play state.
	PlayState string `json:"playState"`

	// Animation's playback rate.
	PlaybackRate float64 `json:"playbackRate"`

	// Animation's start time.
	StartTime float64 `json:"startTime"`

	// Animation's current time.
	CurrentTime float64 `json:"currentTime"`

	// Animation type of Animation. Allowed values: CSSTransition, CSSAnimation, WebAnimation.
	Type string `json:"type"`

	// Optional. Animation's source animation node.
	//
	// This expects an instance of AnimationEffect, but that doesn't omitempty correctly so it must
	// be added manually.
	Source interface{} `json:"source,omitempty"`

	// Optional. A unique ID for Animation representing the sources that triggered this CSS
	// animation/transition.
	CSSID string `json:"cssId,omitempty"`
}

/*
Effect instance

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-AnimationEffect
*/
type Effect struct {
	// Effect's delay.
	Delay float64 `json:"delay"`

	// Effect's end delay.
	EndDelay float64 `json:"endDelay"`

	// Effect's iteration start.
	IterationStart float64 `json:"iterationStart"`

	// Effect's iterations.
	Iterations float64 `json:"iterations"`

	// Effect's iteration duration.
	Duration float64 `json:"duration"`

	// Effect's playback direction.
	Direction string `json:"direction"`

	// Effect's fill mode.
	Fill string `json:"fill"`

	// Optional. Effect's target node.
	//
	// This expects an instance of DOM.BackendNodeID, but that doesn't omitempty correctly so it
	// must be added manually.
	BackendNodeID interface{} `json:"backendNodeId,omitempty"`

	// Optional. Effect's keyframes.
	//
	// This expects an instance of KeyframesRule, but that doesn't omitempty correctly so it must be
	// added manually.
	KeyframesRule interface{} `json:"keyframesRule,omitempty"`

	// Effect's timing function.
	Easing string `json:"easing"`
}

/*
KeyframesRule is the keyframes rule

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-KeyframesRule
*/
type KeyframesRule struct {
	// Optional. CSS keyframed animation's name.
	Name string `json:"name,omitempty"`

	// List of animation keyframes.
	Keyframes []*KeyframeStyle `json:"keyframes"`
}

/*
KeyframeStyle is the keyframe Style

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-KeyframeStyle
*/
type KeyframeStyle struct {
	// Keyframe's time offset.
	Offset string `json:"offset"`

	// Effect's timing function.
	Easing string `json:"easing"`
}
