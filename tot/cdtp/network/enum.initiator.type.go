package network

import (
	"encoding/json"
	"fmt"
)

type initiatorTypeEnum struct {
	Parser  InitiatorTypeEnum
	Script  InitiatorTypeEnum
	Preload InitiatorTypeEnum
	Other   InitiatorTypeEnum
}

var InitiatorType = initiatorTypeEnum{
	Parser:  initiatorTypeParser,
	Script:  initiatorTypeScript,
	Preload: initiatorTypePreload,
	Other:   initiatorTypeOther,
}

/*
InitiatorTypeEnum is the type of this initiator. Allowed values:
	- InitiatorType.Parser
	- InitiatorType.Script
	- InitiatorType.Preload
	- InitiatorType.Other

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-InitiatorType
*/
type InitiatorTypeEnum int

/*
String implements Stringer
*/
func (enum InitiatorTypeEnum) String() string {
	return _initiatorTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum InitiatorTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *InitiatorTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _initiatorTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// initiatorTypeParser represents the "parser" value.
	initiatorTypeParser InitiatorTypeEnum = iota + 1
	// initiatorTypeScript represents the "script" value.
	initiatorTypeScript
	// initiatorTypePreload represents the "preload" value.
	initiatorTypePreload
	// initiatorTypeOther represents the "other" value.
	initiatorTypeOther
)

var _initiatorTypeEnums = map[InitiatorTypeEnum]string{
	initiatorTypeParser:  "parser",
	initiatorTypeScript:  "script",
	initiatorTypePreload: "preload",
	initiatorTypeOther:   "other",
}
