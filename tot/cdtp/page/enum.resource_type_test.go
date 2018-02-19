package page

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/page"
)

func TestEnumResourceType(t *testing.T) {
	var enum page.ResourceTypeEnum
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

	enum = page.ResourceType.Document
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Document"` != string(result) {
		t.Errorf("Expected '\"Document\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Document"`), &enum)
	if page.ResourceType.Document != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Document, enum)
	}

	enum = page.ResourceType.Stylesheet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Stylesheet"` != string(result) {
		t.Errorf("Expected '\"Stylesheet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Stylesheet"`), &enum)
	if page.ResourceType.Stylesheet != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Stylesheet, enum)
	}

	enum = page.ResourceType.Image
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Image"` != string(result) {
		t.Errorf("Expected '\"Image\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Image"`), &enum)
	if page.ResourceType.Image != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Image, enum)
	}
}

func TestEnumResourceType2(t *testing.T) {
	var enum page.ResourceTypeEnum
	var err error
	var result []byte

	enum = page.ResourceType.Media
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Media"` != string(result) {
		t.Errorf("Expected '\"Media\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Media"`), &enum)
	if page.ResourceType.Media != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Media, enum)
	}

	enum = page.ResourceType.Font
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Font"` != string(result) {
		t.Errorf("Expected '\"Font\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Font"`), &enum)
	if page.ResourceType.Font != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Font, enum)
	}

	enum = page.ResourceType.Script
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Script"` != string(result) {
		t.Errorf("Expected '\"Script\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Script"`), &enum)
	if page.ResourceType.Script != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Script, enum)
	}

	enum = page.ResourceType.TextTrack
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"TextTrack"` != string(result) {
		t.Errorf("Expected '\"TextTrack\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"TextTrack"`), &enum)
	if page.ResourceType.TextTrack != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.TextTrack, enum)
	}
}

func TestEnumResourceType3(t *testing.T) {
	var enum page.ResourceTypeEnum
	var err error
	var result []byte

	enum = page.ResourceType.XHR
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"XHR"` != string(result) {
		t.Errorf("Expected '\"XHR\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"XHR"`), &enum)
	if page.ResourceType.XHR != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.XHR, enum)
	}

	enum = page.ResourceType.Fetch
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Fetch"` != string(result) {
		t.Errorf("Expected '\"Fetch\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Fetch"`), &enum)
	if page.ResourceType.Fetch != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Fetch, enum)
	}

	enum = page.ResourceType.EventSource
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"EventSource"` != string(result) {
		t.Errorf("Expected '\"EventSource\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"EventSource"`), &enum)
	if page.ResourceType.EventSource != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.EventSource, enum)
	}

	enum = page.ResourceType.WebSocket
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"WebSocket"` != string(result) {
		t.Errorf("Expected '\"WebSocket\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"WebSocket"`), &enum)
	if page.ResourceType.WebSocket != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.WebSocket, enum)
	}
}

func TestEnumResourceType4(t *testing.T) {
	var enum page.ResourceTypeEnum
	var err error
	var result []byte

	enum = page.ResourceType.Manifest
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Manifest"` != string(result) {
		t.Errorf("Expected '\"Manifest\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Manifest"`), &enum)
	if page.ResourceType.Manifest != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Manifest, enum)
	}

	enum = page.ResourceType.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Other"` != string(result) {
		t.Errorf("Expected '\"Other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Other"`), &enum)
	if page.ResourceType.Other != enum {
		t.Errorf("Expcected %d, got %d", page.ResourceType.Other, enum)
	}
}
