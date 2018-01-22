package page

import (
	"encoding/json"
	"testing"
)

func TestEnumDialogType(t *testing.T) {
	var enum DialogTypeEnum
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

	enum = DialogType.Alert
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"alert"` != string(result) {
		t.Errorf("Expected '\"alert\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"alert"`), &enum)
	if DialogType.Alert != enum {
		t.Errorf("Expcected %d, got %d", DialogType.Alert, enum)
	}

	enum = DialogType.Confirm
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"confirm"` != string(result) {
		t.Errorf("Expected '\"confirm\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"confirm"`), &enum)
	if DialogType.Confirm != enum {
		t.Errorf("Expcected %d, got %d", DialogType.Confirm, enum)
	}

	enum = DialogType.Prompt
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"prompt"` != string(result) {
		t.Errorf("Expected '\"prompt\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"prompt"`), &enum)
	if DialogType.Prompt != enum {
		t.Errorf("Expcected %d, got %d", DialogType.Prompt, enum)
	}

	enum = DialogType.Beforeunload
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"beforeunload"` != string(result) {
		t.Errorf("Expected '\"beforeunload\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"beforeunload"`), &enum)
	if DialogType.Beforeunload != enum {
		t.Errorf("Expcected %d, got %d", DialogType.Beforeunload, enum)
	}
}
