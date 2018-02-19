package network_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumReferrerPolicy(t *testing.T) {
	var enum network.ReferrerPolicyEnum
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

	enum = network.ReferrerPolicy.UnsafeURL
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"unsafe-url"` != string(result) {
		t.Errorf("Expected '\"unsafe-url\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"unsafe-url"`), &enum)
	if network.ReferrerPolicy.UnsafeURL != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.UnsafeURL, enum)
	}

	enum = network.ReferrerPolicy.NoReferrerWhenDowngrade
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"no-referrer-when-downgrade"` != string(result) {
		t.Errorf("Expected '\"no-referrer-when-downgrade\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"no-referrer-when-downgrade"`), &enum)
	if network.ReferrerPolicy.NoReferrerWhenDowngrade != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.NoReferrerWhenDowngrade, enum)
	}

	enum = network.ReferrerPolicy.NoReferrer
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"no-referrer"` != string(result) {
		t.Errorf("Expected '\"no-referrer\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"no-referrer"`), &enum)
	if network.ReferrerPolicy.NoReferrer != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.NoReferrer, enum)
	}
}

func TestEnumReferrerPolicy2(t *testing.T) {
	var enum network.ReferrerPolicyEnum
	var err error
	var result []byte

	enum = network.ReferrerPolicy.Origin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"origin"` != string(result) {
		t.Errorf("Expected '\"origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"origin"`), &enum)
	if network.ReferrerPolicy.Origin != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.Origin, enum)
	}

	enum = network.ReferrerPolicy.OriginWhenCrossOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"origin-when-cross-origin"` != string(result) {
		t.Errorf("Expected '\"origin-when-cross-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"origin-when-cross-origin"`), &enum)
	if network.ReferrerPolicy.OriginWhenCrossOrigin != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.OriginWhenCrossOrigin, enum)
	}

	enum = network.ReferrerPolicy.SameOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"same-origin"` != string(result) {
		t.Errorf("Expected '\"same-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"same-origin"`), &enum)
	if network.ReferrerPolicy.SameOrigin != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.SameOrigin, enum)
	}

	enum = network.ReferrerPolicy.StrictOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"strict-origin"` != string(result) {
		t.Errorf("Expected '\"strict-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"strict-origin"`), &enum)
	if network.ReferrerPolicy.StrictOrigin != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.StrictOrigin, enum)
	}
}

func TestEnumReferrerPolicy3(t *testing.T) {
	var enum network.ReferrerPolicyEnum
	var err error
	var result []byte

	enum = network.ReferrerPolicy.StrictOriginWhenCrossOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"strict-origin-when-cross-origin"` != string(result) {
		t.Errorf("Expected '\"strict-origin-when-cross-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"strict-origin-when-cross-origin"`), &enum)
	if network.ReferrerPolicy.StrictOriginWhenCrossOrigin != enum {
		t.Errorf("Expcected %d, got %d", network.ReferrerPolicy.StrictOriginWhenCrossOrigin, enum)
	}
}
