package css

import (
	"encoding/json"
	"testing"
)

func TestEnumForcedPseudoClasses(t *testing.T) {
	var enum ForcedPseudoClassesEnum
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

	enum = ForcedPseudoClasses.Active
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"active"` != string(result) {
		t.Errorf("Expected '\"active\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"active"`), &enum)
	if ForcedPseudoClasses.Active != enum {
		t.Errorf("Expcected %d, got %d", ForcedPseudoClasses.Active, enum)
	}

	enum = ForcedPseudoClasses.Focus
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"focus"` != string(result) {
		t.Errorf("Expected '\"focus\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"focus"`), &enum)
	if ForcedPseudoClasses.Focus != enum {
		t.Errorf("Expcected %d, got %d", ForcedPseudoClasses.Focus, enum)
	}

	enum = ForcedPseudoClasses.Hover
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"hover"` != string(result) {
		t.Errorf("Expected '\"hover\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"hover"`), &enum)
	if ForcedPseudoClasses.Hover != enum {
		t.Errorf("Expcected %d, got %d", ForcedPseudoClasses.Hover, enum)
	}

	enum = ForcedPseudoClasses.Visited
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"visited"` != string(result) {
		t.Errorf("Expected '\"visited\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"visited"`), &enum)
	if ForcedPseudoClasses.Visited != enum {
		t.Errorf("Expcected %d, got %d", ForcedPseudoClasses.Visited, enum)
	}
}
