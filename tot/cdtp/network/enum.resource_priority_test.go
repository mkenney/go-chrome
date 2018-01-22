package network

import (
	"encoding/json"
	"testing"
)

func TestEnumResourcePriority(t *testing.T) {
	var enum ResourcePriorityEnum
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

	enum = ResourcePriority.VeryLow
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"VeryLow"` != string(result) {
		t.Errorf("Expected '\"VeryLow\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"VeryLow"`), &enum)
	if ResourcePriority.VeryLow != enum {
		t.Errorf("Expcected %d, got %d", ResourcePriority.VeryLow, enum)
	}

	enum = ResourcePriority.Low
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Low"` != string(result) {
		t.Errorf("Expected '\"Low\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Low"`), &enum)
	if ResourcePriority.Low != enum {
		t.Errorf("Expcected %d, got %d", ResourcePriority.Low, enum)
	}

	enum = ResourcePriority.Medium
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Medium"` != string(result) {
		t.Errorf("Expected '\"Medium\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Medium"`), &enum)
	if ResourcePriority.Medium != enum {
		t.Errorf("Expcected %d, got %d", ResourcePriority.Medium, enum)
	}

	enum = ResourcePriority.High
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"High"` != string(result) {
		t.Errorf("Expected '\"High\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"High"`), &enum)
	if ResourcePriority.High != enum {
		t.Errorf("Expcected %d, got %d", ResourcePriority.High, enum)
	}

	enum = ResourcePriority.VeryHigh
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"VeryHigh"` != string(result) {
		t.Errorf("Expected '\"VeryHigh\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"VeryHigh"`), &enum)
	if ResourcePriority.VeryHigh != enum {
		t.Errorf("Expcected %d, got %d", ResourcePriority.VeryHigh, enum)
	}
}
