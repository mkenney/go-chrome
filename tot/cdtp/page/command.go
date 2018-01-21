package page

import (
	"github.com/mkenney/go-chrome/tot/cdtp/debugger"
	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

/*
AddScriptToEvaluateOnLoadParams represents Page.addScriptToEvaluateOnLoad parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
*/
type AddScriptToEvaluateOnLoadParams struct {
	ScriptSource string `json:"scriptSource"`
}

/*
AddScriptToEvaluateOnLoadResult represents the result of calls to Page.addScriptToEvaluateOnLoad.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
*/
type AddScriptToEvaluateOnLoadResult struct {
	// Identifier of the added script.
	Identifier ScriptIdentifier `json:"identifier"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
AddScriptToEvaluateOnNewDocumentParams represents Page.addScriptToEvaluateOnNewDocument
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
type AddScriptToEvaluateOnNewDocumentParams struct {
	Source string `json:"source"`
}

/*
AddScriptToEvaluateOnNewDocumentResult represents the result of calls to
Page.addScriptToEvaluateOnNewDocument.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
type AddScriptToEvaluateOnNewDocumentResult struct {
	// Identifier of the added script.
	Identifier ScriptIdentifier `json:"identifier"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
BringToFrontResult represents the result of calls to Page.bringToFront.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
*/
type BringToFrontResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CaptureScreenshotParams represents Page.captureScreenshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
type CaptureScreenshotParams struct {
	// Optional. Image compression format (defaults to png). Allowed values:
	//	- Format.Jpeg
	//	- Format.Png
	Format FormatEnum `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`

	// Optional. Capture the screenshot of a given region only.
	Clip *Viewport `json:"clip,omitempty"`

	// Optional. Capture the screenshot from the surface, rather than the view.
	// Defaults to true. EXPERIMENTAL.
	FromSurface bool `json:"fromSurface,omitempty"`
}

/*
CaptureScreenshotResult represents the result of calls to Page.captureScreenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
type CaptureScreenshotResult struct {
	// Base64-encoded image data.
	Data string `json:"data"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CreateIsolatedWorldParams represents Page.createIsolatedWorld parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
type CreateIsolatedWorldParams struct {
	// ID of the frame in which the isolated world should be created.
	FrameID FrameID `json:"frameId"`

	// Optional. An optional name which is reported in the Execution Context.
	WorldName string `json:"worldName,omitempty"`

	// Optional. Whether or not universal access should be granted to the
	// isolated world. This is a powerful option, use with caution.
	GrantUniveralAccess bool `json:"grantUniveralAccess,omitempty"`
}

/*
CreateIsolatedWorldResult represents the result of calls to Page.createIsolatedWorld.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
type CreateIsolatedWorldResult struct {
	// Execution context of the isolated world.
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to Page.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Page.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetAppManifestParams represents Page.getAppManifest parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
*/
type GetAppManifestParams struct {
	// Manifest location.
	URL string `json:"url"`

	// Errors.
	Errors []*AppManifestError `json:"errors"`

	// Optional. Manifest content.
	Data string `json:"data,omitempty"`
}

/*
GetAppManifestResult represents the result of calls to Page.getAppManifest.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
*/
type GetAppManifestResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetFrameTreeResult represents the result of calls to Page.getFrameTree.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
*/
type GetFrameTreeResult struct {
	// Present frame tree structure.
	FrameTree *FrameTree `json:"frameTree"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetLayoutMetricsResult represents the result of calls to Page.getLayoutMetrics.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
*/
type GetLayoutMetricsResult struct {
	// Metrics relating to the layout viewport.
	LayoutViewport *LayoutViewport `json:"layoutViewport"`

	// Metrics relating to the visual viewport.
	VisualViewport *VisualViewport `json:"visualViewport"`

	// Size of scrollable area. Rect is a local implementation of DOM.Rect
	ContentSize *Rect `json:"contentSize"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetNavigationHistoryResult represents the result of calls to Page.getNavigationHistory.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
*/
type GetNavigationHistoryResult struct {
	// Index of the current navigation history entry.
	CurrentIndex int `json:"currentIndex"`

	// Array of navigation history entries.
	Entries []*NavigationEntry `json:"entries"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetResourceContentParams represents Page.getResourceContent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
*/
type GetResourceContentParams struct {
	// Frame ID to get resource for.
	FrameID FrameID `json:"frameId"`

	// URL of the resource to get content for.
	URL string `json:"url"`
}

/*
GetResourceContentResult represents the result of calls to Page.getResourceContent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
*/
type GetResourceContentResult struct {
	// Resource content.
	Content string `json:"content"`

	// True, if content was served as base64.
	Base64Encoded bool `json:"base64Encoded"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetResourceTreeResult represents the result of calls to Page.getResourceTree.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
*/
type GetResourceTreeResult struct {
	// Present frame / resource tree structure.
	FrameTree *FrameResourceTree `json:"frameTree"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
HandleJavaScriptDialogParams represents Page.handleJavaScriptDialog parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
type HandleJavaScriptDialogParams struct {
	// Whether to accept or dismiss the dialog.
	Accept bool `json:"accept"`

	// Optional. The text to enter into the dialog prompt before accepting. Used
	// only if this is a prompt dialog.
	PromptText string `json:"promptText,omitempty"`
}

/*
HandleJavaScriptDialogResult represents the result of calls to Page.handleJavaScriptDialog.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
type HandleJavaScriptDialogResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
NavigateParams represents Page.navigate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
type NavigateParams struct {
	// URL to navigate the page to.
	URL string `json:"url"`

	// Optional. Referrer URL.
	Referrer string `json:"referrer,omitempty"`

	// Optional. Intended transition type.
	TransitionType TransitionType `json:"transitionType,omitempty"`
}

/*
NavigateResult represents the result of calls to Page.navigate.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
type NavigateResult struct {
	// Frame ID that has navigated (or failed to navigate).
	FrameID FrameID `json:"frameId"`

	// Loader identifier.
	LoaderID LoaderID `json:"loaderId"`

	// User friendly error message, present if and only if navigation has failed.
	ErrorText string `json:"errorText"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
NavigateToHistoryEntryParams represents Page.navigateToHistoryEntry parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
type NavigateToHistoryEntryParams struct {
	// Unique ID of the entry to navigate to.
	EntryID int `json:"entryId"`
}

/*
NavigateToHistoryEntryResult represents the result of calls to Page.navigateToHistoryEntry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
type NavigateToHistoryEntryResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
PrintToPDFParams represents Page.printToPDF parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
type PrintToPDFParams struct {
	// Optional. Paper orientation. Defaults to false.
	Landscape bool `json:"landscape,omitempty"`

	// Optional. Display header and footer. Defaults to false.
	DisplayHeaderFooter bool `json:"displayHeaderFooter,omitempty"`

	// Optional. Print background graphics. Defaults to false.
	PrintBackground bool `json:"printBackground,omitempty"`

	// Optional. Scale of the webpage rendering. Defaults to 1.
	Scale float64 `json:"scale,omitempty"`

	// Optional. Paper width in inches. Defaults to 8.5 inches.
	PaperWidth float64 `json:"paperWidth,omitempty"`

	// Optional. Paper height in inches. Defaults to 11 inches.
	PaperHeight float64 `json:"paperHeight,omitempty"`

	// Optional. Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop float64 `json:"marginTop,omitempty"`

	// Optional. Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom float64 `json:"marginBottom,omitempty"`

	// Optional. Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft float64 `json:"marginLeft,omitempty"`

	// Optional. Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight float64 `json:"marginRight,omitempty"`

	// Optional. Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the
	// empty string, which means print all pages.
	PageRanges string `json:"pageRanges,omitempty"`

	// Optional. Whether to silently ignore invalid but successfully parsed page
	// ranges, such as '3-2'. Defaults to false.
	IgnoreInvalidPageRanges bool `json:"ignoreInvalidPageRanges,omitempty"`
}

/*
PrintToPDFResult represents the result of calls to Page.printToPDF.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
type PrintToPDFResult struct {
	// Base64-encoded pdf data.
	Data string `json:"data"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ReloadParams represents Page.reload parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
type ReloadParams struct {
	// Optional. If true, browser cache is ignored (as if the user pressed
	// Shift+refresh).
	IgnoreCache bool `json:"ignoreCache,omitempty"`

	// Optional. If set, the script will be injected into all frames of the
	// inspected page after reload.
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
}

/*
ReloadResult represents the result of calls to Page.reload.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
type ReloadResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RemoveScriptToEvaluateOnLoadParams represents Page.removeScriptToEvaluateOnLoad parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
*/
type RemoveScriptToEvaluateOnLoadParams struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

/*
RemoveScriptToEvaluateOnLoadResult represents the result of calls to
Page.removeScriptToEvaluateOnLoad.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
*/
type RemoveScriptToEvaluateOnLoadResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RemoveScriptToEvaluateOnNewDocumentParams represents Page.removeScriptToEvaluateOnNewDocument
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
*/
type RemoveScriptToEvaluateOnNewDocumentParams struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

/*
RemoveScriptToEvaluateOnNewDocumentResult represents the result of calls to
Page.removeScriptToEvaluateOnNewDocument.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
*/
type RemoveScriptToEvaluateOnNewDocumentResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RequestAppBannerResult represents the result of calls to Page.requestAppBanner.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-requestAppBanner
*/
type RequestAppBannerResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ScreencastFrameAckParams represents Page.screencastFrameAck parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck
*/
type ScreencastFrameAckParams struct {
	// Frame number.
	SessionID int `json:"sessionId"`
}

/*
ScreencastFrameAckResult represents the result of calls to Page.screencastFrameAck.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck
*/
type ScreencastFrameAckResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SearchInResourceParams represents Page.searchInResource parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
*/
type SearchInResourceParams struct {
	// Frame ID for resource to search in.
	FrameID FrameID `json:"frameId"`

	// URL of the resource to search in.
	URL string `json:"url"`

	// String to search for.
	Query string `json:"query"`

	// Optional. If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// Optional. If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

/*
SearchInResourceResult represents the result of calls to Page.searchInResource.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
*/
type SearchInResourceResult struct {
	// List of search matches.
	Result []*debugger.SearchMatch `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetAdBlockingEnabledParams represents Page.setAdBlockingEnabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
*/
type SetAdBlockingEnabledParams struct {
	// Whether to block ads.
	Enabled bool `json:"enabled"`
}

/*
SetAdBlockingEnabledResult represents the result of calls to Page.setAdBlockingEnabled.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
*/
type SetAdBlockingEnabledResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetAutoAttachToCreatedPagesParams represents Page.setAutoAttachToCreatedPages parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAutoAttachToCreatedPages
*/
type SetAutoAttachToCreatedPagesParams struct {
	// If true, browser will open a new inspector window for every page created
	// from this one.
	AutoAttach bool `json:"autoAttach"`
}

/*
SetAutoAttachToCreatedPagesResult represents the result of calls to
Page.setAutoAttachToCreatedPages.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAutoAttachToCreatedPages
*/
type SetAutoAttachToCreatedPagesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetDocumentContentParams represents Page.setDocumentContent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
*/
type SetDocumentContentParams struct {
	// Frame ID to set HTML for.
	FrameID FrameID `json:"frameId"`

	// HTML content to set.
	HTML string `json:"html"`
}

/*
SetDocumentContentResult represents the result of calls to Page.setDocumentContent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
*/
type SetDocumentContentResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetDownloadBehaviorParams represents Page.setDownloadBehavior parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
*/
type SetDownloadBehaviorParams struct {
	// Whether to allow all or deny all download requests, or use default Chrome
	// behavior if available (otherwise deny). Allowed values:
	//	- Behavior.Deny
	//	- Behavior.Allow
	//	- Behavior.Default
	Behavior BehaviorEnum `json:"behavior"`

	// Optional. The default path to save downloaded files to. This is required
	// if behavior is set to 'allow'.
	DownloadPath string `json:"downloadPath,omitempty"`
}

/*
SetDownloadBehaviorResult represents the result of calls to Page.setDownloadBehavior.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
*/
type SetDownloadBehaviorResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetLifecycleEventsEnabledParams represents Page.setLifecycleEventsEnabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
*/
type SetLifecycleEventsEnabledParams struct {
	// If true, starts emitting lifecycle events.
	Enabled bool `json:"enabled"`
}

/*
SetLifecycleEventsEnabledResult represents the result of calls to Page.setLifecycleEventsEnabled.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
*/
type SetLifecycleEventsEnabledResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StartScreencastParams represents Page.startScreencast parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
*/
type StartScreencastParams struct {
	// Optional. Image compression format. Allowed values:
	//	- Format.Jpeg
	//	- Format.Png
	Format FormatEnum `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100].
	Quality int `json:"quality,omitempty"`

	// Optional. Maximum screenshot width.
	MaxWidth int `json:"maxWidth,omitempty"`

	// Optional. Maximum screenshot height.
	MaxHeight int `json:"maxHeight,omitempty"`

	// Optional. Send every n-th frame.
	EveryNthFrame int `json:"everyNthFrame,omitempty"`
}

/*
StartScreencastResult represents the result of calls to Page.startScreencast.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
*/
type StartScreencastResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopLoadingResult represents the result of calls to Page.stopLoading.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopLoading
*/
type StopLoadingResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopScreencastResult represents the result of calls to Page.stopScreencast.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopScreencast
*/
type StopScreencastResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
