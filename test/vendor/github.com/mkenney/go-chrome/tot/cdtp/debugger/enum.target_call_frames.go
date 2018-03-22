package debugger

import (
	"encoding/json"
	"fmt"
)

type targetCallFramesEnum struct {
	Any     TargetCallFramesEnum
	Current TargetCallFramesEnum
}

/*
TargetCallFrames provides named acces to the TargetCallFramesEnum values.
*/
var TargetCallFrames = targetCallFramesEnum{
	Any:     targetCallFramesAny,
	Current: targetCallFramesCurrent,
}

/*
TargetCallFramesEnum is optional. Allowed values:
	- TargetCallFrames.Any     "any"
	- TargetCallFrames.Current "current"

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
*/
type TargetCallFramesEnum int

/*
String implements Stringer
*/
func (enum TargetCallFramesEnum) String() string {
	return _targetCallFramesEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum TargetCallFramesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *TargetCallFramesEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _targetCallFramesEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid targetCallFrames value", bytes)
}

const (
	// BreakLocationTypeDebuggerStatement represents the "debuggerStatement" value.
	targetCallFramesAny TargetCallFramesEnum = iota + 1
	// targetCallFramesCall represents the "call" value.
	targetCallFramesCurrent
)

var _targetCallFramesEnums = map[TargetCallFramesEnum]string{
	TargetCallFramesEnum(0): "",
	targetCallFramesAny:     "any",
	targetCallFramesCurrent: "current",
}
