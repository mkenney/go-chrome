package socket

import (
	domDebugger "github.com/mkenney/go-chrome/protocol/dom_debugger"
)

/*
DOMDebuggerProtocol provides a namespace for the Chrome DOMDebugger protocol methods. The
DOMDebugger protocol allows setting breakpoints on particular DOM operations and events. JavaScript
execution will stop on these operations as if there was a regular breakpoint set.

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
) (*domDebugger.GetEventListenersResult, error) {
	command := NewCommand("DOMDebugger.getEventListeners", params)
	result := &domDebugger.GetEventListenersResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
RemoveDOMBreakpoint removes the DOM breakpoint that was set using setDOMBreakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveDOMBreakpoint(
	params *domDebugger.RemoveDOMBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.removeDOMBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveEventListenerBreakpoint(
	params *domDebugger.RemoveEventListenerBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.removeEventListenerBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveInstrumentationBreakpoint removes breakpoint on particular native event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
EXPERIMENTAL.
*/
func (protocol *DOMDebuggerProtocol) RemoveInstrumentationBreakpoint(
	params *domDebugger.RemoveInstrumentationBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.removeInstrumentationBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
*/
func (protocol *DOMDebuggerProtocol) RemoveXHRBreakpoint(
	params *domDebugger.RemoveXHRBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.removeXHRBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetDOMBreakpoint sets a breakpoint on a particular operation with DOM.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetDOMBreakpoint(
	params *domDebugger.SetDOMBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.setDOMBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetEventListenerBreakpoint sets the breakpoint on a particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetEventListenerBreakpoint(
	params *domDebugger.SetEventListenerBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.setEventListenerBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetInstrumentationBreakpoint sets breakpoint on particular native event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
EXPERIMENTAL.
*/
func (protocol *DOMDebuggerProtocol) SetInstrumentationBreakpoint(
	params *domDebugger.SetInstrumentationBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.setInstrumentationBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetXHRBreakpoint sets breakpoint on XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
*/
func (protocol *DOMDebuggerProtocol) SetXHRBreakpoint(
	params *domDebugger.SetXHRBreakpointParams,
) error {
	command := NewCommand("DOMDebugger.setXHRBreakpoint", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}
