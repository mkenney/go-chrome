package cache

import (
	"github.com/mkenney/go-chrome/tot/page"
)

/*
StatusUpdatedEvent represents ApplicationCache.applicationCacheStatusUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
*/
type StatusUpdatedEvent struct {
	// Identifier of the frame containing document whose application cache
	// updated status.
	FrameID page.FrameID `json:"frameId"`

	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Updated application cache status.
	Status int `json:"status"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
NetworkStateUpdatedEvent represents ApplicationCache.applicationCachenetworkStateUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
*/
type NetworkStateUpdatedEvent struct {
	IsNowOnline bool `json:"isNowOnline"`

	// Error information related to this event
	Err error `json:"-"`
}
