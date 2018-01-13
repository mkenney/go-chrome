package socket

import (
	input "github.com/mkenney/go-chrome/tot/cdtp/input"
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
) chan *input.DispatchKeyEventResult {
	resultChan := make(chan *input.DispatchKeyEventResult)
	command := NewCommand(protocol.Socket, "Input.dispatchKeyEvent", params)
	result := &input.DispatchKeyEventResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DispatchMouseEvent dispatches a mouse event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
*/
func (protocol *InputProtocol) DispatchMouseEvent(
	params *input.DispatchMouseEventParams,
) chan *input.DispatchMouseEventResult {
	resultChan := make(chan *input.DispatchMouseEventResult)
	command := NewCommand(protocol.Socket, "Input.dispatchMouseEvent", params)
	result := &input.DispatchMouseEventResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DispatchTouchEvent dispatches a touch event to the page.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
*/
func (protocol *InputProtocol) DispatchTouchEvent(
	params *input.DispatchTouchEventParams,
) chan *input.DispatchTouchEventResult {
	resultChan := make(chan *input.DispatchTouchEventResult)
	command := NewCommand(protocol.Socket, "Input.dispatchTouchEvent", params)
	result := &input.DispatchTouchEventResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
EmulateTouchFromMouseEvent emulates touch event from the mouse event parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
EXPERIMENTAL.
*/
func (protocol *InputProtocol) EmulateTouchFromMouseEvent(
	params *input.EmulateTouchFromMouseEventParams,
) chan *input.EmulateTouchFromMouseEventResult {
	resultChan := make(chan *input.EmulateTouchFromMouseEventResult)
	command := NewCommand(protocol.Socket, "Input.emulateTouchFromMouseEvent", params)
	result := &input.EmulateTouchFromMouseEventResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetIgnoreEvents ignores input events (useful while auditing page).

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-setIgnoreInputEvents
*/
func (protocol *InputProtocol) SetIgnoreEvents(
	params *input.SetIgnoreEventsParams,
) chan *input.SetIgnoreEventsResult {
	resultChan := make(chan *input.SetIgnoreEventsResult)
	command := NewCommand(protocol.Socket, "Input.setIgnoreInputEvents", params)
	result := &input.SetIgnoreEventsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SynthesizePinchGesture synthesizes a pinch gesture over a time period by issuing
appropriate touch events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizePinchGesture
EXPERIMENTAL.
*/
func (protocol *InputProtocol) SynthesizePinchGesture(
	params *input.SynthesizePinchGestureParams,
) chan *input.SynthesizePinchGestureResult {
	resultChan := make(chan *input.SynthesizePinchGestureResult)
	command := NewCommand(protocol.Socket, "Input.synthesizePinchGesture", params)
	result := &input.SynthesizePinchGestureResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SynthesizeScrollGesture synthesizes a scroll gesture over a time period by
issuing appropriate touch events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeScrollGesture
EXPERIMENTAL.
*/
func (protocol *InputProtocol) SynthesizeScrollGesture(
	params *input.SynthesizeScrollGestureParams,
) chan *input.SynthesizeScrollGestureResult {
	resultChan := make(chan *input.SynthesizeScrollGestureResult)
	command := NewCommand(protocol.Socket, "Input.synthesizeScrollGesture", params)
	result := &input.SynthesizeScrollGestureResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SynthesizeTapGesture synthesizes a tap gesture over a time period by issuing
appropriate touch events.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeTapGesture
EXPERIMENTAL.
*/
func (protocol *InputProtocol) SynthesizeTapGesture(
	params *input.SynthesizeTapGestureParams,
) chan *input.SynthesizeTapGestureResult {
	resultChan := make(chan *input.SynthesizeTapGestureResult)
	command := NewCommand(protocol.Socket, "Input.synthesizeTapGesture", params)
	result := &input.SynthesizeTapGestureResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}
