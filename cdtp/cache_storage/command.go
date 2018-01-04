package cacheStorage

/*
DeleteCacheParams represents CacheStorage.deleteCache parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
*/
type DeleteCacheParams struct {
	// Id of cache for deletion.
	CacheID CacheID `json:"cacheId"`
}

/*
DeleteCacheResult represents the result of calls to CacheStorage.deleteCache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
*/
type DeleteCacheResult struct {
	CDTPError error `json:"-"`
}

/*
DeleteEntryParams represents CacheStorage.deleteEntry parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
*/
type DeleteEntryParams struct {
	// ID of cache where the entry will be deleted.
	CacheID CacheID `json:"cacheId"`

	// URL spec of the request.
	Request string `json:"request"`
}

/*
DeleteEntryResult represents the result of calls to CacheStorage.deleteEntry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
*/
type DeleteEntryResult struct {
	CDTPError error `json:"-"`
}

/*
RequestCacheNamesParams represents CacheStorage.requestCacheNames parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
*/
type RequestCacheNamesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

/*
RequestCacheNamesResult represents the result of calls to CacheStorage.requestCacheNames.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
*/
type RequestCacheNamesResult struct {
	// Caches for the security origin.
	Caches []*Cache `json:"caches"`

	CDTPError error `json:"-"`
}

/*
RequestCachedResponseParams represents CacheStorage.requestCachedResponse parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
*/
type RequestCachedResponseParams struct {
	// ID of cache that contains the enty.
	CacheID CacheID `json:"cacheId"`

	// URL spec of the request.
	RequestURL string `json:"requestURL"`
}

/*
RequestCachedResponseResult represents the result of calls to CacheStorage.requestCachedResponse.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
*/
type RequestCachedResponseResult struct {
	// Response read from the cache.
	Response *CachedResponse `json:"response"`

	CDTPError error `json:"-"`
}

/*
RequestEntriesParams represents CacheStorage.requestEntries parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
*/
type RequestEntriesParams struct {
	// ID of cache to get entries from.
	CacheID CacheID `json:"cacheId"`

	// Number of records to skip.
	SkipCount int `json:"skipCount"`

	// Number of records to fetch.
	PageSize int `json:"pageSize"`
}

/*
RequestEntriesResult represents the result of calls to CacheStorage.requestEntries.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
*/
type RequestEntriesResult struct {
	// Array of object store data entries.
	CacheDataEntries []*DataEntry `json:"cacheDataEntries"`

	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`

	CDTPError error `json:"-"`
}
