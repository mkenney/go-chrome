package network_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumInterceptionStage(t *testing.T) {
	var enum network.InterceptionStageEnum
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

	enum = network.InterceptionStage.Request
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Request"` != string(result) {
		t.Errorf("Expected '\"Request\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Request"`), &enum)
	if network.InterceptionStage.Request != enum {
		t.Errorf("Expcected %d, got %d", network.InterceptionStage.Request, enum)
	}

	enum = network.InterceptionStage.HeadersReceived
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"HeadersReceived"` != string(result) {
		t.Errorf("Expected '\"HeadersReceived\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"HeadersReceived"`), &enum)
	if network.InterceptionStage.HeadersReceived != enum {
		t.Errorf("Expcected %d, got %d", network.InterceptionStage.HeadersReceived, enum)
	}
}
