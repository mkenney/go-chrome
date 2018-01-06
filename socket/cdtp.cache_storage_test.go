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

	resultChan := mockSocket.CacheStorage().DeleteCache(&cacheStorage.DeleteCacheParams{
		CacheID: cacheStorage.CacheID("cache-id"),
	})
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"CacheStorage.deleteCache",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestCacheStorageDeleteEntry(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().DeleteEntry(&cacheStorage.DeleteEntryParams{
		CacheID: cacheStorage.CacheID("cache-id"),
		Request: "request",
	})
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"CacheStorage.deleteEntry",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestCacheStorageRequestCacheNames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().RequestCacheNames(&cacheStorage.RequestCacheNamesParams{
		SecurityOrigin: "security-origin",
	})
	mockResult := &cacheStorage.RequestCacheNamesResult{
		Caches: []*cacheStorage.Cache{&cacheStorage.Cache{
			CacheID:        cacheStorage.CacheID("cache-id"),
			SecurityOrigin: "security-origin",
			CacheName:      "cache-name",
		}},
	}
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"CacheStorage.requestCacheNames",
		mockResult,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
	if result.Caches[0].CacheID != mockResult.Caches[0].CacheID {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.Caches[0].CacheID,
			result.Caches[0].CacheID,
		)
	}
}

func TestCacheStorageRequestCachedResponse(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
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
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"CacheStorage.requestCachedResponse",
		mockResult,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
	if result.Response.Body != mockResult.Response.Body {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.Response.Body,
			result.Response.Body,
		)
	}
}

func TestCacheStorageRequestEntries(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CacheStorage().RequestEntries(&cacheStorage.RequestEntriesParams{
		CacheID:   cacheStorage.CacheID("security-origin"),
		SkipCount: 1,
		PageSize:  1,
	})
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
		mockSocket.CurCommandID(),
		&Error{},
		"CacheStorage.RequestEntries",
		mockResult,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
	if result.CacheDataEntries[0].RequestURL != mockResult.CacheDataEntries[0].RequestURL {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.CacheDataEntries[0].RequestURL,
			result.CacheDataEntries[0].RequestURL,
		)
	}
}
