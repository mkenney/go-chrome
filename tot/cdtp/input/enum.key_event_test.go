package input

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/input"
)

func TestEnumKeyEvent(t *testing.T) {
	var enum input.KeyEventEnum
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

	enum = input.KeyEvent.KeyDown
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"keyDown"` != string(result) {
		t.Errorf("Expected '\"keyDown\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"keyDown"`), &enum)
	if input.KeyEvent.KeyDown != enum {
		t.Errorf("Expcected %d, got %d", input.KeyEvent.KeyDown, enum)
	}

	enum = input.KeyEvent.KeyUp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"keyUp"` != string(result) {
		t.Errorf("Expected '\"keyUp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"keyUp"`), &enum)
	if input.KeyEvent.KeyUp != enum {
		t.Errorf("Expcected %d, got %d", input.KeyEvent.KeyUp, enum)
	}

	enum = input.KeyEvent.RawKeyDown
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"rawKeyDown"` != string(result) {
		t.Errorf("Expected '\"rawKeyDown\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"rawKeyDown"`), &enum)
	if input.KeyEvent.RawKeyDown != enum {
		t.Errorf("Expcected %d, got %d", input.KeyEvent.RawKeyDown, enum)
	}
}

func TestEnumKeyEvent2(t *testing.T) {
	var enum input.KeyEventEnum
	var err error
	var result []byte

	enum = input.KeyEvent.Char
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"char"` != string(result) {
		t.Errorf("Expected '\"char\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"char"`), &enum)
	if input.KeyEvent.Char != enum {
		t.Errorf("Expcected %d, got %d", input.KeyEvent.Char, enum)
	}
}
