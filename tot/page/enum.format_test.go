package page

import (
	"encoding/json"
	"testing"
)

func TestEnumFormat(t *testing.T) {
	var enum FormatEnum
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

	enum = Format.Png
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"png"` != string(result) {
		t.Errorf("Expected '\"png\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"png"`), &enum)
	if Format.Png != enum {
		t.Errorf("Expcected %d, got %d", Format.Png, enum)
	}

	enum = Format.Jpeg
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"jpeg"` != string(result) {
		t.Errorf("Expected '\"jpeg\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"jpeg"`), &enum)
	if Format.Jpeg != enum {
		t.Errorf("Expcected %d, got %d", Format.Jpeg, enum)
	}
}
