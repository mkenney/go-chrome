package page

import (
	"encoding/json"
	"testing"
)

func TestEnumReason(t *testing.T) {
	var enum ReasonEnum
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

	enum = Reason.FormSubmissionGet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"formSubmissionGet"` != string(result) {
		t.Errorf("Expected '\"formSubmissionGet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"formSubmissionGet"`), &enum)
	if Reason.FormSubmissionGet != enum {
		t.Errorf("Expcected %d, got %d", Reason.FormSubmissionGet, enum)
	}

	enum = Reason.FormSubmissionPost
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"formSubmissionPost"` != string(result) {
		t.Errorf("Expected '\"formSubmissionPost\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"formSubmissionPost"`), &enum)
	if Reason.FormSubmissionPost != enum {
		t.Errorf("Expcected %d, got %d", Reason.FormSubmissionPost, enum)
	}

	enum = Reason.HTTPHeaderRefresh
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"httpHeaderRefresh"` != string(result) {
		t.Errorf("Expected '\"httpHeaderRefresh\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"httpHeaderRefresh"`), &enum)
	if Reason.HTTPHeaderRefresh != enum {
		t.Errorf("Expcected %d, got %d", Reason.HTTPHeaderRefresh, enum)
	}
}
func TestEnumReason2(t *testing.T) {
	var enum ReasonEnum
	var err error
	var result []byte

	enum = Reason.ScriptInitiated
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"scriptInitiated"` != string(result) {
		t.Errorf("Expected '\"scriptInitiated\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"scriptInitiated"`), &enum)
	if Reason.ScriptInitiated != enum {
		t.Errorf("Expcected %d, got %d", Reason.ScriptInitiated, enum)
	}

	enum = Reason.MetaTagRefresh
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"metaTagRefresh"` != string(result) {
		t.Errorf("Expected '\"metaTagRefresh\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"metaTagRefresh"`), &enum)
	if Reason.MetaTagRefresh != enum {
		t.Errorf("Expcected %d, got %d", Reason.MetaTagRefresh, enum)
	}

	enum = Reason.PageBlockInterstitial
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"pageBlockInterstitial"` != string(result) {
		t.Errorf("Expected '\"pageBlockInterstitial\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"pageBlockInterstitial"`), &enum)
	if Reason.PageBlockInterstitial != enum {
		t.Errorf("Expcected %d, got %d", Reason.PageBlockInterstitial, enum)
	}

	enum = Reason.Reload
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"reload"` != string(result) {
		t.Errorf("Expected '\"reload\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"reload"`), &enum)
	if Reason.Reload != enum {
		t.Errorf("Expcected %d, got %d", Reason.Reload, enum)
	}
}
