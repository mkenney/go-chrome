package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/dom/debugger"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestDOMDebuggerGetEventListeners(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerGetEventListeners")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().GetEventListeners(&debugger.GetEventListenersParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
		Depth:    1,
		Pierce:   true,
	})
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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

	resultChan = mockSocket.DOMDebugger().GetEventListeners(&debugger.GetEventListenersParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
		Depth:    1,
		Pierce:   true,
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerRemoveDOMBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveDOMBreakpoint(&debugger.RemoveDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	mockResult := &debugger.RemoveDOMBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveDOMBreakpoint(&debugger.RemoveDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerRemoveEventListenerBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveEventListenerBreakpoint(&debugger.RemoveEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockResult := &debugger.RemoveEventListenerBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveEventListenerBreakpoint(&debugger.RemoveEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerRemoveInstrumentationBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveInstrumentationBreakpoint(&debugger.RemoveInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockResult := &debugger.RemoveInstrumentationBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveInstrumentationBreakpoint(&debugger.RemoveInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerRemoveXHRBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().RemoveXHRBreakpoint(&debugger.RemoveXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockResult := &debugger.RemoveXHRBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().RemoveXHRBreakpoint(&debugger.RemoveXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerSetDOMBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetDOMBreakpoint(&debugger.SetDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	mockResult := &debugger.SetDOMBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetDOMBreakpoint(&debugger.SetDOMBreakpointParams{
		NodeID: dom.NodeID(1),
		Type:   debugger.DOMBreakpointType("breakpoint type"),
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerSetEventListenerBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetEventListenerBreakpoint(&debugger.SetEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockResult := &debugger.SetEventListenerBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetEventListenerBreakpoint(&debugger.SetEventListenerBreakpointParams{
		EventName:  "event name",
		TargetName: "target name",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerSetInstrumentationBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetInstrumentationBreakpoint(&debugger.SetInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockResult := &debugger.SetInstrumentationBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetInstrumentationBreakpoint(&debugger.SetInstrumentationBreakpointParams{
		EventName: "event name",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestDOMDebuggerSetXHRBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMDebugger().SetXHRBreakpoint(&debugger.SetXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockResult := &debugger.SetXHRBreakpointResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMDebugger().SetXHRBreakpoint(&debugger.SetXHRBreakpointParams{
		URL: "http://xhr.url",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
