package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	debugger "github.com/mkenney/go-chrome/cdtp/debugger"
	page "github.com/mkenney/go-chrome/cdtp/page"
	runtime "github.com/mkenney/go-chrome/cdtp/runtime"
)

func TestPageAddScriptToEvaluateOnLoad(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.AddScriptToEvaluateOnLoadParams{
		ScriptSource: "some-source",
	}
	resultChan := mockSocket.Page().AddScriptToEvaluateOnLoad(params)
	mockResult := &page.AddScriptToEvaluateOnLoadResult{
		Identifier: page.ScriptIdentifier("script-id"),
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
	if mockResult.Identifier != result.Identifier {
		t.Errorf("Expected %s, got %s", mockResult.Identifier, result.Identifier)
	}

	resultChan = mockSocket.Page().AddScriptToEvaluateOnLoad(params)
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

func TestPageAddScriptToEvaluateOnNewDocument(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.AddScriptToEvaluateOnNewDocumentParams{
		Source: "some-source",
	}
	resultChan := mockSocket.Page().AddScriptToEvaluateOnNewDocument(params)
	mockResult := &page.AddScriptToEvaluateOnNewDocumentResult{
		Identifier: page.ScriptIdentifier("script-id"),
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
	if mockResult.Identifier != result.Identifier {
		t.Errorf("Expected %s, got %s", mockResult.Identifier, result.Identifier)
	}

	resultChan = mockSocket.Page().AddScriptToEvaluateOnNewDocument(params)
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

func TestPageBringToFront(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().BringToFront()
	mockResult := &page.BringToFrontResult{}
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

	resultChan = mockSocket.Page().BringToFront()
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

func TestPageCaptureScreenshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.CaptureScreenshotParams{
		Format:  "jpeg",
		Quality: 50,
		Clip: &page.Viewport{
			X:      1,
			Y:      1,
			Width:  1,
			Height: 1,
			Scale:  1,
		},
		FromSurface: true,
	}
	resultChan := mockSocket.Page().CaptureScreenshot(params)
	mockResult := &page.CaptureScreenshotResult{
		Data: "screenshot data",
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
	if mockResult.Data != result.Data {
		t.Errorf("Expected %s, got %s", mockResult.Data, result.Data)
	}

	resultChan = mockSocket.Page().CaptureScreenshot(params)
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

func TestPageCreateIsolatedWorld(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.CreateIsolatedWorldParams{
		FrameID:             page.FrameID("frame-id"),
		WorldName:           "world-name",
		GrantUniveralAccess: true,
	}
	resultChan := mockSocket.Page().CreateIsolatedWorld(params)
	mockResult := &page.CreateIsolatedWorldResult{
		ExecutionContextID: runtime.ExecutionContextID(1),
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
	if mockResult.ExecutionContextID != result.ExecutionContextID {
		t.Errorf("Expected %d, got %d", mockResult.ExecutionContextID, result.ExecutionContextID)
	}

	resultChan = mockSocket.Page().CreateIsolatedWorld(params)
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

func TestPageDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().Disable()
	mockResult := &page.DisableResult{}
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

	resultChan = mockSocket.Page().Disable()
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

func TestPageEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().Enable()
	mockResult := &page.EnableResult{}
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

	resultChan = mockSocket.Page().Enable()
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

func TestPageGetAppManifest(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.GetAppManifestParams{
		URL: "http://some.url",
		Errors: []*page.AppManifestError{{
			Message:  "message",
			Critical: 1,
			Line:     1,
			Column:   1,
		}},
		Data: "some data",
	}
	resultChan := mockSocket.Page().GetAppManifest(params)
	mockResult := &page.GetAppManifestResult{}
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

	resultChan = mockSocket.Page().GetAppManifest(params)
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

func TestPageGetFrameTree(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().GetFrameTree()
	mockResult := &page.GetFrameTreeResult{
		FrameTree: &page.FrameTree{
			Frame: &page.Frame{
				ID:             "frame-id",
				ParentID:       "parent-id",
				LoaderID:       page.LoaderID("loader-id"),
				Name:           "name",
				URL:            "http://some.url",
				SecurityOrigin: "origin",
				MimeType:       "mime/type",
				UnreachableURL: "http://someother.url",
			},
			ChildFrames: []*page.FrameTree{},
		},
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
	if mockResult.FrameTree.Frame.ID != result.FrameTree.Frame.ID {
		t.Errorf("Expected %s, got %s", mockResult.FrameTree.Frame.ID, result.FrameTree.Frame.ID)
	}

	resultChan = mockSocket.Page().GetFrameTree()
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

func TestPageGetLayoutMetrics(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().GetLayoutMetrics()
	mockResult := &page.GetLayoutMetricsResult{
		LayoutViewport: &page.LayoutViewport{
			PageX:        1,
			PageY:        1,
			ClientWidth:  1,
			ClientHeight: 1,
		},
		VisualViewport: &page.VisualViewport{
			OffsetX:      1,
			OffsetY:      1,
			PageX:        1,
			PageY:        1,
			ClientWidth:  1,
			ClientHeight: 1,
			Scale:        1,
		},
		ContentSize: &page.Rect{
			X:      1,
			Y:      1,
			Width:  1,
			Height: 1,
		},
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
	if mockResult.LayoutViewport.PageX != result.LayoutViewport.PageX {
		t.Errorf("Expected %d, got %d", mockResult.LayoutViewport.PageX, result.LayoutViewport.PageX)
	}

	resultChan = mockSocket.Page().GetLayoutMetrics()
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

func TestPageGetNavigationHistory(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().GetNavigationHistory()
	mockResult := &page.GetNavigationHistoryResult{
		CurrentIndex: 1,
		Entries: []*page.NavigationEntry{{
			ID:             1,
			URL:            "http://some.url",
			UserTypedURL:   "http://someother.url",
			Title:          "title",
			TransitionType: page.TransitionType("type"),
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
	if mockResult.CurrentIndex != result.CurrentIndex {
		t.Errorf("Expected %d, got %d", mockResult.CurrentIndex, result.CurrentIndex)
	}

	resultChan = mockSocket.Page().GetNavigationHistory()
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

func TestPageGetResourceContent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.GetResourceContentParams{
		FrameID: page.FrameID("frame-id"),
	}
	resultChan := mockSocket.Page().GetResourceContent(params)
	mockResult := &page.GetResourceContentResult{
		Content:       "content",
		Base64Encoded: true,
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
	if mockResult.Content != result.Content {
		t.Errorf("Expected %s, got %s", mockResult.Content, result.Content)
	}

	resultChan = mockSocket.Page().GetResourceContent(params)
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

func TestPageGetResourceTree(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().GetResourceTree()
	mockResult := &page.GetResourceTreeResult{
		FrameTree: &page.FrameResourceTree{
			Frame: &page.Frame{
				ID:             "frame-id",
				ParentID:       "parent-id",
				LoaderID:       page.LoaderID("loader-id"),
				Name:           "name",
				URL:            "http://some.url",
				SecurityOrigin: "origin",
				MimeType:       "mime/type",
				UnreachableURL: "http://someother.url",
			},
			ChildFrames: []*page.FrameResourceTree{},
			Resources:   []*page.FrameResource{},
		},
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
	if mockResult.FrameTree.Frame.ID != result.FrameTree.Frame.ID {
		t.Errorf("Expected %s, got %s", mockResult.FrameTree.Frame.ID, result.FrameTree.Frame.ID)
	}

	resultChan = mockSocket.Page().GetResourceTree()
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

func TestPageHandleJavaScriptDialog(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.HandleJavaScriptDialogParams{
		Accept:     true,
		PromptText: "prompt text",
	}
	resultChan := mockSocket.Page().HandleJavaScriptDialog(params)
	mockResult := &page.HandleJavaScriptDialogResult{}
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

	resultChan = mockSocket.Page().HandleJavaScriptDialog(params)
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

func TestPageNavigate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.NavigateParams{
		URL:            "http://some.url",
		Referrer:       "http://referrer.url",
		TransitionType: page.TransitionType("transition-type"),
	}
	resultChan := mockSocket.Page().Navigate(params)
	mockResult := &page.NavigateResult{
		FrameID:   page.FrameID("frame-id"),
		LoaderID:  page.LoaderID("loader-id"),
		ErrorText: "error text",
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
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = mockSocket.Page().Navigate(params)
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

func TestPageNavigateToHistoryEntry(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.NavigateToHistoryEntryParams{
		EntryID: 1,
	}
	resultChan := mockSocket.Page().NavigateToHistoryEntry(params)
	mockResult := &page.NavigateToHistoryEntryResult{}
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

	resultChan = mockSocket.Page().NavigateToHistoryEntry(params)
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

func TestPagePrintToPDF(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.PrintToPDFParams{
		Landscape:               true,
		DisplayHeaderFooter:     true,
		PrintBackground:         true,
		Scale:                   1,
		PaperWidth:              1,
		PaperHeight:             1,
		MarginTop:               1,
		MarginBottom:            1,
		MarginLeft:              1,
		MarginRight:             1,
		PageRanges:              "1-2",
		IgnoreInvalidPageRanges: true,
	}
	resultChan := mockSocket.Page().PrintToPDF(params)
	mockResult := &page.PrintToPDFResult{
		Data: "result data",
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
	if mockResult.Data != result.Data {
		t.Errorf("Expected %s, got %s", mockResult.Data, result.Data)
	}

	resultChan = mockSocket.Page().PrintToPDF(params)
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

func TestPageReload(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.ReloadParams{
		IgnoreCache:            true,
		ScriptToEvaluateOnLoad: "some-script",
	}
	resultChan := mockSocket.Page().Reload(params)
	mockResult := &page.ReloadResult{}
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

	resultChan = mockSocket.Page().Reload(params)
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

func TestPageRemoveScriptToEvaluateOnLoad(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.RemoveScriptToEvaluateOnLoadParams{
		Identifier: page.ScriptIdentifier("script-id"),
	}
	resultChan := mockSocket.Page().RemoveScriptToEvaluateOnLoad(params)
	mockResult := &page.RemoveScriptToEvaluateOnLoadResult{}
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

	resultChan = mockSocket.Page().RemoveScriptToEvaluateOnLoad(params)
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

func TestPageRemoveScriptToEvaluateOnNewDocument(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.RemoveScriptToEvaluateOnNewDocumentParams{
		Identifier: page.ScriptIdentifier("script-id"),
	}
	resultChan := mockSocket.Page().RemoveScriptToEvaluateOnNewDocument(params)
	mockResult := &page.RemoveScriptToEvaluateOnNewDocumentResult{}
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

	resultChan = mockSocket.Page().RemoveScriptToEvaluateOnNewDocument(params)
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

func TestPageRequestAppBanner(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().RequestAppBanner()
	mockResult := &page.RequestAppBannerResult{}
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

	resultChan = mockSocket.Page().RequestAppBanner()
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

func TestPageScreencastFrameAck(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.ScreencastFrameAckParams{
		SessionID: 1,
	}
	resultChan := mockSocket.Page().ScreencastFrameAck(params)
	mockResult := &page.ScreencastFrameAckResult{}
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

	resultChan = mockSocket.Page().ScreencastFrameAck(params)
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

func TestPageSearchInResource(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.SearchInResourceParams{
		FrameID:       page.FrameID("frame-id"),
		URL:           "http://some.url",
		Query:         "some query",
		CaseSensitive: true,
		IsRegex:       true,
	}
	resultChan := mockSocket.Page().SearchInResource(params)
	mockResult := &page.SearchInResourceResult{
		Result: []*debugger.SearchMatch{{
			LineNumber:  1,
			LineContent: "some content",
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
	if mockResult.Result[0].LineNumber != result.Result[0].LineNumber {
		t.Errorf("Expected %d, got %d", mockResult.Result[0].LineNumber, result.Result[0].LineNumber)
	}

	resultChan = mockSocket.Page().SearchInResource(params)
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

func TestPageSetAdBlockingEnabled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.SetAdBlockingEnabledParams{
		Enabled: true,
	}
	resultChan := mockSocket.Page().SetAdBlockingEnabled(params)
	mockResult := &page.SetAdBlockingEnabledResult{}
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

	resultChan = mockSocket.Page().SetAdBlockingEnabled(params)
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

func TestPageSetAutoAttachToCreatedPages(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.SetAutoAttachToCreatedPagesParams{
		AutoAttach: true,
	}
	resultChan := mockSocket.Page().SetAutoAttachToCreatedPages(params)
	mockResult := &page.SetAutoAttachToCreatedPagesResult{}
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

	resultChan = mockSocket.Page().SetAutoAttachToCreatedPages(params)
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

func TestPageSetDocumentContent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.SetDocumentContentParams{
		FrameID: page.FrameID("frame-id"),
		HTML:    "<some>html</some>",
	}
	resultChan := mockSocket.Page().SetDocumentContent(params)
	mockResult := &page.SetDocumentContentResult{}
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

	resultChan = mockSocket.Page().SetDocumentContent(params)
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

func TestPageSetDownloadBehavior(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.SetDownloadBehaviorParams{
		Behavior:     "allow",
		DownloadPath: "/some/path",
	}
	resultChan := mockSocket.Page().SetDownloadBehavior(params)
	mockResult := &page.SetDownloadBehaviorResult{}
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

	resultChan = mockSocket.Page().SetDownloadBehavior(params)
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

func TestPageSetLifecycleEventsEnabled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.SetLifecycleEventsEnabledParams{
		Enabled: true,
	}
	resultChan := mockSocket.Page().SetLifecycleEventsEnabled(params)
	mockResult := &page.SetLifecycleEventsEnabledResult{}
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

	resultChan = mockSocket.Page().SetLifecycleEventsEnabled(params)
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

func TestPageStartScreencast(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &page.StartScreencastParams{
		Format:        "jpeg",
		Quality:       1,
		MaxWidth:      1,
		MaxHeight:     1,
		EveryNthFrame: 1,
	}
	resultChan := mockSocket.Page().StartScreencast(params)
	mockResult := &page.StartScreencastResult{}
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

	resultChan = mockSocket.Page().StartScreencast(params)
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

func TestPageStopLoading(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().StopLoading()
	mockResult := &page.StopLoadingResult{}
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

	resultChan = mockSocket.Page().StopLoading()
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

func TestPageStopScreencast(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Page().StopScreencast()
	mockResult := &page.StopScreencastResult{}
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

	resultChan = mockSocket.Page().StopScreencast()
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

func TestPageOnDOMContentEventFired(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.DOMContentEventFiredEvent)
	mockSocket.Page().OnDOMContentEventFired(func(eventData *page.DOMContentEventFiredEvent) {
		resultChan <- eventData
	})
	mockResult := &page.DOMContentEventFiredEvent{
		Timestamp: page.MonotonicTime(time.Now().Unix()),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.domContentEventFired",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Timestamp != result.Timestamp {
		t.Errorf("Expected %d, got %d", mockResult.Timestamp, result.Timestamp)
	}

	resultChan = make(chan *page.DOMContentEventFiredEvent)
	mockSocket.Page().OnDOMContentEventFired(func(eventData *page.DOMContentEventFiredEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.domContentEventFired",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameAttached(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameAttachedEvent)
	mockSocket.Page().OnFrameAttached(func(eventData *page.FrameAttachedEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameAttachedEvent{
		FrameID:       page.FrameID("frame-id"),
		ParentFrameID: page.FrameID("parent-id"),
		Stack:         &runtime.StackTrace{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameAttached",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = make(chan *page.FrameAttachedEvent)
	mockSocket.Page().OnFrameAttached(func(eventData *page.FrameAttachedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameAttached",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameClearedScheduledNavigation(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameClearedScheduledNavigationEvent)
	mockSocket.Page().OnFrameClearedScheduledNavigation(func(eventData *page.FrameClearedScheduledNavigationEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameClearedScheduledNavigationEvent{
		FrameID: page.FrameID("frame-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameClearedScheduledNavigation",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = make(chan *page.FrameClearedScheduledNavigationEvent)
	mockSocket.Page().OnFrameClearedScheduledNavigation(func(eventData *page.FrameClearedScheduledNavigationEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameClearedScheduledNavigation",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameDetached(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameDetachedEvent)
	mockSocket.Page().OnFrameDetached(func(eventData *page.FrameDetachedEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameDetachedEvent{
		FrameID: page.FrameID("frame-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameDetached",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = make(chan *page.FrameDetachedEvent)
	mockSocket.Page().OnFrameDetached(func(eventData *page.FrameDetachedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameDetached",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameNavigated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameNavigatedEvent)
	mockSocket.Page().OnFrameNavigated(func(eventData *page.FrameNavigatedEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameNavigatedEvent{
		Frame: &page.Frame{
			ID:             "frame-id",
			ParentID:       "parent-id",
			LoaderID:       page.LoaderID("loader-id"),
			Name:           "name",
			URL:            "http://some.url",
			SecurityOrigin: "origin",
			MimeType:       "mime/type",
			UnreachableURL: "http://someother.url",
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameNavigated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Frame.ID != result.Frame.ID {
		t.Errorf("Expected %s, got %s", mockResult.Frame.ID, result.Frame.ID)
	}

	resultChan = make(chan *page.FrameNavigatedEvent)
	mockSocket.Page().OnFrameNavigated(func(eventData *page.FrameNavigatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameNavigated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameResized(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameResizedEvent)
	mockSocket.Page().OnFrameResized(func(eventData *page.FrameResizedEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameResizedEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameResized",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *page.FrameResizedEvent)
	mockSocket.Page().OnFrameResized(func(eventData *page.FrameResizedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameResized",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameScheduledNavigation(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameScheduledNavigationEvent)
	mockSocket.Page().OnFrameScheduledNavigation(func(eventData *page.FrameScheduledNavigationEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameScheduledNavigationEvent{
		FrameID: page.FrameID("frame-id"),
		Delay:   1,
		Reason:  "formSubmissionGet",
		URL:     "http://some.url",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameScheduledNavigation",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = make(chan *page.FrameScheduledNavigationEvent)
	mockSocket.Page().OnFrameScheduledNavigation(func(eventData *page.FrameScheduledNavigationEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameScheduledNavigation",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameStartedLoading(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameStartedLoadingEvent)
	mockSocket.Page().OnFrameStartedLoading(func(eventData *page.FrameStartedLoadingEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameStartedLoadingEvent{
		FrameID: page.FrameID("frame-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameStartedLoading",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = make(chan *page.FrameStartedLoadingEvent)
	mockSocket.Page().OnFrameStartedLoading(func(eventData *page.FrameStartedLoadingEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameStartedLoading",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameStoppedLoading(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.FrameStoppedLoadingEvent)
	mockSocket.Page().OnFrameStoppedLoading(func(eventData *page.FrameStoppedLoadingEvent) {
		resultChan <- eventData
	})
	mockResult := &page.FrameStoppedLoadingEvent{
		FrameID: page.FrameID("frame-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.frameStoppedLoading",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = make(chan *page.FrameStoppedLoadingEvent)
	mockSocket.Page().OnFrameStoppedLoading(func(eventData *page.FrameStoppedLoadingEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.frameStoppedLoading",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnInterstitialHidden(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.InterstitialHiddenEvent)
	mockSocket.Page().OnInterstitialHidden(func(eventData *page.InterstitialHiddenEvent) {
		resultChan <- eventData
	})
	mockResult := &page.InterstitialHiddenEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.interstitialHidden",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *page.InterstitialHiddenEvent)
	mockSocket.Page().OnInterstitialHidden(func(eventData *page.InterstitialHiddenEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.interstitialHidden",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnInterstitialShown(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.InterstitialShownEvent)
	mockSocket.Page().OnInterstitialShown(func(eventData *page.InterstitialShownEvent) {
		resultChan <- eventData
	})
	mockResult := &page.InterstitialShownEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.interstitialShown",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *page.InterstitialShownEvent)
	mockSocket.Page().OnInterstitialShown(func(eventData *page.InterstitialShownEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.interstitialShown",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnJavascriptDialogClosed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.JavascriptDialogClosedEvent)
	mockSocket.Page().OnJavascriptDialogClosed(func(eventData *page.JavascriptDialogClosedEvent) {
		resultChan <- eventData
	})
	mockResult := &page.JavascriptDialogClosedEvent{
		Result:    true,
		UserInput: "user input",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.javascriptDialogClosed",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	resultChan = make(chan *page.JavascriptDialogClosedEvent)
	mockSocket.Page().OnJavascriptDialogClosed(func(eventData *page.JavascriptDialogClosedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.javascriptDialogClosed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnJavascriptDialogOpening(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.JavascriptDialogOpeningEvent)
	mockSocket.Page().OnJavascriptDialogOpening(func(eventData *page.JavascriptDialogOpeningEvent) {
		resultChan <- eventData
	})
	mockResult := &page.JavascriptDialogOpeningEvent{
		URL:           "http://some.url",
		Message:       "some message",
		Type:          page.DialogType("dialog type"),
		DefaultPrompt: "some prompt",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.javascriptDialogOpening",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.URL != result.URL {
		t.Errorf("Expected %s, got %s", mockResult.URL, result.URL)
	}

	resultChan = make(chan *page.JavascriptDialogOpeningEvent)
	mockSocket.Page().OnJavascriptDialogOpening(func(eventData *page.JavascriptDialogOpeningEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.javascriptDialogOpening",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnLifecycleEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.LifecycleEventEvent)
	mockSocket.Page().OnLifecycleEvent(func(eventData *page.LifecycleEventEvent) {
		resultChan <- eventData
	})
	mockResult := &page.LifecycleEventEvent{
		FrameID:   page.FrameID("frame-id"),
		LoaderID:  page.LoaderID("loader-id"),
		Name:      "name",
		Timestamp: page.MonotonicTime(time.Now().Unix()),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.lifecycleEvent",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	resultChan = make(chan *page.LifecycleEventEvent)
	mockSocket.Page().OnLifecycleEvent(func(eventData *page.LifecycleEventEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.lifecycleEvent",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnLoadEventFired(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.LoadEventFiredEvent)
	mockSocket.Page().OnLoadEventFired(func(eventData *page.LoadEventFiredEvent) {
		resultChan <- eventData
	})
	mockResult := &page.LoadEventFiredEvent{
		Timestamp: page.MonotonicTime(time.Now().Unix()),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.loadEventFired",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Timestamp != result.Timestamp {
		t.Errorf("Expected %d, got %d", mockResult.Timestamp, result.Timestamp)
	}

	resultChan = make(chan *page.LoadEventFiredEvent)
	mockSocket.Page().OnLoadEventFired(func(eventData *page.LoadEventFiredEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.loadEventFired",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnScreencastFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.ScreencastFrameEvent)
	mockSocket.Page().OnScreencastFrame(func(eventData *page.ScreencastFrameEvent) {
		resultChan <- eventData
	})
	mockResult := &page.ScreencastFrameEvent{
		Data: "data",
		Metadata: &page.ScreencastFrameMetadata{
			OffsetTop:       1,
			PageScaleFactor: 1,
			DeviceWidth:     1,
			DeviceHeight:    1,
			ScrollOffsetX:   1,
			ScrollOffsetY:   1,
			Timestamp:       page.TimeSinceEpoch(time.Now().Unix()),
		},
		SessionID: 1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.screencastFrame",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Data != result.Data {
		t.Errorf("Expected %d, got %d", mockResult.Data, result.Data)
	}

	resultChan = make(chan *page.ScreencastFrameEvent)
	mockSocket.Page().OnScreencastFrame(func(eventData *page.ScreencastFrameEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.screencastFrame",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnScreencastVisibilityChanged(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.ScreencastVisibilityChangedEvent)
	mockSocket.Page().OnScreencastVisibilityChanged(func(eventData *page.ScreencastVisibilityChangedEvent) {
		resultChan <- eventData
	})
	mockResult := &page.ScreencastVisibilityChangedEvent{
		Visible: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.screencastVisibilityChanged",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Visible != result.Visible {
		t.Errorf("Expected %v, got %v", mockResult.Visible, result.Visible)
	}

	resultChan = make(chan *page.ScreencastVisibilityChangedEvent)
	mockSocket.Page().OnScreencastVisibilityChanged(func(eventData *page.ScreencastVisibilityChangedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.screencastVisibilityChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnWindowOpen(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *page.WindowOpenEvent)
	mockSocket.Page().OnWindowOpen(func(eventData *page.WindowOpenEvent) {
		resultChan <- eventData
	})
	mockResult := &page.WindowOpenEvent{
		URL:            "http://some.url",
		WindowName:     "window-name",
		WindowFeatures: []string{"feature1", "feature2"},
		UserGesture:    true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Page.windowOpen",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.URL != result.URL {
		t.Errorf("Expected %s, got %s", mockResult.URL, result.URL)
	}

	resultChan = make(chan *page.WindowOpenEvent)
	mockSocket.Page().OnWindowOpen(func(eventData *page.WindowOpenEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Page.windowOpen",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
