package socket

import (
	"encoding/json"

	storage "github.com/mkenney/go-chrome/protocol/storage"
	log "github.com/sirupsen/logrus"
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
) error {
	command := NewCommand("storage.clearDataForOrigin", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetUsageAndQuota returns usage and quota in bytes.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getUsageAndQuota
*/
func (protocol *StorageProtocol) GetUsageAndQuota(
	params *storage.GetUsageAndQuotaParams,
) (*storage.GetUsageAndQuotaResult, error) {
	command := NewCommand("storage.getUsageAndQuota", params)
	result := &storage.GetUsageAndQuotaResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
TrackCacheStorageForOrigin registers origin to be notified when an update occurs
to its cache storage list.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackCacheStorageForOrigin
*/
func (protocol *StorageProtocol) TrackCacheStorageForOrigin(
	params *storage.TrackCacheStorageForOriginParams,
) error {
	command := NewCommand("storage.trackCacheStorageForOrigin", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
TrackIndexedDBForOrigin registers origin to be notified when an update occurs to
its IndexedDB.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackIndexedDBForOrigin
*/
func (protocol *StorageProtocol) TrackIndexedDBForOrigin(
	params *storage.TrackIndexedDBForOriginParams,
) error {
	command := NewCommand("storage.trackIndexedDBForOrigin", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
UntrackCacheStorageForOrigin unregisters origin from receiving notifications for
cache storage.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackCacheStorageForOrigin
*/
func (protocol *StorageProtocol) UntrackCacheStorageForOrigin(
	params *storage.UntrackCacheStorageForOriginParams,
) error {
	command := NewCommand("storage.untrackCacheStorageForOrigin", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
UntrackIndexedDBForOrigin unregisters origin from receiving notifications for
IndexedDB.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackIndexedDBForOrigin
*/
func (protocol *StorageProtocol) UntrackIndexedDBForOrigin(
	params *storage.UntrackIndexedDBForOriginParams,
) error {
	command := NewCommand("storage.untrackIndexedDBForOrigin", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
