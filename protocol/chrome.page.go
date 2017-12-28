package protocol

import (
	"encoding/json"

	page "github.com/mkenney/go-chrome/protocol/page"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Page is a struct that provides a namespace for the Chrome Page protocol methods. The Page protocol
provides actions and events related to the inspected page belong to the page domain.

https://chromedevtools.github.io/devtools-protocol/tot/Page/
*/
var Page = _page{}

type _page struct{}

/*
AddScriptToEvaluateOnLoad is eprecated, please use addScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
*/
func (_page) AddScriptToEvaluateOnLoad(
	socket sock.Socketer,
	params *page.AddScriptToEvaluateOnLoadParams,
) (page.AddScriptToEvaluateOnLoadResult, error) {
	command := sock.NewCommand("Page.addScriptToEvaluateOnLoad", params)

	result := page.AddScriptToEvaluateOnLoadResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
AddScriptToEvaluateOnNewDocument evaluates given script in every frame upon creation (before loading
frame's scripts).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
func (_page) AddScriptToEvaluateOnNewDocument(
	socket sock.Socketer,
	params *page.AddScriptToEvaluateOnNewDocumentParams,
) (page.AddScriptToEvaluateOnNewDocumentResult, error) {
	command := sock.NewCommand("Page.addScriptToEvaluateOnNewDocument", params)

	result := page.AddScriptToEvaluateOnNewDocumentResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
BringToFront brings page to front (activates tab).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
*/
func (_page) BringToFront(
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
func (_page) CaptureScreenshot(
	socket sock.Socketer,
	params *page.CaptureScreenshotParams,
) (page.CaptureScreenshotResult, error) {
	command := sock.NewCommand("Page.captureScreenshot", params)

	result := page.CaptureScreenshotResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
CreateIsolatedWorld creates an isolated world for the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
func (_page) CreateIsolatedWorld(
	socket sock.Socketer,
	params *page.CreateIsolatedWorldParams,
) (page.CreateIsolatedWorldResult, error) {
	command := sock.NewCommand("Page.createIsolatedWorld", params)

	result := page.CreateIsolatedWorldResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Disable disables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
*/
func (_page) Disable(
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
func (_page) Enable(
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
func (_page) GetAppManifest(
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
func (_page) GetFrameTree(
	socket sock.Socketer,
) (page.GetFrameTreeResult, error) {
	command := sock.NewCommand("Page.getFrameTree", nil)
	result := page.GetFrameTreeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetLayoutMetrics returns metrics relating to the layouting of the page, such as viewport
bounds/scale.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
*/
func (_page) GetLayoutMetrics(
	socket sock.Socketer,
) (page.GetLayoutMetricsResult, error) {
	command := sock.NewCommand("Page.getLayoutMetrics", nil)
	result := page.GetLayoutMetricsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetNavigationHistory returns navigation history for the current page.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
*/
func (_page) GetNavigationHistory(
	socket sock.Socketer,
) (page.GetNavigationHistoryResult, error) {
	command := sock.NewCommand("Page.getNavigationHistory", nil)
	result := page.GetNavigationHistoryResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetResourceContent returns content of the given resource. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
*/
func (_page) GetResourceContent(
	socket sock.Socketer,
	params *page.GetResourceContentParams,
) (page.GetResourceContentResult, error) {
	command := sock.NewCommand("Page.getResourceContent", params)

	result := page.GetResourceContentResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetResourceTree returns present frame / resource tree structure. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
*/
func (_page) GetResourceTree(
	socket sock.Socketer,
) (page.GetResourceTreeResult, error) {
	command := sock.NewCommand("Page.getResourceTree", nil)
	result := page.GetResourceTreeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt,
or onbeforeunload).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
func (_page) HandleJavaScriptDialog(
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
func (_page) Navigate(
	socket sock.Socketer,
	params *page.NavigateParams,
) (page.NavigateResult, error) {
	command := sock.NewCommand("Page.navigate", params)

	result := page.NavigateResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
NavigateToHistoryEntry navigates current page to the given history entry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
func (_page) NavigateToHistoryEntry(
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
func (_page) PrintToPDF(
	socket sock.Socketer,
	params *page.PrintToPDFParams,
) (page.PrintToPDFResult, error) {
	command := sock.NewCommand("Page.printToPDF", params)

	result := page.PrintToPDFResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Reload reloads given page optionally ignoring the cache.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
func (_page) Reload(
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
func (_page) RemoveScriptToEvaluateOnLoad(
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
func (_page) RemoveScriptToEvaluateOnNewDocument(
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
func (_page) RequestAppBanner(
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
func (_page) ScreencastFrameAck(
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
func (_page) SearchInResource(
	socket sock.Socketer,
	params *page.SearchInResourceParams,
) (page.SearchInResourceResult, error) {
	command := sock.NewCommand("Page.searchInResource", params)

	result := page.SearchInResourceResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
*/
func (_page) SetAdBlockingEnabled(
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
func (_page) SetAutoAttachToCreatedPages(
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
func (_page) SetDocumentContent(
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
func (_page) SetDownloadBehavior(
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
func (_page) SetLifecycleEventsEnabled(
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
func (_page) StartScreencast(
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
func (_page) StopLoading(
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
func (_page) StopScreencast(
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
func (_page) OnDOMContentEventFired(
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
func (_page) OnFrameAttached(
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
func (_page) OnFrameClearedScheduledNavigation(
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
func (_page) OnFrameDetached(
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
func (_page) OnFrameNavigated(
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
func (_page) OnFrameResized(
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
func (_page) OnFrameScheduledNavigation(
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
func (_page) OnFrameStartedLoading(
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
func (_page) OnFrameStoppedLoading(
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
func (_page) OnInterstitialHidden(
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
func (_page) OnInterstitialShown(
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
func (_page) OnJavascriptDialogClosed(
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
func (_page) OnJavascriptDialogOpening(
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
func (_page) OnLifecycleEvent(
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
func (_page) OnLoadEventFired(
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
func (_page) OnScreencastFrame(
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
func (_page) OnScreencastVisibilityChanged(
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
func (_page) OnWindowOpen(
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
