package network

import (
	"encoding/json"
	"fmt"
)

type challengeResponseEnum struct {
	Default            ChallengeResponseEnum
	CancelAuth         ChallengeResponseEnum
	ProvideCredentials ChallengeResponseEnum
}

/*
ChallengeResponse provides named acces to the ChallengeResponseEnum values.
*/
var ChallengeResponse = challengeResponseEnum{
	Default:            challengeResponseDefault,
	CancelAuth:         challengeResponseCancelAuth,
	ProvideCredentials: challengeResponseProvideCredentials,
}

/*
ChallengeResponseEnum represents the decision on what to do in response to the
authorization challenge. Default means deferring to the default behavior of the
net stack, which will likely either the Cancel authentication or display a popup
dialog box. Allowed values:
	- ChallengeResponse.Default            "Default"
	- ChallengeResponse.CancelAuth         "CancelAuth"
	- ChallengeResponse.ProvideCredentials "ProvideCredentials"

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallenge
*/
type ChallengeResponseEnum int

/*
String implements Stringer
*/
func (enum ChallengeResponseEnum) String() string {
	return _challengeResponseEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ChallengeResponseEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ChallengeResponseEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _challengeResponseEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// challengeResponseDefault represents the "Default" value.
	challengeResponseDefault ChallengeResponseEnum = iota + 1
	// challengeResponseCancelAuth represents the "CancelAuth" value.
	challengeResponseCancelAuth
	// challengeResponseProvideCredentials represents the "ProvideCredentials" value.
	challengeResponseProvideCredentials
)

var _challengeResponseEnums = map[ChallengeResponseEnum]string{
	challengeResponseDefault:            "Default",
	challengeResponseCancelAuth:         "CancelAuth",
	challengeResponseProvideCredentials: "ProvideCredentials",
}
