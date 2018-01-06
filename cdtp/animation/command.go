package animation

import (
	"github.com/mkenney/go-chrome/cdtp/runtime"
)

/*
DisableResult represents the result of calls to Animation.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-disable
*/
type DisableResult struct {
	CDTPError error `json:"-"`
}

/*
EnableResult represents the result of calls to Animation.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-enable
*/
type EnableResult struct {
	CDTPError error `json:"-"`
}

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
	CurrentTime int64 `json:"currentTime"`

	CDTPError error `json:"-"`
}

/*
GetPlaybackRateResult represents the result of calls to Animation.getPlaybackRate.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
*/
type GetPlaybackRateResult struct {
	// Playback rate for animations on page.
	PlaybackRate float64 `json:"playbackRate"`

	CDTPError error `json:"-"`
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
ReleaseAnimationsResult represents the result of calls to Animation.releaseAnimations.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
*/
type ReleaseAnimationsResult struct {
	CDTPError error `json:"-"`
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

	CDTPError error `json:"-"`
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
SeekAnimationsResult represents the result of calls to Animation.seekAnimations.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
*/
type SeekAnimationsResult struct {
	CDTPError error `json:"-"`
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
SetPausedResult represents the result of calls to Animation.setPaused.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
*/
type SetPausedResult struct {
	CDTPError error `json:"-"`
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
SetPlaybackRateResult represents the result of calls to Animation.setPlaybackRate.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
*/
type SetPlaybackRateResult struct {
	CDTPError error `json:"-"`
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
SetTimingResult represents the result of calls to Animation.setTiming.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
*/
type SetTimingResult struct {
	CDTPError error `json:"-"`
}
