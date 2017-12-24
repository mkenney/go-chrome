package chrome

import (
	"encoding/json"

	network "github.com/mkenney/go-chrome/network"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/sirupsen/logrus"
)

/*
Network - https://chromedevtools.github.io/devtools-protocol/tot/Network/
Allows tracking network activities of the page. It exposes information about http, file, data and
other requests and responses, their headers, bodies, timing, etc.
*/
type Network struct{}

/*
CanClearBrowserCache tells whether clearing browser cache is supported. DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCache
*/
func (Network) CanClearBrowserCache(
	socket *Socket,
) (network.CanClearBrowserCacheResult, error) {
	command := &protocol.Command{
		Method: "Network.canClearBrowserCache",
	}
	result := network.CanClearBrowserCacheResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
CanClearBrowserCookies tells whether clearing browser cookies is supported. DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCookies
*/
func (Network) CanClearBrowserCookies(
	socket *Socket,
) (network.CanClearBrowserCookiesResult, error) {
	command := &protocol.Command{
		Method: "Network.canClearBrowserCookies",
	}
	result := network.CanClearBrowserCookiesResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
CanEmulateConditions tells whether emulation of network conditions is supported. DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canEmulateNetworkConditions
*/
func (Network) CanEmulateConditions(
	socket *Socket,
) (network.CanEmulateConditionsResult, error) {
	command := &protocol.Command{
		Method: "Network.canEmulateNetworkConditions",
	}
	result := network.CanEmulateConditionsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
ClearBrowserCache clears browser cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache
*/
func (Network) ClearBrowserCache(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Network.clearBrowserCache",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ClearBrowserCookies clears browser cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCookies
*/
func (Network) ClearBrowserCookies(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Network.clearBrowserCookies",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ContinueInterceptedRequest response to Network.requestIntercepted which either modifies the request
to continue with any modifications, or blocks it, or completes it with the provided response bytes.
If a network fetch occurs as a result which encounters a redirect an additional
Network.requestIntercepted event will be sent with the same InterceptionID. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-continueInterceptedRequest
*/
func (Network) ContinueInterceptedRequest(
	socket *Socket,
	params *network.ContinueInterceptedRequestParams,
) error {
	command := &protocol.Command{
		Method: "Network.continueInterceptedRequest",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteCookies deletes browser cookies with matching name and url or domain/path pair.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
*/
func (Network) DeleteCookies(
	socket *Socket,
	params *network.DeleteCookiesParams,
) error {
	command := &protocol.Command{
		Method: "Network.deleteCookies",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables network tracking, prevents network events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-disable
*/
func (Network) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Network.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
EmulateConditions activates emulation of network conditions.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions
*/
func (Network) EmulateConditions(
	socket *Socket,
	params *network.EmulateConditionsParams,
) error {
	command := &protocol.Command{
		Method: "Network.emulateNetworkConditions",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables network tracking, network events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
*/
func (Network) Enable(
	socket *Socket,
	params *network.EnableParams,
) error {
	command := &protocol.Command{
		Method: "Network.enable",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetAllCookies returns all browser cookies. Depending on the backend support, will return detailed
cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getAllCookies
*/
func (Network) GetAllCookies(
	socket *Socket,
) (network.GetAllCookiesResult, error) {
	command := &protocol.Command{
		Method: "Network.getAllCookies",
	}
	result := network.GetAllCookiesResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetCertificate returns the DER-encoded certificate. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate
*/
func (Network) GetCertificate(
	socket *Socket,
	params *network.GetCertificateParams,
) (network.GetCertificateResult, error) {
	command := &protocol.Command{
		Method: "Network.getCertificate",
		Params: params,
	}
	result := network.GetCertificateResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetCookies returns all browser cookies for the current URL. Depending on the backend support, will
return detailed cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
*/
func (Network) GetCookies(
	socket *Socket,
	params *network.GetCookiesParams,
) (network.GetCookiesResult, error) {
	command := &protocol.Command{
		Method: "Network.getCookies",
		Params: params,
	}
	result := network.GetCookiesResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetResponseBody returns content served for the given request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
*/
func (Network) GetResponseBody(
	socket *Socket,
	params *network.GetResponseBodyParams,
) (network.GetResponseBodyResult, error) {
	command := &protocol.Command{
		Method: "Network.getResponseBody",
		Params: params,
	}
	result := network.GetResponseBodyResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetResponseBodyForInterception returns content served for the given currently intercepted request.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
*/
func (Network) GetResponseBodyForInterception(
	socket *Socket,
	params *network.GetResponseBodyForInterceptionParams,
) (network.GetResponseBodyForInterceptionResult, error) {
	command := &protocol.Command{
		Method: "Network.getResponseBodyForInterception",
		Params: params,
	}
	result := network.GetResponseBodyForInterceptionResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
ReplayXHR sends a new XMLHttpRequest which is identical to the original one. The following
parameters should be identical: method, url, async, request body, extra headers, withCredentials
attribute, user, password. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR
*/
func (Network) ReplayXHR(
	socket *Socket,
	params *network.ReplayXHRParams,
) error {
	command := &protocol.Command{
		Method: "Network.replayXHR",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SearchInResponseBody searches for given string in response content. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
*/
func (Network) SearchInResponseBody(
	socket *Socket,
	params *network.SearchInResponseBodyParams,
) (network.SearchInResponseBodyResult, error) {
	command := &protocol.Command{
		Method: "Network.searchInResponseBody",
		Params: params,
	}
	result := network.SearchInResponseBodyResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetBlockedURLs blocks URLs from loading. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs
*/
func (Network) SetBlockedURLs(
	socket *Socket,
	params *network.SetBlockedURLsParams,
) error {
	command := &protocol.Command{
		Method: "Network.setBlockedURLs",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetBypassServiceWorker toggles ignoring of service worker for each request. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
*/
func (Network) SetBypassServiceWorker(
	socket *Socket,
	params *network.SetBypassServiceWorkerParams,
) error {
	command := &protocol.Command{
		Method: "Network.setBypassServiceWorker",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetCacheDisabled toggles ignoring cache for each request. If `true`, cache will not be used.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
*/
func (Network) SetCacheDisabled(
	socket *Socket,
	params *network.SetCacheDisabledParams,
) error {
	command := &protocol.Command{
		Method: "Network.setCacheDisabled",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetCookie sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
*/
func (Network) SetCookie(
	socket *Socket,
	params *network.SetCookieParams,
) (network.SetCookieResult, error) {
	command := &protocol.Command{
		Method: "Network.setCookie",
		Params: params,
	}
	result := network.SetCookieResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetCookies sets given cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
*/
func (Network) SetCookies(
	socket *Socket,
	params *network.SetCookiesParams,
) error {
	command := &protocol.Command{
		Method: "Network.setCookies",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDataSizeLimitsForTest is for testing. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
*/
func (Network) SetDataSizeLimitsForTest(
	socket *Socket,
	params *network.SetDataSizeLimitsForTestParams,
) error {
	command := &protocol.Command{
		Method: "Network.setDataSizeLimitsForTest",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetExtraHTTPHeaders specifies whether to always send extra HTTP headers with the requests from this
page.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
*/
func (Network) SetExtraHTTPHeaders(
	socket *Socket,
	params *network.SetExtraHTTPHeadersParams,
) error {
	command := &protocol.Command{
		Method: "Network.setExtraHTTPHeaders",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetRequestInterception sets the requests to intercept that match a the provided patterns and
optionally resource types. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
*/
func (Network) SetRequestInterception(
	socket *Socket,
	params *network.SetRequestInterceptionParams,
) error {
	command := &protocol.Command{
		Method: "Network.setRequestInterception",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetUserAgentOverride allows overriding user agent with the given string.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setUserAgentOverride
*/
func (Network) SetUserAgentOverride(
	socket *Socket,
	params *network.SetUserAgentOverrideParams,
) error {
	command := &protocol.Command{
		Method: "Network.setUserAgentOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnDataReceived adds a handler to the Network.dataReceived event. Network.dataReceived fires when a
data chunk was received over the network.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-dataReceived
*/
func (Network) OnDataReceived(
	socket *Socket,
	callback func(event *network.DataReceivedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.dataReceived",
		func(name string, params []byte) {
			event := &network.DataReceivedEvent{}
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
OnEventSourceMessageReceived adds a handler to the Network.eventSourceMessageReceived event.
Network.eventSourceMessageReceived fires when EventSource message is received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-eventSourceMessageReceived
*/
func (Network) OnEventSourceMessageReceived(
	socket *Socket,
	callback func(event *network.EventSourceMessageReceivedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.eventSourceMessageReceived",
		func(name string, params []byte) {
			event := &network.EventSourceMessageReceivedEvent{}
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
OnLoadingFailed adds a handler to the Network.loadingFailed event. Network.loadingFailed fires when
HTTP request has failed to load.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFailed
*/
func (Network) OnLoadingFailed(
	socket *Socket,
	callback func(event *network.LoadingFailedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.loadingFailed",
		func(name string, params []byte) {
			event := &network.LoadingFailedEvent{}
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
OnLoadingFinished adds a handler to the Network.loadingFinished event. Network.loadingFinished fires
when HTTP request has finished loading.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFinished
*/
func (Network) OnLoadingFinished(
	socket *Socket,
	callback func(event *network.LoadingFinishedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.loadingFinished",
		func(name string, params []byte) {
			event := &network.LoadingFinishedEvent{}
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
OnRequestIntercepted adds a handler to the Network.requestIntercepted event.
Network.requestIntercepted fires when a HTTP request is intercepted and returns details, which must
be either allowed, blocked, modified or mocked. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestIntercepted
*/
func (Network) OnRequestIntercepted(
	socket *Socket,
	callback func(event *network.RequestInterceptedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.requestIntercepted",
		func(name string, params []byte) {
			event := &network.RequestInterceptedEvent{}
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
OnRequestServedFromCache adds a handler to the Network.requestServedFromCache event.
Network.requestServedFromCache fires when request ended up loading from cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestServedFromCache
*/
func (Network) OnRequestServedFromCache(
	socket *Socket,
	callback func(event *network.RequestServedFromCacheEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.requestServedFromCache",
		func(name string, params []byte) {
			event := &network.RequestServedFromCacheEvent{}
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
OnRequestWillBeSent adds a handler to the Network.requestWillBeSent event. Network.requestWillBeSent
fires when the page is about to send HTTP request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestWillBeSent
*/
func (Network) OnRequestWillBeSent(
	socket *Socket,
	callback func(event *network.RequestWillBeSentEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.requestWillBeSent",
		func(name string, params []byte) {
			event := &network.RequestWillBeSentEvent{}
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
OnResourceChangedPriority adds a handler to the Network.resourceChangedPriority event.
Network.resourceChangedPriority fires when resource loading priority is changed EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-resourceChangedPriority
*/
func (Network) OnResourceChangedPriority(
	socket *Socket,
	callback func(event *network.ResourceChangedPriorityEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.resourceChangedPriority",
		func(name string, params []byte) {
			event := &network.ResourceChangedPriorityEvent{}
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
OnResponseReceived adds a handler to the Network.responseReceived event. Network.responseReceived
fires when HTTP response is available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-responseReceived
*/
func (Network) OnResponseReceived(
	socket *Socket,
	callback func(event *network.ResponseReceivedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.responseReceived",
		func(name string, params []byte) {
			event := &network.ResponseReceivedEvent{}
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
OnWebSocketClosed adds a handler to the Network.webSocketClosed event. Network.webSocketClosed
fires when WebSocket is closed.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketClosed
*/
func (Network) OnWebSocketClosed(
	socket *Socket,
	callback func(event *network.WebSocketClosedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.webSocketClosed",
		func(name string, params []byte) {
			event := &network.WebSocketClosedEvent{}
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
OnWebSocketCreated adds a handler to the Network.webSocketCreated event. Network.webSocketCreated
fires upon WebSocket creation.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketCreated
*/
func (Network) OnWebSocketCreated(
	socket *Socket,
	callback func(event *network.WebSocketCreatedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.webSocketCreated",
		func(name string, params []byte) {
			event := &network.WebSocketCreatedEvent{}
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
OnWebSocketFrameError adds a handler to the Network.webSocketFrameError event.
Network.webSocketFrameError fires when a WebSocket frame error occurs.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameError
*/
func (Network) OnWebSocketFrameError(
	socket *Socket,
	callback func(event *network.WebSocketFrameErrorEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.webSocketFrameError",
		func(name string, params []byte) {
			event := &network.WebSocketFrameErrorEvent{}
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
OnWebSocketFrameReceived adds a handler to the Network.webSocketFrameReceived event.
Network.webSocketFrameReceived fires when WebSocket frame is received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameReceived
*/
func (Network) OnWebSocketFrameReceived(
	socket *Socket,
	callback func(event *network.WebSocketFrameReceivedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.webSocketFrameReceived",
		func(name string, params []byte) {
			event := &network.WebSocketFrameReceivedEvent{}
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
OnWebSocketFrameSent adds a handler to the Network.webSocketFrameSent event.
Network.webSocketFrameSent fires when WebSocket frame is sent.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameSent
*/
func (Network) OnWebSocketFrameSent(
	socket *Socket,
	callback func(event *network.WebSocketFrameSentEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.webSocketFrameSent",
		func(name string, params []byte) {
			event := &network.WebSocketFrameSentEvent{}
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
OnWebSocketHandshakeResponseReceived adds a handler to the
Network.webSocketHandshakeResponseReceived event. Network.webSocketHandshakeResponseReceived fires
when WebSocket handshake response becomes available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketHandshakeResponseReceived
*/
func (Network) OnWebSocketHandshakeResponseReceived(
	socket *Socket,
	callback func(event *network.WebSocketHandshakeResponseReceivedEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.webSocketHandshakeResponseReceived",
		func(name string, params []byte) {
			event := &network.WebSocketHandshakeResponseReceivedEvent{}
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
OnWebSocketWillSendHandshakeRequest adds a handler to the
Network.webSocketWillSendHandshakeRequest event. Network.webSocketWillSendHandshakeRequest fires
when WebSocket is about to initiate handshake.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketWillSendHandshakeRequest
*/
func (Network) OnWebSocketWillSendHandshakeRequest(
	socket *Socket,
	callback func(event *network.WebSocketWillSendHandshakeRequestEvent),
) {
	handler := protocol.NewEventHandler(
		"Network.webSocketWillSendHandshakeRequest",
		func(name string, params []byte) {
			event := &network.WebSocketWillSendHandshakeRequestEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
