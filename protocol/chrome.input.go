package protocol

import (
	input "github.com/mkenney/go-chrome/protocol/input"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Input provides a namespace for the Chrome Input protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Input/
*/
var Input = InputProtocol{}

/*
InputProtocol defines the Input protocol methods.
*/
type InputProtocol struct{}

/*
DispatchKeyEvent dispatches a key event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchKeyEvent
*/
func (InputProtocol) DispatchKeyEvent(
	socket sock.Socketer,
	params *input.DispatchKeyEventParams,
) error {
	command := sock.NewCommand("Input.dispatchKeyEvent", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DispatchMouseEvent dispatches a mouse event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
*/
func (InputProtocol) DispatchMouseEvent(
	socket sock.Socketer,
	params *input.DispatchMouseEventParams,
) error {
	command := sock.NewCommand("Input.dispatchMouseEvent", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DispatchTouchEvent dispatches a touch event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
*/
func (InputProtocol) DispatchTouchEvent(
	socket sock.Socketer,
	params *input.DispatchTouchEventParams,
) error {
	command := sock.NewCommand("Input.dispatchTouchEvent", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
EmulateTouchFromMouseEvent emulates touch event from the mouse event parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
EXPERIMENTAL.
*/
func (InputProtocol) EmulateTouchFromMouseEvent(
	socket sock.Socketer,
	params *input.EmulateTouchFromMouseEventParams,
) error {
	command := sock.NewCommand("Input.emulateTouchFromMouseEvent", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetIgnoreEvents ignores input events (useful while auditing page).

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-setIgnoreInputEvents
*/
func (InputProtocol) SetIgnoreEvents(
	socket sock.Socketer,
	params *input.SetIgnoreEventsParams,
) error {
	command := sock.NewCommand("Input.setIgnoreInputEvents", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizePinchGesture synthesizes a pinch gesture over a time period by issuing appropriate touch
events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizePinchGesture
EXPERIMENTAL.
*/
func (InputProtocol) SynthesizePinchGesture(
	socket sock.Socketer,
	params *input.SynthesizePinchGestureParams,
) error {
	command := sock.NewCommand("Input.synthesizePinchGesture", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizeScrollGesture synthesizes a scroll gesture over a time period by issuing appropriate touch
events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeScrollGesture
EXPERIMENTAL.
*/
func (InputProtocol) SynthesizeScrollGesture(
	socket sock.Socketer,
	params *input.SynthesizeScrollGestureParams,
) error {
	command := sock.NewCommand("Input.synthesizeScrollGesture", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizeTapGesture synthesizes a tap gesture over a time period by issuing appropriate touch
events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeTapGesture
EXPERIMENTAL.
*/
func (InputProtocol) SynthesizeTapGesture(
	socket sock.Socketer,
	params *input.SynthesizeTapGestureParams,
) error {
	command := sock.NewCommand("Input.synthesizeTapGesture", params)
	socket.SendCommand(command)
	return command.Error()
}
