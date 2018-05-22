package log

import (
	"encoding/json"
	"fmt"
)

type levelEnum struct {
	Verbose LevelEnum
	Info    LevelEnum
	Warning LevelEnum
	Error   LevelEnum
}

/*
Level provides named acces to the LevelEnum values.
*/
var Level = levelEnum{
	Verbose: levelVerbose,
	Info:    levelInfo,
	Warning: levelWarning,
	Error:   levelError,
}

/*
LevelEnum represents the log entry severity. Allowed values:
	- Level.Verbose "verbose"
	- Level.Info    "info"
	- Level.Warning "warning"
	- Level.Error   "error"

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-LogEntry
*/
type LevelEnum int

/*
String implements Stringer
*/
func (enum LevelEnum) String() string {
	return _levelEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum LevelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *LevelEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _levelEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// levelVerbose represents the "verbose" value.
	levelVerbose LevelEnum = iota + 1
	// levelInfo represents the "info" value.
	levelInfo
	// levelWarning represents the "warning" value.
	levelWarning
	// levelError represents the "error" value.
	levelError
)

var _levelEnums = map[LevelEnum]string{
	levelVerbose: "verbose",
	levelInfo:    "info",
	levelWarning: "warning",
	levelError:   "error",
}
