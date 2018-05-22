package security

import (
	"encoding/json"
	"testing"
)

func TestEnumMixedContentType(t *testing.T) {
	var enum MixedContentTypeEnum
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

	enum = MixedContentType.Blockable
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"blockable"` != string(result) {
		t.Errorf("Expected '\"blockable\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"blockable"`), &enum)
	if MixedContentType.Blockable != enum {
		t.Errorf("Expcected %d, got %d", MixedContentType.Blockable, enum)
	}

	enum = MixedContentType.OptionallyBlockable
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"optionally-blockable"` != string(result) {
		t.Errorf("Expected '\"optionally-blockable\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"optionally-blockable"`), &enum)
	if MixedContentType.OptionallyBlockable != enum {
		t.Errorf("Expcected %d, got %d", MixedContentType.OptionallyBlockable, enum)
	}

	enum = MixedContentType.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if MixedContentType.None != enum {
		t.Errorf("Expcected %d, got %d", MixedContentType.None, enum)
	}
}
