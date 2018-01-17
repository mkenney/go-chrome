package indexedDB

import (
	"encoding/json"
	"fmt"
)

type keyTypeEnum struct {
	Number KeyTypeEnum
	String KeyTypeEnum
	Date   KeyTypeEnum
	Array  KeyTypeEnum
}

var KeyType = keyTypeEnum{
	Number: keyTypeNumber,
	String: keyTypeString,
	Date:   keyTypeDate,
	Array:  keyTypeArray,
}

/*
Key type. Allowed values:
	- number
	- string
	- date
	- array

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#type-ScreenshotParams
*/
type KeyTypeEnum int

/*
String implements Stringer
*/
func (enum KeyTypeEnum) String() string {
	return _keyTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum KeyTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *KeyTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _keyTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// keyTypeNumber represents the "number" value.
	keyTypeNumber KeyTypeEnum = iota + 1
	// keyTypeString represents the "string" value.
	keyTypeString
	// keyTypeDate represents the "date" value.
	keyTypeDate
	// keyTypeArray represents the "array" value.
	keyTypeArray
)

var _keyTypeEnums = map[KeyTypeEnum]string{
	keyTypeNumber: "number",
	keyTypeString: "string",
	keyTypeDate:   "date",
	keyTypeArray:  "array",
}
