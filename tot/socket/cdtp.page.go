package socket

import (
	"encoding/json"

	page "github.com/mkenney/go-chrome/tot/cdtp/page"
)

/*
PageProtocol provides a namespace for the Chrome Page protocol methods. The Page
protocol provides actions and events related to the inspected page belong to the
page domain.

https://chromedevtools.github.io/devtools-protocol/tot/Page/
*/
type PageProtocol struct {
	Socket Socketer
}

/*
AddScriptToEvaluateOnLoad is eprecated, please use addScriptToEvaluateOnNewDocument
instead.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
EXPERIMENTAL. DEPRECATED.
*/
func (protocol *PageProtocol) AddScriptToEvaluateOnLoad(
	params *page.AddScriptToEvaluateOnLoadParams,
) <-chan *page.AddScriptToEvaluateOnLoadResult {
	resultChan := make(chan *page.AddScriptToEvaluateOnLoadResult)
	command := NewCommand(protocol.Socket, "Page.addScriptToEvaluateOnLoad", params)
	result := &page.AddScriptToEvaluateOnLoadResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
AddScriptToEvaluateOnNewDocument evaluates given script in every frame upon
creation (before loading frame's scripts).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
func (protocol *PageProtocol) AddScriptToEvaluateOnNewDocument(
	params *page.AddScriptToEvaluateOnNewDocumentParams,
) <-chan *page.AddScriptToEvaluateOnNewDocumentResult {
	resultChan := make(chan *page.AddScriptToEvaluateOnNewDocumentResult)
	command := NewCommand(protocol.Socket, "Page.addScriptToEvaluateOnNewDocument", params)
	result := &page.AddScriptToEvaluateOnNewDocumentResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
BringToFront brings page to front (activates tab).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
*/
func (protocol *PageProtocol) BringToFront() <-chan *page.BringToFrontResult {
	resultChan := make(chan *page.BringToFrontResult)
	command := NewCommand(protocol.Socket, "Page.bringToFront", nil)
	result := &page.BringToFrontResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
CaptureScreenshot capture a page screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
func (protocol *PageProtocol) CaptureScreenshot(
	params *page.CaptureScreenshotParams,
) <-chan *page.CaptureScreenshotResult {
	resultChan := make(chan *page.CaptureScreenshotResult)
	command := NewCommand(protocol.Socket, "Page.captureScreenshot", params)
	result := &page.CaptureScreenshotResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
CreateIsolatedWorld creates an isolated world for the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
func (protocol *PageProtocol) CreateIsolatedWorld(
	params *page.CreateIsolatedWorldParams,
) <-chan *page.CreateIsolatedWorldResult {
	resultChan := make(chan *page.CreateIsolatedWorldResult)
	command := NewCommand(protocol.Socket, "Page.createIsolatedWorld", params)
	result := &page.CreateIsolatedWorldResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
*/
func (protocol *PageProtocol) Disable() <-chan *page.DisableResult {
	resultChan := make(chan *page.DisableResult)
	command := NewCommand(protocol.Socket, "Page.disable", nil)
	result := &page.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable Ennables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-enable
*/
func (protocol *PageProtocol) Enable() <-chan *page.EnableResult {
	resultChan := make(chan *page.EnableResult)
	command := NewCommand(protocol.Socket, "Page.enable", nil)
	result := &page.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetAppManifest gets the app manifest.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
*/
func (protocol *PageProtocol) GetAppManifest(
	params *page.GetAppManifestParams,
) <-chan *page.GetAppManifestResult {
	resultChan := make(chan *page.GetAppManifestResult)
	command := NewCommand(protocol.Socket, "Page.getAppManifest", params)
	result := &page.GetAppManifestResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetFrameTree returns present frame tree structure.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
*/
func (protocol *PageProtocol) GetFrameTree() <-chan *page.GetFrameTreeResult {
	resultChan := make(chan *page.GetFrameTreeResult)
	command := NewCommand(protocol.Socket, "Page.getFrameTree", nil)
	result := &page.GetFrameTreeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetLayoutMetrics returns metrics relating to the layouting of the page, such as
viewport bounds/scale.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
*/
func (protocol *PageProtocol) GetLayoutMetrics() <-chan *page.GetLayoutMetricsResult {
	resultChan := make(chan *page.GetLayoutMetricsResult)
	command := NewCommand(protocol.Socket, "Page.getLayoutMetrics", nil)
	result := &page.GetLayoutMetricsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetNavigationHistory returns navigation history for the current page.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
*/
func (protocol *PageProtocol) GetNavigationHistory() <-chan *page.GetNavigationHistoryResult {
	resultChan := make(chan *page.GetNavigationHistoryResult)
	command := NewCommand(protocol.Socket, "Page.getNavigationHistory", nil)
	result := &page.GetNavigationHistoryResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetResourceContent returns content of the given resource.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
EXPERIMENTAL.
*/
func (protocol *PageProtocol) GetResourceContent(
	params *page.GetResourceContentParams,
) <-chan *page.GetResourceContentResult {
	resultChan := make(chan *page.GetResourceContentResult)
	command := NewCommand(protocol.Socket, "Page.getResourceContent", params)
	result := &page.GetResourceContentResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetResourceTree returns present frame / resource tree structure.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
EXPERIMENTAL.
*/
func (protocol *PageProtocol) GetResourceTree() <-chan *page.GetResourceTreeResult {
	resultChan := make(chan *page.GetResourceTreeResult)
	command := NewCommand(protocol.Socket, "Page.getResourceTree", nil)
	result := &page.GetResourceTreeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog
(alert, confirm, prompt, or onbeforeunload).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
func (protocol *PageProtocol) HandleJavaScriptDialog(
	params *page.HandleJavaScriptDialogParams,
) <-chan *page.HandleJavaScriptDialogResult {
	resultChan := make(chan *page.HandleJavaScriptDialogResult)
	command := NewCommand(protocol.Socket, "Page.handleJavaScriptDialog", params)
	result := &page.HandleJavaScriptDialogResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Navigate navigates current page to the given URL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
func (protocol *PageProtocol) Navigate(
	params *page.NavigateParams,
) <-chan *page.NavigateResult {
	resultChan := make(chan *page.NavigateResult)
	command := NewCommand(protocol.Socket, "Page.navigate", params)
	result := &page.NavigateResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
NavigateToHistoryEntry navigates current page to the given history entry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
func (protocol *PageProtocol) NavigateToHistoryEntry(
	params *page.NavigateToHistoryEntryParams,
) <-chan *page.NavigateToHistoryEntryResult {
	resultChan := make(chan *page.NavigateToHistoryEntryResult)
	command := NewCommand(protocol.Socket, "Page.navigateToHistoryEntry", params)
	result := &page.NavigateToHistoryEntryResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
PrintToPDF print page as PDF.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
func (protocol *PageProtocol) PrintToPDF(
	params *page.PrintToPDFParams,
) <-chan *page.PrintToPDFResult {
	resultChan := make(chan *page.PrintToPDFResult)
	command := NewCommand(protocol.Socket, "Page.printToPDF", params)
	result := &page.PrintToPDFResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Reload reloads given page optionally ignoring the cache.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
func (protocol *PageProtocol) Reload(
	params *page.ReloadParams,
) <-chan *page.ReloadResult {
	resultChan := make(chan *page.ReloadResult)
	command := NewCommand(protocol.Socket, "Page.reload", params)
	result := &page.ReloadResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveScriptToEvaluateOnLoad deprecated, please use removeScriptToEvaluateOnNewDocument
instead.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
EXPERIMENTAL. DEPRECATED.
*/
func (protocol *PageProtocol) RemoveScriptToEvaluateOnLoad(
	params *page.RemoveScriptToEvaluateOnLoadParams,
) <-chan *page.RemoveScriptToEvaluateOnLoadResult {
	resultChan := make(chan *page.RemoveScriptToEvaluateOnLoadResult)
	command := NewCommand(protocol.Socket, "Page.removeScriptToEvaluateOnLoad", params)
	result := &page.RemoveScriptToEvaluateOnLoadResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveScriptToEvaluateOnNewDocument removes given script from the list.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
*/
func (protocol *PageProtocol) RemoveScriptToEvaluateOnNewDocument(
	params *page.RemoveScriptToEvaluateOnNewDocumentParams,
) <-chan *page.RemoveScriptToEvaluateOnNewDocumentResult {
	resultChan := make(chan *page.RemoveScriptToEvaluateOnNewDocumentResult)
	command := NewCommand(protocol.Socket, "Page.removeScriptToEvaluateOnNewDocument", params)
	result := &page.RemoveScriptToEvaluateOnNewDocumentResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RequestAppBanner is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-requestAppBanner EXPERIMENTAL.
*/
func (protocol *PageProtocol) RequestAppBanner() <-chan *page.RequestAppBannerResult {
	resultChan := make(chan *page.RequestAppBannerResult)
	command := NewCommand(protocol.Socket, "Page.requestAppBanner", nil)
	result := &page.RequestAppBannerResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
ScreencastFrameAck acknowledges that a screencast frame has been received by the
frontend.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck EXPERIMENTAL.
*/
func (protocol *PageProtocol) ScreencastFrameAck(
	params *page.ScreencastFrameAckParams,
) <-chan *page.ScreencastFrameAckResult {
	resultChan := make(chan *page.ScreencastFrameAckResult)
	command := NewCommand(protocol.Socket, "Page.screencastFrameAck", params)
	result := &page.ScreencastFrameAckResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SearchInResource searches for given string in resource content.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SearchInResource(
	params *page.SearchInResourceParams,
) <-chan *page.SearchInResourceResult {
	resultChan := make(chan *page.SearchInResourceResult)
	command := NewCommand(protocol.Socket, "Page.searchInResource", params)
	result := &page.SearchInResourceResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetAdBlockingEnabled(
	params *page.SetAdBlockingEnabledParams,
) <-chan *page.SetAdBlockingEnabledResult {
	resultChan := make(chan *page.SetAdBlockingEnabledResult)
	command := NewCommand(protocol.Socket, "Page.setAdBlockingEnabled", params)
	result := &page.SetAdBlockingEnabledResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetAutoAttachToCreatedPages controls whether browser will open a new inspector
window for connected pages.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAutoAttachToCreatedPages
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetAutoAttachToCreatedPages(
	params *page.SetAutoAttachToCreatedPagesParams,
) <-chan *page.SetAutoAttachToCreatedPagesResult {
	resultChan := make(chan *page.SetAutoAttachToCreatedPagesResult)
	command := NewCommand(protocol.Socket, "Page.setAutoAttachToCreatedPages", params)
	result := &page.SetAutoAttachToCreatedPagesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetDocumentContent sets given markup as the document's HTML.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
*/
func (protocol *PageProtocol) SetDocumentContent(
	params *page.SetDocumentContentParams,
) <-chan *page.SetDocumentContentResult {
	resultChan := make(chan *page.SetDocumentContentResult)
	command := NewCommand(protocol.Socket, "Page.setDocumentContent", params)
	result := &page.SetDocumentContentResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetDownloadBehavior sets the behavior when downloading a file.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetDownloadBehavior(
	params *page.SetDownloadBehaviorParams,
) <-chan *page.SetDownloadBehaviorResult {
	resultChan := make(chan *page.SetDownloadBehaviorResult)
	command := NewCommand(protocol.Socket, "Page.setDownloadBehavior", params)
	result := &page.SetDownloadBehaviorResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetLifecycleEventsEnabled controls whether page will emit lifecycle events.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetLifecycleEventsEnabled(
	params *page.SetLifecycleEventsEnabledParams,
) <-chan *page.SetLifecycleEventsEnabledResult {
	resultChan := make(chan *page.SetLifecycleEventsEnabledResult)
	command := NewCommand(protocol.Socket, "Page.setLifecycleEventsEnabled", params)
	result := &page.SetLifecycleEventsEnabledResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StartScreencast starts sending each frame using the `screencastFrame` event.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
EXPERIMENTAL.
*/
func (protocol *PageProtocol) StartScreencast(
	params *page.StartScreencastParams,
) <-chan *page.StartScreencastResult {
	resultChan := make(chan *page.StartScreencastResult)
	command := NewCommand(protocol.Socket, "Page.startScreencast", params)
	result := &page.StartScreencastResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StopLoading force the page stop all navigations and pending resource fetches.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopLoading
*/
func (protocol *PageProtocol) StopLoading() <-chan *page.StopLoadingResult {
	resultChan := make(chan *page.StopLoadingResult)
	command := NewCommand(protocol.Socket, "Page.stopLoading", nil)
	result := &page.StopLoadingResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StopScreencast stops sending each frame in the `screencastFrame`.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopScreencast
EXPERIMENTAL.
*/
func (protocol *PageProtocol) StopScreencast() <-chan *page.StopScreencastResult {
	resultChan := make(chan *page.StopScreencastResult)
	command := NewCommand(protocol.Socket, "Page.stopScreencast", nil)
	result := &page.StopScreencastResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnDOMContentEventFired adds a handler to the Page.domContentEventFired event.
Page.domContentEventFired fires when a content event occurs in the DOM.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-domContentEventFired
*/
func (protocol *PageProtocol) OnDOMContentEventFired(
	callback func(event *page.DOMContentEventFiredEvent),
) {
	handler := NewEventHandler(
		"Page.domContentEventFired",
		func(response *Response) {
			event := &page.DOMContentEventFiredEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameAttached adds a handler to the Page.frameAttached event. Page.frameAttached
fires when a frame has been attached to its parent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameAttached
*/
func (protocol *PageProtocol) OnFrameAttached(
	callback func(event *page.FrameAttachedEvent),
) {
	handler := NewEventHandler(
		"Page.frameAttached",
		func(response *Response) {
			event := &page.FrameAttachedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameClearedScheduledNavigation adds a handler to the Page.frameClearedScheduledNavigation
event. Page.frameClearedScheduledNavigation fires when a frame no longer has a
scheduled navigation.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameClearedScheduledNavigation
EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameClearedScheduledNavigation(
	callback func(event *page.FrameClearedScheduledNavigationEvent),
) {
	handler := NewEventHandler(
		"Page.frameClearedScheduledNavigation",
		func(response *Response) {
			event := &page.FrameClearedScheduledNavigationEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameDetached adds a handler to the Page.frameDetached event. Page.frameDetached
fires when a frame has been detached from its parent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameDetached
*/
func (protocol *PageProtocol) OnFrameDetached(
	callback func(event *page.FrameDetachedEvent),
) {
	handler := NewEventHandler(
		"Page.frameDetached",
		func(response *Response) {
			event := &page.FrameDetachedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameNavigated adds a handler to the Page.frameNavigated event. Page.frameNavigated
fires once navigation of the frame has completed. Frame is now associated with
the new loader.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameNavigated
*/
func (protocol *PageProtocol) OnFrameNavigated(
	callback func(event *page.FrameNavigatedEvent),
) {
	handler := NewEventHandler(
		"Page.frameNavigated",
		func(response *Response) {
			event := &page.FrameNavigatedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameResized adds a handler to the Page.frameResized event. Page.frameResized
fires when frame is resized.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameResized
EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameResized(
	callback func(event *page.FrameResizedEvent),
) {
	handler := NewEventHandler(
		"Page.frameResized",
		func(response *Response) {
			event := &page.FrameResizedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameScheduledNavigation adds a handler to the Page.frameScheduledNavigation
event. Page.frameScheduledNavigation fires when frame schedules a potential
navigation.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameScheduledNavigation
EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameScheduledNavigation(
	callback func(event *page.FrameScheduledNavigationEvent),
) {
	handler := NewEventHandler(
		"Page.frameScheduledNavigation",
		func(response *Response) {
			event := &page.FrameScheduledNavigationEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameStartedLoading adds a handler to the Page.frameStartedLoading event.
Page.frameStartedLoading fires when frame has started loading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStartedLoading
EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameStartedLoading(
	callback func(event *page.FrameStartedLoadingEvent),
) {
	handler := NewEventHandler(
		"Page.frameStartedLoading",
		func(response *Response) {
			event := &page.FrameStartedLoadingEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameStoppedLoading adds a handler to the Page.frameStoppedLoading event.
Page.frameStoppedLoading fires when frame has stopped loading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStoppedLoading
EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameStoppedLoading(
	callback func(event *page.FrameStoppedLoadingEvent),
) {
	handler := NewEventHandler(
		"Page.frameStoppedLoading",
		func(response *Response) {
			event := &page.FrameStoppedLoadingEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnInterstitialHidden adds a handler to the Page.interstitialHidden event.
Page.interstitialHidden fires when interstitial page was hidden.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialHidden
*/
func (protocol *PageProtocol) OnInterstitialHidden(
	callback func(event *page.InterstitialHiddenEvent),
) {
	handler := NewEventHandler(
		"Page.interstitialHidden",
		func(response *Response) {
			event := &page.InterstitialHiddenEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnInterstitialShown adds a handler to the Page.interstitialShown event.
Page.interstitialShown fires when interstitial page was shown.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialShown
*/
func (protocol *PageProtocol) OnInterstitialShown(
	callback func(event *page.InterstitialShownEvent),
) {
	handler := NewEventHandler(
		"Page.interstitialShown",
		func(response *Response) {
			event := &page.InterstitialShownEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnJavascriptDialogClosed adds a handler to the Page.javascriptDialogClosed
event. Page.javascriptDialogClosed fires when a JavaScript initiated dialog
(alert, confirm, prompt, or onbeforeunload) has been closed.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogClosed
*/
func (protocol *PageProtocol) OnJavascriptDialogClosed(
	callback func(event *page.JavascriptDialogClosedEvent),
) {
	handler := NewEventHandler(
		"Page.javascriptDialogClosed",
		func(response *Response) {
			event := &page.JavascriptDialogClosedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnJavascriptDialogOpening adds a handler to the Page.javascriptDialogOpening
event. Page.javascriptDialogOpening fires when a JavaScript initiated dialog
(alert, confirm, prompt, or onbeforeunload) is about to open.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogOpening
*/
func (protocol *PageProtocol) OnJavascriptDialogOpening(
	callback func(event *page.JavascriptDialogOpeningEvent),
) {
	handler := NewEventHandler(
		"Page.javascriptDialogOpening",
		func(response *Response) {
			event := &page.JavascriptDialogOpeningEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLifecycleEvent adds a handler to the Page.lifecycleEvent event. Page.lifecycleEvent
fires for top level page lifecycle events such as navigation, load, paint, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-lifecycleEvent
*/
func (protocol *PageProtocol) OnLifecycleEvent(
	callback func(event *page.LifecycleEventEvent),
) {
	handler := NewEventHandler(
		"Page.lifecycleEvent",
		func(response *Response) {
			event := &page.LifecycleEventEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLoadEventFired adds a handler to the Page.loadEventFired event. Page.loadEventFired
fires when the page has finished loading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-loadEventFired
*/
func (protocol *PageProtocol) OnLoadEventFired(
	callback func(event *page.LoadEventFiredEvent),
) {
	handler := NewEventHandler(
		"Page.loadEventFired",
		func(response *Response) {
			event := &page.LoadEventFiredEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnScreencastFrame adds a handler to the Page.screencastFrame event. Page.screencastFrame
fires when compressed image data is requested by the `startScreencast` method.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastFrame
EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnScreencastFrame(
	callback func(event *page.ScreencastFrameEvent),
) {
	handler := NewEventHandler(
		"Page.screencastFrame",
		func(response *Response) {
			event := &page.ScreencastFrameEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnScreencastVisibilityChanged adds a handler to the Page.screencastVisibilityChanged
event. Page.screencastVisibilityChanged fires when the page with currently
enabled screencast was shown or hidden.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastVisibilityChanged
EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnScreencastVisibilityChanged(
	callback func(event *page.ScreencastVisibilityChangedEvent),
) {
	handler := NewEventHandler(
		"Page.screencastVisibilityChanged",
		func(response *Response) {
			event := &page.ScreencastVisibilityChangedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWindowOpen adds a handler to the Page.windowOpen event. Page.windowOpen fires
when a new window is going to be opened, via window.open(), link click, form
submission, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-windowOpen
*/
func (protocol *PageProtocol) OnWindowOpen(
	callback func(event *page.WindowOpenEvent),
) {
	handler := NewEventHandler(
		"Page.windowOpen",
		func(response *Response) {
			event := &page.WindowOpenEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
