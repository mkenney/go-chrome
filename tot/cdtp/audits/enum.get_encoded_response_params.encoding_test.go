package audits

import (
	"encoding/json"
	"testing"
)

func TestEnumEncoding(t *testing.T) {
	var enum EncodingEnum
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

	enum = Encoding.Webp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"webp"` != string(result) {
		t.Errorf("Expected '\"webp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"webp"`), &enum)
	if Encoding.Webp != enum {
		t.Errorf("Expcected %d, got %d", Encoding.Webp, enum)
	}

	enum = Encoding.Jpeg
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"jpeg"` != string(result) {
		t.Errorf("Expected '\"jpeg\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"jpeg"`), &enum)
	if Encoding.Jpeg != enum {
		t.Errorf("Expcected %d, got %d", Encoding.Jpeg, enum)
	}

	enum = Encoding.Png
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"png"` != string(result) {
		t.Errorf("Expected '\"png\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"png"`), &enum)
	if Encoding.Png != enum {
		t.Errorf("Expcected %d, got %d", Encoding.Png, enum)
	}
}
