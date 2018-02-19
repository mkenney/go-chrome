package runtime

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestEnumObjectType(t *testing.T) {
	var enum runtime.ObjectTypeEnum
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

	enum = runtime.ObjectType.Object
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"object"` != string(result) {
		t.Errorf("Expected '\"object\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"object"`), &enum)
	if runtime.ObjectType.Object != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.Object, enum)
	}

	enum = runtime.ObjectType.Function
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"function"` != string(result) {
		t.Errorf("Expected '\"function\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"function"`), &enum)
	if runtime.ObjectType.Function != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.Function, enum)
	}

	enum = runtime.ObjectType.Undefined
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"undefined"` != string(result) {
		t.Errorf("Expected '\"undefined\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"undefined"`), &enum)
	if runtime.ObjectType.Undefined != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.Undefined, enum)
	}
}

func TestEnumObjectType2(t *testing.T) {
	var enum runtime.ObjectTypeEnum
	var err error
	var result []byte

	enum = runtime.ObjectType.String
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"string"` != string(result) {
		t.Errorf("Expected '\"string\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"string"`), &enum)
	if runtime.ObjectType.String != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.String, enum)
	}

	enum = runtime.ObjectType.Number
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"number"` != string(result) {
		t.Errorf("Expected '\"number\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"number"`), &enum)
	if runtime.ObjectType.Number != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.Number, enum)
	}

	enum = runtime.ObjectType.Boolean
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"boolean"` != string(result) {
		t.Errorf("Expected '\"boolean\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"boolean"`), &enum)
	if runtime.ObjectType.Boolean != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.Boolean, enum)
	}

	enum = runtime.ObjectType.Symbol
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"symbol"` != string(result) {
		t.Errorf("Expected '\"symbol\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"symbol"`), &enum)
	if runtime.ObjectType.Symbol != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.Symbol, enum)
	}
}

func TestEnumObjectType3(t *testing.T) {
	var enum runtime.ObjectTypeEnum
	var err error
	var result []byte

	enum = runtime.ObjectType.Accessor
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"accessor"` != string(result) {
		t.Errorf("Expected '\"accessor\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"accessor"`), &enum)
	if runtime.ObjectType.Accessor != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectType.Accessor, enum)
	}
}
