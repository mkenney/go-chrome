package tree_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/layer/tree"
)

func TestEnumRectType(t *testing.T) {
	var enum tree.RectTypeEnum
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

	enum = tree.RectType.RepaintsOnScroll
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"RepaintsOnScroll"` != string(result) {
		t.Errorf("Expected '\"RepaintsOnScroll\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"RepaintsOnScroll"`), &enum)
	if tree.RectType.RepaintsOnScroll != enum {
		t.Errorf("Expcected %d, got %d", tree.RectType.RepaintsOnScroll, enum)
	}

	enum = tree.RectType.TouchEventHandler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"TouchEventHandler"` != string(result) {
		t.Errorf("Expected '\"TouchEventHandler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"TouchEventHandler"`), &enum)
	if tree.RectType.TouchEventHandler != enum {
		t.Errorf("Expcected %d, got %d", tree.RectType.TouchEventHandler, enum)
	}

	enum = tree.RectType.WheelEventHandler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"WheelEventHandler"` != string(result) {
		t.Errorf("Expected '\"WheelEventHandler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"WheelEventHandler"`), &enum)
	if tree.RectType.WheelEventHandler != enum {
		t.Errorf("Expcected %d, got %d", tree.RectType.WheelEventHandler, enum)
	}
}
