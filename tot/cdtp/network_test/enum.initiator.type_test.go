package network_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumInitiatorType(t *testing.T) {
	var enum network.InitiatorTypeEnum
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

	enum = network.InitiatorType.Parser
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"parser"` != string(result) {
		t.Errorf("Expected '\"parser\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"parser"`), &enum)
	if network.InitiatorType.Parser != enum {
		t.Errorf("Expcected %d, got %d", network.InitiatorType.Parser, enum)
	}

	enum = network.InitiatorType.Script
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"script"` != string(result) {
		t.Errorf("Expected '\"script\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"script"`), &enum)
	if network.InitiatorType.Script != enum {
		t.Errorf("Expcected %d, got %d", network.InitiatorType.Script, enum)
	}

	enum = network.InitiatorType.Preload
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"preload"` != string(result) {
		t.Errorf("Expected '\"preload\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"preload"`), &enum)
	if network.InitiatorType.Preload != enum {
		t.Errorf("Expcected %d, got %d", network.InitiatorType.Preload, enum)
	}
}

func TestEnumInitiatorType2(t *testing.T) {
	var enum network.InitiatorTypeEnum
	var err error
	var result []byte

	enum = network.InitiatorType.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if network.InitiatorType.Other != enum {
		t.Errorf("Expcected %d, got %d", network.InitiatorType.Other, enum)
	}
}
