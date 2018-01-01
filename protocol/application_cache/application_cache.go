/*
Package applicationCache provides type definitions for use with the Chrome
ApplicationCache protocol

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/
*/
package applicationCache

import (
	"github.com/mkenney/go-chrome/protocol/page"
)

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
	Resources []*Resource `json:"resources"`
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
