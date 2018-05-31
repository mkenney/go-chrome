/*
Package storage provides type definitions for use with the Chrome CacheStorage protocol

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/
*/
package storage

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
