/*
Package animation provides type definitions for use with the Chrome Animation
protocol

https://chromedevtools.github.io/devtools-protocol/tot/Animation/
*/
package animation

import (
	"github.com/mkenney/go-chrome/tot/cdtp/dom"
)

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
	StartTime int64 `json:"startTime"`

	// Animation's current time.
	CurrentTime int64 `json:"currentTime"`

	// Animation type of Animation. Allowed values: CSSTransition, CSSAnimation,
	// WebAnimation.
	Type string `json:"type"`

	// Optional. Animation's source animation node.
	Source *Effect `json:"source,omitempty"`

	// Optional. A unique ID for Animation representing the sources that
	// triggered this CSS animation/transition.
	CSSID string `json:"cssId,omitempty"`
}

/*
Effect instance

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-AnimationEffect
*/
type Effect struct {
	// Effect's delay.
	Delay int64 `json:"delay"`

	// Effect's end delay.
	EndDelay int64 `json:"endDelay"`

	// Effect's iteration start.
	IterationStart int64 `json:"iterationStart"`

	// Effect's iterations.
	Iterations int64 `json:"iterations"`

	// Effect's iteration duration.
	Duration int64 `json:"duration"`

	// Effect's playback direction.
	Direction string `json:"direction"`

	// Effect's fill mode.
	Fill string `json:"fill"`

	// Optional. Effect's target node.
	BackendNodeID dom.BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. Effect's keyframes.
	KeyframesRule *KeyframesRule `json:"keyframesRule,omitempty"`

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
