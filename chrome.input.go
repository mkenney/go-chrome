package chrome

import "app/chrome/protocol"

/*
Input - https://chromedevtools.github.io/devtools-protocol/tot/Input/
*/
type Input struct{}

/*
DispatchKeyEvent dispatches a key event to the page.
*/
func (Input) DispatchKeyEvent(
	socket *Socket,
	params *input.DispatchKeyEventParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.dispatchKeyEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
DispatchMouseEvent dispatches a mouse event to the page.
*/
func (Input) DispatchMouseEvent(
	socket *Socket,
	params *input.DispatchMouseEventParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.dispatchMouseEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
DispatchTouchEvent dispatches a touch event to the page.
*/
func (Input) DispatchTouchEvent(
	socket *Socket,
	params *input.DispatchTouchEventParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.dispatchTouchEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
EmulateTouchFromMouseEvent emulates touch event from the mouse event parameters. EXPERIMENTAL
*/
func (Input) EmulateTouchFromMouseEvent(
	socket *Socket,
	params *input.EmulateTouchFromMouseEventParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.emulateTouchFromMouseEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetIgnoreInputEvents ignores input events (useful while auditing page).
*/
func (Input) SetIgnoreInputEvents(
	socket *Socket,
	params *input.SetIgnoreInputEventsParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.setIgnoreInputEvents",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SynthesizePinchGesture synthesizes a pinch gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL
*/
func (Input) SynthesizePinchGesture(
	socket *Socket,
	params *input.SynthesizePinchGestureParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.synthesizePinchGesture",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SynthesizeScrollGesture synthesizes a scroll gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL
*/
func (Input) SynthesizeScrollGesture(
	socket *Socket,
	params *input.SynthesizeScrollGestureParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.synthesizeScrollGesture",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SynthesizeTapGesture synthesizes a tap gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL
*/
func (Input) SynthesizeTapGesture(
	socket *Socket,
	params *input.SynthesizeTapGestureParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "Input.synthesizeTapGesture",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}
