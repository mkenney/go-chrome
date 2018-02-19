package db

import (
	"encoding/json"
	"testing"
)

func TestEnumKeyType(t *testing.T) {
	var enum KeyTypeEnum
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

	enum = KeyType.Number
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"number"` != string(result) {
		t.Errorf("Expected '\"number\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"number"`), &enum)
	if KeyType.Number != enum {
		t.Errorf("Expcected %d, got %d", KeyType.Number, enum)
	}

	enum = KeyType.String
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"string"` != string(result) {
		t.Errorf("Expected '\"string\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"string"`), &enum)
	if KeyType.String != enum {
		t.Errorf("Expcected %d, got %d", KeyType.String, enum)
	}

	enum = KeyType.Date
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"date"` != string(result) {
		t.Errorf("Expected '\"date\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"date"`), &enum)
	if KeyType.Date != enum {
		t.Errorf("Expcected %d, got %d", KeyType.Date, enum)
	}
}

func TestEnumKeyType2(t *testing.T) {
	var enum KeyTypeEnum
	var err error
	var result []byte

	enum = KeyType.Array
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"array"` != string(result) {
		t.Errorf("Expected '\"array\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"array"`), &enum)
	if KeyType.Array != enum {
		t.Errorf("Expcected %d, got %d", KeyType.Array, enum)
	}
}
