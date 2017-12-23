package CacheStorage

/*
DeleteCacheParams represents CacheStorage.deleteCache parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
*/
type DeleteCacheParams struct {
	// Id of cache for deletion.
	CacheID CacheID `json:"cacheId"`
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
	Caches []Cache `json:"caches"`
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
	Response CachedResponse `json:"response"`
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
	CacheDataEntries []DataEntry `json:"cacheDataEntries"`

	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

/*
CacheID is the unique identifier of the Cache object.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-CacheId
*/
type CacheID string

/*
DataEntry is a data entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-DataEntry
*/
type DataEntry struct {
	// Request URL.
	RequestURL string `json:"requestURL"`

	// Request method.
	RequestMethod string `json:"requestMethod"`

	// Request headers.
	RequestHeaders []*Header `json:"requestHeaders"`

	// Number of seconds since epoch.
	ResponseTime float64 `json:"responseTime"`

	// HTTP response status code.
	ResponseStatus int `json:"responseStatus"`

	// HTTP response status text.
	ResponseStatusText string `json:"responseStatusText"`

	// Response headers.
	ResponseHeaders []*Header `json:"responseHeaders"`
}

/*
Cache is a cache identifier.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-Cache
*/
type Cache struct {
	// An opaque unique ID of the cache.
	CacheID CacheID `json:"cacheId"`

	// Security origin of the cache.
	SecurityOrigin string `json:"securityOrigin"`

	// The name of the cache.
	CacheName string `json:"cacheName"`
}

/*
Header is a single header value

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-Header
*/
type Header struct {
	// Header name
	Name string `json:"name"`

	// Header value
	Value string `json:"value"`
}

/*
CachedResponse represents a cached response

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-CachedResponse
*/
type CachedResponse struct {
	// Entry content, base64-encoded.
	Body string `json:"body"`
}
