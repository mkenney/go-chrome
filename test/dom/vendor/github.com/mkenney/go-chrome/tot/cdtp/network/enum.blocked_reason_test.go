package network

import (
	"encoding/json"
	"testing"
)

func TestEnumBlockedReason(t *testing.T) {
	var enum BlockedReasonEnum
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

	enum = BlockedReason.Csp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"csp"` != string(result) {
		t.Errorf("Expected '\"csp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"csp"`), &enum)
	if BlockedReason.Csp != enum {
		t.Errorf("Expcected %d, got %d", BlockedReason.Csp, enum)
	}

	enum = BlockedReason.MixedContent
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mixed-content"` != string(result) {
		t.Errorf("Expected '\"mixed-content\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mixed-content"`), &enum)
	if BlockedReason.MixedContent != enum {
		t.Errorf("Expcected %d, got %d", BlockedReason.MixedContent, enum)
	}

	enum = BlockedReason.Origin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"origin"` != string(result) {
		t.Errorf("Expected '\"origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"origin"`), &enum)
	if BlockedReason.Origin != enum {
		t.Errorf("Expcected %d, got %d", BlockedReason.Origin, enum)
	}
}

func TestEnumBlockedReason2(t *testing.T) {
	var enum BlockedReasonEnum
	var err error
	var result []byte

	enum = BlockedReason.Inspector
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"inspector"` != string(result) {
		t.Errorf("Expected '\"inspector\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"inspector"`), &enum)
	if BlockedReason.Inspector != enum {
		t.Errorf("Expcected %d, got %d", BlockedReason.Inspector, enum)
	}

	enum = BlockedReason.SubresourceFilter
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"subresource-filter"` != string(result) {
		t.Errorf("Expected '\"subresource-filter\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"subresource-filter"`), &enum)
	if BlockedReason.SubresourceFilter != enum {
		t.Errorf("Expcected %d, got %d", BlockedReason.SubresourceFilter, enum)
	}

	enum = BlockedReason.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if BlockedReason.Other != enum {
		t.Errorf("Expcected %d, got %d", BlockedReason.Other, enum)
	}
}
