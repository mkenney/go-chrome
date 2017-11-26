package chrome

import (
	"app/chrome/protocol"
)

type Animation struct{}

/*
Enable animation domain notifications.
*/
func (Animation) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Animation.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
Disable animation domain notifications.
*/
func (Animation) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Animation.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
GetPlaybackRate gets the playback rate of the document timeline.
*/
func (Animation) GetPlaybackRate(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Animation.getPlaybackRate"
	socket.SendCommand(command)
	return command.Err
}

/*
SetPlaybackRate sets the playback rate of the document timeline.
*/
func (Animation) SetPlaybackRate(socket *Socket, params *protocol.SetPlaybackRateParams) error {
	command := new(protocol.Command)
	command.method = "Animation.setPlaybackRate"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
GetCurrentTime returns the current time of the an animation.
*/
func (Animation) GetCurrentTime(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Animation.getCurrentTime"
	socket.SendCommand(command)
	return command.Err
}

/*
SetPaused sets the paused state of a set of animations.
*/
func (Animation) SetPaused(socket *Socket, params *protocol.SetPausedParams) error {
	command := new(protocol.Command)
	command.method = "Animation.setPaused"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetTiming sets the timing of an animation node.
*/
func (Animation) SetTiming(socket *Socket, params *protocol.SetTimingParams) error {
	command := new(protocol.Command)
	command.method = "Animation.setTiming"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SeekAnimations seeks a set of animations to a particular time within each animation.
*/
func (Animation) SeekAnimations(socket *Socket, params *protocol.SeekAnimationsParams) error {
	command := new(protocol.Command)
	command.method = "Animation.seekAnimations"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ReleaseAnimations releases a set of animations to no longer be manipulated.
*/
func (Animation) ReleaseAnimations(socket *Socket, params *protocol.ReleaseAnimationsParams) error {
	command := new(protocol.Command)
	command.method = "Animation.releaseAnimations"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ResolveAnimation gets the remote object of the Animation.
*/
func (Animation) ResolveAnimation(socket *Socket, params *protocol.ResolveAnimationParams) error {
	command := new(protocol.Command)
	command.method = "Animation.resolveAnimation"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
