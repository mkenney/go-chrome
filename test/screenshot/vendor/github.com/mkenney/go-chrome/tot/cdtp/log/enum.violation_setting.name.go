package log

import (
	"encoding/json"
	"fmt"
)

type nameEnum struct {
	LongTask          NameEnum
	LongLayout        NameEnum
	BlockedEvent      NameEnum
	BlockedParser     NameEnum
	DiscouragedAPIUse NameEnum
	Handler           NameEnum
	RecurringHandler  NameEnum
}

/*
Name provides named acces to the NameEnum values.
*/
var Name = nameEnum{
	LongTask:          nameLongTask,
	LongLayout:        nameLongLayout,
	BlockedEvent:      nameBlockedEvent,
	BlockedParser:     nameBlockedParser,
	DiscouragedAPIUse: nameDiscouragedAPIUse,
	Handler:           nameHandler,
	RecurringHandler:  nameRecurringHandler,
}

/*
NameEnum represents the violation type. Allowed values:
	- Name.LongTask          "longTask"
	- Name.LongLayout        "longLayout"
	- Name.BlockedEvent      "blockedEvent"
	- Name.BlockedParser     "blockedParser"
	- Name.DiscouragedAPIUse "discouragedAPIUse"
	- Name.Handler           "handler"
	- Name.RecurringHandler  "recurringHandler"

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-ViolationSetting
*/
type NameEnum int

/*
String implements Stringer
*/
func (enum NameEnum) String() string {
	return _nameEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum NameEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *NameEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _nameEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// nameLongTask represents the "longTask" value.
	nameLongTask NameEnum = iota + 1
	// nameLongLayout represents the "longLayout" value.
	nameLongLayout
	// nameBlockedEvent represents the "blockedEvent" value.
	nameBlockedEvent
	// nameBlockedParser represents the "blockedParser" value.
	nameBlockedParser
	// nameDiscouragedAPIUse represents the "discouragedAPIUse" value.
	nameDiscouragedAPIUse
	// nameHandler represents the "handler" value.
	nameHandler
	// nameRecurringHandler represents the "recurringHandler" value.
	nameRecurringHandler
)

var _nameEnums = map[NameEnum]string{
	nameLongTask:          "longTask",
	nameLongLayout:        "longLayout",
	nameBlockedEvent:      "blockedEvent",
	nameBlockedParser:     "blockedParser",
	nameDiscouragedAPIUse: "discouragedAPIUse",
	nameHandler:           "handler",
	nameRecurringHandler:  "recurringHandler",
}
