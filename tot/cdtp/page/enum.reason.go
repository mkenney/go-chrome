package page

import (
	"encoding/json"
	"fmt"
)

type reasonEnum struct {
	FormSubmissionGet     ReasonEnum
	FormSubmissionPost    ReasonEnum
	HttpHeaderRefresh     ReasonEnum
	ScriptInitiated       ReasonEnum
	MetaTagRefresh        ReasonEnum
	PageBlockInterstitial ReasonEnum
	Reload                ReasonEnum
}

var Reason = reasonEnum{
	FormSubmissionGet:     reasonFormSubmissionGet,
	FormSubmissionPost:    reasonFormSubmissionPost,
	HttpHeaderRefresh:     reasonHttpHeaderRefresh,
	ScriptInitiated:       reasonScriptInitiated,
	MetaTagRefresh:        reasonMetaTagRefresh,
	PageBlockInterstitial: reasonPageBlockInterstitial,
	Reload:                reasonReload,
}

/*
The reason for the navigation. Allowed values:
	- Reason.FormSubmissionGet
	- Reason.FormSubmissionPost
	- Reason.HttpHeaderRefresh
	- Reason.ScriptInitiated
	- Reason.MetaTagRefresh
	- Reason.PageBlockInterstitial
	- Reason.Reload

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameScheduledNavigation
*/
type ReasonEnum int

/*
String implements Stringer
*/
func (enum ReasonEnum) String() string {
	return _reasonEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ReasonEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ReasonEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _reasonEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// reasonFormSubmissionGet represents the "formSubmissionGet" value.
	reasonFormSubmissionGet ReasonEnum = iota + 1
	// reasonFormSubmissionPost represents the "formSubmissionPost" value.
	reasonFormSubmissionPost
	// reasonHttpHeaderRefresh represents the "httpHeaderRefresh" value.
	reasonHttpHeaderRefresh
	// reasonScriptInitiated represents the "scriptInitiated" value.
	reasonScriptInitiated
	// reasonMetaTagRefresh represents the "metaTagRefresh" value.
	reasonMetaTagRefresh
	// reasonPageBlockInterstitial represents the "pageBlockInterstitial" value.
	reasonPageBlockInterstitial
	// reasonReload represents the "reload" value.
	reasonReload
)

var _reasonEnums = map[ReasonEnum]string{
	reasonFormSubmissionGet:     "formSubmissionGet",
	reasonFormSubmissionPost:    "formSubmissionPost",
	reasonHttpHeaderRefresh:     "httpHeaderRefresh",
	reasonScriptInitiated:       "scriptInitiated",
	reasonMetaTagRefresh:        "metaTagRefresh",
	reasonPageBlockInterstitial: "pageBlockInterstitial",
	reasonReload:                "reload",
}
