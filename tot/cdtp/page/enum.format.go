package page

import (
	"encoding/json"
	"fmt"
)

type formatEnum struct {
	Png  FormatEnum
	Jpeg FormatEnum
}

var Format = formatEnum{
	Png:  formatPng,
	Jpeg: formatJpeg,
}

/*
Format defines the Javascript dialog type. Allowed values:
	- Format.Png
	- Format.Jpeg

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
*/
type FormatEnum int

/*
String implements Stringer
*/
func (enum FormatEnum) String() string {
	return _formatEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum FormatEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *FormatEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _formatEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// formatPng represents the "png" value.
	formatPng FormatEnum = iota + 1
	// formatJpeg represents the "jpeg" value.
	formatJpeg
)

var _formatEnums = map[FormatEnum]string{
	FormatEnum(0): "",
	formatPng:     "png",
	formatJpeg:    "jpeg",
}
