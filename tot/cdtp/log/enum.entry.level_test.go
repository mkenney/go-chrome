package log

import (
	"encoding/json"
	"testing"
)

func TestEnumLevel(t *testing.T) {
	var enum LevelEnum
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

	enum = Level.Verbose
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"verbose"` != string(result) {
		t.Errorf("Expected '\"verbose\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"verbose"`), &enum)
	if Level.Verbose != enum {
		t.Errorf("Expcected %d, got %d", Level.Verbose, enum)
	}

	enum = Level.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if Level.Info != enum {
		t.Errorf("Expcected %d, got %d", Level.Info, enum)
	}

	enum = Level.Warning
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"warning"` != string(result) {
		t.Errorf("Expected '\"warning\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"warning"`), &enum)
	if Level.Warning != enum {
		t.Errorf("Expcected %d, got %d", Level.Warning, enum)
	}

	enum = Level.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if Level.Error != enum {
		t.Errorf("Expcected %d, got %d", Level.Error, enum)
	}
}
