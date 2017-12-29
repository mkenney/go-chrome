package protocol

import (
	"encoding/json"

	storage "github.com/mkenney/go-chrome/protocol/storage"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Storage is a struct that provides a namespace for the Chrome Storage protocol methods. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/
*/
var Storage = _storage{}

type _storage struct{}

/*
ClearDataForOrigin clears storage for origin.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin
*/
func (_storage) ClearDataForOrigin(
	socket sock.Socketer,
	params *storage.ClearDataForOriginParams,
) error {
	command := sock.NewCommand("storage.clearDataForOrigin", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetUsageAndQuota returns usage and quota in bytes.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getUsageAndQuota
*/
func (_storage) GetUsageAndQuota(
	socket sock.Socketer,
	params *storage.GetUsageAndQuotaParams,
) (*storage.GetUsageAndQuotaResult, error) {
	command := sock.NewCommand("storage.getUsageAndQuota", params)
	result := &storage.GetUsageAndQuotaResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
TrackCacheStorageForOrigin registers origin to be notified when an update occurs to its cache
storage list.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackCacheStorageForOrigin
*/
func (_storage) TrackCacheStorageForOrigin(
	socket sock.Socketer,
	params *storage.TrackCacheStorageForOriginParams,
) error {
	command := sock.NewCommand("storage.trackCacheStorageForOrigin", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
TrackIndexedDBForOrigin registers origin to be notified when an update occurs to its IndexedDB.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackIndexedDBForOrigin
*/
func (_storage) TrackIndexedDBForOrigin(
	socket sock.Socketer,
	params *storage.TrackIndexedDBForOriginParams,
) error {
	command := sock.NewCommand("storage.trackIndexedDBForOrigin", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
UntrackCacheStorageForOrigin unregisters origin from receiving notifications for cache storage.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackCacheStorageForOrigin
*/
func (_storage) UntrackCacheStorageForOrigin(
	socket sock.Socketer,
	params *storage.UntrackCacheStorageForOriginParams,
) error {
	command := sock.NewCommand("storage.untrackCacheStorageForOrigin", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
UntrackIndexedDBForOrigin unregisters origin from receiving notifications for IndexedDB.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackIndexedDBForOrigin
*/
func (_storage) UntrackIndexedDBForOrigin(
	socket sock.Socketer,
	params *storage.UntrackIndexedDBForOriginParams,
) error {
	command := sock.NewCommand("storage.untrackIndexedDBForOrigin", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnCacheStorageContentUpdated adds a handler to the Storage.cacheStorageContentUpdated event.
Storage.cacheStorageContentUpdated fires when a cache's contents have been modified.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageContentUpdated
*/
func (_storage) OnCacheStorageContentUpdated(
	socket sock.Socketer,
	callback func(event *storage.CacheStorageContentUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"Storage.cacheStorageContentUpdated",
		func(response *sock.Response) {
			event := &storage.CacheStorageContentUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnCacheStorageListUpdated adds a handler to the Storage.cacheStorageListUpdated event.
Storage.cacheStorageListUpdated fires when cache has been added/deleted.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageListUpdated
*/
func (_storage) OnCacheStorageListUpdated(
	socket sock.Socketer,
	callback func(event *storage.CacheStorageListUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"Storage.cacheStorageListUpdated",
		func(response *sock.Response) {
			event := &storage.CacheStorageListUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnIndexedDBContentUpdated adds a handler to the Storage.indexedDBContentUpdated event.
Storage.indexedDBContentUpdated fires when the origin's IndexedDB object store has been modified.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBContentUpdated
*/
func (_storage) OnIndexedDBContentUpdated(
	socket sock.Socketer,
	callback func(event *storage.IndexedDBContentUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"Storage.indexedDBContentUpdated",
		func(response *sock.Response) {
			event := &storage.IndexedDBContentUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnIndexedDBListUpdated adds a handler to the Storage.indexedDBListUpdated event.
Storage.indexedDBListUpdated fires when the origin's IndexedDB database list has been modified.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBListUpdated
*/
func (_storage) OnIndexedDBListUpdated(
	socket sock.Socketer,
	callback func(event *storage.IndexedDBListUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"Storage.indexedDBListUpdated",
		func(response *sock.Response) {
			event := &storage.IndexedDBListUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
