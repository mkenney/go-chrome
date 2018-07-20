package page

import (
	"encoding/json"
	"fmt"
)

type dialogTypeEnum struct {
	Alert        DialogTypeEnum
	Confirm      DialogTypeEnum
	Prompt       DialogTypeEnum
	Beforeunload DialogTypeEnum
}

/*
DialogType provides named acces to the DialogTypeEnum values.
*/
var DialogType = dialogTypeEnum{
	Alert:        dialogTypeAlert,
	Confirm:      dialogTypeConfirm,
	Prompt:       dialogTypePrompt,
	Beforeunload: dialogTypeBeforeunload,
}

/*
DialogTypeEnum defines the Javascript dialog type. Allowed values:
	- DialogType.Alert        "alert"
	- DialogType.Confirm      "confirm"
	- DialogType.Prompt       "prompt"
	- DialogType.Beforeunload "beforeunload"

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-DialogType
*/
type DialogTypeEnum int

/*
String implements Stringer
*/
func (enum DialogTypeEnum) String() string {
	return _dialogTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum DialogTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *DialogTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _dialogTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// dialogTypeAlert represents the "alert" value.
	dialogTypeAlert DialogTypeEnum = iota + 1
	// dialogTypeConfirm represents the "confirm" value.
	dialogTypeConfirm
	// dialogTypePrompt represents the "prompt" value.
	dialogTypePrompt
	// dialogTypeBeforeunload represents the "beforeunload" value.
	dialogTypeBeforeunload
)

var _dialogTypeEnums = map[DialogTypeEnum]string{
	dialogTypeAlert:        "alert",
	dialogTypeConfirm:      "confirm",
	dialogTypePrompt:       "prompt",
	dialogTypeBeforeunload: "beforeunload",
}
