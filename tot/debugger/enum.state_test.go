package debugger

import (
	"encoding/json"
	"testing"
)

func TestEnumState(t *testing.T) {
	var enum StateEnum
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

	enum = State.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if State.None != enum {
		t.Errorf("Expcected %d, got %d", State.None, enum)
	}

	enum = State.Uncaught
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"uncaught"` != string(result) {
		t.Errorf("Expected '\"uncaught\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"uncaught"`), &enum)
	if State.Uncaught != enum {
		t.Errorf("Expcected %d, got %d", State.Uncaught, enum)
	}

	enum = State.All
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"all"` != string(result) {
		t.Errorf("Expected '\"all\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"all"`), &enum)
	if State.All != enum {
		t.Errorf("Expcected %d, got %d", State.All, enum)
	}
}
