package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/debugger"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestDebuggerContinueToLocation(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerContinueToLocation")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().ContinueToLocation(&debugger.ContinueToLocationParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		TargetCallFrames: debugger.TargetCallFrames.Any,
	})
	mockResult := &debugger.ContinueToLocationResult{}
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

	resultChan = mockSocket.Debugger().ContinueToLocation(&debugger.ContinueToLocationParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		TargetCallFrames: debugger.TargetCallFrames.Any,
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

func TestDebuggerDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerDisable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().Disable()
	mockResult := &debugger.DisableResult{}
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

	resultChan = mockSocket.Debugger().Disable()
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

func TestDebuggerEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerEnable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().Enable()
	mockResult := &debugger.EnableResult{}
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

	resultChan = mockSocket.Debugger().Enable()
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

func TestDebuggerEvaluateOnCallFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerEvaluateOnCallFrame")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().EvaluateOnCallFrame(&debugger.EvaluateOnCallFrameParams{
		CallFrameID:           debugger.CallFrameID("call-frame-id"),
		Expression:            "expression",
		ObjectGroup:           "object-group",
		IncludeCommandLineAPI: true,
		Silent:                true,
		ReturnByValue:         true,
		GeneratePreview:       true,
		ThrowOnSideEffect:     true,
	})
	mockResult := &debugger.EvaluateOnCallFrameResult{
		Result: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			Subtype:             runtime.ObjectSubtype.Array,
			ClassName:           "class-name",
			Value:               1,
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "description",
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
			Preview: &runtime.ObjectPreview{
				Type:        runtime.ObjectType.Object,
				Subtype:     runtime.ObjectSubtype.Array,
				Description: "description",
				Overflow:    true,
				Properties:  []*runtime.PropertyPreview{},
				Entries:     []*runtime.EntryPreview{},
			},
			CustomPreview: &runtime.CustomPreview{},
		},
		ExceptionDetails: &runtime.ExceptionDetails{
			ExceptionID:  1,
			Text:         "exception text",
			LineNumber:   1,
			ColumnNumber: 1,
			ScriptID:     runtime.ScriptID("script-id"),
			URL:          "http://exception.url",
			StackTrace:   &runtime.StackTrace{},
			Exception: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ExecutionContextID: runtime.ExecutionContextID(1),
		},
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
	if mockResult.Result.Type != result.Result.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.Result.Type, result.Result.Type)
	}

	resultChan = mockSocket.Debugger().EvaluateOnCallFrame(&debugger.EvaluateOnCallFrameParams{
		CallFrameID:           debugger.CallFrameID("call-frame-id"),
		Expression:            "expression",
		ObjectGroup:           "object-group",
		IncludeCommandLineAPI: true,
		Silent:                true,
		ReturnByValue:         true,
		GeneratePreview:       true,
		ThrowOnSideEffect:     true,
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

func TestDebuggerGetPossibleBreakpoints(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerGetPossibleBreakpoints")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().GetPossibleBreakpoints(&debugger.GetPossibleBreakpointsParams{
		Start: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		End: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   2,
			ColumnNumber: 2,
		},
		RestrictToFunction: true,
	})
	mockResult := &debugger.GetPossibleBreakpointsResult{
		Locations: []*debugger.BreakLocation{{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
			Type:         debugger.BreakLocationType.DebuggerStatement,
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
	if mockResult.Locations[0].ScriptID != result.Locations[0].ScriptID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Locations[0].ScriptID, result.Locations[0].ScriptID)
	}

	resultChan = mockSocket.Debugger().GetPossibleBreakpoints(&debugger.GetPossibleBreakpointsParams{
		Start: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		End: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   2,
			ColumnNumber: 2,
		},
		RestrictToFunction: true,
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

func TestDebuggerGetScriptSource(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerGetScriptSource")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().GetScriptSource(&debugger.GetScriptSourceParams{
		ScriptID: runtime.ScriptID("script-id"),
	})
	mockResult := &debugger.GetScriptSourceResult{
		ScriptSource: "script-source",
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
	if mockResult.ScriptSource != result.ScriptSource {
		t.Errorf("Expected '%s', got '%s'", mockResult.ScriptSource, result.ScriptSource)
	}

	resultChan = mockSocket.Debugger().GetScriptSource(&debugger.GetScriptSourceParams{
		ScriptID: runtime.ScriptID("script-id"),
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

func TestDebuggerGetStackTrace(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerGetStackTrace")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().GetStackTrace(&debugger.GetStackTraceParams{
		StackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
	})
	mockResult := &debugger.GetStackTraceResult{
		StackTrace: &runtime.StackTrace{
			Description: "description",
			CallFrames:  []*runtime.CallFrame{{}},
			Parent:      &runtime.StackTrace{},
			ParentID: &runtime.StackTraceID{
				ID:         "parent-stack-trace-id",
				DebuggerID: runtime.UniqueDebuggerID("parent-unique-debugger-id"),
			},
		},
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
	if mockResult.StackTrace.Description != result.StackTrace.Description {
		t.Errorf("Expected '%s', got '%s'", mockResult.StackTrace.Description, result.StackTrace.Description)
	}

	resultChan = mockSocket.Debugger().GetStackTrace(&debugger.GetStackTraceParams{
		StackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
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

func TestDebuggerPause(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerPause")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().Pause()
	mockResult := &debugger.PauseResult{}
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

	resultChan = mockSocket.Debugger().Pause()
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

func TestDebuggerPauseOnAsyncCall(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerPauseOnAsyncCall")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().PauseOnAsyncCall(&debugger.PauseOnAsyncCallParams{
		ParentStackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
	})
	mockResult := &debugger.PauseOnAsyncCallResult{}
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

	resultChan = mockSocket.Debugger().PauseOnAsyncCall(&debugger.PauseOnAsyncCallParams{
		ParentStackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
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

func TestDebuggerRemoveBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerRemoveBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().RemoveBreakpoint(&debugger.RemoveBreakpointParams{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
	})
	mockResult := &debugger.RemoveBreakpointResult{}
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

	resultChan = mockSocket.Debugger().RemoveBreakpoint(&debugger.RemoveBreakpointParams{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
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

func TestDebuggerRestartFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerRestartFrame")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().RestartFrame(&debugger.RestartFrameParams{
		CallFrameID: debugger.CallFrameID("call-frame-id"),
	})
	mockResult := &debugger.RestartFrameResult{
		CallFrames: []*debugger.CallFrame{{
			CallFrameID:  debugger.CallFrameID("call-frame-id"),
			FunctionName: "someFunc",
			FunctionLocation: &debugger.Location{
				ScriptID:     runtime.ScriptID("script-id"),
				LineNumber:   1,
				ColumnNumber: 1,
			},
			Location: &debugger.Location{
				ScriptID:     runtime.ScriptID("script-id"),
				LineNumber:   2,
				ColumnNumber: 4,
			},
			URL: "http://frame.url",
			ScopeChain: []*debugger.Scope{{
				Type: debugger.ScopeType.Global,
				Object: &runtime.RemoteObject{
					Type: runtime.ObjectType.Accessor,
				},
				Name: "scope-name",
				StartLocation: &debugger.Location{
					ScriptID:     runtime.ScriptID("script-id"),
					LineNumber:   2,
					ColumnNumber: 4,
				},
				EndLocation: &debugger.Location{
					ScriptID:     runtime.ScriptID("script-id"),
					LineNumber:   2,
					ColumnNumber: 4,
				},
			}},
			This: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ReturnValue: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
		}},
		AsyncStackTrace: &runtime.StackTrace{},
		AsyncStackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
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
	if mockResult.CallFrames[0].CallFrameID != result.CallFrames[0].CallFrameID {
		t.Errorf("Expected '%s', got '%s'", mockResult.CallFrames[0].CallFrameID, result.CallFrames[0].CallFrameID)
	}

	resultChan = mockSocket.Debugger().RestartFrame(&debugger.RestartFrameParams{
		CallFrameID: debugger.CallFrameID("call-frame-id"),
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

func TestDebuggerResume(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerResume")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().Resume()
	mockResult := &debugger.ResumeResult{}
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

	resultChan = mockSocket.Debugger().Resume()
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

func TestDebuggerScheduleStepIntoAsync(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerScheduleStepIntoAsync")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().ScheduleStepIntoAsync()
	mockResult := &debugger.ScheduleStepIntoAsyncResult{}
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

	resultChan = mockSocket.Debugger().ScheduleStepIntoAsync()
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

func TestDebuggerSearchInContent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSearchInContent")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SearchInContent(&debugger.SearchInContentParams{
		ScriptID:      runtime.ScriptID("script-id"),
		Query:         "search string",
		CaseSensitive: true,
		IsRegex:       true,
	})
	mockResult := &debugger.SearchInContentResult{
		Result: []*debugger.SearchMatch{{
			LineNumber:  1,
			LineContent: "line with search string in it",
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
	if mockResult.Result[0].LineNumber != result.Result[0].LineNumber {
		t.Errorf("Expected '%d', got '%d'", mockResult.Result[0].LineNumber, result.Result[0].LineNumber)
	}

	resultChan = mockSocket.Debugger().SearchInContent(&debugger.SearchInContentParams{
		ScriptID:      runtime.ScriptID("script-id"),
		Query:         "search string",
		CaseSensitive: true,
		IsRegex:       true,
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

func TestDebuggerSetAsyncCallStackDepth(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetAsyncCallStackDepth")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetAsyncCallStackDepth(&debugger.SetAsyncCallStackDepthParams{
		MaxDepth: 1,
	})
	mockResult := &debugger.SetAsyncCallStackDepthResult{}
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

	resultChan = mockSocket.Debugger().SetAsyncCallStackDepth(&debugger.SetAsyncCallStackDepthParams{
		MaxDepth: 1,
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

func TestDebuggerSetBlackboxPatterns(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetBlackboxPatterns")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetBlackboxPatterns(&debugger.SetBlackboxPatternsParams{
		Patterns: []string{"pattern 1", "pattern 2"},
	})
	mockResult := &debugger.SetBlackboxPatternsResult{}
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

	resultChan = mockSocket.Debugger().SetBlackboxPatterns(&debugger.SetBlackboxPatternsParams{
		Patterns: []string{"pattern 1", "pattern 2"},
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

func TestDebuggerSetBlackboxedRanges(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetBlackboxedRanges")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetBlackboxedRanges(&debugger.SetBlackboxedRangesParams{
		ScriptID: runtime.ScriptID("script-id"),
		Positions: []*debugger.ScriptPosition{{
			LineNumber:   1,
			ColumnNumber: 1,
		}},
	})
	mockResult := &debugger.SetBlackboxedRangesResult{}
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

	resultChan = mockSocket.Debugger().SetBlackboxedRanges(&debugger.SetBlackboxedRangesParams{
		ScriptID: runtime.ScriptID("script-id"),
		Positions: []*debugger.ScriptPosition{{
			LineNumber:   1,
			ColumnNumber: 1,
		}},
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

func TestDebuggerSetBreakpoint(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetBreakpoint")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetBreakpoint(&debugger.SetBreakpointParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		Condition: "breakpoint-condition",
	})
	mockResult := &debugger.SetBreakpointResult{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
		ActualLocation: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   2,
			ColumnNumber: 4,
		},
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
	if mockResult.BreakpointID != result.BreakpointID {
		t.Errorf("Expected '%s', got '%s'", mockResult.BreakpointID, result.BreakpointID)
	}

	resultChan = mockSocket.Debugger().SetBreakpoint(&debugger.SetBreakpointParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		Condition: "breakpoint-condition",
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

func TestDebuggerSetBreakpointByURL(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetBreakpointByURL")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetBreakpointByURL(&debugger.SetBreakpointByURLParams{
		LineNumber:   1,
		URL:          "http://some.url",
		URLRegex:     "some regex",
		ScriptHash:   "some hash",
		ColumnNumber: 1,
		Condition:    "some condition",
	})
	mockResult := &debugger.SetBreakpointByURLResult{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
		Locations: []*debugger.Location{{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   2,
			ColumnNumber: 4,
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
	if mockResult.BreakpointID != result.BreakpointID {
		t.Errorf("Expected '%s', got '%s'", mockResult.BreakpointID, result.BreakpointID)
	}

	resultChan = mockSocket.Debugger().SetBreakpointByURL(&debugger.SetBreakpointByURLParams{
		LineNumber:   1,
		URL:          "http://some.url",
		URLRegex:     "some regex",
		ScriptHash:   "some hash",
		ColumnNumber: 1,
		Condition:    "some condition",
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

func TestDebuggerSetBreakpointsActive(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetBreakpointsActive")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetBreakpointsActive(&debugger.SetBreakpointsActiveParams{
		Active: true,
	})
	mockResult := &debugger.SetBreakpointsActiveResult{}
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

	resultChan = mockSocket.Debugger().SetBreakpointsActive(&debugger.SetBreakpointsActiveParams{
		Active: true,
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

func TestDebuggerSetPauseOnExceptions(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetPauseOnExceptions")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetPauseOnExceptions(&debugger.SetPauseOnExceptionsParams{
		State: debugger.State.None,
	})
	mockResult := &debugger.SetPauseOnExceptionsResult{}
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

	resultChan = mockSocket.Debugger().SetPauseOnExceptions(&debugger.SetPauseOnExceptionsParams{
		State: debugger.State.None,
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

func TestDebuggerSetReturnValue(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetReturnValue")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetReturnValue(&debugger.SetReturnValueParams{
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
	})
	mockResult := &debugger.SetReturnValueResult{}
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

	resultChan = mockSocket.Debugger().SetReturnValue(&debugger.SetReturnValueParams{
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
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

func TestDebuggerSetScriptSource(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetScriptSource")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetScriptSource(&debugger.SetScriptSourceParams{
		ScriptID:     runtime.ScriptID("script-id"),
		ScriptSource: "http://script.source",
		DryRun:       true,
	})
	mockResult := &debugger.SetScriptSourceResult{
		CallFrames: []*debugger.CallFrame{{
			CallFrameID:  debugger.CallFrameID("call-frame-id"),
			FunctionName: "someFunc",
			FunctionLocation: &debugger.Location{
				ScriptID:     runtime.ScriptID("script-id"),
				LineNumber:   1,
				ColumnNumber: 1,
			},
			Location: &debugger.Location{
				ScriptID:     runtime.ScriptID("script-id"),
				LineNumber:   2,
				ColumnNumber: 4,
			},
			URL: "http://frame.url",
			ScopeChain: []*debugger.Scope{{
				Type: debugger.ScopeType.Global,
				Object: &runtime.RemoteObject{
					Type: runtime.ObjectType.Accessor,
				},
				Name: "scope-name",
				StartLocation: &debugger.Location{
					ScriptID:     runtime.ScriptID("script-id"),
					LineNumber:   2,
					ColumnNumber: 4,
				},
				EndLocation: &debugger.Location{
					ScriptID:     runtime.ScriptID("script-id"),
					LineNumber:   2,
					ColumnNumber: 4,
				},
			}},
			This: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ReturnValue: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
		}},
		StackChanged: true,
		AsyncStackTrace: &runtime.StackTrace{
			Description: "description",
			CallFrames:  []*runtime.CallFrame{{}},
			Parent:      &runtime.StackTrace{},
			ParentID: &runtime.StackTraceID{
				ID:         "parent-stack-trace-id",
				DebuggerID: runtime.UniqueDebuggerID("parent-unique-debugger-id"),
			},
		},
		AsyncStackTraceID: runtime.StackTraceID{
			ID:         "stacktrace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
		ExceptionDetails: &runtime.ExceptionDetails{
			ExceptionID:  1,
			Text:         "exception text",
			LineNumber:   1,
			ColumnNumber: 1,
			ScriptID:     runtime.ScriptID("script-id"),
			URL:          "http://exception.url",
			StackTrace:   &runtime.StackTrace{},
			Exception: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ExecutionContextID: runtime.ExecutionContextID(1),
		},
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
	if mockResult.CallFrames[0].CallFrameID != result.CallFrames[0].CallFrameID {
		t.Errorf("Expected '%s', got '%s'", mockResult.CallFrames[0].CallFrameID, result.CallFrames[0].CallFrameID)
	}

	resultChan = mockSocket.Debugger().SetScriptSource(&debugger.SetScriptSourceParams{
		ScriptID:     runtime.ScriptID("script-id"),
		ScriptSource: "http://script.source",
		DryRun:       true,
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

func TestDebuggerSetSkipAllPauses(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetSkipAllPauses")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetSkipAllPauses(&debugger.SetSkipAllPausesParams{
		Skip: true,
	})
	mockResult := &debugger.SetSkipAllPausesResult{}
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

	resultChan = mockSocket.Debugger().SetSkipAllPauses(&debugger.SetSkipAllPausesParams{
		Skip: true,
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

func TestDebuggerSetVariableValue(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerSetVariableValue")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().SetVariableValue(&debugger.SetVariableValueParams{
		ScopeNumber:  1,
		VariableName: "varname",
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
		CallFrameID: debugger.CallFrameID("call-frame-id"),
	})
	mockResult := &debugger.SetVariableValueResult{}
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

	resultChan = mockSocket.Debugger().SetVariableValue(&debugger.SetVariableValueParams{
		ScopeNumber:  1,
		VariableName: "varname",
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
		CallFrameID: debugger.CallFrameID("call-frame-id"),
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

func TestDebuggerStepInto(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerStepInto")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().StepInto(&debugger.StepIntoParams{
		BreakOnAsyncCall: true,
	})
	mockResult := &debugger.StepIntoResult{}
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

	resultChan = mockSocket.Debugger().StepInto(&debugger.StepIntoParams{
		BreakOnAsyncCall: true,
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

func TestDebuggerStepOut(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerStepOut")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().StepOut()
	mockResult := &debugger.StepOutResult{}
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

	resultChan = mockSocket.Debugger().StepOut()
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

func TestDebuggerStepOver(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerStepOver")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Debugger().StepOver()
	mockResult := &debugger.StepOverResult{}
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

	resultChan = mockSocket.Debugger().StepOver()
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

func TestDebuggerOnBreakpointResolved(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerOnBreakpointResolved")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *debugger.BreakpointResolvedEvent)
	mockSocket.Debugger().OnBreakpointResolved(func(eventData *debugger.BreakpointResolvedEvent) {
		resultChan <- eventData
	})
	mockResult := &debugger.BreakpointResolvedEvent{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Debugger.breakpointResolved",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *debugger.BreakpointResolvedEvent)
	mockSocket.Debugger().OnBreakpointResolved(func(eventData *debugger.BreakpointResolvedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Debugger.breakpointResolved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnPaused(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerOnPaused")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *debugger.PausedEvent)
	mockSocket.Debugger().OnPaused(func(eventData *debugger.PausedEvent) {
		resultChan <- eventData
	})
	mockResult := &debugger.PausedEvent{
		CallFrames: []*debugger.CallFrame{{
			CallFrameID: debugger.CallFrameID("call-frame-id"),
		}},
		Reason:                "reason",
		Data:                  map[string]string{"key": "value"},
		HitBreakpoints:        []string{"breakpoint1"},
		AsyncStackTrace:       &runtime.StackTrace{},
		AsyncStackTraceID:     &runtime.StackTraceID{},
		AsyncCallStackTraceID: &runtime.StackTraceID{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Debugger.paused",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *debugger.PausedEvent)
	mockSocket.Debugger().OnPaused(func(eventData *debugger.PausedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Debugger.paused",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnResumed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerOnResumed")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *debugger.ResumedEvent)
	mockSocket.Debugger().OnResumed(func(eventData *debugger.ResumedEvent) {
		resultChan <- eventData
	})
	mockResult := &debugger.ResumedEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Debugger.resumed",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *debugger.ResumedEvent)
	mockSocket.Debugger().OnResumed(func(eventData *debugger.ResumedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Debugger.resumed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnScriptFailedToParse(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerOnScriptFailedToParse")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *debugger.ScriptFailedToParseEvent)
	mockSocket.Debugger().OnScriptFailedToParse(func(eventData *debugger.ScriptFailedToParseEvent) {
		resultChan <- eventData
	})
	mockResult := &debugger.ScriptFailedToParseEvent{
		ScriptID:                runtime.ScriptID("script-id"),
		URL:                     "http://script.url",
		StartLine:               1,
		StartColumn:             1,
		EndLine:                 2,
		EndColumn:               10,
		ExecutionContextID:      runtime.ExecutionContextID(1),
		Hash:                    "some hash",
		ExecutionContextAuxData: map[string]string{"key": "value"},
		SourceMapURL:            "http://source-map.url",
		HasSourceURL:            true,
		IsModule:                true,
		Length:                  1,
		StackTrace:              &runtime.StackTrace{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Debugger.scriptFailedToParse",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *debugger.ScriptFailedToParseEvent)
	mockSocket.Debugger().OnScriptFailedToParse(func(eventData *debugger.ScriptFailedToParseEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Debugger.scriptFailedToParse",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnScriptParsed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDebuggerOnScriptParsed")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *debugger.ScriptParsedEvent)
	mockSocket.Debugger().OnScriptParsed(func(eventData *debugger.ScriptParsedEvent) {
		resultChan <- eventData
	})
	mockResult := &debugger.ScriptParsedEvent{
		ScriptID:                runtime.ScriptID("script-id"),
		URL:                     "http://script.url",
		StartLine:               1,
		StartColumn:             1,
		EndLine:                 2,
		EndColumn:               10,
		ExecutionContextID:      runtime.ExecutionContextID(1),
		Hash:                    "some hash",
		ExecutionContextAuxData: map[string]string{"key": "value"},
		IsLiveEdit:              true,
		SourceMapURL:            "http://source-map.url",
		HasSourceURL:            true,
		IsModule:                true,
		Length:                  1,
		StackTrace:              &runtime.StackTrace{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Debugger.scriptParsed",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *debugger.ScriptParsedEvent)
	mockSocket.Debugger().OnScriptParsed(func(eventData *debugger.ScriptParsedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Debugger.scriptParsed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
