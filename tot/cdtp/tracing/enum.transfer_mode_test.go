package tracing

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/tracing"
)

func TestEnumTransferMode(t *testing.T) {
	var enum tracing.TransferModeEnum
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

	enum = tracing.TransferMode.ReportEvents
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ReportEvents"` != string(result) {
		t.Errorf("Expected '\"ReportEvents\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ReportEvents"`), &enum)
	if tracing.TransferMode.ReportEvents != enum {
		t.Errorf("Expcected %d, got %d", tracing.TransferMode.ReportEvents, enum)
	}

	enum = tracing.TransferMode.ReturnAsStream
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ReturnAsStream"` != string(result) {
		t.Errorf("Expected '\"ReturnAsStream\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ReturnAsStream"`), &enum)
	if tracing.TransferMode.ReturnAsStream != enum {
		t.Errorf("Expcected %d, got %d", tracing.TransferMode.ReturnAsStream, enum)
	}
}
