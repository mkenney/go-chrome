package runtime

import (
	"encoding/json"
	"fmt"
)

type unserializableValueEnum struct {
	Infinity    UnserializableValueEnum
	NaN         UnserializableValueEnum
	NegInfinity UnserializableValueEnum
	NegZero     UnserializableValueEnum
}

/*
UnserializableValue provides named acces to the UnserializableValueEnum values.
*/
var UnserializableValue = unserializableValueEnum{
	Infinity:    unserializableValueInfinity,
	NaN:         unserializableValueNaN,
	NegInfinity: unserializableValueNegInfinity,
	NegZero:     unserializableValueNegZero,
}

/*
UnserializableValueEnum is a primitive value which cannot be JSON-stringified.
Allowed values:
	- UnserializableValue.Infinity    "Infinity"
	- UnserializableValue.NaN         "NaN"
	- UnserializableValue.NegInfinity "-Infinity"
	- UnserializableValue.NegZero     "-0"

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-UnserializableValue
*/
type UnserializableValueEnum int

/*
String implements Stringer
*/
func (enum UnserializableValueEnum) String() string {
	return _unserializableValueEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum UnserializableValueEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *UnserializableValueEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _unserializableValueEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// unserializableValueInfinity represents the "Infinity" value.
	unserializableValueInfinity UnserializableValueEnum = iota + 1
	// unserializableValueNaN represents the "NaN" value.
	unserializableValueNaN
	// unserializableValueNegInfinity represents the "-Infinity" value.
	unserializableValueNegInfinity
	// unserializableValueNegZero represents the "-0" value.
	unserializableValueNegZero
)

var _unserializableValueEnums = map[UnserializableValueEnum]string{
	unserializableValueInfinity:    "Infinity",
	unserializableValueNaN:         "NaN",
	unserializableValueNegInfinity: "-Infinity",
	unserializableValueNegZero:     "-0",
}
