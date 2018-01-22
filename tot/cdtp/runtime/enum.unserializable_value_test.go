package runtime

import (
	"encoding/json"
	"testing"
)

func TestEnumUnserializableValue(t *testing.T) {
	var enum UnserializableValueEnum
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

	enum = UnserializableValue.Infinity
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Infinity"` != string(result) {
		t.Errorf("Expected '\"Infinity\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Infinity"`), &enum)
	if UnserializableValue.Infinity != enum {
		t.Errorf("Expcected %d, got %d", UnserializableValue.Infinity, enum)
	}

	enum = UnserializableValue.NaN
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"NaN"` != string(result) {
		t.Errorf("Expected '\"NaN\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"NaN"`), &enum)
	if UnserializableValue.NaN != enum {
		t.Errorf("Expcected %d, got %d", UnserializableValue.NaN, enum)
	}

	enum = UnserializableValue.NegInfinity
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"-Infinity"` != string(result) {
		t.Errorf("Expected '\"-Infinity\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"-Infinity"`), &enum)
	if UnserializableValue.NegInfinity != enum {
		t.Errorf("Expcected %d, got %d", UnserializableValue.NegInfinity, enum)
	}

	enum = UnserializableValue.NegZero
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"-0"` != string(result) {
		t.Errorf("Expected '\"-0\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"-0"`), &enum)
	if UnserializableValue.NegZero != enum {
		t.Errorf("Expcected %d, got %d", UnserializableValue.NegZero, enum)
	}
}
