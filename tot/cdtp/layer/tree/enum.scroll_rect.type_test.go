package tree

import (
	"encoding/json"
	"testing"
)

func TestEnumRectType(t *testing.T) {
	var enum RectTypeEnum
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

	enum = RectType.RepaintsOnScroll
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"RepaintsOnScroll"` != string(result) {
		t.Errorf("Expected '\"RepaintsOnScroll\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"RepaintsOnScroll"`), &enum)
	if RectType.RepaintsOnScroll != enum {
		t.Errorf("Expcected %d, got %d", RectType.RepaintsOnScroll, enum)
	}

	enum = RectType.TouchEventHandler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"TouchEventHandler"` != string(result) {
		t.Errorf("Expected '\"TouchEventHandler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"TouchEventHandler"`), &enum)
	if RectType.TouchEventHandler != enum {
		t.Errorf("Expcected %d, got %d", RectType.TouchEventHandler, enum)
	}

	enum = RectType.WheelEventHandler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"WheelEventHandler"` != string(result) {
		t.Errorf("Expected '\"WheelEventHandler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"WheelEventHandler"`), &enum)
	if RectType.WheelEventHandler != enum {
		t.Errorf("Expcected %d, got %d", RectType.WheelEventHandler, enum)
	}
}
