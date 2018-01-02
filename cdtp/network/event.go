package network

import (
	"github.com/mkenney/go-chrome/cdtp/page"
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
	Type *page.ResourceType `json:"type"`

	// User friendly error message.
	ErrorText string `json:"errorText"`

	// Optional. True if loading was canceled.
	Canceled bool `json:"canceled,omitempty"`

	// Optional. The reason why loading was blocked, if any.
	BlockedReason BlockedReason `json:"blockedReason,omitempty"`
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
	ResourceType page.ResourceType `json:"resourceType"`

	// Whether this is a navigation request, which can abort the navigation
	// completely.
	IsNavigationRequest bool `json:"isNavigationRequest"`

	// Optional. Redirect location, only sent if a redirect was intercepted.
	RedirectURL string `json:"redirectUrl,omitempty"`

	// Optional. Details of the Authorization Challenge encountered. If this is
	// set then continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge *AuthChallenge `json:"authChallenge,omitempty"`

	// Optional. Response error if intercepted at response stage or if redirect
	// occurred while intercepting request.
	ResponseErrorReason ErrorReason `json:"responseErrorReason,omitempty"`

	// Optional. Response code if intercepted at response stage or if redirect
	// occurred while intercepting request or auth retry occurred.
	ResponseStatusCode int `json:"responseStatusCode,omitempty"`

	// Optional. Response headers if intercepted at the response stage or if
	// redirect occurred while intercepting request or auth retry occurred.
	ResponseHeaders Headers `json:"responseHeaders,omitempty"`
}

/*
RequestServedFromCacheEvent represents Network.requestServedFromCache event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestServedFromCache
*/
type RequestServedFromCacheEvent struct {
	// Request identifier.
	RequestID RequestID `json:"requestId"`
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
	Type page.ResourceType `json:"type,omitempty"`

	// Optional. Frame identifier.
	FrameID page.FrameID `json:"frameId,omitempty"`
}

/*
ResourceChangedPriorityEvent represents Network.resourceChangedPriority event data.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-resourceChangedPriority
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
	Type page.ResourceType `json:"type"`

	// Response data.
	Response Response `json:"response"`

	// Optional. Frame identifier.
	FrameID page.FrameID `json:"frameId,omitempty"`
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
}
