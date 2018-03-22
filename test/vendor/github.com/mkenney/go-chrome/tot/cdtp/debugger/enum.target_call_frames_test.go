package debugger

import (
	"encoding/json"
	"testing"
)

func TestEnumTargetCallFrames(t *testing.T) {
	var enum TargetCallFramesEnum
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

	enum = TargetCallFrames.Any
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"any"` != string(result) {
		t.Errorf("Expected '\"any\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"any"`), &enum)
	if TargetCallFrames.Any != enum {
		t.Errorf("Expcected %d, got %d", TargetCallFrames.Any, enum)
	}

	enum = TargetCallFrames.Current
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"current"` != string(result) {
		t.Errorf("Expected '\"current\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"current"`), &enum)
	if TargetCallFrames.Current != enum {
		t.Errorf("Expcected %d, got %d", TargetCallFrames.Current, enum)
	}
}
