package debugger

import (
	"encoding/json"
	"testing"
)

func TestEnumBreakLocationType(t *testing.T) {
	var enum BreakLocationTypeEnum
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

	enum = BreakLocationType.DebuggerStatement
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"debuggerStatement"` != string(result) {
		t.Errorf("Expected '\"debuggerStatement\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"debuggerStatement"`), &enum)
	if BreakLocationType.DebuggerStatement != enum {
		t.Errorf("Expcected %d, got %d", BreakLocationType.DebuggerStatement, enum)
	}

	enum = BreakLocationType.Call
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"call"` != string(result) {
		t.Errorf("Expected '\"call\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"call"`), &enum)
	if BreakLocationType.Call != enum {
		t.Errorf("Expcected %d, got %d", BreakLocationType.Call, enum)
	}

	enum = BreakLocationType.Return
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"return"` != string(result) {
		t.Errorf("Expected '\"return\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"return"`), &enum)
	if BreakLocationType.Return != enum {
		t.Errorf("Expcected %d, got %d", BreakLocationType.Return, enum)
	}
}
