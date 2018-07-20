package css

import (
	"encoding/json"
	"fmt"
)

type styleSheetOriginEnum struct {
	Injected  StyleSheetOriginEnum
	UserAgent StyleSheetOriginEnum
	Inspector StyleSheetOriginEnum
	Log       StyleSheetOriginEnum
}

/*
StyleSheetOrigin provides named acces to the StyleSheetOriginEnum values.
*/
var StyleSheetOrigin = styleSheetOriginEnum{
	Injected:  StyleSheetOriginInjected,
	UserAgent: StyleSheetOriginUserAgent,
	Inspector: StyleSheetOriginInspector,
	Log:       StyleSheetOriginLog,
}

/*
StyleSheetOriginEnum specifies the stylesheet origin. Allowed values:
	- StyleSheetOrigin.Injected  "injected" for stylesheets injected via
	  extension
	- StyleSheetOrigin.UserAgent "user-agent" for user-agent stylesheets
	- StyleSheetOrigin.Inspector "inspector" for stylesheets created by the
	  inspector (i.e. those holding the "via inspector" rules)
	- StyleSheetOrigin.Log       "regular" for regular stylesheets.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-StyleSheetOrigin
*/
type StyleSheetOriginEnum int

/*
String implements Stringer
*/
func (source StyleSheetOriginEnum) String() string {
	return _styleSheetOriginEnums[source]
}

/*
MarshalJSON implements json.Marshaler
*/
func (source StyleSheetOriginEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(source.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (source *StyleSheetOriginEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _styleSheetOriginEnums {
		if v == val {
			*source = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid source value", bytes)
}

const (
	// StyleSheetOriginInjected represents the "injected" value.
	StyleSheetOriginInjected StyleSheetOriginEnum = iota + 1
	// StyleSheetOriginUserAgent represents the "user-agent" value.
	StyleSheetOriginUserAgent
	// StyleSheetOriginInspector represents the "inspector" value.
	StyleSheetOriginInspector
	// StyleSheetOriginLog represents the "log" value.
	StyleSheetOriginLog
)

var _styleSheetOriginEnums = map[StyleSheetOriginEnum]string{
	StyleSheetOriginInjected:  "injected",
	StyleSheetOriginUserAgent: "user-agent",
	StyleSheetOriginInspector: "inspector",
	StyleSheetOriginLog:       "log",
}
