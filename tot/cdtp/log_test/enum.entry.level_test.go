package log_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/log"
)

func TestEnumLevel(t *testing.T) {
	var enum log.LevelEnum
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

	enum = log.Level.Verbose
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"verbose"` != string(result) {
		t.Errorf("Expected '\"verbose\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"verbose"`), &enum)
	if log.Level.Verbose != enum {
		t.Errorf("Expcected %d, got %d", log.Level.Verbose, enum)
	}

	enum = log.Level.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if log.Level.Info != enum {
		t.Errorf("Expcected %d, got %d", log.Level.Info, enum)
	}

	enum = log.Level.Warning
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"warning"` != string(result) {
		t.Errorf("Expected '\"warning\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"warning"`), &enum)
	if log.Level.Warning != enum {
		t.Errorf("Expcected %d, got %d", log.Level.Warning, enum)
	}
}

func TestEnumLevel2(t *testing.T) {
	var enum log.LevelEnum
	var err error
	var result []byte

	enum = log.Level.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if log.Level.Error != enum {
		t.Errorf("Expcected %d, got %d", log.Level.Error, enum)
	}
}
