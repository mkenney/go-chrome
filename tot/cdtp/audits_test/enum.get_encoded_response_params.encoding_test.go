package audits_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/audits"
)

func TestEnumEncoding(t *testing.T) {
	var enum audits.EncodingEnum
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

	enum = audits.Encoding.Webp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"webp"` != string(result) {
		t.Errorf("Expected '\"webp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"webp"`), &enum)
	if audits.Encoding.Webp != enum {
		t.Errorf("Expcected %d, got %d", audits.Encoding.Webp, enum)
	}

	enum = audits.Encoding.Jpeg
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"jpeg"` != string(result) {
		t.Errorf("Expected '\"jpeg\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"jpeg"`), &enum)
	if audits.Encoding.Jpeg != enum {
		t.Errorf("Expcected %d, got %d", audits.Encoding.Jpeg, enum)
	}

	enum = audits.Encoding.Png
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"png"` != string(result) {
		t.Errorf("Expected '\"png\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"png"`), &enum)
	if audits.Encoding.Png != enum {
		t.Errorf("Expcected %d, got %d", audits.Encoding.Png, enum)
	}
}
