package log

import (
	"encoding/json"
	"testing"
)

func TestEnumSource(t *testing.T) {
	var enum SourceEnum
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

	enum = Source.XML
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"xml"` != string(result) {
		t.Errorf("Expected '\"xml\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"xml"`), &enum)
	if Source.XML != enum {
		t.Errorf("Expcected %d, got %d", Source.XML, enum)
	}

	enum = Source.Javascript
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"javascript"` != string(result) {
		t.Errorf("Expected '\"javascript\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"javascript"`), &enum)
	if Source.Javascript != enum {
		t.Errorf("Expcected %d, got %d", Source.Javascript, enum)
	}

	enum = Source.Network
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"network"` != string(result) {
		t.Errorf("Expected '\"network\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"network"`), &enum)
	if Source.Network != enum {
		t.Errorf("Expcected %d, got %d", Source.Network, enum)
	}

	enum = Source.Storage
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"storage"` != string(result) {
		t.Errorf("Expected '\"storage\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"storage"`), &enum)
	if Source.Storage != enum {
		t.Errorf("Expcected %d, got %d", Source.Storage, enum)
	}

	enum = Source.Appcache
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"appcache"` != string(result) {
		t.Errorf("Expected '\"appcache\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"appcache"`), &enum)
	if Source.Appcache != enum {
		t.Errorf("Expcected %d, got %d", Source.Appcache, enum)
	}

	enum = Source.Rendering
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"rendering"` != string(result) {
		t.Errorf("Expected '\"rendering\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"rendering"`), &enum)
	if Source.Rendering != enum {
		t.Errorf("Expcected %d, got %d", Source.Rendering, enum)
	}

	enum = Source.Security
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"security"` != string(result) {
		t.Errorf("Expected '\"security\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"security"`), &enum)
	if Source.Security != enum {
		t.Errorf("Expcected %d, got %d", Source.Security, enum)
	}

	enum = Source.Deprecation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"deprecation"` != string(result) {
		t.Errorf("Expected '\"deprecation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"deprecation"`), &enum)
	if Source.Deprecation != enum {
		t.Errorf("Expcected %d, got %d", Source.Deprecation, enum)
	}

	enum = Source.Worker
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"worker"` != string(result) {
		t.Errorf("Expected '\"worker\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"worker"`), &enum)
	if Source.Worker != enum {
		t.Errorf("Expcected %d, got %d", Source.Worker, enum)
	}

	enum = Source.Violation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"violation"` != string(result) {
		t.Errorf("Expected '\"violation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"violation"`), &enum)
	if Source.Violation != enum {
		t.Errorf("Expcected %d, got %d", Source.Violation, enum)
	}

	enum = Source.Intervention
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"intervention"` != string(result) {
		t.Errorf("Expected '\"intervention\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"intervention"`), &enum)
	if Source.Intervention != enum {
		t.Errorf("Expcected %d, got %d", Source.Intervention, enum)
	}

	enum = Source.Recommendation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"recommendation"` != string(result) {
		t.Errorf("Expected '\"recommendation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"recommendation"`), &enum)
	if Source.Recommendation != enum {
		t.Errorf("Expcected %d, got %d", Source.Recommendation, enum)
	}

	enum = Source.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if Source.Other != enum {
		t.Errorf("Expcected %d, got %d", Source.Other, enum)
	}
}
