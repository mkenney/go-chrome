package console

import (
	"encoding/json"
	"fmt"
)

type messageLevelEnum struct {
	Log     MessageLevelEnum
	Warning MessageLevelEnum
	Error   MessageLevelEnum
	Debug   MessageLevelEnum
	Info    MessageLevelEnum
}

/*
MessageLevel provides named acces to the MessageLevelEnum values.
*/
var MessageLevel = messageLevelEnum{
	Log:     MessageLevelLog,
	Warning: MessageLevelWarning,
	Error:   MessageLevelError,
	Debug:   MessageLevelDebug,
	Info:    MessageLevelInfo,
}

/*
MessageLevelEnum represents the message severity. Allowed values:
	- MessageLevel.Log     "log"
	- MessageLevel.Warning "warning"
	- MessageLevel.Error   "error"
	- MessageLevel.Debug   "debug"
	- MessageLevel.Info    "info"

https://chromedevtools.github.io/devtools-protocol/tot/Console/#type-ConsoleMessage
*/
type MessageLevelEnum int

/*
String implements Stringer
*/
func (enum MessageLevelEnum) String() string {
	return _messageLevelEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum MessageLevelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *MessageLevelEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _messageLevelEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid source value", bytes)
}

const (
	// MessageLevelLog represents the "log" value.
	MessageLevelLog MessageLevelEnum = iota + 1
	// MessageLevelWarning represents the "warning" value.
	MessageLevelWarning
	// MessageLevelError represents the "error" value.
	MessageLevelError
	// MessageLevelDebug represents the "debug" value.
	MessageLevelDebug
	// MessageLevelInfo represents the "info" value.
	MessageLevelInfo
)

var _messageLevelEnums = map[MessageLevelEnum]string{
	MessageLevelLog:     "log",
	MessageLevelWarning: "warning",
	MessageLevelError:   "error",
	MessageLevelDebug:   "debug",
	MessageLevelInfo:    "info",
}
