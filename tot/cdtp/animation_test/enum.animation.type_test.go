package animation_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/animation"
)

func TestEnumType(t *testing.T) {
	var enum animation.TypeEnum
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

	enum = animation.Type.CSSAnimation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CSSAnimation"` != string(result) {
		t.Errorf("Expected '\"CSSAnimation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CSSAnimation"`), &enum)
	if animation.Type.CSSAnimation != enum {
		t.Errorf("Expcected %d, got %d", animation.Type.CSSAnimation, enum)
	}

	enum = animation.Type.CSSTransition
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CSSTransition"` != string(result) {
		t.Errorf("Expected '\"CSSTransition\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CSSTransition"`), &enum)
	if animation.Type.CSSTransition != enum {
		t.Errorf("Expcected %d, got %d", animation.Type.CSSTransition, enum)
	}

	enum = animation.Type.WebAnimation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"WebAnimation"` != string(result) {
		t.Errorf("Expected '\"WebAnimation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"WebAnimation"`), &enum)
	if animation.Type.WebAnimation != enum {
		t.Errorf("Expcected %d, got %d", animation.Type.WebAnimation, enum)
	}
}
