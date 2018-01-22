package css

import (
	"encoding/json"
	"fmt"
)

type forcedPseudoClassesEnum struct {
	Active  ForcedPseudoClassesEnum
	Focus   ForcedPseudoClassesEnum
	Hover   ForcedPseudoClassesEnum
	Visited ForcedPseudoClassesEnum
}

/*
ForcedPseudoClasses provides named acces to the ForcedPseudoClassesEnum values.
*/
var ForcedPseudoClasses = forcedPseudoClassesEnum{
	Active:  forcedPseudoClassesActive,
	Focus:   forcedPseudoClassesFocus,
	Hover:   forcedPseudoClassesHover,
	Visited: forcedPseudoClassesVisited,
}

/*
ForcedPseudoClassesEnum represents element pseudo classes to force when
computing the element's style. Allowed values:
	- ForcedPseudoClasses.Active  "active"
	- ForcedPseudoClasses.Focus   "focus"
	- ForcedPseudoClasses.Hover   "hover"
	- ForcedPseudoClasses.Visited "visited"

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
*/
type ForcedPseudoClassesEnum int

/*
String implements Stringer
*/
func (enum ForcedPseudoClassesEnum) String() string {
	return _forcedPseudoClassesEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ForcedPseudoClassesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ForcedPseudoClassesEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _forcedPseudoClassesEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid ForcedPseudoClass value", bytes)
}

const (
	// forcedPseudoClassesActive represents the "active" value.
	forcedPseudoClassesActive ForcedPseudoClassesEnum = iota + 1
	// forcedPseudoClassesFocus represents the "focus" value.
	forcedPseudoClassesFocus
	// forcedPseudoClassesHover represents the "hover" value.
	forcedPseudoClassesHover
	// forcedPseudoClassesVisited represents the "visited" value.
	forcedPseudoClassesVisited
)

var _forcedPseudoClassesEnums = map[ForcedPseudoClassesEnum]string{
	forcedPseudoClassesActive:  "active",
	forcedPseudoClassesFocus:   "focus",
	forcedPseudoClassesHover:   "hover",
	forcedPseudoClassesVisited: "visited",
}
