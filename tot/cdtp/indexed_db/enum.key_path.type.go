package indexedDB

import (
	"encoding/json"
	"fmt"
)

type keyPathTypeEnum struct {
	Null   KeyPathTypeEnum
	String KeyPathTypeEnum
	Array  KeyPathTypeEnum
}

/*
KeyPathType provides named acces to the KeyPathTypeEnum values.
*/
var KeyPathType = keyPathTypeEnum{
	Null:   keyPathTypeNull,
	String: keyPathTypeString,
	Array:  keyPathTypeArray,
}

/*
KeyPathTypeEnum represents the key path type. Allowed values:
	- KeyPathType.Null   "null"
	- KeyPathType.String "string"
	- KeyPathType.Array  "array"

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#type-ScreenshotParams
*/
type KeyPathTypeEnum int

/*
String implements Stringer
*/
func (enum KeyPathTypeEnum) String() string {
	return _keyPathTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum KeyPathTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *KeyPathTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _keyPathTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// keyPathTypeNull represents the "null" value.
	keyPathTypeNull KeyPathTypeEnum = iota + 1
	// keyPathTypeString represents the "string" value.
	keyPathTypeString
	// keyPathTypeArray represents the "array" value.
	keyPathTypeArray
)

var _keyPathTypeEnums = map[KeyPathTypeEnum]string{
	keyPathTypeNull:   "null",
	keyPathTypeString: "string",
	keyPathTypeArray:  "array",
}
