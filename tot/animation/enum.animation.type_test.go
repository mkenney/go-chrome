package animation

import (
	"encoding/json"
	"testing"
)

func TestEnumType(t *testing.T) {
	var enum TypeEnum
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

	enum = Type.CSSAnimation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CSSAnimation"` != string(result) {
		t.Errorf("Expected '\"CSSAnimation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CSSAnimation"`), &enum)
	if Type.CSSAnimation != enum {
		t.Errorf("Expcected %d, got %d", Type.CSSAnimation, enum)
	}

	enum = Type.CSSTransition
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CSSTransition"` != string(result) {
		t.Errorf("Expected '\"CSSTransition\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CSSTransition"`), &enum)
	if Type.CSSTransition != enum {
		t.Errorf("Expcected %d, got %d", Type.CSSTransition, enum)
	}

	enum = Type.WebAnimation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"WebAnimation"` != string(result) {
		t.Errorf("Expected '\"WebAnimation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"WebAnimation"`), &enum)
	if Type.WebAnimation != enum {
		t.Errorf("Expcected %d, got %d", Type.WebAnimation, enum)
	}
}
