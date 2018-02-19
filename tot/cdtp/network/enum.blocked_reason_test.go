package network

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumBlockedReason(t *testing.T) {
	var enum network.BlockedReasonEnum
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

	enum = network.BlockedReason.Csp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"csp"` != string(result) {
		t.Errorf("Expected '\"csp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"csp"`), &enum)
	if network.BlockedReason.Csp != enum {
		t.Errorf("Expcected %d, got %d", network.BlockedReason.Csp, enum)
	}

	enum = network.BlockedReason.MixedContent
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"mixed-content"` != string(result) {
		t.Errorf("Expected '\"mixed-content\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"mixed-content"`), &enum)
	if network.BlockedReason.MixedContent != enum {
		t.Errorf("Expcected %d, got %d", network.BlockedReason.MixedContent, enum)
	}

	enum = network.BlockedReason.Origin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"origin"` != string(result) {
		t.Errorf("Expected '\"origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"origin"`), &enum)
	if network.BlockedReason.Origin != enum {
		t.Errorf("Expcected %d, got %d", network.BlockedReason.Origin, enum)
	}
}

func TestEnumBlockedReason2(t *testing.T) {
	var enum network.BlockedReasonEnum
	var err error
	var result []byte

	enum = network.BlockedReason.Inspector
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"inspector"` != string(result) {
		t.Errorf("Expected '\"inspector\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"inspector"`), &enum)
	if network.BlockedReason.Inspector != enum {
		t.Errorf("Expcected %d, got %d", network.BlockedReason.Inspector, enum)
	}

	enum = network.BlockedReason.SubresourceFilter
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"subresource-filter"` != string(result) {
		t.Errorf("Expected '\"subresource-filter\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"subresource-filter"`), &enum)
	if network.BlockedReason.SubresourceFilter != enum {
		t.Errorf("Expcected %d, got %d", network.BlockedReason.SubresourceFilter, enum)
	}

	enum = network.BlockedReason.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if network.BlockedReason.Other != enum {
		t.Errorf("Expcected %d, got %d", network.BlockedReason.Other, enum)
	}
}
