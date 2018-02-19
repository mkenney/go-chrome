package tracing

import (
	"encoding/json"
	"testing"
)

func TestEnumTransferMode(t *testing.T) {
	var enum TransferModeEnum
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

	enum = TransferMode.ReportEvents
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ReportEvents"` != string(result) {
		t.Errorf("Expected '\"ReportEvents\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ReportEvents"`), &enum)
	if TransferMode.ReportEvents != enum {
		t.Errorf("Expcected %d, got %d", TransferMode.ReportEvents, enum)
	}

	enum = TransferMode.ReturnAsStream
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ReturnAsStream"` != string(result) {
		t.Errorf("Expected '\"ReturnAsStream\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ReturnAsStream"`), &enum)
	if TransferMode.ReturnAsStream != enum {
		t.Errorf("Expcected %d, got %d", TransferMode.ReturnAsStream, enum)
	}
}
