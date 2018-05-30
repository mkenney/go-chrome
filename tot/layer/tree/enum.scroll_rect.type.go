package tree

import (
	"encoding/json"
	"fmt"
)

type rectTypeEnum struct {
	RepaintsOnScroll  RectTypeEnum
	TouchEventHandler RectTypeEnum
	WheelEventHandler RectTypeEnum
}

/*
RectType provides named acces to the RectTypeEnum values.
*/
var RectType = rectTypeEnum{
	RepaintsOnScroll:  rectTypeRepaintsOnScroll,
	TouchEventHandler: rectTypeTouchEventHandler,
	WheelEventHandler: rectTypeWheelEventHandler,
}

/*
RectTypeEnum represents the reason for a rectangle to force scrolling on the
main thread. Allowed values:
	- RectType.RepaintsOnScroll  "RepaintsOnScroll"
	- RectType.TouchEventHandler "TouchEventHandler"
	- RectType.WheelEventHandler "WheelEventHandler"

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-ScrollRect
*/
type RectTypeEnum int

/*
String implements Stringer
*/
func (enum RectTypeEnum) String() string {
	return _rectTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum RectTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *RectTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _rectTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// rectTypeRepaintsOnScroll represents the "RepaintsOnScroll" value.
	rectTypeRepaintsOnScroll RectTypeEnum = iota + 1
	// rectTypeTouchEventHandler represents the "TouchEventHandler" value.
	rectTypeTouchEventHandler
	// rectTypeWheelEventHandler represents the "WheelEventHandler" value.
	rectTypeWheelEventHandler
)

var _rectTypeEnums = map[RectTypeEnum]string{
	rectTypeRepaintsOnScroll:  "RepaintsOnScroll",
	rectTypeTouchEventHandler: "TouchEventHandler",
	rectTypeWheelEventHandler: "WheelEventHandler",
}
