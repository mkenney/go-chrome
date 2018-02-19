package page_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/page"
)

func TestEnumDialogType(t *testing.T) {
	var enum page.DialogTypeEnum
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

	enum = page.DialogType.Alert
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"alert"` != string(result) {
		t.Errorf("Expected '\"alert\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"alert"`), &enum)
	if page.DialogType.Alert != enum {
		t.Errorf("Expcected %d, got %d", page.DialogType.Alert, enum)
	}

	enum = page.DialogType.Confirm
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"confirm"` != string(result) {
		t.Errorf("Expected '\"confirm\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"confirm"`), &enum)
	if page.DialogType.Confirm != enum {
		t.Errorf("Expcected %d, got %d", page.DialogType.Confirm, enum)
	}

	enum = page.DialogType.Prompt
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"prompt"` != string(result) {
		t.Errorf("Expected '\"prompt\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"prompt"`), &enum)
	if page.DialogType.Prompt != enum {
		t.Errorf("Expcected %d, got %d", page.DialogType.Prompt, enum)
	}
}

func TestEnumDialogType2(t *testing.T) {
	var enum page.DialogTypeEnum
	var err error
	var result []byte

	enum = page.DialogType.Beforeunload
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"beforeunload"` != string(result) {
		t.Errorf("Expected '\"beforeunload\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"beforeunload"`), &enum)
	if page.DialogType.Beforeunload != enum {
		t.Errorf("Expcected %d, got %d", page.DialogType.Beforeunload, enum)
	}
}
