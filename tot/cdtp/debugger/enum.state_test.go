package debugger

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/debugger"
)

func TestEnumState(t *testing.T) {
	var enum debugger.StateEnum
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

	enum = debugger.State.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if debugger.State.None != enum {
		t.Errorf("Expcected %d, got %d", debugger.State.None, enum)
	}

	enum = debugger.State.Uncaught
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"uncaught"` != string(result) {
		t.Errorf("Expected '\"uncaught\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"uncaught"`), &enum)
	if debugger.State.Uncaught != enum {
		t.Errorf("Expcected %d, got %d", debugger.State.Uncaught, enum)
	}

	enum = debugger.State.All
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"all"` != string(result) {
		t.Errorf("Expected '\"all\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"all"`), &enum)
	if debugger.State.All != enum {
		t.Errorf("Expcected %d, got %d", debugger.State.All, enum)
	}
}
