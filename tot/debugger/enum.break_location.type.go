package debugger

import (
	"encoding/json"
	"fmt"
)

type breakLocationTypeEnum struct {
	DebuggerStatement BreakLocationTypeEnum
	Call              BreakLocationTypeEnum
	Return            BreakLocationTypeEnum
}

/*
BreakLocationType provides named acces to the BreakLocationTypeEnum values.
*/
var BreakLocationType = breakLocationTypeEnum{
	DebuggerStatement: BreakLocationTypeDebuggerStatement,
	Call:              BreakLocationTypeCall,
	Return:            BreakLocationTypeReturn,
}

/*
BreakLocationTypeEnum is Optional. Allowed values:
	- BreakLocationType.DebuggerStatement "debuggerStatement"
	- BreakLocationType.Call              "call"
	- BreakLocationType.Return            "return"

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-BreakLocation
*/
type BreakLocationTypeEnum int

/*
String implements Stringer
*/
func (enum BreakLocationTypeEnum) String() string {
	return _breakLocationTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum BreakLocationTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *BreakLocationTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _breakLocationTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid breakLocationType value", bytes)
}

const (
	// BreakLocationTypeDebuggerStatement represents the "debuggerStatement" value.
	BreakLocationTypeDebuggerStatement BreakLocationTypeEnum = iota + 1
	// BreakLocationTypeCall represents the "call" value.
	BreakLocationTypeCall
	// BreakLocationTypeReturn represents the "return" value.
	BreakLocationTypeReturn
)

var _breakLocationTypeEnums = map[BreakLocationTypeEnum]string{
	BreakLocationTypeEnum(0):           "",
	BreakLocationTypeDebuggerStatement: "debuggerStatement",
	BreakLocationTypeCall:              "call",
	BreakLocationTypeReturn:            "return",
}
