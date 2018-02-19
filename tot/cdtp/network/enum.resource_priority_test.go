package network

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumResourcePriority(t *testing.T) {
	var enum network.ResourcePriorityEnum
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

	enum = network.ResourcePriority.VeryLow
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"VeryLow"` != string(result) {
		t.Errorf("Expected '\"VeryLow\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"VeryLow"`), &enum)
	if network.ResourcePriority.VeryLow != enum {
		t.Errorf("Expcected %d, got %d", network.ResourcePriority.VeryLow, enum)
	}

	enum = network.ResourcePriority.Low
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Low"` != string(result) {
		t.Errorf("Expected '\"Low\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Low"`), &enum)
	if network.ResourcePriority.Low != enum {
		t.Errorf("Expcected %d, got %d", network.ResourcePriority.Low, enum)
	}

	enum = network.ResourcePriority.Medium
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Medium"` != string(result) {
		t.Errorf("Expected '\"Medium\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Medium"`), &enum)
	if network.ResourcePriority.Medium != enum {
		t.Errorf("Expcected %d, got %d", network.ResourcePriority.Medium, enum)
	}
}

func TestEnumResourcePriority2(t *testing.T) {
	var enum network.ResourcePriorityEnum
	var err error
	var result []byte

	enum = network.ResourcePriority.High
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"High"` != string(result) {
		t.Errorf("Expected '\"High\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"High"`), &enum)
	if network.ResourcePriority.High != enum {
		t.Errorf("Expcected %d, got %d", network.ResourcePriority.High, enum)
	}

	enum = network.ResourcePriority.VeryHigh
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"VeryHigh"` != string(result) {
		t.Errorf("Expected '\"VeryHigh\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"VeryHigh"`), &enum)
	if network.ResourcePriority.VeryHigh != enum {
		t.Errorf("Expcected %d, got %d", network.ResourcePriority.VeryHigh, enum)
	}
}
