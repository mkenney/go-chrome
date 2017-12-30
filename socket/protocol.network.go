package socket

import (
	"encoding/json"

	network "github.com/mkenney/go-chrome/protocol/network"
	log "github.com/sirupsen/logrus"
)

/*
NetworkProtocol provides a namespace for the Chrome Network protocol methods. The Network protocol
allows tracking network activities of the page. It exposes information about http, file, data and
other requests and responses, their headers, bodies, timing, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Network/
*/
type NetworkProtocol struct {
	Socket Socketer
}

/*
CanClearBrowserCache tells whether clearing browser cache is supported. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCache
*/
func (protocol *NetworkProtocol) CanClearBrowserCache() (*network.CanClearBrowserCacheResult, error) {
	command := NewCommand("Network.canClearBrowserCache", nil)
	result := &network.CanClearBrowserCacheResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CanClearBrowserCookies tells whether clearing browser cookies is supported. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCookies
*/
func (protocol *NetworkProtocol) CanClearBrowserCookies() (*network.CanClearBrowserCookiesResult, error) {
	command := NewCommand("Network.canClearBrowserCookies", nil)
	result := &network.CanClearBrowserCookiesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CanEmulateConditions tells whether emulation of network conditions is supported. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canEmulateNetworkConditions
*/
func (protocol *NetworkProtocol) CanEmulateConditions() (*network.CanEmulateConditionsResult, error) {
	command := NewCommand("Network.canEmulateNetworkConditions", nil)
	result := &network.CanEmulateConditionsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
ClearBrowserCache clears browser cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache
*/
func (protocol *NetworkProtocol) ClearBrowserCache() error {
	command := NewCommand("Network.clearBrowserCache", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ClearBrowserCookies clears browser cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCookies
*/
func (protocol *NetworkProtocol) ClearBrowserCookies() error {
	command := NewCommand("Network.clearBrowserCookies", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ContinueInterceptedRequest response to Network.requestIntercepted which either modifies the request
to continue with any modifications, or blocks it, or completes it with the provided response bytes.
If a network fetch occurs as a result which encounters a redirect an additional
Network.requestIntercepted event will be sent with the same InterceptionID.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-continueInterceptedRequest
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) ContinueInterceptedRequest(
	params *network.ContinueInterceptedRequestParams,
) error {
	command := NewCommand("Network.continueInterceptedRequest", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DeleteCookies deletes browser cookies with matching name and url or domain/path pair.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
*/
func (protocol *NetworkProtocol) DeleteCookies(
	params *network.DeleteCookiesParams,
) error {
	command := NewCommand("Network.deleteCookies", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables network tracking, prevents network events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-disable
*/
func (protocol *NetworkProtocol) Disable() error {
	command := NewCommand("Network.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
EmulateConditions activates emulation of network conditions.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions
*/
func (protocol *NetworkProtocol) EmulateConditions(
	params *network.EmulateConditionsParams,
) error {
	command := NewCommand("Network.emulateNetworkConditions", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables network tracking, network events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
*/
func (protocol *NetworkProtocol) Enable(
	params *network.EnableParams,
) error {
	command := NewCommand("Network.enable", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetAllCookies returns all browser cookies. Depending on the backend support, will return detailed
cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getAllCookies
*/
func (protocol *NetworkProtocol) GetAllCookies() (*network.GetAllCookiesResult, error) {
	command := NewCommand("Network.getAllCookies", nil)
	result := &network.GetAllCookiesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetCertificate returns the DER-encoded certificate.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) GetCertificate(
	params *network.GetCertificateParams,
) (*network.GetCertificateResult, error) {
	command := NewCommand("Network.getCertificate", params)
	result := &network.GetCertificateResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetCookies returns all browser cookies for the current URL. Depending on the backend support, will
return detailed cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
*/
func (protocol *NetworkProtocol) GetCookies(
	params *network.GetCookiesParams,
) (*network.GetCookiesResult, error) {
	command := NewCommand("Network.getCookies", params)
	result := &network.GetCookiesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetResponseBody returns content served for the given request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
*/
func (protocol *NetworkProtocol) GetResponseBody(
	params *network.GetResponseBodyParams,
) (*network.GetResponseBodyResult, error) {
	command := NewCommand("Network.getResponseBody", params)
	result := &network.GetResponseBodyResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetResponseBodyForInterception returns content served for the given currently intercepted request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) GetResponseBodyForInterception(
	params *network.GetResponseBodyForInterceptionParams,
) (*network.GetResponseBodyForInterceptionResult, error) {
	command := NewCommand("Network.getResponseBodyForInterception", params)
	result := &network.GetResponseBodyForInterceptionResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
ReplayXHR sends a new XMLHttpRequest which is identical to the original one. The following
parameters should be identical: method, url, async, request body, extra headers, withCredentials
attribute, user, password.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) ReplayXHR(
	params *network.ReplayXHRParams,
) error {
	command := NewCommand("Network.replayXHR", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SearchInResponseBody searches for given string in response content.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SearchInResponseBody(
	params *network.SearchInResponseBodyParams,
) (*network.SearchInResponseBodyResult, error) {
	command := NewCommand("Network.searchInResponseBody", params)
	result := &network.SearchInResponseBodyResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetBlockedURLs blocks URLs from loading.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetBlockedURLs(
	params *network.SetBlockedURLsParams,
) error {
	command := NewCommand("Network.setBlockedURLs", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetBypassServiceWorker toggles ignoring of service worker for each request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetBypassServiceWorker(
	params *network.SetBypassServiceWorkerParams,
) error {
	command := NewCommand("Network.setBypassServiceWorker", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetCacheDisabled toggles ignoring cache for each request. If `true`, cache will not be used.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
*/
func (protocol *NetworkProtocol) SetCacheDisabled(
	params *network.SetCacheDisabledParams,
) error {
	command := NewCommand("Network.setCacheDisabled", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetCookie sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
*/
func (protocol *NetworkProtocol) SetCookie(
	params *network.SetCookieParams,
) (*network.SetCookieResult, error) {
	command := NewCommand("Network.setCookie", params)
	result := &network.SetCookieResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetCookies sets given cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
*/
func (protocol *NetworkProtocol) SetCookies(
	params *network.SetCookiesParams,
) error {
	command := NewCommand("Network.setCookies", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetDataSizeLimitsForTest is for testing.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetDataSizeLimitsForTest(
	params *network.SetDataSizeLimitsForTestParams,
) error {
	command := NewCommand("Network.setDataSizeLimitsForTest", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetExtraHTTPHeaders specifies whether to always send extra HTTP headers with the requests from this
page.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
*/
func (protocol *NetworkProtocol) SetExtraHTTPHeaders(
	params *network.SetExtraHTTPHeadersParams,
) error {
	command := NewCommand("Network.setExtraHTTPHeaders", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetRequestInterception sets the requests to intercept that match a the provided patterns and
optionally resource types.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetRequestInterception(
	params *network.SetRequestInterceptionParams,
) error {
	command := NewCommand("Network.setRequestInterception", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetUserAgentOverride allows overriding user agent with the given string.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setUserAgentOverride
*/
func (protocol *NetworkProtocol) SetUserAgentOverride(
	params *network.SetUserAgentOverrideParams,
) error {
	command := NewCommand("Network.setUserAgentOverride", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnDataReceived adds a handler to the Network.dataReceived event. Network.dataReceived fires when a
data chunk was received over the network.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-dataReceived
*/
func (protocol *NetworkProtocol) OnDataReceived(
	callback func(event *network.DataReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.dataReceived",
		func(response *Response) {
			event := &network.DataReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnEventSourceMessageReceived adds a handler to the Network.eventSourceMessageReceived event.
Network.eventSourceMessageReceived fires when EventSource message is received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-eventSourceMessageReceived
*/
func (protocol *NetworkProtocol) OnEventSourceMessageReceived(
	callback func(event *network.EventSourceMessageReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.eventSourceMessageReceived",
		func(response *Response) {
			event := &network.EventSourceMessageReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLoadingFailed adds a handler to the Network.loadingFailed event. Network.loadingFailed fires when
HTTP request has failed to load.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFailed
*/
func (protocol *NetworkProtocol) OnLoadingFailed(
	callback func(event *network.LoadingFailedEvent),
) {
	handler := NewEventHandler(
		"Network.loadingFailed",
		func(response *Response) {
			event := &network.LoadingFailedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLoadingFinished adds a handler to the Network.loadingFinished event. Network.loadingFinished fires
when HTTP request has finished loading.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFinished
*/
func (protocol *NetworkProtocol) OnLoadingFinished(
	callback func(event *network.LoadingFinishedEvent),
) {
	handler := NewEventHandler(
		"Network.loadingFinished",
		func(response *Response) {
			event := &network.LoadingFinishedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnRequestIntercepted adds a handler to the Network.requestIntercepted event.
Network.requestIntercepted fires when a HTTP request is intercepted and returns details, which must
be either allowed, blocked, modified or mocked.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestIntercepted
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) OnRequestIntercepted(
	callback func(event *network.RequestInterceptedEvent),
) {
	handler := NewEventHandler(
		"Network.requestIntercepted",
		func(response *Response) {
			event := &network.RequestInterceptedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnRequestServedFromCache adds a handler to the Network.requestServedFromCache event.
Network.requestServedFromCache fires when request ended up loading from cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestServedFromCache
*/
func (protocol *NetworkProtocol) OnRequestServedFromCache(
	callback func(event *network.RequestServedFromCacheEvent),
) {
	handler := NewEventHandler(
		"Network.requestServedFromCache",
		func(response *Response) {
			event := &network.RequestServedFromCacheEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnRequestWillBeSent adds a handler to the Network.requestWillBeSent event. Network.requestWillBeSent
fires when the page is about to send HTTP request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestWillBeSent
*/
func (protocol *NetworkProtocol) OnRequestWillBeSent(
	callback func(event *network.RequestWillBeSentEvent),
) {
	handler := NewEventHandler(
		"Network.requestWillBeSent",
		func(response *Response) {
			event := &network.RequestWillBeSentEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnResourceChangedPriority adds a handler to the Network.resourceChangedPriority event.
Network.resourceChangedPriority fires when resource loading priority is changed.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-resourceChangedPriority
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) OnResourceChangedPriority(
	callback func(event *network.ResourceChangedPriorityEvent),
) {
	handler := NewEventHandler(
		"Network.resourceChangedPriority",
		func(response *Response) {
			event := &network.ResourceChangedPriorityEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnResponseReceived adds a handler to the Network.responseReceived event. Network.responseReceived
fires when HTTP response is available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-responseReceived
*/
func (protocol *NetworkProtocol) OnResponseReceived(
	callback func(event *network.ResponseReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.responseReceived",
		func(response *Response) {
			event := &network.ResponseReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketClosed adds a handler to the Network.webSocketClosed event. Network.webSocketClosed
fires when WebSocket is closed.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketClosed
*/
func (protocol *NetworkProtocol) OnWebSocketClosed(
	callback func(event *network.WebSocketClosedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketClosed",
		func(response *Response) {
			event := &network.WebSocketClosedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketCreated adds a handler to the Network.webSocketCreated event. Network.webSocketCreated
fires upon WebSocket creation.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketCreated
*/
func (protocol *NetworkProtocol) OnWebSocketCreated(
	callback func(event *network.WebSocketCreatedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketCreated",
		func(response *Response) {
			event := &network.WebSocketCreatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketFrameError adds a handler to the Network.webSocketFrameError event.
Network.webSocketFrameError fires when a WebSocket frame error occurs.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameError
*/
func (protocol *NetworkProtocol) OnWebSocketFrameError(
	callback func(event *network.WebSocketFrameErrorEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketFrameError",
		func(response *Response) {
			event := &network.WebSocketFrameErrorEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketFrameReceived adds a handler to the Network.webSocketFrameReceived event.
Network.webSocketFrameReceived fires when WebSocket frame is received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameReceived
*/
func (protocol *NetworkProtocol) OnWebSocketFrameReceived(
	callback func(event *network.WebSocketFrameReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketFrameReceived",
		func(response *Response) {
			event := &network.WebSocketFrameReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketFrameSent adds a handler to the Network.webSocketFrameSent event.
Network.webSocketFrameSent fires when WebSocket frame is sent.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameSent
*/
func (protocol *NetworkProtocol) OnWebSocketFrameSent(
	callback func(event *network.WebSocketFrameSentEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketFrameSent",
		func(response *Response) {
			event := &network.WebSocketFrameSentEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketHandshakeResponseReceived adds a handler to the
Network.webSocketHandshakeResponseReceived event. Network.webSocketHandshakeResponseReceived fires
when WebSocket handshake response becomes available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketHandshakeResponseReceived
*/
func (protocol *NetworkProtocol) OnWebSocketHandshakeResponseReceived(
	callback func(event *network.WebSocketHandshakeResponseReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketHandshakeResponseReceived",
		func(response *Response) {
			event := &network.WebSocketHandshakeResponseReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketWillSendHandshakeRequest adds a handler to the
Network.webSocketWillSendHandshakeRequest event. Network.webSocketWillSendHandshakeRequest fires
when WebSocket is about to initiate handshake.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketWillSendHandshakeRequest
*/
func (protocol *NetworkProtocol) OnWebSocketWillSendHandshakeRequest(
	callback func(event *network.WebSocketWillSendHandshakeRequestEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketWillSendHandshakeRequest",
		func(response *Response) {
			event := &network.WebSocketWillSendHandshakeRequestEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
