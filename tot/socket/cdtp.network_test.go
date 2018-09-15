package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/debugger"
	"github.com/mkenney/go-chrome/tot/network"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/runtime"
	"github.com/mkenney/go-chrome/tot/security"
)

func TestNetworkCanClearBrowserCache(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &network.CanClearBrowserCacheResult{
		Result: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().CanClearBrowserCache()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().CanClearBrowserCache()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkCanClearBrowserCookies(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &network.CanClearBrowserCookiesResult{
		Result: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().CanClearBrowserCookies()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().CanClearBrowserCookies()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkCanEmulateConditions(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &network.CanEmulateConditionsResult{
		Result: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().CanEmulateConditions()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().CanEmulateConditions()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkClearBrowserCache(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &network.ClearBrowserCacheResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().ClearBrowserCache()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().ClearBrowserCache()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkClearBrowserCookies(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &network.ClearBrowserCookiesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().ClearBrowserCookies()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().ClearBrowserCookies()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkContinueInterceptedRequest(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.ContinueInterceptedRequestParams{
		InterceptionID: network.InterceptionID("interception-id"),
		ErrorReason:    network.ErrorReason.Failed,
		RawResponse:    "raw response",
		URL:            "http://some.url",
		Method:         "someMethod",
		PostData:       "post data",
		Headers:        network.Headers{"header": "value"},
		AuthChallengeResponse: &network.AuthChallengeResponse{
			Response: network.ChallengeResponse.Default,
			Username: "username",
			Password: "password",
		},
	}
	mockResult := &network.ContinueInterceptedRequestResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().ContinueInterceptedRequest(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().ContinueInterceptedRequest(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkDeleteCookies(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.DeleteCookiesParams{
		Name:   "name",
		URL:    "http://some.url",
		Domain: "some.url",
		Path:   "/",
	}
	mockResult := &network.DeleteCookiesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().DeleteCookies(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().DeleteCookies(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &network.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkEmulateConditions(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.EmulateConditionsParams{
		Offline:            true,
		Latency:            1,
		DownloadThroughput: 1,
		UploadThroughput:   1,
		ConnectionType:     network.ConnectionType.None,
	}
	mockResult := &network.EmulateConditionsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().EmulateConditions(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().EmulateConditions(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
func TestNetworkEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.EnableParams{
		MaxTotalBufferSize:    1,
		MaxResourceBufferSize: 1,
	}
	mockResult := &network.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().Enable(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().Enable(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetAllCookies(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &network.GetAllCookiesResult{
		Cookies: []*network.Cookie{{
			Name:     "name",
			Value:    "value",
			Domain:   "domain",
			Path:     "/",
			Expires:  time.Now().Unix() + 10,
			Size:     1,
			HTTPOnly: true,
			Secure:   true,
			Session:  true,
			SameSite: network.CookieSameSite.Strict,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().GetAllCookies()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Cookies[0].Name != result.Cookies[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Cookies[0].Name, result.Cookies[0].Name)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().GetAllCookies()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetCertificate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.GetCertificateParams{
		Origin: "origin",
	}
	mockResult := &network.GetCertificateResult{
		TableNames: []string{"name1", "name2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().GetCertificate(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.TableNames[0] != result.TableNames[0] {
		t.Errorf("Expected %s, got %s", mockResult.TableNames[0], result.TableNames[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().GetCertificate(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetCookies(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.GetCookiesParams{
		URLs: []string{"url1", "url2"},
	}
	mockResult := &network.GetCookiesResult{
		Cookies: []*network.Cookie{{
			Name:     "name",
			Value:    "value",
			Domain:   "domain",
			Path:     "/",
			Expires:  time.Now().Unix() + 10,
			Size:     1,
			HTTPOnly: true,
			Secure:   true,
			Session:  true,
			SameSite: network.CookieSameSite.Strict,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().GetCookies(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Cookies[0].Name != result.Cookies[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Cookies[0].Name, result.Cookies[0].Name)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().GetCookies(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetResponseBody(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.GetResponseBodyParams{
		RequestID: network.RequestID("request-id"),
	}
	mockResult := &network.GetResponseBodyResult{
		Body:          "body data",
		Base64Encoded: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().GetResponseBody(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Body != result.Body {
		t.Errorf("Expected %s, got %s", mockResult.Body, result.Body)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().GetResponseBody(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetResponseBodyForInterception(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.GetResponseBodyForInterceptionParams{
		InterceptionID: network.InterceptionID("interception-id"),
	}
	mockResult := &network.GetResponseBodyForInterceptionResult{
		Body:          "body data",
		Base64Encoded: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().GetResponseBodyForInterception(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Body != result.Body {
		t.Errorf("Expected %s, got %s", mockResult.Body, result.Body)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().GetResponseBodyForInterception(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkReplayXHR(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.ReplayXHRParams{
		RequestID: network.RequestID("request-id"),
	}
	mockResult := &network.ReplayXHRResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().ReplayXHR(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().ReplayXHR(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSearchInResponseBody(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SearchInResponseBodyParams{
		RequestID:     network.RequestID("request-id"),
		Query:         "query string",
		CaseSensitive: true,
		IsRegex:       true,
	}
	mockResult := &network.SearchInResponseBodyResult{
		Result: []*debugger.SearchMatch{{
			LineNumber:  1,
			LineContent: "content",
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SearchInResponseBody(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].LineNumber != result.Result[0].LineNumber {
		t.Errorf("Expected %d, got %d", mockResult.Result[0].LineNumber, result.Result[0].LineNumber)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SearchInResponseBody(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetBlockedURLs(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetBlockedURLsParams{
		URLs: []string{"http://url.1", "http://url.2"},
	}
	mockResult := &network.SetBlockedURLsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetBlockedURLs(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetBlockedURLs(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetBypassServiceWorker(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetBypassServiceWorkerParams{
		Bypass: true,
	}
	mockResult := &network.SetBypassServiceWorkerResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetBypassServiceWorker(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetBypassServiceWorker(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetCacheDisabled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetCacheDisabledParams{
		CacheDisabled: true,
	}
	mockResult := &network.SetCacheDisabledResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetCacheDisabled(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetCacheDisabled(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetCookie(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetCookieParams{
		Name:     "name",
		Value:    "value",
		URL:      "http://some.url",
		Domain:   "some.url",
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
		SameSite: network.CookieSameSite.Strict,
		Expires:  network.TimeSinceEpoch(time.Now().Unix()),
	}
	mockResult := &network.SetCookieResult{
		Success: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetCookie(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Success != result.Success {
		t.Errorf("Expected %v, got %v", mockResult.Success, result.Success)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetCookie(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetCookies(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetCookiesParams{
		Cookies: []*network.SetCookieParams{{
			Name:     "name",
			Value:    "value",
			URL:      "http://some.url",
			Domain:   "some.url",
			Path:     "/",
			Secure:   true,
			HTTPOnly: true,
			SameSite: network.CookieSameSite.Strict,
			Expires:  network.TimeSinceEpoch(time.Now().Unix()),
		}},
	}
	mockResult := &network.SetCookiesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetCookies(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetCookies(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetDataSizeLimitsForTest(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetDataSizeLimitsForTestParams{
		MaxTotalSize:    1,
		MaxResourceSize: 1,
	}
	mockResult := &network.SetDataSizeLimitsForTestResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetDataSizeLimitsForTest(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetDataSizeLimitsForTest(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetExtraHTTPHeaders(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetExtraHTTPHeadersParams{
		Headers: network.Headers{"header1": "value1", "header2": "value2"},
	}
	mockResult := &network.SetExtraHTTPHeadersResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetExtraHTTPHeaders(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetExtraHTTPHeaders(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetRequestInterception(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetRequestInterceptionParams{
		Patterns: []*network.RequestPattern{{
			URLPattern:        "url pattern",
			ResourceType:      page.ResourceType.Document,
			InterceptionStage: network.InterceptionStage.Request,
		}},
	}
	mockResult := &network.SetRequestInterceptionResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetRequestInterception(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetRequestInterception(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetUserAgentOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &network.SetUserAgentOverrideParams{
		UserAgent: "user-agent",
	}
	mockResult := &network.SetUserAgentOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Network().SetUserAgentOverride(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Network().SetUserAgentOverride(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnDataReceived(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()

	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.DataReceivedEvent)
	soc.Network().OnDataReceived(func(eventData *network.DataReceivedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.DataReceivedEvent{
		RequestID:         network.RequestID("request-id"),
		Timestamp:         network.MonotonicTime(1),
		DataLength:        1,
		EncodedDataLength: 1,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.dataReceived",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.dataReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnEventSourceMessageReceived(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.EventSourceMessageReceivedEvent)
	soc.Network().OnEventSourceMessageReceived(func(eventData *network.EventSourceMessageReceivedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.EventSourceMessageReceivedEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
		EventName: "event-name",
		EventID:   "event-id",
		Data:      "some data",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.eventSourceMessageReceived",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.eventSourceMessageReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnLoadingFailed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.LoadingFailedEvent)
	soc.Network().OnLoadingFailed(func(eventData *network.LoadingFailedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.LoadingFailedEvent{
		RequestID:     network.RequestID("request-id"),
		Timestamp:     network.MonotonicTime(1),
		Type:          page.ResourceType.Document,
		ErrorText:     "error text",
		Canceled:      true,
		BlockedReason: network.BlockedReason.Csp,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.loadingFailed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.loadingFailed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnLoadingFinished(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.LoadingFinishedEvent)
	soc.Network().OnLoadingFinished(func(eventData *network.LoadingFinishedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.LoadingFinishedEvent{
		RequestID:         network.RequestID("request-id"),
		Timestamp:         network.MonotonicTime(1),
		EncodedDataLength: 1,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.loadingFinished",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.loadingFinished",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnRequestIntercepted(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.RequestInterceptedEvent)
	soc.Network().OnRequestIntercepted(func(eventData *network.RequestInterceptedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.RequestInterceptedEvent{
		InterceptionID: network.InterceptionID("interception-id"),
		Request: &network.Request{
			URL:              "http://some.url",
			Method:           "GET",
			Headers:          network.Headers{"header1": "value1"},
			PostData:         "post data",
			MixedContentType: security.MixedContentType.Blockable,
			InitialPriority:  network.ResourcePriority.VeryLow,
			ReferrerPolicy:   network.ReferrerPolicy.UnsafeURL,
			IsLinkPreload:    true,
		},
		FrameID:             page.FrameID("frame-id"),
		ResourceType:        page.ResourceType.Document,
		IsNavigationRequest: true,
		RedirectURL:         "http://some.other.url",
		AuthChallenge: &network.AuthChallenge{
			Source: network.Source.Server,
			Origin: "origin",
			Scheme: "scheme",
			Realm:  "realm",
		},
		ResponseErrorReason: network.ErrorReason.Failed,
		ResponseStatusCode:  400,
		ResponseHeaders:     network.Headers{"header1": "value1"},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.requestIntercepted",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.InterceptionID != result.InterceptionID {
		t.Errorf("Expected %s, got %s", mockResult.InterceptionID, result.InterceptionID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.requestIntercepted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnRequestServedFromCache(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.RequestServedFromCacheEvent)
	soc.Network().OnRequestServedFromCache(func(eventData *network.RequestServedFromCacheEvent) {
		resultChan <- eventData
	})

	mockResult := &network.RequestServedFromCacheEvent{
		RequestID: network.RequestID("request-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.requestServedFromCache",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.requestServedFromCache",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnRequestWillBeSent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.RequestWillBeSentEvent)
	soc.Network().OnRequestWillBeSent(func(eventData *network.RequestWillBeSentEvent) {
		resultChan <- eventData
	})

	mockResult := &network.RequestWillBeSentEvent{
		RequestID:   network.RequestID("request-id"),
		LoaderID:    network.LoaderID("loader-id"),
		DocumentURL: "http://some.url",
		Request: &network.Request{
			URL:              "http://some.url",
			Method:           "GET",
			Headers:          network.Headers{"header1": "value1"},
			PostData:         "post data",
			MixedContentType: security.MixedContentType.Blockable,
			InitialPriority:  network.ResourcePriority.VeryLow,
			ReferrerPolicy:   network.ReferrerPolicy.UnsafeURL,
			IsLinkPreload:    true,
		},
		Timestamp: network.MonotonicTime(1),
		WallTime:  network.TimeSinceEpoch(time.Now().Unix()),
		Initiator: &network.Initiator{
			Type:       network.InitiatorType.Parser,
			Stack:      &runtime.StackTrace{},
			URL:        "http://some.url",
			LineNumber: 1,
		},
		RedirectResponse: &network.Response{
			URL:                "http://some.url",
			Status:             200,
			StatusText:         "OK",
			Headers:            network.Headers{},
			HeadersText:        "header text",
			MimeType:           "mime type",
			RequestHeaders:     network.Headers{},
			RequestHeadersText: "header text",
			ConnectionReused:   true,
			ConnectionID:       1,
			RemoteIPAddress:    "1.2.3.4",
			RemotePort:         8080,
			FromDiskCache:      true,
			FromServiceWorker:  true,
			EncodedDataLength:  1,
			Timing: &network.ResourceTiming{
				RequestTime:       1,
				ProxyStart:        1,
				ProxyEnd:          1,
				DNSStart:          1,
				DNSEnd:            1,
				ConnectStart:      1,
				ConnectEnd:        1,
				SSLStart:          1,
				SSLEnd:            1,
				WorkerStart:       1,
				WorkerReady:       1,
				SendStart:         1,
				SendEnd:           1,
				PushStart:         1,
				PushEnd:           1,
				ReceiveHeadersEnd: 1,
			},
			Protocol:      "protocol",
			SecurityState: security.State.Unknown,
			SecurityDetails: &network.SecurityDetails{
				Protocol:         "TLS 1.2",
				KeyExchange:      "key-exchange",
				KeyExchangeGroup: "key-exchange-group",
				Cipher:           "cipher",
				Mac:              "mac address",
				CertificateID:    security.CertificateID(1),
				SubjectName:      "subject name",
				SanList:          []string{"one", "two"},
				Issuer:           "issuer",
				ValidFrom:        network.TimeSinceEpoch(time.Now().Unix()),
				ValidTo:          network.TimeSinceEpoch(time.Now().Unix() + 10),
				SignedCertificateTimestampList: []*network.SignedCertificateTimestamp{{
					Status:             "status",
					Origin:             "origin",
					LogDescription:     "description",
					LogID:              "log-id",
					Timestamp:          network.TimeSinceEpoch(time.Now().Unix()),
					HashAlgorithm:      "alg",
					SignatureAlgorithm: "alg",
					SignatureData:      "data",
				}},
			},
		},
		Type:    page.ResourceType.Document,
		FrameID: page.FrameID("frame-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.requestWillBeSent",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.requestWillBeSent",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnResourceChangedPriority(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.ResourceChangedPriorityEvent)
	soc.Network().OnResourceChangedPriority(func(eventData *network.ResourceChangedPriorityEvent) {
		resultChan <- eventData
	})

	mockResult := &network.ResourceChangedPriorityEvent{
		RequestID:   network.RequestID("request-id"),
		NewPriority: network.ResourcePriority.VeryLow,
		Timestamp:   network.MonotonicTime(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.resourceChangedPriority",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.resourceChangedPriority",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnResponseReceived(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.ResponseReceivedEvent)
	soc.Network().OnResponseReceived(func(eventData *network.ResponseReceivedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.ResponseReceivedEvent{
		RequestID: network.RequestID("request-id"),
		LoaderID:  network.LoaderID("loader-id"),
		Timestamp: network.MonotonicTime(1),
		Type:      page.ResourceType.Document,
		Response:  &network.Response{},
		FrameID:   page.FrameID("frame-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.responseReceived",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.responseReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketClosed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.WebSocketClosedEvent)
	soc.Network().OnWebSocketClosed(func(eventData *network.WebSocketClosedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.WebSocketClosedEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.webSocketClosed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.webSocketClosed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketCreated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.WebSocketCreatedEvent)
	soc.Network().OnWebSocketCreated(func(eventData *network.WebSocketCreatedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.WebSocketCreatedEvent{
		RequestID:    network.RequestID("request-id"),
		Timestamp:    network.MonotonicTime(1),
		ErrorMessage: "error message",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.webSocketCreated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.webSocketCreated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketFrameError(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.WebSocketFrameErrorEvent)
	soc.Network().OnWebSocketFrameError(func(eventData *network.WebSocketFrameErrorEvent) {
		resultChan <- eventData
	})

	mockResult := &network.WebSocketFrameErrorEvent{
		RequestID:    network.RequestID("request-id"),
		Timestamp:    network.MonotonicTime(1),
		ErrorMessage: "error message",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.webSocketFrameError",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.webSocketFrameError",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketFrameReceived(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.WebSocketFrameReceivedEvent)
	soc.Network().OnWebSocketFrameReceived(func(eventData *network.WebSocketFrameReceivedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.WebSocketFrameReceivedEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
		Response: &network.WebSocketFrame{
			Opcode:      1,
			Mask:        true,
			PayloadData: "payload",
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.webSocketFrameReceived",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.webSocketFrameReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketFrameSent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.WebSocketFrameSentEvent)
	soc.Network().OnWebSocketFrameSent(func(eventData *network.WebSocketFrameSentEvent) {
		resultChan <- eventData
	})

	mockResult := &network.WebSocketFrameSentEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
		Response: &network.WebSocketFrame{
			Opcode:      1,
			Mask:        true,
			PayloadData: "payload",
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.webSocketFrameSent",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.webSocketFrameSent",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketHandshakeResponseReceived(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.WebSocketHandshakeResponseReceivedEvent)
	soc.Network().OnWebSocketHandshakeResponseReceived(func(eventData *network.WebSocketHandshakeResponseReceivedEvent) {
		resultChan <- eventData
	})

	mockResult := &network.WebSocketHandshakeResponseReceivedEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
		Response: &network.WebSocketFrame{
			Opcode:      1,
			Mask:        true,
			PayloadData: "payload",
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.webSocketHandshakeResponseReceived",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.webSocketHandshakeResponseReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketWillSendHandshakeRequest(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *network.WebSocketWillSendHandshakeRequestEvent)
	soc.Network().OnWebSocketWillSendHandshakeRequest(func(eventData *network.WebSocketWillSendHandshakeRequestEvent) {
		resultChan <- eventData
	})

	mockResult := &network.WebSocketWillSendHandshakeRequestEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
		WallTime:  network.TimeSinceEpoch(time.Now().Unix()),
		Request: &network.WebSocketRequest{
			Headers: network.Headers{"header1": "value1"},
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Network.webSocketWillSendHandshakeRequest",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Network.webSocketWillSendHandshakeRequest",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
