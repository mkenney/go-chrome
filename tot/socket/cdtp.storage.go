package socket

import (
	"encoding/json"

	storage "github.com/mkenney/go-chrome/tot/cdtp/storage"
)

/*
StorageProtocol provides a namespace for the Chrome Storage protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/ EXPERIMENTAL.
*/
type StorageProtocol struct {
	Socket Socketer
}

/*
ClearDataForOrigin clears storage for origin.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin
*/
func (protocol *StorageProtocol) ClearDataForOrigin(
	params *storage.ClearDataForOriginParams,
) chan *storage.ClearDataForOriginResult {
	resultChan := make(chan *storage.ClearDataForOriginResult)
	command := NewCommand(protocol.Socket, "Storage.clearDataForOrigin", params)
	result := &storage.ClearDataForOriginResult{}

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
GetUsageAndQuota returns usage and quota in bytes.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getUsageAndQuota
*/
func (protocol *StorageProtocol) GetUsageAndQuota(
	params *storage.GetUsageAndQuotaParams,
) chan *storage.GetUsageAndQuotaResult {
	resultChan := make(chan *storage.GetUsageAndQuotaResult)
	command := NewCommand(protocol.Socket, "storage.getUsageAndQuota", params)
	result := &storage.GetUsageAndQuotaResult{}

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
TrackCacheStorageForOrigin registers origin to be notified when an update occurs
to its cache storage list.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackCacheStorageForOrigin
*/
func (protocol *StorageProtocol) TrackCacheStorageForOrigin(
	params *storage.TrackCacheStorageForOriginParams,
) chan *storage.TrackCacheStorageForOriginResult {
	resultChan := make(chan *storage.TrackCacheStorageForOriginResult)
	command := NewCommand(protocol.Socket, "Storage.trackCacheStorageForOrigin", params)
	result := &storage.TrackCacheStorageForOriginResult{}

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
TrackIndexedDBForOrigin registers origin to be notified when an update occurs to
its IndexedDB.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackIndexedDBForOrigin
*/
func (protocol *StorageProtocol) TrackIndexedDBForOrigin(
	params *storage.TrackIndexedDBForOriginParams,
) chan *storage.TrackIndexedDBForOriginResult {
	resultChan := make(chan *storage.TrackIndexedDBForOriginResult)
	command := NewCommand(protocol.Socket, "Storage.trackIndexedDBForOrigin", params)
	result := &storage.TrackIndexedDBForOriginResult{}

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
UntrackCacheStorageForOrigin unregisters origin from receiving notifications for
cache storage.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackCacheStorageForOrigin
*/
func (protocol *StorageProtocol) UntrackCacheStorageForOrigin(
	params *storage.UntrackCacheStorageForOriginParams,
) chan *storage.UntrackCacheStorageForOriginResult {
	resultChan := make(chan *storage.UntrackCacheStorageForOriginResult)
	command := NewCommand(protocol.Socket, "Storage.untrackCacheStorageForOrigin", params)
	result := &storage.UntrackCacheStorageForOriginResult{}

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
UntrackIndexedDBForOrigin unregisters origin from receiving notifications for
IndexedDB.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackIndexedDBForOrigin
*/
func (protocol *StorageProtocol) UntrackIndexedDBForOrigin(
	params *storage.UntrackIndexedDBForOriginParams,
) chan *storage.UntrackIndexedDBForOriginResult {
	resultChan := make(chan *storage.UntrackIndexedDBForOriginResult)
	command := NewCommand(protocol.Socket, "Storage.untrackIndexedDBForOrigin", params)
	result := &storage.UntrackIndexedDBForOriginResult{}

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
OnCacheStorageContentUpdated adds a handler to the Storage.cacheStorageContentUpdated
event. Storage.cacheStorageContentUpdated fires when a cache's contents have
been modified.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageContentUpdated
*/
func (protocol *StorageProtocol) OnCacheStorageContentUpdated(
	callback func(event *storage.CacheStorageContentUpdatedEvent),
) {
	handler := NewEventHandler(
		"Storage.cacheStorageContentUpdated",
		func(response *Response) {
			event := &storage.CacheStorageContentUpdatedEvent{}
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
OnCacheStorageListUpdated adds a handler to the Storage.cacheStorageListUpdated
event. Storage.cacheStorageListUpdated fires when cache has been added/deleted.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageListUpdated
*/
func (protocol *StorageProtocol) OnCacheStorageListUpdated(
	callback func(event *storage.CacheStorageListUpdatedEvent),
) {
	handler := NewEventHandler(
		"Storage.cacheStorageListUpdated",
		func(response *Response) {
			event := &storage.CacheStorageListUpdatedEvent{}
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
OnIndexedDBContentUpdated adds a handler to the Storage.indexedDBContentUpdated
event. Storage.indexedDBContentUpdated fires when the origin's IndexedDB object
store has been modified.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBContentUpdated
*/
func (protocol *StorageProtocol) OnIndexedDBContentUpdated(
	callback func(event *storage.IndexedDBContentUpdatedEvent),
) {
	handler := NewEventHandler(
		"Storage.indexedDBContentUpdated",
		func(response *Response) {
			event := &storage.IndexedDBContentUpdatedEvent{}
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
OnIndexedDBListUpdated adds a handler to the Storage.indexedDBListUpdated event.
Storage.indexedDBListUpdated fires when the origin's IndexedDB database list has
been modified.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBListUpdated
*/
func (protocol *StorageProtocol) OnIndexedDBListUpdated(
	callback func(event *storage.IndexedDBListUpdatedEvent),
) {
	handler := NewEventHandler(
		"Storage.indexedDBListUpdated",
		func(response *Response) {
			event := &storage.IndexedDBListUpdatedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
