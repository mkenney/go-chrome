package console

import (
	"encoding/json"
	"testing"
)

func TestEnumMessageSource(t *testing.T) {
	var enum MessageSourceEnum
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

	enum = MessageSource.XML
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"xml"` != string(result) {
		t.Errorf("Expected '\"xml\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"xml"`), &enum)
	if MessageSource.XML != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.XML, enum)
	}

	enum = MessageSource.Javascript
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"javascript"` != string(result) {
		t.Errorf("Expected '\"javascript\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"javascript"`), &enum)
	if MessageSource.Javascript != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Javascript, enum)
	}

	enum = MessageSource.Network
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"network"` != string(result) {
		t.Errorf("Expected '\"network\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"network"`), &enum)
	if MessageSource.Network != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Network, enum)
	}

	enum = MessageSource.ConsoleAPI
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"console-api"` != string(result) {
		t.Errorf("Expected '\"console-api\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"console-api"`), &enum)
	if MessageSource.ConsoleAPI != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.ConsoleAPI, enum)
	}

	enum = MessageSource.Storage
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"storage"` != string(result) {
		t.Errorf("Expected '\"storage\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"storage"`), &enum)
	if MessageSource.Storage != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Storage, enum)
	}

	enum = MessageSource.Appcache
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"appcache"` != string(result) {
		t.Errorf("Expected '\"appcache\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"appcache"`), &enum)
	if MessageSource.Appcache != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Appcache, enum)
	}

	enum = MessageSource.Rendering
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"rendering"` != string(result) {
		t.Errorf("Expected '\"rendering\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"rendering"`), &enum)
	if MessageSource.Rendering != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Rendering, enum)
	}

	enum = MessageSource.Security
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"security"` != string(result) {
		t.Errorf("Expected '\"security\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"security"`), &enum)
	if MessageSource.Security != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Security, enum)
	}

	enum = MessageSource.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if MessageSource.Other != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Other, enum)
	}

	enum = MessageSource.Deprecation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"deprecation"` != string(result) {
		t.Errorf("Expected '\"deprecation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"deprecation"`), &enum)
	if MessageSource.Deprecation != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Deprecation, enum)
	}

	enum = MessageSource.Worker
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"worker"` != string(result) {
		t.Errorf("Expected '\"worker\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"worker"`), &enum)
	if MessageSource.Worker != enum {
		t.Errorf("Expcected %d, got %d", MessageSource.Worker, enum)
	}
}
