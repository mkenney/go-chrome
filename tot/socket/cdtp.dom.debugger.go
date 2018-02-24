package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/dom/debugger"
)

/*
DOMDebuggerProtocol provides a namespace for the Chrome DOMDebugger protocol
methods. The DOMDebugger protocol allows setting breakpoints on particular DOM
operations and events. JavaScript execution will stop on these operations as if
there was a regular breakpoint set.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/
*/
type DOMDebuggerProtocol struct {
	Socket Socketer
}

/*
GetEventListeners returns event listeners of the given object.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
*/
func (protocol *DOMDebuggerProtocol) GetEventListeners(
	params *debugger.GetEventListenersParams,
) <-chan *debugger.GetEventListenersResult {
	resultChan := make(chan *debugger.GetEventListenersResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.getEventListeners", params)
	result := &debugger.GetEventListenersResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveDOMBreakpoint removes the DOM breakpoint that was set using
setDOMBreakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveDOMBreakpoint(
	params *debugger.RemoveDOMBreakpointParams,
) <-chan *debugger.RemoveDOMBreakpointResult {
	resultChan := make(chan *debugger.RemoveDOMBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeDOMBreakpoint", params)
	result := &debugger.RemoveDOMBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveEventListenerBreakpoint(
	params *debugger.RemoveEventListenerBreakpointParams,
) <-chan *debugger.RemoveEventListenerBreakpointResult {
	resultChan := make(chan *debugger.RemoveEventListenerBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeEventListenerBreakpoint", params)
	result := &debugger.RemoveEventListenerBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveInstrumentationBreakpoint removes breakpoint on particular native event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
EXPERIMENTAL.
*/
func (protocol *DOMDebuggerProtocol) RemoveInstrumentationBreakpoint(
	params *debugger.RemoveInstrumentationBreakpointParams,
) <-chan *debugger.RemoveInstrumentationBreakpointResult {
	resultChan := make(chan *debugger.RemoveInstrumentationBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeInstrumentationBreakpoint", params)
	result := &debugger.RemoveInstrumentationBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveXHRBreakpoint(
	params *debugger.RemoveXHRBreakpointParams,
) <-chan *debugger.RemoveXHRBreakpointResult {
	resultChan := make(chan *debugger.RemoveXHRBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeXHRBreakpoint", params)
	result := &debugger.RemoveXHRBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetDOMBreakpoint sets a breakpoint on a particular operation with DOM.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetDOMBreakpoint(
	params *debugger.SetDOMBreakpointParams,
) <-chan *debugger.SetDOMBreakpointResult {
	resultChan := make(chan *debugger.SetDOMBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setDOMBreakpoint", params)
	result := &debugger.SetDOMBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetEventListenerBreakpoint sets the breakpoint on a particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetEventListenerBreakpoint(
	params *debugger.SetEventListenerBreakpointParams,
) <-chan *debugger.SetEventListenerBreakpointResult {
	resultChan := make(chan *debugger.SetEventListenerBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setEventListenerBreakpoint", params)
	result := &debugger.SetEventListenerBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetInstrumentationBreakpoint sets breakpoint on particular native event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
EXPERIMENTAL.
*/
func (protocol *DOMDebuggerProtocol) SetInstrumentationBreakpoint(
	params *debugger.SetInstrumentationBreakpointParams,
) <-chan *debugger.SetInstrumentationBreakpointResult {
	resultChan := make(chan *debugger.SetInstrumentationBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setInstrumentationBreakpoint", params)
	result := &debugger.SetInstrumentationBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetXHRBreakpoint sets breakpoint on XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetXHRBreakpoint(
	params *debugger.SetXHRBreakpointParams,
) <-chan *debugger.SetXHRBreakpointResult {
	resultChan := make(chan *debugger.SetXHRBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setXHRBreakpoint", params)
	result := &debugger.SetXHRBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}
