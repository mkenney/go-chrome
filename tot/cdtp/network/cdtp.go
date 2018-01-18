/*
Package network provides type definitions for use with the Chrome Network protocol

https://chromedevtools.github.io/devtools-protocol/tot/Network/
*/
package network

import (
	"github.com/mkenney/go-chrome/tot/cdtp/page"
	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
	"github.com/mkenney/go-chrome/tot/cdtp/security"
)

/*
LoaderID is the Unique loader identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoaderId
*/
type LoaderID string

/*
RequestID is the unique request identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-RequestId
*/
type RequestID string

/*
InterceptionID is the unique intercepted request identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-InterceptionId
*/
type InterceptionID string

/*
TimeSinceEpoch represents UTC time in seconds, counted from January 1, 1970.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TimeSinceEpoch
*/
type TimeSinceEpoch int64

/*
MonotonicTime is the monotonically increasing time in seconds since an arbitrary point in the past.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-MonotonicTime
*/
type MonotonicTime int

/*
Headers contains request / response headers as keys / values of JSON object.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Headers
*/
type Headers map[string]string

/*
ResourceTiming defines the timing information for the request

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ResourceTiming
*/
type ResourceTiming struct {
	// Timing's requestTime is a baseline in seconds, while the other numbers
	// are ticks in milliseconds relatively to this requestTime.
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

	// Started running ServiceWorker. EXPERIMENTAL.
	WorkerStart int `json:"workerStart"`

	// Finished Starting ServiceWorker. EXPERIMENTAL.
	WorkerReady int `json:"workerReady"`

	// Started sending request.
	SendStart int `json:"sendStart"`

	// Finished sending request.
	SendEnd int `json:"sendEnd"`

	// Time the server started pushing request. EXPERIMENTAL.
	PushStart int `json:"pushStart"`

	// Time the server finished pushing request. EXPERIMENTAL.
	PushEnd int `json:"pushEnd"`

	// Finished receiving response headers.
	ReceiveHeadersEnd int `json:"receiveHeadersEnd"`
}

/*
Request represents the HTTP request data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Request
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
	MixedContentType security.MixedContentType `json:"mixedContentType,omitempty"`

	// Priority of the resource request at the time request is sent. Allowed
	// values:
	//	- ResourcePriority.VeryLow
	//	- ResourcePriority.Low
	//	- ResourcePriority.Medium
	//	- ResourcePriority.High
	//	- ResourcePriority.VeryHigh
	InitialPriority ResourcePriorityEnum `json:"initialPriority"`

	// The referrer policy of the request, as defined in
	// https://www.w3.org/TR/referrer-policy/ Allowed values:
	//	- ReferrerPolicy.UnsafeUrl
	//	- ReferrerPolicy.NoReferrerWhenDowngrade
	//	- ReferrerPolicy.NoReferrer
	//	- ReferrerPolicy.Origin
	//	- ReferrerPolicy.OriginWhenCrossOrigin
	//	- ReferrerPolicy.SameOrigin
	//	- ReferrerPolicy.StrictOrigin
	//	- ReferrerPolicy.StrictOriginWhenCrossOrigin
	ReferrerPolicy ReferrerPolicyEnum `json:"referrerPolicy"`

	// Optional. Whether is loaded via link preload.
	IsLinkPreload bool `json:"isLinkPreload,omitempty"`
}

/*
SignedCertificateTimestamp contains details of a signed certificate timestamp (SCT).

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedCertificateTimestamp
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

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SecurityDetails
*/
type SecurityDetails struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string `json:"protocol"`

	// Key Exchange used by the connection, or the empty string if not
	// applicable.
	KeyExchange string `json:"keyExchange"`

	// Optional. (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`

	// Cipher name.
	Cipher string `json:"cipher"`

	// Optional. TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac string `json:"mac,omitempty"`

	// Certificate ID value.
	CertificateID security.CertificateID `json:"certificateId"`

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
Response contains HTTP response data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Response
*/
type Response struct {

	// Response URL. This URL can be different from CachedResource.url in case
	// of redirect.
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

	// Optional. Refined HTTP request headers that were actually transmitted
	// over the network.
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
	Timing *ResourceTiming `json:"timing,omitempty"`

	// Optional. Protocol used to fetch this request.
	Protocol string `json:"protocol,omitempty"`

	// Security state of the request resource.
	SecurityState security.State `json:"securityState"`

	// Optional. Security details for the request.
	SecurityDetails *SecurityDetails `json:"securityDetails,omitempty"`
}

/*
WebSocketRequest contains WebSocket request data

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketRequest
*/
type WebSocketRequest struct {
	Headers Headers `json:"headers"`
}

/*
WebSocketResponse contains WebSocket response data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketResponse
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

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketFrame
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

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CachedResource
*/
type CachedResource struct {
	// Resource URL. This is the url of the original network request.
	URL string `json:"url"`

	// Type of this resource.
	Type page.ResourceType `json:"type"`

	// Optional. Cached response data.
	Response *Response `json:"response,omitempty"`

	// Cached response body size.
	BodySize int `json:"bodySize"`
}

/*
Initiator contains information about the request initiator.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Initiator
*/
type Initiator struct {
	// Type of this initiator. Allowed values:
	//	- InitiatorType.Parser
	//	- InitiatorType.Script
	//	- InitiatorType.Preload
	//	- InitiatorType.Other
	Type InitiatorTypeEnum `json:"type"`

	// Optional. Initiator JavaScript stack trace, set for Script only.
	Stack *runtime.StackTrace `json:"stack,omitempty"`

	// Optional. Initiator URL, set for Parser type or for Script type (when
	// script is importing module).
	URL string `json:"url,omitempty"`

	// Optional. Initiator line number, set for Parser type or for Script type
	// (when script is importing module) (0-based).
	LineNumber int `json:"lineNumber,omitempty"`
}

/*
Cookie object

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Cookie
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
	Expires int64 `json:"expires"`

	// Cookie size.
	Size int `json:"size"`

	// True if cookie is http-only.
	HTTPOnly bool `json:"httpOnly"`

	// True if cookie is secure.
	Secure bool `json:"secure"`

	// True in case of session cookie.
	Session bool `json:"session"`

	// Optional. Cookie SameSite type. Allowed values:
	//	- CookieSameSite.Strict
	//	- CookieSameSite.Lax
	SameSite CookieSameSiteEnum `json:"sameSite,omitempty"`
}

/*
AuthChallenge is an authorization challenge for HTTP status code 401 or 407. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallenge
*/
type AuthChallenge struct {
	// Optional. Source of the authentication challenge. Allowed values:
	//	- Source.Server
	//	- Source.Proxy
	Source SourceEnum `json:"source,omitempty"`

	// Origin of the challenger.
	Origin string `json:"origin"`

	// The authentication scheme used, such as basic or digest.
	Scheme string `json:"scheme"`

	// The realm of the challenge. May be empty.
	Realm string `json:"realm"`
}

/*
AuthChallengeResponse is the response to an AuthChallenge. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallengeResponse
*/
type AuthChallengeResponse struct {
	// The decision on what to do in response to the authorization challenge.
	// Default means deferring to the default behavior of the net stack, which
	// will likely either the Cancel authentication or display a popup dialog
	// box. Allowed values:
	//	- ChallengeResponse.Default
	//	- ChallengeResponse.CancelAuth
	//	- ChallengeResponse.ProvideCredentials
	Response ChallengeResponseEnum `json:"response"`

	// Optional. The username to provide, possibly empty. Should only be set if
	// response is ProvideCredentials.
	Username string `json:"username,omitempty"`

	// Optional. The password to provide, possibly empty. Should only be set if
	// response is ProvideCredentials.
	Password string `json:"password,omitempty"`
}

/*
RequestPattern is the request pattern for interception. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-RequestPattern
*/
type RequestPattern struct {
	// Optional. Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed.
	// Escape character is backslash. Omitting is equivalent to "*".
	URLPattern string `json:"urlPattern,omitempty"`

	// Optional. If set, only requests for matching resource types will be
	// intercepted.
	ResourceType page.ResourceType `json:"resourceType,omitempty"`

	// Optional. Stage at which to begin intercepting requests. Default is
	// Request. Allowed values:
	//	- InterceptionStage.Request
	//	- InterceptionStage.HeadersReceived
	InterceptionStage InterceptionStageEnum `json:"interceptionStage,omitempty"`
}
