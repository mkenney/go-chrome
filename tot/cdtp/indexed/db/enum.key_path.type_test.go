package db

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/indexed/db"
)

func TestEnumKeyPathType(t *testing.T) {
	var enum db.KeyPathTypeEnum
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

	enum = db.KeyPathType.Null
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"null"` != string(result) {
		t.Errorf("Expected '\"null\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"null"`), &enum)
	if db.KeyPathType.Null != enum {
		t.Errorf("Expcected %d, got %d", db.KeyPathType.Null, enum)
	}

	enum = db.KeyPathType.String
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"string"` != string(result) {
		t.Errorf("Expected '\"string\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"string"`), &enum)
	if db.KeyPathType.String != enum {
		t.Errorf("Expcected %d, got %d", db.KeyPathType.String, enum)
	}

	enum = db.KeyPathType.Array
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"array"` != string(result) {
		t.Errorf("Expected '\"array\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"array"`), &enum)
	if db.KeyPathType.Array != enum {
		t.Errorf("Expcected %d, got %d", db.KeyPathType.Array, enum)
	}
}
