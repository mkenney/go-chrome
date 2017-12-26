/*
Package Storage provides type definitions for use with the Chrome Storage protocol

https://chromedevtools.github.io/devtools-protocol/tot/Storage/
*/
package Storage

/*
ClearDataForOriginParams represents Storage.clearDataForOrigin parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin
*/
type ClearDataForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`

	// Comma separated origin names.
	Types string `json:"storageTypes"`
}

/*
GetUsageAndQuotaParams represents Storage.getUsageAndQuota parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getUsageAndQuota
*/
type GetUsageAndQuotaParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
GetUsageAndQuotaResult represents the result of calls to Storage.getUsageAndQuota.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getUsageAndQuota
*/
type GetUsageAndQuotaResult struct {
	// Storage usage (bytes).
	Usage float64 `json:"usage"`

	// Storage quota (bytes).
	Quota float64 `json:"quota"`

	// Storage usage per type (bytes).
	UsageBreakdown []UsageForType `json:"usageBreakdown"`
}

/*
TrackCacheStorageForOriginParams represents Storage.trackCacheStorageForOrigin parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackCacheStorageForOrigin
*/
type TrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
TrackIndexedDBForOriginParams represents Storage.trackIndexedDBForOrigin parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackIndexedDBForOrigin
*/
type TrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
UntrackCacheStorageForOriginParams represents Storage.untrackCacheStorageForOrigin parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackCacheStorageForOrigin
*/
type UntrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
UntrackIndexedDBForOriginParams represents Storage.untrackIndexedDBForOrigin parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackIndexedDBForOrigin
*/
type UntrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
CacheStorageContentUpdatedEvent represents Storage.cacheStorageContentUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageContentUpdated
*/
type CacheStorageContentUpdatedEvent struct {
	// Origin to update.
	Origin string `json:"origin"`

	// Name of cache in origin.
	CacheName string `json:"cacheName"`
}

/*
CacheStorageListUpdatedEvent represents Storage.cacheStorageListUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageListUpdated
*/
type CacheStorageListUpdatedEvent struct {
	// Origin to update.
	Origin string `json:"origin"`
}

/*
IndexedDBContentUpdatedEvent represents Storage.indexedDBContentUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBContentUpdated
*/
type IndexedDBContentUpdatedEvent struct {
	// Origin to update.
	Origin string `json:"origin"`

	// Database to update.
	DatabaseName string `json:"databaseName"`

	// ObjectStore to update.
	ObjectStoreName string `json:"objectStoreName"`
}

/*
IndexedDBListUpdatedEvent represents Storage.indexedDBListUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBListUpdated
*/
type IndexedDBListUpdatedEvent struct {
	// Origin to update.
	Origin string `json:"origin"`
}

/*
Type is an enum of possible storage types.

ALLOWED VALUES
	- appcache
	- cookies
	- file_systems
	- indexeddb
	- local_storage
	- shader_cache
	- websql
	- service_workers
	- cache_storage
	- all
	- other

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-StorageType
*/
type Type string

/*
UsageForType is usage for a storage type.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-UsageForType
*/
type UsageForType struct {
	// Name of storage type.
	Type Type `json:"storageType"`

	// Storage usage (bytes).
	Usage int `json:"usage"`
}
