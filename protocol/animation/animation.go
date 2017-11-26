package Animation

import (
	DOM "app/chrome/protocol/dom"
)

/*
Animation instance
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
	PlaybackRate int `json:"playbackRate"`

	// Animation's start time.
	StartTime int `json:"startTime"`

	// Animation's current time.
	CurrentTime int `json:"currentTime"`

	// Animation type of Animation. Allowed values: CSSTransition, CSSAnimation, WebAnimation.
	Type string `json:"type"`

	// Optional. Animation's source animation node.
	Source AnimationEffect `json:"source,omitempty"`

	// Optional. A unique ID for Animation representing the sources that triggered this CSS
	// animation/transition.
	CSSID string `json:"cssId,omitempty"`
}

/*
AnimationEffect instance
*/
type AnimationEffect struct {
	// AnimationEffect's delay.
	Delay int `json:"delay"`

	// AnimationEffect's end delay.
	EndDelay int `json:"endDelay"`

	// AnimationEffect's iteration start.
	IterationStart int `json:"iterationStart"`

	// AnimationEffect's iterations.
	Iterations int `json:"iterations"`

	// AnimationEffect's iteration duration.
	Duration int `json:"duration"`

	// AnimationEffect's playback direction.
	Direction string `json:"direction"`

	// AnimationEffect's fill mode.
	Fill string `json:"fill"`

	// Optional. AnimationEffect's target node.
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. AnimationEffect's keyframes.
	KeyframesRule KeyframesRule `json:"keyframesRule,omitempty"`

	// AnimationEffect's timing function.
	Easing string `json:"easing"`
}

/*
KeyframesRule is the keyframes rule
*/
type KeyframesRule struct {
	// Optional. CSS keyframed animation's name.
	Name string `json:"name,omitempty"`

	// List of animation keyframes.
	Keyframes []*KeyframeStyle `json:"keyframes"`
}

/*
KeyframeStyle is the keyframe Style
*/
type KeyframeStyle struct {
	// Keyframe's time offset.
	Offset string `json:"offset"`

	// AnimationEffect's timing function.
	Easing string `json:"easing"`
}

/*
SetPlaybackRateParams represents Animation.setPaused parameters
*/
type SetPlaybackRateParams struct {
	// Playback rate for animations on page.
	PlaybackRate int `json:"playbackRate"`
}

/*
SetPausedParams represents Animation.setPaused parameters
*/
type SetPausedParams struct {
	// Animations to set the pause state of.
	Animations []string `json:"animations"`

	// Paused state to set to.
	Paused bool `json:"paused"`
}

/*
SetTimingParams represents Animation.setTiming parameters
*/
type SetTimingParams struct {
	// Animation id.
	AnimationId string `json:"animationId"`

	// Duration of the animation.
	Duration int `json:"duration"`

	// Delay of the animation.
	Delay int `json:"delay"`
}

/*
SeekAnimationsParams represents Animation.seekAnimations parameters
*/
type SeekAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`

	// Set the current time of each animation.
	CurrentTime int `json:"currentTime"`
}

/*
ReleaseAnimationsParams represents Animation.releaseAnimations parameters
*/
type ReleaseAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

/*
ResolveAnimationParams represents Animation.resolveAnimation parameters
*/
type ResolveAnimationParams struct {
	// Animation id.
	AnimationId string `json:"animationId"`
}
