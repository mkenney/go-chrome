package socket

import (
	"testing"

	cacheStorage "github.com/mkenney/go-chrome/tot/cache/storage"
)

func TestCacheStorageDeleteCache(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CacheStorage().DeleteCache(&cacheStorage.DeleteCacheParams{
		CacheID: cacheStorage.CacheID("cache-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CacheStorage().DeleteCache(&cacheStorage.DeleteCacheParams{
		CacheID: cacheStorage.CacheID("cache-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCacheStorageDeleteEntry(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}
	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CacheStorage().DeleteEntry(&cacheStorage.DeleteEntryParams{
		CacheID: cacheStorage.CacheID("cache-id"),
		Request: "request",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CacheStorage().DeleteEntry(&cacheStorage.DeleteEntryParams{
		CacheID: cacheStorage.CacheID("cache-id"),
		Request: "request",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCacheStorageRequestCacheNames(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &cacheStorage.RequestCacheNamesResult{
		Caches: []*cacheStorage.Cache{{
			CacheID:        cacheStorage.CacheID("cache-id"),
			SecurityOrigin: "security-origin",
			CacheName:      "cache-name",
		}},
	}

	chrome.AddData(MockData{
		0,
		&Error{},
		mockResult,
		"",
	})
	result := <-soc.CacheStorage().RequestCacheNames(&cacheStorage.RequestCacheNamesParams{
		SecurityOrigin: "security-origin",
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CacheStorage().RequestCacheNames(&cacheStorage.RequestCacheNamesParams{
		SecurityOrigin: "security-origin",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCacheStorageRequestCachedResponse(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &cacheStorage.RequestCachedResponseResult{
		Response: &cacheStorage.CachedResponse{
			Body: "body",
		},
	}

	chrome.AddData(MockData{
		0,
		&Error{},
		mockResult,
		"",
	})
	result := <-soc.CacheStorage().RequestCachedResponse(&cacheStorage.RequestCachedResponseParams{
		CacheID:    cacheStorage.CacheID("security-origin"),
		RequestURL: soc.URL().String(),
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CacheStorage().RequestCachedResponse(&cacheStorage.RequestCachedResponseParams{
		CacheID:    cacheStorage.CacheID("security-origin"),
		RequestURL: soc.URL().String(),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCacheStorageRequestEntries(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &cacheStorage.RequestEntriesResult{
		CacheDataEntries: []*cacheStorage.DataEntry{{
			RequestURL:    soc.URL().String(),
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

	chrome.AddData(MockData{
		0,
		&Error{},
		mockResult,
		"",
	})
	result := <-soc.CacheStorage().RequestEntries(&cacheStorage.RequestEntriesParams{
		CacheID:   cacheStorage.CacheID("security-origin"),
		SkipCount: 1,
		PageSize:  1,
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CacheStorage().RequestEntries(&cacheStorage.RequestEntriesParams{
		CacheID:   cacheStorage.CacheID("security-origin"),
		SkipCount: 1,
		PageSize:  1,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
