package css_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/css"
)

func TestEnumSource(t *testing.T) {
	var enum css.SourceEnum
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

	enum = css.Source.MediaRule
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mediaRule"` != string(result) {
		t.Errorf("Expected '\"mediaRule\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mediaRule"`), &enum)
	if css.Source.MediaRule != enum {
		t.Errorf("Expcected %d, got %d", css.Source.MediaRule, enum)
	}

	enum = css.Source.ImportRule
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"importRule"` != string(result) {
		t.Errorf("Expected '\"importRule\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"importRule"`), &enum)
	if css.Source.ImportRule != enum {
		t.Errorf("Expcected %d, got %d", css.Source.ImportRule, enum)
	}

	enum = css.Source.LinkedSheet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"linkedSheet"` != string(result) {
		t.Errorf("Expected '\"linkedSheet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"linkedSheet"`), &enum)
	if css.Source.LinkedSheet != enum {
		t.Errorf("Expcected %d, got %d", css.Source.LinkedSheet, enum)
	}
}

func TestEnumSource2(t *testing.T) {
	var enum css.SourceEnum
	var err error
	var result []byte

	enum = css.Source.InlineSheet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"inlineSheet"` != string(result) {
		t.Errorf("Expected '\"inlineSheet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"inlineSheet"`), &enum)
	if css.Source.InlineSheet != enum {
		t.Errorf("Expcected %d, got %d", css.Source.InlineSheet, enum)
	}
}
