package network

import (
	"encoding/json"
	"testing"
)

func TestEnumConnectionType(t *testing.T) {
	var enum ConnectionTypeEnum
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

	enum = ConnectionType.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if ConnectionType.None != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.None, enum)
	}

	enum = ConnectionType.Cellular2g
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cellular2g"` != string(result) {
		t.Errorf("Expected '\"cellular2g\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cellular2g"`), &enum)
	if ConnectionType.Cellular2g != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Cellular2g, enum)
	}

	enum = ConnectionType.Cellular3g
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cellular3g"` != string(result) {
		t.Errorf("Expected '\"cellular3g\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cellular3g"`), &enum)
	if ConnectionType.Cellular3g != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Cellular3g, enum)
	}
}

func TestEnumConnectionType2(t *testing.T) {
	var enum ConnectionTypeEnum
	var err error
	var result []byte

	enum = ConnectionType.Cellular4g
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cellular4g"` != string(result) {
		t.Errorf("Expected '\"cellular4g\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cellular4g"`), &enum)
	if ConnectionType.Cellular4g != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Cellular4g, enum)
	}

	enum = ConnectionType.Bluetooth
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"bluetooth"` != string(result) {
		t.Errorf("Expected '\"bluetooth\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"bluetooth"`), &enum)
	if ConnectionType.Bluetooth != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Bluetooth, enum)
	}

	enum = ConnectionType.Ethernet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ethernet"` != string(result) {
		t.Errorf("Expected '\"ethernet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ethernet"`), &enum)
	if ConnectionType.Ethernet != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Ethernet, enum)
	}

	enum = ConnectionType.Wifi
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"wifi"` != string(result) {
		t.Errorf("Expected '\"wifi\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"wifi"`), &enum)
	if ConnectionType.Wifi != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Wifi, enum)
	}
}

func TestEnumConnectionType3(t *testing.T) {
	var enum ConnectionTypeEnum
	var err error
	var result []byte

	enum = ConnectionType.Wimax
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"wimax"` != string(result) {
		t.Errorf("Expected '\"wimax\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"wimax"`), &enum)
	if ConnectionType.Wimax != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Wimax, enum)
	}

	enum = ConnectionType.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if ConnectionType.Other != enum {
		t.Errorf("Expcected %d, got %d", ConnectionType.Other, enum)
	}
}
