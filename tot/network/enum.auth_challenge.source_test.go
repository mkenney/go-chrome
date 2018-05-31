package network

import (
	"encoding/json"
	"testing"
)

func TestEnumSource(t *testing.T) {
	var enum SourceEnum
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

	enum = Source.Server
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Server"` != string(result) {
		t.Errorf("Expected '\"Server\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Server"`), &enum)
	if Source.Server != enum {
		t.Errorf("Expcected %d, got %d", Source.Server, enum)
	}

	enum = Source.Proxy
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Proxy"` != string(result) {
		t.Errorf("Expected '\"Proxy\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Proxy"`), &enum)
	if Source.Proxy != enum {
		t.Errorf("Expcected %d, got %d", Source.Proxy, enum)
	}
}
