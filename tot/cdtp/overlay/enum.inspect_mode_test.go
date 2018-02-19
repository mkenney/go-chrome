package overlay

import (
	"encoding/json"
	"testing"
)

func TestEnumInspectMode(t *testing.T) {
	var enum InspectModeEnum
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

	enum = InspectMode.SearchForNode
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"searchForNode"` != string(result) {
		t.Errorf("Expected '\"searchForNode\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"searchForNode"`), &enum)
	if InspectMode.SearchForNode != enum {
		t.Errorf("Expcected %d, got %d", InspectMode.SearchForNode, enum)
	}

	enum = InspectMode.SearchForUAShadowDOM
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"searchForUAShadowDOM"` != string(result) {
		t.Errorf("Expected '\"searchForUAShadowDOM\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"searchForUAShadowDOM"`), &enum)
	if InspectMode.SearchForUAShadowDOM != enum {
		t.Errorf("Expcected %d, got %d", InspectMode.SearchForUAShadowDOM, enum)
	}

	enum = InspectMode.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if InspectMode.None != enum {
		t.Errorf("Expcected %d, got %d", InspectMode.None, enum)
	}
}
