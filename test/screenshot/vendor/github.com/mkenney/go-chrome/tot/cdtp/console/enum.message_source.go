package console

import (
	"encoding/json"
	"fmt"
)

type messageSourceEnum struct {
	XML         MessageSourceEnum
	Javascript  MessageSourceEnum
	Network     MessageSourceEnum
	ConsoleAPI  MessageSourceEnum
	Storage     MessageSourceEnum
	Appcache    MessageSourceEnum
	Rendering   MessageSourceEnum
	Security    MessageSourceEnum
	Other       MessageSourceEnum
	Deprecation MessageSourceEnum
	Worker      MessageSourceEnum
}

/*
MessageSource provides named acces to the MessageSourceEnum values.
*/
var MessageSource = messageSourceEnum{
	XML:         MessageSourceXML,
	Javascript:  MessageSourceJavascript,
	Network:     MessageSourceNetwork,
	ConsoleAPI:  MessageSourceConsoleAPI,
	Storage:     MessageSourceStorage,
	Appcache:    MessageSourceAppcache,
	Rendering:   MessageSourceRendering,
	Security:    MessageSourceSecurity,
	Other:       MessageSourceOther,
	Deprecation: MessageSourceDeprecation,
	Worker:      MessageSourceWorker,
}

/*
MessageSourceEnum represents the message source. Allowed values:
	- MessageSource.SourceXML         "xml"
	- MessageSource.SourceJavascript  "javascript"
	- MessageSource.SourceNetwork     "network"
	- MessageSource.SourceConsoleAPI  "console-api"
	- MessageSource.SourceStorage     "storage"
	- MessageSource.SourceAppcache    "appcache"
	- MessageSource.SourceRendering   "rendering"
	- MessageSource.SourceSecurity    "security"
	- MessageSource.SourceOther       "other"
	- MessageSource.SourceDeprecation "deprecation"
	- MessageSource.SourceWorker      "worker"
*/
type MessageSourceEnum int

/*
String implements Stringer
*/
func (enum MessageSourceEnum) String() string {
	return _sourceEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum MessageSourceEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *MessageSourceEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _sourceEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid source value", bytes)
}

const (
	// MessageSourceXML represents the "xml" value
	MessageSourceXML MessageSourceEnum = iota + 1
	// MessageSourceJavascript represents the "javascript" value
	MessageSourceJavascript
	// MessageSourceNetwork represents the "network" value
	MessageSourceNetwork
	// MessageSourceConsoleAPI represents the "console-api" value
	MessageSourceConsoleAPI
	// MessageSourceStorage represents the "storage" value
	MessageSourceStorage
	// MessageSourceAppcache represents the "appcache" value
	MessageSourceAppcache
	// MessageSourceRendering represents the "rendering" value
	MessageSourceRendering
	// MessageSourceSecurity represents the "security" value
	MessageSourceSecurity
	// MessageSourceOther represents the "other" value
	MessageSourceOther
	// MessageSourceDeprecation represents the "deprecation" value
	MessageSourceDeprecation
	// MessageSourceWorker represents the "worker" value
	MessageSourceWorker
)

var _sourceEnums = map[MessageSourceEnum]string{
	MessageSourceXML:         "xml",
	MessageSourceJavascript:  "javascript",
	MessageSourceNetwork:     "network",
	MessageSourceConsoleAPI:  "console-api",
	MessageSourceStorage:     "storage",
	MessageSourceAppcache:    "appcache",
	MessageSourceRendering:   "rendering",
	MessageSourceSecurity:    "security",
	MessageSourceOther:       "other",
	MessageSourceDeprecation: "deprecation",
	MessageSourceWorker:      "worker",
}
