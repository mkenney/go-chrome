package tracing

import (
	"encoding/json"
	"fmt"
)

type streamCompressionEnum struct {
	None StreamCompressionEnum
	Gzip StreamCompressionEnum
}

/*
StreamCompression provides named acces to the StreamCompressionEnum values.
*/
var StreamCompression = streamCompressionEnum{
	None: streamCompressionNone,
	Gzip: streamCompressionGzip,
}

/*
StreamCompressionEnum represents the compression type to use for traces returned
via streams. Allowed values:
	- StreamCompression.None "none"
	- StreamCompression.Gzip "gzip"

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-StreamCompression
*/
type StreamCompressionEnum int

/*
String implements Stringer
*/
func (enum StreamCompressionEnum) String() string {
	return _streamCompressionEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum StreamCompressionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *StreamCompressionEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _streamCompressionEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// streamCompressionNone represents the "none" value.
	streamCompressionNone StreamCompressionEnum = iota + 1
	// streamCompressionGzip represents the "gzip" value.
	streamCompressionGzip
)

var _streamCompressionEnums = map[StreamCompressionEnum]string{
	streamCompressionNone: "none",
	streamCompressionGzip: "gzip",
}
