package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Network - https://chromedevtools.github.io/devtools-protocol/tot/Network/
Allows tracking network activities of the page. It exposes information about http, file, data and
other requests and responses, their headers, bodies, timing, etc.
*/
type Network struct{}

/*
CanClearBrowserCache tells whether clearing browser cache is supported. DEPRECATED
*/
func (Network) CanClearBrowserCache(
	socket *Socket,
) (netowrk.CanClearBrowserCacheResult, error) {
	command := &protocol.Command{
		method: "Network.canClearBrowserCache",
	}
	socket.SendCommand(command)
	return netowrk.CanClearBrowserCacheResult, command.Err
}

/*
CanClearBrowserCookies tells whether clearing browser cookies is supported. DEPRECATED
*/
func (Network) CanClearBrowserCookies(
	socket *Socket,
) (netowrk.CanClearBrowserCookiesResult, error) {
	command := &protocol.Command{
		method: "Network.canClearBrowserCookies",
	}
	socket.SendCommand(command)
	return netowrk.CanClearBrowserCookiesResult, command.Err
}

/*
CanEmulateNetworkConditions tells whether emulation of network conditions is supported. DEPRECATED
*/
func (Network) CanEmulateNetworkConditions(
	socket *Socket,
) (netowrk.CanEmulateNetworkConditionsResult, error) {
	command := &protocol.Command{
		method: "Network.canEmulateNetworkConditions",
	}
	socket.SendCommand(command)
	return netowrk.CanEmulateNetworkConditionsResult, command.Err
}

/*
ClearBrowserCache clears browser cache.
*/
func (Network) ClearBrowserCache(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.clearBrowserCache",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
ClearBrowserCookies clears browser cookies.
*/
func (Network) ClearBrowserCookies(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.clearBrowserCookies",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
ContinueInterceptedRequest response to Network.requestIntercepted which either modifies the request
to continue with any modifications, or blocks it, or completes it with the provided response bytes.
If a network fetch occurs as a result which encounters a redirect an additional
Network.requestIntercepted event will be sent with the same InterceptionID. EXPERIMENTAL
*/
func (Network) ContinueInterceptedRequest(
	socket *Socket,
	params *netowrk.ContinueInterceptedRequestParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.continueInterceptedRequest",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
DeleteCookies deletes browser cookies with matching name and url or domain/path pair.
*/
func (Network) DeleteCookies(
	socket *Socket,
	params *netowrk.DeleteCookiesParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.deleteCookies",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Disable disables network tracking, prevents network events from being sent to the client.
*/
func (Network) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
EmulateNetworkConditions activates emulation of network conditions.
*/
func (Network) EmulateNetworkConditions(
	socket *Socket,
	params *netowrk.EmulateNetworkConditionsParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.emulateNetworkConditions",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables network tracking, network events will now be delivered to the client.
*/
func (Network) Enable(
	socket *Socket,
	params *netowrk.EnableParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.enable",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetAllCookies returns all browser cookies. Depending on the backend support, will return detailed
cookie information in the `cookies` field.
*/
func (Network) GetAllCookies(
	socket *Socket,
) (netowrk.GetAllCookiesResult, error) {
	command := &protocol.Command{
		method: "Network.getAllCookies",
	}
	socket.SendCommand(command)
	return netowrk.GetAllCookiesResult, command.Err
}

/*
GetCertificate returns the DER-encoded certificate. EXPERIMENTAL
*/
func (Network) GetCertificate(
	socket *Socket,
	params *netowrk.GetCertificateParams,
) (netowrk.GetCertificateResult, error) {
	command := &protocol.Command{
		method: "Network.getCertificate",
		params: params,
	}
	socket.SendCommand(command)
	return netowrk.GetCertificateResult, command.Err
}

/*
GetCookies returns all browser cookies for the current URL. Depending on the backend support, will
return detailed cookie information in the `cookies` field.
*/
func (Network) GetCookies(
	socket *Socket,
	params *netowrk.GetCookiesParams,
) (netowrk.GetCookiesResult, error) {
	command := &protocol.Command{
		method: "Network.getCookies",
		params: params,
	}
	socket.SendCommand(command)
	return netowrk.GetCookiesResult, command.Err
}

/*
GetResponseBody returns content served for the given request.
*/
func (Network) GetResponseBody(
	socket *Socket,
	params *netowrk.GetResponseBodyParams,
) (netowrk.GetResponseBodyResult, error) {
	command := &protocol.Command{
		method: "Network.getResponseBody",
		params: params,
	}
	socket.SendCommand(command)
	return netowrk.GetResponseBodyResult, command.Err
}

/*
GetResponseBodyForInterception returns content served for the given currently intercepted request.
EXPERIMENTAL
*/
func (Network) GetResponseBodyForInterception(
	socket *Socket,
	params *netowrk.GetResponseBodyForInterceptionParams,
) (netowrk.GetResponseBodyForInterceptionResult, error) {
	command := &protocol.Command{
		method: "Network.getResponseBodyForInterception",
		params: params,
	}
	socket.SendCommand(command)
	return netowrk.GetResponseBodyForInterceptionResult, command.Err
}

/*
ReplayXHR sends a new XMLHttpRequest which is identical to the original one. The following
parameters should be identical: method, url, async, request body, extra headers, withCredentials
attribute, user, password. EXPERIMENTAL
*/
func (Network) ReplayXHR(
	socket *Socket,
	params *netowrk.ReplayXHRParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.replayXHR",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SearchInResponseBody searches for given string in response content. EXPERIMENTAL
*/
func (Network) SearchInResponseBody(
	socket *Socket,
	params *netowrk.SearchInResponseBodyParams,
) (netowrk.SearchInResponseBodyResult, error) {
	command := &protocol.Command{
		method: "Network.searchInResponseBody",
		params: params,
	}
	socket.SendCommand(command)
	return netowrk.SearchInResponseBodyResult, command.Err
}

/*
SetBlockedURLs blocks URLs from loading. EXPERIMENTAL
*/
func (Network) SetBlockedURLs(
	socket *Socket,
	params *netowrk.SetBlockedURLsParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setBlockedURLs",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetBypassServiceWorker toggles ignoring of service worker for each request. EXPERIMENTAL
*/
func (Network) SetBypassServiceWorker(
	socket *Socket,
	params *netowrk.SetBypassServiceWorkerParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setBypassServiceWorker",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetCacheDisabled toggles ignoring cache for each request. If `true`, cache will not be used.
*/
func (Network) SetCacheDisabled(
	socket *Socket,
	params *netowrk.SetCacheDisabledParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setCacheDisabled",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetCookie sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
*/
func (Network) SetCookie(
	socket *Socket,
	params *netowrk.SetCookieParams,
) (netowrk.SetCookieResult, error) {
	command := &protocol.Command{
		method: "Network.setCookie",
		params: params,
	}
	socket.SendCommand(command)
	return netowrk.SetCookieResult, command.Err
}

/*
SetCookies sets given cookies.
*/
func (Network) SetCookies(
	socket *Socket,
	params *netowrk.SetCookiesParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setCookies",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetDataSizeLimitsForTest is for testing. EXPERIMENTAL
*/
func (Network) SetDataSizeLimitsForTest(
	socket *Socket,
	params *netowrk.SetDataSizeLimitsForTestParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setDataSizeLimitsForTest",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetExtraHTTPHeaders specifies whether to always send extra HTTP headers with the requests from this
page.
*/
func (Network) SetExtraHTTPHeaders(
	socket *Socket,
	params *netowrk.SetExtraHTTPHeadersParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setExtraHTTPHeaders",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetRequestInterception sets the requests to intercept that match a the provided patterns and
optionally resource types. EXPERIMENTAL
*/
func (Network) SetRequestInterception(
	socket *Socket,
	params *netowrk.SetRequestInterceptionParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setRequestInterception",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetUserAgentOverride allows overriding user agent with the given string.
*/
func (Network) SetUserAgentOverride(
	socket *Socket,
	params *netowrk.SetUserAgentOverrideParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Network.setUserAgentOverride",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
OnDataReceived adds a handler to the Network.dataReceived event. Network.dataReceived fires when a
data chunk was received over the network.
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
OnRequestWillBeSent adds a handler to the Network.requestWillBeSent event.
Network.requestWillBeSent fires when the page is about to send HTTP request.
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
OnWebSocketFrameReceived adds a handler to the Network.webSocketFrameReceived event.
Network.webSocketFrameReceived fires when WebSocket frame is received.
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
