package runtime

import (
	"encoding/json"
	"testing"
)

func TestEnumObjectType(t *testing.T) {
	var enum ObjectTypeEnum
	var err error
	var result []byte

	err = json.Unmarshal([]byte(`""`), &enum)
	if nil == err {
		t.Errorf("Expected error, got nil")
	}

	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `""` != string(result) {
		t.Errorf("Expected empty JSON string, got '%s'", result)
	}

	enum = ObjectType.Object
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"object"` != string(result) {
		t.Errorf("Expected '\"object\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"object"`), &enum)
	if ObjectType.Object != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.Object, enum)
	}

	enum = ObjectType.Function
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"function"` != string(result) {
		t.Errorf("Expected '\"function\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"function"`), &enum)
	if ObjectType.Function != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.Function, enum)
	}

	enum = ObjectType.Undefined
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"undefined"` != string(result) {
		t.Errorf("Expected '\"undefined\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"undefined"`), &enum)
	if ObjectType.Undefined != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.Undefined, enum)
	}
}

func TestEnumObjectType2(t *testing.T) {
	var enum ObjectTypeEnum
	var err error
	var result []byte

	enum = ObjectType.String
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"string"` != string(result) {
		t.Errorf("Expected '\"string\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"string"`), &enum)
	if ObjectType.String != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.String, enum)
	}

	enum = ObjectType.Number
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"number"` != string(result) {
		t.Errorf("Expected '\"number\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"number"`), &enum)
	if ObjectType.Number != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.Number, enum)
	}

	enum = ObjectType.Boolean
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"boolean"` != string(result) {
		t.Errorf("Expected '\"boolean\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"boolean"`), &enum)
	if ObjectType.Boolean != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.Boolean, enum)
	}

	enum = ObjectType.Symbol
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"symbol"` != string(result) {
		t.Errorf("Expected '\"symbol\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"symbol"`), &enum)
	if ObjectType.Symbol != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.Symbol, enum)
	}
}

func TestEnumObjectType3(t *testing.T) {
	var enum ObjectTypeEnum
	var err error
	var result []byte

	enum = ObjectType.Accessor
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"accessor"` != string(result) {
		t.Errorf("Expected '\"accessor\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"accessor"`), &enum)
	if ObjectType.Accessor != enum {
		t.Errorf("Expcected %d, got %d", ObjectType.Accessor, enum)
	}
}
