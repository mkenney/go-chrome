package input

import (
	"encoding/json"
	"fmt"
)

type keyEventEnum struct {
	KeyDown    KeyEventEnum
	KeyUp      KeyEventEnum
	RawKeyDown KeyEventEnum
	Char       KeyEventEnum
}

var KeyEvent = keyEventEnum{
	KeyDown:    keyEventKeyDown,
	KeyUp:      keyEventKeyUp,
	RawKeyDown: keyEventRawKeyDown,
	Char:       keyEventChar,
}

/*
Type of the key event. Allowed values:
	- keyDown
	- keyUp
	- rawKeyDown
	- char

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchKeyEvent
*/
type KeyEventEnum int

/*
String implements Stringer
*/
func (enum KeyEventEnum) String() string {
	return _keyEventEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum KeyEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *KeyEventEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _keyEventEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// keyEventKeyDown represents the "keyDown" value.
	keyEventKeyDown KeyEventEnum = iota + 1
	// keyEventKeyUp represents the "keyUp" value.
	keyEventKeyUp
	// keyEventRawKeyDown represents the "rawKeyDown" value.
	keyEventRawKeyDown
	// keyEventChar represents the "char" value.
	keyEventChar
)

var _keyEventEnums = map[KeyEventEnum]string{
	keyEventKeyDown:    "keyDown",
	keyEventKeyUp:      "keyUp",
	keyEventRawKeyDown: "rawKeyDown",
	keyEventChar:       "char",
}
