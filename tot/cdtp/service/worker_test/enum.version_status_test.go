package worker_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/service/worker"
)

func TestEnumVersionStatus(t *testing.T) {
	var enum worker.VersionStatusEnum
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

	enum = worker.VersionStatus.New
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"new"` != string(result) {
		t.Errorf("Expected '\"new\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"new"`), &enum)
	if worker.VersionStatus.New != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionStatus.New, enum)
	}

	enum = worker.VersionStatus.Installing
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"installing"` != string(result) {
		t.Errorf("Expected '\"installing\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"installing"`), &enum)
	if worker.VersionStatus.Installing != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionStatus.Installing, enum)
	}

	enum = worker.VersionStatus.Installed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"installed"` != string(result) {
		t.Errorf("Expected '\"installed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"installed"`), &enum)
	if worker.VersionStatus.Installed != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionStatus.Installed, enum)
	}
}

func TestEnumVersionStatus2(t *testing.T) {
	var enum worker.VersionStatusEnum
	var err error
	var result []byte

	enum = worker.VersionStatus.Activating
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"activating"` != string(result) {
		t.Errorf("Expected '\"activating\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"activating"`), &enum)
	if worker.VersionStatus.Activating != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionStatus.Activating, enum)
	}

	enum = worker.VersionStatus.Activated
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"activated"` != string(result) {
		t.Errorf("Expected '\"activated\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"activated"`), &enum)
	if worker.VersionStatus.Activated != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionStatus.Activated, enum)
	}

	enum = worker.VersionStatus.Redundant
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"redundant"` != string(result) {
		t.Errorf("Expected '\"redundant\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"redundant"`), &enum)
	if worker.VersionStatus.Redundant != enum {
		t.Errorf("Expcected %d, got %d", worker.VersionStatus.Redundant, enum)
	}
}
