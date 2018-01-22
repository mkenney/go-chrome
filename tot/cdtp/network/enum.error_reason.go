package network

import (
	"encoding/json"
	"fmt"
)

type errorReasonEnum struct {
	Failed               ErrorReasonEnum
	Aborted              ErrorReasonEnum
	TimedOut             ErrorReasonEnum
	AccessDenied         ErrorReasonEnum
	ConnectionClosed     ErrorReasonEnum
	ConnectionReset      ErrorReasonEnum
	ConnectionRefused    ErrorReasonEnum
	ConnectionAborted    ErrorReasonEnum
	ConnectionFailed     ErrorReasonEnum
	NameNotResolved      ErrorReasonEnum
	InternetDisconnected ErrorReasonEnum
	AddressUnreachable   ErrorReasonEnum
}

var ErrorReason = errorReasonEnum{
	Failed:               errorReasonFailed,
	Aborted:              errorReasonAborted,
	TimedOut:             errorReasonTimedOut,
	AccessDenied:         errorReasonAccessDenied,
	ConnectionClosed:     errorReasonConnectionClosed,
	ConnectionReset:      errorReasonConnectionReset,
	ConnectionRefused:    errorReasonConnectionRefused,
	ConnectionAborted:    errorReasonConnectionAborted,
	ConnectionFailed:     errorReasonConnectionFailed,
	NameNotResolved:      errorReasonNameNotResolved,
	InternetDisconnected: errorReasonInternetDisconnected,
	AddressUnreachable:   errorReasonAddressUnreachable,
}

/*
ErrorReasonEnum is the network level fetch failure reason. Allowed values:
	- ErrorReason.Failed
	- ErrorReason.Aborted
	- ErrorReason.TimedOut
	- ErrorReason.AccessDenied
	- ErrorReason.ConnectionClosed
	- ErrorReason.ConnectionReset
	- ErrorReason.ConnectionRefused
	- ErrorReason.ConnectionAborted
	- ErrorReason.ConnectionFailed
	- ErrorReason.NameNotResolved
	- ErrorReason.InternetDisconnected
	- ErrorReason.AddressUnreachable

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ErrorReason
*/
type ErrorReasonEnum int

/*
String implements Stringer
*/
func (enum ErrorReasonEnum) String() string {
	return _errorReasonEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ErrorReasonEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ErrorReasonEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _errorReasonEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// errorReasonFailed represents the "Failed" value.
	errorReasonFailed ErrorReasonEnum = iota + 1
	// errorReasonAborted represents the "Aborted" value.
	errorReasonAborted
	// errorReasonTimedOut represents the "TimedOut" value.
	errorReasonTimedOut
	// errorReasonAccessDenied represents the "AccessDenied" value.
	errorReasonAccessDenied
	// errorReasonConnectionClosed represents the "ConnectionClosed" value.
	errorReasonConnectionClosed
	// errorReasonConnectionReset represents the "ConnectionReset" value.
	errorReasonConnectionReset
	// errorReasonConnectionRefused represents the "ConnectionRefused" value.
	errorReasonConnectionRefused
	// errorReasonConnectionAborted represents the "ConnectionAborted" value.
	errorReasonConnectionAborted
	// errorReasonConnectionFailed represents the "ConnectionFailed" value.
	errorReasonConnectionFailed
	// errorReasonNameNotResolved represents the "NameNotResolved" value.
	errorReasonNameNotResolved
	// errorReasonInternetDisconnected represents the "InternetDisconnected" value.
	errorReasonInternetDisconnected
	// errorReasonAddressUnreachable represents the "AddressUnreachable" value.
	errorReasonAddressUnreachable
)

var _errorReasonEnums = map[ErrorReasonEnum]string{
	errorReasonFailed:               "Failed",
	errorReasonAborted:              "Aborted",
	errorReasonTimedOut:             "TimedOut",
	errorReasonAccessDenied:         "AccessDenied",
	errorReasonConnectionClosed:     "ConnectionClosed",
	errorReasonConnectionReset:      "ConnectionReset",
	errorReasonConnectionRefused:    "ConnectionRefused",
	errorReasonConnectionAborted:    "ConnectionAborted",
	errorReasonConnectionFailed:     "ConnectionFailed",
	errorReasonNameNotResolved:      "NameNotResolved",
	errorReasonInternetDisconnected: "InternetDisconnected",
	errorReasonAddressUnreachable:   "AddressUnreachable",
}
