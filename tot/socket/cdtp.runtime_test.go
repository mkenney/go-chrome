package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestRuntimeAwaitPromise(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeAwaitPromise")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.AwaitPromiseParams{
		PromiseObjectID: runtime.RemoteObjectID("remote-object-id"),
		ReturnByValue:   true,
		GeneratePreview: true,
	}
	resultChan := mockSocket.Runtime().AwaitPromise(params)
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
		t.Errorf("Expected %s, got %s", mockResult.Result.Type, result.Result.Type)
	}

	resultChan = mockSocket.Runtime().AwaitPromise(params)
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

func TestRuntimeCallFunctionOn(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeCallFunctionOn")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.Runtime().CallFunctionOn(params)
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
		t.Errorf("Expected %s, got %s", mockResult.Result.Type, result.Result.Type)
	}

	resultChan = mockSocket.Runtime().CallFunctionOn(params)
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

func TestRuntimeCompileScript(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeCompileScript")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.CompileScriptParams{
		Expression:         "expression",
		SourceURL:          "http://some.url",
		PersistScript:      true,
		ExecutionContextID: runtime.ExecutionContextID(1),
	}
	resultChan := mockSocket.Runtime().CompileScript(params)
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
	if mockResult.ScriptID != result.ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.ScriptID, result.ScriptID)
	}

	resultChan = mockSocket.Runtime().CompileScript(params)
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

func TestRuntimeDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeDisable")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Runtime().Disable()
	mockResult := &runtime.DisableResult{}
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

	resultChan = mockSocket.Runtime().Disable()
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

func TestRuntimeDiscardConsoleEntries(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeDiscardConsoleEntries")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Runtime().DiscardConsoleEntries()
	mockResult := &runtime.DiscardConsoleEntriesResult{}
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

	resultChan = mockSocket.Runtime().DiscardConsoleEntries()
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

func TestRuntimeEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeEnable")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Runtime().Enable()
	mockResult := &runtime.EnableResult{}
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

	resultChan = mockSocket.Runtime().Enable()
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

func TestRuntimeEvaluate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeEvaluate")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.Runtime().Evaluate(params)
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
		t.Errorf("Expected %s, got %s", mockResult.Result.Type, result.Result.Type)
	}

	resultChan = mockSocket.Runtime().Evaluate(params)
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

func TestRuntimeGetProperties(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeGetProperties")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.GetPropertiesParams{
		ObjectID:               runtime.RemoteObjectID("remote-object-id"),
		OwnProperties:          true,
		AccessorPropertiesOnly: true,
		GeneratePreview:        true,
	}
	resultChan := mockSocket.Runtime().GetProperties(params)
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
	if mockResult.Result[0].Name != result.Result[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].Name, result.Result[0].Name)
	}

	resultChan = mockSocket.Runtime().GetProperties(params)
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

func TestRuntimeGlobalLexicalScopeNames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeGlobalLexicalScopeNames")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.GlobalLexicalScopeNamesParams{
		ExecutionContextID: runtime.ExecutionContextID(1),
	}
	resultChan := mockSocket.Runtime().GlobalLexicalScopeNames(params)
	mockResult := &runtime.GlobalLexicalScopeNamesResult{
		Names: []string{"name1", "name2"},
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
	if mockResult.Names[0] != result.Names[0] {
		t.Errorf("Expected %s, got %s", mockResult.Names[0], result.Names[0])
	}

	resultChan = mockSocket.Runtime().GlobalLexicalScopeNames(params)
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

func TestRuntimeQueryObjects(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeQueryObjects")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.QueryObjectsParams{
		PrototypeObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.Runtime().QueryObjects(params)
	mockResult := &runtime.QueryObjectsResult{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
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
	if mockResult.ObjectID != result.ObjectID {
		t.Errorf("Expected %s, got %s", mockResult.ObjectID, result.ObjectID)
	}

	resultChan = mockSocket.Runtime().QueryObjects(params)
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

func TestRuntimeReleaseObject(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeReleaseObject")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.ReleaseObjectParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.Runtime().ReleaseObject(params)
	mockResult := &runtime.ReleaseObjectResult{}
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

	resultChan = mockSocket.Runtime().ReleaseObject(params)
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

func TestRuntimeReleaseObjectGroup(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeReleaseObjectGroup")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.ReleaseObjectGroupParams{
		ObjectGroup: "object-group",
	}
	resultChan := mockSocket.Runtime().ReleaseObjectGroup(params)
	mockResult := &runtime.ReleaseObjectGroupResult{}
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

	resultChan = mockSocket.Runtime().ReleaseObjectGroup(params)
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

func TestRuntimeRunIfWaitingForDebugger(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeRunIfWaitingForDebugger")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Runtime().RunIfWaitingForDebugger()
	mockResult := &runtime.RunIfWaitingForDebuggerResult{}
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

	resultChan = mockSocket.Runtime().RunIfWaitingForDebugger()
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

func TestRuntimeRunScript(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeRunScript")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.Runtime().RunScript(params)
	mockResult := &runtime.RunScriptResult{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
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
	if mockResult.ObjectID != result.ObjectID {
		t.Errorf("Expected %s, got %s", mockResult.ObjectID, result.ObjectID)
	}

	resultChan = mockSocket.Runtime().RunScript(params)
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

func TestRuntimeSetCustomObjectFormatterEnabled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeSetCustomObjectFormatterEnabled")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &runtime.SetCustomObjectFormatterEnabledParams{
		Result: &runtime.RemoteObject{
			Type: runtime.ObjectType.Accessor,
		},
		ExceptionDetails: &runtime.ExceptionDetails{},
	}
	resultChan := mockSocket.Runtime().SetCustomObjectFormatterEnabled(params)
	mockResult := &runtime.SetCustomObjectFormatterEnabledResult{}
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

	resultChan = mockSocket.Runtime().SetCustomObjectFormatterEnabled(params)
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

func TestRuntimeOnConsoleAPICalled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeOnConsoleAPICalled")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *runtime.ConsoleAPICalledEvent)
	mockSocket.Runtime().OnConsoleAPICalled(func(eventData *runtime.ConsoleAPICalledEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Runtime.consoleAPICalled",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Type != result.Type {
		t.Errorf("Expected %s, got %s", mockResult.Type, result.Type)
	}

	resultChan = make(chan *runtime.ConsoleAPICalledEvent)
	mockSocket.Runtime().OnConsoleAPICalled(func(eventData *runtime.ConsoleAPICalledEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Runtime.consoleAPICalled",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExceptionRevoked(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeOnExceptionRevoked")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *runtime.ExceptionRevokedEvent)
	mockSocket.Runtime().OnExceptionRevoked(func(eventData *runtime.ExceptionRevokedEvent) {
		resultChan <- eventData
	})
	mockResult := &runtime.ExceptionRevokedEvent{
		Reason:      "reason",
		ExceptionID: 1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Runtime.exceptionRevoked",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Reason != result.Reason {
		t.Errorf("Expected %s, got %s", mockResult.Reason, result.Reason)
	}

	resultChan = make(chan *runtime.ExceptionRevokedEvent)
	mockSocket.Runtime().OnExceptionRevoked(func(eventData *runtime.ExceptionRevokedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Runtime.exceptionRevoked",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExceptionThrown(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeOnExceptionThrown")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *runtime.ExceptionThrownEvent)
	mockSocket.Runtime().OnExceptionThrown(func(eventData *runtime.ExceptionThrownEvent) {
		resultChan <- eventData
	})
	mockResult := &runtime.ExceptionThrownEvent{
		Timestamp:        runtime.Timestamp(time.Now().Unix()),
		ExceptionDetails: &runtime.ExceptionDetails{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Runtime.exceptionThrown",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Timestamp != result.Timestamp {
		t.Errorf("Expected %d, got %d", mockResult.Timestamp, result.Timestamp)
	}

	resultChan = make(chan *runtime.ExceptionThrownEvent)
	mockSocket.Runtime().OnExceptionThrown(func(eventData *runtime.ExceptionThrownEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Runtime.exceptionThrown",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExecutionContextCreated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeOnExecutionContextCreated")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *runtime.ExecutionContextCreatedEvent)
	mockSocket.Runtime().OnExecutionContextCreated(func(eventData *runtime.ExecutionContextCreatedEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Runtime.executionContextCreated",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Context.ID != result.Context.ID {
		t.Errorf("Expected %d, got %d", mockResult.Context.ID, result.Context.ID)
	}

	resultChan = make(chan *runtime.ExecutionContextCreatedEvent)
	mockSocket.Runtime().OnExecutionContextCreated(func(eventData *runtime.ExecutionContextCreatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Runtime.executionContextCreated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExecutionContextDestroyed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeOnExecutionContextDestroyed")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *runtime.ExecutionContextDestroyedEvent)
	mockSocket.Runtime().OnExecutionContextDestroyed(func(eventData *runtime.ExecutionContextDestroyedEvent) {
		resultChan <- eventData
	})
	mockResult := &runtime.ExecutionContextDestroyedEvent{
		ExecutionContextID: runtime.ExecutionContextID(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Runtime.executionContextDestroyed",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ExecutionContextID != result.ExecutionContextID {
		t.Errorf("Expected %d, got %d", mockResult.ExecutionContextID, result.ExecutionContextID)
	}

	resultChan = make(chan *runtime.ExecutionContextDestroyedEvent)
	mockSocket.Runtime().OnExecutionContextDestroyed(func(eventData *runtime.ExecutionContextDestroyedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Runtime.executionContextDestroyed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnExecutionContextsCleared(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeOnExecutionContextsCleared")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *runtime.ExecutionContextsClearedEvent)
	mockSocket.Runtime().OnExecutionContextsCleared(func(eventData *runtime.ExecutionContextsClearedEvent) {
		resultChan <- eventData
	})
	mockResult := &runtime.ExecutionContextsClearedEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Runtime.executionContextsCleared",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *runtime.ExecutionContextsClearedEvent)
	mockSocket.Runtime().OnExecutionContextsCleared(func(eventData *runtime.ExecutionContextsClearedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Runtime.executionContextsCleared",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestRuntimeOnInspectRequested(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestRuntimeOnInspectRequested")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *runtime.InspectRequestedEvent)
	mockSocket.Runtime().OnInspectRequested(func(eventData *runtime.InspectRequestedEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Runtime.inspectRequested",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Object.Type != result.Object.Type {
		t.Errorf("Expected %s, got %s", mockResult.Object.Type, result.Object.Type)
	}

	resultChan = make(chan *runtime.InspectRequestedEvent)
	mockSocket.Runtime().OnInspectRequested(func(eventData *runtime.InspectRequestedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Runtime.inspectRequested",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
