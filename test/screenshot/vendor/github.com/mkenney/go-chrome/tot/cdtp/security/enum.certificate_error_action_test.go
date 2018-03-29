package security

import (
	"encoding/json"
	"testing"
)

func TestEnumCertificateErrorAction(t *testing.T) {
	var enum CertificateErrorActionEnum
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

	enum = CertificateErrorAction.Continue
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"continue"` != string(result) {
		t.Errorf("Expected '\"continue\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"continue"`), &enum)
	if CertificateErrorAction.Continue != enum {
		t.Errorf("Expcected %d, got %d", CertificateErrorAction.Continue, enum)
	}

	enum = CertificateErrorAction.Cancel
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cancel"` != string(result) {
		t.Errorf("Expected '\"cancel\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cancel"`), &enum)
	if CertificateErrorAction.Cancel != enum {
		t.Errorf("Expcected %d, got %d", CertificateErrorAction.Cancel, enum)
	}
}
