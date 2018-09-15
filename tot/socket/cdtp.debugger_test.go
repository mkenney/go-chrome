package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/debugger"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestDebuggerContinueToLocation(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.ContinueToLocationResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().ContinueToLocation(&debugger.ContinueToLocationParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		TargetCallFrames: debugger.TargetCallFrames.Any,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().ContinueToLocation(&debugger.ContinueToLocationParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		TargetCallFrames: debugger.TargetCallFrames.Any,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerEvaluateOnCallFrame(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().EvaluateOnCallFrame(&debugger.EvaluateOnCallFrameParams{
		CallFrameID:           debugger.CallFrameID("call-frame-id"),
		Expression:            "expression",
		ObjectGroup:           "object-group",
		IncludeCommandLineAPI: true,
		Silent:                true,
		ReturnByValue:         true,
		GeneratePreview:       true,
		ThrowOnSideEffect:     true,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result.Type != result.Result.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.Result.Type, result.Result.Type)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().EvaluateOnCallFrame(&debugger.EvaluateOnCallFrameParams{
		CallFrameID:           debugger.CallFrameID("call-frame-id"),
		Expression:            "expression",
		ObjectGroup:           "object-group",
		IncludeCommandLineAPI: true,
		Silent:                true,
		ReturnByValue:         true,
		GeneratePreview:       true,
		ThrowOnSideEffect:     true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerGetPossibleBreakpoints(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.GetPossibleBreakpointsResult{
		Locations: []*debugger.BreakLocation{{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
			Type:         debugger.BreakLocationType.DebuggerStatement,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().GetPossibleBreakpoints(&debugger.GetPossibleBreakpointsParams{
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
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Locations[0].ScriptID != result.Locations[0].ScriptID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Locations[0].ScriptID, result.Locations[0].ScriptID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().GetPossibleBreakpoints(&debugger.GetPossibleBreakpointsParams{
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
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerGetScriptSource(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.GetScriptSourceResult{
		ScriptSource: "script-source",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().GetScriptSource(&debugger.GetScriptSourceParams{
		ScriptID: runtime.ScriptID("script-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ScriptSource != result.ScriptSource {
		t.Errorf("Expected '%s', got '%s'", mockResult.ScriptSource, result.ScriptSource)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().GetScriptSource(&debugger.GetScriptSourceParams{
		ScriptID: runtime.ScriptID("script-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerGetStackTrace(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().GetStackTrace(&debugger.GetStackTraceParams{
		StackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.StackTrace.Description != result.StackTrace.Description {
		t.Errorf("Expected '%s', got '%s'", mockResult.StackTrace.Description, result.StackTrace.Description)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().GetStackTrace(&debugger.GetStackTraceParams{
		StackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerPause(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.PauseResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().Pause()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().Pause()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerPauseOnAsyncCall(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.PauseOnAsyncCallResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().PauseOnAsyncCall(&debugger.PauseOnAsyncCallParams{
		ParentStackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().PauseOnAsyncCall(&debugger.PauseOnAsyncCallParams{
		ParentStackTraceID: runtime.StackTraceID{
			ID:         "stack-trace-id",
			DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
		},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerRemoveBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.RemoveBreakpointResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().RemoveBreakpoint(&debugger.RemoveBreakpointParams{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().RemoveBreakpoint(&debugger.RemoveBreakpointParams{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerRestartFrame(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().RestartFrame(&debugger.RestartFrameParams{
		CallFrameID: debugger.CallFrameID("call-frame-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.CallFrames[0].CallFrameID != result.CallFrames[0].CallFrameID {
		t.Errorf("Expected '%s', got '%s'", mockResult.CallFrames[0].CallFrameID, result.CallFrames[0].CallFrameID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().RestartFrame(&debugger.RestartFrameParams{
		CallFrameID: debugger.CallFrameID("call-frame-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerResume(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.ResumeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().Resume()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().Resume()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerScheduleStepIntoAsync(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.ScheduleStepIntoAsyncResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().ScheduleStepIntoAsync()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().ScheduleStepIntoAsync()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSearchInContent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SearchInContentResult{
		Result: []*debugger.SearchMatch{{
			LineNumber:  1,
			LineContent: "line with search string in it",
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SearchInContent(&debugger.SearchInContentParams{
		ScriptID:      runtime.ScriptID("script-id"),
		Query:         "search string",
		CaseSensitive: true,
		IsRegex:       true,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].LineNumber != result.Result[0].LineNumber {
		t.Errorf("Expected '%d', got '%d'", mockResult.Result[0].LineNumber, result.Result[0].LineNumber)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SearchInContent(&debugger.SearchInContentParams{
		ScriptID:      runtime.ScriptID("script-id"),
		Query:         "search string",
		CaseSensitive: true,
		IsRegex:       true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetAsyncCallStackDepth(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetAsyncCallStackDepthResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetAsyncCallStackDepth(&debugger.SetAsyncCallStackDepthParams{
		MaxDepth: 1,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetAsyncCallStackDepth(&debugger.SetAsyncCallStackDepthParams{
		MaxDepth: 1,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetBlackboxPatterns(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetBlackboxPatternsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetBlackboxPatterns(&debugger.SetBlackboxPatternsParams{
		Patterns: []string{"pattern 1", "pattern 2"},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetBlackboxPatterns(&debugger.SetBlackboxPatternsParams{
		Patterns: []string{"pattern 1", "pattern 2"},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetBlackboxedRanges(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetBlackboxedRangesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetBlackboxedRanges(&debugger.SetBlackboxedRangesParams{
		ScriptID: runtime.ScriptID("script-id"),
		Positions: []*debugger.ScriptPosition{{
			LineNumber:   1,
			ColumnNumber: 1,
		}},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetBlackboxedRanges(&debugger.SetBlackboxedRangesParams{
		ScriptID: runtime.ScriptID("script-id"),
		Positions: []*debugger.ScriptPosition{{
			LineNumber:   1,
			ColumnNumber: 1,
		}},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetBreakpoint(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetBreakpointResult{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
		ActualLocation: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   2,
			ColumnNumber: 4,
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetBreakpoint(&debugger.SetBreakpointParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		Condition: "breakpoint-condition",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.BreakpointID != result.BreakpointID {
		t.Errorf("Expected '%s', got '%s'", mockResult.BreakpointID, result.BreakpointID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetBreakpoint(&debugger.SetBreakpointParams{
		Location: &debugger.Location{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   1,
			ColumnNumber: 1,
		},
		Condition: "breakpoint-condition",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetBreakpointByURL(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetBreakpointByURLResult{
		BreakpointID: debugger.BreakpointID("breakpoint-id"),
		Locations: []*debugger.Location{{
			ScriptID:     runtime.ScriptID("script-id"),
			LineNumber:   2,
			ColumnNumber: 4,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetBreakpointByURL(&debugger.SetBreakpointByURLParams{
		LineNumber:   1,
		URL:          "http://some.url",
		URLRegex:     "some regex",
		ScriptHash:   "some hash",
		ColumnNumber: 1,
		Condition:    "some condition",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.BreakpointID != result.BreakpointID {
		t.Errorf("Expected '%s', got '%s'", mockResult.BreakpointID, result.BreakpointID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetBreakpointByURL(&debugger.SetBreakpointByURLParams{
		LineNumber:   1,
		URL:          "http://some.url",
		URLRegex:     "some regex",
		ScriptHash:   "some hash",
		ColumnNumber: 1,
		Condition:    "some condition",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetBreakpointsActive(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetBreakpointsActiveResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetBreakpointsActive(&debugger.SetBreakpointsActiveParams{
		Active: true,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetBreakpointsActive(&debugger.SetBreakpointsActiveParams{
		Active: true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetPauseOnExceptions(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetPauseOnExceptionsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetPauseOnExceptions(&debugger.SetPauseOnExceptionsParams{
		State: debugger.State.None,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetPauseOnExceptions(&debugger.SetPauseOnExceptionsParams{
		State: debugger.State.None,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetReturnValue(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetReturnValueResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetReturnValue(&debugger.SetReturnValueParams{
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetReturnValue(&debugger.SetReturnValueParams{
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetScriptSource(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetScriptSource(&debugger.SetScriptSourceParams{
		ScriptID:     runtime.ScriptID("script-id"),
		ScriptSource: "http://script.source",
		DryRun:       true,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.CallFrames[0].CallFrameID != result.CallFrames[0].CallFrameID {
		t.Errorf("Expected '%s', got '%s'", mockResult.CallFrames[0].CallFrameID, result.CallFrames[0].CallFrameID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetScriptSource(&debugger.SetScriptSourceParams{
		ScriptID:     runtime.ScriptID("script-id"),
		ScriptSource: "http://script.source",
		DryRun:       true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetSkipAllPauses(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetSkipAllPausesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetSkipAllPauses(&debugger.SetSkipAllPausesParams{
		Skip: true,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetSkipAllPauses(&debugger.SetSkipAllPausesParams{
		Skip: true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerSetVariableValue(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.SetVariableValueResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().SetVariableValue(&debugger.SetVariableValueParams{
		ScopeNumber:  1,
		VariableName: "varname",
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
		CallFrameID: debugger.CallFrameID("call-frame-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().SetVariableValue(&debugger.SetVariableValueParams{
		ScopeNumber:  1,
		VariableName: "varname",
		NewValue: &runtime.CallArgument{
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		},
		CallFrameID: debugger.CallFrameID("call-frame-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerStepInto(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.StepIntoResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().StepInto(&debugger.StepIntoParams{
		BreakOnAsyncCall: true,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().StepInto(&debugger.StepIntoParams{
		BreakOnAsyncCall: true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerStepOut(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.StepOutResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().StepOut()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().StepOut()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerStepOver(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &debugger.StepOverResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Debugger().StepOver()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Debugger().StepOver()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnBreakpointResolved(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *debugger.BreakpointResolvedEvent)
	soc.Debugger().OnBreakpointResolved(func(eventData *debugger.BreakpointResolvedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Debugger.breakpointResolved",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Debugger.breakpointResolved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnPaused(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *debugger.PausedEvent)
	soc.Debugger().OnPaused(func(eventData *debugger.PausedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Debugger.paused",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Debugger.paused",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnResumed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *debugger.ResumedEvent)
	soc.Debugger().OnResumed(func(eventData *debugger.ResumedEvent) {
		resultChan <- eventData
	})

	mockResult := &debugger.ResumedEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Debugger.resumed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Debugger.resumed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnScriptFailedToParse(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *debugger.ScriptFailedToParseEvent)
	soc.Debugger().OnScriptFailedToParse(func(eventData *debugger.ScriptFailedToParseEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Debugger.scriptFailedToParse",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Debugger.scriptFailedToParse",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDebuggerOnScriptParsed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *debugger.ScriptParsedEvent)
	soc.Debugger().OnScriptParsed(func(eventData *debugger.ScriptParsedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Debugger.scriptParsed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Debugger.scriptParsed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
