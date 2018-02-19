package network_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumChallengeResponse(t *testing.T) {
	var enum network.ChallengeResponseEnum
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

	enum = network.ChallengeResponse.Default
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Default"` != string(result) {
		t.Errorf("Expected '\"Default\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Default"`), &enum)
	if network.ChallengeResponse.Default != enum {
		t.Errorf("Expcected %d, got %d", network.ChallengeResponse.Default, enum)
	}

	enum = network.ChallengeResponse.CancelAuth
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"CancelAuth"` != string(result) {
		t.Errorf("Expected '\"CancelAuth\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"CancelAuth"`), &enum)
	if network.ChallengeResponse.CancelAuth != enum {
		t.Errorf("Expcected %d, got %d", network.ChallengeResponse.CancelAuth, enum)
	}

	enum = network.ChallengeResponse.ProvideCredentials
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ProvideCredentials"` != string(result) {
		t.Errorf("Expected '\"ProvideCredentials\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ProvideCredentials"`), &enum)
	if network.ChallengeResponse.ProvideCredentials != enum {
		t.Errorf("Expcected %d, got %d", network.ChallengeResponse.ProvideCredentials, enum)
	}
}
