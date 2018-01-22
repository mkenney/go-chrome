package network

import (
	"encoding/json"
	"testing"
)

func TestEnumReferrerPolicy(t *testing.T) {
	var enum ReferrerPolicyEnum
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

	enum = ReferrerPolicy.UnsafeURL
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"unsafe-url"` != string(result) {
		t.Errorf("Expected '\"unsafe-url\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"unsafe-url"`), &enum)
	if ReferrerPolicy.UnsafeURL != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.UnsafeURL, enum)
	}

	enum = ReferrerPolicy.NoReferrerWhenDowngrade
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"no-referrer-when-downgrade"` != string(result) {
		t.Errorf("Expected '\"no-referrer-when-downgrade\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"no-referrer-when-downgrade"`), &enum)
	if ReferrerPolicy.NoReferrerWhenDowngrade != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.NoReferrerWhenDowngrade, enum)
	}

	enum = ReferrerPolicy.NoReferrer
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"no-referrer"` != string(result) {
		t.Errorf("Expected '\"no-referrer\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"no-referrer"`), &enum)
	if ReferrerPolicy.NoReferrer != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.NoReferrer, enum)
	}
}

func TestEnumReferrerPolicy2(t *testing.T) {
	var enum ReferrerPolicyEnum
	var err error
	var result []byte

	enum = ReferrerPolicy.Origin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"origin"` != string(result) {
		t.Errorf("Expected '\"origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"origin"`), &enum)
	if ReferrerPolicy.Origin != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.Origin, enum)
	}

	enum = ReferrerPolicy.OriginWhenCrossOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"origin-when-cross-origin"` != string(result) {
		t.Errorf("Expected '\"origin-when-cross-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"origin-when-cross-origin"`), &enum)
	if ReferrerPolicy.OriginWhenCrossOrigin != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.OriginWhenCrossOrigin, enum)
	}

	enum = ReferrerPolicy.SameOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"same-origin"` != string(result) {
		t.Errorf("Expected '\"same-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"same-origin"`), &enum)
	if ReferrerPolicy.SameOrigin != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.SameOrigin, enum)
	}

	enum = ReferrerPolicy.StrictOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"strict-origin"` != string(result) {
		t.Errorf("Expected '\"strict-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"strict-origin"`), &enum)
	if ReferrerPolicy.StrictOrigin != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.StrictOrigin, enum)
	}
}

func TestEnumReferrerPolicy3(t *testing.T) {
	var enum ReferrerPolicyEnum
	var err error
	var result []byte

	enum = ReferrerPolicy.StrictOriginWhenCrossOrigin
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"strict-origin-when-cross-origin"` != string(result) {
		t.Errorf("Expected '\"strict-origin-when-cross-origin\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"strict-origin-when-cross-origin"`), &enum)
	if ReferrerPolicy.StrictOriginWhenCrossOrigin != enum {
		t.Errorf("Expcected %d, got %d", ReferrerPolicy.StrictOriginWhenCrossOrigin, enum)
	}
}
