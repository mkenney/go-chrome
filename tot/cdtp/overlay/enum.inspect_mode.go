package overlay

import (
	"encoding/json"
	"fmt"
)

type inspectModeEnum struct {
	SearchForNode        InspectModeEnum
	SearchForUAShadowDOM InspectModeEnum
	None                 InspectModeEnum
}

var InspectMode = inspectModeEnum{
	SearchForNode:        inspectModeSearchForNode,
	SearchForUAShadowDOM: inspectModeSearchForUAShadowDOM,
	None:                 inspectModeNone,
}

/*
InspectMode is the inspect mode. Allowed values:
	- InspectMode.SearchForNode
	- InspectMode.SearchForUAShadowDOM
	- InspectMode.None

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-InspectMode
*/
type InspectModeEnum int

/*
String implements Stringer
*/
func (enum InspectModeEnum) String() string {
	return _inspectModeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum InspectModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *InspectModeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _inspectModeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// inspectModeSearchForNode represents the "searchForNode" value.
	inspectModeSearchForNode InspectModeEnum = iota + 1
	// inspectModeSearchForUAShadowDOM represents the "searchForUAShadowDOM" value.
	inspectModeSearchForUAShadowDOM
	// inspectModeNone represents the "none" value.
	inspectModeNone
)

var _inspectModeEnums = map[InspectModeEnum]string{
	inspectModeSearchForNode:        "searchForNode",
	inspectModeSearchForUAShadowDOM: "searchForUAShadowDOM",
	inspectModeNone:                 "none",
}
