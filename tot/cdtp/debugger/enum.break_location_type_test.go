package debugger

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/debugger"
)

func TestEnumBreakLocationType(t *testing.T) {
	var enum debugger.BreakLocationTypeEnum
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

	enum = debugger.BreakLocationType.DebuggerStatement
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"debuggerStatement"` != string(result) {
		t.Errorf("Expected '\"debuggerStatement\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"debuggerStatement"`), &enum)
	if debugger.BreakLocationType.DebuggerStatement != enum {
		t.Errorf("Expcected %d, got %d", debugger.BreakLocationType.DebuggerStatement, enum)
	}

	enum = debugger.BreakLocationType.Call
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"call"` != string(result) {
		t.Errorf("Expected '\"call\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"call"`), &enum)
	if debugger.BreakLocationType.Call != enum {
		t.Errorf("Expcected %d, got %d", debugger.BreakLocationType.Call, enum)
	}

	enum = debugger.BreakLocationType.Return
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"return"` != string(result) {
		t.Errorf("Expected '\"return\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"return"`), &enum)
	if debugger.BreakLocationType.Return != enum {
		t.Errorf("Expcected %d, got %d", debugger.BreakLocationType.Return, enum)
	}
}
