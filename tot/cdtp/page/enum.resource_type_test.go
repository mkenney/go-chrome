package page

import (
	"encoding/json"
	"testing"
)

func TestEnumResourceType(t *testing.T) {
	var enum ResourceTypeEnum
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

	enum = ResourceType.Document
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Document"` != string(result) {
		t.Errorf("Expected '\"Document\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Document"`), &enum)
	if ResourceType.Document != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Document, enum)
	}

	enum = ResourceType.Stylesheet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Stylesheet"` != string(result) {
		t.Errorf("Expected '\"Stylesheet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Stylesheet"`), &enum)
	if ResourceType.Stylesheet != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Stylesheet, enum)
	}

	enum = ResourceType.Image
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Image"` != string(result) {
		t.Errorf("Expected '\"Image\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Image"`), &enum)
	if ResourceType.Image != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Image, enum)
	}

	enum = ResourceType.Media
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Media"` != string(result) {
		t.Errorf("Expected '\"Media\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Media"`), &enum)
	if ResourceType.Media != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Media, enum)
	}

	enum = ResourceType.Font
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Font"` != string(result) {
		t.Errorf("Expected '\"Font\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Font"`), &enum)
	if ResourceType.Font != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Font, enum)
	}

	enum = ResourceType.Script
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Script"` != string(result) {
		t.Errorf("Expected '\"Script\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Script"`), &enum)
	if ResourceType.Script != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Script, enum)
	}

	enum = ResourceType.TextTrack
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"TextTrack"` != string(result) {
		t.Errorf("Expected '\"TextTrack\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"TextTrack"`), &enum)
	if ResourceType.TextTrack != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.TextTrack, enum)
	}

	enum = ResourceType.XHR
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"XHR"` != string(result) {
		t.Errorf("Expected '\"XHR\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"XHR"`), &enum)
	if ResourceType.XHR != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.XHR, enum)
	}

	enum = ResourceType.Fetch
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Fetch"` != string(result) {
		t.Errorf("Expected '\"Fetch\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Fetch"`), &enum)
	if ResourceType.Fetch != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Fetch, enum)
	}

	enum = ResourceType.EventSource
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"EventSource"` != string(result) {
		t.Errorf("Expected '\"EventSource\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"EventSource"`), &enum)
	if ResourceType.EventSource != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.EventSource, enum)
	}

	enum = ResourceType.WebSocket
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"WebSocket"` != string(result) {
		t.Errorf("Expected '\"WebSocket\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"WebSocket"`), &enum)
	if ResourceType.WebSocket != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.WebSocket, enum)
	}

	enum = ResourceType.Manifest
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Manifest"` != string(result) {
		t.Errorf("Expected '\"Manifest\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Manifest"`), &enum)
	if ResourceType.Manifest != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Manifest, enum)
	}

	enum = ResourceType.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Other"` != string(result) {
		t.Errorf("Expected '\"Other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Other"`), &enum)
	if ResourceType.Other != enum {
		t.Errorf("Expcected %d, got %d", ResourceType.Other, enum)
	}
}
