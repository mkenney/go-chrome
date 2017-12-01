package Network

import (
	Page "app/chrome/page"
	Runtime "app/chrome/runtime"
	Security "app/chrome/security"
	"fmt"
)

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

	// Optional. Stage at wich to begin intercepting requests. Default is Request.
	InterceptionStage InterceptionStage `json:"interceptionStage,omitempty"`
}
