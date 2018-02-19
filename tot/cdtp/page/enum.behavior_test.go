package page

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/page"
)

func TestEnumBehavior(t *testing.T) {
	var enum page.BehaviorEnum
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

	enum = page.Behavior.Deny
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"deny"` != string(result) {
		t.Errorf("Expected '\"deny\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"deny"`), &enum)
	if page.Behavior.Deny != enum {
		t.Errorf("Expcected %d, got %d", page.Behavior.Deny, enum)
	}

	enum = page.Behavior.Allow
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"allow"` != string(result) {
		t.Errorf("Expected '\"allow\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"allow"`), &enum)
	if page.Behavior.Allow != enum {
		t.Errorf("Expcected %d, got %d", page.Behavior.Allow, enum)
	}

	enum = page.Behavior.Default
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"default"` != string(result) {
		t.Errorf("Expected '\"default\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"default"`), &enum)
	if page.Behavior.Default != enum {
		t.Errorf("Expcected %d, got %d", page.Behavior.Default, enum)
	}
}
