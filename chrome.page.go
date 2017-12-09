package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Page - https://chromedevtools.github.io/devtools-protocol/tot/Page/
Actions and events related to the inspected page belong to the page domain.
*/
type Page struct{}

/*
AddScriptToEvaluateOnLoad is eprecated, please use addScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL DEPRECATED
*/
func (Page) AddScriptToEvaluateOnLoad(
	socket *Socket,
	params *Page.AddScriptToEvaluateOnLoadParams,
) (Page.AddScriptToEvaluateOnLoadResult, error) {
	command := &protocol.Command{
		method: "Page.addScriptToEvaluateOnLoad",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.AddScriptToEvaluateOnLoadResult), command.Err
}

/*
AddScriptToEvaluateOnNewDocument evaluates given script in every frame upon creation (before loading
frame's scripts).
*/
func (Page) AddScriptToEvaluateOnNewDocument(
	socket *Socket,
	params *Page.AddScriptToEvaluateOnNewDocumentParams,
) (Page.AddScriptToEvaluateOnNewDocumentResult, error) {
	command := &protocol.Command{
		method: "Page.addScriptToEvaluateOnNewDocument",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.AddScriptToEvaluateOnNewDocumentResult), command.Err
}

/*
BringToFront brings page to front (activates tab).
*/
func (Page) BringToFront(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.bringToFront",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
CaptureScreenshot capture a page screenshot.
*/
func (Page) CaptureScreenshot(
	socket *Socket,
	params *Page.CaptureScreenshotParams,
) (Page.CaptureScreenshotResult, error) {
	command := &protocol.Command{
		method: "Page.CaptureScreenshot",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.CaptureScreenshotResult), command.Err
}

/*
CreateIsolatedWorld creates an isolated world for the given frame.
*/
func (Page) CreateIsolatedWorld(
	socket *Socket,
	params *Page.CreateIsolatedWorldParams,
) (Page.CreateIsolatedWorldResult, error) {
	command := &protocol.Command{
		method: "Page.createIsolatedWorld",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.CreateIsolatedWorldResult), command.Err
}

/*
Disable disables page domain notifications.
*/
func (Page) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable Ennables page domain notifications.
*/
func (Page) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetAppManifest gets the app manifest.
*/
func (Page) GetAppManifest(
	socket *Socket,
	params *Page.GetAppManifestParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.getAppManifest",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetFrameTree returns present frame tree structure.
*/
func (Page) GetFrameTree(
	socket *Socket,
) (Page.GetFrameTreeResult, error) {
	command := &protocol.Command{
		method: "Page.getFrameTree",
	}
	socket.SendCommand(command)
	return command.Result.(Page.GetFrameTreeResult), command.Err
}

/*
GetLayoutMetrics returns metrics relating to the layouting of the page, such as viewport
bounds/scale.
*/
func (Page) GetLayoutMetrics(
	socket *Socket,
) (Page.GetLayoutMetricsResult, error) {
	command := &protocol.Command{
		method: "Page.getLayoutMetrics",
	}
	socket.SendCommand(command)
	return command.Result.(Page.GetLayoutMetricsResult), command.Err
}

/*
GetNavigationHistory returns navigation history for the current page.
*/
func (Page) GetNavigationHistory(
	socket *Socket,
) (Page.GetNavigationHistoryResult, error) {
	command := &protocol.Command{
		method: "Page.getNavigationHistory",
	}
	socket.SendCommand(command)
	return command.Result.(Page.GetNavigationHistoryResult), command.Err
}

/*
GetResourceContent returns content of the given resource. EXPERIMENTAL
*/
func (Page) GetResourceContent(
	socket *Socket,
	params *Page.GetResourceContentParams,
) (Page.GetResourceContentResult, error) {
	command := &protocol.Command{
		method: "Page.getResourceContent",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.GetResourceContentResult), command.Err
}

/*
GetResourceTree returns present frame / resource tree structure. EXPERIMENTAL
*/
func (Page) GetResourceTree(
	socket *Socket,
) (Page.GetResourceTreeResult, error) {
	command := &protocol.Command{
		method: "Page.getResourceTree",
	}
	socket.SendCommand(command)
	return command.Result.(Page.GetResourceTreeResult), command.Err
}

/*
HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt,
or onbeforeunload).
*/
func (Page) HandleJavaScriptDialog(
	socket *Socket,
	params *Page.HandleJavaScriptDialogParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.handleJavaScriptDialog",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Navigate navigates current page to the given URL.
*/
func (Page) Navigate(
	socket *Socket,
	params *Page.NavigateParams,
) (Page.NavigateResult, error) {
	command := &protocol.Command{
		method: "Page.navigate",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.NavigateResult), command.Err
}

/*
NavigateToHistoryEntry navigates current page to the given history entry.
*/
func (Page) NavigateToHistoryEntry(
	socket *Socket,
	params *Page.NavigateToHistoryEntryParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.navigateToHistoryEntry",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
PrintToPDF print page as PDF.
*/
func (Page) PrintToPDF(
	socket *Socket,
	params *Page.PrintToPDFParams,
) (Page.PrintToPDFResult, error) {
	command := &protocol.Command{
		method: "Page.printToPDF",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.PrintToPDFResult), command.Err
}

/*
Reload reloads given page optionally ignoring the cache.
*/
func (Page) Reload(
	socket *Socket,
	params *Page.ReloadParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.reload",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveScriptToEvaluateOnLoad deprecated, please use removeScriptToEvaluateOnNewDocument instead.
EXPERIMENTAL DEPRECATED
*/
func (Page) RemoveScriptToEvaluateOnLoad(
	socket *Socket,
	params *Page.RemoveScriptToEvaluateOnLoadParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.removeScriptToEvaluateOnLoad",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveScriptToEvaluateOnNewDocument removes given script from the list.
*/
func (Page) RemoveScriptToEvaluateOnNewDocument(
	socket *Socket,
	params *Page.RemoveScriptToEvaluateOnNewDocumentParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.removeScriptToEvaluateOnNewDocument",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RequestAppBanner EXPERIMENTAL
*/
func (Page) RequestAppBanner(
	socket *Socket,
	params *Page.RequestAppBannerParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.requestAppBanner",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
ScreencastFrameAck acknowledges that a screencast frame has been received by the frontend.
EXPERIMENTAL
*/
func (Page) ScreencastFrameAck(
	socket *Socket,
	params *Page.ScreencastFrameAckParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.screencastFrameAck",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SearchInResource searches for given string in resource content. EXPERIMENTAL
*/
func (Page) SearchInResource(
	socket *Socket,
	params *Page.SearchInResourceParams,
) (Page.SearchInResourceResult, error) {
	command := &protocol.Command{
		method: "Page.searchInResource",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(Page.SearchInResourceResult), command.Err
}

/*
SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites. EXPERIMENTAL
*/
func (Page) SetAdBlockingEnabled(
	socket *Socket,
	params *Page.SetAdBlockingEnabledParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.setAdBlockingEnabled",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetAutoAttachToCreatedPages controls whether browser will open a new inspector window for connected
pages. EXPERIMENTAL
*/
func (Page) SetAutoAttachToCreatedPages(
	socket *Socket,
	params *Page.SetAutoAttachToCreatedPagesParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.setAutoAttachToCreatedPages",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetDocumentContent sets given markup as the document's HTML.
*/
func (Page) SetDocumentContent(
	socket *Socket,
	params *Page.SetDocumentContentParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.setDocumentContent",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetDownloadBehavior sets the behavior when downloading a file. EXPERIMENTAL
*/
func (Page) SetDownloadBehavior(
	socket *Socket,
	params *Page.SetDownloadBehaviorParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.setDownloadBehavior",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetLifecycleEventsEnabled controls whether page will emit lifecycle events. EXPERIMENTAL
*/
func (Page) SetLifecycleEventsEnabled(
	socket *Socket,
	params *Page.SetLifecycleEventsEnabledParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.setLifecycleEventsEnabled",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StartScreencast starts sending each frame using the `screencastFrame` event. EXPERIMENTAL
*/
func (Page) StartScreencast(
	socket *Socket,
	params *Page.StartScreencastParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.startScreencast",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StopLoading force the page stop all navigations and pending resource fetches.
*/
func (Page) StopLoading(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.stopLoading",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StopScreencast stops sending each frame in the `screencastFrame`. EXPERIMENTAL
*/
func (Page) StopScreencast(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Page.stopScreencast",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
OnDOMContentEventFired adds a handler to the Page.domContentEventFired event.
Page.domContentEventFired fires when a content event occurs in the DOM.
*/
func (Page) OnDOMContentEventFired(
	socket *Socket,
	callback func(event *page.DOMContentEventFiredEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnFrameAttached(
	socket *Socket,
	callback func(event *page.FrameAttachedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnFrameClearedScheduledNavigation(
	socket *Socket,
	callback func(event *page.FrameClearedScheduledNavigationEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnFrameDetached(
	socket *Socket,
	callback func(event *page.FrameDetachedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnFrameNavigated(
	socket *Socket,
	callback func(event *page.FrameNavigatedEvent),
) {
	handler := protocol.NewEventHandler(
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
schedules a potential navigation. EXPERIMENTAL
*/
func (Page) OnFrameResized(
	socket *Socket,
	callback func(event *page.FrameResizedEvent),
) {
	handler := protocol.NewEventHandler(
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
OnFrameStartedLoading adds a handler to the Page.frameStartedLoading event. Page.frameStartedLoading
fires when frame has started loading. EXPERIMENTAL
*/
func (Page) OnFrameStartedLoading(
	socket *Socket,
	callback func(event *page.FrameStartedLoadingEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnFrameStoppedLoading(
	socket *Socket,
	callback func(event *page.FrameStoppedLoadingEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnInterstitialHidden(
	socket *Socket,
	callback func(event *page.InterstitialHiddenEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnInterstitialShown(
	socket *Socket,
	callback func(event *page.InterstitialShownEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnJavascriptDialogClosed(
	socket *Socket,
	callback func(event *page.JavascriptDialogClosedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnJavascriptDialogOpening(
	socket *Socket,
	callback func(event *page.JavascriptDialogOpeningEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnLifecycleEvent(
	socket *Socket,
	callback func(event *page.LifecycleEventEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnLoadEventFired(
	socket *Socket,
	callback func(event *page.LoadEventFiredEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnScreencastFrame(
	socket *Socket,
	callback func(event *page.ScreencastFrameEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnScreencastVisibilityChanged(
	socket *Socket,
	callback func(event *page.ScreencastVisibilityChangedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (Page) OnWindowOpen(
	socket *Socket,
	callback func(event *page.WindowOpenEvent),
) {
	handler := protocol.NewEventHandler(
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
