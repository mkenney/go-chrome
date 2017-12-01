package chrome

import (
	domdebugger "app/chrome/domdebugger"
	"app/chrome/protocol"
)

/*
DOMDebugger allows setting breakpoints on particular DOM operations and events. JavaScript execution
will stop on these operations as if there was a regular breakpoint set.
*/
type DOMDebugger struct{}

/*
SetDOMBreakpoint sets a breakpoint on a particular operation with DOM.
*/
func (DOMDebugger) SetDOMBreakpoint(socket *Socket, params *domdebugger.SetDOMBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.setDOMBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveDOMBreakpoint removes the DOM breakpoint that was set using setDOMBreakpoint.
*/
func (DOMDebugger) RemoveDOMBreakpoint(socket *Socket, params *domdebugger.RemoveDOMBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.removeDOMBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetEventListenerBreakpoint sets the breakpoint on a particular DOM event.
*/
func (DOMDebugger) SetEventListenerBreakpoint(socket *Socket, params *domdebugger.SetEventListenerBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.setEventListenerBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.
*/
func (DOMDebugger) RemoveEventListenerBreakpoint(socket *Socket, params *domdebugger.RemoveEventListenerBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.removeEventListenerBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetInstrumentationBreakpoint sets breakpoint on particular native event. EXPERIMENTAL
*/
func (DOMDebugger) SetInstrumentationBreakpoint(socket *Socket, params *domdebugger.SetInstrumentationBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.setInstrumentationBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveInstrumentationBreakpoint removes breakpoint on particular native event. EXPERIMENTAL
*/
func (DOMDebugger) RemoveInstrumentationBreakpoint(socket *Socket, params *domdebugger.RemoveInstrumentationBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.removeInstrumentationBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetXHRBreakpoint sets breakpoint on XMLHttpRequest.
*/
func (DOMDebugger) SetXHRBreakpoint(socket *Socket, params *domdebugger.SetXHRBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.setXHRBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.
*/
func (DOMDebugger) RemoveXHRBreakpoint(socket *Socket, params *domdebugger.RemoveXHRBreakpointParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.removeXHRBreakpoint"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
GetEventListeners returns event listeners of the given object.
*/
func (DOMDebugger) GetEventListeners(socket *Socket, params *domdebugger.GetEventListenersParams) error {
	command := new(protocol.Command)
	command.method = "DOMDebugger.getEventListeners"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
