package css_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/css"
)

func TestEnumStyleSheetOrigin(t *testing.T) {
	var enum css.StyleSheetOriginEnum
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

	enum = css.StyleSheetOrigin.Injected
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"injected"` != string(result) {
		t.Errorf("Expected '\"injected\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"injected"`), &enum)
	if css.StyleSheetOrigin.Injected != enum {
		t.Errorf("Expcected %d, got %d", css.StyleSheetOrigin.Injected, enum)
	}

	enum = css.StyleSheetOrigin.UserAgent
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"user-agent"` != string(result) {
		t.Errorf("Expected '\"user-agent\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"user-agent"`), &enum)
	if css.StyleSheetOrigin.UserAgent != enum {
		t.Errorf("Expcected %d, got %d", css.StyleSheetOrigin.UserAgent, enum)
	}

	enum = css.StyleSheetOrigin.Inspector
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"inspector"` != string(result) {
		t.Errorf("Expected '\"inspector\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"inspector"`), &enum)
	if css.StyleSheetOrigin.Inspector != enum {
		t.Errorf("Expcected %d, got %d", css.StyleSheetOrigin.Inspector, enum)
	}
}

func TestEnumStyleSheetOrigin2(t *testing.T) {
	var enum css.StyleSheetOriginEnum
	var err error
	var result []byte

	enum = css.StyleSheetOrigin.Log
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"log"` != string(result) {
		t.Errorf("Expected '\"log\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"log"`), &enum)
	if css.StyleSheetOrigin.Log != enum {
		t.Errorf("Expcected %d, got %d", css.StyleSheetOrigin.Log, enum)
	}
}
