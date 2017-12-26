/*
Package applicationCache provides type definitions for use with the Chrome ApplicationCache protocol

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/
*/
package applicationCache

import (
	"github.com/mkenney/go-chrome/protocol/page"
)

/*
GetForFrameParams represents ApplicationCache.getApplicationCacheForFrame parameters

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
type GetForFrameParams struct {
	// Identifier of the frame containing document whose application cache is retrieved.
	FrameID page.FrameID `json:"frameId"`
}

/*
GetForFrameResult represents the result of calls to ApplicationCache.getApplicationCacheForFrame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
type GetForFrameResult struct {
	// Relevant application cache data for the document in given frame.
	ApplicationCache ApplicationCache `json:"applicationCache"`
}

/*
GetFramesWithManifestsResult represents the result of calls to
ApplicationCache.getFramesWithManifests.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
type GetFramesWithManifestsResult struct {
	// Array of frame identifiers with manifest urls for each frame containing a document associated
	// with some application cache.
	FrameIDs []FrameWithManifest `json:"frameIds"`
}

/*
GetManifestForFrameParams represents ApplicationCache.getFramesWithManifests parameters

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
type GetManifestForFrameParams struct {
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameID page.FrameID `json:"frameId"`
}

/*
GetManifestForFrameResult represents the result of calls to ApplicationCache.getManifestForFrame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
type GetManifestForFrameResult struct {
	// Manifest URL for document in the given frame.
	ManifestURL string `json:"manifestURL"`
}

/*
StatusUpdatedEvent represents ApplicationCache.applicationCacheStatusUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
*/
type StatusUpdatedEvent struct {
	// Identifier of the frame containing document whose application cache updated status.
	FrameID page.FrameID `json:"frameId"`

	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Updated application cache status.
	Status int `json:"status"`
}

/*
NetworkStateUpdatedEvent represents ApplicationCache.applicationCachenetworkStateUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
*/
type NetworkStateUpdatedEvent struct {
	IsNowOnline bool `json:"isNowOnline"`
}

/*
Resource contains detailed application cache resource information.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-ApplicationCacheResource
*/
type Resource struct {
	// Resource URL.
	URL string `json:"url"`

	// Resource size.
	Size int `json:"size"`

	// Resource type.
	Type string `json:"type"`
}

/*
ApplicationCache contains detailed application cache information.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-ApplicationCache
*/
type ApplicationCache struct {
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Application cache size.
	Size float64 `json:"size"`

	// Application cache creation time.
	CreationTime float64 `json:"creationTime"`

	// Application cache update time.
	UpdateTime float64 `json:"updateTime"`

	// Application cache resources.
	Resources []Resource `json:"resources"`
}

/*
FrameWithManifest is a frame identifier / manifest URL pair.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-FrameWithManifest
*/
type FrameWithManifest struct {
	// Frame identifier.
	FrameID page.FrameID `json:"frameId"`

	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Application cache status.
	Status int `json:"status"`
}
