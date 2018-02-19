package input

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/input"
)

func TestEnumButtonEvent(t *testing.T) {
	var enum input.ButtonEventEnum
	var err error
	var result []byte

	err = json.Unmarshal([]byte(`""`), &enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}

	err = json.Unmarshal([]byte(`"invalid value"`), &enum)
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

	enum = input.ButtonEvent.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if input.ButtonEvent.None != enum {
		t.Errorf("Expcected %d, got %d", input.ButtonEvent.None, enum)
	}

	enum = input.ButtonEvent.Left
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"left"` != string(result) {
		t.Errorf("Expected '\"left\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"left"`), &enum)
	if input.ButtonEvent.Left != enum {
		t.Errorf("Expcected %d, got %d", input.ButtonEvent.Left, enum)
	}

	enum = input.ButtonEvent.Middle
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"middle"` != string(result) {
		t.Errorf("Expected '\"middle\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"middle"`), &enum)
	if input.ButtonEvent.Middle != enum {
		t.Errorf("Expcected %d, got %d", input.ButtonEvent.Middle, enum)
	}
}

func TestEnumButtonEvent2(t *testing.T) {
	var enum input.ButtonEventEnum
	var err error
	var result []byte

	enum = input.ButtonEvent.Right
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"right"` != string(result) {
		t.Errorf("Expected '\"right\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"right"`), &enum)
	if input.ButtonEvent.Right != enum {
		t.Errorf("Expcected %d, got %d", input.ButtonEvent.Right, enum)
	}
}
