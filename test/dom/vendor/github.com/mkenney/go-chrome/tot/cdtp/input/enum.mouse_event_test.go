package input

import (
	"encoding/json"
	"testing"
)

func TestEnumMouseEvent(t *testing.T) {
	var enum MouseEventEnum
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

	enum = MouseEvent.MousePressed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mousePressed"` != string(result) {
		t.Errorf("Expected '\"mousePressed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mousePressed"`), &enum)
	if MouseEvent.MousePressed != enum {
		t.Errorf("Expcected %d, got %d", MouseEvent.MousePressed, enum)
	}

	enum = MouseEvent.MouseReleased
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mouseReleased"` != string(result) {
		t.Errorf("Expected '\"mouseReleased\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mouseReleased"`), &enum)
	if MouseEvent.MouseReleased != enum {
		t.Errorf("Expcected %d, got %d", MouseEvent.MouseReleased, enum)
	}

	enum = MouseEvent.MouseMoved
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mouseMoved"` != string(result) {
		t.Errorf("Expected '\"mouseMoved\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mouseMoved"`), &enum)
	if MouseEvent.MouseMoved != enum {
		t.Errorf("Expcected %d, got %d", MouseEvent.MouseMoved, enum)
	}
}

func TestEnumMouseEvent2(t *testing.T) {
	var enum MouseEventEnum
	var err error
	var result []byte

	enum = MouseEvent.MouseWheel
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mouseWheel"` != string(result) {
		t.Errorf("Expected '\"mouseWheel\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mouseWheel"`), &enum)
	if MouseEvent.MouseWheel != enum {
		t.Errorf("Expcected %d, got %d", MouseEvent.MouseWheel, enum)
	}
}
