package chrome

import (
	input "app/chrome/input"
	"app/chrome/protocol"
)

/*
Input domain
*/
type Input struct{}

/*
DispatchKeyEvent dispatches a key event to the page.
*/
func (Input) DispatchKeyEvent(socket *Socket, params *input.DispatchKeyEventParams) error {
	command := new(protocol.Command)
	command.method = "Input.dispatchKeyEvent"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
DispatchMouseEvent dispatches a mouse event to the page.
*/
func (Input) DispatchMouseEvent(socket *Socket, params *input.DispatchMouseEventParams) error {
	command := new(protocol.Command)
	command.method = "Input.dispatchMouseEvent"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
DispatchTouchEvent dispatches a touch event to the page.
*/
func (Input) DispatchTouchEvent(socket *Socket, params *input.DispatchTouchEventParams) error {
	command := new(protocol.Command)
	command.method = "Input.dispatchTouchEvent"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
EmulateTouchFromMouseEvent emulates touch event from the mouse event parameters. EXPERIMENTAL
*/
func (Input) EmulateTouchFromMouseEvent(socket *Socket, params *input.EmulateTouchFromMouseEventParams) error {
	command := new(protocol.Command)
	command.method = "Input.emulateTouchFromMouseEvent"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetIgnoreInputEvents ignores input events (useful while auditing page).
*/
func (Input) SetIgnoreInputEvents(socket *Socket, params *input.SetIgnoreInputEventsParams) error {
	command := new(protocol.Command)
	command.method = "Input.setIgnoreInputEvents"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SynthesizePinchGesture synthesizes a pinch gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL
*/
func (Input) SynthesizePinchGesture(socket *Socket, params *input.SynthesizePinchGestureParams) error {
	command := new(protocol.Command)
	command.method = "Input.synthesizePinchGesture"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SynthesizeScrollGesture synthesizes a scroll gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL
*/
func (Input) SynthesizeScrollGesture(socket *Socket, params *input.SynthesizeScrollGestureParams) error {
	command := new(protocol.Command)
	command.method = "Input.synthesizeScrollGesture"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SynthesizeTapGesture synthesizes a tap gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL
*/
func (Input) SynthesizeTapGesture(socket *Socket, params *input.SynthesizeTapGestureParams) error {
	command := new(protocol.Command)
	command.method = "Input.synthesizeTapGesture"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
