package tracing

import (
	"encoding/json"
	"fmt"
)

type transferModeEnum struct {
	ReportEvents   TransferModeEnum
	ReturnAsStream TransferModeEnum
}

var TransferMode = transferModeEnum{
	ReportEvents:   transferModeReportEvents,
	ReturnAsStream: transferModeReturnAsStream,
}

/*
Optional. Whether to report trace events as series of dataCollected events or to
save trace to a stream (defaults to `ReportEvents`). Allowed values:
	- TransferMode.ReportEvents
	- TransferMode.ReturnAsStream

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-TraceConfig
*/
type TransferModeEnum int

/*
String implements Stringer
*/
func (enum TransferModeEnum) String() string {
	return _transferModeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum TransferModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *TransferModeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _transferModeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// transferModeReportEvents represents the "ReportEvents" value.
	transferModeReportEvents TransferModeEnum = iota + 1
	// transferModeReturnAsStream represents the "ReturnAsStream" value.
	transferModeReturnAsStream
)

var _transferModeEnums = map[TransferModeEnum]string{
	TransferModeEnum(0):        "",
	transferModeReportEvents:   "ReportEvents",
	transferModeReturnAsStream: "ReturnAsStream",
}
