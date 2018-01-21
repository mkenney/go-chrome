package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	dom "github.com/mkenney/go-chrome/tot/cdtp/dom"
	"github.com/mkenney/go-chrome/tot/cdtp/dom_debugger"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestDOMDebuggerGetEventListeners(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().GetEventListeners(&domDebugger.GetEventListenersParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
		Depth:    1,
		Pierce:   true,
	})
	mockResult := &domDebugger.GetEventListenersResult{
		Listeners: []*domDebugger.EventListener{{
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Listeners[0].BackendNodeID != result.Listeners[0].BackendNodeID {
		t.Errorf("Expected '%d', got '%d'", mockResult.Listeners[0].BackendNodeID, result.Listeners[0].BackendNodeID)
	}

	resultChan = mockSocket.DOMDebugger().GetEventListeners(&domDebugger.GetEventListenersParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
		Depth:    1,
		Pierce:   true,
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveDOMBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveDOMBreakpoint(&domDebugger.RemoveDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   domDebugger.DOMBreakpointType("breakpoint type"),
	})
	mockResult := &domDebugger.RemoveDOMBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveDOMBreakpoint(&domDebugger.RemoveDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   domDebugger.DOMBreakpointType("breakpoint type"),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveEventListenerBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveEventListenerBreakpoint(&domDebugger.RemoveEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockResult := &domDebugger.RemoveEventListenerBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveEventListenerBreakpoint(&domDebugger.RemoveEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveInstrumentationBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveInstrumentationBreakpoint(&domDebugger.RemoveInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockResult := &domDebugger.RemoveInstrumentationBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveInstrumentationBreakpoint(&domDebugger.RemoveInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerRemoveXHRBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveXHRBreakpoint(&domDebugger.RemoveXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockResult := &domDebugger.RemoveXHRBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveXHRBreakpoint(&domDebugger.RemoveXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetDOMBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetDOMBreakpoint(&domDebugger.SetDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   domDebugger.DOMBreakpointType("breakpoint type"),
	})
	mockResult := &domDebugger.SetDOMBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetDOMBreakpoint(&domDebugger.SetDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   domDebugger.DOMBreakpointType("breakpoint type"),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetEventListenerBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetEventListenerBreakpoint(&domDebugger.SetEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockResult := &domDebugger.SetEventListenerBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetEventListenerBreakpoint(&domDebugger.SetEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetInstrumentationBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetInstrumentationBreakpoint(&domDebugger.SetInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockResult := &domDebugger.SetInstrumentationBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetInstrumentationBreakpoint(&domDebugger.SetInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDebuggerSetXHRBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetXHRBreakpoint(&domDebugger.SetXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockResult := &domDebugger.SetXHRBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetXHRBreakpoint(&domDebugger.SetXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
