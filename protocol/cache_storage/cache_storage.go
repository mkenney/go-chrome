package CacheStorage

/*
CacheID is the unique identifier of the Cache object.
*/
type CacheID string

/*
CacheDataEntry is a data entry.
*/
type CacheDataEntry struct {
	// Request URL.
	RequestURL string `json:"requestURL"`

	// Request method.
	RequestMethod string `json:"requestMethod"`

	// Request headers.
	RequestHeaders []*Header `json:"requestHeaders"`

	// Number of seconds since epoch.
	ResponseTime int `json:"responseTime"`

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
