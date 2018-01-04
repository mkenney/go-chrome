package socket

import (
	"encoding/json"

	domDebugger "github.com/mkenney/go-chrome/cdtp/dom_debugger"
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
	params *domDebugger.GetEventListenersParams,
) chan *domDebugger.GetEventListenersResult {
	resultChan := make(chan *domDebugger.GetEventListenersResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.getEventListeners", params)
	result := &domDebugger.GetEventListenersResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RemoveDOMBreakpoint removes the DOM breakpoint that was set using
setDOMBreakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveDOMBreakpoint(
	params *domDebugger.RemoveDOMBreakpointParams,
) chan *domDebugger.RemoveDOMBreakpointResult {
	resultChan := make(chan *domDebugger.RemoveDOMBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeDOMBreakpoint", params)
	result := &domDebugger.RemoveDOMBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveEventListenerBreakpoint(
	params *domDebugger.RemoveEventListenerBreakpointParams,
) chan *domDebugger.RemoveEventListenerBreakpointResult {
	resultChan := make(chan *domDebugger.RemoveEventListenerBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeEventListenerBreakpoint", params)
	result := &domDebugger.RemoveEventListenerBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RemoveInstrumentationBreakpoint removes breakpoint on particular native event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
EXPERIMENTAL.
*/
func (protocol *DOMDebuggerProtocol) RemoveInstrumentationBreakpoint(
	params *domDebugger.RemoveInstrumentationBreakpointParams,
) chan *domDebugger.RemoveInstrumentationBreakpointResult {
	resultChan := make(chan *domDebugger.RemoveInstrumentationBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeInstrumentationBreakpoint", params)
	result := &domDebugger.RemoveInstrumentationBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveXHRBreakpoint(
	params *domDebugger.RemoveXHRBreakpointParams,
) chan *domDebugger.RemoveXHRBreakpointResult {
	resultChan := make(chan *domDebugger.RemoveXHRBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.removeXHRBreakpoint", params)
	result := &domDebugger.RemoveXHRBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetDOMBreakpoint sets a breakpoint on a particular operation with DOM.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetDOMBreakpoint(
	params *domDebugger.SetDOMBreakpointParams,
) chan *domDebugger.SetDOMBreakpointResult {
	resultChan := make(chan *domDebugger.SetDOMBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setDOMBreakpoint", params)
	result := &domDebugger.SetDOMBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetEventListenerBreakpoint sets the breakpoint on a particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetEventListenerBreakpoint(
	params *domDebugger.SetEventListenerBreakpointParams,
) chan *domDebugger.SetEventListenerBreakpointResult {
	resultChan := make(chan *domDebugger.SetEventListenerBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setEventListenerBreakpoint", params)
	result := &domDebugger.SetEventListenerBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetInstrumentationBreakpoint sets breakpoint on particular native event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
EXPERIMENTAL.
*/
func (protocol *DOMDebuggerProtocol) SetInstrumentationBreakpoint(
	params *domDebugger.SetInstrumentationBreakpointParams,
) chan *domDebugger.SetInstrumentationBreakpointResult {
	resultChan := make(chan *domDebugger.SetInstrumentationBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setInstrumentationBreakpoint", params)
	result := &domDebugger.SetInstrumentationBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetXHRBreakpoint sets breakpoint on XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetXHRBreakpoint(
	params *domDebugger.SetXHRBreakpointParams,
) chan *domDebugger.SetXHRBreakpointResult {
	resultChan := make(chan *domDebugger.SetXHRBreakpointResult)
	command := NewCommand(protocol.Socket, "DOMDebugger.setXHRBreakpoint", params)
	result := &domDebugger.SetXHRBreakpointResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}
