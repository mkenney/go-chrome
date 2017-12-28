package protocol

import (
	input "github.com/mkenney/go-chrome/protocol/input"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Input is a struct that provides a namespace for the Chrome Input protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Input/
*/
var Input = _input{}

type _input struct{}

/*
DispatchKeyEvent dispatches a key event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchKeyEvent
*/
func (_input) DispatchKeyEvent(
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
func (_input) DispatchMouseEvent(
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
func (_input) DispatchTouchEvent(
	socket sock.Socketer,
	params *input.DispatchTouchEventParams,
) error {
	command := sock.NewCommand("Input.dispatchTouchEvent", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
EmulateTouchFromMouseEvent emulates touch event from the mouse event parameters. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
*/
func (_input) EmulateTouchFromMouseEvent(
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
func (_input) SetIgnoreEvents(
	socket sock.Socketer,
	params *input.SetIgnoreEventsParams,
) error {
	command := sock.NewCommand("Input.setIgnoreInputEvents", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizePinchGesture synthesizes a pinch gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizePinchGesture
*/
func (_input) SynthesizePinchGesture(
	socket sock.Socketer,
	params *input.SynthesizePinchGestureParams,
) error {
	command := sock.NewCommand("Input.synthesizePinchGesture", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizeScrollGesture synthesizes a scroll gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeScrollGesture
*/
func (_input) SynthesizeScrollGesture(
	socket sock.Socketer,
	params *input.SynthesizeScrollGestureParams,
) error {
	command := sock.NewCommand("Input.synthesizeScrollGesture", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizeTapGesture synthesizes a tap gesture over a time period by issuing appropriate touch
events. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeTapGesture
*/
func (_input) SynthesizeTapGesture(
	socket sock.Socketer,
	params *input.SynthesizeTapGestureParams,
) error {
	command := sock.NewCommand("Input.synthesizeTapGesture", params)
	socket.SendCommand(command)
	return command.Error()
}
