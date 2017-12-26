package protocol

import (
	"encoding/json"

	page "github.com/mkenney/go-chrome/protocol/page"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Page - https://chromedevtools.github.io/devtools-protocol/tot/Page/
Actions and events related to the inspected page belong to the page domain.
*/
type Page struct{}

/*
AddScriptToEvaluateOnLoad is eprecated, please use addScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
*/
func (Page) AddScriptToEvaluateOnLoad(
	socket *sock.Socket,
	params *page.AddScriptToEvaluateOnLoadParams,
) (page.AddScriptToEvaluateOnLoadResult, error) {
	command := &sock.Command{
		Method: "Page.addScriptToEvaluateOnLoad",
		Params: params,
	}
	result := page.AddScriptToEvaluateOnLoadResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
AddScriptToEvaluateOnNewDocument evaluates given script in every frame upon creation (before loading
frame's scripts).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
func (Page) AddScriptToEvaluateOnNewDocument(
	socket *sock.Socket,
	params *page.AddScriptToEvaluateOnNewDocumentParams,
) (page.AddScriptToEvaluateOnNewDocumentResult, error) {
	command := &sock.Command{
		Method: "Page.addScriptToEvaluateOnNewDocument",
		Params: params,
	}
	result := page.AddScriptToEvaluateOnNewDocumentResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
BringToFront brings page to front (activates tab).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
*/
func (Page) BringToFront(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Page.bringToFront",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CaptureScreenshot capture a page screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
func (Page) CaptureScreenshot(
	socket *sock.Socket,
	params *page.CaptureScreenshotParams,
) (page.CaptureScreenshotResult, error) {
	command := &sock.Command{
		Method: "Page.captureScreenshot",
		Params: params,
	}
	result := page.CaptureScreenshotResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
CreateIsolatedWorld creates an isolated world for the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
func (Page) CreateIsolatedWorld(
	socket *sock.Socket,
	params *page.CreateIsolatedWorldParams,
) (page.CreateIsolatedWorldResult, error) {
	command := &sock.Command{
		Method: "Page.createIsolatedWorld",
		Params: params,
	}
	result := page.CreateIsolatedWorldResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
Disable disables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
*/
func (Page) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Page.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable Ennables page domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-enable
*/
func (Page) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Page.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetAppManifest gets the app manifest.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
*/
func (Page) GetAppManifest(
	socket *sock.Socket,
	params *page.GetAppManifestParams,
) error {
	command := &sock.Command{
		Method: "Page.getAppManifest",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetFrameTree returns present frame tree structure.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
*/
func (Page) GetFrameTree(
	socket *sock.Socket,
) (page.GetFrameTreeResult, error) {
	command := &sock.Command{
		Method: "Page.getFrameTree",
	}
	result := page.GetFrameTreeResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetLayoutMetrics returns metrics relating to the layouting of the page, such as viewport
bounds/scale.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
*/
func (Page) GetLayoutMetrics(
	socket *sock.Socket,
) (page.GetLayoutMetricsResult, error) {
	command := &sock.Command{
		Method: "Page.getLayoutMetrics",
	}
	result := page.GetLayoutMetricsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetNavigationHistory returns navigation history for the current page.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
*/
func (Page) GetNavigationHistory(
	socket *sock.Socket,
) (page.GetNavigationHistoryResult, error) {
	command := &sock.Command{
		Method: "Page.getNavigationHistory",
	}
	result := page.GetNavigationHistoryResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetResourceContent returns content of the given resource. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
*/
func (Page) GetResourceContent(
	socket *sock.Socket,
	params *page.GetResourceContentParams,
) (page.GetResourceContentResult, error) {
	command := &sock.Command{
		Method: "Page.getResourceContent",
		Params: params,
	}
	result := page.GetResourceContentResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetResourceTree returns present frame / resource tree structure. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
*/
func (Page) GetResourceTree(
	socket *sock.Socket,
) (page.GetResourceTreeResult, error) {
	command := &sock.Command{
		Method: "Page.getResourceTree",
	}
	result := page.GetResourceTreeResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt,
or onbeforeunload).

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
func (Page) HandleJavaScriptDialog(
	socket *sock.Socket,
	params *page.HandleJavaScriptDialogParams,
) error {
	command := &sock.Command{
		Method: "Page.handleJavaScriptDialog",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Navigate navigates current page to the given URL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
func (Page) Navigate(
	socket *sock.Socket,
	params *page.NavigateParams,
) (page.NavigateResult, error) {
	command := &sock.Command{
		Method: "Page.navigate",
		Params: params,
	}
	result := page.NavigateResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
NavigateToHistoryEntry navigates current page to the given history entry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
func (Page) NavigateToHistoryEntry(
	socket *sock.Socket,
	params *page.NavigateToHistoryEntryParams,
) error {
	command := &sock.Command{
		Method: "Page.navigateToHistoryEntry",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
PrintToPDF print page as PDF.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
func (Page) PrintToPDF(
	socket *sock.Socket,
	params *page.PrintToPDFParams,
) (page.PrintToPDFResult, error) {
	command := &sock.Command{
		Method: "Page.printToPDF",
		Params: params,
	}
	result := page.PrintToPDFResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
Reload reloads given page optionally ignoring the cache.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
func (Page) Reload(
	socket *sock.Socket,
	params *page.ReloadParams,
) error {
	command := &sock.Command{
		Method: "Page.reload",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveScriptToEvaluateOnLoad deprecated, please use removeScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
*/
func (Page) RemoveScriptToEvaluateOnLoad(
	socket *sock.Socket,
	params *page.RemoveScriptToEvaluateOnLoadParams,
) error {
	command := &sock.Command{
		Method: "Page.removeScriptToEvaluateOnLoad",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveScriptToEvaluateOnNewDocument removes given script from the list.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
*/
func (Page) RemoveScriptToEvaluateOnNewDocument(
	socket *sock.Socket,
	params *page.RemoveScriptToEvaluateOnNewDocumentParams,
) error {
	command := &sock.Command{
		Method: "Page.removeScriptToEvaluateOnNewDocument",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestAppBanner EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-requestAppBanner
*/
func (Page) RequestAppBanner(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Page.requestAppBanner",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ScreencastFrameAck acknowledges that a screencast frame has been received by the frontend.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck
*/
func (Page) ScreencastFrameAck(
	socket *sock.Socket,
	params *page.ScreencastFrameAckParams,
) error {
	command := &sock.Command{
		Method: "Page.screencastFrameAck",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SearchInResource searches for given string in resource content. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
*/
func (Page) SearchInResource(
	socket *sock.Socket,
	params *page.SearchInResourceParams,
) (page.SearchInResourceResult, error) {
	command := &sock.Command{
		Method: "Page.searchInResource",
		Params: params,
	}
	result := page.SearchInResourceResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
*/
func (Page) SetAdBlockingEnabled(
	socket *sock.Socket,
	params *page.SetAdBlockingEnabledParams,
) error {
	command := &sock.Command{
		Method: "Page.setAdBlockingEnabled",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAutoAttachToCreatedPages controls whether browser will open a new inspector window for connected
pages. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAutoAttachToCreatedPages
*/
func (Page) SetAutoAttachToCreatedPages(
	socket *sock.Socket,
	params *page.SetAutoAttachToCreatedPagesParams,
) error {
	command := &sock.Command{
		Method: "Page.setAutoAttachToCreatedPages",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDocumentContent sets given markup as the document's HTML.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
*/
func (Page) SetDocumentContent(
	socket *sock.Socket,
	params *page.SetDocumentContentParams,
) error {
	command := &sock.Command{
		Method: "Page.setDocumentContent",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDownloadBehavior sets the behavior when downloading a file. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
*/
func (Page) SetDownloadBehavior(
	socket *sock.Socket,
	params *page.SetDownloadBehaviorParams,
) error {
	command := &sock.Command{
		Method: "Page.setDownloadBehavior",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetLifecycleEventsEnabled controls whether page will emit lifecycle events. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
*/
func (Page) SetLifecycleEventsEnabled(
	socket *sock.Socket,
	params *page.SetLifecycleEventsEnabledParams,
) error {
	command := &sock.Command{
		Method: "Page.setLifecycleEventsEnabled",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartScreencast starts sending each frame using the `screencastFrame` event. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
*/
func (Page) StartScreencast(
	socket *sock.Socket,
	params *page.StartScreencastParams,
) error {
	command := &sock.Command{
		Method: "Page.startScreencast",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopLoading force the page stop all navigations and pending resource fetches.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopLoading
*/
func (Page) StopLoading(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Page.stopLoading",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopScreencast stops sending each frame in the `screencastFrame`. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopScreencast
*/
func (Page) StopScreencast(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Page.stopScreencast",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnDOMContentEventFired adds a handler to the Page.domContentEventFired event.
Page.domContentEventFired fires when a content event occurs in the DOM.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-domContentEventFired
*/
func (Page) OnDOMContentEventFired(
	socket *sock.Socket,
	callback func(event *page.DOMContentEventFiredEvent),
) {
	handler := sock.NewEventHandler(
		"Page.domContentEventFired",
		func(name string, params []byte) {
			event := &page.DOMContentEventFiredEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnFrameAttached(
	socket *sock.Socket,
	callback func(event *page.FrameAttachedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameAttached",
		func(name string, params []byte) {
			event := &page.FrameAttachedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameClearedScheduledNavigation
*/
func (Page) OnFrameClearedScheduledNavigation(
	socket *sock.Socket,
	callback func(event *page.FrameClearedScheduledNavigationEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameClearedScheduledNavigation",
		func(name string, params []byte) {
			event := &page.FrameClearedScheduledNavigationEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnFrameDetached(
	socket *sock.Socket,
	callback func(event *page.FrameDetachedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameDetached",
		func(name string, params []byte) {
			event := &page.FrameDetachedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnFrameNavigated(
	socket *sock.Socket,
	callback func(event *page.FrameNavigatedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameNavigated",
		func(name string, params []byte) {
			event := &page.FrameNavigatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
is resized. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameResized
*/
func (Page) OnFrameResized(
	socket *sock.Socket,
	callback func(event *page.FrameResizedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameResized",
		func(name string, params []byte) {
			event := &page.FrameResizedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
Page.frameScheduledNavigation fires when frame schedules a potential navigation. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameScheduledNavigation
*/
func (Page) OnFrameScheduledNavigation(
	socket *sock.Socket,
	callback func(event *page.FrameScheduledNavigationEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameScheduledNavigation",
		func(name string, params []byte) {
			event := &page.FrameScheduledNavigationEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
fires when frame has started loading. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStartedLoading
*/
func (Page) OnFrameStartedLoading(
	socket *sock.Socket,
	callback func(event *page.FrameStartedLoadingEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameStartedLoading",
		func(name string, params []byte) {
			event := &page.FrameStartedLoadingEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
fires when frame has stopped loading. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStoppedLoading
*/
func (Page) OnFrameStoppedLoading(
	socket *sock.Socket,
	callback func(event *page.FrameStoppedLoadingEvent),
) {
	handler := sock.NewEventHandler(
		"Page.frameStoppedLoading",
		func(name string, params []byte) {
			event := &page.FrameStoppedLoadingEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnInterstitialHidden(
	socket *sock.Socket,
	callback func(event *page.InterstitialHiddenEvent),
) {
	handler := sock.NewEventHandler(
		"Page.interstitialHidden",
		func(name string, params []byte) {
			event := &page.InterstitialHiddenEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnInterstitialShown(
	socket *sock.Socket,
	callback func(event *page.InterstitialShownEvent),
) {
	handler := sock.NewEventHandler(
		"Page.interstitialShown",
		func(name string, params []byte) {
			event := &page.InterstitialShownEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnJavascriptDialogClosed(
	socket *sock.Socket,
	callback func(event *page.JavascriptDialogClosedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.javascriptDialogClosed",
		func(name string, params []byte) {
			event := &page.JavascriptDialogClosedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnJavascriptDialogOpening(
	socket *sock.Socket,
	callback func(event *page.JavascriptDialogOpeningEvent),
) {
	handler := sock.NewEventHandler(
		"Page.javascriptDialogOpening",
		func(name string, params []byte) {
			event := &page.JavascriptDialogOpeningEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnLifecycleEvent(
	socket *sock.Socket,
	callback func(event *page.LifecycleEventEvent),
) {
	handler := sock.NewEventHandler(
		"Page.lifecycleEvent",
		func(name string, params []byte) {
			event := &page.LifecycleEventEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnLoadEventFired(
	socket *sock.Socket,
	callback func(event *page.LoadEventFiredEvent),
) {
	handler := sock.NewEventHandler(
		"Page.loadEventFired",
		func(name string, params []byte) {
			event := &page.LoadEventFiredEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
compressed image data is requested by the `startScreencast` method. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastFrame
*/
func (Page) OnScreencastFrame(
	socket *sock.Socket,
	callback func(event *page.ScreencastFrameEvent),
) {
	handler := sock.NewEventHandler(
		"Page.screencastFrame",
		func(name string, params []byte) {
			event := &page.ScreencastFrameEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
hidden. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastVisibilityChanged
*/
func (Page) OnScreencastVisibilityChanged(
	socket *sock.Socket,
	callback func(event *page.ScreencastVisibilityChangedEvent),
) {
	handler := sock.NewEventHandler(
		"Page.screencastVisibilityChanged",
		func(name string, params []byte) {
			event := &page.ScreencastVisibilityChangedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Page) OnWindowOpen(
	socket *sock.Socket,
	callback func(event *page.WindowOpenEvent),
) {
	handler := sock.NewEventHandler(
		"Page.windowOpen",
		func(name string, params []byte) {
			event := &page.WindowOpenEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
