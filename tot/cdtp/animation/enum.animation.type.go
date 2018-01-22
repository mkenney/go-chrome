package animation

import (
	"encoding/json"
	"fmt"
)

type typeEnum struct {
	CSSTransition TypeEnum
	CSSAnimation  TypeEnum
	WebAnimation  TypeEnum
}

/*
Type provides named acces to the TypeEnum values.
*/
var Type = typeEnum{
	CSSTransition: typeCSSTransition,
	CSSAnimation:  typeCSSAnimation,
	WebAnimation:  typeWebAnimation,
}

/*
TypeEnum represents the type of Animation. Allowed values:
	- Type.CSSTransition "CSSTransition"
	- Type.CSSAnimation  "CSSAnimation"
	- Type.WebAnimation  "WebAnimation"

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-Animation
*/
type TypeEnum int

/*
String implements Stringer
*/
func (enum TypeEnum) String() string {
	return _typeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum TypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *TypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _typeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid Type value", bytes)
}

const (
	// typeCSSTransition represents the "CSSTransition" value.
	typeCSSTransition TypeEnum = iota + 1
	// typeCSSAnimation represents the "CSSAnimation" value.
	typeCSSAnimation
	// typeWebAnimation represents the "WebAnimation" value.
	typeWebAnimation
)

var _typeEnums = map[TypeEnum]string{
	typeCSSTransition: "CSSTransition",
	typeCSSAnimation:  "CSSAnimation",
	typeWebAnimation:  "WebAnimation",
}
