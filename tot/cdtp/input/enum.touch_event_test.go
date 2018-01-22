package input

import (
	"encoding/json"
	"testing"
)

func TestEnumTouchEvent(t *testing.T) {
	var enum TouchEventEnum
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

	enum = TouchEvent.TouchStart
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchStart"` != string(result) {
		t.Errorf("Expected '\"touchStart\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchStart"`), &enum)
	if TouchEvent.TouchStart != enum {
		t.Errorf("Expcected %d, got %d", TouchEvent.TouchStart, enum)
	}

	enum = TouchEvent.TouchEnd
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchEnd"` != string(result) {
		t.Errorf("Expected '\"touchEnd\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchEnd"`), &enum)
	if TouchEvent.TouchEnd != enum {
		t.Errorf("Expcected %d, got %d", TouchEvent.TouchEnd, enum)
	}

	enum = TouchEvent.TouchMove
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchMove"` != string(result) {
		t.Errorf("Expected '\"touchMove\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchMove"`), &enum)
	if TouchEvent.TouchMove != enum {
		t.Errorf("Expcected %d, got %d", TouchEvent.TouchMove, enum)
	}

	enum = TouchEvent.TouchCancel
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"touchCancel"` != string(result) {
		t.Errorf("Expected '\"touchCancel\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"touchCancel"`), &enum)
	if TouchEvent.TouchCancel != enum {
		t.Errorf("Expcected %d, got %d", TouchEvent.TouchCancel, enum)
	}
}
