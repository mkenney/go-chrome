package network

import (
	"encoding/json"
	"fmt"
)

type referrerPolicyEnum struct {
	UnsafeUrl                   ReferrerPolicyEnum
	NoReferrerWhenDowngrade     ReferrerPolicyEnum
	NoReferrer                  ReferrerPolicyEnum
	Origin                      ReferrerPolicyEnum
	OriginWhenCrossOrigin       ReferrerPolicyEnum
	SameOrigin                  ReferrerPolicyEnum
	StrictOrigin                ReferrerPolicyEnum
	StrictOriginWhenCrossOrigin ReferrerPolicyEnum
}

var ReferrerPolicy = referrerPolicyEnum{
	UnsafeUrl:                   referrerPolicyUnsafeUrl,
	NoReferrerWhenDowngrade:     referrerPolicyNoReferrerWhenDowngrade,
	NoReferrer:                  referrerPolicyNoReferrer,
	Origin:                      referrerPolicyOrigin,
	OriginWhenCrossOrigin:       referrerPolicyOriginWhenCrossOrigin,
	SameOrigin:                  referrerPolicySameOrigin,
	StrictOrigin:                referrerPolicyStrictOrigin,
	StrictOriginWhenCrossOrigin: referrerPolicyStrictOriginWhenCrossOrigin,
}

/*
The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
Allowed values:
	- ReferrerPolicy.UnsafeUrl                   "unsafe-url"
	- ReferrerPolicy.NoReferrerWhenDowngrade     "no-referrer-when-downgrade"
	- ReferrerPolicy.NoReferrer                  "no-referrer"
	- ReferrerPolicy.Origin                      "origin"
	- ReferrerPolicy.OriginWhenCrossOrigin       "origin-when-cross-origin"
	- ReferrerPolicy.SameOrigin                  "same-origin"
	- ReferrerPolicy.StrictOrigin                "strict-origin"
	- ReferrerPolicy.StrictOriginWhenCrossOrigin "strict-origin-when-cross-origin"

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ResourcePriority
*/
type ReferrerPolicyEnum int

/*
String implements Stringer
*/
func (enum ReferrerPolicyEnum) String() string {
	return _referrerPolicyEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ReferrerPolicyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ReferrerPolicyEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _referrerPolicyEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// referrerPolicyUnsafeUrl represents the "unsafe-url" value.
	referrerPolicyUnsafeUrl ReferrerPolicyEnum = iota + 1
	// referrerPolicyNoReferrerWhenDowngrade represents the "no-referrer-when-downgrade" value.
	referrerPolicyNoReferrerWhenDowngrade
	// referrerPolicyNoReferrer represents the "no-referrer" value.
	referrerPolicyNoReferrer
	// referrerPolicyOrigin represents the "origin" value.
	referrerPolicyOrigin
	// referrerPolicyOriginWhenCrossOrigin represents the "origin-when-cross-origin" value.
	referrerPolicyOriginWhenCrossOrigin
	// referrerPolicySameOrigin represents the "same-origin" value.
	referrerPolicySameOrigin
	// referrerPolicyStrictOrigin represents the "strict-origin" value.
	referrerPolicyStrictOrigin
	// referrerPolicyStrictOriginWhenCrossOrigin represents the "strict-origin-when-cross-origin" value.
	referrerPolicyStrictOriginWhenCrossOrigin
)

var _referrerPolicyEnums = map[ReferrerPolicyEnum]string{
	referrerPolicyUnsafeUrl:                   "unsafe-url",
	referrerPolicyNoReferrerWhenDowngrade:     "no-referrer-when-downgrade",
	referrerPolicyNoReferrer:                  "no-referrer",
	referrerPolicyOrigin:                      "origin",
	referrerPolicyOriginWhenCrossOrigin:       "origin-when-cross-origin",
	referrerPolicySameOrigin:                  "same-origin",
	referrerPolicyStrictOrigin:                "strict-origin",
	referrerPolicyStrictOriginWhenCrossOrigin: "strict-origin-when-cross-origin",
}
