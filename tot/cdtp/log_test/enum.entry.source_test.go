package log_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/log"
)

func TestEnumSource(t *testing.T) {
	var enum log.SourceEnum
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

	enum = log.Source.XML
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"xml"` != string(result) {
		t.Errorf("Expected '\"xml\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"xml"`), &enum)
	if log.Source.XML != enum {
		t.Errorf("Expcected %d, got %d", log.Source.XML, enum)
	}

	enum = log.Source.Javascript
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"javascript"` != string(result) {
		t.Errorf("Expected '\"javascript\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"javascript"`), &enum)
	if log.Source.Javascript != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Javascript, enum)
	}

	enum = log.Source.Network
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"network"` != string(result) {
		t.Errorf("Expected '\"network\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"network"`), &enum)
	if log.Source.Network != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Network, enum)
	}
}

func TestEnumSource2(t *testing.T) {
	var enum log.SourceEnum
	var err error
	var result []byte

	enum = log.Source.Storage
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"storage"` != string(result) {
		t.Errorf("Expected '\"storage\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"storage"`), &enum)
	if log.Source.Storage != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Storage, enum)
	}

	enum = log.Source.Appcache
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"appcache"` != string(result) {
		t.Errorf("Expected '\"appcache\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"appcache"`), &enum)
	if log.Source.Appcache != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Appcache, enum)
	}

	enum = log.Source.Rendering
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"rendering"` != string(result) {
		t.Errorf("Expected '\"rendering\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"rendering"`), &enum)
	if log.Source.Rendering != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Rendering, enum)
	}

	enum = log.Source.Security
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"security"` != string(result) {
		t.Errorf("Expected '\"security\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"security"`), &enum)
	if log.Source.Security != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Security, enum)
	}
}

func TestEnumSource3(t *testing.T) {
	var enum log.SourceEnum
	var err error
	var result []byte

	enum = log.Source.Deprecation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"deprecation"` != string(result) {
		t.Errorf("Expected '\"deprecation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"deprecation"`), &enum)
	if log.Source.Deprecation != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Deprecation, enum)
	}

	enum = log.Source.Worker
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"worker"` != string(result) {
		t.Errorf("Expected '\"worker\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"worker"`), &enum)
	if log.Source.Worker != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Worker, enum)
	}

	enum = log.Source.Violation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"violation"` != string(result) {
		t.Errorf("Expected '\"violation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"violation"`), &enum)
	if log.Source.Violation != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Violation, enum)
	}

	enum = log.Source.Intervention
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"intervention"` != string(result) {
		t.Errorf("Expected '\"intervention\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"intervention"`), &enum)
	if log.Source.Intervention != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Intervention, enum)
	}
}

func TestEnumSource4(t *testing.T) {
	var enum log.SourceEnum
	var err error
	var result []byte

	enum = log.Source.Recommendation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"recommendation"` != string(result) {
		t.Errorf("Expected '\"recommendation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"recommendation"`), &enum)
	if log.Source.Recommendation != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Recommendation, enum)
	}

	enum = log.Source.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if log.Source.Other != enum {
		t.Errorf("Expcected %d, got %d", log.Source.Other, enum)
	}
}
