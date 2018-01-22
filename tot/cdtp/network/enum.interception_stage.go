package network

import (
	"encoding/json"
	"fmt"
)

type interceptionStageEnum struct {
	Request         InterceptionStageEnum
	HeadersReceived InterceptionStageEnum
}

/*
InterceptionStage provides named acces to the InterceptionStageEnum values.
*/
var InterceptionStage = interceptionStageEnum{
	Request:         interceptionStageRequest,
	HeadersReceived: interceptionStageHeadersReceived,
}

/*
InterceptionStageEnum represents stages of the interception to begin
intercepting. Request will intercept before the request is sent. Response will
intercept after the response is received. EXPERIMENTAL. Allowed values:
	- InterceptionStage.Request         "Request"
	- InterceptionStage.HeadersReceived "HeadersReceived"

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-InterceptionStage
*/
type InterceptionStageEnum int

/*
String implements Stringer
*/
func (enum InterceptionStageEnum) String() string {
	return _interceptionStageEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum InterceptionStageEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *InterceptionStageEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _interceptionStageEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// interceptionStageRequest represents the "Request" value.
	interceptionStageRequest InterceptionStageEnum = iota + 1
	// interceptionStageHeadersReceived represents the "HeadersReceived" value.
	interceptionStageHeadersReceived
)

var _interceptionStageEnums = map[InterceptionStageEnum]string{
	interceptionStageRequest:         "Request",
	interceptionStageHeadersReceived: "HeadersReceived",
}
