package network

import (
	"encoding/json"
	"testing"
)

func TestEnumInterceptionStage(t *testing.T) {
	var enum InterceptionStageEnum
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

	enum = InterceptionStage.Request
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Request"` != string(result) {
		t.Errorf("Expected '\"Request\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Request"`), &enum)
	if InterceptionStage.Request != enum {
		t.Errorf("Expcected %d, got %d", InterceptionStage.Request, enum)
	}

	enum = InterceptionStage.HeadersReceived
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"HeadersReceived"` != string(result) {
		t.Errorf("Expected '\"HeadersReceived\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"HeadersReceived"`), &enum)
	if InterceptionStage.HeadersReceived != enum {
		t.Errorf("Expcected %d, got %d", InterceptionStage.HeadersReceived, enum)
	}
}
