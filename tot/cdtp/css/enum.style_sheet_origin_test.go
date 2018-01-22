package css

import (
	"encoding/json"
	"testing"
)

func TestEnumStyleSheetOrigin(t *testing.T) {
	var enum StyleSheetOriginEnum
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

	enum = StyleSheetOrigin.Injected
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"injected"` != string(result) {
		t.Errorf("Expected '\"injected\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"injected"`), &enum)
	if StyleSheetOrigin.Injected != enum {
		t.Errorf("Expcected %d, got %d", StyleSheetOrigin.Injected, enum)
	}

	enum = StyleSheetOrigin.UserAgent
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"user-agent"` != string(result) {
		t.Errorf("Expected '\"user-agent\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"user-agent"`), &enum)
	if StyleSheetOrigin.UserAgent != enum {
		t.Errorf("Expcected %d, got %d", StyleSheetOrigin.UserAgent, enum)
	}

	enum = StyleSheetOrigin.Inspector
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"inspector"` != string(result) {
		t.Errorf("Expected '\"inspector\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"inspector"`), &enum)
	if StyleSheetOrigin.Inspector != enum {
		t.Errorf("Expcected %d, got %d", StyleSheetOrigin.Inspector, enum)
	}

	enum = StyleSheetOrigin.Log
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"log"` != string(result) {
		t.Errorf("Expected '\"log\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"log"`), &enum)
	if StyleSheetOrigin.Log != enum {
		t.Errorf("Expcected %d, got %d", StyleSheetOrigin.Log, enum)
	}
}
