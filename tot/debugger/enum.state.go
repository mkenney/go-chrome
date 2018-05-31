package debugger

import (
	"encoding/json"
	"fmt"
)

type stateEnum struct {
	None     StateEnum
	Uncaught StateEnum
	All      StateEnum
}

/*
State provides named acces to the StateEnum values.
*/
var State = stateEnum{
	None:     stateNone,
	Uncaught: stateUncaught,
	All:      stateAll,
}

/*
StateEnum represents the pause on exceptions mode. Allowed values:
	- State.None     "none"
	- State.Uncaught "uncaught"
	- State.All      "all"

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
*/
type StateEnum int

/*
String implements Stringer
*/
func (enum StateEnum) String() string {
	return _stateEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum StateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *StateEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _stateEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid state value", bytes)
}

const (
	// stateNone represents the "none" value.
	stateNone StateEnum = iota + 1
	// stateUncaught represents the "uncaught" value.
	stateUncaught
	// stateAll represents the "all" value.
	stateAll
)

var _stateEnums = map[StateEnum]string{
	stateNone:     "none",
	stateUncaught: "uncaught",
	stateAll:      "all",
}
