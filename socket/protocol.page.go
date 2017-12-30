package socket

import (
	"encoding/json"

	page "github.com/mkenney/go-chrome/protocol/page"
	log "github.com/sirupsen/logrus"
)

/*
PageProtocol provides a namespace for the Chrome Page protocol methods. The Page protocol provides
actions and events related to the inspected page belong to the page domain.

https://chromedevtools.github.io/devtools-protocol/tot/Page/
*/
type PageProtocol struct {
	Socket Socketer
}

/*
AddScriptToEvaluateOnLoad is eprecated, please use addScriptToEvaluateOnNewDocument instead.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
EXPERIMENTAL. DEPRECATED.
*/
func (protocol *PageProtocol) AddScriptToEvaluateOnLoad(
	params *page.AddScriptToEvaluateOnLoadParams,
) (*page.AddScriptToEvaluateOnLoadResult, error) {
	command := NewCommand("Page.addScriptToEvaluateOnLoad", params)
	result := &page.AddScriptToEvaluateOnLoadResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
AddScriptToEvaluateOnNewDocument evaluates given script in every frame upon creation (before loading
frame's scripts).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
func (protocol *PageProtocol) AddScriptToEvaluateOnNewDocument(
	params *page.AddScriptToEvaluateOnNewDocumentParams,
) (*page.AddScriptToEvaluateOnNewDocumentResult, error) {
	command := NewCommand("Page.addScriptToEvaluateOnNewDocument", params)
	result := &page.AddScriptToEvaluateOnNewDocumentResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
BringToFront brings page to front (activates tab).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
*/
func (protocol *PageProtocol) BringToFront() error {
	command := NewCommand("Page.bringToFront", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
CaptureScreenshot capture a page screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
func (protocol *PageProtocol) CaptureScreenshot(
	params *page.CaptureScreenshotParams,
) (*page.CaptureScreenshotResult, error) {
	command := NewCommand("Page.captureScreenshot", params)
	result := &page.CaptureScreenshotResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CreateIsolatedWorld creates an isolated world for the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
func (protocol *PageProtocol) CreateIsolatedWorld(
	params *page.CreateIsolatedWorldParams,
) (*page.CreateIsolatedWorldResult, error) {
	command := NewCommand("Page.createIsolatedWorld", params)
	result := &page.CreateIsolatedWorldResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Disable disables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
*/
func (protocol *PageProtocol) Disable() error {
	command := NewCommand("Page.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable Ennables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-enable
*/
func (protocol *PageProtocol) Enable() error {
	command := NewCommand("Page.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetAppManifest gets the app manifest.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
*/
func (protocol *PageProtocol) GetAppManifest(
	params *page.GetAppManifestParams,
) error {
	command := NewCommand("Page.getAppManifest", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetFrameTree returns present frame tree structure.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
*/
func (protocol *PageProtocol) GetFrameTree() (*page.GetFrameTreeResult, error) {
	command := NewCommand("Page.getFrameTree", nil)
	result := &page.GetFrameTreeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetLayoutMetrics returns metrics relating to the layouting of the page, such as viewport
bounds/scale.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
*/
func (protocol *PageProtocol) GetLayoutMetrics() (*page.GetLayoutMetricsResult, error) {
	command := NewCommand("Page.getLayoutMetrics", nil)
	result := &page.GetLayoutMetricsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetNavigationHistory returns navigation history for the current page.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
*/
func (protocol *PageProtocol) GetNavigationHistory() (*page.GetNavigationHistoryResult, error) {
	command := NewCommand("Page.getNavigationHistory", nil)
	result := &page.GetNavigationHistoryResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetResourceContent returns content of the given resource.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent EXPERIMENTAL.
*/
func (protocol *PageProtocol) GetResourceContent(
	params *page.GetResourceContentParams,
) (*page.GetResourceContentResult, error) {
	command := NewCommand("Page.getResourceContent", params)
	result := &page.GetResourceContentResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetResourceTree returns present frame / resource tree structure.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree EXPERIMENTAL.
*/
func (protocol *PageProtocol) GetResourceTree() (*page.GetResourceTreeResult, error) {
	command := NewCommand("Page.getResourceTree", nil)
	result := &page.GetResourceTreeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt,
or onbeforeunload).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
func (protocol *PageProtocol) HandleJavaScriptDialog(
	params *page.HandleJavaScriptDialogParams,
) error {
	command := NewCommand("Page.handleJavaScriptDialog", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Navigate navigates current page to the given URL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
func (protocol *PageProtocol) Navigate(
	params *page.NavigateParams,
) (*page.NavigateResult, error) {
	command := NewCommand("Page.navigate", params)
	result := &page.NavigateResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
NavigateToHistoryEntry navigates current page to the given history entry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
func (protocol *PageProtocol) NavigateToHistoryEntry(
	params *page.NavigateToHistoryEntryParams,
) error {
	command := NewCommand("Page.navigateToHistoryEntry", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
PrintToPDF print page as PDF.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
func (protocol *PageProtocol) PrintToPDF(
	params *page.PrintToPDFParams,
) (*page.PrintToPDFResult, error) {
	command := NewCommand("Page.printToPDF", params)
	result := &page.PrintToPDFResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Reload reloads given page optionally ignoring the cache.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
func (protocol *PageProtocol) Reload(
	params *page.ReloadParams,
) error {
	command := NewCommand("Page.reload", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveScriptToEvaluateOnLoad deprecated, please use removeScriptToEvaluateOnNewDocument instead.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
EXPERIMENTAL. DEPRECATED.
*/
func (protocol *PageProtocol) RemoveScriptToEvaluateOnLoad(
	params *page.RemoveScriptToEvaluateOnLoadParams,
) error {
	command := NewCommand("Page.removeScriptToEvaluateOnLoad", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveScriptToEvaluateOnNewDocument removes given script from the list.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
*/
func (protocol *PageProtocol) RemoveScriptToEvaluateOnNewDocument(
	params *page.RemoveScriptToEvaluateOnNewDocumentParams,
) error {
	command := NewCommand("Page.removeScriptToEvaluateOnNewDocument", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RequestAppBanner is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-requestAppBanner EXPERIMENTAL.
*/
func (protocol *PageProtocol) RequestAppBanner() error {
	command := NewCommand("Page.requestAppBanner", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ScreencastFrameAck acknowledges that a screencast frame has been received by the frontend.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck EXPERIMENTAL.
*/
func (protocol *PageProtocol) ScreencastFrameAck(
	params *page.ScreencastFrameAckParams,
) error {
	command := NewCommand("Page.screencastFrameAck", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SearchInResource searches for given string in resource content.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource EXPERIMENTAL.
*/
func (protocol *PageProtocol) SearchInResource(
	params *page.SearchInResourceParams,
) (*page.SearchInResourceResult, error) {
	command := NewCommand("Page.searchInResource", params)
	result := &page.SearchInResourceResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetAdBlockingEnabled(
	params *page.SetAdBlockingEnabledParams,
) error {
	command := NewCommand("Page.setAdBlockingEnabled", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetAutoAttachToCreatedPages controls whether browser will open a new inspector window for connected
pages.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAutoAttachToCreatedPages
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetAutoAttachToCreatedPages(
	params *page.SetAutoAttachToCreatedPagesParams,
) error {
	command := NewCommand("Page.setAutoAttachToCreatedPages", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetDocumentContent sets given markup as the document's HTML.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
*/
func (protocol *PageProtocol) SetDocumentContent(
	params *page.SetDocumentContentParams,
) error {
	command := NewCommand("Page.setDocumentContent", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetDownloadBehavior sets the behavior when downloading a file.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetDownloadBehavior(
	params *page.SetDownloadBehaviorParams,
) error {
	command := NewCommand("Page.setDownloadBehavior", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetLifecycleEventsEnabled controls whether page will emit lifecycle events.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
EXPERIMENTAL.
*/
func (protocol *PageProtocol) SetLifecycleEventsEnabled(
	params *page.SetLifecycleEventsEnabledParams,
) error {
	command := NewCommand("Page.setLifecycleEventsEnabled", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StartScreencast starts sending each frame using the `screencastFrame` event.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast EXPERIMENTAL.
*/
func (protocol *PageProtocol) StartScreencast(
	params *page.StartScreencastParams,
) error {
	command := NewCommand("Page.startScreencast", params)

	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopLoading force the page stop all navigations and pending resource fetches.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopLoading
*/
func (protocol *PageProtocol) StopLoading() error {
	command := NewCommand("Page.stopLoading", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopScreencast stops sending each frame in the `screencastFrame`.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopScreencast EXPERIMENTAL.
*/
func (protocol *PageProtocol) StopScreencast() error {
	command := NewCommand("Page.stopScreencast", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameAttached adds a handler to the Page.frameAttached event. Page.frameAttached fires when a
frame has been attached to its parent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameAttached
*/
func (protocol *PageProtocol) OnFrameAttached(
	callback func(event *page.FrameAttachedEvent),
) {
	handler := NewEventHandler(
		"Page.frameAttached",
		func(response *Response) {
			event := &page.FrameAttachedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameClearedScheduledNavigation adds a handler to the Page.frameClearedScheduledNavigation event.
Page.frameClearedScheduledNavigation fires when a frame no longer has a scheduled navigation.

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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameDetached adds a handler to the Page.frameDetached event. Page.frameDetached fires when a
frame has been detached from its parent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameDetached
*/
func (protocol *PageProtocol) OnFrameDetached(
	callback func(event *page.FrameDetachedEvent),
) {
	handler := NewEventHandler(
		"Page.frameDetached",
		func(response *Response) {
			event := &page.FrameDetachedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameNavigated adds a handler to the Page.frameNavigated event. Page.frameNavigated fires once
navigation of the frame has completed. Frame is now associated with the new loader.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameNavigated
*/
func (protocol *PageProtocol) OnFrameNavigated(
	callback func(event *page.FrameNavigatedEvent),
) {
	handler := NewEventHandler(
		"Page.frameNavigated",
		func(response *Response) {
			event := &page.FrameNavigatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameResized adds a handler to the Page.frameResized event. Page.frameResized fires when frame
is resized.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameResized EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameResized(
	callback func(event *page.FrameResizedEvent),
) {
	handler := NewEventHandler(
		"Page.frameResized",
		func(response *Response) {
			event := &page.FrameResizedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameScheduledNavigation adds a handler to the Page.frameScheduledNavigation event.
Page.frameScheduledNavigation fires when frame schedules a potential navigation.

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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameStartedLoading adds a handler to the Page.frameStartedLoading event. Page.frameStartedLoading
fires when frame has started loading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStartedLoading EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameStartedLoading(
	callback func(event *page.FrameStartedLoadingEvent),
) {
	handler := NewEventHandler(
		"Page.frameStartedLoading",
		func(response *Response) {
			event := &page.FrameStartedLoadingEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnFrameStoppedLoading adds a handler to the Page.frameStoppedLoading event. Page.frameStoppedLoading
fires when frame has stopped loading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStoppedLoading EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnFrameStoppedLoading(
	callback func(event *page.FrameStoppedLoadingEvent),
) {
	handler := NewEventHandler(
		"Page.frameStoppedLoading",
		func(response *Response) {
			event := &page.FrameStoppedLoadingEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnInterstitialHidden adds a handler to the Page.interstitialHidden event. Page.interstitialHidden
fires when interstitial page was hidden.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialHidden
*/
func (protocol *PageProtocol) OnInterstitialHidden(
	callback func(event *page.InterstitialHiddenEvent),
) {
	handler := NewEventHandler(
		"Page.interstitialHidden",
		func(response *Response) {
			event := &page.InterstitialHiddenEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnInterstitialShown adds a handler to the Page.interstitialShown event. Page.interstitialShown fires
when interstitial page was shown.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialShown
*/
func (protocol *PageProtocol) OnInterstitialShown(
	callback func(event *page.InterstitialShownEvent),
) {
	handler := NewEventHandler(
		"Page.interstitialShown",
		func(response *Response) {
			event := &page.InterstitialShownEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnJavascriptDialogClosed adds a handler to the Page.javascriptDialogClosed event.
Page.javascriptDialogClosed fires when a JavaScript initiated dialog (alert, confirm, prompt, or
onbeforeunload) has been closed.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogClosed
*/
func (protocol *PageProtocol) OnJavascriptDialogClosed(
	callback func(event *page.JavascriptDialogClosedEvent),
) {
	handler := NewEventHandler(
		"Page.javascriptDialogClosed",
		func(response *Response) {
			event := &page.JavascriptDialogClosedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnJavascriptDialogOpening adds a handler to the Page.javascriptDialogOpening event.
Page.javascriptDialogOpening fires when a JavaScript initiated dialog (alert, confirm, prompt, or
onbeforeunload) is about to open.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogOpening
*/
func (protocol *PageProtocol) OnJavascriptDialogOpening(
	callback func(event *page.JavascriptDialogOpeningEvent),
) {
	handler := NewEventHandler(
		"Page.javascriptDialogOpening",
		func(response *Response) {
			event := &page.JavascriptDialogOpeningEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLifecycleEvent adds a handler to the Page.lifecycleEvent event. Page.lifecycleEvent fires for top
level page lifecycle events such as navigation, load, paint, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-lifecycleEvent
*/
func (protocol *PageProtocol) OnLifecycleEvent(
	callback func(event *page.LifecycleEventEvent),
) {
	handler := NewEventHandler(
		"Page.lifecycleEvent",
		func(response *Response) {
			event := &page.LifecycleEventEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLoadEventFired adds a handler to the Page.loadEventFired event. Page.loadEventFired fires when the
page has finished loading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-loadEventFired
*/
func (protocol *PageProtocol) OnLoadEventFired(
	callback func(event *page.LoadEventFiredEvent),
) {
	handler := NewEventHandler(
		"Page.loadEventFired",
		func(response *Response) {
			event := &page.LoadEventFiredEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnScreencastFrame adds a handler to the Page.screencastFrame event. Page.screencastFrame fires when
compressed image data is requested by the `startScreencast` method.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastFrame EXPERIMENTAL.
*/
func (protocol *PageProtocol) OnScreencastFrame(
	callback func(event *page.ScreencastFrameEvent),
) {
	handler := NewEventHandler(
		"Page.screencastFrame",
		func(response *Response) {
			event := &page.ScreencastFrameEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnScreencastVisibilityChanged adds a handler to the Page.screencastVisibilityChanged event.
Page.screencastVisibilityChanged fires when the page with currently enabled screencast was shown or
hidden.

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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWindowOpen adds a handler to the Page.windowOpen event. Page.windowOpen fires when a new window is
going to be opened, via window.open(), link click, form submission, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-windowOpen
*/
func (protocol *PageProtocol) OnWindowOpen(
	callback func(event *page.WindowOpenEvent),
) {
	handler := NewEventHandler(
		"Page.windowOpen",
		func(response *Response) {
			event := &page.WindowOpenEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
