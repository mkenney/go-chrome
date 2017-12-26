package chrome

import (
	"encoding/json"

	domDebugger "github.com/mkenney/go-chrome/dom_debugger"
	"github.com/mkenney/go-chrome/protocol"
)

/*
DOMDebugger - https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/
Allows setting breakpoints on particular DOM operations and events. JavaScript execution will stop
on these operations as if there was a regular breakpoint set.
*/
type DOMDebugger struct{}

/*
GetEventListeners returns event listeners of the given object.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
*/
func (DOMDebugger) GetEventListeners(
	socket *Socket,
	params *domDebugger.GetEventListenersParams,
) (domDebugger.GetEventListenersResult, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.getEventListeners",
		Params: params,
	}
	result := domDebugger.GetEventListenersResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
RemoveDOMBreakpoint removes the DOM breakpoint that was set using setDOMBreakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
*/
func (DOMDebugger) RemoveDOMBreakpoint(
	socket *Socket,
	params *domDebugger.RemoveDOMBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.removeDOMBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
*/
func (DOMDebugger) RemoveEventListenerBreakpoint(
	socket *Socket,
	params *domDebugger.RemoveEventListenerBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.removeEventListenerBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveInstrumentationBreakpoint removes breakpoint on particular native event. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
*/
func (DOMDebugger) RemoveInstrumentationBreakpoint(
	socket *Socket,
	params *domDebugger.RemoveInstrumentationBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.removeInstrumentationBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
*/
func (DOMDebugger) RemoveXHRBreakpoint(
	socket *Socket,
	params *domDebugger.RemoveXHRBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.removeXHRBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDOMBreakpoint sets a breakpoint on a particular operation with DOM.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
*/
func (DOMDebugger) SetDOMBreakpoint(
	socket *Socket,
	params *domDebugger.SetDOMBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.setDOMBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetEventListenerBreakpoint sets the breakpoint on a particular DOM event.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
*/
func (DOMDebugger) SetEventListenerBreakpoint(
	socket *Socket,
	params *domDebugger.SetEventListenerBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.setEventListenerBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetInstrumentationBreakpoint sets breakpoint on particular native event. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
*/
func (DOMDebugger) SetInstrumentationBreakpoint(
	socket *Socket,
	params *domDebugger.SetInstrumentationBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.setInstrumentationBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetXHRBreakpoint sets breakpoint on XMLHttpRequest.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
*/
func (DOMDebugger) SetXHRBreakpoint(
	socket *Socket,
	params *domDebugger.SetXHRBreakpointParams,
) error {
	command := &protocol.Command{
		Method: "DOMDebugger.setXHRBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
