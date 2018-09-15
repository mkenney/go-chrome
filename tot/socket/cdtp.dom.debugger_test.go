package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/dom/debugger"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestDOMDebuggerGetEventListeners(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.GetEventListenersResult{
		Listeners: []*debugger.EventListener{{
			Type:         "listener-type",
			UseCapture:   true,
			Passive:      true,
			Once:         true,
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
			Handler: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			OriginalHandler: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			BackendNodeID: dom.BackendNodeID(1),
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().GetEventListeners(&debugger.GetEventListenersParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
		Depth:    1,
		Pierce:   true,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Listeners[0].BackendNodeID != result.Listeners[0].BackendNodeID {
		t.Errorf("Expected '%d', got '%d'", mockResult.Listeners[0].BackendNodeID, result.Listeners[0].BackendNodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().GetEventListeners(&debugger.GetEventListenersParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
		Depth:    1,
		Pierce:   true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveDOMBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.RemoveDOMBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().RemoveDOMBreakpoint(&debugger.RemoveDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().RemoveDOMBreakpoint(&debugger.RemoveDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveEventListenerBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.RemoveEventListenerBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().RemoveEventListenerBreakpoint(&debugger.RemoveEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().RemoveEventListenerBreakpoint(&debugger.RemoveEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveInstrumentationBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.RemoveInstrumentationBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().RemoveInstrumentationBreakpoint(&debugger.RemoveInstrumentationBreakpointParams{
		EventName: "event name",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().RemoveInstrumentationBreakpoint(&debugger.RemoveInstrumentationBreakpointParams{
		EventName: "event name",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveXHRBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.RemoveXHRBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().RemoveXHRBreakpoint(&debugger.RemoveXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().RemoveXHRBreakpoint(&debugger.RemoveXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetDOMBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetDOMBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().SetDOMBreakpoint(&debugger.SetDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().SetDOMBreakpoint(&debugger.SetDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetEventListenerBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetEventListenerBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().SetEventListenerBreakpoint(&debugger.SetEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().SetEventListenerBreakpoint(&debugger.SetEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetInstrumentationBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetInstrumentationBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().SetInstrumentationBreakpoint(&debugger.SetInstrumentationBreakpointParams{
		EventName: "event name",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().SetInstrumentationBreakpoint(&debugger.SetInstrumentationBreakpointParams{
		EventName: "event name",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetXHRBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetXHRBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMDebugger().SetXHRBreakpoint(&debugger.SetXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMDebugger().SetXHRBreakpoint(&debugger.SetXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
