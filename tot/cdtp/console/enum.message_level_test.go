package console

import (
	"encoding/json"
	"testing"
)

func TestEnumMessageLevel(t *testing.T) {
	var enum MessageLevelEnum
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

	enum = MessageLevel.Log
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"log"` != string(result) {
		t.Errorf("Expected '\"log\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"log"`), &enum)
	if MessageLevel.Log != enum {
		t.Errorf("Expcected %d, got %d", MessageLevel.Log, enum)
	}

	enum = MessageLevel.Warning
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"warning"` != string(result) {
		t.Errorf("Expected '\"jpeg\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"warning"`), &enum)
	if MessageLevel.Warning != enum {
		t.Errorf("Expcected %d, got %d", MessageLevel.Warning, enum)
	}

	enum = MessageLevel.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if MessageLevel.Error != enum {
		t.Errorf("Expcected %d, got %d", MessageLevel.Error, enum)
	}

	enum = MessageLevel.Debug
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"debug"` != string(result) {
		t.Errorf("Expected '\"debug\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"debug"`), &enum)
	if MessageLevel.Debug != enum {
		t.Errorf("Expcected %d, got %d", MessageLevel.Debug, enum)
	}

	enum = MessageLevel.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if MessageLevel.Info != enum {
		t.Errorf("Expcected %d, got %d", MessageLevel.Info, enum)
	}
}
