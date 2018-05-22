package security

import (
	"encoding/json"
	"fmt"
)

type certificateErrorActionEnum struct {
	Continue CertificateErrorActionEnum
	Cancel   CertificateErrorActionEnum
}

/*
CertificateErrorAction provides named acces to the CertificateErrorActionEnum values.
*/
var CertificateErrorAction = certificateErrorActionEnum{
	Continue: certificateErrorActionContinue,
	Cancel:   certificateErrorActionCancel,
}

/*
CertificateErrorActionEnum describes the action to take when a certificate error
occurs. 'continue' will continue processing the request and 'cancel' will cancel
the request. Allowed values:
	- CertificateErrorAction.Continue "continue"
	- CertificateErrorAction.Cancel   "cancel"

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-CertificateErrorAction
*/
type CertificateErrorActionEnum int

/*
String implements Stringer
*/
func (enum CertificateErrorActionEnum) String() string {
	return _certificateErrorActionEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum CertificateErrorActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *CertificateErrorActionEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _certificateErrorActionEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// certificateErrorActionContinue represents the "continue" value.
	certificateErrorActionContinue CertificateErrorActionEnum = iota + 1
	// certificateErrorActionCancel represents the "cancel" value.
	certificateErrorActionCancel
)

var _certificateErrorActionEnums = map[CertificateErrorActionEnum]string{
	certificateErrorActionContinue: "continue",
	certificateErrorActionCancel:   "cancel",
}
