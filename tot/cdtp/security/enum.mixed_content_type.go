package security

import (
	"encoding/json"
	"fmt"
)

type mixedContentTypeEnum struct {
	Blockable           MixedContentTypeEnum
	OptionallyBlockable MixedContentTypeEnum
	None                MixedContentTypeEnum
}

var MixedContentType = mixedContentTypeEnum{
	Blockable:           mixedContentTypeBlockable,
	OptionallyBlockable: mixedContentTypeOptionallyBlockable,
	None:                mixedContentTypeNone,
}

/*
MixedContentType is a description of mixed content (HTTP resources on HTTPS
pages), as defined by https://www.w3.org/TR/mixed-content/#categories
Allowed values:
	- MixedContentType.Blockable           "blockable"
	- MixedContentType.OptionallyBlockable "optionally-blockable"
	- MixedContentType.None                "none"

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-MixedContentType
*/
type MixedContentTypeEnum int

/*
String implements Stringer
*/
func (enum MixedContentTypeEnum) String() string {
	return _mixedContentTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum MixedContentTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *MixedContentTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _mixedContentTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// mixedContentTypeBlockable represents the "blockable" value.
	mixedContentTypeBlockable MixedContentTypeEnum = iota + 1
	// mixedContentTypeOptionallyBlockable represents the "optionally-blockable" value.
	mixedContentTypeOptionallyBlockable
	// mixedContentTypeNone represents the "none" value.
	mixedContentTypeNone
)

var _mixedContentTypeEnums = map[MixedContentTypeEnum]string{
	mixedContentTypeBlockable:           "blockable",
	mixedContentTypeOptionallyBlockable: "optionally-blockable",
	mixedContentTypeNone:                "none",
}
