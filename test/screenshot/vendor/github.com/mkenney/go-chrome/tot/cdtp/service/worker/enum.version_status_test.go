package worker

import (
	"encoding/json"
	"testing"
)

func TestEnumVersionStatus(t *testing.T) {
	var enum VersionStatusEnum
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

	enum = VersionStatus.New
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"new"` != string(result) {
		t.Errorf("Expected '\"new\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"new"`), &enum)
	if VersionStatus.New != enum {
		t.Errorf("Expcected %d, got %d", VersionStatus.New, enum)
	}

	enum = VersionStatus.Installing
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"installing"` != string(result) {
		t.Errorf("Expected '\"installing\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"installing"`), &enum)
	if VersionStatus.Installing != enum {
		t.Errorf("Expcected %d, got %d", VersionStatus.Installing, enum)
	}

	enum = VersionStatus.Installed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"installed"` != string(result) {
		t.Errorf("Expected '\"installed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"installed"`), &enum)
	if VersionStatus.Installed != enum {
		t.Errorf("Expcected %d, got %d", VersionStatus.Installed, enum)
	}
}

func TestEnumVersionStatus2(t *testing.T) {
	var enum VersionStatusEnum
	var err error
	var result []byte

	enum = VersionStatus.Activating
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"activating"` != string(result) {
		t.Errorf("Expected '\"activating\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"activating"`), &enum)
	if VersionStatus.Activating != enum {
		t.Errorf("Expcected %d, got %d", VersionStatus.Activating, enum)
	}

	enum = VersionStatus.Activated
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"activated"` != string(result) {
		t.Errorf("Expected '\"activated\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"activated"`), &enum)
	if VersionStatus.Activated != enum {
		t.Errorf("Expcected %d, got %d", VersionStatus.Activated, enum)
	}

	enum = VersionStatus.Redundant
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"redundant"` != string(result) {
		t.Errorf("Expected '\"redundant\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"redundant"`), &enum)
	if VersionStatus.Redundant != enum {
		t.Errorf("Expcected %d, got %d", VersionStatus.Redundant, enum)
	}
}
