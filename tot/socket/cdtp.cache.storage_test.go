package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	cacheStorage "github.com/mkenney/go-chrome/tot/cache/storage"
)

func TestCacheStorageDeleteCache(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestCacheStorageDeleteCache")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().DeleteCache(&cacheStorage.DeleteCacheParams{
		CacheID: cacheStorage.CacheID("cache-id"),
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:    mockSocket.CurCommandID(),
		Error: &Error{},
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.CacheStorage().DeleteCache(&cacheStorage.DeleteCacheParams{
		CacheID: cacheStorage.CacheID("cache-id"),
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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

func TestCacheStorageDeleteEntry(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestCacheStorageDeleteEntry")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().DeleteEntry(&cacheStorage.DeleteEntryParams{
		CacheID: cacheStorage.CacheID("cache-id"),
		Request: "request",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:    mockSocket.CurCommandID(),
		Error: &Error{},
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.CacheStorage().DeleteEntry(&cacheStorage.DeleteEntryParams{
		CacheID: cacheStorage.CacheID("cache-id"),
		Request: "request",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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

func TestCacheStorageRequestCacheNames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestCacheStorageRequestCacheNames")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().RequestCacheNames(&cacheStorage.RequestCacheNamesParams{
		SecurityOrigin: "security-origin",
	})
	mockResult := &cacheStorage.RequestCacheNamesResult{
		Caches: []*cacheStorage.Cache{{
			CacheID:        cacheStorage.CacheID("cache-id"),
			SecurityOrigin: "security-origin",
			CacheName:      "cache-name",
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.Caches[0].CacheID != mockResult.Caches[0].CacheID {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.Caches[0].CacheID,
			result.Caches[0].CacheID,
		)
	}

	resultChan = mockSocket.CacheStorage().RequestCacheNames(&cacheStorage.RequestCacheNamesParams{
		SecurityOrigin: "security-origin",
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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

func TestCacheStorageRequestCachedResponse(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestCacheStorageRequestCachedResponse")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().RequestCachedResponse(&cacheStorage.RequestCachedResponseParams{
		CacheID:    cacheStorage.CacheID("security-origin"),
		RequestURL: mockSocket.URL().String(),
	})
	mockResult := &cacheStorage.RequestCachedResponseResult{
		Response: &cacheStorage.CachedResponse{
			Body: "body",
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.Response.Body != mockResult.Response.Body {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.Response.Body,
			result.Response.Body,
		)
	}

	resultChan = mockSocket.CacheStorage().RequestCachedResponse(&cacheStorage.RequestCachedResponseParams{
		CacheID:    cacheStorage.CacheID("security-origin"),
		RequestURL: mockSocket.URL().String(),
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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

func TestCacheStorageRequestEntries(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestCacheStorageRequestEntries")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().RequestEntries(&cacheStorage.RequestEntriesParams{
		CacheID:   cacheStorage.CacheID("security-origin"),
		SkipCount: 1,
		PageSize:  1,
	})
	mockResult := &cacheStorage.RequestEntriesResult{
		CacheDataEntries: []*cacheStorage.DataEntry{{
			RequestURL:    mockSocket.URL().String(),
			RequestMethod: "POST",
			RequestHeaders: []*cacheStorage.Header{{
				Name:  "Header",
				Value: "Value",
			}},
			ResponseTime:       1,
			ResponseStatus:     1,
			ResponseStatusText: "response text",
			ResponseHeaders: []*cacheStorage.Header{{
				Name:  "Header",
				Value: "Value",
			}},
		}},
		HasMore: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.CacheDataEntries[0].RequestURL != mockResult.CacheDataEntries[0].RequestURL {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.CacheDataEntries[0].RequestURL,
			result.CacheDataEntries[0].RequestURL,
		)
	}

	resultChan = mockSocket.CacheStorage().RequestEntries(&cacheStorage.RequestEntriesParams{
		CacheID:   cacheStorage.CacheID("security-origin"),
		SkipCount: 1,
		PageSize:  1,
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
