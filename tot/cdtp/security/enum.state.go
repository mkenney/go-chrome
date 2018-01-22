package security

import (
	"encoding/json"
	"fmt"
)

type stateEnum struct {
	Unknown  StateEnum
	Neutral  StateEnum
	Insecure StateEnum
	Secure   StateEnum
	Info     StateEnum
}

var State = stateEnum{
	Unknown:  stateUnknown,
	Neutral:  stateNeutral,
	Insecure: stateInsecure,
	Secure:   stateSecure,
	Info:     stateInfo,
}

/*
State is the security level of a page or resource. Allowed values:
	- State.Unknown
	- State.Neutral
	- State.Insecure
	- State.Secure
	- State.Info

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SecurityState
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

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// stateUnknown represents the "unknown" value.
	stateUnknown StateEnum = iota + 1
	// stateNeutral represents the "neutral" value.
	stateNeutral
	// stateInsecure represents the "insecure" value.
	stateInsecure
	// stateSecure represents the "secure" value.
	stateSecure
	// stateInfo represents the "info" value.
	stateInfo
)

var _stateEnums = map[StateEnum]string{
	stateUnknown:  "unknown",
	stateNeutral:  "neutral",
	stateInsecure: "insecure",
	stateSecure:   "secure",
	stateInfo:     "info",
}
