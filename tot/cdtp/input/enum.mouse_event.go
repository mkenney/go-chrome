package input

import (
	"encoding/json"
	"fmt"
)

type mouseEventEnum struct {
	MousePressed  MouseEventEnum
	MouseReleased MouseEventEnum
	MouseMoved    MouseEventEnum
	MouseWheel    MouseEventEnum
}

/*
MouseEvent provides named acces to the MouseEventEnum values.
*/
var MouseEvent = mouseEventEnum{
	MousePressed:  mouseEventMousePressed,
	MouseReleased: mouseEventMouseReleased,
	MouseMoved:    mouseEventMouseMoved,
	MouseWheel:    mouseEventMouseWheel,
}

/*
MouseEventEnum represents the type of the mouse event. Allowed values:
	- MouseEvent.MousePressed  "mousePressed"
	- MouseEvent.MouseReleased "mouseReleased"
	- MouseEvent.MouseMoved    "mouseMoved"
	- MouseEvent.MouseWheel    "mouseWheel"

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
*/
type MouseEventEnum int

/*
String implements Stringer
*/
func (enum MouseEventEnum) String() string {
	return _mouseEventEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum MouseEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *MouseEventEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _mouseEventEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// mouseEventMousePressed represents the "mousePressed" value.
	mouseEventMousePressed MouseEventEnum = iota + 1
	// mouseEventMouseReleased represents the "mouseReleased" value.
	mouseEventMouseReleased
	// mouseEventMouseMoved represents the "mouseMoved" value.
	mouseEventMouseMoved
	// mouseEventMouseWheel represents the "mouseWheel" value.
	mouseEventMouseWheel
)

var _mouseEventEnums = map[MouseEventEnum]string{
	mouseEventMousePressed:  "mousePressed",
	mouseEventMouseReleased: "mouseReleased",
	mouseEventMouseMoved:    "mouseMoved",
	mouseEventMouseWheel:    "mouseWheel",
}
