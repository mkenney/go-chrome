package network

import (
	"encoding/json"
	"testing"
)

func TestEnumCookieSameSite(t *testing.T) {
	var enum CookieSameSiteEnum
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

	enum = CookieSameSite.Strict
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Strict"` != string(result) {
		t.Errorf("Expected '\"Strict\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Strict"`), &enum)
	if CookieSameSite.Strict != enum {
		t.Errorf("Expcected %d, got %d", CookieSameSite.Strict, enum)
	}

	enum = CookieSameSite.Lax
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Lax"` != string(result) {
		t.Errorf("Expected '\"Lax\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Lax"`), &enum)
	if CookieSameSite.Lax != enum {
		t.Errorf("Expcected %d, got %d", CookieSameSite.Lax, enum)
	}
}
