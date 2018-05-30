package css

import (
	"encoding/json"
	"fmt"
)

type sourceEnum struct {
	MediaRule   SourceEnum
	ImportRule  SourceEnum
	LinkedSheet SourceEnum
	InlineSheet SourceEnum
}

/*
Source provides named acces to the SourceEnum values.
*/
var Source = sourceEnum{
	MediaRule:   SourceMediaRule,
	ImportRule:  SourceImportRule,
	LinkedSheet: SourceLinkedSheet,
	InlineSheet: SourceInlineSheet,
}

/*
SourceEnum represents the source of the media query. Allowed values:
	- Source.MediaRule   "mediaRule" if specified by a @media rule
	- Source.ImportRule  "importRule" if specified by an @import rule
	- Source.LinkedSheet "linkedSheet" if specified by a "media" attribute in a
	  linked stylesheet's LINK tag
	- Source.InlineSheet "inlineSheet" if specified by a "media" attribute in an
	  inline stylesheet's STYLE tag.

Allowed values: mediaRule, importRule, linkedSheet, inlineSheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSMedia
*/
type SourceEnum int

/*
String implements Stringer
*/
func (enum SourceEnum) String() string {
	return _sourceEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum SourceEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *SourceEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _sourceEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid source value", bytes)
}

const (
	// SourceMediaRule represents the "mediaRule" value.
	SourceMediaRule SourceEnum = iota + 1
	// SourceImportRule represents the "importRule" value.
	SourceImportRule
	// SourceLinkedSheet represents the "linkedSheet" value.
	SourceLinkedSheet
	// SourceInlineSheet represents the "inlineSheet" value.
	SourceInlineSheet
)

var _sourceEnums = map[SourceEnum]string{
	SourceMediaRule:   "mediaRule",
	SourceImportRule:  "importRule",
	SourceLinkedSheet: "linkedSheet",
	SourceInlineSheet: "inlineSheet",
}
