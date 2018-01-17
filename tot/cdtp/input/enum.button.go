package input

import (
	"encoding/json"
	"fmt"
)

type buttonEventEnum struct {
	None   ButtonEventEnum
	Left   ButtonEventEnum
	Middle ButtonEventEnum
	Right  ButtonEventEnum
}

var ButtonEvent = buttonEventEnum{
	None:   buttonEventNone,
	Left:   buttonEventLeft,
	Middle: buttonEventMiddle,
	Right:  buttonEventRight,
}

/*
Mouse button (default: "none"). Allowed values:
	- none
	- left
	- middle
	- right

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
*/
type ButtonEventEnum int

/*
String implements Stringer
*/
func (enum ButtonEventEnum) String() string {
	return _buttonEventEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ButtonEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ButtonEventEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _buttonEventEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// buttonEventNone represents the "none" value.
	buttonEventNone ButtonEventEnum = iota + 1
	// buttonEventLeft represents the "left" value.
	buttonEventLeft
	// buttonEventMiddle represents the "middle" value.
	buttonEventMiddle
	// buttonEventRight represents the "right" value.
	buttonEventRight
)

var _buttonEventEnums = map[ButtonEventEnum]string{
	ButtonEventEnum(0): "",
	buttonEventNone:    "none",
	buttonEventLeft:    "left",
	buttonEventMiddle:  "middle",
	buttonEventRight:   "right",
}
