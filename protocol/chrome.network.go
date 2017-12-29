package protocol

import (
	"encoding/json"

	network "github.com/mkenney/go-chrome/protocol/network"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Network provides a namespace for the Chrome Network protocol methods. The Network protocol allows
tracking network activities of the page. It exposes information about http, file, data and other
requests and responses, their headers, bodies, timing, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Network/
*/
var Network = NetworkProtocol{}

/*
NetworkProtocol defines the Network protocol methods.
*/
type NetworkProtocol struct{}

/*
CanClearBrowserCache tells whether clearing browser cache is supported. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCache
*/
func (NetworkProtocol) CanClearBrowserCache(
	socket sock.Socketer,
) (*network.CanClearBrowserCacheResult, error) {
	command := sock.NewCommand("Network.canClearBrowserCache", nil)
	result := &network.CanClearBrowserCacheResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
CanClearBrowserCookies tells whether clearing browser cookies is supported. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCookies
*/
func (NetworkProtocol) CanClearBrowserCookies(
	socket sock.Socketer,
) (*network.CanClearBrowserCookiesResult, error) {
	command := sock.NewCommand("Network.canClearBrowserCookies", nil)
	result := &network.CanClearBrowserCookiesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
CanEmulateConditions tells whether emulation of network conditions is supported. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canEmulateNetworkConditions
*/
func (NetworkProtocol) CanEmulateConditions(
	socket sock.Socketer,
) (*network.CanEmulateConditionsResult, error) {
	command := sock.NewCommand("Network.canEmulateNetworkConditions", nil)
	result := &network.CanEmulateConditionsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
ClearBrowserCache clears browser cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache
*/
func (NetworkProtocol) ClearBrowserCache(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Network.clearBrowserCache", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ClearBrowserCookies clears browser cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCookies
*/
func (NetworkProtocol) ClearBrowserCookies(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Network.clearBrowserCookies", nil)
	socket.SendCommand(command)
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
func (NetworkProtocol) ContinueInterceptedRequest(
	socket sock.Socketer,
	params *network.ContinueInterceptedRequestParams,
) error {
	command := sock.NewCommand("Network.continueInterceptedRequest", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DeleteCookies deletes browser cookies with matching name and url or domain/path pair.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
*/
func (NetworkProtocol) DeleteCookies(
	socket sock.Socketer,
	params *network.DeleteCookiesParams,
) error {
	command := sock.NewCommand("Network.deleteCookies", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables network tracking, prevents network events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-disable
*/
func (NetworkProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Network.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
EmulateConditions activates emulation of network conditions.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions
*/
func (NetworkProtocol) EmulateConditions(
	socket sock.Socketer,
	params *network.EmulateConditionsParams,
) error {
	command := sock.NewCommand("Network.emulateNetworkConditions", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables network tracking, network events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
*/
func (NetworkProtocol) Enable(
	socket sock.Socketer,
	params *network.EnableParams,
) error {
	command := sock.NewCommand("Network.enable", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetAllCookies returns all browser cookies. Depending on the backend support, will return detailed
cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getAllCookies
*/
func (NetworkProtocol) GetAllCookies(
	socket sock.Socketer,
) (*network.GetAllCookiesResult, error) {
	command := sock.NewCommand("Network.getAllCookies", nil)
	result := &network.GetAllCookiesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetCertificate returns the DER-encoded certificate.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate EXPERIMENTAL.
*/
func (NetworkProtocol) GetCertificate(
	socket sock.Socketer,
	params *network.GetCertificateParams,
) (*network.GetCertificateResult, error) {
	command := sock.NewCommand("Network.getCertificate", params)
	result := &network.GetCertificateResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetCookies returns all browser cookies for the current URL. Depending on the backend support, will
return detailed cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
*/
func (NetworkProtocol) GetCookies(
	socket sock.Socketer,
	params *network.GetCookiesParams,
) (*network.GetCookiesResult, error) {
	command := sock.NewCommand("Network.getCookies", params)
	result := &network.GetCookiesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetResponseBody returns content served for the given request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
*/
func (NetworkProtocol) GetResponseBody(
	socket sock.Socketer,
	params *network.GetResponseBodyParams,
) (*network.GetResponseBodyResult, error) {
	command := sock.NewCommand("Network.getResponseBody", params)
	result := &network.GetResponseBodyResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetResponseBodyForInterception returns content served for the given currently intercepted request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
EXPERIMENTAL.
*/
func (NetworkProtocol) GetResponseBodyForInterception(
	socket sock.Socketer,
	params *network.GetResponseBodyForInterceptionParams,
) (*network.GetResponseBodyForInterceptionResult, error) {
	command := sock.NewCommand("Network.getResponseBodyForInterception", params)
	result := &network.GetResponseBodyForInterceptionResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
ReplayXHR sends a new XMLHttpRequest which is identical to the original one. The following
parameters should be identical: method, url, async, request body, extra headers, withCredentials
attribute, user, password.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR EXPERIMENTAL.
*/
func (NetworkProtocol) ReplayXHR(
	socket sock.Socketer,
	params *network.ReplayXHRParams,
) error {
	command := sock.NewCommand("Network.replayXHR", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SearchInResponseBody searches for given string in response content.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
EXPERIMENTAL.
*/
func (NetworkProtocol) SearchInResponseBody(
	socket sock.Socketer,
	params *network.SearchInResponseBodyParams,
) (*network.SearchInResponseBodyResult, error) {
	command := sock.NewCommand("Network.searchInResponseBody", params)
	result := &network.SearchInResponseBodyResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
SetBlockedURLs blocks URLs from loading.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs EXPERIMENTAL.
*/
func (NetworkProtocol) SetBlockedURLs(
	socket sock.Socketer,
	params *network.SetBlockedURLsParams,
) error {
	command := sock.NewCommand("Network.setBlockedURLs", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetBypassServiceWorker toggles ignoring of service worker for each request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
EXPERIMENTAL.
*/
func (NetworkProtocol) SetBypassServiceWorker(
	socket sock.Socketer,
	params *network.SetBypassServiceWorkerParams,
) error {
	command := sock.NewCommand("Network.setBypassServiceWorker", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetCacheDisabled toggles ignoring cache for each request. If `true`, cache will not be used.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
*/
func (NetworkProtocol) SetCacheDisabled(
	socket sock.Socketer,
	params *network.SetCacheDisabledParams,
) error {
	command := sock.NewCommand("Network.setCacheDisabled", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetCookie sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
*/
func (NetworkProtocol) SetCookie(
	socket sock.Socketer,
	params *network.SetCookieParams,
) (*network.SetCookieResult, error) {
	command := sock.NewCommand("Network.setCookie", params)
	result := &network.SetCookieResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
SetCookies sets given cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
*/
func (NetworkProtocol) SetCookies(
	socket sock.Socketer,
	params *network.SetCookiesParams,
) error {
	command := sock.NewCommand("Network.setCookies", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetDataSizeLimitsForTest is for testing.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
EXPERIMENTAL.
*/
func (NetworkProtocol) SetDataSizeLimitsForTest(
	socket sock.Socketer,
	params *network.SetDataSizeLimitsForTestParams,
) error {
	command := sock.NewCommand("Network.setDataSizeLimitsForTest", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetExtraHTTPHeaders specifies whether to always send extra HTTP headers with the requests from this
page.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
*/
func (NetworkProtocol) SetExtraHTTPHeaders(
	socket sock.Socketer,
	params *network.SetExtraHTTPHeadersParams,
) error {
	command := sock.NewCommand("Network.setExtraHTTPHeaders", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetRequestInterception sets the requests to intercept that match a the provided patterns and
optionally resource types.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
EXPERIMENTAL.
*/
func (NetworkProtocol) SetRequestInterception(
	socket sock.Socketer,
	params *network.SetRequestInterceptionParams,
) error {
	command := sock.NewCommand("Network.setRequestInterception", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetUserAgentOverride allows overriding user agent with the given string.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setUserAgentOverride
*/
func (NetworkProtocol) SetUserAgentOverride(
	socket sock.Socketer,
	params *network.SetUserAgentOverrideParams,
) error {
	command := sock.NewCommand("Network.setUserAgentOverride", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnDataReceived adds a handler to the Network.dataReceived event. Network.dataReceived fires when a
data chunk was received over the network.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-dataReceived
*/
func (NetworkProtocol) OnDataReceived(
	socket sock.Socketer,
	callback func(event *network.DataReceivedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.dataReceived",
		func(response *sock.Response) {
			event := &network.DataReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnEventSourceMessageReceived adds a handler to the Network.eventSourceMessageReceived event.
Network.eventSourceMessageReceived fires when EventSource message is received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-eventSourceMessageReceived
*/
func (NetworkProtocol) OnEventSourceMessageReceived(
	socket sock.Socketer,
	callback func(event *network.EventSourceMessageReceivedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.eventSourceMessageReceived",
		func(response *sock.Response) {
			event := &network.EventSourceMessageReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLoadingFailed adds a handler to the Network.loadingFailed event. Network.loadingFailed fires when
HTTP request has failed to load.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFailed
*/
func (NetworkProtocol) OnLoadingFailed(
	socket sock.Socketer,
	callback func(event *network.LoadingFailedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.loadingFailed",
		func(response *sock.Response) {
			event := &network.LoadingFailedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLoadingFinished adds a handler to the Network.loadingFinished event. Network.loadingFinished fires
when HTTP request has finished loading.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFinished
*/
func (NetworkProtocol) OnLoadingFinished(
	socket sock.Socketer,
	callback func(event *network.LoadingFinishedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.loadingFinished",
		func(response *sock.Response) {
			event := &network.LoadingFinishedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnRequestIntercepted adds a handler to the Network.requestIntercepted event.
Network.requestIntercepted fires when a HTTP request is intercepted and returns details, which must
be either allowed, blocked, modified or mocked.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestIntercepted
EXPERIMENTAL.
*/
func (NetworkProtocol) OnRequestIntercepted(
	socket sock.Socketer,
	callback func(event *network.RequestInterceptedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.requestIntercepted",
		func(response *sock.Response) {
			event := &network.RequestInterceptedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnRequestServedFromCache adds a handler to the Network.requestServedFromCache event.
Network.requestServedFromCache fires when request ended up loading from cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestServedFromCache
*/
func (NetworkProtocol) OnRequestServedFromCache(
	socket sock.Socketer,
	callback func(event *network.RequestServedFromCacheEvent),
) {
	handler := sock.NewEventHandler(
		"Network.requestServedFromCache",
		func(response *sock.Response) {
			event := &network.RequestServedFromCacheEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnRequestWillBeSent adds a handler to the Network.requestWillBeSent event. Network.requestWillBeSent
fires when the page is about to send HTTP request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestWillBeSent
*/
func (NetworkProtocol) OnRequestWillBeSent(
	socket sock.Socketer,
	callback func(event *network.RequestWillBeSentEvent),
) {
	handler := sock.NewEventHandler(
		"Network.requestWillBeSent",
		func(response *sock.Response) {
			event := &network.RequestWillBeSentEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnResourceChangedPriority adds a handler to the Network.resourceChangedPriority event.
Network.resourceChangedPriority fires when resource loading priority is changed.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-resourceChangedPriority
EXPERIMENTAL.
*/
func (NetworkProtocol) OnResourceChangedPriority(
	socket sock.Socketer,
	callback func(event *network.ResourceChangedPriorityEvent),
) {
	handler := sock.NewEventHandler(
		"Network.resourceChangedPriority",
		func(response *sock.Response) {
			event := &network.ResourceChangedPriorityEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnResponseReceived adds a handler to the Network.responseReceived event. Network.responseReceived
fires when HTTP response is available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-responseReceived
*/
func (NetworkProtocol) OnResponseReceived(
	socket sock.Socketer,
	callback func(event *network.ResponseReceivedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.responseReceived",
		func(response *sock.Response) {
			event := &network.ResponseReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWebSocketClosed adds a handler to the Network.webSocketClosed event. Network.webSocketClosed
fires when WebSocket is closed.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketClosed
*/
func (NetworkProtocol) OnWebSocketClosed(
	socket sock.Socketer,
	callback func(event *network.WebSocketClosedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.webSocketClosed",
		func(response *sock.Response) {
			event := &network.WebSocketClosedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWebSocketCreated adds a handler to the Network.webSocketCreated event. Network.webSocketCreated
fires upon WebSocket creation.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketCreated
*/
func (NetworkProtocol) OnWebSocketCreated(
	socket sock.Socketer,
	callback func(event *network.WebSocketCreatedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.webSocketCreated",
		func(response *sock.Response) {
			event := &network.WebSocketCreatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWebSocketFrameError adds a handler to the Network.webSocketFrameError event.
Network.webSocketFrameError fires when a WebSocket frame error occurs.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameError
*/
func (NetworkProtocol) OnWebSocketFrameError(
	socket sock.Socketer,
	callback func(event *network.WebSocketFrameErrorEvent),
) {
	handler := sock.NewEventHandler(
		"Network.webSocketFrameError",
		func(response *sock.Response) {
			event := &network.WebSocketFrameErrorEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWebSocketFrameReceived adds a handler to the Network.webSocketFrameReceived event.
Network.webSocketFrameReceived fires when WebSocket frame is received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameReceived
*/
func (NetworkProtocol) OnWebSocketFrameReceived(
	socket sock.Socketer,
	callback func(event *network.WebSocketFrameReceivedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.webSocketFrameReceived",
		func(response *sock.Response) {
			event := &network.WebSocketFrameReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWebSocketFrameSent adds a handler to the Network.webSocketFrameSent event.
Network.webSocketFrameSent fires when WebSocket frame is sent.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameSent
*/
func (NetworkProtocol) OnWebSocketFrameSent(
	socket sock.Socketer,
	callback func(event *network.WebSocketFrameSentEvent),
) {
	handler := sock.NewEventHandler(
		"Network.webSocketFrameSent",
		func(response *sock.Response) {
			event := &network.WebSocketFrameSentEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWebSocketHandshakeResponseReceived adds a handler to the
Network.webSocketHandshakeResponseReceived event. Network.webSocketHandshakeResponseReceived fires
when WebSocket handshake response becomes available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketHandshakeResponseReceived
*/
func (NetworkProtocol) OnWebSocketHandshakeResponseReceived(
	socket sock.Socketer,
	callback func(event *network.WebSocketHandshakeResponseReceivedEvent),
) {
	handler := sock.NewEventHandler(
		"Network.webSocketHandshakeResponseReceived",
		func(response *sock.Response) {
			event := &network.WebSocketHandshakeResponseReceivedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnWebSocketWillSendHandshakeRequest adds a handler to the
Network.webSocketWillSendHandshakeRequest event. Network.webSocketWillSendHandshakeRequest fires
when WebSocket is about to initiate handshake.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketWillSendHandshakeRequest
*/
func (NetworkProtocol) OnWebSocketWillSendHandshakeRequest(
	socket sock.Socketer,
	callback func(event *network.WebSocketWillSendHandshakeRequestEvent),
) {
	handler := sock.NewEventHandler(
		"Network.webSocketWillSendHandshakeRequest",
		func(response *sock.Response) {
			event := &network.WebSocketWillSendHandshakeRequestEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
