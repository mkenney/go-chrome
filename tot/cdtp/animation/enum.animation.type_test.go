package animation

import (
	"encoding/json"
	"testing"
)

func TestEnumAnimationType(t *testing.T) {
	var enum AnimationTypeEnum
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

	enum = AnimationType.CSSAnimation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CSSAnimation"` != string(result) {
		t.Errorf("Expected '\"CSSAnimation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CSSAnimation"`), &enum)
	if AnimationType.CSSAnimation != enum {
		t.Errorf("Expcected %d, got %d", AnimationType.CSSAnimation, enum)
	}

	enum = AnimationType.CSSTransition
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CSSTransition"` != string(result) {
		t.Errorf("Expected '\"CSSTransition\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CSSTransition"`), &enum)
	if AnimationType.CSSTransition != enum {
		t.Errorf("Expcected %d, got %d", AnimationType.CSSTransition, enum)
	}

	enum = AnimationType.WebAnimation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"WebAnimation"` != string(result) {
		t.Errorf("Expected '\"WebAnimation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"WebAnimation"`), &enum)
	if AnimationType.WebAnimation != enum {
		t.Errorf("Expcected %d, got %d", AnimationType.WebAnimation, enum)
	}
}
