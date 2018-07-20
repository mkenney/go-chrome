package network

import (
	"github.com/mkenney/go-chrome/tot/page"
)

/*
DataReceivedEvent represents Network.dataReceived event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-dataReceived
*/
type DataReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Data chunk length.
	DataLength int `json:"dataLength"`

	// Actual bytes received (might be less than dataLength for compressed
	// encodings).
	EncodedDataLength int `json:"encodedDataLength"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
EventSourceMessageReceivedEvent represents Network.eventSourceMessageReceived event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-eventSourceMessageReceived
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

	// Error information related to this event
	Err error `json:"-"`
}

/*
LoadingFailedEvent represents Network.loadingFailed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFailed
*/
type LoadingFailedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Resource type.
	Type page.ResourceTypeEnum `json:"type"`

	// User friendly error message.
	ErrorText string `json:"errorText"`

	// Optional. True if loading was canceled.
	Canceled bool `json:"canceled,omitempty"`

	// Optional. The reason why loading was blocked, if any. Allowed values:
	//	- BlockedReason.Csp
	//	- BlockedReason.MixedContent
	//	- BlockedReason.Origin
	//	- BlockedReason.Inspector
	//	- BlockedReason.SubresourceFilter
	//	- BlockedReason.Other
	BlockedReason BlockedReasonEnum `json:"blockedReason,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
LoadingFinishedEvent represents Network.loadingFinished event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFinished
*/
type LoadingFinishedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Total number of bytes received for this request.
	EncodedDataLength float64 `json:"encodedDataLength"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
RequestInterceptedEvent represents Network.requestIntercepted event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestIntercepted
*/
type RequestInterceptedEvent struct {
	// Each request the page makes will have a unique id, however if any
	// redirects are encountered while processing that fetch, they will be
	// reported with the same id as the original fetch. Likewise if HTTP
	// authentication is needed then the same fetch id will be used.
	InterceptionID InterceptionID `json:"interceptionId"`

	// desc.
	Request *Request `json:"request"`

	// The ID of the frame that initiated the request.
	FrameID page.FrameID `json:"frameId"`

	// How the requested resource will be used.
	ResourceType page.ResourceTypeEnum `json:"resourceType"`

	// Whether this is a navigation request, which can abort the navigation
	// completely.
	IsNavigationRequest bool `json:"isNavigationRequest"`

	// Optional. Redirect location, only sent if a redirect was intercepted.
	RedirectURL string `json:"redirectUrl,omitempty"`

	// Optional. Details of the Authorization Challenge encountered. If this is
	// set then continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge *AuthChallenge `json:"authChallenge,omitempty"`

	// Optional. Response error if intercepted at response stage or if redirect
	// occurred while intercepting request. Allowed values:
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
	ResponseErrorReason ErrorReasonEnum `json:"responseErrorReason,omitempty"`

	// Optional. Response code if intercepted at response stage or if redirect
	// occurred while intercepting request or auth retry occurred.
	ResponseStatusCode int `json:"responseStatusCode,omitempty"`

	// Optional. Response headers if intercepted at the response stage or if
	// redirect occurred while intercepting request or auth retry occurred.
	ResponseHeaders Headers `json:"responseHeaders,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
RequestServedFromCacheEvent represents Network.requestServedFromCache event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestServedFromCache
*/
type RequestServedFromCacheEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
RequestWillBeSentEvent represents Network.requestWillBeSent event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestWillBeSent
*/
type RequestWillBeSentEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID LoaderID `json:"loaderId"`

	// URL of the document this request is loaded for.
	DocumentURL string `json:"documentURL"`

	// Request data.
	Request *Request `json:"request"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Timestamp.
	WallTime TimeSinceEpoch `json:"wallTime"`

	// Request initiator.
	Initiator *Initiator `json:"initiator"`

	// Redirect response data.
	RedirectResponse *Response `json:"redirectResponse"`

	// Optional. Type of this resource.
	Type page.ResourceTypeEnum `json:"type,omitempty"`

	// Optional. Frame identifier.
	FrameID page.FrameID `json:"frameId,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ResourceChangedPriorityEvent represents Network.resourceChangedPriority event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-resourceChangedPriority
*/
type ResourceChangedPriorityEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// New priority. Allowed values:
	//	- ResourcePriority.VeryLow
	//	- ResourcePriority.Low
	//	- ResourcePriority.Medium
	//	- ResourcePriority.High
	//	- ResourcePriority.VeryHigh
	NewPriority ResourcePriorityEnum `json:"newPriority"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ResponseReceivedEvent represents Network.responseReceived event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-responseReceived
*/
type ResponseReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID LoaderID `json:"loaderId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Resource type.
	Type page.ResourceTypeEnum `json:"type"`

	// Response data.
	Response *Response `json:"response"`

	// Optional. Frame identifier.
	FrameID page.FrameID `json:"frameId,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WebSocketClosedEvent represents Network.webSocketClosed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketClosed
*/
type WebSocketClosedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WebSocketCreatedEvent represents Network.webSocketCreated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketCreated
*/
type WebSocketCreatedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Optional. WebSocket frame error message.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WebSocketFrameErrorEvent represents Network.webSocketFrameError event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameError
*/
type WebSocketFrameErrorEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// Optional. WebSocket frame error message.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WebSocketFrameReceivedEvent represents Network.webSocketFrameReceived event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameReceived
*/
type WebSocketFrameReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// WebSocket response data.
	Response *WebSocketFrame `json:"response"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WebSocketFrameSentEvent represents Network.webSocketFrameSent event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameSent
*/
type WebSocketFrameSentEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// WebSocket response data.
	Response *WebSocketFrame `json:"response"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WebSocketHandshakeResponseReceivedEvent represents Network.webSocketHandshakeResponseReceived event
data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketHandshakeResponseReceived
*/
type WebSocketHandshakeResponseReceivedEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// WebSocket response data.
	Response *WebSocketFrame `json:"response"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WebSocketWillSendHandshakeRequestEvent represents Network.webSocketWillSendHandshakeRequest event
data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketWillSendHandshakeRequest
*/
type WebSocketWillSendHandshakeRequestEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`

	// Timestamp.
	Timestamp MonotonicTime `json:"timestamp"`

	// UTC Timestamp.
	WallTime TimeSinceEpoch `json:"wallTime"`

	// WebSocket request data.
	Request *WebSocketRequest `json:"request"`

	// Error information related to this event
	Err error `json:"-"`
}
