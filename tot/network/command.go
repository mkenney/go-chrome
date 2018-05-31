package network

import (
	"github.com/mkenney/go-chrome/tot/debugger"
)

/*
CanClearBrowserCacheResult represents the result of calls to Network.canClearBrowserCache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCache
*/
type CanClearBrowserCacheResult struct {
	// True if browser cache can be cleared.
	Result bool `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CanClearBrowserCookiesResult represents the result of calls to Network.canClearBrowserCookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCookies
*/
type CanClearBrowserCookiesResult struct {
	// True if browser cookies can be cleared.
	Result bool `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CanEmulateConditionsResult represents the result of calls to Network.canEmulateNetworkConditions.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canEmulateNetworkConditions
*/
type CanEmulateConditionsResult struct {
	// True if emulation of network conditions is supported.
	Result bool `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ClearBrowserCacheResult represents the result of calls to Network.clearBrowserCache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache
*/
type ClearBrowserCacheResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ClearBrowserCookiesResult represents the result of calls to Network.clearBrowserCookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCookies
*/
type ClearBrowserCookiesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ContinueInterceptedRequestParams represents Network.continueInterceptedRequest parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-continueInterceptedRequest
*/
type ContinueInterceptedRequestParams struct {
	// The interception ID.
	InterceptionID InterceptionID `json:"interceptionId"`

	// Optional. If set this causes the request to fail with the given reason.
	// Passing `Aborted` for requests marked with `isNavigationRequest` also
	// cancels the navigation. Must not be set in response to an AuthChallenge.
	// Allowed values:
	//	- ErrorReason.Failed
	//	- ErrorReason.Aborted
	//	- ErrorReason.TimedOut
	//	- ErrorReason.AccessDenied
	//	- ErrorReason.ConnectionClosed
	//	- ErrorReason.ConnectionReset
	//	- ErrorReason.ConnectionRefused
	//	- ErrorReason.ConnectionAborted
	//	- ErrorReason.ConnectionFailed
	//	- ErrorReason.NameNotResolved
	//	- ErrorReason.InternetDisconnected
	//	- ErrorReason.AddressUnreachable
	ErrorReason ErrorReasonEnum `json:"errorReason,omitempty"`

	// Optional. If set the requests completes using with the provided base64
	// encoded raw response, including HTTP status line and headers etc... Must
	// not be set in response to an AuthChallenge.
	RawResponse string `json:"rawResponse,omitempty"`

	// Optional. f set the request url will be modified in a way that's not
	// observable by page. Must not be set in response to an AuthChallenge.
	URL string `json:"url,omitempty"`

	// Optional. If set this allows the request method to be overridden. Must
	// not be set in response to an AuthChallenge.
	Method string `json:"method,omitempty"`

	// Optional. If set this allows postData to be set. Must not be set in
	// response to an AuthChallenge.
	PostData string `json:"postData,omitempty"`

	// Optional. If set this allows the request headers to be changed. Must not
	// be set in response to an AuthChallenge.
	Headers Headers `json:"headers,omitempty"`

	// Optional. Response to a requestIntercepted with an AuthChallenge. Must
	// not be set otherwise.
	AuthChallengeResponse *AuthChallengeResponse `json:"authChallengeResponse,omitempty"`
}

/*
ContinueInterceptedRequestResult represents the result of calls to
Network.continueInterceptedRequest.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-continueInterceptedRequest
*/
type ContinueInterceptedRequestResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DeleteCookiesParams represents Network.deleteCookies parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
*/
type DeleteCookiesParams struct {
	// Name of the cookies to remove.
	Name string `json:"name"`

	// Optional. If specified, deletes all the cookies with the given name where
	// domain and path match provided URL.
	URL string `json:"url,omitempty"`

	// Optional. If specified, deletes only cookies with the exact domain.
	Domain string `json:"domain,omitempty"`

	// Optional. If specified, deletes only cookies with the exact path.
	Path string `json:"path,omitempty"`
}

/*
DeleteCookiesResult represents the result of calls to Network.deleteCookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
*/
type DeleteCookiesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to Network.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EmulateConditionsParams represents Network.emulateNetworkConditions parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions
*/
type EmulateConditionsParams struct {
	// True to emulate internet disconnection.
	Offline bool `json:"offline"`

	// Minimum latency from request sent to response headers received (ms).
	Latency float64 `json:"latency"`

	// Maximal aggregated download throughput (bytes/sec). -1 disables download
	// throttling.
	DownloadThroughput float64 `json:"downloadThroughput"`

	// Maximal aggregated upload throughput (bytes/sec). -1 disables upload
	// throttling.
	UploadThroughput float64 `json:"uploadThroughput"`

	// Optional. Connection type if known. Allowed values:
	//	- ConnectionType.None
	//	- ConnectionType.Cellular2g
	//	- ConnectionType.Cellular3g
	//	- ConnectionType.Cellular4g
	//	- ConnectionType.Bluetooth
	//	- ConnectionType.Ethernet
	//	- ConnectionType.Wifi
	//	- ConnectionType.Wimax
	//	- ConnectionType.Other
	ConnectionType ConnectionTypeEnum `json:"connectionType,omitempty"`
}

/*
EmulateConditionsResult represents the result of calls to Network.emulateConditions.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateConditions
*/
type EmulateConditionsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableParams represents Network.enable parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
*/
type EnableParams struct {
	// Optional. Buffer size in bytes to use when preserving network payloads
	// (XHRs, etc). EXPERIMENTAL.
	MaxTotalBufferSize int `json:"maxTotalBufferSize,omitempty"`

	// Optional. Per-resource buffer size in bytes to use when preserving
	// network payloads (XHRs, etc). EXPERIMENTAL.
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
}

/*
EnableResult represents the result of calls to Network.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetAllCookiesResult represents the result of calls to Network.getAllCookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getAllCookies
*/
type GetAllCookiesResult struct {
	// Array of cookie objects.
	Cookies []*Cookie `json:"cookies"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetCertificateParams represents Network.getCertificate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate
*/
type GetCertificateParams struct {
	// Origin to get certificate for.
	Origin string `json:"origin"`
}

/*
GetCertificateResult represents the result of calls to Network.getCertificate.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate
*/
type GetCertificateResult struct {
	// The list of table names
	TableNames []string `json:"tableNames"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetCookiesParams represents Network.getCookies parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
*/
type GetCookiesParams struct {
	// Optional. The list of URLs for which applicable cookies will be fetched.
	URLs []string `json:"urls,omitempty"`
}

/*
GetCookiesResult represents the result of calls to Network.getCookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
*/
type GetCookiesResult struct {
	// Array of cookie objects.
	Cookies []*Cookie `json:"cookies"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetResponseBodyParams represents Network.getResponseBody parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
*/
type GetResponseBodyParams struct {
	// Identifier of the network request to get content for.
	RequestID RequestID `json:"requestId"`
}

/*
GetResponseBodyResult represents the result of calls to Network.getResponseBody.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
*/
type GetResponseBodyResult struct {
	// Response body.
	Body string `json:"body"`

	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetResponseBodyForInterceptionParams represents Network.getResponseBodyForInterception parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
*/
type GetResponseBodyForInterceptionParams struct {
	// Identifier for the intercepted request to get body for.
	InterceptionID InterceptionID `json:"interceptionId"`
}

/*
GetResponseBodyForInterceptionResult represents the result of calls to
Network.getResponseBodyForInterception.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
*/
type GetResponseBodyForInterceptionResult struct {
	// Response body.
	Body string `json:"body"`

	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ReplayXHRParams represents Network.replayXHR parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR
*/
type ReplayXHRParams struct {
	// Identifier of XHR to replay.
	RequestID RequestID `json:"requestId"`
}

/*
ReplayXHRResult represents the result of calls to Network.replayXHR.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR
*/
type ReplayXHRResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SearchInResponseBodyParams represents Network.searchInResponseBody parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
*/
type SearchInResponseBodyParams struct {
	// Identifier of the network response to search.
	RequestID RequestID `json:"requestId"`

	// String to search for.
	Query string `json:"query"`

	// Optional. If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// Optional. If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

/*
SearchInResponseBodyResult represents the result of calls to Network.searchInResponseBody.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
*/
type SearchInResponseBodyResult struct {
	// List of search matches.
	Result []*debugger.SearchMatch `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetBlockedURLsParams represents Network.setBlockedURLs parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs
*/
type SetBlockedURLsParams struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	URLs []string `json:"urls"`
}

/*
SetBlockedURLsResult represents the result of calls to Network.setBlockedURLs.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs
*/
type SetBlockedURLsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetBypassServiceWorkerParams represents Network.setBypassServiceWorker parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
*/
type SetBypassServiceWorkerParams struct {
	// Bypass service worker and load from network.
	Bypass bool `json:"bypass"`
}

/*
SetBypassServiceWorkerResult represents the result of calls to Network.setBypassServiceWorker.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
*/
type SetBypassServiceWorkerResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetCacheDisabledParams represents Network.setCacheDisabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
*/
type SetCacheDisabledParams struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

/*
SetCacheDisabledResult represents the result of calls to Network.setCacheDisabled.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
*/
type SetCacheDisabledResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetCookieParams represents Network.setCookie parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
*/
type SetCookieParams struct {
	// Cookie name.
	Name string `json:"name"`

	// Cookie value.
	Value string `json:"value"`

	// Optional. The request-URI to associate with the setting of the cookie.
	// This value can affect the default domain and path values of the created
	// cookie.
	URL string `json:"url,omitempty"`

	// Optional. Cookie domain.
	Domain string `json:"domain,omitempty"`

	// Optional. Cookie path.
	Path string `json:"path,omitempty"`

	// Optional. True if cookie is secure.
	Secure bool `json:"secure,omitempty"`

	// Optional. True if cookie is http-only.
	HTTPOnly bool `json:"httpOnly,omitempty"`

	// Optional. Cookie SameSite type. Allowed values:
	//	- CookieSameSite.Strict
	//	- CookieSameSite.Lax
	SameSite CookieSameSiteEnum `json:"sameSite,omitempty"`

	// Optional. Cookie expiration date, session cookie if not set.
	Expires TimeSinceEpoch `json:"expires,omitempty"`
}

/*
SetCookieResult represents the result of calls to Network.setCookie.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
*/
type SetCookieResult struct {
	// True if successfully set cookie.
	Success bool `json:"success"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetCookiesParams represents Network.setCookies parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
*/
type SetCookiesParams struct {
	// Cookies to be set.
	Cookies []*SetCookieParams `json:"cookies"`
}

/*
SetCookiesResult represents the result of calls to Network.setCookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
*/
type SetCookiesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetDataSizeLimitsForTestParams represents Network.setDataSizeLimitsForTest parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
*/
type SetDataSizeLimitsForTestParams struct {
	// Maximum total buffer size.
	MaxTotalSize int `json:"maxTotalSize"`

	// Maximum per-resource size.
	MaxResourceSize int `json:"maxResourceSize"`
}

/*
SetDataSizeLimitsForTestResult represents the result of calls to
Network.setDataSizeLimitsForTest.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
*/
type SetDataSizeLimitsForTestResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetExtraHTTPHeadersParams represents Network.setExtraHTTPHeaders parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
*/
type SetExtraHTTPHeadersParams struct {
	// Map with extra HTTP headers.
	Headers Headers `json:"headers"`
}

/*
SetExtraHTTPHeadersResult represents the result of calls to Network.setExtraHTTPHeaders.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
*/
type SetExtraHTTPHeadersResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetRequestInterceptionParams represents Network.setRequestInterception parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
*/
type SetRequestInterceptionParams struct {
	// Requests matching any of these patterns will be forwarded and wait for
	// the corresponding continueInterceptedRequest call.
	Patterns []*RequestPattern `json:"patterns"`
}

/*
SetRequestInterceptionResult represents the result of calls to Network.setRequestInterception.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
*/
type SetRequestInterceptionResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetUserAgentOverrideParams represents Network.setUserAgentOverride parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setUserAgentOverride
*/
type SetUserAgentOverrideParams struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
}

/*
SetUserAgentOverrideResult represents the result of calls to Network.setUserAgentOverride.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setUserAgentOverride
*/
type SetUserAgentOverrideResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
