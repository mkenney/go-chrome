package worker_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/service/worker"
)

func TestEnumVersionRunningStatus(t *testing.T) {
	var enum worker.VersionRunningStatusEnum
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

	enum = worker.VersionRunningStatus.Stopped
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"stopped"` != string(result) {
		t.Errorf("Expected '\"stopped\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"stopped"`), &enum)
	if worker.VersionRunningStatus.Stopped != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionRunningStatus.Stopped, enum)
	}

	enum = worker.VersionRunningStatus.Starting
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"starting"` != string(result) {
		t.Errorf("Expected '\"starting\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"starting"`), &enum)
	if worker.VersionRunningStatus.Starting != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionRunningStatus.Starting, enum)
	}

	enum = worker.VersionRunningStatus.Running
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"running"` != string(result) {
		t.Errorf("Expected '\"running\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"running"`), &enum)
	if worker.VersionRunningStatus.Running != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionRunningStatus.Running, enum)
	}
}

func TestEnumVersionRunningStatus2(t *testing.T) {
	var enum worker.VersionRunningStatusEnum
	var err error
	var result []byte

	enum = worker.VersionRunningStatus.Stopping
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"stopping"` != string(result) {
		t.Errorf("Expected '\"stopping\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"stopping"`), &enum)
	if worker.VersionRunningStatus.Stopping != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionRunningStatus.Stopping, enum)
	}
}
