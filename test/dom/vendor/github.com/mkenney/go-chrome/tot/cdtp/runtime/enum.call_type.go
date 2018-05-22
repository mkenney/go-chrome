package runtime

import (
	"encoding/json"
	"fmt"
)

type callTypeEnum struct {
	Log                 CallTypeEnum
	Debug               CallTypeEnum
	Info                CallTypeEnum
	Error               CallTypeEnum
	Warning             CallTypeEnum
	Dir                 CallTypeEnum
	Dirxml              CallTypeEnum
	Table               CallTypeEnum
	Trace               CallTypeEnum
	Clear               CallTypeEnum
	StartGroup          CallTypeEnum
	StartGroupCollapsed CallTypeEnum
	EndGroup            CallTypeEnum
	Assert              CallTypeEnum
	Profile             CallTypeEnum
	ProfileEnd          CallTypeEnum
	Count               CallTypeEnum
	TimeEnd             CallTypeEnum
}

/*
CallType provides named acces to the CallTypeEnum values.
*/
var CallType = callTypeEnum{
	Log:                 callTypeLog,
	Debug:               callTypeDebug,
	Info:                callTypeInfo,
	Error:               callTypeError,
	Warning:             callTypeWarning,
	Dir:                 callTypeDir,
	Dirxml:              callTypeDirxml,
	Table:               callTypeTable,
	Trace:               callTypeTrace,
	Clear:               callTypeClear,
	StartGroup:          callTypeStartGroup,
	StartGroupCollapsed: callTypeStartGroupCollapsed,
	EndGroup:            callTypeEndGroup,
	Assert:              callTypeAssert,
	Profile:             callTypeProfile,
	ProfileEnd:          callTypeProfileEnd,
	Count:               callTypeCount,
	TimeEnd:             callTypeTimeEnd,
}

/*
CallTypeEnum represents the type of the call. Allowed values:
	- CallType.Log                 "log"
	- CallType.Debug               "debug"
	- CallType.Info                "info"
	- CallType.Error               "error"
	- CallType.Warning             "warning"
	- CallType.Dir                 "dir"
	- CallType.Dirxml              "dirxml"
	- CallType.Table               "table"
	- CallType.Trace               "trace"
	- CallType.Clear               "clear"
	- CallType.StartGroup          "startGroup"
	- CallType.StartGroupCollapsed "startGroupCollapsed"
	- CallType.EndGroup            "endGroup"
	- CallType.Assert              "assert"
	- CallType.Profile             "profile"
	- CallType.ProfileEnd          "profileEnd"
	- CallType.Count               "count"
	- CallType.TimeEnd             "timeEnd"

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-consoleAPICalled
*/
type CallTypeEnum int

/*
String implements Stringer
*/
func (enum CallTypeEnum) String() string {
	return _callTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum CallTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *CallTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _callTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// callTypeLog represents the "log" value.
	callTypeLog CallTypeEnum = iota + 1
	// callTypeDebug represents the "debug" value.
	callTypeDebug
	// callTypeInfo represents the "info" value.
	callTypeInfo
	// callTypeError represents the "error" value.
	callTypeError
	// callTypeWarning represents the "warning" value.
	callTypeWarning
	// callTypeDir represents the "dir" value.
	callTypeDir
	// callTypeDirxml represents the "dirxml" value.
	callTypeDirxml
	// callTypeTable represents the "table" value.
	callTypeTable
	// callTypeTrace represents the "trace" value.
	callTypeTrace
	// callTypeClear represents the "clear" value.
	callTypeClear
	// callTypeStartGroup represents the "startGroup" value.
	callTypeStartGroup
	// callTypeStartGroupCollapsed represents the "startGroupCollapsed" value.
	callTypeStartGroupCollapsed
	// callTypeEndGroup represents the "endGroup" value.
	callTypeEndGroup
	// callTypeAssert represents the "assert" value.
	callTypeAssert
	// callTypeProfile represents the "profile" value.
	callTypeProfile
	// callTypeProfileEnd represents the "profileEnd" value.
	callTypeProfileEnd
	// callTypeCount represents the "count" value.
	callTypeCount
	// callTypeTimeEnd represents the "timeEnd" value.
	callTypeTimeEnd
)

var _callTypeEnums = map[CallTypeEnum]string{
	callTypeLog:                 "log",
	callTypeDebug:               "debug",
	callTypeInfo:                "info",
	callTypeError:               "error",
	callTypeWarning:             "warning",
	callTypeDir:                 "dir",
	callTypeDirxml:              "dirxml",
	callTypeTable:               "table",
	callTypeTrace:               "trace",
	callTypeClear:               "clear",
	callTypeStartGroup:          "startGroup",
	callTypeStartGroupCollapsed: "startGroupCollapsed",
	callTypeEndGroup:            "endGroup",
	callTypeAssert:              "assert",
	callTypeProfile:             "profile",
	callTypeProfileEnd:          "profileEnd",
	callTypeCount:               "count",
	callTypeTimeEnd:             "timeEnd",
}
