package network

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumCookieSameSite(t *testing.T) {
	var enum network.CookieSameSiteEnum
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

	enum = network.CookieSameSite.Strict
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Strict"` != string(result) {
		t.Errorf("Expected '\"Strict\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Strict"`), &enum)
	if network.CookieSameSite.Strict != enum {
		t.Errorf("Expcected %d, got %d", network.CookieSameSite.Strict, enum)
	}

	enum = network.CookieSameSite.Lax
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Lax"` != string(result) {
		t.Errorf("Expected '\"Lax\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Lax"`), &enum)
	if network.CookieSameSite.Lax != enum {
		t.Errorf("Expcected %d, got %d", network.CookieSameSite.Lax, enum)
	}
}
