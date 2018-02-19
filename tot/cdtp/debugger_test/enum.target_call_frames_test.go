package debugger_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/debugger"
)

func TestEnumTargetCallFrames(t *testing.T) {
	var enum debugger.TargetCallFramesEnum
	var err error
	var result []byte

	err = json.Unmarshal([]byte(`""`), &enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}

	err = json.Unmarshal([]byte(`"invalid value"`), &enum)
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

	enum = debugger.TargetCallFrames.Any
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"any"` != string(result) {
		t.Errorf("Expected '\"any\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"any"`), &enum)
	if debugger.TargetCallFrames.Any != enum {
		t.Errorf("Expcected %d, got %d", debugger.TargetCallFrames.Any, enum)
	}

	enum = debugger.TargetCallFrames.Current
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"current"` != string(result) {
		t.Errorf("Expected '\"current\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"current"`), &enum)
	if debugger.TargetCallFrames.Current != enum {
		t.Errorf("Expcected %d, got %d", debugger.TargetCallFrames.Current, enum)
	}
}
