package network

import (
	"encoding/json"
	"fmt"
)

type sourceEnum struct {
	Server SourceEnum
	Proxy  SourceEnum
}

/*
Source provides named acces to the SourceEnum values.
*/
var Source = sourceEnum{
	Server: sourceServer,
	Proxy:  sourceProxy,
}

/*
SourceEnum is optional. Source of the authentication challenge. Allowed values:
	- Source.Server "Server"
	- Source.Proxy  "Proxy"

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallenge
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
	// sourceServer represents the "Server" value.
	sourceServer SourceEnum = iota + 1
	// sourceProxy represents the "Proxy" value.
	sourceProxy
)

var _sourceEnums = map[SourceEnum]string{
	SourceEnum(0): "",
	sourceServer:  "Server",
	sourceProxy:   "Proxy",
}
