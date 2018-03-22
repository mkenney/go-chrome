package log

import (
	"encoding/json"
	"fmt"
)

type sourceEnum struct {
	XML            SourceEnum
	Javascript     SourceEnum
	Network        SourceEnum
	Storage        SourceEnum
	Appcache       SourceEnum
	Rendering      SourceEnum
	Security       SourceEnum
	Deprecation    SourceEnum
	Worker         SourceEnum
	Violation      SourceEnum
	Intervention   SourceEnum
	Recommendation SourceEnum
	Other          SourceEnum
}

/*
Source provides named acces to the SourceEnum values.
*/
var Source = sourceEnum{
	XML:            sourceXML,
	Javascript:     sourceJavascript,
	Network:        sourceNetwork,
	Storage:        sourceStorage,
	Appcache:       sourceAppcache,
	Rendering:      sourceRendering,
	Security:       sourceSecurity,
	Deprecation:    sourceDeprecation,
	Worker:         sourceWorker,
	Violation:      sourceViolation,
	Intervention:   sourceIntervention,
	Recommendation: sourceRecommendation,
	Other:          sourceOther,
}

/*
SourceEnum represents the log entry source. Allowed values:
	- Source.XML            "xml"
	- Source.Javascript     "javascript"
	- Source.Network        "network"
	- Source.Storage        "storage"
	- Source.Appcache       "appcache"
	- Source.Rendering      "rendering"
	- Source.Security       "security"
	- Source.Deprecation    "deprecation"
	- Source.Worker         "worker"
	- Source.Violation      "violation"
	- Source.Intervention   "intervention"
	- Source.Recommendation "recommendation"
	- Source.Other          "other"

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-LogEntry
*/
type SourceEnum int

/*
String implements Stringer
*/
func (enum SourceEnum) String() string {
	return _sourceEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum SourceEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *SourceEnum) UnmarshalJSON(bytes []byte) error {
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

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// sourceXML represents the "xml" value.
	sourceXML SourceEnum = iota + 1
	// sourceJavascript represents the "javascript" value.
	sourceJavascript
	// sourceNetwork represents the "network" value.
	sourceNetwork
	// sourceStorage represents the "storage" value.
	sourceStorage
	// sourceAppcache represents the "appcache" value.
	sourceAppcache
	// sourceRendering represents the "rendering" value.
	sourceRendering
	// sourceSecurity represents the "security" value.
	sourceSecurity
	// sourceDeprecation represents the "deprecation" value.
	sourceDeprecation
	// sourceWorker represents the "worker" value.
	sourceWorker
	// sourceViolation represents the "violation" value.
	sourceViolation
	// sourceIntervention represents the "intervention" value.
	sourceIntervention
	// sourceRecommendation represents the "recommendation" value.
	sourceRecommendation
	// sourceOther represents the "other" value.
	sourceOther
)

var _sourceEnums = map[SourceEnum]string{
	sourceXML:            "xml",
	sourceJavascript:     "javascript",
	sourceNetwork:        "network",
	sourceStorage:        "storage",
	sourceAppcache:       "appcache",
	sourceRendering:      "rendering",
	sourceSecurity:       "security",
	sourceDeprecation:    "deprecation",
	sourceWorker:         "worker",
	sourceViolation:      "violation",
	sourceIntervention:   "intervention",
	sourceRecommendation: "recommendation",
	sourceOther:          "other",
}
