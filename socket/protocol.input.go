package socket

import (
	input "github.com/mkenney/go-chrome/protocol/input"
)

/*
InputProtocol provides a namespace for the Chrome Input protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Input/
*/
type InputProtocol struct {
	Socket Socketer
}

/*
DispatchKeyEvent dispatches a key event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchKeyEvent
*/
func (protocol *InputProtocol) DispatchKeyEvent(
	params *input.DispatchKeyEventParams,
) error {
	command := NewCommand("Input.dispatchKeyEvent", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DispatchMouseEvent dispatches a mouse event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
*/
func (protocol *InputProtocol) DispatchMouseEvent(
	params *input.DispatchMouseEventParams,
) error {
	command := NewCommand("Input.dispatchMouseEvent", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DispatchTouchEvent dispatches a touch event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
*/
func (protocol *InputProtocol) DispatchTouchEvent(
	params *input.DispatchTouchEventParams,
) error {
	command := NewCommand("Input.dispatchTouchEvent", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
EmulateTouchFromMouseEvent emulates touch event from the mouse event parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
EXPERIMENTAL.
*/
func (protocol *InputProtocol) EmulateTouchFromMouseEvent(
	params *input.EmulateTouchFromMouseEventParams,
) error {
	command := NewCommand("Input.emulateTouchFromMouseEvent", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetIgnoreEvents ignores input events (useful while auditing page).

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-setIgnoreInputEvents
*/
func (protocol *InputProtocol) SetIgnoreEvents(
	params *input.SetIgnoreEventsParams,
) error {
	command := NewCommand("Input.setIgnoreInputEvents", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizePinchGesture synthesizes a pinch gesture over a time period by issuing appropriate touch
events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizePinchGesture
EXPERIMENTAL.
*/
func (protocol *InputProtocol) SynthesizePinchGesture(
	params *input.SynthesizePinchGestureParams,
) error {
	command := NewCommand("Input.synthesizePinchGesture", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizeScrollGesture synthesizes a scroll gesture over a time period by issuing appropriate touch
events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeScrollGesture
EXPERIMENTAL.
*/
func (protocol *InputProtocol) SynthesizeScrollGesture(
	params *input.SynthesizeScrollGestureParams,
) error {
	command := NewCommand("Input.synthesizeScrollGesture", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SynthesizeTapGesture synthesizes a tap gesture over a time period by issuing appropriate touch
events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeTapGesture
EXPERIMENTAL.
*/
func (protocol *InputProtocol) SynthesizeTapGesture(
	params *input.SynthesizeTapGestureParams,
) error {
	command := NewCommand("Input.synthesizeTapGesture", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}
