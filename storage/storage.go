package Storage

import (
	"fmt"
)

/*
ClearDataForOriginParams represents Storage.clearDataForOrigin parameters.
*/
type ClearDataForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`

	// Comma separated origin names.
	StorageTypes string `json:"storageTypes"`
}

/*
GetUsageAndQuotaParams represents Storage.getUsageAndQuota parameters.
*/
type GetUsageAndQuotaParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
GetUsageAndQuotaResult represents the result of calls to Storage.getUsageAndQuota.
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
*/
type TrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
TrackIndexedDBForOriginParams represents Storage.trackIndexedDBForOrigin parameters.
*/
type TrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
UntrackCacheStorageForOriginParams represents Storage.untrackCacheStorageForOrigin parameters.
*/
type UntrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
UntrackIndexedDBForOriginParams represents Storage.untrackIndexedDBForOrigin parameters.
*/
type UntrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

/*
cacheStorageContentUpdatedEvent represents Storage.cacheStorageContentUpdated event data.
*/
type cacheStorageContentUpdatedEvent struct {
	// Origin to update.
	Origin string `json:"origin"`

	// Name of cache in origin.
	CacheName string `json:"cacheName"`
}

/*
CacheStorageListUpdatedEvent represents Storage.cacheStorageListUpdated event data.
*/
type CacheStorageListUpdatedEvent struct {
	// Origin to update.
	Origin string `json:"origin"`
}

/*
IndexedDBContentUpdatedEvent represents Storage.indexedDBContentUpdated event data.
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
*/
type IndexedDBListUpdatedEvent struct {
	// Origin to update.
	Origin string `json:"origin"`
}

/*
StorageType is an enum of possible storage types.
*/
type StorageType int

const (
	_appcache StorageType = iota
	_cookies
	_fileSystems
	_indexeddb
	_localStorage
	_shaderCache
	_websql
	_serviceWorkers
	_cacheStorage
	_all
	_other
)

func (a StorageType) String() string {
	if a == 0 {
		return "appcache"
	}
	if a == 1 {
		return "cookies"
	}
	if a == 2 {
		return "file_systems"
	}
	if a == 3 {
		return "indexeddb"
	}
	if a == 4 {
		return "local_storage"
	}
	if a == 5 {
		return "shader_cache"
	}
	if a == 6 {
		return "websql"
	}
	if a == 7 {
		return "service_workers"
	}
	if a == 8 {
		return "cache_storage"
	}
	if a == 9 {
		return "all"
	}
	if a == 10 {
		return "other"
	}
	panic(fmt.Errorf("Invalid StorageType %d", a))
}

/*
UsageForType is usage for a storage type.
*/
type UsageForType struct {
	// Name of storage type.
	StorageType StorageType `json:"storageType"`

	// Storage usage (bytes).
	Usage int `json:"usage"`
}
