package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	storage "github.com/mkenney/go-chrome/tot/cdtp/storage"
	"github.com/mkenney/go-chrome/tot/socket"
)

func TestStorageClearDataForOrigin(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &storage.ClearDataForOriginParams{
		Origin: "origin",
		Types:  "type1,type2",
	}
	resultChan := mockSocket.Storage().ClearDataForOrigin(params)
	mockResult := &storage.ClearDataForOriginResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Storage().ClearDataForOrigin(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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

func TestStorageGetUsageAndQuota(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &storage.GetUsageAndQuotaParams{
		Origin: "origin",
	}
	resultChan := mockSocket.Storage().GetUsageAndQuota(params)
	mockResult := &storage.GetUsageAndQuotaResult{
		Usage: 1,
		Quota: 1,
		UsageBreakdown: []*storage.UsageForType{{
			Type:  storage.Type.Appcache,
			Usage: 1,
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Usage != result.Usage {
		t.Errorf("Expected %d, got %d", mockResult.Usage, result.Usage)
	}

	resultChan = mockSocket.Storage().GetUsageAndQuota(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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

func TestStorageTrackCacheStorageForOrigin(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &storage.TrackCacheStorageForOriginParams{
		Origin: "origin",
	}
	resultChan := mockSocket.Storage().TrackCacheStorageForOrigin(params)
	mockResult := &storage.TrackCacheStorageForOriginResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Storage().TrackCacheStorageForOrigin(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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

func TestStorageTrackIndexedDBForOrigin(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &storage.TrackIndexedDBForOriginParams{
		Origin: "origin",
	}
	resultChan := mockSocket.Storage().TrackIndexedDBForOrigin(params)
	mockResult := &storage.TrackIndexedDBForOriginResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Storage().TrackIndexedDBForOrigin(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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

func TestStorageUntrackCacheStorageForOrigin(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &storage.UntrackCacheStorageForOriginParams{
		Origin: "origin",
	}
	resultChan := mockSocket.Storage().UntrackCacheStorageForOrigin(params)
	mockResult := &storage.UntrackCacheStorageForOriginResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Storage().UntrackCacheStorageForOrigin(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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

func TestStorageUntrackIndexedDBForOrigin(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &storage.UntrackIndexedDBForOriginParams{
		Origin: "origin",
	}
	resultChan := mockSocket.Storage().UntrackIndexedDBForOrigin(params)
	mockResult := &storage.UntrackIndexedDBForOriginResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Storage().UntrackIndexedDBForOrigin(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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

func TestStorageOnCacheStorageContentUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.CacheStorageContentUpdatedEvent)
	mockSocket.Storage().OnCacheStorageContentUpdated(func(eventData *storage.CacheStorageContentUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.CacheStorageContentUpdatedEvent{
		Origin:    "origin",
		CacheName: "cache-name",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Storage.cacheStorageContentUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	resultChan = make(chan *storage.CacheStorageContentUpdatedEvent)
	mockSocket.Storage().OnCacheStorageContentUpdated(func(eventData *storage.CacheStorageContentUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Storage.cacheStorageContentUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageOnCacheStorageListUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.CacheStorageListUpdatedEvent)
	mockSocket.Storage().OnCacheStorageListUpdated(func(eventData *storage.CacheStorageListUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.CacheStorageListUpdatedEvent{
		Origin: "origin",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Storage.cacheStorageListUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	resultChan = make(chan *storage.CacheStorageListUpdatedEvent)
	mockSocket.Storage().OnCacheStorageListUpdated(func(eventData *storage.CacheStorageListUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Storage.cacheStorageListUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageOnIndexedDBContentUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.IndexedDBContentUpdatedEvent)
	mockSocket.Storage().OnIndexedDBContentUpdated(func(eventData *storage.IndexedDBContentUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.IndexedDBContentUpdatedEvent{
		Origin:          "origin",
		DatabaseName:    "dbname",
		ObjectStoreName: "storename",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Storage.indexedDBContentUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	resultChan = make(chan *storage.IndexedDBContentUpdatedEvent)
	mockSocket.Storage().OnIndexedDBContentUpdated(func(eventData *storage.IndexedDBContentUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Storage.indexedDBContentUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestStorageOnIndexedDBListUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.IndexedDBListUpdatedEvent)
	mockSocket.Storage().OnIndexedDBListUpdated(func(eventData *storage.IndexedDBListUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.IndexedDBListUpdatedEvent{
		Origin: "origin",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Storage.indexedDBListUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Origin != result.Origin {
		t.Errorf("Expected %s, got %s", mockResult.Origin, result.Origin)
	}

	resultChan = make(chan *storage.IndexedDBListUpdatedEvent)
	mockSocket.Storage().OnIndexedDBListUpdated(func(eventData *storage.IndexedDBListUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Storage.indexedDBListUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
