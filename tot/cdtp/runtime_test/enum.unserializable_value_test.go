package runtime_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestEnumUnserializableValue(t *testing.T) {
	var enum runtime.UnserializableValueEnum
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

	enum = runtime.UnserializableValue.Infinity
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Infinity"` != string(result) {
		t.Errorf("Expected '\"Infinity\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Infinity"`), &enum)
	if runtime.UnserializableValue.Infinity != enum {
		t.Errorf("Expcected %d, got %d", runtime.UnserializableValue.Infinity, enum)
	}

	enum = runtime.UnserializableValue.NaN
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"NaN"` != string(result) {
		t.Errorf("Expected '\"NaN\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"NaN"`), &enum)
	if runtime.UnserializableValue.NaN != enum {
		t.Errorf("Expcected %d, got %d", runtime.UnserializableValue.NaN, enum)
	}

	enum = runtime.UnserializableValue.NegInfinity
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"-Infinity"` != string(result) {
		t.Errorf("Expected '\"-Infinity\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"-Infinity"`), &enum)
	if runtime.UnserializableValue.NegInfinity != enum {
		t.Errorf("Expcected %d, got %d", runtime.UnserializableValue.NegInfinity, enum)
	}
}

func TestEnumUnserializableValue2(t *testing.T) {
	var enum runtime.UnserializableValueEnum
	var err error
	var result []byte

	enum = runtime.UnserializableValue.NegZero
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"-0"` != string(result) {
		t.Errorf("Expected '\"-0\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"-0"`), &enum)
	if runtime.UnserializableValue.NegZero != enum {
		t.Errorf("Expcected %d, got %d", runtime.UnserializableValue.NegZero, enum)
	}
}
