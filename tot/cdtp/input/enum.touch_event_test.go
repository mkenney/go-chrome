package input

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/input"
)

func TestEnumTouchEvent(t *testing.T) {
	var enum input.TouchEventEnum
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

	enum = input.TouchEvent.TouchStart
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchStart"` != string(result) {
		t.Errorf("Expected '\"touchStart\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchStart"`), &enum)
	if input.TouchEvent.TouchStart != enum {
		t.Errorf("Expcected %d, got %d", input.TouchEvent.TouchStart, enum)
	}

	enum = input.TouchEvent.TouchEnd
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchEnd"` != string(result) {
		t.Errorf("Expected '\"touchEnd\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchEnd"`), &enum)
	if input.TouchEvent.TouchEnd != enum {
		t.Errorf("Expcected %d, got %d", input.TouchEvent.TouchEnd, enum)
	}

	enum = input.TouchEvent.TouchMove
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchMove"` != string(result) {
		t.Errorf("Expected '\"touchMove\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchMove"`), &enum)
	if input.TouchEvent.TouchMove != enum {
		t.Errorf("Expcected %d, got %d", input.TouchEvent.TouchMove, enum)
	}
}

func TestEnumTouchEvent2(t *testing.T) {
	var enum input.TouchEventEnum
	var err error
	var result []byte

	enum = input.TouchEvent.TouchCancel
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchCancel"` != string(result) {
		t.Errorf("Expected '\"touchCancel\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchCancel"`), &enum)
	if input.TouchEvent.TouchCancel != enum {
		t.Errorf("Expcected %d, got %d", input.TouchEvent.TouchCancel, enum)
	}
}
