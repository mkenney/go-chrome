package db

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/indexed/db"
)

func TestEnumKeyType(t *testing.T) {
	var enum db.KeyTypeEnum
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

	enum = db.KeyType.Number
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"number"` != string(result) {
		t.Errorf("Expected '\"number\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"number"`), &enum)
	if db.KeyType.Number != enum {
		t.Errorf("Expcected %d, got %d", db.KeyType.Number, enum)
	}

	enum = db.KeyType.String
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"string"` != string(result) {
		t.Errorf("Expected '\"string\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"string"`), &enum)
	if db.KeyType.String != enum {
		t.Errorf("Expcected %d, got %d", db.KeyType.String, enum)
	}

	enum = db.KeyType.Date
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"date"` != string(result) {
		t.Errorf("Expected '\"date\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"date"`), &enum)
	if db.KeyType.Date != enum {
		t.Errorf("Expcected %d, got %d", db.KeyType.Date, enum)
	}
}

func TestEnumKeyType2(t *testing.T) {
	var enum db.KeyTypeEnum
	var err error
	var result []byte

	enum = db.KeyType.Array
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"array"` != string(result) {
		t.Errorf("Expected '\"array\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"array"`), &enum)
	if db.KeyType.Array != enum {
		t.Errorf("Expcected %d, got %d", db.KeyType.Array, enum)
	}
}
