package network

import (
	"encoding/json"
	"testing"
)

func TestEnumInitiatorType(t *testing.T) {
	var enum InitiatorTypeEnum
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

	enum = InitiatorType.Parser
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"parser"` != string(result) {
		t.Errorf("Expected '\"parser\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"parser"`), &enum)
	if InitiatorType.Parser != enum {
		t.Errorf("Expcected %d, got %d", InitiatorType.Parser, enum)
	}

	enum = InitiatorType.Script
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"script"` != string(result) {
		t.Errorf("Expected '\"script\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"script"`), &enum)
	if InitiatorType.Script != enum {
		t.Errorf("Expcected %d, got %d", InitiatorType.Script, enum)
	}

	enum = InitiatorType.Preload
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"preload"` != string(result) {
		t.Errorf("Expected '\"preload\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"preload"`), &enum)
	if InitiatorType.Preload != enum {
		t.Errorf("Expcected %d, got %d", InitiatorType.Preload, enum)
	}

	enum = InitiatorType.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if InitiatorType.Other != enum {
		t.Errorf("Expcected %d, got %d", InitiatorType.Other, enum)
	}
}
