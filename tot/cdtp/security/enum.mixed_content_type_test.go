package security

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/security"
)

func TestEnumMixedContentType(t *testing.T) {
	var enum security.MixedContentTypeEnum
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

	enum = security.MixedContentType.Blockable
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"blockable"` != string(result) {
		t.Errorf("Expected '\"blockable\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"blockable"`), &enum)
	if security.MixedContentType.Blockable != enum {
		t.Errorf("Expcected %d, got %d", security.MixedContentType.Blockable, enum)
	}

	enum = security.MixedContentType.OptionallyBlockable
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"optionally-blockable"` != string(result) {
		t.Errorf("Expected '\"optionally-blockable\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"optionally-blockable"`), &enum)
	if security.MixedContentType.OptionallyBlockable != enum {
		t.Errorf("Expcected %d, got %d", security.MixedContentType.OptionallyBlockable, enum)
	}

	enum = security.MixedContentType.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if security.MixedContentType.None != enum {
		t.Errorf("Expcected %d, got %d", security.MixedContentType.None, enum)
	}
}
