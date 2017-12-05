package CacheStorage

/*
DeleteCacheParams represents CacheStorage.deleteCache parameters.
*/
type DeleteCacheParams struct {
	// Id of cache for deletion.
	CacheID CacheID `json:"cacheId"`
}

/*
DeleteEntryParams represents CacheStorage.deleteEntry parameters.
*/
type DeleteEntryParams struct {
	// ID of cache where the entry will be deleted.
	CacheID CacheID `json:"cacheId"`

	// URL spec of the request.
	Request string `json:"request"`
}

/*
RequestCacheNamesParams represents CacheStorage.requestCacheNames parameters.
*/
type RequestCacheNamesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

/*
RequestCacheNamesResult represents the result of calls to CacheStorage.requestCacheNames.
*/
type RequestCacheNamesResult struct {
	// Caches for the security origin.
	Caches []Cache `json:"caches"`
}

/*
RequestCachedResponseParams represents CacheStorage.requestCachedResponse parameters.
*/
type RequestCachedResponseParams struct {
	// ID of cache that contains the enty.
	CacheID CacheID `json:"cacheId"`

	// URL spec of the request.
	RequestURL string `json:"requestURL"`
}

/*
RequestCachedResponseResult represents the result of calls to CacheStorage.requestCachedResponse.
*/
type RequestCachedResponseResult struct {
	// Response read from the cache.
	Response CachedResponse `json:"response"`
}

/*
RequestEntriesParams represents CacheStorage.requestEntries parameters.
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
*/
type RequestEntriesResult struct {
	// Array of object store data entries.
	CacheDataEntries []DataEntry `json:"cacheDataEntries"`

	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

/*
CacheID is the unique identifier of the Cache object.
*/
type CacheID string

/*
DataEntry is a data entry.
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
*/
type Header struct {
	// Header name
	Name string `json:"name"`

	// Header value
	Value string `json:"value"`
}

/*
CachedResponse represents a cached response
*/
type CachedResponse struct {
	// Entry content, base64-encoded.
	Body string `json:"body"`
}
