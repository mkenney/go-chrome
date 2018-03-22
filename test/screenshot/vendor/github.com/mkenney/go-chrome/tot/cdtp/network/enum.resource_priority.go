package network

import (
	"encoding/json"
	"fmt"
)

type resourcePriorityEnum struct {
	VeryLow  ResourcePriorityEnum
	Low      ResourcePriorityEnum
	Medium   ResourcePriorityEnum
	High     ResourcePriorityEnum
	VeryHigh ResourcePriorityEnum
}

/*
ResourcePriority provides named acces to the ResourcePriorityEnum values.
*/
var ResourcePriority = resourcePriorityEnum{
	VeryLow:  resourcePriorityVeryLow,
	Low:      resourcePriorityLow,
	Medium:   resourcePriorityMedium,
	High:     resourcePriorityHigh,
	VeryHigh: resourcePriorityVeryHigh,
}

/*
ResourcePriorityEnum represents the loading priority of a resource request.
Allowed values:
	- ResourcePriority.VeryLow  "VeryLow"
	- ResourcePriority.Low      "Low"
	- ResourcePriority.Medium   "Medium"
	- ResourcePriority.High     "High"
	- ResourcePriority.VeryHigh "VeryHigh"

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ResourcePriority
*/
type ResourcePriorityEnum int

/*
String implements Stringer
*/
func (enum ResourcePriorityEnum) String() string {
	return _resourcePriorityEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ResourcePriorityEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ResourcePriorityEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _resourcePriorityEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// resourcePriorityVeryLow represents the "VeryLow" value.
	resourcePriorityVeryLow ResourcePriorityEnum = iota + 1
	// resourcePriorityLow represents the "Low" value.
	resourcePriorityLow
	// resourcePriorityMedium represents the "Medium" value.
	resourcePriorityMedium
	// resourcePriorityHigh represents the "High" value.
	resourcePriorityHigh
	// resourcePriorityVeryHigh represents the "VeryHigh" value.
	resourcePriorityVeryHigh
)

var _resourcePriorityEnums = map[ResourcePriorityEnum]string{
	resourcePriorityVeryLow:  "VeryLow",
	resourcePriorityLow:      "Low",
	resourcePriorityMedium:   "Medium",
	resourcePriorityHigh:     "High",
	resourcePriorityVeryHigh: "VeryHigh",
}
