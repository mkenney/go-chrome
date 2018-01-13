/*
Package page provides type definitions for use with the Chrome Page protocol

https://chromedevtools.github.io/devtools-protocol/tot/Page/
*/
package page

/*
LoaderID is the Unique loader identifier.
This is a duplicate of Network.LoaderID to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoaderId
*/
type LoaderID string

/*
MonotonicTime is the monotonically increasing time in seconds since an arbitrary point in the past.
This is a duplicate of Network.MonotonicTime to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-MonotonicTime
*/
type MonotonicTime int

/*
Rect defines a rectangle.
This is a duplicate of DOM.Rect to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Rect
*/
type Rect struct {
	// X coordinate.
	X float64 `json:"x"`

	// Y coordinate.
	Y float64 `json:"y"`

	// Rectangle width.
	Width float64 `json:"width"`

	// Rectangle height.
	Height float64 `json:"height"`
}

/*
TimeSinceEpoch represents UTC time in seconds, counted from January 1, 1970.
This is a duplicate of DOM.Rect to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TimeSinceEpoch
*/
type TimeSinceEpoch int

/*
AppManifestError defines an error that occurs while parsing an app manifest.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AppManifestError
*/
type AppManifestError struct {
	// Error message.
	Message string `json:"message"`

	// If criticial, this is a non-recoverable parse error.
	Critical int `json:"critical"`

	// Error line.
	Line int `json:"line"`

	// Error column.
	Column int `json:"column"`
}

/*
DialogType defines the Javascript dialog type.

ALLOWED VALUES
	- alert
	- confirm
	- prompt
	- beforeunload

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-DialogType
*/
type DialogType string

/*
Frame details information about the Frame on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Frame
*/
type Frame struct {
	// Frame unique identifier.
	ID string `json:"id"`

	// Optional. Parent frame identifier.
	ParentID string `json:"parentId,omitempty"`

	// Identifier of the loader associated with this frame.
	LoaderID LoaderID `json:"loaderId"`

	// Optional. Frame's name as specified in the tag.
	Name string `json:"name,omitempty"`

	// Frame document's URL.
	URL string `json:"url"`

	// Frame document's security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Frame document's mimeType as determined by the browser.
	MimeType string `json:"mimeType"`

	// Optional. If the frame failed to load, this contains the URL that could
	// not be loaded. EXPERIMENTAL.
	UnreachableURL string `json:"unreachableUrl,omitempty"`
}

/*
FrameID is a unique frame identifier

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameId
*/
type FrameID string

/*
FrameResource provides information about the Resource on the page. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResource
*/
type FrameResource struct {
	// Resource URL.
	URL string `json:"url"`

	// Type of this resource.
	Type ResourceType `json:"type"`

	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`

	// Optional. last-modified timestamp as reported by server.
	LastModified TimeSinceEpoch `json:"lastModified,omitempty"`

	// Optional. Resource content size.
	ContentSize int `json:"contentSize,omitempty"`

	// Optional. True if the resource failed to load.
	Failed bool `json:"failed,omitempty"`

	// Optional. True if the resource was canceled during loading.
	Canceled bool `json:"canceled,omitempty"`
}

/*
FrameResourceTree provides information about the Frame hierarchy along with their cached resources.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResourceTree
*/
type FrameResourceTree struct {
	// Frame information for this tree item.
	Frame *Frame `json:"frame"`

	// Optional. Child frames.
	ChildFrames []*FrameResourceTree `json:"childFrames,omitempty"`

	// Information about frame resources.
	Resources []*FrameResource `json:"resources"`
}

/*
FrameTree provides information about the Frame hierarchy.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameTree
*/
type FrameTree struct {
	// Frame information for this tree item.
	Frame *Frame `json:"frame"`

	// Optional. Child frames.
	ChildFrames []*FrameTree `json:"childFrames,omitempty"`
}

/*
LayoutViewport defines layout viewport position and dimensions.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-LayoutViewport
*/
type LayoutViewport struct {
	// Horizontal offset relative to the document (CSS pixels).
	PageX int `json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY int `json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int `json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"`
}

/*
NavigationEntry defines a navigation history entry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-NavigationEntry
*/
type NavigationEntry struct {
	// Unique id of the navigation history entry.
	ID int `json:"id"`

	// URL of the navigation history entry.
	URL string `json:"url"`

	// URL that the user typed in the url bar.
	UserTypedURL string `json:"userTypedURL"`

	// Title of the navigation history entry.
	Title string `json:"title"`

	// Transition type.
	TransitionType TransitionType `json:"transitionType"`
}

/*
ResourceType is the resource type as it was perceived by the rendering engine.

ALLOWED VALUES
	- Document
	- Stylesheet
	- Image
	- Media
	- Font
	- Script
	- TextTrack
	- XHR
	- Fetch
	- EventSource
	- WebSocket
	- Manifest
	- Other

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ResourceType
*/
type ResourceType string

/*
ScreencastFrameMetadata provides screencast frame metadata. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScreencastFrameMetadata
*/
type ScreencastFrameMetadata struct {
	// Top offset in DIP.
	OffsetTop int `json:"offsetTop"`

	// Page scale factor.
	PageScaleFactor int `json:"pageScaleFactor"`

	// Device screen width in DIP.
	DeviceWidth int `json:"deviceWidth"`

	// Device screen height in DIP.
	DeviceHeight int `json:"deviceHeight"`

	// Position of horizontal scroll in CSS pixels.
	ScrollOffsetX int `json:"scrollOffsetX"`

	// Position of vertical scroll in CSS pixels.
	ScrollOffsetY int `json:"scrollOffsetY"`

	// Optional. Frame swap timestamp.
	Timestamp TimeSinceEpoch `json:"timestamp,omitempty"`
}

/*
ScriptIdentifier is the unique script identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScriptIdentifier
*/
type ScriptIdentifier string

/*
TransitionType is the transition type.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-TransitionType
*/
type TransitionType string

/*
Viewport defines the viewport for capturing screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Viewport
*/
type Viewport struct {
	// Required. X offset in CSS pixels.
	X int `json:"x"`

	// Required. Y offset in CSS pixels.
	Y int `json:"y"`

	// Required. Rectangle width in CSS pixels
	Width int `json:"width"`

	// Required. Rectangle height in CSS pixels
	Height int `json:"height"`

	// Required. Page scale factor.
	Scale int `json:"scale"`
}

/*
VisualViewport defines visual viewport position, dimensions, and scale.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-VisualViewport
*/
type VisualViewport struct {
	// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetX int `json:"offsetX"`

	// Vertical offset relative to the layout viewport (CSS pixels).
	OffsetY int `json:"offsetY"`

	// Horizontal offset relative to the document (CSS pixels).
	PageX int `json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY int `json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int `json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"`

	// Scale relative to the ideal viewport (size at width=device-width).
	Scale int `json:"scale"`
}
