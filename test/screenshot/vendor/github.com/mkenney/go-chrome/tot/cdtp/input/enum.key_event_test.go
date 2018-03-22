package input

import (
	"encoding/json"
	"testing"
)

func TestEnumKeyEvent(t *testing.T) {
	var enum KeyEventEnum
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

	enum = KeyEvent.KeyDown
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"keyDown"` != string(result) {
		t.Errorf("Expected '\"keyDown\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"keyDown"`), &enum)
	if KeyEvent.KeyDown != enum {
		t.Errorf("Expcected %d, got %d", KeyEvent.KeyDown, enum)
	}

	enum = KeyEvent.KeyUp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"keyUp"` != string(result) {
		t.Errorf("Expected '\"keyUp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"keyUp"`), &enum)
	if KeyEvent.KeyUp != enum {
		t.Errorf("Expcected %d, got %d", KeyEvent.KeyUp, enum)
	}

	enum = KeyEvent.RawKeyDown
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"rawKeyDown"` != string(result) {
		t.Errorf("Expected '\"rawKeyDown\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"rawKeyDown"`), &enum)
	if KeyEvent.RawKeyDown != enum {
		t.Errorf("Expcected %d, got %d", KeyEvent.RawKeyDown, enum)
	}
}

func TestEnumKeyEvent2(t *testing.T) {
	var enum KeyEventEnum
	var err error
	var result []byte

	enum = KeyEvent.Char
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"char"` != string(result) {
		t.Errorf("Expected '\"char\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"char"`), &enum)
	if KeyEvent.Char != enum {
		t.Errorf("Expcected %d, got %d", KeyEvent.Char, enum)
	}
}
