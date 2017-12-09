package Network

import (
	Debugger "app/chrome/debugger"
	Page "app/chrome/page"
	Runtime "app/chrome/runtime"
	Security "app/chrome/security"
	"fmt"
)

/*
CanClearBrowserCacheResult represents the result of calls to Network.canClearBrowserCache.
*/
type CanClearBrowserCacheResult struct {
	// True if browser cache can be cleared.
	Result bool `json:"result"`
}

/*
CanClearBrowserCookiesResult represents the result of calls to Network.canClearBrowserCookies.
*/
type CanClearBrowserCookiesResult struct {
	// True if browser cookies can be cleared.
	Result bool `json:"result"`
}

/*
CanEmulateNetworkConditionsResult represents the result of calls to Network.canEmulateNetworkConditions.
*/
type CanEmulateNetworkConditionsResult struct {
	// True if emulation of network conditions is supported.
	Result bool `json:"result"`
}

/*
ContinueInterceptedRequestParams represents Network.continueInterceptedRequest parameters.
*/
type ContinueInterceptedRequestParams struct {
	// The interception ID.
	InterceptionID InterceptionID `json:"interceptionId"`

	// Optional. If set this causes the request to fail with the given reason. Passing `Aborted` for
	// requests marked with `isNavigationRequest` also cancels the navigation. Must not be set in
	// response to an AuthChallenge.
	ErrorReason ErrorReason `json:"errorReason,omitempty"`

	// Optional. If set the requests completes using with the provided base64 encoded raw response,
	// including HTTP status line and headers etc... Must not be set in response to an
	// AuthChallenge.
	RawResponse string `json:"rawResponse,omitempty"`

	// IOptional. f set the request url will be modified in a way that's not observable by page.
	// Must not be set in response to an AuthChallenge.
	URL string `json:"url,omitempty"`

	// Optional. If set this allows the request method to be overridden. Must not be set in response
	// to an AuthChallenge.
	Method string `json:"method,omitempty"`

	// Optional. If set this allows postData to be set. Must not be set in response to an
	// AuthChallenge.
	PostData string `json:"postData,omitempty"`

	// Optional. If set this allows the request headers to be changed. Must not be set in response
	// to an AuthChallenge.
	Headers Headers `json:"headers,omitempty"`

	// Optional. Response to a requestIntercepted with an AuthChallenge. Must not be set otherwise.
	AuthChallengeResponse AuthChallengeResponse `json:"authChallengeResponse,omitempty"`
}

/*
DeleteCookiesParams represents Network.deleteCookies parameters.
*/
type DeleteCookiesParams struct {
	// Name of the cookies to remove.
	Name string `json:"name"`

	// Optional. If specified, deletes all the cookies with the given name where domain and path
	// match provided URL.
	URL string `json:"url,omitempty"`

	// Optional. If specified, deletes only cookies with the exact domain.
	Domain string `json:"domain,omitempty"`

	// Optional. If specified, deletes only cookies with the exact path.
	Path string `json:"path,omitempty"`
}

/*
EmulateNetworkConditionsParams represents Network.emulateNetworkConditions parameters.
*/
type EmulateNetworkConditionsParams struct {
	// True to emulate internet disconnection.
	Offline bool `json:"offline"`

	// Minimum latency from request sent to response headers received (ms).
	Latency float64 `json:"latency"`

	// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	DownloadThroughput float64 `json:"downloadThroughput"`

	// Maximal aggregated upload throughput (bytes/sec). -1 disables upload throttling.
	UploadThroughput float64 `json:"uploadThroughput"`

	// Optional. Connection type if known.
	ConnectionType ConnectionType `json:"connectionType,omitempty"`
}

/*
EnableParams represents Network.enable parameters.
*/
type EnableParams struct {
	// Optional. Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	// EXPERIMENTAL
	MaxTotalBufferSize int `json:"maxTotalBufferSize,omitempty"`

	// Optional. Per-resource buffer size in bytes to use when preserving network payloads (XHRs,
	// etc). EXPERIMENTAL
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
}

/*
GetAllCookiesResult represents the result of calls to Network.getAllCookies.
*/
type GetAllCookiesResult struct {
	// Array of cookie objects.
	Cookies []Cookie `json:"cookies"`
}

/*
GetCertificateParams represents Network.getCertificate parameters.
*/
type GetCertificateParams struct {
	// Origin to get certificate for.
	Origin string `json:"origin"`
}

/*
GetCertificateResult represents the result of calls to Network.getCertificate.
*/
type GetCertificateResult struct {
	TableNames []string `json:"tableNames"`
}

/*
GetCookiesParams represents Network.getCookies parameters.
*/
type GetCookiesParams struct {
	// Optional. The list of URLs for which applicable cookies will be fetched.
	URLs []string `json:"urls,omitempty"`
}

/*
GetCookiesResult represents the result of calls to Network.getCookies.
*/
type GetCookiesResult struct {
	// Array of cookie objects.
	Cookies []Cookie `json:"cookies"`
}

/*
GetResponseBodyParams represents Network.getResponseBody parameters.
*/
type GetResponseBodyParams struct {
	// Identifier of the network request to get content for.
	RequestID RequestID `json:"requestId"`
}

/*
GetResponseBodyResult represents the result of calls to Network.getResponseBody.
*/
type GetResponseBodyResult struct {
	// Response body.
	Body string `json:"body"`

	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

/*
GetResponseBodyForInterceptionParams represents Network.getResponseBodyForInterception parameters.
*/
type GetResponseBodyForInterceptionParams struct {
	// Identifier for the intercepted request to get body for.
	InterceptionID InterceptionID `json:"interceptionId"`
}

/*
GetResponseBodyForInterceptionResult represents the result of calls to
Network.getResponseBodyForInterception.
*/
type GetResponseBodyForInterceptionResult struct {
	// Response body.
	Body string `json:"body"`

	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

/*
ReplayXHRParams represents Network.replayXHR parameters.
*/
type ReplayXHRParams struct {
	// Identifier of XHR to replay.
	RequestID RequestID `json:"requestId"`
}

/*
SearchInResponseBodyParams represents Network.searchInResponseBody parameters.
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
*/
type SearchInResponseBodyResult struct {
	// List of search matches.
	Result []Debugger.SearchMatch `json:"result"`
}

/*
SetBlockedURLsParams represents Network.setBlockedURLs parameters.
*/
type SetBlockedURLsParams struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	URLs []string `json:"urls"`
}

/*
SetBypassServiceWorkerParams represents Network.setBypassServiceWorker parameters.
*/
type SetBypassServiceWorkerParams struct {
	// Bypass service worker and load from network.
	Bypass bool `json:"bypass"`
}

/*
SetCacheDisabledParams represents Network.setCacheDisabled parameters.
*/
type SetCacheDisabledParams struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

/*
SetCookieParams represents Network.setCookie parameters.
*/
type SetCookieParams struct {
	// Cookie name.
	Name string `json:"name"`

	// Cookie value.
	Value string `json:"value"`

	// Optional. The request-URI to associate with the setting of the cookie. This value can affect
	// the default domain and path values of the created cookie.
	URL string `json:"url,omitempty"`

	// Optional. Cookie domain.
	Domain string `json:"domain,omitempty"`

	// Optional. Cookie path.
	Path string `json:"path,omitempty"`

	// Optional. True if cookie is secure.
	Secure bool `json:"secure,omitempty"`

	// Optional. True if cookie is http-only.
	HTTPOnly bool `json:"httpOnly,omitempty"`

	// Optional. Cookie SameSite type.
	SameSite CookieSameSite `json:"sameSite,omitempty"`

	// Optional. Cookie expiration date, session cookie if not set.
	Expires TimeSinceEpoch `json:"expires,omitempty"`
}

/*
SetCookieResult represents the result of calls to Network.setCookie.
*/
type SetCookieResult struct {
	// True if successfully set cookie.
	Success bool `json:"success"`
}

/*
SetCookiesParams represents Network.setCookies parameters.
*/
type SetCookiesParams struct {
	// Cookies to be set.
	Cookies []CookieParam `json:"cookies"`
}

/*
SetDataSizeLimitsForTestParams represents Network.setDataSizeLimitsForTest parameters.
*/
type SetDataSizeLimitsForTestParams struct {
	// Maximum total buffer size.
	MaxTotalSize int `json:"maxTotalSize"`

	// Maximum per-resource size.
	MaxResourceSize int `json:"maxResourceSize"`
}

/*
SetExtraHTTPHeadersParams represents Network.setExtraHTTPHeaders parameters.
*/
type SetExtraHTTPHeadersParams struct {
	// Map with extra HTTP headers.
	Headers Headers `json:"headers"`
}

/*
SetRequestInterceptionParams represents Network.setRequestInterception parameters.
*/
type SetRequestInterceptionParams struct {
	// Requests matching any of these patterns will be forwarded and wait for the corresponding
	// continueInterceptedRequest call.
	Patterns []RequestPattern `json:"patterns"`
}

/*
SetUserAgentOverrideParams represents Network.setUserAgentOverride parameters.
*/
type SetUserAgentOverrideParams struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
}

/*
DataReceivedEvent represents Network.dataReceived event data.
*/
type DataReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Data chunk length.
	DataLength int `json:"dataLength"`

	// Actual bytes received (might be less than dataLength for compressed encodings).
	EncodedDataLength int `json:"encodedDataLength"`
}

/*
EventSourceMessageReceivedEvent represents Network.eventSourceMessageReceived event data.
*/
type EventSourceMessageReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Message type.
	EventName string `json:"eventName"`

	// Message identifier.
	EventID string `json:"eventId"`

	// Message content.
	Data string `json:"data"`
}

/*
LoadingFailedEvent represents Network.loadingFailed event data.
*/
type LoadingFailedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Resource type.
	Type Page.ResourceType `json:"type"`

	// User friendly error message.
	ErrorText string `json:"errorText"`

	// Optional. True if loading was canceled.
	Canceled bool `json:"canceled,omitempty"`

	// Optional. The reason why loading was blocked, if any.
	BlockedReason BlockedReason `json:"blockedReason,omitempty"`
}

/*
LoadingFinishedEvent represents Network.loadingFinished event data.
*/
type LoadingFinishedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Total number of bytes received for this request.
	EncodedDataLength float64 `json:"encodedDataLength"`
}

/*
RequestInterceptedEvent represents Network.requestIntercepted event data.
*/
type RequestInterceptedEvent struct {
	// Each request the page makes will have a unique id, however if any redirects are encountered
	// while processing that fetch, they will be reported with the same id as the original fetch.
	// Likewise if HTTP authentication is needed then the same fetch id will be used.
	InterceptionID InterceptionID `json:"interceptionId"`

	// desc.
	Request Request `json:"request"`

	// The ID of the frame that initiated the request.
	FrameID Page.FrameID `json:"frameId"`

	// How the requested resource will be used.
	ResourceType Page.ResourceType `json:"resourceType"`

	// Whether this is a navigation request, which can abort the navigation completely.
	IsNavigationRequest bool `json:"isNavigationRequest"`

	// Optional. Redirect location, only sent if a redirect was intercepted.
	RedirectURL string `json:"redirectUrl,omitempty"`

	// Optional. Details of the Authorization Challenge encountered. If this is set then
	// continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge AuthChallenge `json:"authChallenge,omitempty"`

	// Optional. Response error if intercepted at response stage or if redirect occurred while
	// intercepting request.
	ResponseErrorReason ErrorReason `json:"responseErrorReason,omitempty"`

	// Optional. Response code if intercepted at response stage or if redirect occurred while
	// intercepting request or auth retry occurred.
	ResponseStatusCode int `json:"responseStatusCode,omitempty"`

	// Optional. Response headers if intercepted at the response stage or if redirect occurred while
	// intercepting request or auth retry occurred.
	ResponseHeaders Headers `json:"responseHeaders,omitempty"`
}

/*
RequestServedFromCacheEvent represents Network.requestServedFromCache event data.
*/
type RequestServedFromCacheEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`
}

/*
RequestWillBeSentEvent represents Network.requestWillBeSent event data.
*/
type RequestWillBeSentEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID LoaderID `json:"loaderId"`

	// URL of the document this request is loaded for.
	DocumentURL string `json:"documentURL"`

	// Request data.
	Request Request `json:"request"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Timestamp.
	WallTime TimeSinceEpoch `json:"wallTime"`

	// Request initiator.
	Initiator Initiator `json:"initiator"`

	// Redirect response data.
	RedirectResponse Response `json:"redirectResponse"`

	// Optional. Type of this resource.
	Type Page.ResourceType `json:"type,omitempty"`

	// Optional. Frame identifier.
	FrameID Page.FrameID `json:"frameId,omitempty"`
}

/*
ResourceChangedPriorityEvent represents Network.resourceChangedPriority event data.
*/
type ResourceChangedPriorityEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// New priority.
	NewPriority ResourcePriority `json:"newPriority"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`
}

/*
ResponseReceivedEvent represents Network.responseReceived event data.
*/
type ResponseReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID LoaderID `json:"loaderId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Resource type.
	Type Page.ResourceType `json:"type"`

	// Response data.
	Response Response `json:"response"`

	// Optional. Frame identifier.
	FrameID Page.FrameID `json:"frameId,omitempty"`
}

/*
WebSocketClosedEvent represents Network.webSocketClosed event data.
*/
type WebSocketClosedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`
}

/*
WebSocketCreatedEvent represents Network.webSocketCreated event data.
*/
type WebSocketCreatedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Optional. WebSocket frame error message.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

/*
WebSocketFrameReceivedEvent represents Network.webSocketFrameReceived event data.
*/
type WebSocketFrameReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// WebSocket response data.
	Response WebSocketFrame `json:"response"`
}

/*
WebSocketFrameSentEvent represents Network.webSocketFrameSent event data.
*/
type WebSocketFrameSentEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// WebSocket response data.
	Response WebSocketFrame `json:"response"`
}

/*
WebSocketHandshakeResponseReceivedEvent represents Network.webSocketHandshakeResponseReceived event
data.
*/
type WebSocketHandshakeResponseReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// WebSocket response data.
	Response WebSocketFrame `json:"response"`
}

/*
WebSocketWillSendHandshakeRequestEvent represents Network.webSocketWillSendHandshakeRequest event
data.
*/
type WebSocketWillSendHandshakeRequestEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// UTC Timestamp.
	WallTime TimeSinceEpoch `json:"wallTime"`

	// WebSocket request data.
	Request WebSocketRequest `json:"request"`
}

/*
LoaderID is the Unique loader identifier.
*/
type LoaderID string

/*
RequestID is the unique request identifier.
*/
type RequestID string

/*
InterceptionID is the unique intercepted request identifier.
*/
type InterceptionID string

/*
ErrorReason is the network level fetch failure reason.
*/
type ErrorReason string

func (s ErrorReason) String() string {
	str := string(s)
	if str == "Failed" ||
		str == "Aborted" ||
		str == "TimedOut" ||
		str == "AccessDenied" ||
		str == "ConnectionClosed" ||
		str == "ConnectionReset" ||
		str == "ConnectionRefused" ||
		str == "ConnectionAborted" ||
		str == "ConnectionFailed" ||
		str == "NameNotResolved" ||
		str == "InternetDisconnected" ||
		str == "AddressUnreachable" {
		return str
	}
	panic(fmt.Errorf("Invalid ErrorReason '%s'", str))
}

/*
TimeSinceEpoch represents UTC time in seconds, counted from January 1, 1970.
*/
type TimeSinceEpoch int

/*
MonotonicTime is the monotonically increasing time in seconds since an arbitrary point in the past.
*/
type MonotonicTime int

/*
Headers contains request / response headers as keys / values of JSON object.
*/
type Headers map[string]string

/*
ConnectionType is the underlying connection technology that the browser is supposedly using.
*/
type ConnectionType string

func (s ConnectionType) String() string {
	str := string(s)
	if str == "none" ||
		str == "cellular2g" ||
		str == "cellular3g" ||
		str == "cellular4g" ||
		str == "bluetooth" ||
		str == "ethernet" ||
		str == "wifi" ||
		str == "wimax" ||
		str == "other" {
		return str
	}
	panic(fmt.Errorf("Invalid ConnectionType '%s'", str))
}

/*
CookieSameSite represents the cookie's 'SameSite' status
https://tools.ietf.org/html/draft-west-first-party-cookies
*/
type CookieSameSite string

func (s CookieSameSite) String() string {
	str := string(s)
	if str == "Strict" ||
		str == "Lax" {
		return str
	}
	panic(fmt.Errorf("Invalid CookieSameSite '%s'", str))
}

/*
ResourceTiming defines the timing information for the request
*/
type ResourceTiming struct {
	// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in
	// milliseconds relatively to this requestTime.
	RequestTime int `json:"requestTime"`

	// Started resolving proxy.
	ProxyStart int `json:"proxyStart"`

	// Finished resolving proxy.
	ProxyEnd int `json:"proxyEnd"`

	// Started DNS address resolve.
	DNSStart int `json:"dnsStart"`

	// Finished DNS address resolve.
	DNSEnd int `json:"dnsEnd"`

	// Started connecting to the remote host.
	ConnectStart int `json:"connectStart"`

	// Connected to the remote host.
	ConnectEnd int `json:"connectEnd"`

	// Started SSL handshake.
	SSLStart int `json:"sslStart"`

	// Finished SSL handshake.
	SSLEnd int `json:"sslEnd"`

	// Started running ServiceWorker. EXPERIMENTAL
	WorkerStart int `json:"workerStart"`

	// Finished Starting ServiceWorker. EXPERIMENTAL
	WorkerReady int `json:"workerReady"`

	// Started sending request.
	SendStart int `json:"sendStart"`

	// Finished sending request.
	SendEnd int `json:"sendEnd"`

	// Time the server started pushing request. EXPERIMENTAL
	PushStart int `json:"pushStart"`

	// Time the server finished pushing request. EXPERIMENTAL
	PushEnd int `json:"pushEnd"`

	// Finished receiving response headers.
	ReceiveHeadersEnd int `json:"receiveHeadersEnd"`
}

/*
ResourcePriority represents the loading priority of a resource request.
*/
type ResourcePriority string

func (s ResourcePriority) String() string {
	str := string(s)
	if str == "VeryLow" ||
		str == "Low" ||
		str == "Medium" ||
		str == "High" {
		return str
	}
	panic(fmt.Errorf("Invalid ResourcePriority '%s'", str))
}

/*
Request represents the HTTP request data.
*/
type Request struct {
	// Request URL.
	URL string `json:"url"`

	// HTTP request method.
	Method string `json:"method"`

	// HTTP request headers.
	Headers Headers `json:"headers"`

	// Optional. HTTP POST request data.
	PostData string `json:"postData,omitempty"`

	// Optional. The mixed content type of the request.
	MixedContentType Security.MixedContentType `json:"mixedContentType,omitempty"`

	// Priority of the resource request at the time request is sent.
	InitialPriority ResourcePriority `json:"initialPriority"`

	// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	// Allowed values: unsafe-url, no-referrer-when-downgrade, no-referrer, origin,
	// origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin.
	ReferrerPolicy string `json:"referrerPolicy"`

	// Optional. Whether is loaded via link preload.
	IsLinkPreload bool `json:"isLinkPreload,omitempty"`
}

/*
SignedCertificateTimestamp contains details of a signed certificate timestamp (SCT).
*/
type SignedCertificateTimestamp struct {

	// Validation status.
	Status string `json:"status"`

	// Origin.
	Origin string `json:"origin"`

	// Log name / description.
	LogDescription string `json:"logDescription"`

	// Log ID.
	LogID string `json:"logId"`

	// Issuance date.
	Timestamp TimeSinceEpoch `json:"timestamp"`

	// Hash algorithm.
	HashAlgorithm string `json:"hashAlgorithm"`

	// Signature algorithm.
	SignatureAlgorithm string `json:"signatureAlgorithm"`

	// Signature data.
	SignatureData string `json:"signatureData"`
}

/*
SecurityDetails defines security details about a request.
*/
type SecurityDetails struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string `json:"protocol"`

	// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange string `json:"keyExchange"`

	// Optional. (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`

	// Cipher name.
	Cipher string `json:"cipher"`

	// Optional. TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac string `json:"mac,omitempty"`

	// Certificate ID value.
	CertificateID Security.CertificateID `json:"certificateId"`

	// Certificate subject name.
	SubjectName string `json:"subjectName"`

	// Subject Alternative Name (SAN) DNS names and IP addresses.
	SanList []string `json:"sanList"`

	// Name of the issuing CA.
	Issuer string `json:"issuer"`

	// Certificate valid from date.
	ValidFrom TimeSinceEpoch `json:"validFrom"`

	// Certificate valid to (expiration) date
	ValidTo TimeSinceEpoch `json:"validTo"`

	// List of signed certificate timestamps (SCTs).
	SignedCertificateTimestampList []*SignedCertificateTimestamp `json:"signedCertificateTimestampList"`
}

/*
BlockedReason defines the reason why request was blocked.
*/
type BlockedReason int

const (
	_csp BlockedReason = iota
	_mixedContent
	_origin
	_inspector
	_subresourceFilter
	_other
)

func (s BlockedReason) String() string {
	str := string(s)
	if str == "csp" ||
		str == "mixed-content" ||
		str == "origin" ||
		str == "inspector" ||
		str == "subresource-filter" ||
		str == "other" {
		return str
	}
	panic(fmt.Errorf("Invalid BlockedReason '%s'", str))
}

/*
Response contains HTTP response data.
*/
type Response struct {

	// Response URL. This URL can be different from CachedResource.url in case of redirect.
	URL string `json:"url"`

	// HTTP response status code.
	Status int `json:"status"`

	// HTTP response status text.
	StatusText string `json:"statusText"`

	// HTTP response headers.
	Headers Headers `json:"headers"`

	// Optional. HTTP response headers text.
	HeadersText string `json:"headersText,omitempty"`

	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`

	// Optional. Refined HTTP request headers that were actually transmitted over the network.
	RequestHeaders Headers `json:"requestHeaders,omitempty"`

	// Optional. HTTP request headers text.
	RequestHeadersText string `json:"requestHeadersText,omitempty"`

	// Specifies whether physical connection was actually reused for this request.
	ConnectionReused bool `json:"connectionReused"`

	// Physical connection id that was actually used for this request.
	ConnectionID int `json:"connectionId"`

	// Optional. Remote IP address.
	RemoteIPAddress string `json:"remoteIPAddress,omitempty"`

	// Optional. Remote port.
	RemotePort int `json:"remotePort,omitempty"`

	// Optional. Specifies that the request was served from the disk cache.
	FromDiskCache bool `json:"fromDiskCache,omitempty"`

	// Optional. Specifies that the request was served from the ServiceWorker.
	FromServiceWorker bool `json:"fromServiceWorker,omitempty"`

	// Total number of bytes received for this request so far.
	EncodedDataLength int `json:"encodedDataLength"`

	// Optional. Timing information for the given request.
	Timing ResourceTiming `json:"timing,omitempty"`

	// Optional. Protocol used to fetch this request.
	Protocol string `json:"protocol,omitempty"`

	// Security state of the request resource.
	SecurityState Security.SecurityState `json:"securityState"`

	// Optional. Security details for the request.
	SecurityDetails SecurityDetails `json:"securityDetails,omitempty"`
}

/*
WebSocketRequest contains WebSocket request data
*/
type WebSocketRequest struct {
	Headers Headers `json:"headers"`
}

/*
WebSocketResponse contains WebSocket response data.
*/
type WebSocketResponse struct {
	// HTTP response status code.
	Status int `json:"status"`

	// HTTP response status text.
	StatusText string `json:"statusText"`

	// HTTP response headers.
	Headers Headers `json:"headers"`

	// Optional. HTTP response headers text.
	HeadersText string `json:"headersText,omitempty"`

	// Optional. HTTP request headers.
	RequestHeaders Headers `json:"requestHeaders,omitempty"`

	// Optional. HTTP request headers text.
	RequestHeadersText string `json:"requestHeadersText,omitempty"`
}

/*
WebSocketFrame contains WebSocket frame data.
*/
type WebSocketFrame struct {
	// WebSocket frame opcode.
	Opcode int `json:"opcode"`

	// WebSocket frame mask.
	Mask bool `json:"mask"`

	// WebSocket frame payload data.
	PayloadData string `json:"payloadData"`
}

/*
CachedResource contains information about the cached resource.
*/
type CachedResource struct {
	// Resource URL. This is the url of the original network request.
	URL string `json:"url"`

	// Type of this resource.
	Type Page.ResourceType `json:"type"`

	// Optional. Cached response data.
	Response Response `json:"response,omitempty"`

	// Cached response body size.
	BodySize int `json:"bodySize"`
}

/*
Initiator contains information about the request initiator.
*/
type Initiator struct {
	// Type of this initiator. Allowed values: parser, script, preload, other.
	Type string `json:"type"`

	// Optional. Initiator JavaScript stack trace, set for Script only.
	Stack Runtime.StackTrace `json:"stack,omitempty"`

	// Optional. Initiator URL, set for Parser type or for Script type (when script is importing
	// module).
	URL string `json:"url,omitempty"`

	// Optional. Initiator line number, set for Parser type or for Script type (when script is
	// importing module) (0-based).
	LineNumber int `json:"lineNumber,omitempty"`
}

/*
Cookie object
*/
type Cookie struct {
	// Cookie name.
	Name string `json:"name"`

	// Cookie value.
	Value string `json:"value"`

	// Cookie domain.
	Domain string `json:"domain"`

	// Cookie path.
	Path string `json:"path"`

	// Cookie expiration date as the number of seconds since the UNIX epoch.
	Expires int `json:"expires"`

	// Cookie size.
	Size int `json:"size"`

	// True if cookie is http-only.
	HTTPOnly bool `json:"httpOnly"`

	// True if cookie is secure.
	Secure bool `json:"secure"`

	// True in case of session cookie.
	Session bool `json:"session"`

	// Optional. Cookie SameSite type.
	SameSite CookieSameSite `json:"sameSite,omitempty"`
}

/*
CookieParam is a cookie parameter
*/
type CookieParam struct {
	// Cookie name.
	Name string `json:"name"`

	// Cookie value.
	Value string `json:"value"`

	// Optional. The request-URI to associate with the setting of the cookie. This value can affect
	// the default domain and path values of the created cookie.
	URL string `json:"url,omitempty"`

	// Optional. Cookie domain.
	Domain string `json:"domain,omitempty"`

	// Optional. Cookie path.
	Path string `json:"path,omitempty"`

	// Optional. True if cookie is secure.
	Secure bool `json:"secure,omitempty"`

	// Optional. True if cookie is http-only.
	HTTPOnly bool `json:"httpOnly,omitempty"`

	// Optional. Cookie SameSite type.
	SameSite CookieSameSite `json:"sameSite,omitempty"`

	// Optional. Cookie expiration date, session cookie if not set.
	Expires TimeSinceEpoch `json:"expires,omitempty"`
}

/*
AuthChallenge is an authorization challenge for HTTP status code 401 or 407. EXPERIMENTAL
*/
type AuthChallenge struct {
	// Optional. Source of the authentication challenge. Allowed values: Server, Proxy.
	Source string `json:"source,omitempty"`

	// Origin of the challenger.
	Origin string `json:"origin"`

	// The authentication scheme used, such as basic or digest.
	Scheme string `json:"scheme"`

	// The realm of the challenge. May be empty.
	Realm string `json:"realm"`
}

/*
AuthChallengeResponse is the response to an AuthChallenge. EXPERIMENTAL
*/
type AuthChallengeResponse struct {
	// The decision on what to do in response to the authorization challenge. Default means
	// deferring to the default behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box. Allowed values: Default, CancelAuth,
	// ProvideCredentials.
	Response string `json:"response"`

	// Optional. The username to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Username string `json:"username,omitempty"`

	// Optional. The password to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Password string `json:"password,omitempty"`
}

/*
InterceptionStage represents stages of the interception to begin intercepting. Request will
intercept before the request is sent. Response will intercept after the response is received.
EXPERIMENTAL
*/
type InterceptionStage string

func (s InterceptionStage) String() string {
	str := string(s)
	if str == "Request" ||
		str == "HeadersReceived" {
		return str
	}
	panic(fmt.Errorf("Invalid InterceptionStage '%s'", str))
}

/*
RequestPattern is the request pattern for interception. EXPERIMENTAL
*/
type RequestPattern struct {
	// Optional. Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character
	// is backslash. Omitting is equivalent to "*".
	URLPattern string `json:"urlPattern,omitempty"`

	// Optional. If set, only requests for matching resource types will be intercepted.
	ResourceType Page.ResourceType `json:"resourceType,omitempty"`

	// Optional. Stage at which to begin intercepting requests. Default is Request.
	InterceptionStage InterceptionStage `json:"interceptionStage,omitempty"`
}
