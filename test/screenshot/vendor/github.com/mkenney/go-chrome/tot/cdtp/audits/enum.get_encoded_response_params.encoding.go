package audits

import (
	"encoding/json"
	"fmt"
)

type encodingEnum struct {
	Webp EncodingEnum
	Jpeg EncodingEnum
	Png  EncodingEnum
}

/*
Encoding provides named acces to the EncodingEnum values.
*/
var Encoding = encodingEnum{
	Webp: encodingWebp,
	Jpeg: encodingJpeg,
	Png:  encodingPng,
}

/*
EncodingEnum represents the encoding to use. Allowed values:
	- Encoding.Webp "webp"
	- Encoding.Jpeg "jpeg"
	- Encoding.Png  "png"

https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
*/
type EncodingEnum int

/*
String implements Stringer
*/
func (enum EncodingEnum) String() string {
	return _encodingEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum EncodingEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *EncodingEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _encodingEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid encoding value", bytes)
}

const (
	// encodingWebp represents the "webp" value.
	encodingWebp EncodingEnum = iota + 1
	// encodingJpeg represents the "jpeg" value.
	encodingJpeg
	// encodingPng represents the "png" value.
	encodingPng
)

var _encodingEnums = map[EncodingEnum]string{
	encodingWebp: "webp",
	encodingJpeg: "jpeg",
	encodingPng:  "png",
}
