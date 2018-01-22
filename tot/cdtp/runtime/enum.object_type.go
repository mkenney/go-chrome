package runtime

import (
	"encoding/json"
	"fmt"
)

type objectTypeEnum struct {
	Object    ObjectTypeEnum
	Function  ObjectTypeEnum
	Undefined ObjectTypeEnum
	String    ObjectTypeEnum
	Number    ObjectTypeEnum
	Boolean   ObjectTypeEnum
	Symbol    ObjectTypeEnum
	Accessor  ObjectTypeEnum
}

var ObjectType = objectTypeEnum{
	Object:    objectTypeObject,
	Function:  objectTypeFunction,
	Undefined: objectTypeUndefined,
	String:    objectTypeString,
	Number:    objectTypeNumber,
	Boolean:   objectTypeBoolean,
	Symbol:    objectTypeSymbol,
	Accessor:  objectTypeAccessor,
}

/*
Object type. For properties, "accessor" means that the property itself is an
accessor property. Allowed values:
	- ObjectType.Object    "object"
	- ObjectType.Function  "function"
	- ObjectType.Undefined "undefined"
	- ObjectType.String    "string"
	- ObjectType.Number    "number"
	- ObjectType.Boolean   "boolean"
	- ObjectType.Symbol    "symbol"
	- ObjectType.Accessor  "accessor"

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-RemoteObject
https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ObjectPreview
https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-PropertyPreview
*/
type ObjectTypeEnum int

/*
String implements Stringer
*/
func (enum ObjectTypeEnum) String() string {
	return _objectTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ObjectTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ObjectTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _objectTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// objectTypeObject represents the "object" value.
	objectTypeObject ObjectTypeEnum = iota + 1
	// objectTypeFunction represents the "function" value.
	objectTypeFunction
	// objectTypeUndefined represents the "undefined" value.
	objectTypeUndefined
	// objectTypeString represents the "string" value.
	objectTypeString
	// objectTypeNumber represents the "number" value.
	objectTypeNumber
	// objectTypeBoolean represents the "boolean" value.
	objectTypeBoolean
	// objectTypeSymbol represents the "symbol" value.
	objectTypeSymbol
	// objectTypeAccessor represents the "accessor" value.
	objectTypeAccessor
)

var _objectTypeEnums = map[ObjectTypeEnum]string{
	objectTypeObject:    "object",
	objectTypeFunction:  "function",
	objectTypeUndefined: "undefined",
	objectTypeString:    "string",
	objectTypeNumber:    "number",
	objectTypeBoolean:   "boolean",
	objectTypeSymbol:    "symbol",
	objectTypeAccessor:  "accessor",
}
