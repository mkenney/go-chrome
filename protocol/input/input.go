/*
Package input provides type definitions for use with the Chrome Input protocol

https://chromedevtools.github.io/devtools-protocol/tot/Input/
*/
package input

/*
TouchPoint is a touch point.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-TouchPoint
*/
type TouchPoint struct {
	// X coordinate of the event relative to the main frame's viewport in CSS
	// pixels.
	X int `json:"x"`

	// Y coordinate of the event relative to the main frame's viewport in CSS
	// pixels. 0 refers to the top of the viewport and Y increases as it
	// proceeds towards the bottom of the viewport.
	Y int `json:"y"`

	// Optional. X radius of the touch area (default: 1.0).
	RadiusX int `json:"radiusX,omitempty"`

	// Optional. Y radius of the touch area (default: 1.0).
	RadiusY int `json:"radiusY,omitempty"`

	// Optional. Rotation angle (default: 0.0).
	RotationAngle int `json:"rotationAngle,omitempty"`

	// Optional. Force (default: 1.0).
	Force int `json:"force,omitempty"`

	// Optional. Identifier used to track touch sources between events, must be
	// unique within an event.
	ID int `json:"id,omitempty"`
}

/*
GestureSourceType is a gesture source type. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-GestureSourceType
*/
type GestureSourceType string

/*
TimeSinceEpoch is UTC time in seconds, counted from January 1, 1970.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-TimeSinceEpoch
*/
type TimeSinceEpoch int
