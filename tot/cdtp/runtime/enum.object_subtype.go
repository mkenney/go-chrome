package runtime

import (
	"encoding/json"
	"fmt"
)

type objectSubtypeEnum struct {
	Array      ObjectSubtypeEnum
	Null       ObjectSubtypeEnum
	Node       ObjectSubtypeEnum
	Regexp     ObjectSubtypeEnum
	Date       ObjectSubtypeEnum
	Map        ObjectSubtypeEnum
	Set        ObjectSubtypeEnum
	Weakmap    ObjectSubtypeEnum
	Weakset    ObjectSubtypeEnum
	Iterator   ObjectSubtypeEnum
	Generator  ObjectSubtypeEnum
	Error      ObjectSubtypeEnum
	Proxy      ObjectSubtypeEnum
	Promise    ObjectSubtypeEnum
	Typedarray ObjectSubtypeEnum
}

var ObjectSubtype = objectSubtypeEnum{
	Array:      objectSubtypeArray,
	Null:       objectSubtypeNull,
	Node:       objectSubtypeNode,
	Regexp:     objectSubtypeRegexp,
	Date:       objectSubtypeDate,
	Map:        objectSubtypeMap,
	Set:        objectSubtypeSet,
	Weakmap:    objectSubtypeWeakmap,
	Weakset:    objectSubtypeWeakset,
	Iterator:   objectSubtypeIterator,
	Generator:  objectSubtypeGenerator,
	Error:      objectSubtypeError,
	Proxy:      objectSubtypeProxy,
	Promise:    objectSubtypePromise,
	Typedarray: objectSubtypeTypedarray,
}

/*
Object subtype hint. Specified for object type values only. Allowed values:
	- ObjectSubtype.Array      "array"
	- ObjectSubtype.Null       "null"
	- ObjectSubtype.Node       "node"
	- ObjectSubtype.Regexp     "regexp"
	- ObjectSubtype.Date       "date"
	- ObjectSubtype.Map        "map"
	- ObjectSubtype.Set        "set"
	- ObjectSubtype.Weakmap    "weakmap"
	- ObjectSubtype.Weakset    "weakset"
	- ObjectSubtype.Iterator   "iterator"
	- ObjectSubtype.Generator  "generator"
	- ObjectSubtype.Error      "error"
	- ObjectSubtype.Proxy      "proxy"
	- ObjectSubtype.Promise    "promise"
	- ObjectSubtype.Typedarray "typedarray"

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-RemoteObject
https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ObjectPreview
https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-PropertyPreview
*/
type ObjectSubtypeEnum int

/*
String implements Stringer
*/
func (enum ObjectSubtypeEnum) String() string {
	return _objectSubtypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ObjectSubtypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ObjectSubtypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _objectSubtypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// objectSubtypeArray represents the "array" value.
	objectSubtypeArray ObjectSubtypeEnum = iota + 1
	// objectSubtypeNull represents the "null" value.
	objectSubtypeNull
	// objectSubtypeNode represents the "node" value.
	objectSubtypeNode
	// objectSubtypeRegexp represents the "regexp" value.
	objectSubtypeRegexp
	// objectSubtypeDate represents the "date" value.
	objectSubtypeDate
	// objectSubtypeMap represents the "map" value.
	objectSubtypeMap
	// objectSubtypeSet represents the "set" value.
	objectSubtypeSet
	// objectSubtypeWeakmap represents the "weakmap" value.
	objectSubtypeWeakmap
	// objectSubtypeWeakset represents the "weakset" value.
	objectSubtypeWeakset
	// objectSubtypeIterator represents the "iterator" value.
	objectSubtypeIterator
	// objectSubtypeGenerator represents the "generator" value.
	objectSubtypeGenerator
	// objectSubtypeError represents the "error" value.
	objectSubtypeError
	// objectSubtypeProxy represents the "proxy" value.
	objectSubtypeProxy
	// objectSubtypePromise represents the "promise" value.
	objectSubtypePromise
	// objectSubtypeTypedarray represents the "typedarray" value.
	objectSubtypeTypedarray
)

var _objectSubtypeEnums = map[ObjectSubtypeEnum]string{
	ObjectSubtypeEnum(0):    "",
	objectSubtypeArray:      "array",
	objectSubtypeNull:       "null",
	objectSubtypeNode:       "node",
	objectSubtypeRegexp:     "regexp",
	objectSubtypeDate:       "date",
	objectSubtypeMap:        "map",
	objectSubtypeSet:        "set",
	objectSubtypeWeakmap:    "weakmap",
	objectSubtypeWeakset:    "weakset",
	objectSubtypeIterator:   "iterator",
	objectSubtypeGenerator:  "generator",
	objectSubtypeError:      "error",
	objectSubtypeProxy:      "proxy",
	objectSubtypePromise:    "promise",
	objectSubtypeTypedarray: "typedarray",
}
