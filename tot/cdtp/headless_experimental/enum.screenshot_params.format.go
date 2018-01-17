package headlessExperimental

import (
	"encoding/json"
	"fmt"
)

type formatEnum struct {
	Jpeg FormatEnum
	Png  FormatEnum
}

var Format = formatEnum{
	Jpeg: formatJpeg,
	Png:  formatPng,
}

/*
Optional. Image compression format (defaults to png). Allowed values:
	- jpeg
	- png

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#type-ScreenshotParams
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

	return fmt.Errorf("%s is not a valid format value", bytes)
}

const (
	// formatJpeg represents the "jpeg" value.
	formatJpeg FormatEnum = iota + 1
	// formatPng represents the "png" value.
	formatPng
)

var _formatEnums = map[FormatEnum]string{
	FormatEnum(0): "",
	formatJpeg:    "jpeg",
	formatPng:     "png",
}
