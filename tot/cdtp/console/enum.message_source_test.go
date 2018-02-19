package console

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/console"
)

func TestEnumMessageSource(t *testing.T) {
	var enum console.MessageSourceEnum
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

	enum = console.MessageSource.XML
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"xml"` != string(result) {
		t.Errorf("Expected '\"xml\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"xml"`), &enum)
	if console.MessageSource.XML != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.XML, enum)
	}

	enum = console.MessageSource.Javascript
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"javascript"` != string(result) {
		t.Errorf("Expected '\"javascript\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"javascript"`), &enum)
	if console.MessageSource.Javascript != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Javascript, enum)
	}

	enum = console.MessageSource.Network
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"network"` != string(result) {
		t.Errorf("Expected '\"network\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"network"`), &enum)
	if console.MessageSource.Network != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Network, enum)
	}
}

func TestEnumMessageSource2(t *testing.T) {
	var enum console.MessageSourceEnum
	var err error
	var result []byte

	enum = console.MessageSource.ConsoleAPI
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"console-api"` != string(result) {
		t.Errorf("Expected '\"console-api\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"console-api"`), &enum)
	if console.MessageSource.ConsoleAPI != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.ConsoleAPI, enum)
	}

	enum = console.MessageSource.Storage
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"storage"` != string(result) {
		t.Errorf("Expected '\"storage\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"storage"`), &enum)
	if console.MessageSource.Storage != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Storage, enum)
	}

	enum = console.MessageSource.Appcache
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"appcache"` != string(result) {
		t.Errorf("Expected '\"appcache\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"appcache"`), &enum)
	if console.MessageSource.Appcache != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Appcache, enum)
	}

	enum = console.MessageSource.Rendering
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"rendering"` != string(result) {
		t.Errorf("Expected '\"rendering\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"rendering"`), &enum)
	if console.MessageSource.Rendering != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Rendering, enum)
	}
}

func TestEnumMessageSource3(t *testing.T) {
	var enum console.MessageSourceEnum
	var err error
	var result []byte

	enum = console.MessageSource.Security
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"security"` != string(result) {
		t.Errorf("Expected '\"security\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"security"`), &enum)
	if console.MessageSource.Security != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Security, enum)
	}

	enum = console.MessageSource.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if console.MessageSource.Other != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Other, enum)
	}

	enum = console.MessageSource.Deprecation
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"deprecation"` != string(result) {
		t.Errorf("Expected '\"deprecation\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"deprecation"`), &enum)
	if console.MessageSource.Deprecation != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Deprecation, enum)
	}

	enum = console.MessageSource.Worker
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"worker"` != string(result) {
		t.Errorf("Expected '\"worker\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"worker"`), &enum)
	if console.MessageSource.Worker != enum {
		t.Errorf("Expcected %d, got %d", console.MessageSource.Worker, enum)
	}
}
