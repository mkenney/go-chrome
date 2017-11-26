package ApplicationCache

import (
	Page "app/chrome/protocol/page"
)

/*
ApplicationCacheResource contains detailed application cache resource information.
*/
type ApplicationCacheResource struct {
	// Resource URL.
	URL string `json:"url"`

	// Resource size.
	Size int `json:"size"`

	// Resource type.
	Type string `json:"type"`
}

/*
ApplicationCache contains detailed application cache information.
*/
type ApplicationCache struct {
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Application cache size.
	Size int `json:"size"`

	// Application cache creation time.
	CreationTime int `json:"creationTime"`

	// Application cache update time.
	UpdateTime int `json:"updateTime"`

	// Application cache resources.
	Resources []*ApplicationCacheResource `json:"resources"`
}

/*
FrameWithManifest is a frame identifier / manifest URL pair.
*/
type FrameWithManifest struct {
	// Frame identifier.
	FrameID Page.FrameID `json:"frameId"`

	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Application cache status.
	Status int `json:"status"`
}

/*
GetManifestForFrameParams represents ApplicationCache.getFramesWithManifests parameters
*/
type GetManifestForFrameParams struct {
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameID Page.FrameID `json:"frameId"`
}

/*
GetApplicationCacheForFrameParams represents ApplicationCache.getApplicationCacheForFrame parameters
*/
type GetApplicationCacheForFrameParams struct {
	// Identifier of the frame containing document whose application cache is retrieved.
	FrameID Page.FrameID `json:"frameId"`
}

/*
ApplicationCacheStatusUpdatedEvent represents ApplicationCache.applicationCacheStatusUpdated
event data.
*/
type ApplicationCacheStatusUpdatedEvent struct {
	// Identifier of the frame containing document whose application cache updated status.
	FrameID Page.FrameID `json:"frameId"`

	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Updated application cache status.
	Status int `json:"status"`
}
