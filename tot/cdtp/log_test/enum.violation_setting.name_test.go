package log_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/log"
)

func TestEnumName(t *testing.T) {
	var enum log.NameEnum
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

	enum = log.Name.LongTask
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"longTask"` != string(result) {
		t.Errorf("Expected '\"longTask\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"longTask"`), &enum)
	if log.Name.LongTask != enum {
		t.Errorf("Expcected %d, got %d", log.Name.LongTask, enum)
	}

	enum = log.Name.LongLayout
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"longLayout"` != string(result) {
		t.Errorf("Expected '\"longLayout\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"longLayout"`), &enum)
	if log.Name.LongLayout != enum {
		t.Errorf("Expcected %d, got %d", log.Name.LongLayout, enum)
	}

	enum = log.Name.BlockedEvent
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"blockedEvent"` != string(result) {
		t.Errorf("Expected '\"blockedEvent\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"blockedEvent"`), &enum)
	if log.Name.BlockedEvent != enum {
		t.Errorf("Expcected %d, got %d", log.Name.BlockedEvent, enum)
	}
}

func TestEnumName2(t *testing.T) {
	var enum log.NameEnum
	var err error
	var result []byte

	enum = log.Name.BlockedParser
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"blockedParser"` != string(result) {
		t.Errorf("Expected '\"blockedParser\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"blockedParser"`), &enum)
	if log.Name.BlockedParser != enum {
		t.Errorf("Expcected %d, got %d", log.Name.BlockedParser, enum)
	}

	enum = log.Name.DiscouragedAPIUse
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"discouragedAPIUse"` != string(result) {
		t.Errorf("Expected '\"discouragedAPIUse\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"discouragedAPIUse"`), &enum)
	if log.Name.DiscouragedAPIUse != enum {
		t.Errorf("Expcected %d, got %d", log.Name.DiscouragedAPIUse, enum)
	}

	enum = log.Name.Handler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"handler"` != string(result) {
		t.Errorf("Expected '\"handler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"handler"`), &enum)
	if log.Name.Handler != enum {
		t.Errorf("Expcected %d, got %d", log.Name.Handler, enum)
	}

	enum = log.Name.RecurringHandler
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"recurringHandler"` != string(result) {
		t.Errorf("Expected '\"recurringHandler\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"recurringHandler"`), &enum)
	if log.Name.RecurringHandler != enum {
		t.Errorf("Expcected %d, got %d", log.Name.RecurringHandler, enum)
	}
}
