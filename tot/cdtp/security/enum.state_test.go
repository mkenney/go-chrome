package security

import (
	"encoding/json"
	"testing"
)

func TestEnumState(t *testing.T) {
	var enum StateEnum
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

	enum = State.Unknown
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"unknown"` != string(result) {
		t.Errorf("Expected '\"unknown\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"unknown"`), &enum)
	if State.Unknown != enum {
		t.Errorf("Expcected %d, got %d", State.Unknown, enum)
	}

	enum = State.Neutral
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"neutral"` != string(result) {
		t.Errorf("Expected '\"neutral\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"neutral"`), &enum)
	if State.Neutral != enum {
		t.Errorf("Expcected %d, got %d", State.Neutral, enum)
	}

	enum = State.Insecure
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"insecure"` != string(result) {
		t.Errorf("Expected '\"insecure\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"insecure"`), &enum)
	if State.Insecure != enum {
		t.Errorf("Expcected %d, got %d", State.Insecure, enum)
	}
}

func TestEnumState2(t *testing.T) {
	var enum StateEnum
	var err error
	var result []byte

	enum = State.Secure
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"secure"` != string(result) {
		t.Errorf("Expected '\"secure\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"secure"`), &enum)
	if State.Secure != enum {
		t.Errorf("Expcected %d, got %d", State.Secure, enum)
	}

	enum = State.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if State.Info != enum {
		t.Errorf("Expcected %d, got %d", State.Info, enum)
	}
}
