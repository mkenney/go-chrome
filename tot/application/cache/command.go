package cache

import (
	"github.com/mkenney/go-chrome/tot/page"
)

/*
EnableResult represents the result of calls to ApplicationCache.enable.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetForFrameParams represents ApplicationCache.getApplicationCacheForFrame parameters

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
type GetForFrameParams struct {
	// Identifier of the frame containing document whose application cache is
	// retrieved.
	FrameID page.FrameID `json:"frameId"`
}

/*
GetForFrameResult represents the result of calls to ApplicationCache.getApplicationCacheForFrame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
type GetForFrameResult struct {
	// Relevant application cache data for the document in given frame.
	ApplicationCache *ApplicationCache `json:"applicationCache"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetFramesWithManifestsResult represents the result of calls to
ApplicationCache.getFramesWithManifests.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
type GetFramesWithManifestsResult struct {
	// Array of frame identifiers with manifest urls for each frame containing a
	// document associated with some application cache.
	FrameIDs []*FrameWithManifest `json:"frameIds"`

	// Error information related to executing this method
	Err error `json:"-"`
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

	// Error information related to executing this method
	Err error `json:"-"`
}
