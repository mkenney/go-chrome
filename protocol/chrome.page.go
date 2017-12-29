package protocol

import (
	"encoding/json"

	page "github.com/mkenney/go-chrome/protocol/page"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Page provides a namespace for the Chrome Page protocol methods. The Page protocol provides actions
and events related to the inspected page belong to the page domain.

https://chromedevtools.github.io/devtools-protocol/tot/Page/
*/
var Page = PageProtocol{}

/*
PageProtocol defines the Page protocol methods.
*/
type PageProtocol struct{}

/*
AddScriptToEvaluateOnLoad is eprecated, please use addScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
*/
func (PageProtocol) AddScriptToEvaluateOnLoad(
	socket sock.Socketer,
	params *page.AddScriptToEvaluateOnLoadParams,
) (*page.AddScriptToEvaluateOnLoadResult, error) {
	command := sock.NewCommand("Page.addScriptToEvaluateOnLoad", params)
	result := &page.AddScriptToEvaluateOnLoadResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
AddScriptToEvaluateOnNewDocument evaluates given script in every frame upon creation (before loading
frame's scripts).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
func (PageProtocol) AddScriptToEvaluateOnNewDocument(
	socket sock.Socketer,
	params *page.AddScriptToEvaluateOnNewDocumentParams,
) (*page.AddScriptToEvaluateOnNewDocumentResult, error) {
	command := sock.NewCommand("Page.addScriptToEvaluateOnNewDocument", params)
	result := &page.AddScriptToEvaluateOnNewDocumentResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
BringToFront brings page to front (activates tab).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
*/
func (PageProtocol) BringToFront(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Page.bringToFront", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
CaptureScreenshot capture a page screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
func (PageProtocol) CaptureScreenshot(
	socket sock.Socketer,
	params *page.CaptureScreenshotParams,
) (*page.CaptureScreenshotResult, error) {
	command := sock.NewCommand("Page.captureScreenshot", params)
	result := &page.CaptureScreenshotResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
CreateIsolatedWorld creates an isolated world for the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
func (PageProtocol) CreateIsolatedWorld(
	socket sock.Socketer,
	params *page.CreateIsolatedWorldParams,
) (*page.CreateIsolatedWorldResult, error) {
	command := sock.NewCommand("Page.createIsolatedWorld", params)
	result := &page.CreateIsolatedWorldResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
Disable disables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
*/
func (PageProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Page.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable Ennables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-enable
*/
func (PageProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Page.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetAppManifest gets the app manifest.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
*/
func (PageProtocol) GetAppManifest(
	socket sock.Socketer,
	params *page.GetAppManifestParams,
) error {
	command := sock.NewCommand("Page.getAppManifest", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
GetFrameTree returns present frame tree structure.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
*/
func (PageProtocol) GetFrameTree(
	socket sock.Socketer,
) (*page.GetFrameTreeResult, error) {
	command := sock.NewCommand("Page.getFrameTree", nil)
	result := &page.GetFrameTreeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetLayoutMetrics returns metrics relating to the layouting of the page, such as viewport
bounds/scale.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
*/
func (PageProtocol) GetLayoutMetrics(
	socket sock.Socketer,
) (*page.GetLayoutMetricsResult, error) {
	command := sock.NewCommand("Page.getLayoutMetrics", nil)
	result := &page.GetLayoutMetricsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetNavigationHistory returns navigation history for the current page.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
*/
func (PageProtocol) GetNavigationHistory(
	socket sock.Socketer,
) (*page.GetNavigationHistoryResult, error) {
	command := sock.NewCommand("Page.getNavigationHistory", nil)
	result := &page.GetNavigationHistoryResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetResourceContent returns content of the given resource. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
*/
func (PageProtocol) GetResourceContent(
	socket sock.Socketer,
	params *page.GetResourceContentParams,
) (*page.GetResourceContentResult, error) {
	command := sock.NewCommand("Page.getResourceContent", params)
	result := &page.GetResourceContentResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetResourceTree returns present frame / resource tree structure. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
*/
func (PageProtocol) GetResourceTree(
	socket sock.Socketer,
) (*page.GetResourceTreeResult, error) {
	command := sock.NewCommand("Page.getResourceTree", nil)
	result := &page.GetResourceTreeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt,
or onbeforeunload).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
func (PageProtocol) HandleJavaScriptDialog(
	socket sock.Socketer,
	params *page.HandleJavaScriptDialogParams,
) error {
	command := sock.NewCommand("Page.handleJavaScriptDialog", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
Navigate navigates current page to the given URL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
func (PageProtocol) Navigate(
	socket sock.Socketer,
	params *page.NavigateParams,
) (*page.NavigateResult, error) {
	command := sock.NewCommand("Page.navigate", params)
	result := &page.NavigateResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
NavigateToHistoryEntry navigates current page to the given history entry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
func (PageProtocol) NavigateToHistoryEntry(
	socket sock.Socketer,
	params *page.NavigateToHistoryEntryParams,
) error {
	command := sock.NewCommand("Page.navigateToHistoryEntry", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
PrintToPDF print page as PDF.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
func (PageProtocol) PrintToPDF(
	socket sock.Socketer,
	params *page.PrintToPDFParams,
) (*page.PrintToPDFResult, error) {
	command := sock.NewCommand("Page.printToPDF", params)
	result := &page.PrintToPDFResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
Reload reloads given page optionally ignoring the cache.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
func (PageProtocol) Reload(
	socket sock.Socketer,
	params *page.ReloadParams,
) error {
	command := sock.NewCommand("Page.reload", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveScriptToEvaluateOnLoad deprecated, please use removeScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
*/
func (PageProtocol) RemoveScriptToEvaluateOnLoad(
	socket sock.Socketer,
	params *page.RemoveScriptToEvaluateOnLoadParams,
) error {
	command := sock.NewCommand("Page.removeScriptToEvaluateOnLoad", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveScriptToEvaluateOnNewDocument removes given script from the list.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
*/
func (PageProtocol) RemoveScriptToEvaluateOnNewDocument(
	socket sock.Socketer,
	params *page.RemoveScriptToEvaluateOnNewDocumentParams,
) error {
	command := sock.NewCommand("Page.removeScriptToEvaluateOnNewDocument", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
RequestAppBanner EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-requestAppBanner
*/
func (PageProtocol) RequestAppBanner(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Page.requestAppBanner", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ScreencastFrameAck acknowledges that a screencast frame has been received by the frontend.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck
*/
func (PageProtocol) ScreencastFrameAck(
	socket sock.Socketer,
	params *page.ScreencastFrameAckParams,
) error {
	command := sock.NewCommand("Page.screencastFrameAck", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
SearchInResource searches for given string in resource content. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
*/
func (PageProtocol) SearchInResource(
	socket sock.Socketer,
	params *page.SearchInResourceParams,
) (*page.SearchInResourceResult, error) {
	command := sock.NewCommand("Page.searchInResource", params)
	result := &page.SearchInResourceResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
*/
func (PageProtocol) SetAdBlockingEnabled(
	socket sock.Socketer,
	params *page.SetAdBlockingEnabledParams,
) error {
	command := sock.NewCommand("Page.setAdBlockingEnabled", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
SetAutoAttachToCreatedPages controls whether browser will open a new inspector window for connected
pages. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAutoAttachToCreatedPages
*/
func (PageProtocol) SetAutoAttachToCreatedPages(
	socket sock.Socketer,
	params *page.SetAutoAttachToCreatedPagesParams,
) error {
	command := sock.NewCommand("Page.setAutoAttachToCreatedPages", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
SetDocumentContent sets given markup as the document's HTML.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
*/
func (PageProtocol) SetDocumentContent(
	socket sock.Socketer,
	params *page.SetDocumentContentParams,
) error {
	command := sock.NewCommand("Page.setDocumentContent", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
SetDownloadBehavior sets the behavior when downloading a file. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
*/
func (PageProtocol) SetDownloadBehavior(
	socket sock.Socketer,
	params *page.SetDownloadBehaviorParams,
) error {
	command := sock.NewCommand("Page.setDownloadBehavior", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
SetLifecycleEventsEnabled controls whether page will emit lifecycle events. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
*/
func (PageProtocol) SetLifecycleEventsEnabled(
	socket sock.Socketer,
	params *page.SetLifecycleEventsEnabledParams,
) error {
	command := sock.NewCommand("Page.setLifecycleEventsEnabled", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
StartScreencast starts sending each frame using the `screencastFrame` event. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
*/
func (PageProtocol) StartScreencast(
	socket sock.Socketer,
	params *page.StartScreencastParams,
) error {
	command := sock.NewCommand("Page.startScreencast", params)

	socket.SendCommand(command)
	return command.Error()
}

/*
StopLoading force the page stop all navigations and pending resource fetches.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopLoading
*/
func (PageProtocol) StopLoading(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Page.stopLoading", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopScreencast stops sending each frame in the `screencastFrame`. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopScreencast
*/
func (PageProtocol) StopScreencast(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Page.stopScreencast", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnDOMContentEventFired adds a handler to the Page.domContentEventFired event.
Page.domContentEventFired fires when a content event occurs in the DOM.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-domContentEventFired
*/
func (PageProtocol) OnDOMContentEventFired(
	socket sock.Socketer,
	callback func(event *page.DOMContentEventFiredEvent),
) {
	handler := sock.NewEventHandler(
		"Page.domContentEventFired",
		func(response *sock.Response) {
			event := &page.DOMContentEventFiredEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameAttached adds a handler to the Page.frameAttached event. Page.frameAttached fires when a
frame has been attached to its parent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameAttached
*/
func (PageProtocol) OnFrameAttached(
	socket sock.Socketer,
	callback func(event *page.FrameAttachedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameAttached",
		func(response *sock.Response) {
			event := &page.FrameAttachedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameClearedScheduledNavigation adds a handler to the Page.frameClearedScheduledNavigation event.
Page.frameClearedScheduledNavigation fires when a frame no longer has a scheduled navigation.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameClearedScheduledNavigation
*/
func (PageProtocol) OnFrameClearedScheduledNavigation(
	socket sock.Socketer,
	callback func(event *page.FrameClearedScheduledNavigationEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameClearedScheduledNavigation",
		func(response *sock.Response) {
			event := &page.FrameClearedScheduledNavigationEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameDetached adds a handler to the Page.frameDetached event. Page.frameDetached fires when a
frame has been detached from its parent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameDetached
*/
func (PageProtocol) OnFrameDetached(
	socket sock.Socketer,
	callback func(event *page.FrameDetachedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameDetached",
		func(response *sock.Response) {
			event := &page.FrameDetachedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameNavigated adds a handler to the Page.frameNavigated event. Page.frameNavigated fires once
navigation of the frame has completed. Frame is now associated with the new loader.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameNavigated
*/
func (PageProtocol) OnFrameNavigated(
	socket sock.Socketer,
	callback func(event *page.FrameNavigatedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameNavigated",
		func(response *sock.Response) {
			event := &page.FrameNavigatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameResized adds a handler to the Page.frameResized event. Page.frameResized fires when frame
is resized. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameResized
*/
func (PageProtocol) OnFrameResized(
	socket sock.Socketer,
	callback func(event *page.FrameResizedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameResized",
		func(response *sock.Response) {
			event := &page.FrameResizedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameScheduledNavigation adds a handler to the Page.frameScheduledNavigation event.
Page.frameScheduledNavigation fires when frame schedules a potential navigation. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameScheduledNavigation
*/
func (PageProtocol) OnFrameScheduledNavigation(
	socket sock.Socketer,
	callback func(event *page.FrameScheduledNavigationEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameScheduledNavigation",
		func(response *sock.Response) {
			event := &page.FrameScheduledNavigationEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameStartedLoading adds a handler to the Page.frameStartedLoading event. Page.frameStartedLoading
fires when frame has started loading. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStartedLoading
*/
func (PageProtocol) OnFrameStartedLoading(
	socket sock.Socketer,
	callback func(event *page.FrameStartedLoadingEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameStartedLoading",
		func(response *sock.Response) {
			event := &page.FrameStartedLoadingEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnFrameStoppedLoading adds a handler to the Page.frameStoppedLoading event. Page.frameStoppedLoading
fires when frame has stopped loading. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStoppedLoading
*/
func (PageProtocol) OnFrameStoppedLoading(
	socket sock.Socketer,
	callback func(event *page.FrameStoppedLoadingEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameStoppedLoading",
		func(response *sock.Response) {
			event := &page.FrameStoppedLoadingEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnInterstitialHidden adds a handler to the Page.interstitialHidden event. Page.interstitialHidden
fires when interstitial page was hidden.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialHidden
*/
func (PageProtocol) OnInterstitialHidden(
	socket sock.Socketer,
	callback func(event *page.InterstitialHiddenEvent),
) {
	handler := sock.NewEventHandler(
		"Page.interstitialHidden",
		func(response *sock.Response) {
			event := &page.InterstitialHiddenEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnInterstitialShown adds a handler to the Page.interstitialShown event. Page.interstitialShown fires
when interstitial page was shown.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialShown
*/
func (PageProtocol) OnInterstitialShown(
	socket sock.Socketer,
	callback func(event *page.InterstitialShownEvent),
) {
	handler := sock.NewEventHandler(
		"Page.interstitialShown",
		func(response *sock.Response) {
			event := &page.InterstitialShownEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnJavascriptDialogClosed adds a handler to the Page.javascriptDialogClosed event.
Page.javascriptDialogClosed fires when a JavaScript initiated dialog (alert, confirm, prompt, or
onbeforeunload) has been closed.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogClosed
*/
func (PageProtocol) OnJavascriptDialogClosed(
	socket sock.Socketer,
	callback func(event *page.JavascriptDialogClosedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.javascriptDialogClosed",
		func(response *sock.Response) {
			event := &page.JavascriptDialogClosedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnJavascriptDialogOpening adds a handler to the Page.javascriptDialogOpening event.
Page.javascriptDialogOpening fires when a JavaScript initiated dialog (alert, confirm, prompt, or
onbeforeunload) is about to open.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogOpening
*/
func (PageProtocol) OnJavascriptDialogOpening(
	socket sock.Socketer,
	callback func(event *page.JavascriptDialogOpeningEvent),
) {
	handler := sock.NewEventHandler(
		"Page.javascriptDialogOpening",
		func(response *sock.Response) {
			event := &page.JavascriptDialogOpeningEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLifecycleEvent adds a handler to the Page.lifecycleEvent event. Page.lifecycleEvent fires for top
level page lifecycle events such as navigation, load, paint, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-lifecycleEvent
*/
func (PageProtocol) OnLifecycleEvent(
	socket sock.Socketer,
	callback func(event *page.LifecycleEventEvent),
) {
	handler := sock.NewEventHandler(
		"Page.lifecycleEvent",
		func(response *sock.Response) {
			event := &page.LifecycleEventEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLoadEventFired adds a handler to the Page.loadEventFired event. Page.loadEventFired fires when the
page has finished loading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-loadEventFired
*/
func (PageProtocol) OnLoadEventFired(
	socket sock.Socketer,
	callback func(event *page.LoadEventFiredEvent),
) {
	handler := sock.NewEventHandler(
		"Page.loadEventFired",
		func(response *sock.Response) {
			event := &page.LoadEventFiredEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScreencastFrame adds a handler to the Page.screencastFrame event. Page.screencastFrame fires when
compressed image data is requested by the `startScreencast` method. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastFrame
*/
func (PageProtocol) OnScreencastFrame(
	socket sock.Socketer,
	callback func(event *page.ScreencastFrameEvent),
) {
	handler := sock.NewEventHandler(
		"Page.screencastFrame",
		func(response *sock.Response) {
			event := &page.ScreencastFrameEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScreencastVisibilityChanged adds a handler to the Page.screencastVisibilityChanged event.
Page.screencastVisibilityChanged fires when the page with currently enabled screencast was shown or
hidden. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastVisibilityChanged
*/
func (PageProtocol) OnScreencastVisibilityChanged(
	socket sock.Socketer,
	callback func(event *page.ScreencastVisibilityChangedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.screencastVisibilityChanged",
		func(response *sock.Response) {
			event := &page.ScreencastVisibilityChangedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWindowOpen adds a handler to the Page.windowOpen event. Page.windowOpen fires when a new window is
going to be opened, via window.open(), link click, form submission, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-windowOpen
*/
func (PageProtocol) OnWindowOpen(
	socket sock.Socketer,
	callback func(event *page.WindowOpenEvent),
) {
	handler := sock.NewEventHandler(
		"Page.windowOpen",
		func(response *sock.Response) {
			event := &page.WindowOpenEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
