package console_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/console"
)

func TestEnumMessageLevel(t *testing.T) {
	var enum console.MessageLevelEnum
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

	enum = console.MessageLevel.Log
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"log"` != string(result) {
		t.Errorf("Expected '\"log\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"log"`), &enum)
	if console.MessageLevel.Log != enum {
		t.Errorf("Expcected %d, got %d", console.MessageLevel.Log, enum)
	}

	enum = console.MessageLevel.Warning
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"warning"` != string(result) {
		t.Errorf("Expected '\"jpeg\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"warning"`), &enum)
	if console.MessageLevel.Warning != enum {
		t.Errorf("Expcected %d, got %d", console.MessageLevel.Warning, enum)
	}

	enum = console.MessageLevel.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if console.MessageLevel.Error != enum {
		t.Errorf("Expcected %d, got %d", console.MessageLevel.Error, enum)
	}
}

func TestEnumMessageLevel2(t *testing.T) {
	var enum console.MessageLevelEnum
	var err error
	var result []byte

	enum = console.MessageLevel.Debug
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"debug"` != string(result) {
		t.Errorf("Expected '\"debug\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"debug"`), &enum)
	if console.MessageLevel.Debug != enum {
		t.Errorf("Expcected %d, got %d", console.MessageLevel.Debug, enum)
	}

	enum = console.MessageLevel.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if console.MessageLevel.Info != enum {
		t.Errorf("Expcected %d, got %d", console.MessageLevel.Info, enum)
	}
}
