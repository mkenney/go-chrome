package protocol

import (
	"encoding/json"

	domDebugger "github.com/mkenney/go-chrome/protocol/dom_debugger"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
DOMDebugger is a struct that provides a namespace for the Chrome DOMDebugger protocol methods. The
DOMDebugger protocol allows setting breakpoints on particular DOM operations and events. JavaScript
execution will stop on these operations as if there was a regular breakpoint set.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/
*/
var DOMDebugger = _domDebugger{}

type _domDebugger struct{}

/*
GetEventListeners returns event listeners of the given object.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
*/
func (_domDebugger) GetEventListeners(
	socket sock.Socketer,
	params *domDebugger.GetEventListenersParams,
) (domDebugger.GetEventListenersResult, error) {
	command := sock.NewCommand("DOMDebugger.getEventListeners", params)
	result := domDebugger.GetEventListenersResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
RemoveDOMBreakpoint removes the DOM breakpoint that was set using setDOMBreakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
*/
func (_domDebugger) RemoveDOMBreakpoint(
	socket sock.Socketer,
	params *domDebugger.RemoveDOMBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.removeDOMBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
*/
func (_domDebugger) RemoveEventListenerBreakpoint(
	socket sock.Socketer,
	params *domDebugger.RemoveEventListenerBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.removeEventListenerBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveInstrumentationBreakpoint removes breakpoint on particular native event. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
*/
func (_domDebugger) RemoveInstrumentationBreakpoint(
	socket sock.Socketer,
	params *domDebugger.RemoveInstrumentationBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.removeInstrumentationBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
*/
func (_domDebugger) RemoveXHRBreakpoint(
	socket sock.Socketer,
	params *domDebugger.RemoveXHRBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.removeXHRBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetDOMBreakpoint sets a breakpoint on a particular operation with DOM.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
*/
func (_domDebugger) SetDOMBreakpoint(
	socket sock.Socketer,
	params *domDebugger.SetDOMBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.setDOMBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetEventListenerBreakpoint sets the breakpoint on a particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
*/
func (_domDebugger) SetEventListenerBreakpoint(
	socket sock.Socketer,
	params *domDebugger.SetEventListenerBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.setEventListenerBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetInstrumentationBreakpoint sets breakpoint on particular native event. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
*/
func (_domDebugger) SetInstrumentationBreakpoint(
	socket sock.Socketer,
	params *domDebugger.SetInstrumentationBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.setInstrumentationBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetXHRBreakpoint sets breakpoint on XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
*/
func (_domDebugger) SetXHRBreakpoint(
	socket sock.Socketer,
	params *domDebugger.SetXHRBreakpointParams,
) error {
	command := sock.NewCommand("DOMDebugger.setXHRBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}
