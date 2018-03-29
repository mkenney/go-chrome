package log

import (
	"encoding/json"
	"testing"
)

func TestEnumName(t *testing.T) {
	var enum NameEnum
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

	enum = Name.LongTask
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"longTask"` != string(result) {
		t.Errorf("Expected '\"longTask\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"longTask"`), &enum)
	if Name.LongTask != enum {
		t.Errorf("Expcected %d, got %d", Name.LongTask, enum)
	}

	enum = Name.LongLayout
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"longLayout"` != string(result) {
		t.Errorf("Expected '\"longLayout\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"longLayout"`), &enum)
	if Name.LongLayout != enum {
		t.Errorf("Expcected %d, got %d", Name.LongLayout, enum)
	}

	enum = Name.BlockedEvent
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"blockedEvent"` != string(result) {
		t.Errorf("Expected '\"blockedEvent\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"blockedEvent"`), &enum)
	if Name.BlockedEvent != enum {
		t.Errorf("Expcected %d, got %d", Name.BlockedEvent, enum)
	}
}

func TestEnumName2(t *testing.T) {
	var enum NameEnum
	var err error
	var result []byte

	enum = Name.BlockedParser
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"blockedParser"` != string(result) {
		t.Errorf("Expected '\"blockedParser\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"blockedParser"`), &enum)
	if Name.BlockedParser != enum {
		t.Errorf("Expcected %d, got %d", Name.BlockedParser, enum)
	}

	enum = Name.DiscouragedAPIUse
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"discouragedAPIUse"` != string(result) {
		t.Errorf("Expected '\"discouragedAPIUse\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"discouragedAPIUse"`), &enum)
	if Name.DiscouragedAPIUse != enum {
		t.Errorf("Expcected %d, got %d", Name.DiscouragedAPIUse, enum)
	}

	enum = Name.Handler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"handler"` != string(result) {
		t.Errorf("Expected '\"handler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"handler"`), &enum)
	if Name.Handler != enum {
		t.Errorf("Expcected %d, got %d", Name.Handler, enum)
	}

	enum = Name.RecurringHandler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"recurringHandler"` != string(result) {
		t.Errorf("Expected '\"recurringHandler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"recurringHandler"`), &enum)
	if Name.RecurringHandler != enum {
		t.Errorf("Expcected %d, got %d", Name.RecurringHandler, enum)
	}
}
