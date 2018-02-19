package css_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/css"
)

func TestEnumForcedPseudoClasses(t *testing.T) {
	var enum css.ForcedPseudoClassesEnum
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

	enum = css.ForcedPseudoClasses.Active
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"active"` != string(result) {
		t.Errorf("Expected '\"active\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"active"`), &enum)
	if css.ForcedPseudoClasses.Active != enum {
		t.Errorf("Expcected %d, got %d", css.ForcedPseudoClasses.Active, enum)
	}

	enum = css.ForcedPseudoClasses.Focus
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"focus"` != string(result) {
		t.Errorf("Expected '\"focus\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"focus"`), &enum)
	if css.ForcedPseudoClasses.Focus != enum {
		t.Errorf("Expcected %d, got %d", css.ForcedPseudoClasses.Focus, enum)
	}

	enum = css.ForcedPseudoClasses.Hover
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"hover"` != string(result) {
		t.Errorf("Expected '\"hover\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"hover"`), &enum)
	if css.ForcedPseudoClasses.Hover != enum {
		t.Errorf("Expcected %d, got %d", css.ForcedPseudoClasses.Hover, enum)
	}
}

func TestEnumForcedPseudoClasses2(t *testing.T) {
	var enum css.ForcedPseudoClassesEnum
	var err error
	var result []byte

	enum = css.ForcedPseudoClasses.Visited
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"visited"` != string(result) {
		t.Errorf("Expected '\"visited\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"visited"`), &enum)
	if css.ForcedPseudoClasses.Visited != enum {
		t.Errorf("Expcected %d, got %d", css.ForcedPseudoClasses.Visited, enum)
	}
}
