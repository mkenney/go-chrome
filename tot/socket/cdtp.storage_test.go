package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/storage"
)

func TestStorageClearDataForOrigin(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &storage.ClearDataForOriginParams{
		Origin: "origin",
		Types:  "type1,type2",
	}
	mockResult := &storage.ClearDataForOriginResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Storage().ClearDataForOrigin(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Storage().ClearDataForOrigin(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageGetUsageAndQuota(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &storage.GetUsageAndQuotaParams{
		Origin: "origin",
	}
	mockResult := &storage.GetUsageAndQuotaResult{
		Usage: 1,
		Quota: 1,
		UsageBreakdown: []*storage.UsageForType{{
			Type:  storage.Type.Appcache,
			Usage: 1,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Storage().GetUsageAndQuota(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Usage != result.Usage {
		t.Errorf("Expected %d, got %d", mockResult.Usage, result.Usage)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Storage().GetUsageAndQuota(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageTrackCacheStorageForOrigin(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &storage.TrackCacheStorageForOriginParams{
		Origin: "origin",
	}
	mockResult := &storage.TrackCacheStorageForOriginResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Storage().TrackCacheStorageForOrigin(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Storage().TrackCacheStorageForOrigin(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageTrackIndexedDBForOrigin(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &storage.TrackIndexedDBForOriginParams{
		Origin: "origin",
	}
	mockResult := &storage.TrackIndexedDBForOriginResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Storage().TrackIndexedDBForOrigin(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Storage().TrackIndexedDBForOrigin(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageUntrackCacheStorageForOrigin(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &storage.UntrackCacheStorageForOriginParams{
		Origin: "origin",
	}
	mockResult := &storage.UntrackCacheStorageForOriginResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Storage().UntrackCacheStorageForOrigin(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Storage().UntrackCacheStorageForOrigin(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageUntrackIndexedDBForOrigin(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &storage.UntrackIndexedDBForOriginParams{
		Origin: "origin",
	}
	mockResult := &storage.UntrackIndexedDBForOriginResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Storage().UntrackIndexedDBForOrigin(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Storage().UntrackIndexedDBForOrigin(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageOnCacheStorageContentUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.CacheStorageContentUpdatedEvent)
	soc.Storage().OnCacheStorageContentUpdated(func(eventData *storage.CacheStorageContentUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &storage.CacheStorageContentUpdatedEvent{
		Origin:    "origin",
		CacheName: "cache-name",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Storage.cacheStorageContentUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Storage.cacheStorageContentUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageOnCacheStorageListUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.CacheStorageListUpdatedEvent)
	soc.Storage().OnCacheStorageListUpdated(func(eventData *storage.CacheStorageListUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &storage.CacheStorageListUpdatedEvent{
		Origin: "origin",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Storage.cacheStorageListUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Storage.cacheStorageListUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageOnIndexedDBContentUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.IndexedDBContentUpdatedEvent)
	soc.Storage().OnIndexedDBContentUpdated(func(eventData *storage.IndexedDBContentUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &storage.IndexedDBContentUpdatedEvent{
		Origin:          "origin",
		DatabaseName:    "dbname",
		ObjectStoreName: "storename",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Storage.indexedDBContentUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Storage.indexedDBContentUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageOnIndexedDBListUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.IndexedDBListUpdatedEvent)
	soc.Storage().OnIndexedDBListUpdated(func(eventData *storage.IndexedDBListUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &storage.IndexedDBListUpdatedEvent{
		Origin: "origin",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Storage.indexedDBListUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Storage.indexedDBListUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
