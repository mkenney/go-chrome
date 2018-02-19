package emulation_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/emulation"
)

func TestEnumConfiguration(t *testing.T) {
	var enum emulation.ConfigurationEnum
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

	enum = emulation.Configuration.Mobile
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mobile"` != string(result) {
		t.Errorf("Expected '\"mobile\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mobile"`), &enum)
	if emulation.Configuration.Mobile != enum {
		t.Errorf("Expcected %d, got %d", emulation.Configuration.Mobile, enum)
	}

	enum = emulation.Configuration.Desktop
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"desktop"` != string(result) {
		t.Errorf("Expected '\"desktop\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"desktop"`), &enum)
	if emulation.Configuration.Desktop != enum {
		t.Errorf("Expcected %d, got %d", emulation.Configuration.Desktop, enum)
	}
}
