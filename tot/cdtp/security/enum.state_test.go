package security

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/security"
)

func TestEnumState(t *testing.T) {
	var enum security.StateEnum
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

	enum = security.State.Unknown
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"unknown"` != string(result) {
		t.Errorf("Expected '\"unknown\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"unknown"`), &enum)
	if security.State.Unknown != enum {
		t.Errorf("Expcected %d, got %d", security.State.Unknown, enum)
	}

	enum = security.State.Neutral
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"neutral"` != string(result) {
		t.Errorf("Expected '\"neutral\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"neutral"`), &enum)
	if security.State.Neutral != enum {
		t.Errorf("Expcected %d, got %d", security.State.Neutral, enum)
	}

	enum = security.State.Insecure
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"insecure"` != string(result) {
		t.Errorf("Expected '\"insecure\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"insecure"`), &enum)
	if security.State.Insecure != enum {
		t.Errorf("Expcected %d, got %d", security.State.Insecure, enum)
	}
}

func TestEnumState2(t *testing.T) {
	var enum security.StateEnum
	var err error
	var result []byte

	enum = security.State.Secure
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"secure"` != string(result) {
		t.Errorf("Expected '\"secure\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"secure"`), &enum)
	if security.State.Secure != enum {
		t.Errorf("Expcected %d, got %d", security.State.Secure, enum)
	}

	enum = security.State.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if security.State.Info != enum {
		t.Errorf("Expcected %d, got %d", security.State.Info, enum)
	}
}
