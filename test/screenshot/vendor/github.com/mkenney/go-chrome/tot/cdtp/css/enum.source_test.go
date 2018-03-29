package css

import (
	"encoding/json"
	"testing"
)

func TestEnumSource(t *testing.T) {
	var enum SourceEnum
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

	enum = Source.MediaRule
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mediaRule"` != string(result) {
		t.Errorf("Expected '\"mediaRule\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mediaRule"`), &enum)
	if Source.MediaRule != enum {
		t.Errorf("Expcected %d, got %d", Source.MediaRule, enum)
	}

	enum = Source.ImportRule
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"importRule"` != string(result) {
		t.Errorf("Expected '\"importRule\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"importRule"`), &enum)
	if Source.ImportRule != enum {
		t.Errorf("Expcected %d, got %d", Source.ImportRule, enum)
	}

	enum = Source.LinkedSheet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"linkedSheet"` != string(result) {
		t.Errorf("Expected '\"linkedSheet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"linkedSheet"`), &enum)
	if Source.LinkedSheet != enum {
		t.Errorf("Expcected %d, got %d", Source.LinkedSheet, enum)
	}
}

func TestEnumSource2(t *testing.T) {
	var enum SourceEnum
	var err error
	var result []byte

	enum = Source.InlineSheet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"inlineSheet"` != string(result) {
		t.Errorf("Expected '\"inlineSheet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"inlineSheet"`), &enum)
	if Source.InlineSheet != enum {
		t.Errorf("Expcected %d, got %d", Source.InlineSheet, enum)
	}
}
