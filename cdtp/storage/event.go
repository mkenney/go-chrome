package storage

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
