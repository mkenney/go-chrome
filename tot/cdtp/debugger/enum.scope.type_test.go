package debugger

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/debugger"
)

func TestEnumScopeType(t *testing.T) {
	var enum debugger.ScopeTypeEnum
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

	enum = debugger.ScopeType.Global
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"global"` != string(result) {
		t.Errorf("Expected '\"global\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"global"`), &enum)
	if debugger.ScopeType.Global != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Global, enum)
	}

	enum = debugger.ScopeType.Local
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"local"` != string(result) {
		t.Errorf("Expected '\"local\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"local"`), &enum)
	if debugger.ScopeType.Local != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Local, enum)
	}

	enum = debugger.ScopeType.With
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"with"` != string(result) {
		t.Errorf("Expected '\"with\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"with"`), &enum)
	if debugger.ScopeType.With != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.With, enum)
	}
}
func TestEnumScopeType2(t *testing.T) {
	var enum debugger.ScopeTypeEnum
	var err error
	var result []byte

	enum = debugger.ScopeType.Closure
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"closure"` != string(result) {
		t.Errorf("Expected '\"debuggerstatementclosure got '%s'", result)
	}
	json.Unmarshal([]byte(`"debuggerstatementclosure"`), &enum)
	if debugger.ScopeType.Closure != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Closure, enum)
	}

	enum = debugger.ScopeType.Catch
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"catch"` != string(result) {
		t.Errorf("Expected '\"catch\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"catch"`), &enum)
	if debugger.ScopeType.Catch != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Catch, enum)
	}

	enum = debugger.ScopeType.Block
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"block"` != string(result) {
		t.Errorf("Expected '\"block\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"block"`), &enum)
	if debugger.ScopeType.Block != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Block, enum)
	}

	enum = debugger.ScopeType.Script
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"script"` != string(result) {
		t.Errorf("Expected '\"script\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"script"`), &enum)
	if debugger.ScopeType.Script != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Script, enum)
	}
}

func TestEnumScopeType3(t *testing.T) {
	var enum debugger.ScopeTypeEnum
	var err error
	var result []byte

	enum = debugger.ScopeType.Eval
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"eval"` != string(result) {
		t.Errorf("Expected '\"eval\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"eval"`), &enum)
	if debugger.ScopeType.Eval != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Eval, enum)
	}

	enum = debugger.ScopeType.Module
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"module"` != string(result) {
		t.Errorf("Expected '\"module\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"module"`), &enum)
	if debugger.ScopeType.Module != enum {
		t.Errorf("Expcected %d, got %d", debugger.ScopeType.Module, enum)
	}
}
