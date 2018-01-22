package emulation

import (
	"encoding/json"
	"testing"
)

func TestEnumOrientationType(t *testing.T) {
	var enum OrientationTypeEnum
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

	enum = OrientationType.PortraitPrimary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"portraitPrimary"` != string(result) {
		t.Errorf("Expected '\"portraitPrimary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"portraitPrimary"`), &enum)
	if OrientationType.PortraitPrimary != enum {
		t.Errorf("Expcected %d, got %d", OrientationType.PortraitPrimary, enum)
	}

	enum = OrientationType.PortraitSecondary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"portraitSecondary"` != string(result) {
		t.Errorf("Expected '\"portraitSecondary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"portraitSecondary"`), &enum)
	if OrientationType.PortraitSecondary != enum {
		t.Errorf("Expcected %d, got %d", OrientationType.PortraitSecondary, enum)
	}

	enum = OrientationType.LandscapePrimary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"landscapePrimary"` != string(result) {
		t.Errorf("Expected '\"landscapePrimary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"landscapePrimary"`), &enum)
	if OrientationType.LandscapePrimary != enum {
		t.Errorf("Expcected %d, got %d", OrientationType.LandscapePrimary, enum)
	}

	enum = OrientationType.LandscapeSecondary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"landscapeSecondary"` != string(result) {
		t.Errorf("Expected '\"landscapeSecondary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"landscapeSecondary"`), &enum)
	if OrientationType.LandscapeSecondary != enum {
		t.Errorf("Expcected %d, got %d", OrientationType.LandscapeSecondary, enum)
	}
}
