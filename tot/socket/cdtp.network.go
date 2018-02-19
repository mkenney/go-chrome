package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

/*
NetworkProtocol provides a namespace for the Chrome Network protocol methods.
The Network protocol allows tracking network activities of the page. It exposes
information about http, file, data and other requests and responses, their
headers, bodies, timing, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Network/
*/
type NetworkProtocol struct {
	Socket Socketer
}

/*
CanClearBrowserCache tells whether clearing browser cache is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCache
DEPRECATED.
*/
func (protocol *NetworkProtocol) CanClearBrowserCache() <-chan *network.CanClearBrowserCacheResult {
	resultChan := make(chan *network.CanClearBrowserCacheResult)
	command := NewCommand(protocol.Socket, "Network.canClearBrowserCache", nil)
	result := &network.CanClearBrowserCacheResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
CanClearBrowserCookies tells whether clearing browser cookies is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCookies
DEPRECATED.
*/
func (protocol *NetworkProtocol) CanClearBrowserCookies() <-chan *network.CanClearBrowserCookiesResult {
	resultChan := make(chan *network.CanClearBrowserCookiesResult)
	command := NewCommand(protocol.Socket, "Network.canClearBrowserCookies", nil)
	result := &network.CanClearBrowserCookiesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
CanEmulateConditions tells whether emulation of network conditions is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canEmulateNetworkConditions
DEPRECATED.
*/
func (protocol *NetworkProtocol) CanEmulateConditions() <-chan *network.CanEmulateConditionsResult {
	resultChan := make(chan *network.CanEmulateConditionsResult)
	command := NewCommand(protocol.Socket, "Network.canEmulateNetworkConditions", nil)
	result := &network.CanEmulateConditionsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
ClearBrowserCache clears browser cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache
*/
func (protocol *NetworkProtocol) ClearBrowserCache() <-chan *network.ClearBrowserCacheResult {
	resultChan := make(chan *network.ClearBrowserCacheResult)
	command := NewCommand(protocol.Socket, "Network.clearBrowserCache", nil)
	result := &network.ClearBrowserCacheResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
ClearBrowserCookies clears browser cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCookies
*/
func (protocol *NetworkProtocol) ClearBrowserCookies() <-chan *network.ClearBrowserCookiesResult {
	resultChan := make(chan *network.ClearBrowserCookiesResult)
	result := &network.ClearBrowserCookiesResult{}
	command := NewCommand(protocol.Socket, "Network.clearBrowserCookies", nil)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
ContinueInterceptedRequest response to Network.requestIntercepted which either
modifies the request to continue with any modifications, or blocks it, or
completes it with the provided response bytes. If a network fetch occurs as a
result which encounters a redirect an additional Network.requestIntercepted
event will be sent with the same InterceptionID.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-continueInterceptedRequest
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) ContinueInterceptedRequest(
	params *network.ContinueInterceptedRequestParams,
) <-chan *network.ContinueInterceptedRequestResult {
	resultChan := make(chan *network.ContinueInterceptedRequestResult)
	result := &network.ContinueInterceptedRequestResult{}
	command := NewCommand(protocol.Socket, "Network.continueInterceptedRequest", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
DeleteCookies deletes browser cookies with matching name and url or domain/path
pair.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
*/
func (protocol *NetworkProtocol) DeleteCookies(
	params *network.DeleteCookiesParams,
) <-chan *network.DeleteCookiesResult {
	resultChan := make(chan *network.DeleteCookiesResult)
	result := &network.DeleteCookiesResult{}
	command := NewCommand(protocol.Socket, "Network.deleteCookies", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables network tracking, prevents network events from being sent to
the client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-disable
*/
func (protocol *NetworkProtocol) Disable() <-chan *network.DisableResult {
	resultChan := make(chan *network.DisableResult)
	result := &network.DisableResult{}
	command := NewCommand(protocol.Socket, "Network.disable", nil)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
EmulateConditions activates emulation of network conditions.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions
*/
func (protocol *NetworkProtocol) EmulateConditions(
	params *network.EmulateConditionsParams,
) <-chan *network.EmulateConditionsResult {
	resultChan := make(chan *network.EmulateConditionsResult)
	result := &network.EmulateConditionsResult{}
	command := NewCommand(protocol.Socket, "Network.emulateNetworkConditions", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables network tracking, network events will now be delivered to the
client.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
*/
func (protocol *NetworkProtocol) Enable(
	params *network.EnableParams,
) <-chan *network.EnableResult {
	resultChan := make(chan *network.EnableResult)
	result := &network.EnableResult{}
	command := NewCommand(protocol.Socket, "Network.enable", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetAllCookies returns all browser cookies. Depending on the backend support,
will return detailed cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getAllCookies
*/
func (protocol *NetworkProtocol) GetAllCookies() <-chan *network.GetAllCookiesResult {
	resultChan := make(chan *network.GetAllCookiesResult)
	command := NewCommand(protocol.Socket, "Network.getAllCookies", nil)
	result := &network.GetAllCookiesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetCertificate returns the DER-encoded certificate.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) GetCertificate(
	params *network.GetCertificateParams,
) <-chan *network.GetCertificateResult {
	resultChan := make(chan *network.GetCertificateResult)
	command := NewCommand(protocol.Socket, "Network.getCertificate", params)
	result := &network.GetCertificateResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetCookies returns all browser cookies for the current URL. Depending on the
backend support, will return detailed cookie information in the `cookies` field.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
*/
func (protocol *NetworkProtocol) GetCookies(
	params *network.GetCookiesParams,
) <-chan *network.GetCookiesResult {
	resultChan := make(chan *network.GetCookiesResult)
	command := NewCommand(protocol.Socket, "Network.getCookies", params)
	result := &network.GetCookiesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetResponseBody returns content served for the given request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
*/
func (protocol *NetworkProtocol) GetResponseBody(
	params *network.GetResponseBodyParams,
) <-chan *network.GetResponseBodyResult {
	resultChan := make(chan *network.GetResponseBodyResult)
	command := NewCommand(protocol.Socket, "Network.getResponseBody", params)
	result := &network.GetResponseBodyResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetResponseBodyForInterception returns content served for the given currently
intercepted request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) GetResponseBodyForInterception(
	params *network.GetResponseBodyForInterceptionParams,
) <-chan *network.GetResponseBodyForInterceptionResult {
	resultChan := make(chan *network.GetResponseBodyForInterceptionResult)
	command := NewCommand(protocol.Socket, "Network.getResponseBodyForInterception", params)
	result := &network.GetResponseBodyForInterceptionResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
ReplayXHR sends a new XMLHttpRequest which is identical to the original one. The
following parameters should be identical: method, url, async, request body,
extra headers, withCredentials attribute, user, password.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) ReplayXHR(
	params *network.ReplayXHRParams,
) <-chan *network.ReplayXHRResult {
	resultChan := make(chan *network.ReplayXHRResult)
	result := &network.ReplayXHRResult{}
	command := NewCommand(protocol.Socket, "Network.replayXHR", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SearchInResponseBody searches for given string in response content.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SearchInResponseBody(
	params *network.SearchInResponseBodyParams,
) <-chan *network.SearchInResponseBodyResult {
	resultChan := make(chan *network.SearchInResponseBodyResult)
	command := NewCommand(protocol.Socket, "Network.searchInResponseBody", params)
	result := &network.SearchInResponseBodyResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetBlockedURLs blocks URLs from loading.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetBlockedURLs(
	params *network.SetBlockedURLsParams,
) <-chan *network.SetBlockedURLsResult {
	resultChan := make(chan *network.SetBlockedURLsResult)
	result := &network.SetBlockedURLsResult{}
	command := NewCommand(protocol.Socket, "Network.setBlockedURLs", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetBypassServiceWorker toggles ignoring of service worker for each request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetBypassServiceWorker(
	params *network.SetBypassServiceWorkerParams,
) <-chan *network.SetBypassServiceWorkerResult {
	resultChan := make(chan *network.SetBypassServiceWorkerResult)
	result := &network.SetBypassServiceWorkerResult{}
	command := NewCommand(protocol.Socket, "Network.setBypassServiceWorker", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetCacheDisabled toggles ignoring cache for each request. If `true`, cache will
not be used.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
*/
func (protocol *NetworkProtocol) SetCacheDisabled(
	params *network.SetCacheDisabledParams,
) <-chan *network.SetCacheDisabledResult {
	resultChan := make(chan *network.SetCacheDisabledResult)
	result := &network.SetCacheDisabledResult{}
	command := NewCommand(protocol.Socket, "Network.setCacheDisabled", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetCookie sets a cookie with the given cookie data; may overwrite equivalent
cookies if they exist.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
*/
func (protocol *NetworkProtocol) SetCookie(
	params *network.SetCookieParams,
) <-chan *network.SetCookieResult {
	resultChan := make(chan *network.SetCookieResult)
	command := NewCommand(protocol.Socket, "Network.setCookie", params)
	result := &network.SetCookieResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetCookies sets given cookies.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
*/
func (protocol *NetworkProtocol) SetCookies(
	params *network.SetCookiesParams,
) <-chan *network.SetCookiesResult {
	resultChan := make(chan *network.SetCookiesResult)
	result := &network.SetCookiesResult{}
	command := NewCommand(protocol.Socket, "Network.setCookies", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetDataSizeLimitsForTest is for testing.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetDataSizeLimitsForTest(
	params *network.SetDataSizeLimitsForTestParams,
) <-chan *network.SetDataSizeLimitsForTestResult {
	resultChan := make(chan *network.SetDataSizeLimitsForTestResult)
	result := &network.SetDataSizeLimitsForTestResult{}
	command := NewCommand(protocol.Socket, "Network.setDataSizeLimitsForTest", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetExtraHTTPHeaders specifies whether to always send extra HTTP headers with the
requests from this page.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
*/
func (protocol *NetworkProtocol) SetExtraHTTPHeaders(
	params *network.SetExtraHTTPHeadersParams,
) <-chan *network.SetExtraHTTPHeadersResult {
	resultChan := make(chan *network.SetExtraHTTPHeadersResult)
	result := &network.SetExtraHTTPHeadersResult{}
	command := NewCommand(protocol.Socket, "Network.setExtraHTTPHeaders", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetRequestInterception sets the requests to intercept that match a the provided
patterns and optionally resource types.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
EXPERIMENTAL.
*/
func (protocol *NetworkProtocol) SetRequestInterception(
	params *network.SetRequestInterceptionParams,
) <-chan *network.SetRequestInterceptionResult {
	resultChan := make(chan *network.SetRequestInterceptionResult)
	result := &network.SetRequestInterceptionResult{}
	command := NewCommand(protocol.Socket, "Network.setRequestInterception", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetUserAgentOverride allows overriding user agent with the given string.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setUserAgentOverride
*/
func (protocol *NetworkProtocol) SetUserAgentOverride(
	params *network.SetUserAgentOverrideParams,
) <-chan *network.SetUserAgentOverrideResult {
	resultChan := make(chan *network.SetUserAgentOverrideResult)
	result := &network.SetUserAgentOverrideResult{}
	command := NewCommand(protocol.Socket, "Network.setUserAgentOverride", params)

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnDataReceived adds a handler to the Network.dataReceived event. Network.dataReceived
fires when a data chunk was received over the network.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-dataReceived
*/
func (protocol *NetworkProtocol) OnDataReceived(
	callback func(event *network.DataReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.dataReceived",
		func(response *Response) {
			event := &network.DataReceivedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnEventSourceMessageReceived adds a handler to the Network.eventSourceMessageReceived
event. Network.eventSourceMessageReceived fires when EventSource message is
received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-eventSourceMessageReceived
*/
func (protocol *NetworkProtocol) OnEventSourceMessageReceived(
	callback func(event *network.EventSourceMessageReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.eventSourceMessageReceived",
		func(response *Response) {
			event := &network.EventSourceMessageReceivedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLoadingFailed adds a handler to the Network.loadingFailed event. Network.loadingFailed
fires when HTTP request has failed to load.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFailed
*/
func (protocol *NetworkProtocol) OnLoadingFailed(
	callback func(event *network.LoadingFailedEvent),
) {
	handler := NewEventHandler(
		"Network.loadingFailed",
		func(response *Response) {
			event := &network.LoadingFailedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnLoadingFinished adds a handler to the Network.loadingFinished event.
Network.loadingFinished fires when HTTP request has finished loading.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFinished
*/
func (protocol *NetworkProtocol) OnLoadingFinished(
	callback func(event *network.LoadingFinishedEvent),
) {
	handler := NewEventHandler(
		"Network.loadingFinished",
		func(response *Response) {
			event := &network.LoadingFinishedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnRequestIntercepted adds a handler to the Network.requestIntercepted event.
Network.requestIntercepted fires when a HTTP request is intercepted and returns
details, which must be either allowed, blocked, modified or mocked.

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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnRequestServedFromCache adds a handler to the Network.requestServedFromCache
event. Network.requestServedFromCache fires when request ended up loading from
cache.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestServedFromCache
*/
func (protocol *NetworkProtocol) OnRequestServedFromCache(
	callback func(event *network.RequestServedFromCacheEvent),
) {
	handler := NewEventHandler(
		"Network.requestServedFromCache",
		func(response *Response) {
			event := &network.RequestServedFromCacheEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnRequestWillBeSent adds a handler to the Network.requestWillBeSent event.
Network.requestWillBeSent fires when the page is about to send HTTP request.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestWillBeSent
*/
func (protocol *NetworkProtocol) OnRequestWillBeSent(
	callback func(event *network.RequestWillBeSentEvent),
) {
	handler := NewEventHandler(
		"Network.requestWillBeSent",
		func(response *Response) {
			event := &network.RequestWillBeSentEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnResourceChangedPriority adds a handler to the Network.resourceChangedPriority
event. Network.resourceChangedPriority fires when resource loading priority is
changed.

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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnResponseReceived adds a handler to the Network.responseReceived event.
Network.responseReceived fires when HTTP response is available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-responseReceived
*/
func (protocol *NetworkProtocol) OnResponseReceived(
	callback func(event *network.ResponseReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.responseReceived",
		func(response *Response) {
			event := &network.ResponseReceivedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketClosed adds a handler to the Network.webSocketClosed event.
Network.webSocketClosed fires when WebSocket is closed.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketClosed
*/
func (protocol *NetworkProtocol) OnWebSocketClosed(
	callback func(event *network.WebSocketClosedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketClosed",
		func(response *Response) {
			event := &network.WebSocketClosedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketCreated adds a handler to the Network.webSocketCreated event.
Network.webSocketCreated fires upon WebSocket creation.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketCreated
*/
func (protocol *NetworkProtocol) OnWebSocketCreated(
	callback func(event *network.WebSocketCreatedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketCreated",
		func(response *Response) {
			event := &network.WebSocketCreatedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketFrameReceived adds a handler to the Network.webSocketFrameReceived
event. Network.webSocketFrameReceived fires when WebSocket frame is received.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameReceived
*/
func (protocol *NetworkProtocol) OnWebSocketFrameReceived(
	callback func(event *network.WebSocketFrameReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketFrameReceived",
		func(response *Response) {
			event := &network.WebSocketFrameReceivedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketHandshakeResponseReceived adds a handler to the Network.webSocketHandshakeResponseReceived
event. Network.webSocketHandshakeResponseReceived fires when WebSocket handshake
response becomes available.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketHandshakeResponseReceived
*/
func (protocol *NetworkProtocol) OnWebSocketHandshakeResponseReceived(
	callback func(event *network.WebSocketHandshakeResponseReceivedEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketHandshakeResponseReceived",
		func(response *Response) {
			event := &network.WebSocketHandshakeResponseReceivedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnWebSocketWillSendHandshakeRequest adds a handler to the Network.webSocketWillSendHandshakeRequest
event. Network.webSocketWillSendHandshakeRequest fires when WebSocket is about
to initiate handshake.

https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketWillSendHandshakeRequest
*/
func (protocol *NetworkProtocol) OnWebSocketWillSendHandshakeRequest(
	callback func(event *network.WebSocketWillSendHandshakeRequestEvent),
) {
	handler := NewEventHandler(
		"Network.webSocketWillSendHandshakeRequest",
		func(response *Response) {
			event := &network.WebSocketWillSendHandshakeRequestEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
