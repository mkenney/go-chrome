package chrome

import "app/chrome/protocol"

/*
DOMDebugger - https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/
Allows setting breakpoints on particular DOM operations and events. JavaScript execution will stop
on these operations as if there was a regular breakpoint set.
*/
type DOMDebugger struct{}

/*
GetEventListeners returns event listeners of the given object.
*/
func (DOMDebugger) GetEventListeners(
	socket *Socket,
	params *dom_debugger.GetEventListenersParams,
) (dom_debugger.GetEventListenersResult, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.getEventListeners",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom_debugger.GetEventListenersResult), command.Err
}

/*
RemoveDOMBreakpoint removes the DOM breakpoint that was set using setDOMBreakpoint.
*/
func (DOMDebugger) RemoveDOMBreakpoint(
	socket *Socket,
	params *dom_debugger.RemoveDOMBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.removeDOMBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.
*/
func (DOMDebugger) RemoveEventListenerBreakpoint(
	socket *Socket,
	params *dom_debugger.RemoveEventListenerBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.removeEventListenerBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveInstrumentationBreakpoint removes breakpoint on particular native event. EXPERIMENTAL
*/
func (DOMDebugger) RemoveInstrumentationBreakpoint(
	socket *Socket,
	params *dom_debugger.RemoveInstrumentationBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.removeInstrumentationBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.
*/
func (DOMDebugger) RemoveXHRBreakpoint(
	socket *Socket,
	params *dom_debugger.RemoveXHRBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.removeXHRBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetDOMBreakpoint sets a breakpoint on a particular operation with DOM.
*/
func (DOMDebugger) SetDOMBreakpoint(
	socket *Socket,
	params *dom_debugger.SetDOMBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.setDOMBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetEventListenerBreakpoint sets the breakpoint on a particular DOM event.
*/
func (DOMDebugger) SetEventListenerBreakpoint(
	socket *Socket,
	params *dom_debugger.SetEventListenerBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.setEventListenerBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetInstrumentationBreakpoint sets breakpoint on particular native event. EXPERIMENTAL
*/
func (DOMDebugger) SetInstrumentationBreakpoint(
	socket *Socket,
	params *dom_debugger.SetInstrumentationBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.setInstrumentationBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetXHRBreakpoint sets breakpoint on XMLHttpRequest.
*/
func (DOMDebugger) SetXHRBreakpoint(
	socket *Socket,
	params *dom_debugger.SetXHRBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "DOMDebugger.setXHRBreakpoint",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}
