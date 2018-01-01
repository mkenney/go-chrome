package domDebugger

import (
	"github.com/mkenney/go-chrome/protocol/dom"
	"github.com/mkenney/go-chrome/protocol/runtime"
)

/*
GetEventListenersParams represents DOMDebugger.getEventListeners parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
*/
type GetEventListenersParams struct {
	// ID of the object to return listeners for.
	ObjectID runtime.RemoteObjectID `json:"objectId"`

	// Optional. The maximum depth at which Node children should be retrieved,
	// defaults to 1. Use -1 for the entire subtree or provide an integer larger
	// than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed
	// when returning the subtree (default is false). Reports listeners for all
	// contexts if pierce is enabled.
	Pierce bool `json:"pierce,omitempty"`
}

/*
GetEventListenersResult represents the result of calls to DOMDebugger.getEventListeners.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
*/
type GetEventListenersResult struct {
	// Array of relevant listeners.
	Listeners []*EventListener `json:"listeners"`
}

/*
RemoveDOMBreakpointParams represents DOMDebugger.removeDOMBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
*/
type RemoveDOMBreakpointParams struct {
	// ID of the node to remove breakpoint from.
	NodeID dom.NodeID `json:"nodeId"`

	// Type of the breakpoint to remove.
	Type DOMBreakpointType `json:"type"`
}

/*
RemoveEventListenerBreakpointParams represents DOMDebugger.removeEventListenerBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
*/
type RemoveEventListenerBreakpointParams struct {
	// Event name.
	EventName string `json:"eventName"`

	// Optional. EventTarget interface name. EXPERIMENTAL
	TargetName string `json:"targetName,omitempty"`
}

/*
RemoveInstrumentationBreakpointParams represents DOMDebugger.removeInstrumentationBreakpoint
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
*/
type RemoveInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

/*
RemoveXHRBreakpointParams represents DOMDebugger.removeXHRBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
*/
type RemoveXHRBreakpointParams struct {
	// Resource URL substring.
	URL string `json:"url"`
}

/*
SetDOMBreakpointParams represents DOMDebugger.setDOMBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
*/
type SetDOMBreakpointParams struct {
	// ID of the node to set breakpoint on.
	NodeID dom.NodeID `json:"nodeId"`

	// Type of the operation to stop upon.
	Type DOMBreakpointType `json:"type"`
}

/*
SetEventListenerBreakpointParams represents DOMDebugger.setEventListenerBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
*/
type SetEventListenerBreakpointParams struct {
	// DOM Event name to stop on (any DOM event will do).
	EventName string `json:"eventName"`

	// Optional. EventTarget interface name to stop on. If equal to "*" or not
	// provided, will stop on any EventTarget. EXPERIMENTAL
	TargetName string `json:"targetName,omitempty"`
}

/*
SetInstrumentationBreakpointParams represents DOMDebugger.setInstrumentationBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
*/
type SetInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

/*
SetXHRBreakpointParams represents DOMDebugger.setXHRBreakpoint parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
*/
type SetXHRBreakpointParams struct {
	// Resource URL substring. All XHRs having this substring in the URL will
	// get stopped upon.
	URL string `json:"url"`
}
