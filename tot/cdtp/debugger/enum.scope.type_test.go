package debugger

import (
	"encoding/json"
	"testing"
)

func TestEnumScopeType(t *testing.T) {
	var enum ScopeTypeEnum
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

	enum = ScopeType.Global
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"global"` != string(result) {
		t.Errorf("Expected '\"global\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"global"`), &enum)
	if ScopeType.Global != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Global, enum)
	}

	enum = ScopeType.Local
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"local"` != string(result) {
		t.Errorf("Expected '\"local\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"local"`), &enum)
	if ScopeType.Local != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Local, enum)
	}

	enum = ScopeType.With
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"with"` != string(result) {
		t.Errorf("Expected '\"with\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"with"`), &enum)
	if ScopeType.With != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.With, enum)
	}
}
func TestEnumScopeType2(t *testing.T) {
	var enum ScopeTypeEnum
	var err error
	var result []byte

	enum = ScopeType.Closure
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"closure"` != string(result) {
		t.Errorf("Expected '\"debuggerstatementclosure got '%s'", result)
	}
	json.Unmarshal([]byte(`"debuggerstatementclosure"`), &enum)
	if ScopeType.Closure != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Closure, enum)
	}

	enum = ScopeType.Catch
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"catch"` != string(result) {
		t.Errorf("Expected '\"catch\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"catch"`), &enum)
	if ScopeType.Catch != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Catch, enum)
	}

	enum = ScopeType.Block
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"block"` != string(result) {
		t.Errorf("Expected '\"block\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"block"`), &enum)
	if ScopeType.Block != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Block, enum)
	}

	enum = ScopeType.Script
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"script"` != string(result) {
		t.Errorf("Expected '\"script\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"script"`), &enum)
	if ScopeType.Script != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Script, enum)
	}
}

func TestEnumScopeType3(t *testing.T) {
	var enum ScopeTypeEnum
	var err error
	var result []byte

	enum = ScopeType.Eval
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"eval"` != string(result) {
		t.Errorf("Expected '\"eval\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"eval"`), &enum)
	if ScopeType.Eval != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Eval, enum)
	}

	enum = ScopeType.Module
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"module"` != string(result) {
		t.Errorf("Expected '\"module\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"module"`), &enum)
	if ScopeType.Module != enum {
		t.Errorf("Expcected %d, got %d", ScopeType.Module, enum)
	}
}
