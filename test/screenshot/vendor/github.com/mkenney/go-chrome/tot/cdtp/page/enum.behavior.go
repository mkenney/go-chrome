package page

import (
	"encoding/json"
	"fmt"
)

type behaviorEnum struct {
	Deny    BehaviorEnum
	Allow   BehaviorEnum
	Default BehaviorEnum
}

/*
Behavior provides named acces to the BehaviorEnum values.
*/
var Behavior = behaviorEnum{
	Deny:    behaviorDeny,
	Allow:   behaviorAllow,
	Default: behaviorDefault,
}

/*
BehaviorEnum represents whether to allow all or deny all download requests, or
use default Chrome behavior if available (otherwise deny). Allowed values:
	- Behavior.Deny    "deny"
	- Behavior.Allow   "allow"
	- Behavior.Default "default"

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
*/
type BehaviorEnum int

/*
String implements Stringer
*/
func (enum BehaviorEnum) String() string {
	return _behaviorEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum BehaviorEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *BehaviorEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _behaviorEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// behaviorDeny represents the "deny" value.
	behaviorDeny BehaviorEnum = iota + 1
	// behaviorAllow represents the "allow" value.
	behaviorAllow
	// behaviorDefault represents the "default" value.
	behaviorDefault
)

var _behaviorEnums = map[BehaviorEnum]string{
	behaviorDeny:    "deny",
	behaviorAllow:   "allow",
	behaviorDefault: "default",
}
