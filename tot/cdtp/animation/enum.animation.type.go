package animation

import (
	"encoding/json"
	"fmt"
)

type animationTypeEnum struct {
	CSSTransition AnimationTypeEnum
	CSSAnimation  AnimationTypeEnum
	WebAnimation  AnimationTypeEnum
}

var AnimationType = animationTypeEnum{
	CSSTransition: animationTypeCSSTransition,
	CSSAnimation:  animationTypeCSSAnimation,
	WebAnimation:  animationTypeWebAnimation,
}

/*
Animation type of Animation. Allowed values:
	- CSSTransition
	- CSSAnimation
	- WebAnimation

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-Animation
*/
type AnimationTypeEnum int

/*
String implements Stringer
*/
func (enum AnimationTypeEnum) String() string {
	return _animationTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum AnimationTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *AnimationTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _animationTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid animationType value", bytes)
}

const (
	// animationTypeCSSTransition represents the "CSSTransition" value.
	animationTypeCSSTransition AnimationTypeEnum = iota + 1
	// animationTypeCSSAnimation represents the "CSSAnimation" value.
	animationTypeCSSAnimation
	// animationTypeWebAnimation represents the "WebAnimation" value.
	animationTypeWebAnimation
)

var _animationTypeEnums = map[AnimationTypeEnum]string{
	animationTypeCSSTransition: "CSSTransition",
	animationTypeCSSAnimation:  "CSSAnimation",
	animationTypeWebAnimation:  "WebAnimation",
}
