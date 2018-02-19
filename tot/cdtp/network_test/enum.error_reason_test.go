package network_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/network"
)

func TestEnumErrorReason(t *testing.T) {
	var enum network.ErrorReasonEnum
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

	enum = network.ErrorReason.Failed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Failed"` != string(result) {
		t.Errorf("Expected '\"Failed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Failed"`), &enum)
	if network.ErrorReason.Failed != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.Failed, enum)
	}

	enum = network.ErrorReason.Aborted
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Aborted"` != string(result) {
		t.Errorf("Expected '\"Aborted\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Aborted"`), &enum)
	if network.ErrorReason.Aborted != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.Aborted, enum)
	}

	enum = network.ErrorReason.TimedOut
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"TimedOut"` != string(result) {
		t.Errorf("Expected '\"TimedOut\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"TimedOut"`), &enum)
	if network.ErrorReason.TimedOut != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.TimedOut, enum)
	}
}

func TestEnumErrorReason2(t *testing.T) {
	var enum network.ErrorReasonEnum
	var err error
	var result []byte

	enum = network.ErrorReason.AccessDenied
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"AccessDenied"` != string(result) {
		t.Errorf("Expected '\"AccessDenied\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"AccessDenied"`), &enum)
	if network.ErrorReason.AccessDenied != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.AccessDenied, enum)
	}

	enum = network.ErrorReason.ConnectionClosed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionClosed"` != string(result) {
		t.Errorf("Expected '\"ConnectionClosed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionClosed"`), &enum)
	if network.ErrorReason.ConnectionClosed != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.ConnectionClosed, enum)
	}

	enum = network.ErrorReason.ConnectionReset
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionReset"` != string(result) {
		t.Errorf("Expected '\"ConnectionReset\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionReset"`), &enum)
	if network.ErrorReason.ConnectionReset != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.ConnectionReset, enum)
	}

	enum = network.ErrorReason.ConnectionRefused
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionRefused"` != string(result) {
		t.Errorf("Expected '\"ConnectionRefused\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionRefused"`), &enum)
	if network.ErrorReason.ConnectionRefused != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.ConnectionRefused, enum)
	}
}

func TestEnumErrorReason3(t *testing.T) {
	var enum network.ErrorReasonEnum
	var err error
	var result []byte

	enum = network.ErrorReason.ConnectionAborted
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionAborted"` != string(result) {
		t.Errorf("Expected '\"ConnectionAborted\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionAborted"`), &enum)
	if network.ErrorReason.ConnectionAborted != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.ConnectionAborted, enum)
	}

	enum = network.ErrorReason.ConnectionFailed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionFailed"` != string(result) {
		t.Errorf("Expected '\"ConnectionFailed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionFailed"`), &enum)
	if network.ErrorReason.ConnectionFailed != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.ConnectionFailed, enum)
	}

	enum = network.ErrorReason.NameNotResolved
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"NameNotResolved"` != string(result) {
		t.Errorf("Expected '\"NameNotResolved\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"NameNotResolved"`), &enum)
	if network.ErrorReason.NameNotResolved != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.NameNotResolved, enum)
	}

	enum = network.ErrorReason.InternetDisconnected
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"InternetDisconnected"` != string(result) {
		t.Errorf("Expected '\"InternetDisconnected\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"InternetDisconnected"`), &enum)
	if network.ErrorReason.InternetDisconnected != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.InternetDisconnected, enum)
	}
}

func TestEnumErrorReason4(t *testing.T) {
	var enum network.ErrorReasonEnum
	var err error
	var result []byte

	enum = network.ErrorReason.AddressUnreachable
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"AddressUnreachable"` != string(result) {
		t.Errorf("Expected '\"AddressUnreachable\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"AddressUnreachable"`), &enum)
	if network.ErrorReason.AddressUnreachable != enum {
		t.Errorf("Expcected %d, got %d", network.ErrorReason.AddressUnreachable, enum)
	}
}
