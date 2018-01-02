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
