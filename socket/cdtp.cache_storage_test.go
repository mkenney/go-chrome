package socket

import (
	"net/url"
	"testing"

	cacheStorage "github.com/mkenney/go-chrome/cdtp/cache_storage"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, err := log.ParseLevel("debug")
	if nil == err {
		log.SetLevel(level)
	}
}

func TestCacheStorageDeleteCache(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"CacheStorage.deleteCache",
		nil,
	)
	go func() {
		err := mockSocket.CacheStorage().DeleteCache(&cacheStorage.DeleteCacheParams{
			CacheID: cacheStorage.CacheID("cache-id"),
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
	}()
}

func TestCacheStorageDeleteEntry(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"CacheStorage.deleteEntry",
		nil,
	)
	go func() {
		err := mockSocket.CacheStorage().DeleteEntry(&cacheStorage.DeleteEntryParams{
			CacheID: cacheStorage.CacheID("cache-id"),
			Request: "request",
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
	}()
}

func TestCacheStorageRequestCacheNames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockResult := &cacheStorage.RequestCacheNamesResult{
		Caches: []*cacheStorage.Cache{&cacheStorage.Cache{
			CacheID:        cacheStorage.CacheID("cache-id"),
			SecurityOrigin: "security-origin",
			CacheName:      "cache-name",
		}},
	}
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"CacheStorage.requestCacheNames",
		mockResult,
	)
	go func() {
		result, err := mockSocket.CacheStorage().RequestCacheNames(&cacheStorage.RequestCacheNamesParams{
			SecurityOrigin: "security-origin",
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.Caches[0].CacheID != mockResult.Caches[0].CacheID {
			t.Errorf(
				"Expected %s, received %s",
				mockResult.Caches[0].CacheID,
				result.Caches[0].CacheID,
			)
		}
	}()
}

func TestCacheStorageRequestCachedResponse(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockResult := &cacheStorage.RequestCachedResponseResult{
		Response: &cacheStorage.CachedResponse{
			Body: "body",
		},
	}
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"CacheStorage.requestCachedResponse",
		mockResult,
	)
	go func() {
		result, err := mockSocket.CacheStorage().RequestCachedResponse(&cacheStorage.RequestCachedResponseParams{
			CacheID:    cacheStorage.CacheID("security-origin"),
			RequestURL: mockSocket.URL().String(),
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.Response != mockResult.Response {
			t.Errorf(
				"Expected %s, received %s",
				mockResult.Response,
				result.Response,
			)
		}
	}()
}

func TestCacheStorageRequestEntries(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockResult := &cacheStorage.RequestEntriesResult{
		CacheDataEntries: []*cacheStorage.DataEntry{&cacheStorage.DataEntry{
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
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"CacheStorage.RequestEntries",
		mockResult,
	)
	//go func() {
	result, err := mockSocket.CacheStorage().RequestEntries(&cacheStorage.RequestEntriesParams{
		CacheID:   cacheStorage.CacheID("security-origin"),
		SkipCount: 1,
		PageSize:  1,
	})
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}
	if result.CacheDataEntries[0].RequestURL != mockResult.CacheDataEntries[0].RequestURL {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.CacheDataEntries[0].RequestURL,
			result.CacheDataEntries[0].RequestURL,
		)
	}
	//}()
}
