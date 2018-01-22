package network

import (
	"encoding/json"
	"testing"
)

func TestEnumChallengeResponse(t *testing.T) {
	var enum ChallengeResponseEnum
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

	enum = ChallengeResponse.Default
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Default"` != string(result) {
		t.Errorf("Expected '\"Default\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Default"`), &enum)
	if ChallengeResponse.Default != enum {
		t.Errorf("Expcected %d, got %d", ChallengeResponse.Default, enum)
	}

	enum = ChallengeResponse.CancelAuth
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CancelAuth"` != string(result) {
		t.Errorf("Expected '\"CancelAuth\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CancelAuth"`), &enum)
	if ChallengeResponse.CancelAuth != enum {
		t.Errorf("Expcected %d, got %d", ChallengeResponse.CancelAuth, enum)
	}

	enum = ChallengeResponse.ProvideCredentials
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ProvideCredentials"` != string(result) {
		t.Errorf("Expected '\"ProvideCredentials\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ProvideCredentials"`), &enum)
	if ChallengeResponse.ProvideCredentials != enum {
		t.Errorf("Expcected %d, got %d", ChallengeResponse.ProvideCredentials, enum)
	}
}
