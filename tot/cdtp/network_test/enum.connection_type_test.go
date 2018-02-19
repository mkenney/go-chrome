package network_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumConnectionType(t *testing.T) {
	var enum network.ConnectionTypeEnum
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

	enum = network.ConnectionType.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if network.ConnectionType.None != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.None, enum)
	}

	enum = network.ConnectionType.Cellular2g
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cellular2g"` != string(result) {
		t.Errorf("Expected '\"cellular2g\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cellular2g"`), &enum)
	if network.ConnectionType.Cellular2g != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Cellular2g, enum)
	}

	enum = network.ConnectionType.Cellular3g
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cellular3g"` != string(result) {
		t.Errorf("Expected '\"cellular3g\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cellular3g"`), &enum)
	if network.ConnectionType.Cellular3g != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Cellular3g, enum)
	}
}

func TestEnumConnectionType2(t *testing.T) {
	var enum network.ConnectionTypeEnum
	var err error
	var result []byte

	enum = network.ConnectionType.Cellular4g
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"cellular4g"` != string(result) {
		t.Errorf("Expected '\"cellular4g\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"cellular4g"`), &enum)
	if network.ConnectionType.Cellular4g != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Cellular4g, enum)
	}

	enum = network.ConnectionType.Bluetooth
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"bluetooth"` != string(result) {
		t.Errorf("Expected '\"bluetooth\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"bluetooth"`), &enum)
	if network.ConnectionType.Bluetooth != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Bluetooth, enum)
	}

	enum = network.ConnectionType.Ethernet
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ethernet"` != string(result) {
		t.Errorf("Expected '\"ethernet\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ethernet"`), &enum)
	if network.ConnectionType.Ethernet != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Ethernet, enum)
	}

	enum = network.ConnectionType.Wifi
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"wifi"` != string(result) {
		t.Errorf("Expected '\"wifi\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"wifi"`), &enum)
	if network.ConnectionType.Wifi != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Wifi, enum)
	}
}

func TestEnumConnectionType3(t *testing.T) {
	var enum network.ConnectionTypeEnum
	var err error
	var result []byte

	enum = network.ConnectionType.Wimax
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"wimax"` != string(result) {
		t.Errorf("Expected '\"wimax\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"wimax"`), &enum)
	if network.ConnectionType.Wimax != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Wimax, enum)
	}

	enum = network.ConnectionType.Other
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"other"` != string(result) {
		t.Errorf("Expected '\"other\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"other"`), &enum)
	if network.ConnectionType.Other != enum {
		t.Errorf("Expcected %d, got %d", network.ConnectionType.Other, enum)
	}
}
