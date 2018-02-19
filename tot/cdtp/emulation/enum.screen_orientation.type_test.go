package emulation

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/emulation"
)

func TestEnumOrientationType(t *testing.T) {
	var enum emulation.OrientationTypeEnum
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

	enum = emulation.OrientationType.PortraitPrimary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"portraitPrimary"` != string(result) {
		t.Errorf("Expected '\"portraitPrimary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"portraitPrimary"`), &enum)
	if emulation.OrientationType.PortraitPrimary != enum {
		t.Errorf("Expcected %d, got %d", emulation.OrientationType.PortraitPrimary, enum)
	}

	enum = emulation.OrientationType.PortraitSecondary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"portraitSecondary"` != string(result) {
		t.Errorf("Expected '\"portraitSecondary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"portraitSecondary"`), &enum)
	if emulation.OrientationType.PortraitSecondary != enum {
		t.Errorf("Expcected %d, got %d", emulation.OrientationType.PortraitSecondary, enum)
	}

	enum = emulation.OrientationType.LandscapePrimary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"landscapePrimary"` != string(result) {
		t.Errorf("Expected '\"landscapePrimary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"landscapePrimary"`), &enum)
	if emulation.OrientationType.LandscapePrimary != enum {
		t.Errorf("Expcected %d, got %d", emulation.OrientationType.LandscapePrimary, enum)
	}
}

func TestEnumOrientationType2(t *testing.T) {
	var enum emulation.OrientationTypeEnum
	var err error
	var result []byte

	enum = emulation.OrientationType.LandscapeSecondary
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"landscapeSecondary"` != string(result) {
		t.Errorf("Expected '\"landscapeSecondary\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"landscapeSecondary"`), &enum)
	if emulation.OrientationType.LandscapeSecondary != enum {
		t.Errorf("Expcected %d, got %d", emulation.OrientationType.LandscapeSecondary, enum)
	}
}
