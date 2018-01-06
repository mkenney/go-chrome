package storage

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
ClearDataForOriginResult represents the result of calls to Storage.clearDataForOrigin.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin
*/
type ClearDataForOriginResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
	UsageBreakdown []*UsageForType `json:"usageBreakdown"`

	// Error information related to executing this method
	Err error `json:"-"`
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
TrackCacheStorageForOriginResult represents the result of calls to
Storage.trackCacheStorageForOrigin.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackCacheStorageForOrigin
*/
type TrackCacheStorageForOriginResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
TrackIndexedDBForOriginResult represents the result of calls to Storage.trackIndexedDBForOrigin.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackIndexedDBForOrigin
*/
type TrackIndexedDBForOriginResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
UntrackCacheStorageForOriginResult represents the result of calls to
Storage.untrackCacheStorageForOrigin.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackCacheStorageForOrigin
*/
type UntrackCacheStorageForOriginResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
UntrackIndexedDBForOriginResult represents the result of calls to
Storage.untrackIndexedDBForOrigin.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackIndexedDBForOrigin
*/
type UntrackIndexedDBForOriginResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}
