package page

import (
	"encoding/json"
	"fmt"
)

type reasonEnum struct {
	FormSubmissionGet     ReasonEnum
	FormSubmissionPost    ReasonEnum
	HTTPHeaderRefresh     ReasonEnum
	ScriptInitiated       ReasonEnum
	MetaTagRefresh        ReasonEnum
	PageBlockInterstitial ReasonEnum
	Reload                ReasonEnum
}

/*
Reason provides named acces to the ReasonEnum values.
*/
var Reason = reasonEnum{
	FormSubmissionGet:     reasonFormSubmissionGet,
	FormSubmissionPost:    reasonFormSubmissionPost,
	HTTPHeaderRefresh:     reasonHTTPHeaderRefresh,
	ScriptInitiated:       reasonScriptInitiated,
	MetaTagRefresh:        reasonMetaTagRefresh,
	PageBlockInterstitial: reasonPageBlockInterstitial,
	Reload:                reasonReload,
}

/*
ReasonEnum represents the reason for the navigation. Allowed values:
	- Reason.FormSubmissionGet     "formSubmissionGet"
	- Reason.FormSubmissionPost    "formSubmissionPost"
	- Reason.HTTPHeaderRefresh     "httpHeaderRefresh"
	- Reason.ScriptInitiated       "scriptInitiated"
	- Reason.MetaTagRefresh        "metaTagRefresh"
	- Reason.PageBlockInterstitial "pageBlockInterstitial"
	- Reason.Reload                "reload"

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
	// reasonHTTPHeaderRefresh represents the "httpHeaderRefresh" value.
	reasonHTTPHeaderRefresh
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
	reasonHTTPHeaderRefresh:     "httpHeaderRefresh",
	reasonScriptInitiated:       "scriptInitiated",
	reasonMetaTagRefresh:        "metaTagRefresh",
	reasonPageBlockInterstitial: "pageBlockInterstitial",
	reasonReload:                "reload",
}
