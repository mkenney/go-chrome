package emulation

import (
	"encoding/json"
	"fmt"
)

type orientationTypeEnum struct {
	PortraitPrimary    OrientationTypeEnum
	PortraitSecondary  OrientationTypeEnum
	LandscapePrimary   OrientationTypeEnum
	LandscapeSecondary OrientationTypeEnum
}

/*
OrientationType provides named acces to the OrientationTypeEnum values.
*/
var OrientationType = orientationTypeEnum{
	PortraitPrimary:    orientationTypePortraitPrimary,
	PortraitSecondary:  orientationTypePortraitSecondary,
	LandscapePrimary:   orientationTypeLandscapePrimary,
	LandscapeSecondary: orientationTypeLandscapeSecondary,
}

/*
OrientationTypeEnum represents the orientation type. Allowed values:
	- OrientationType.PortraitPrimary    "portraitPrimary"
	- OrientationType.PortraitSecondary  "portraitSecondary"
	- OrientationType.LandscapePrimary   "landscapePrimary"
	- OrientationType.LandscapeSecondary "landscapeSecondary"

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-ScreenOrientation
*/
type OrientationTypeEnum int

/*
String implements Stringer
*/
func (enum OrientationTypeEnum) String() string {
	return _orientationTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum OrientationTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *OrientationTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _orientationTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid orientationType value", bytes)
}

const (
	// orientationTypePortraitPrimary represents the "portraitPrimary" value.
	orientationTypePortraitPrimary OrientationTypeEnum = iota + 1
	// orientationTypePortraitSecondary represents the "portraitSecondary" value.
	orientationTypePortraitSecondary
	// orientationTypeLandscapePrimary represents the "landscapePrimary" value.
	orientationTypeLandscapePrimary
	// orientationTypeLandscapeSecondary represents the "landscapeSecondary" value.
	orientationTypeLandscapeSecondary
)

var _orientationTypeEnums = map[OrientationTypeEnum]string{
	orientationTypePortraitPrimary:    "portraitPrimary",
	orientationTypePortraitSecondary:  "portraitSecondary",
	orientationTypeLandscapePrimary:   "landscapePrimary",
	orientationTypeLandscapeSecondary: "landscapeSecondary",
}
