package Chrome

import (
	input "github.com/mkenney/go-chrome/input"
	"github.com/mkenney/go-chrome/protocol"
)

/*
Input - https://chromedevtools.github.io/devtools-protocol/tot/Input/
*/
type Input struct{}

/*
DispatchKeyEvent dispatches a key event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchKeyEvent
*/
func (Input) DispatchKeyEvent(
	socket *Socket,
	params *input.DispatchKeyEventParams,
) error {
	command := &protocol.Command{
		Method: "Input.dispatchKeyEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DispatchMouseEvent dispatches a mouse event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
*/
func (Input) DispatchMouseEvent(
	socket *Socket,
	params *input.DispatchMouseEventParams,
) error {
	command := &protocol.Command{
		Method: "Input.dispatchMouseEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DispatchTouchEvent dispatches a touch event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
*/
func (Input) DispatchTouchEvent(
	socket *Socket,
	params *input.DispatchTouchEventParams,
) error {
	command := &protocol.Command{
		Method: "Input.dispatchTouchEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
EmulateTouchFromMouseEvent emulates touch event from the mouse event parameters. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
*/
func (Input) EmulateTouchFromMouseEvent(
	socket *Socket,
	params *input.EmulateTouchFromMouseEventParams,
) error {
	command := &protocol.Command{
		Method: "Input.emulateTouchFromMouseEvent",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetIgnoreEvents ignores input events (useful while auditing page).

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-setIgnoreInputEvents
*/
func (Input) SetIgnoreEvents(
	socket *Socket,
	params *input.SetIgnoreEventsParams,
) error {
	command := &protocol.Command{
		Method: "Input.setIgnoreInputEvents",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SynthesizePinchGesture synthesizes a pinch gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizePinchGesture
*/
func (Input) SynthesizePinchGesture(
	socket *Socket,
	params *input.SynthesizePinchGestureParams,
) error {
	command := &protocol.Command{
		Method: "Input.synthesizePinchGesture",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SynthesizeScrollGesture synthesizes a scroll gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeScrollGesture
*/
func (Input) SynthesizeScrollGesture(
	socket *Socket,
	params *input.SynthesizeScrollGestureParams,
) error {
	command := &protocol.Command{
		Method: "Input.synthesizeScrollGesture",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SynthesizeTapGesture synthesizes a tap gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeTapGesture
*/
func (Input) SynthesizeTapGesture(
	socket *Socket,
	params *input.SynthesizeTapGestureParams,
) error {
	command := &protocol.Command{
		Method: "Input.synthesizeTapGesture",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
