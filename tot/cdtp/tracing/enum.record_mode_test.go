package tracing

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/tracing"
)

func TestEnumRecordMode(t *testing.T) {
	var enum tracing.RecordModeEnum
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

	enum = tracing.RecordMode.RecordUntilFull
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"recordUntilFull"` != string(result) {
		t.Errorf("Expected '\"recordUntilFull\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"recordUntilFull"`), &enum)
	if tracing.RecordMode.RecordUntilFull != enum {
		t.Errorf("Expcected %d, got %d", tracing.RecordMode.RecordUntilFull, enum)
	}

	enum = tracing.RecordMode.RecordContinuously
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"recordContinuously"` != string(result) {
		t.Errorf("Expected '\"recordContinuously\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"recordContinuously"`), &enum)
	if tracing.RecordMode.RecordContinuously != enum {
		t.Errorf("Expcected %d, got %d", tracing.RecordMode.RecordContinuously, enum)
	}

	enum = tracing.RecordMode.RecordAsMuchAsPossible
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"recordAsMuchAsPossible"` != string(result) {
		t.Errorf("Expected '\"recordAsMuchAsPossible\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"recordAsMuchAsPossible"`), &enum)
	if tracing.RecordMode.RecordAsMuchAsPossible != enum {
		t.Errorf("Expcected %d, got %d", tracing.RecordMode.RecordAsMuchAsPossible, enum)
	}
}

func TestEnumRecordMode2(t *testing.T) {
	var enum tracing.RecordModeEnum
	var err error
	var result []byte

	enum = tracing.RecordMode.EchoToConsole
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"echoToConsole"` != string(result) {
		t.Errorf("Expected '\"echoToConsole\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"echoToConsole"`), &enum)
	if tracing.RecordMode.EchoToConsole != enum {
		t.Errorf("Expcected %d, got %d", tracing.RecordMode.EchoToConsole, enum)
	}
}
