package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	debugger "github.com/mkenney/go-chrome/tot/cdtp/debugger"
	network "github.com/mkenney/go-chrome/tot/cdtp/network"
	page "github.com/mkenney/go-chrome/tot/cdtp/page"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
	security "github.com/mkenney/go-chrome/tot/cdtp/security"
)

func TestNetworkCanClearBrowserCache(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Network().CanClearBrowserCache()
	mockResult := &network.CanClearBrowserCacheResult{
		Result: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	resultChan = mockSocket.Network().CanClearBrowserCache()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkCanClearBrowserCookies(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Network().CanClearBrowserCookies()
	mockResult := &network.CanClearBrowserCookiesResult{
		Result: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	resultChan = mockSocket.Network().CanClearBrowserCookies()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkCanEmulateConditions(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Network().CanEmulateConditions()
	mockResult := &network.CanEmulateConditionsResult{
		Result: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	resultChan = mockSocket.Network().CanEmulateConditions()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkClearBrowserCache(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Network().ClearBrowserCache()
	mockResult := &network.ClearBrowserCacheResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().ClearBrowserCache()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkClearBrowserCookies(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Network().ClearBrowserCookies()
	mockResult := &network.ClearBrowserCookiesResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().ClearBrowserCookies()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkContinueInterceptedRequest(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.ContinueInterceptedRequestParams{
		InterceptionID: network.InterceptionID("interception-id"),
		ErrorReason:    network.ErrorReason("error-reason"),
		RawResponse:    "raw response",
		URL:            "http://some.url",
		Method:         "someMethod",
		PostData:       "post data",
		Headers:        network.Headers{"header": "value"},
		AuthChallengeResponse: &network.AuthChallengeResponse{
			Response: "Default",
			Username: "username",
			Password: "password",
		},
	}
	resultChan := mockSocket.Network().ContinueInterceptedRequest(params)
	mockResult := &network.ContinueInterceptedRequestResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().ContinueInterceptedRequest(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkDeleteCookies(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.DeleteCookiesParams{
		Name:   "name",
		URL:    "http://some.url",
		Domain: "some.url",
		Path:   "/",
	}
	resultChan := mockSocket.Network().DeleteCookies(params)
	mockResult := &network.DeleteCookiesResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().DeleteCookies(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Network().Disable()
	mockResult := &network.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().Disable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkEmulateConditions(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.EmulateConditionsParams{
		Offline:            true,
		Latency:            1,
		DownloadThroughput: 1,
		UploadThroughput:   1,
		ConnectionType:     network.ConnectionType("connection-type"),
	}
	resultChan := mockSocket.Network().EmulateConditions(params)
	mockResult := &network.EmulateConditionsResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().EmulateConditions(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
func TestNetworkEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.EnableParams{
		MaxTotalBufferSize:    1,
		MaxResourceBufferSize: 1,
	}
	resultChan := mockSocket.Network().Enable(params)
	mockResult := &network.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().Enable(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetAllCookies(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Network().GetAllCookies()
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
			SameSite: network.CookieSameSite("cookie-same-site"),
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Cookies[0].Name != result.Cookies[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Cookies[0].Name, result.Cookies[0].Name)
	}

	resultChan = mockSocket.Network().GetAllCookies()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetCertificate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.GetCertificateParams{
		Origin: "origin",
	}
	resultChan := mockSocket.Network().GetCertificate(params)
	mockResult := &network.GetCertificateResult{
		TableNames: []string{"name1", "name2"},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.TableNames[0] != result.TableNames[0] {
		t.Errorf("Expected %s, got %s", mockResult.TableNames[0], result.TableNames[0])
	}

	resultChan = mockSocket.Network().GetCertificate(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetCookies(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.GetCookiesParams{
		URLs: []string{"url1", "url2"},
	}
	resultChan := mockSocket.Network().GetCookies(params)
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
			SameSite: network.CookieSameSite("cookie-same-site"),
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Cookies[0].Name != result.Cookies[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Cookies[0].Name, result.Cookies[0].Name)
	}

	resultChan = mockSocket.Network().GetCookies(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetResponseBody(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.GetResponseBodyParams{
		RequestID: network.RequestID("request-id"),
	}
	resultChan := mockSocket.Network().GetResponseBody(params)
	mockResult := &network.GetResponseBodyResult{
		Body:          "body data",
		Base64Encoded: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Body != result.Body {
		t.Errorf("Expected %s, got %s", mockResult.Body, result.Body)
	}

	resultChan = mockSocket.Network().GetResponseBody(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkGetResponseBodyForInterception(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.GetResponseBodyForInterceptionParams{
		InterceptionID: network.InterceptionID("interception-id"),
	}
	resultChan := mockSocket.Network().GetResponseBodyForInterception(params)
	mockResult := &network.GetResponseBodyForInterceptionResult{
		Body:          "body data",
		Base64Encoded: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Body != result.Body {
		t.Errorf("Expected %s, got %s", mockResult.Body, result.Body)
	}

	resultChan = mockSocket.Network().GetResponseBodyForInterception(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkReplayXHR(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.ReplayXHRParams{
		RequestID: network.RequestID("request-id"),
	}
	resultChan := mockSocket.Network().ReplayXHR(params)
	mockResult := &network.ReplayXHRResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().ReplayXHR(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSearchInResponseBody(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SearchInResponseBodyParams{
		RequestID:     network.RequestID("request-id"),
		Query:         "query string",
		CaseSensitive: true,
		IsRegex:       true,
	}
	resultChan := mockSocket.Network().SearchInResponseBody(params)
	mockResult := &network.SearchInResponseBodyResult{
		Result: []*debugger.SearchMatch{{
			LineNumber:  1,
			LineContent: "content",
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].LineNumber != result.Result[0].LineNumber {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].LineNumber, result.Result[0].LineNumber)
	}

	resultChan = mockSocket.Network().SearchInResponseBody(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetBlockedURLs(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetBlockedURLsParams{
		URLs: []string{"http://url.1", "http://url.2"},
	}
	resultChan := mockSocket.Network().SetBlockedURLs(params)
	mockResult := &network.SetBlockedURLsResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetBlockedURLs(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetBypassServiceWorker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetBypassServiceWorkerParams{
		Bypass: true,
	}
	resultChan := mockSocket.Network().SetBypassServiceWorker(params)
	mockResult := &network.SetBypassServiceWorkerResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetBypassServiceWorker(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetCacheDisabled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetCacheDisabledParams{
		CacheDisabled: true,
	}
	resultChan := mockSocket.Network().SetCacheDisabled(params)
	mockResult := &network.SetCacheDisabledResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetCacheDisabled(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetCookie(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetCookieParams{
		Name:     "name",
		Value:    "value",
		URL:      "http://some.url",
		Domain:   "some.url",
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
		SameSite: network.CookieSameSite("same site"),
		Expires:  network.TimeSinceEpoch(time.Now().Unix()),
	}
	resultChan := mockSocket.Network().SetCookie(params)
	mockResult := &network.SetCookieResult{
		Success: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Success != result.Success {
		t.Errorf("Expected %v, got %v", mockResult.Success, result.Success)
	}

	resultChan = mockSocket.Network().SetCookie(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetCookies(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetCookiesParams{
		Cookies: []*network.SetCookieParams{{
			Name:     "name",
			Value:    "value",
			URL:      "http://some.url",
			Domain:   "some.url",
			Path:     "/",
			Secure:   true,
			HTTPOnly: true,
			SameSite: network.CookieSameSite("same site"),
			Expires:  network.TimeSinceEpoch(time.Now().Unix()),
		}},
	}
	resultChan := mockSocket.Network().SetCookies(params)
	mockResult := &network.SetCookiesResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetCookies(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetDataSizeLimitsForTest(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetDataSizeLimitsForTestParams{
		MaxTotalSize:    1,
		MaxResourceSize: 1,
	}
	resultChan := mockSocket.Network().SetDataSizeLimitsForTest(params)
	mockResult := &network.SetDataSizeLimitsForTestResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetDataSizeLimitsForTest(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetExtraHTTPHeaders(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetExtraHTTPHeadersParams{
		Headers: network.Headers{"header1": "value1", "header2": "value2"},
	}
	resultChan := mockSocket.Network().SetExtraHTTPHeaders(params)
	mockResult := &network.SetExtraHTTPHeadersResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetExtraHTTPHeaders(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetRequestInterception(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetRequestInterceptionParams{
		Patterns: []*network.RequestPattern{{
			URLPattern:        "url pattern",
			ResourceType:      page.ResourceType("resource-type"),
			InterceptionStage: network.InterceptionStage("interception-stage"),
		}},
	}
	resultChan := mockSocket.Network().SetRequestInterception(params)
	mockResult := &network.SetRequestInterceptionResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetRequestInterception(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkSetUserAgentOverride(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &network.SetUserAgentOverrideParams{
		UserAgent: "user-agent",
	}
	resultChan := mockSocket.Network().SetUserAgentOverride(params)
	mockResult := &network.SetUserAgentOverrideResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Network().SetUserAgentOverride(params)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnDataReceived(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.DataReceivedEvent)
	mockSocket.Network().OnDataReceived(func(eventData *network.DataReceivedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.DataReceivedEvent{
		RequestID:         network.RequestID("request-id"),
		Timestamp:         network.MonotonicTime(1),
		DataLength:        1,
		EncodedDataLength: 1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.dataReceived",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.DataReceivedEvent)
	mockSocket.Network().OnDataReceived(func(eventData *network.DataReceivedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.dataReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnEventSourceMessageReceived(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.EventSourceMessageReceivedEvent)
	mockSocket.Network().OnEventSourceMessageReceived(func(eventData *network.EventSourceMessageReceivedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.EventSourceMessageReceivedEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
		EventName: "event-name",
		EventID:   "event-id",
		Data:      "some data",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.eventSourceMessageReceived",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.EventSourceMessageReceivedEvent)
	mockSocket.Network().OnEventSourceMessageReceived(func(eventData *network.EventSourceMessageReceivedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.eventSourceMessageReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnLoadingFailed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.LoadingFailedEvent)
	mockSocket.Network().OnLoadingFailed(func(eventData *network.LoadingFailedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.LoadingFailedEvent{
		RequestID:     network.RequestID("request-id"),
		Timestamp:     network.MonotonicTime(1),
		Type:          page.ResourceType("resource-type"),
		ErrorText:     "error text",
		Canceled:      true,
		BlockedReason: network.BlockedReason("blocked-reason"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.loadingFailed",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.LoadingFailedEvent)
	mockSocket.Network().OnLoadingFailed(func(eventData *network.LoadingFailedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.loadingFailed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnLoadingFinished(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.LoadingFinishedEvent)
	mockSocket.Network().OnLoadingFinished(func(eventData *network.LoadingFinishedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.LoadingFinishedEvent{
		RequestID:         network.RequestID("request-id"),
		Timestamp:         network.MonotonicTime(1),
		EncodedDataLength: 1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.loadingFinished",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.LoadingFinishedEvent)
	mockSocket.Network().OnLoadingFinished(func(eventData *network.LoadingFinishedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.loadingFinished",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnRequestIntercepted(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.RequestInterceptedEvent)
	mockSocket.Network().OnRequestIntercepted(func(eventData *network.RequestInterceptedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.RequestInterceptedEvent{
		InterceptionID: network.InterceptionID("interception-id"),
		Request: &network.Request{
			URL:              "http://some.url",
			Method:           "GET",
			Headers:          network.Headers{"header1": "value1"},
			PostData:         "post data",
			MixedContentType: security.MixedContentType("mixed-content-type"),
			InitialPriority:  network.ResourcePriority("initial-priority"),
			ReferrerPolicy:   "unsafe-url",
			IsLinkPreload:    true,
		},
		FrameID:             page.FrameID("frame-id"),
		ResourceType:        page.ResourceType("resource-type"),
		IsNavigationRequest: true,
		RedirectURL:         "http://some.other.url",
		AuthChallenge: &network.AuthChallenge{
			Source: "server",
			Origin: "origin",
			Scheme: "scheme",
			Realm:  "realm",
		},
		ResponseErrorReason: network.ErrorReason("reason"),
		ResponseStatusCode:  400,
		ResponseHeaders:     network.Headers{"header1": "value1"},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.requestIntercepted",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.InterceptionID != result.InterceptionID {
		t.Errorf("Expected %s, got %s", mockResult.InterceptionID, result.InterceptionID)
	}

	resultChan = make(chan *network.RequestInterceptedEvent)
	mockSocket.Network().OnRequestIntercepted(func(eventData *network.RequestInterceptedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.requestIntercepted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnRequestServedFromCache(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.RequestServedFromCacheEvent)
	mockSocket.Network().OnRequestServedFromCache(func(eventData *network.RequestServedFromCacheEvent) {
		resultChan <- eventData
	})
	mockResult := &network.RequestServedFromCacheEvent{
		RequestID: network.RequestID("request-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.requestServedFromCache",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.RequestServedFromCacheEvent)
	mockSocket.Network().OnRequestServedFromCache(func(eventData *network.RequestServedFromCacheEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.requestServedFromCache",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnRequestWillBeSent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.RequestWillBeSentEvent)
	mockSocket.Network().OnRequestWillBeSent(func(eventData *network.RequestWillBeSentEvent) {
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
			MixedContentType: security.MixedContentType("mixed-content-type"),
			InitialPriority:  network.ResourcePriority("initial-priority"),
			ReferrerPolicy:   "unsafe-url",
			IsLinkPreload:    true,
		},
		Timestamp: network.MonotonicTime(1),
		WallTime:  network.TimeSinceEpoch(time.Now().Unix()),
		Initiator: &network.Initiator{
			Type:       "type",
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
			SecurityState: security.State("state"),
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
		Type:    page.ResourceType("resource-type"),
		FrameID: page.FrameID("frame-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.requestWillBeSent",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.RequestWillBeSentEvent)
	mockSocket.Network().OnRequestWillBeSent(func(eventData *network.RequestWillBeSentEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.requestWillBeSent",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnResourceChangedPriority(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.ResourceChangedPriorityEvent)
	mockSocket.Network().OnResourceChangedPriority(func(eventData *network.ResourceChangedPriorityEvent) {
		resultChan <- eventData
	})
	mockResult := &network.ResourceChangedPriorityEvent{
		RequestID:   network.RequestID("request-id"),
		NewPriority: network.ResourcePriority("resource-priority"),
		Timestamp:   network.MonotonicTime(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.resourceChangedPriority",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.ResourceChangedPriorityEvent)
	mockSocket.Network().OnResourceChangedPriority(func(eventData *network.ResourceChangedPriorityEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.resourceChangedPriority",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnResponseReceived(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.ResponseReceivedEvent)
	mockSocket.Network().OnResponseReceived(func(eventData *network.ResponseReceivedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.ResponseReceivedEvent{
		RequestID: network.RequestID("request-id"),
		LoaderID:  network.LoaderID("loader-id"),
		Timestamp: network.MonotonicTime(1),
		Type:      page.ResourceType("resource-type"),
		Response:  &network.Response{},
		FrameID:   page.FrameID("frame-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.responseReceived",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.ResponseReceivedEvent)
	mockSocket.Network().OnResponseReceived(func(eventData *network.ResponseReceivedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.responseReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketClosed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.WebSocketClosedEvent)
	mockSocket.Network().OnWebSocketClosed(func(eventData *network.WebSocketClosedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.WebSocketClosedEvent{
		RequestID: network.RequestID("request-id"),
		Timestamp: network.MonotonicTime(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.webSocketClosed",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.WebSocketClosedEvent)
	mockSocket.Network().OnWebSocketClosed(func(eventData *network.WebSocketClosedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.webSocketClosed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketCreated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.WebSocketCreatedEvent)
	mockSocket.Network().OnWebSocketCreated(func(eventData *network.WebSocketCreatedEvent) {
		resultChan <- eventData
	})
	mockResult := &network.WebSocketCreatedEvent{
		RequestID:    network.RequestID("request-id"),
		Timestamp:    network.MonotonicTime(1),
		ErrorMessage: "error message",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.webSocketCreated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.WebSocketCreatedEvent)
	mockSocket.Network().OnWebSocketCreated(func(eventData *network.WebSocketCreatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.webSocketCreated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketFrameError(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.WebSocketFrameErrorEvent)
	mockSocket.Network().OnWebSocketFrameError(func(eventData *network.WebSocketFrameErrorEvent) {
		resultChan <- eventData
	})
	mockResult := &network.WebSocketFrameErrorEvent{
		RequestID:    network.RequestID("request-id"),
		Timestamp:    network.MonotonicTime(1),
		ErrorMessage: "error message",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.webSocketFrameError",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.WebSocketFrameErrorEvent)
	mockSocket.Network().OnWebSocketFrameError(func(eventData *network.WebSocketFrameErrorEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.webSocketFrameError",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketFrameReceived(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.WebSocketFrameReceivedEvent)
	mockSocket.Network().OnWebSocketFrameReceived(func(eventData *network.WebSocketFrameReceivedEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.webSocketFrameReceived",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.WebSocketFrameReceivedEvent)
	mockSocket.Network().OnWebSocketFrameReceived(func(eventData *network.WebSocketFrameReceivedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.webSocketFrameReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketFrameSent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.WebSocketFrameSentEvent)
	mockSocket.Network().OnWebSocketFrameSent(func(eventData *network.WebSocketFrameSentEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.webSocketFrameSent",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.WebSocketFrameSentEvent)
	mockSocket.Network().OnWebSocketFrameSent(func(eventData *network.WebSocketFrameSentEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.webSocketFrameSent",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketHandshakeResponseReceived(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.WebSocketHandshakeResponseReceivedEvent)
	mockSocket.Network().OnWebSocketHandshakeResponseReceived(func(eventData *network.WebSocketHandshakeResponseReceivedEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.webSocketHandshakeResponseReceived",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.WebSocketHandshakeResponseReceivedEvent)
	mockSocket.Network().OnWebSocketHandshakeResponseReceived(func(eventData *network.WebSocketHandshakeResponseReceivedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.webSocketHandshakeResponseReceived",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestNetworkOnWebSocketWillSendHandshakeRequest(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *network.WebSocketWillSendHandshakeRequestEvent)
	mockSocket.Network().OnWebSocketWillSendHandshakeRequest(func(eventData *network.WebSocketWillSendHandshakeRequestEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Network.webSocketWillSendHandshakeRequest",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.RequestID != result.RequestID {
		t.Errorf("Expected %s, got %s", mockResult.RequestID, result.RequestID)
	}

	resultChan = make(chan *network.WebSocketWillSendHandshakeRequestEvent)
	mockSocket.Network().OnWebSocketWillSendHandshakeRequest(func(eventData *network.WebSocketWillSendHandshakeRequestEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Network.webSocketWillSendHandshakeRequest",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
