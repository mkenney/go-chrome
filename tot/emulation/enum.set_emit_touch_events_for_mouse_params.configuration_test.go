package emulation

import (
	"encoding/json"
	"testing"
)

func TestEnumConfiguration(t *testing.T) {
	var enum ConfigurationEnum
	var err error
	var result []byte

	err = json.Unmarshal([]byte(`""`), &enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}

	err = json.Unmarshal([]byte(`"invalid value"`), &enum)
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

	enum = Configuration.Mobile
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mobile"` != string(result) {
		t.Errorf("Expected '\"mobile\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mobile"`), &enum)
	if Configuration.Mobile != enum {
		t.Errorf("Expcected %d, got %d", Configuration.Mobile, enum)
	}

	enum = Configuration.Desktop
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"desktop"` != string(result) {
		t.Errorf("Expected '\"desktop\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"desktop"`), &enum)
	if Configuration.Desktop != enum {
		t.Errorf("Expcected %d, got %d", Configuration.Desktop, enum)
	}
}
