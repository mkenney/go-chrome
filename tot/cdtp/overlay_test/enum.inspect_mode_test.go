package overlay_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/overlay"
)

func TestEnumInspectMode(t *testing.T) {
	var enum overlay.InspectModeEnum
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

	enum = overlay.InspectMode.SearchForNode
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"searchForNode"` != string(result) {
		t.Errorf("Expected '\"searchForNode\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"searchForNode"`), &enum)
	if overlay.InspectMode.SearchForNode != enum {
		t.Errorf("Expcected %d, got %d", overlay.InspectMode.SearchForNode, enum)
	}

	enum = overlay.InspectMode.SearchForUAShadowDOM
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"searchForUAShadowDOM"` != string(result) {
		t.Errorf("Expected '\"searchForUAShadowDOM\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"searchForUAShadowDOM"`), &enum)
	if overlay.InspectMode.SearchForUAShadowDOM != enum {
		t.Errorf("Expcected %d, got %d", overlay.InspectMode.SearchForUAShadowDOM, enum)
	}

	enum = overlay.InspectMode.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if overlay.InspectMode.None != enum {
		t.Errorf("Expcected %d, got %d", overlay.InspectMode.None, enum)
	}
}
