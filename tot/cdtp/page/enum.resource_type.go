package page

import (
	"encoding/json"
	"fmt"
)

type resourceTypeEnum struct {
	Document    ResourceTypeEnum
	Stylesheet  ResourceTypeEnum
	Image       ResourceTypeEnum
	Media       ResourceTypeEnum
	Font        ResourceTypeEnum
	Script      ResourceTypeEnum
	TextTrack   ResourceTypeEnum
	XHR         ResourceTypeEnum
	Fetch       ResourceTypeEnum
	EventSource ResourceTypeEnum
	WebSocket   ResourceTypeEnum
	Manifest    ResourceTypeEnum
	Other       ResourceTypeEnum
}

var ResourceType = resourceTypeEnum{
	Document:    resourceTypeDocument,
	Stylesheet:  resourceTypeStylesheet,
	Image:       resourceTypeImage,
	Media:       resourceTypeMedia,
	Font:        resourceTypeFont,
	Script:      resourceTypeScript,
	TextTrack:   resourceTypeTextTrack,
	XHR:         resourceTypeXHR,
	Fetch:       resourceTypeFetch,
	EventSource: resourceTypeEventSource,
	WebSocket:   resourceTypeWebSocket,
	Manifest:    resourceTypeManifest,
	Other:       resourceTypeOther,
}

/*
ResourceType is the resource type as it was perceived by the rendering engine.
Allowed Values:
	- ResourceType.Document
	- ResourceType.Stylesheet
	- ResourceType.Image
	- ResourceType.Media
	- ResourceType.Font
	- ResourceType.Script
	- ResourceType.TextTrack
	- ResourceType.XHR
	- ResourceType.Fetch
	- ResourceType.EventSource
	- ResourceType.WebSocket
	- ResourceType.Manifest
	- ResourceType.Other

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ResourceType
*/
type ResourceTypeEnum int

/*
String implements Stringer
*/
func (enum ResourceTypeEnum) String() string {
	return _resourceTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ResourceTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ResourceTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _resourceTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// resourceTypeDocument represents the "Document" value.
	resourceTypeDocument ResourceTypeEnum = iota + 1
	// resourceTypeStylesheet represents the "Stylesheet" value.
	resourceTypeStylesheet
	// resourceTypeImage represents the "Image" value.
	resourceTypeImage
	// resourceTypeMedia represents the "Media" value.
	resourceTypeMedia
	// resourceTypeFont represents the "Font" value.
	resourceTypeFont
	// resourceTypeScript represents the "Script" value.
	resourceTypeScript
	// resourceTypeTextTrack represents the "TextTrack" value.
	resourceTypeTextTrack
	// resourceTypeXHR represents the "XHR" value.
	resourceTypeXHR
	// resourceTypeFetch represents the "Fetch" value.
	resourceTypeFetch
	// resourceTypeEventSource represents the "EventSource" value.
	resourceTypeEventSource
	// resourceTypeWebSocket represents the "WebSocket" value.
	resourceTypeWebSocket
	// resourceTypeManifest represents the "Manifest" value.
	resourceTypeManifest
	// resourceTypeOther represents the "Other" value.
	resourceTypeOther
)

var _resourceTypeEnums = map[ResourceTypeEnum]string{
	resourceTypeDocument:    "Document",
	resourceTypeStylesheet:  "Stylesheet",
	resourceTypeImage:       "Image",
	resourceTypeMedia:       "Media",
	resourceTypeFont:        "Font",
	resourceTypeScript:      "Script",
	resourceTypeTextTrack:   "TextTrack",
	resourceTypeXHR:         "XHR",
	resourceTypeFetch:       "Fetch",
	resourceTypeEventSource: "EventSource",
	resourceTypeWebSocket:   "WebSocket",
	resourceTypeManifest:    "Manifest",
	resourceTypeOther:       "Other",
}
