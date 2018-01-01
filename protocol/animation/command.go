package animation

import (
	"github.com/mkenney/go-chrome/protocol/runtime"
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
	RemoteObject *runtime.RemoteObject `json:"remoteObject"`
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
