package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/debugger"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestPageAddScriptToEvaluateOnLoad(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.AddScriptToEvaluateOnLoadParams{
		ScriptSource: "some-source",
	}
	mockResult := &page.AddScriptToEvaluateOnLoadResult{
		Identifier: page.ScriptIdentifier("script-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().AddScriptToEvaluateOnLoad(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Identifier != result.Identifier {
		t.Errorf("Expected %s, got %s", mockResult.Identifier, result.Identifier)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().AddScriptToEvaluateOnLoad(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageAddScriptToEvaluateOnNewDocument(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.AddScriptToEvaluateOnNewDocumentParams{
		Source: "some-source",
	}
	mockResult := &page.AddScriptToEvaluateOnNewDocumentResult{
		Identifier: page.ScriptIdentifier("script-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().AddScriptToEvaluateOnNewDocument(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Identifier != result.Identifier {
		t.Errorf("Expected %s, got %s", mockResult.Identifier, result.Identifier)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().AddScriptToEvaluateOnNewDocument(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageBringToFront(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &page.BringToFrontResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().BringToFront()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().BringToFront()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageCaptureScreenshot(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.CaptureScreenshotParams{
		Format:  page.Format.Jpeg,
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
	mockResult := &page.CaptureScreenshotResult{
		Data: "screenshot data",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().CaptureScreenshot(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Data != result.Data {
		t.Errorf("Expected %s, got %s", mockResult.Data, result.Data)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().CaptureScreenshot(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageCreateIsolatedWorld(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.CreateIsolatedWorldParams{
		FrameID:             page.FrameID("frame-id"),
		WorldName:           "world-name",
		GrantUniveralAccess: true,
	}
	mockResult := &page.CreateIsolatedWorldResult{
		ExecutionContextID: runtime.ExecutionContextID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().CreateIsolatedWorld(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ExecutionContextID != result.ExecutionContextID {
		t.Errorf("Expected %d, got %d", mockResult.ExecutionContextID, result.ExecutionContextID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().CreateIsolatedWorld(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &page.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &page.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageGetAppManifest(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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
	mockResult := &page.GetAppManifestResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().GetAppManifest(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().GetAppManifest(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageGetFrameTree(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().GetFrameTree()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.FrameTree.Frame.ID != result.FrameTree.Frame.ID {
		t.Errorf("Expected %s, got %s", mockResult.FrameTree.Frame.ID, result.FrameTree.Frame.ID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().GetFrameTree()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageGetLayoutMetrics(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().GetLayoutMetrics()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.LayoutViewport.PageX != result.LayoutViewport.PageX {
		t.Errorf("Expected %d, got %d", mockResult.LayoutViewport.PageX, result.LayoutViewport.PageX)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().GetLayoutMetrics()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageGetNavigationHistory(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().GetNavigationHistory()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.CurrentIndex != result.CurrentIndex {
		t.Errorf("Expected %d, got %d", mockResult.CurrentIndex, result.CurrentIndex)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().GetNavigationHistory()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageGetResourceContent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.GetResourceContentParams{
		FrameID: page.FrameID("frame-id"),
	}
	mockResult := &page.GetResourceContentResult{
		Content:       "content",
		Base64Encoded: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().GetResourceContent(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Content != result.Content {
		t.Errorf("Expected %s, got %s", mockResult.Content, result.Content)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().GetResourceContent(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageGetResourceTree(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().GetResourceTree()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.FrameTree.Frame.ID != result.FrameTree.Frame.ID {
		t.Errorf("Expected %s, got %s", mockResult.FrameTree.Frame.ID, result.FrameTree.Frame.ID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().GetResourceTree()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageHandleJavaScriptDialog(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.HandleJavaScriptDialogParams{
		Accept:     true,
		PromptText: "prompt text",
	}
	mockResult := &page.HandleJavaScriptDialogResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().HandleJavaScriptDialog(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().HandleJavaScriptDialog(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageNavigate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.NavigateParams{
		URL:            "http://some.url",
		Referrer:       "http://referrer.url",
		TransitionType: page.TransitionType("transition-type"),
	}
	mockResult := &page.NavigateResult{
		FrameID:   page.FrameID("frame-id"),
		LoaderID:  page.LoaderID("loader-id"),
		ErrorText: "error text",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().Navigate(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().Navigate(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageNavigateToHistoryEntry(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.NavigateToHistoryEntryParams{
		EntryID: 1,
	}
	mockResult := &page.NavigateToHistoryEntryResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().NavigateToHistoryEntry(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().NavigateToHistoryEntry(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPagePrintToPDF(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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
	mockResult := &page.PrintToPDFResult{
		Data: "result data",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().PrintToPDF(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Data != result.Data {
		t.Errorf("Expected %s, got %s", mockResult.Data, result.Data)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().PrintToPDF(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageReload(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.ReloadParams{
		IgnoreCache:            true,
		ScriptToEvaluateOnLoad: "some-script",
	}
	mockResult := &page.ReloadResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().Reload(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().Reload(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageRemoveScriptToEvaluateOnLoad(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.RemoveScriptToEvaluateOnLoadParams{
		Identifier: page.ScriptIdentifier("script-id"),
	}
	mockResult := &page.RemoveScriptToEvaluateOnLoadResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().RemoveScriptToEvaluateOnLoad(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().RemoveScriptToEvaluateOnLoad(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageRemoveScriptToEvaluateOnNewDocument(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.RemoveScriptToEvaluateOnNewDocumentParams{
		Identifier: page.ScriptIdentifier("script-id"),
	}
	mockResult := &page.RemoveScriptToEvaluateOnNewDocumentResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().RemoveScriptToEvaluateOnNewDocument(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().RemoveScriptToEvaluateOnNewDocument(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageRequestAppBanner(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &page.RequestAppBannerResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().RequestAppBanner()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().RequestAppBanner()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageScreencastFrameAck(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.ScreencastFrameAckParams{
		SessionID: 1,
	}
	mockResult := &page.ScreencastFrameAckResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().ScreencastFrameAck(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().ScreencastFrameAck(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageSearchInResource(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.SearchInResourceParams{
		FrameID:       page.FrameID("frame-id"),
		URL:           "http://some.url",
		Query:         "some query",
		CaseSensitive: true,
		IsRegex:       true,
	}
	mockResult := &page.SearchInResourceResult{
		Result: []*debugger.SearchMatch{{
			LineNumber:  1,
			LineContent: "some content",
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().SearchInResource(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].LineNumber != result.Result[0].LineNumber {
		t.Errorf("Expected %d, got %d", mockResult.Result[0].LineNumber, result.Result[0].LineNumber)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().SearchInResource(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageSetAdBlockingEnabled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.SetAdBlockingEnabledParams{
		Enabled: true,
	}
	mockResult := &page.SetAdBlockingEnabledResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().SetAdBlockingEnabled(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().SetAdBlockingEnabled(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageSetAutoAttachToCreatedPages(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.SetAutoAttachToCreatedPagesParams{
		AutoAttach: true,
	}
	mockResult := &page.SetAutoAttachToCreatedPagesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().SetAutoAttachToCreatedPages(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().SetAutoAttachToCreatedPages(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageSetDocumentContent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.SetDocumentContentParams{
		FrameID: page.FrameID("frame-id"),
		HTML:    "<some>html</some>",
	}
	mockResult := &page.SetDocumentContentResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().SetDocumentContent(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().SetDocumentContent(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageSetDownloadBehavior(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.SetDownloadBehaviorParams{
		Behavior:     page.Behavior.Allow,
		DownloadPath: "/some/path",
	}
	mockResult := &page.SetDownloadBehaviorResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().SetDownloadBehavior(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().SetDownloadBehavior(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageSetLifecycleEventsEnabled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.SetLifecycleEventsEnabledParams{
		Enabled: true,
	}
	mockResult := &page.SetLifecycleEventsEnabledResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().SetLifecycleEventsEnabled(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().SetLifecycleEventsEnabled(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageStartScreencast(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &page.StartScreencastParams{
		Format:        page.Format.Jpeg,
		Quality:       1,
		MaxWidth:      1,
		MaxHeight:     1,
		EveryNthFrame: 1,
	}
	mockResult := &page.StartScreencastResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().StartScreencast(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().StartScreencast(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageStopLoading(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &page.StopLoadingResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().StopLoading()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().StopLoading()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageStopScreencast(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &page.StopScreencastResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Page().StopScreencast()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Page().StopScreencast()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnDOMContentEventFired(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.DOMContentEventFiredEvent)
	soc.Page().OnDOMContentEventFired(func(eventData *page.DOMContentEventFiredEvent) {
		resultChan <- eventData
	})

	mockResult := &page.DOMContentEventFiredEvent{
		Timestamp: page.MonotonicTime(time.Now().Unix()),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.domContentEventFired",
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
		Method: "Page.domContentEventFired",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameAttached(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.FrameAttachedEvent)
	soc.Page().OnFrameAttached(func(eventData *page.FrameAttachedEvent) {
		resultChan <- eventData
	})

	mockResult := &page.FrameAttachedEvent{
		FrameID:       page.FrameID("frame-id"),
		ParentFrameID: page.FrameID("parent-id"),
		Stack:         &runtime.StackTrace{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameAttached",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameAttached",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameClearedScheduledNavigation(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.FrameClearedScheduledNavigationEvent)
	soc.Page().OnFrameClearedScheduledNavigation(func(eventData *page.FrameClearedScheduledNavigationEvent) {
		resultChan <- eventData
	})

	mockResult := &page.FrameClearedScheduledNavigationEvent{
		FrameID: page.FrameID("frame-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameClearedScheduledNavigation",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameClearedScheduledNavigation",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameDetached(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &page.FrameDetachedEvent{
		FrameID: page.FrameID("frame-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameDetached",
	})
	resultChan := make(chan *page.FrameDetachedEvent)
	soc.Page().OnFrameDetached(func(eventData *page.FrameDetachedEvent) {
		resultChan <- eventData
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameDetached",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameNavigated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.FrameNavigatedEvent)
	soc.Page().OnFrameNavigated(func(eventData *page.FrameNavigatedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameNavigated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Frame.ID != result.Frame.ID {
		t.Errorf("Expected %s, got %s", mockResult.Frame.ID, result.Frame.ID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameNavigated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameResized(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.FrameResizedEvent)
	soc.Page().OnFrameResized(func(eventData *page.FrameResizedEvent) {
		resultChan <- eventData
	})

	mockResult := &page.FrameResizedEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameResized",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameResized",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameScheduledNavigation(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.FrameScheduledNavigationEvent)
	soc.Page().OnFrameScheduledNavigation(func(eventData *page.FrameScheduledNavigationEvent) {
		resultChan <- eventData
	})

	mockResult := &page.FrameScheduledNavigationEvent{
		FrameID: page.FrameID("frame-id"),
		Delay:   1,
		Reason:  page.Reason.FormSubmissionGet,
		URL:     "http://some.url",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameScheduledNavigation",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameScheduledNavigation",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameStartedLoading(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.FrameStartedLoadingEvent)
	soc.Page().OnFrameStartedLoading(func(eventData *page.FrameStartedLoadingEvent) {
		resultChan <- eventData
	})

	mockResult := &page.FrameStartedLoadingEvent{
		FrameID: page.FrameID("frame-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameStartedLoading",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameStartedLoading",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnFrameStoppedLoading(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.FrameStoppedLoadingEvent)
	soc.Page().OnFrameStoppedLoading(func(eventData *page.FrameStoppedLoadingEvent) {
		resultChan <- eventData
	})

	mockResult := &page.FrameStoppedLoadingEvent{
		FrameID: page.FrameID("frame-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.frameStoppedLoading",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.frameStoppedLoading",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnInterstitialHidden(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.InterstitialHiddenEvent)
	soc.Page().OnInterstitialHidden(func(eventData *page.InterstitialHiddenEvent) {
		resultChan <- eventData
	})

	mockResult := &page.InterstitialHiddenEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.interstitialHidden",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.interstitialHidden",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnInterstitialShown(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.InterstitialShownEvent)
	soc.Page().OnInterstitialShown(func(eventData *page.InterstitialShownEvent) {
		resultChan <- eventData
	})

	mockResult := &page.InterstitialShownEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.interstitialShown",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.interstitialShown",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnJavascriptDialogClosed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.JavascriptDialogClosedEvent)
	soc.Page().OnJavascriptDialogClosed(func(eventData *page.JavascriptDialogClosedEvent) {
		resultChan <- eventData
	})

	mockResult := &page.JavascriptDialogClosedEvent{
		Result:    true,
		UserInput: "user input",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.javascriptDialogClosed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.javascriptDialogClosed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnJavascriptDialogOpening(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.JavascriptDialogOpeningEvent)
	soc.Page().OnJavascriptDialogOpening(func(eventData *page.JavascriptDialogOpeningEvent) {
		resultChan <- eventData
	})

	mockResult := &page.JavascriptDialogOpeningEvent{
		URL:           "http://some.url",
		Message:       "some message",
		Type:          page.DialogType.Alert,
		DefaultPrompt: "some prompt",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.javascriptDialogOpening",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.URL != result.URL {
		t.Errorf("Expected %s, got %s", mockResult.URL, result.URL)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.javascriptDialogOpening",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnLifecycleEvent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.LifecycleEventEvent)
	soc.Page().OnLifecycleEvent(func(eventData *page.LifecycleEventEvent) {
		resultChan <- eventData
	})

	mockResult := &page.LifecycleEventEvent{
		FrameID:   page.FrameID("frame-id"),
		LoaderID:  page.LoaderID("loader-id"),
		Name:      "name",
		Timestamp: page.MonotonicTime(time.Now().Unix()),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.lifecycleEvent",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.FrameID != result.FrameID {
		t.Errorf("Expected %s, got %s", mockResult.FrameID, result.FrameID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.lifecycleEvent",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnLoadEventFired(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.LoadEventFiredEvent)
	soc.Page().OnLoadEventFired(func(eventData *page.LoadEventFiredEvent) {
		resultChan <- eventData
	})

	mockResult := &page.LoadEventFiredEvent{
		Timestamp: page.MonotonicTime(time.Now().Unix()),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.loadEventFired",
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
		Method: "Page.loadEventFired",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnScreencastFrame(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.ScreencastFrameEvent)
	soc.Page().OnScreencastFrame(func(eventData *page.ScreencastFrameEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.screencastFrame",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Data != result.Data {
		t.Errorf("Expected %s, got %s", mockResult.Data, result.Data)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.screencastFrame",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnScreencastVisibilityChanged(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.ScreencastVisibilityChangedEvent)
	soc.Page().OnScreencastVisibilityChanged(func(eventData *page.ScreencastVisibilityChangedEvent) {
		resultChan <- eventData
	})

	mockResult := &page.ScreencastVisibilityChangedEvent{
		Visible: true,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.screencastVisibilityChanged",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Visible != result.Visible {
		t.Errorf("Expected %v, got %v", mockResult.Visible, result.Visible)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.screencastVisibilityChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPageOnWindowOpen(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *page.WindowOpenEvent)
	soc.Page().OnWindowOpen(func(eventData *page.WindowOpenEvent) {
		resultChan <- eventData
	})

	mockResult := &page.WindowOpenEvent{
		URL:            "http://some.url",
		WindowName:     "window-name",
		WindowFeatures: []string{"feature1", "feature2"},
		UserGesture:    true,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Page.windowOpen",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.URL != result.URL {
		t.Errorf("Expected %s, got %s", mockResult.URL, result.URL)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Page.windowOpen",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
