package tracing

import (
	"encoding/json"
	"fmt"
)

type recordModeEnum struct {
	RecordUntilFull        RecordModeEnum
	RecordContinuously     RecordModeEnum
	RecordAsMuchAsPossible RecordModeEnum
	EchoToConsole          RecordModeEnum
}

var RecordMode = recordModeEnum{
	RecordUntilFull:        recordModeRecordUntilFull,
	RecordContinuously:     recordModeRecordContinuously,
	RecordAsMuchAsPossible: recordModeRecordAsMuchAsPossible,
	EchoToConsole:          recordModeEchoToConsole,
}

/*
Optional. Controls how the trace buffer stores data. Allowed values:
	- RecordMode.RecordUntilFull        "recordUntilFull"
	- RecordMode.RecordContinuously     "recordContinuously"
	- RecordMode.RecordAsMuchAsPossible "recordAsMuchAsPossible"
	- RecordMode.EchoToConsole          "echoToConsole"

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-TraceConfig
*/
type RecordModeEnum int

/*
String implements Stringer
*/
func (enum RecordModeEnum) String() string {
	return _recordModeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum RecordModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *RecordModeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _recordModeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// recordModeRecordUntilFull represents the "recordUntilFull" value.
	recordModeRecordUntilFull RecordModeEnum = iota + 1
	// recordModeRecordContinuously represents the "recordContinuously" value.
	recordModeRecordContinuously
	// recordModeRecordAsMuchAsPossible represents the "recordAsMuchAsPossible" value.
	recordModeRecordAsMuchAsPossible
	// recordModeEchoToConsole represents the "echoToConsole" value.
	recordModeEchoToConsole
)

var _recordModeEnums = map[RecordModeEnum]string{
	RecordModeEnum(0):                "",
	recordModeRecordUntilFull:        "recordUntilFull",
	recordModeRecordContinuously:     "recordContinuously",
	recordModeRecordAsMuchAsPossible: "recordAsMuchAsPossible",
	recordModeEchoToConsole:          "echoToConsole",
}
