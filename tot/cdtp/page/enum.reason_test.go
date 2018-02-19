package page

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/page"
)

func TestEnumReason(t *testing.T) {
	var enum page.ReasonEnum
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

	enum = page.Reason.FormSubmissionGet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"formSubmissionGet"` != string(result) {
		t.Errorf("Expected '\"formSubmissionGet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"formSubmissionGet"`), &enum)
	if page.Reason.FormSubmissionGet != enum {
		t.Errorf("Expcected %d, got %d", page.Reason.FormSubmissionGet, enum)
	}

	enum = page.Reason.FormSubmissionPost
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"formSubmissionPost"` != string(result) {
		t.Errorf("Expected '\"formSubmissionPost\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"formSubmissionPost"`), &enum)
	if page.Reason.FormSubmissionPost != enum {
		t.Errorf("Expcected %d, got %d", page.Reason.FormSubmissionPost, enum)
	}

	enum = page.Reason.HTTPHeaderRefresh
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"httpHeaderRefresh"` != string(result) {
		t.Errorf("Expected '\"httpHeaderRefresh\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"httpHeaderRefresh"`), &enum)
	if page.Reason.HTTPHeaderRefresh != enum {
		t.Errorf("Expcected %d, got %d", page.Reason.HTTPHeaderRefresh, enum)
	}
}
func TestEnumReason2(t *testing.T) {
	var enum page.ReasonEnum
	var err error
	var result []byte

	enum = page.Reason.ScriptInitiated
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"scriptInitiated"` != string(result) {
		t.Errorf("Expected '\"scriptInitiated\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"scriptInitiated"`), &enum)
	if page.Reason.ScriptInitiated != enum {
		t.Errorf("Expcected %d, got %d", page.Reason.ScriptInitiated, enum)
	}

	enum = page.Reason.MetaTagRefresh
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"metaTagRefresh"` != string(result) {
		t.Errorf("Expected '\"metaTagRefresh\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"metaTagRefresh"`), &enum)
	if page.Reason.MetaTagRefresh != enum {
		t.Errorf("Expcected %d, got %d", page.Reason.MetaTagRefresh, enum)
	}

	enum = page.Reason.PageBlockInterstitial
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"pageBlockInterstitial"` != string(result) {
		t.Errorf("Expected '\"pageBlockInterstitial\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"pageBlockInterstitial"`), &enum)
	if page.Reason.PageBlockInterstitial != enum {
		t.Errorf("Expcected %d, got %d", page.Reason.PageBlockInterstitial, enum)
	}

	enum = page.Reason.Reload
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"reload"` != string(result) {
		t.Errorf("Expected '\"reload\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"reload"`), &enum)
	if page.Reason.Reload != enum {
		t.Errorf("Expcected %d, got %d", page.Reason.Reload, enum)
	}
}
