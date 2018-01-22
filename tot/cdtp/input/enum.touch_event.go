package input

import (
	"encoding/json"
	"fmt"
)

type touchEventEnum struct {
	TouchStart  TouchEventEnum
	TouchEnd    TouchEventEnum
	TouchMove   TouchEventEnum
	TouchCancel TouchEventEnum
}

var TouchEvent = touchEventEnum{
	TouchStart:  touchEventTouchStart,
	TouchEnd:    touchEventTouchEnd,
	TouchMove:   touchEventTouchMove,
	TouchCancel: touchEventTouchCancel,
}

/*
Type of the touch event. TouchEnd and TouchCancel must not contain any touch
points, while TouchStart and TouchMove must contains at least one. Allowed
values:
	- touchStart
	- touchEnd
	- touchMove
	- touchCancel

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
*/
type TouchEventEnum int

/*
String implements Stringer
*/
func (enum TouchEventEnum) String() string {
	return _touchEventEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum TouchEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *TouchEventEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _touchEventEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// touchEventTouchStart represents the "touchStart" value.
	touchEventTouchStart TouchEventEnum = iota + 1
	// touchEventTouchEnd represents the "touchEnd" value.
	touchEventTouchEnd
	// touchEventTouchMove represents the "touchMove" value.
	touchEventTouchMove
	// touchEventTouchCancel represents the "touchCancel" value.
	touchEventTouchCancel
)

var _touchEventEnums = map[TouchEventEnum]string{
	touchEventTouchStart:  "touchStart",
	touchEventTouchEnd:    "touchEnd",
	touchEventTouchMove:   "touchMove",
	touchEventTouchCancel: "touchCancel",
}
