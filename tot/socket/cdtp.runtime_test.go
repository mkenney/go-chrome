package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestRuntimeAwaitPromise(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.AwaitPromiseParams{
		PromiseObjectID: runtime.RemoteObjectID("remote-object-id"),
		ReturnByValue:   true,
		GeneratePreview: true,
	}
	mockResult := &runtime.AwaitPromiseResult{
		Result: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			Subtype:             runtime.ObjectSubtype.Array,
			ClassName:           "class-name",
			Value:               "value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "description",
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
			Preview: &runtime.ObjectPreview{
				Type:        runtime.ObjectType.Object,
				Subtype:     runtime.ObjectSubtype.Array,
				Description: "description",
				Overflow:    true,
				Properties: []*runtime.PropertyPreview{{
					Name:  "name",
					Type:  runtime.ObjectType.Object,
					Value: "value",
					ValuePreview: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Subtype: runtime.ObjectSubtype.Array,
				}},
				Entries: []*runtime.EntryPreview{{
					Key: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Value: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
				}},
			},
			CustomPreview: &runtime.CustomPreview{
				Header:                     "header",
				HasBody:                    true,
				FormatterObjectID:          runtime.RemoteObjectID("remote-object-id"),
				BindRemoteObjectFunctionID: runtime.RemoteObjectID("remote-object-id"),
				ConfigObjectID:             runtime.RemoteObjectID("remote-object-id"),
			},
		},
		ExceptionDetails: &runtime.ExceptionDetails{
			ExceptionID:  1,
			Text:         "text",
			LineNumber:   1,
			ColumnNumber: 1,
			ScriptID:     runtime.ScriptID("script-id"),
			URL:          "http://some.url",
			StackTrace: &runtime.StackTrace{
				Description: "description",
				CallFrames: []*runtime.CallFrame{{
					FunctionName: "functionName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://some.url",
					LineNumber:   1,
					ColumnNumber: 1,
				}},
				Parent: &runtime.StackTrace{},
				ParentID: &runtime.StackTraceID{
					ID:         "stack-trace-id",
					DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
				},
			},
			Exception: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ExecutionContextID: runtime.ExecutionContextID(1),
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().AwaitPromise(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result.Type != result.Result.Type {
		t.Errorf("Expected %s, got %s", mockResult.Result.Type, result.Result.Type)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().AwaitPromise(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeCallFunctionOn(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.CallFunctionOnParams{
		FunctionDeclaration: "function(){}",
		ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		Arguments: []*runtime.CallArgument{{
			Value:               "value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
		}},
		Silent:             true,
		ReturnByValue:      true,
		GeneratePreview:    true,
		UserGesture:        true,
		AwaitPromise:       true,
		ExecutionContextID: runtime.ExecutionContextID(1),
		ObjectGroup:        "object group",
	}
	mockResult := &runtime.CallFunctionOnResult{
		Result: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			Subtype:             runtime.ObjectSubtype.Array,
			ClassName:           "class-name",
			Value:               "value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "description",
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
			Preview: &runtime.ObjectPreview{
				Type:        runtime.ObjectType.Object,
				Subtype:     runtime.ObjectSubtype.Array,
				Description: "description",
				Overflow:    true,
				Properties: []*runtime.PropertyPreview{{
					Name:  "name",
					Type:  runtime.ObjectType.Object,
					Value: "value",
					ValuePreview: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Subtype: runtime.ObjectSubtype.Array,
				}},
				Entries: []*runtime.EntryPreview{{
					Key: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Value: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
				}},
			},
			CustomPreview: &runtime.CustomPreview{
				Header:                     "header",
				HasBody:                    true,
				FormatterObjectID:          runtime.RemoteObjectID("remote-object-id"),
				BindRemoteObjectFunctionID: runtime.RemoteObjectID("remote-object-id"),
				ConfigObjectID:             runtime.RemoteObjectID("remote-object-id"),
			},
		},
		ExceptionDetails: &runtime.ExceptionDetails{
			ExceptionID:  1,
			Text:         "text",
			LineNumber:   1,
			ColumnNumber: 1,
			ScriptID:     runtime.ScriptID("script-id"),
			URL:          "http://some.url",
			StackTrace: &runtime.StackTrace{
				Description: "description",
				CallFrames: []*runtime.CallFrame{{
					FunctionName: "functionName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://some.url",
					LineNumber:   1,
					ColumnNumber: 1,
				}},
				Parent: &runtime.StackTrace{},
				ParentID: &runtime.StackTraceID{
					ID:         "stack-trace-id",
					DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
				},
			},
			Exception: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ExecutionContextID: runtime.ExecutionContextID(1),
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().CallFunctionOn(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result.Type != result.Result.Type {
		t.Errorf("Expected %s, got %s", mockResult.Result.Type, result.Result.Type)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().CallFunctionOn(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeCompileScript(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.CompileScriptParams{
		Expression:         "expression",
		SourceURL:          "http://some.url",
		PersistScript:      true,
		ExecutionContextID: runtime.ExecutionContextID(1),
	}
	mockResult := &runtime.CompileScriptResult{
		ScriptID: runtime.ScriptID("script-id"),
		ExceptionDetails: &runtime.ExceptionDetails{
			ExceptionID:  1,
			Text:         "text",
			LineNumber:   1,
			ColumnNumber: 1,
			ScriptID:     runtime.ScriptID("script-id"),
			URL:          "http://some.url",
			StackTrace: &runtime.StackTrace{
				Description: "description",
				CallFrames: []*runtime.CallFrame{{
					FunctionName: "functionName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://some.url",
					LineNumber:   1,
					ColumnNumber: 1,
				}},
				Parent: &runtime.StackTrace{},
				ParentID: &runtime.StackTraceID{
					ID:         "stack-trace-id",
					DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
				},
			},
			Exception: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ExecutionContextID: runtime.ExecutionContextID(1),
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().CompileScript(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ScriptID != result.ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.ScriptID, result.ScriptID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().CompileScript(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &runtime.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeDiscardConsoleEntries(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &runtime.DiscardConsoleEntriesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().DiscardConsoleEntries()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().DiscardConsoleEntries()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &runtime.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeEvaluate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.EvaluateParams{
		Expression:            "expression",
		ObjectGroup:           "object-group",
		IncludeCommandLineAPI: true,
		Silent:                true,
		ContextID:             runtime.ExecutionContextID(1),
		ReturnByValue:         true,
		GeneratePreview:       true,
		UserGesture:           true,
		AwaitPromise:          true,
	}
	mockResult := &runtime.EvaluateResult{
		Result: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			Subtype:             runtime.ObjectSubtype.Array,
			ClassName:           "class-name",
			Value:               "value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "description",
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
			Preview: &runtime.ObjectPreview{
				Type:        runtime.ObjectType.Object,
				Subtype:     runtime.ObjectSubtype.Array,
				Description: "description",
				Overflow:    true,
				Properties: []*runtime.PropertyPreview{{
					Name:  "name",
					Type:  runtime.ObjectType.Object,
					Value: "value",
					ValuePreview: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Subtype: runtime.ObjectSubtype.Array,
				}},
				Entries: []*runtime.EntryPreview{{
					Key: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Value: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
				}},
			},
			CustomPreview: &runtime.CustomPreview{
				Header:                     "header",
				HasBody:                    true,
				FormatterObjectID:          runtime.RemoteObjectID("remote-object-id"),
				BindRemoteObjectFunctionID: runtime.RemoteObjectID("remote-object-id"),
				ConfigObjectID:             runtime.RemoteObjectID("remote-object-id"),
			},
		},
		ExceptionDetails: &runtime.ExceptionDetails{
			ExceptionID:  1,
			Text:         "text",
			LineNumber:   1,
			ColumnNumber: 1,
			ScriptID:     runtime.ScriptID("script-id"),
			URL:          "http://some.url",
			StackTrace: &runtime.StackTrace{
				Description: "description",
				CallFrames: []*runtime.CallFrame{{
					FunctionName: "functionName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://some.url",
					LineNumber:   1,
					ColumnNumber: 1,
				}},
				Parent: &runtime.StackTrace{},
				ParentID: &runtime.StackTraceID{
					ID:         "stack-trace-id",
					DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
				},
			},
			Exception: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ExecutionContextID: runtime.ExecutionContextID(1),
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().Evaluate(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result.Type != result.Result.Type {
		t.Errorf("Expected %s, got %s", mockResult.Result.Type, result.Result.Type)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().Evaluate(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeGetProperties(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.GetPropertiesParams{
		ObjectID:               runtime.RemoteObjectID("remote-object-id"),
		OwnProperties:          true,
		AccessorPropertiesOnly: true,
		GeneratePreview:        true,
	}
	mockResult := &runtime.GetPropertiesResult{
		Result: []*runtime.PropertyDescriptor{{
			Name: "name",
			Value: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			Writable: true,
			Get: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			Set: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			Configurable: true,
			Enumerable:   true,
			WasThrown:    true,
			IsOwn:        true,
			Symbol: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
		}},
		InternalProperties: []*runtime.InternalPropertyDescriptor{{
			Name: "name",
			Value: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
		}},
		ExceptionDetails: &runtime.ExceptionDetails{
			ExceptionID:  1,
			Text:         "text",
			LineNumber:   1,
			ColumnNumber: 1,
			ScriptID:     runtime.ScriptID("script-id"),
			URL:          "http://some.url",
			StackTrace: &runtime.StackTrace{
				Description: "description",
				CallFrames: []*runtime.CallFrame{{
					FunctionName: "functionName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://some.url",
					LineNumber:   1,
					ColumnNumber: 1,
				}},
				Parent: &runtime.StackTrace{},
				ParentID: &runtime.StackTraceID{
					ID:         "stack-trace-id",
					DebuggerID: runtime.UniqueDebuggerID("unique-debugger-id"),
				},
			},
			Exception: &runtime.RemoteObject{
				Type: runtime.ObjectType.Accessor,
			},
			ExecutionContextID: runtime.ExecutionContextID(1),
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().GetProperties(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].Name != result.Result[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].Name, result.Result[0].Name)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().GetProperties(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeGlobalLexicalScopeNames(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.GlobalLexicalScopeNamesParams{
		ExecutionContextID: runtime.ExecutionContextID(1),
	}
	mockResult := &runtime.GlobalLexicalScopeNamesResult{
		Names: []string{"name1", "name2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().GlobalLexicalScopeNames(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Names[0] != result.Names[0] {
		t.Errorf("Expected %s, got %s", mockResult.Names[0], result.Names[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().GlobalLexicalScopeNames(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeQueryObjects(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.QueryObjectsParams{
		PrototypeObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &runtime.QueryObjectsResult{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().QueryObjects(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ObjectID != result.ObjectID {
		t.Errorf("Expected %s, got %s", mockResult.ObjectID, result.ObjectID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().QueryObjects(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeReleaseObject(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.ReleaseObjectParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &runtime.ReleaseObjectResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().ReleaseObject(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().ReleaseObject(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeReleaseObjectGroup(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.ReleaseObjectGroupParams{
		ObjectGroup: "object-group",
	}
	mockResult := &runtime.ReleaseObjectGroupResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().ReleaseObjectGroup(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().ReleaseObjectGroup(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeRunIfWaitingForDebugger(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &runtime.RunIfWaitingForDebuggerResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().RunIfWaitingForDebugger()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().RunIfWaitingForDebugger()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeRunScript(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.RunScriptParams{
		ScriptID:              runtime.ScriptID("script-id"),
		ExecutionContextID:    runtime.ExecutionContextID(1),
		ObjectGroup:           "object-group",
		Silent:                true,
		IncludeCommandLineAPI: true,
		ReturnByValue:         true,
		GeneratePreview:       true,
		AwaitPromise:          true,
	}
	mockResult := &runtime.RunScriptResult{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().RunScript(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ObjectID != result.ObjectID {
		t.Errorf("Expected %s, got %s", mockResult.ObjectID, result.ObjectID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().RunScript(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeSetCustomObjectFormatterEnabled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &runtime.SetCustomObjectFormatterEnabledParams{
		Result: &runtime.RemoteObject{
			Type: runtime.ObjectType.Accessor,
		},
		ExceptionDetails: &runtime.ExceptionDetails{},
	}
	mockResult := &runtime.SetCustomObjectFormatterEnabledResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Runtime().SetCustomObjectFormatterEnabled(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Runtime().SetCustomObjectFormatterEnabled(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnConsoleAPICalled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *runtime.ConsoleAPICalledEvent)
	soc.Runtime().OnConsoleAPICalled(func(eventData *runtime.ConsoleAPICalledEvent) {
		resultChan <- eventData
	})

	mockResult := &runtime.ConsoleAPICalledEvent{
		Type:               runtime.CallType.Assert,
		Args:               []*runtime.RemoteObject{{}},
		ExecutionContextID: runtime.ExecutionContextID(1),
		Timestamp:          runtime.Timestamp(time.Now().Unix()),
		StackTrace:         &runtime.StackTrace{},
		Context:            "context",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Runtime.consoleAPICalled",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Type != result.Type {
		t.Errorf("Expected %s, got %s", mockResult.Type, result.Type)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Runtime.consoleAPICalled",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExceptionRevoked(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *runtime.ExceptionRevokedEvent)
	soc.Runtime().OnExceptionRevoked(func(eventData *runtime.ExceptionRevokedEvent) {
		resultChan <- eventData
	})

	mockResult := &runtime.ExceptionRevokedEvent{
		Reason:      "reason",
		ExceptionID: 1,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Runtime.exceptionRevoked",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Reason != result.Reason {
		t.Errorf("Expected %s, got %s", mockResult.Reason, result.Reason)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Runtime.exceptionRevoked",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExceptionThrown(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *runtime.ExceptionThrownEvent)
	soc.Runtime().OnExceptionThrown(func(eventData *runtime.ExceptionThrownEvent) {
		resultChan <- eventData
	})

	mockResult := &runtime.ExceptionThrownEvent{
		Timestamp:        runtime.Timestamp(time.Now().Unix()),
		ExceptionDetails: &runtime.ExceptionDetails{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Runtime.exceptionThrown",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Timestamp != result.Timestamp {
		t.Errorf("Expected %d, got %d", mockResult.Timestamp, result.Timestamp)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Runtime.exceptionThrown",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExecutionContextCreated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *runtime.ExecutionContextCreatedEvent)
	soc.Runtime().OnExecutionContextCreated(func(eventData *runtime.ExecutionContextCreatedEvent) {
		resultChan <- eventData
	})

	mockResult := &runtime.ExecutionContextCreatedEvent{
		Context: &runtime.ExecutionContextDescription{
			ID:      runtime.ExecutionContextID(1),
			Origin:  "origin",
			Name:    "name",
			AuxData: map[string]string{"key": "value"},
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Runtime.executionContextCreated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Context.ID != result.Context.ID {
		t.Errorf("Expected %d, got %d", mockResult.Context.ID, result.Context.ID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Runtime.executionContextCreated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExecutionContextDestroyed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *runtime.ExecutionContextDestroyedEvent)
	soc.Runtime().OnExecutionContextDestroyed(func(eventData *runtime.ExecutionContextDestroyedEvent) {
		resultChan <- eventData
	})

	mockResult := &runtime.ExecutionContextDestroyedEvent{
		ExecutionContextID: runtime.ExecutionContextID(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Runtime.executionContextDestroyed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ExecutionContextID != result.ExecutionContextID {
		t.Errorf("Expected %d, got %d", mockResult.ExecutionContextID, result.ExecutionContextID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Runtime.executionContextDestroyed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExecutionContextsCleared(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *runtime.ExecutionContextsClearedEvent)
	soc.Runtime().OnExecutionContextsCleared(func(eventData *runtime.ExecutionContextsClearedEvent) {
		resultChan <- eventData
	})

	mockResult := &runtime.ExecutionContextsClearedEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Runtime.executionContextsCleared",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Runtime.executionContextsCleared",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnInspectRequested(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *runtime.InspectRequestedEvent)
	soc.Runtime().OnInspectRequested(func(eventData *runtime.InspectRequestedEvent) {
		resultChan <- eventData
	})

	mockResult := &runtime.InspectRequestedEvent{
		Object: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			Subtype:             runtime.ObjectSubtype.Array,
			ClassName:           "class-name",
			Value:               "value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "description",
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
			Preview: &runtime.ObjectPreview{
				Type:        runtime.ObjectType.Object,
				Subtype:     runtime.ObjectSubtype.Array,
				Description: "description",
				Overflow:    true,
				Properties: []*runtime.PropertyPreview{{
					Name:  "name",
					Type:  runtime.ObjectType.Object,
					Value: "value",
					ValuePreview: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Subtype: runtime.ObjectSubtype.Array,
				}},
				Entries: []*runtime.EntryPreview{{
					Key: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
					Value: &runtime.ObjectPreview{
						Type: runtime.ObjectType.Accessor,
					},
				}},
			},
			CustomPreview: &runtime.CustomPreview{
				Header:                     "header",
				HasBody:                    true,
				FormatterObjectID:          runtime.RemoteObjectID("remote-object-id"),
				BindRemoteObjectFunctionID: runtime.RemoteObjectID("remote-object-id"),
				ConfigObjectID:             runtime.RemoteObjectID("remote-object-id"),
			},
		},
		Hints: map[string]string{"key": "value"},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Runtime.inspectRequested",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Object.Type != result.Object.Type {
		t.Errorf("Expected %s, got %s", mockResult.Object.Type, result.Object.Type)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Runtime.inspectRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
