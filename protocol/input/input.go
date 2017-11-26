package Input

import (
	"fmt"
)

/*
TouchPoint is a touch point.
*/
type TouchPoint struct {
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X int `json:"x"`

	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to
	// the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y int `json:"y"`

	// Optional. X radius of the touch area (default: 1.0).
	RadiusX int `json:"radiusX,omitempty"`

	// Optional. Y radius of the touch area (default: 1.0).
	RadiusY int `json:"radiusY,omitempty"`

	// Optional. Rotation angle (default: 0.0).
	RotationAngle int `json:"rotationAngle,omitempty"`

	// Optional. Force (default: 1.0).
	Force int `json:"force,omitempty"`

	// Optional. Identifier used to track touch sources between events, must be unique within an
	// event.
	ID int `json:"id,omitempty"`
}

/*
GestureSourceType is a gesture source type. EXPERIMENTAL
*/
type GestureSourceType int

const (
	_default GestureSourceType = iota
	_touch
	_mouse
)

func (a GestureSourceType) String() string {
	if a == 0 {
		return "default"
	}
	if a == 1 {
		return "touch"
	}
	if a == 2 {
		return "mouse"
	}
	panic(fmt.Errorf("Invalid GestureSourceType %d", a))
}

/*
TimeSinceEpoch is UTC time in seconds, counted from January 1, 1970.
*/
type TimeSinceEpoch int
