package security_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/security"
)

func TestEnumCertificateErrorAction(t *testing.T) {
	var enum security.CertificateErrorActionEnum
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

	enum = security.CertificateErrorAction.Continue
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"continue"` != string(result) {
		t.Errorf("Expected '\"continue\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"continue"`), &enum)
	if security.CertificateErrorAction.Continue != enum {
		t.Errorf("Expcected %d, got %d", security.CertificateErrorAction.Continue, enum)
	}

	enum = security.CertificateErrorAction.Cancel
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cancel"` != string(result) {
		t.Errorf("Expected '\"cancel\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cancel"`), &enum)
	if security.CertificateErrorAction.Cancel != enum {
		t.Errorf("Expcected %d, got %d", security.CertificateErrorAction.Cancel, enum)
	}
}
