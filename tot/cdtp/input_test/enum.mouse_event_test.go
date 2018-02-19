package input_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/input"
)

func TestEnumMouseEvent(t *testing.T) {
	var enum input.MouseEventEnum
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

	enum = input.MouseEvent.MousePressed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mousePressed"` != string(result) {
		t.Errorf("Expected '\"mousePressed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mousePressed"`), &enum)
	if input.MouseEvent.MousePressed != enum {
		t.Errorf("Expcected %d, got %d", input.MouseEvent.MousePressed, enum)
	}

	enum = input.MouseEvent.MouseReleased
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mouseReleased"` != string(result) {
		t.Errorf("Expected '\"mouseReleased\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mouseReleased"`), &enum)
	if input.MouseEvent.MouseReleased != enum {
		t.Errorf("Expcected %d, got %d", input.MouseEvent.MouseReleased, enum)
	}

	enum = input.MouseEvent.MouseMoved
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mouseMoved"` != string(result) {
		t.Errorf("Expected '\"mouseMoved\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mouseMoved"`), &enum)
	if input.MouseEvent.MouseMoved != enum {
		t.Errorf("Expcected %d, got %d", input.MouseEvent.MouseMoved, enum)
	}
}

func TestEnumMouseEvent2(t *testing.T) {
	var enum input.MouseEventEnum
	var err error
	var result []byte

	enum = input.MouseEvent.MouseWheel
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mouseWheel"` != string(result) {
		t.Errorf("Expected '\"mouseWheel\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mouseWheel"`), &enum)
	if input.MouseEvent.MouseWheel != enum {
		t.Errorf("Expcected %d, got %d", input.MouseEvent.MouseWheel, enum)
	}
}
