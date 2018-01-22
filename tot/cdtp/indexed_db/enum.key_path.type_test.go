package indexedDB

import (
	"encoding/json"
	"testing"
)

func TestEnumKeyPathType(t *testing.T) {
	var enum KeyPathTypeEnum
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

	enum = KeyPathType.Null
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"null"` != string(result) {
		t.Errorf("Expected '\"null\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"null"`), &enum)
	if KeyPathType.Null != enum {
		t.Errorf("Expcected %d, got %d", KeyPathType.Null, enum)
	}

	enum = KeyPathType.String
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"string"` != string(result) {
		t.Errorf("Expected '\"string\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"string"`), &enum)
	if KeyPathType.String != enum {
		t.Errorf("Expcected %d, got %d", KeyPathType.String, enum)
	}

	enum = KeyPathType.Array
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"array"` != string(result) {
		t.Errorf("Expected '\"array\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"array"`), &enum)
	if KeyPathType.Array != enum {
		t.Errorf("Expcected %d, got %d", KeyPathType.Array, enum)
	}
}
