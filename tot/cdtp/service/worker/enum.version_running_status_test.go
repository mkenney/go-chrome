package worker

import (
	"encoding/json"
	"testing"
)

func TestEnumVersionRunningStatus(t *testing.T) {
	var enum VersionRunningStatusEnum
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

	enum = VersionRunningStatus.Stopped
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"stopped"` != string(result) {
		t.Errorf("Expected '\"stopped\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"stopped"`), &enum)
	if VersionRunningStatus.Stopped != enum {
		t.Errorf("Expcected %d, got %d", VersionRunningStatus.Stopped, enum)
	}

	enum = VersionRunningStatus.Starting
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"starting"` != string(result) {
		t.Errorf("Expected '\"starting\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"starting"`), &enum)
	if VersionRunningStatus.Starting != enum {
		t.Errorf("Expcected %d, got %d", VersionRunningStatus.Starting, enum)
	}

	enum = VersionRunningStatus.Running
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"running"` != string(result) {
		t.Errorf("Expected '\"running\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"running"`), &enum)
	if VersionRunningStatus.Running != enum {
		t.Errorf("Expcected %d, got %d", VersionRunningStatus.Running, enum)
	}
}

func TestEnumVersionRunningStatus2(t *testing.T) {
	var enum VersionRunningStatusEnum
	var err error
	var result []byte

	enum = VersionRunningStatus.Stopping
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"stopping"` != string(result) {
		t.Errorf("Expected '\"stopping\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"stopping"`), &enum)
	if VersionRunningStatus.Stopping != enum {
		t.Errorf("Expcected %d, got %d", VersionRunningStatus.Stopping, enum)
	}
}
