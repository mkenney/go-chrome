package network

import (
	"encoding/json"
	"testing"
)

func TestEnumErrorReason(t *testing.T) {
	var enum ErrorReasonEnum
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

	enum = ErrorReason.Failed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Failed"` != string(result) {
		t.Errorf("Expected '\"Failed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Failed"`), &enum)
	if ErrorReason.Failed != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.Failed, enum)
	}

	enum = ErrorReason.Aborted
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"Aborted"` != string(result) {
		t.Errorf("Expected '\"Aborted\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"Aborted"`), &enum)
	if ErrorReason.Aborted != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.Aborted, enum)
	}

	enum = ErrorReason.TimedOut
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"TimedOut"` != string(result) {
		t.Errorf("Expected '\"TimedOut\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"TimedOut"`), &enum)
	if ErrorReason.TimedOut != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.TimedOut, enum)
	}
}

func TestEnumErrorReason2(t *testing.T) {
	var enum ErrorReasonEnum
	var err error
	var result []byte

	enum = ErrorReason.AccessDenied
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"AccessDenied"` != string(result) {
		t.Errorf("Expected '\"AccessDenied\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"AccessDenied"`), &enum)
	if ErrorReason.AccessDenied != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.AccessDenied, enum)
	}

	enum = ErrorReason.ConnectionClosed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionClosed"` != string(result) {
		t.Errorf("Expected '\"ConnectionClosed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionClosed"`), &enum)
	if ErrorReason.ConnectionClosed != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.ConnectionClosed, enum)
	}

	enum = ErrorReason.ConnectionReset
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionReset"` != string(result) {
		t.Errorf("Expected '\"ConnectionReset\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionReset"`), &enum)
	if ErrorReason.ConnectionReset != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.ConnectionReset, enum)
	}

	enum = ErrorReason.ConnectionRefused
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionRefused"` != string(result) {
		t.Errorf("Expected '\"ConnectionRefused\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionRefused"`), &enum)
	if ErrorReason.ConnectionRefused != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.ConnectionRefused, enum)
	}
}

func TestEnumErrorReason3(t *testing.T) {
	var enum ErrorReasonEnum
	var err error
	var result []byte

	enum = ErrorReason.ConnectionAborted
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionAborted"` != string(result) {
		t.Errorf("Expected '\"ConnectionAborted\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionAborted"`), &enum)
	if ErrorReason.ConnectionAborted != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.ConnectionAborted, enum)
	}

	enum = ErrorReason.ConnectionFailed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"ConnectionFailed"` != string(result) {
		t.Errorf("Expected '\"ConnectionFailed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"ConnectionFailed"`), &enum)
	if ErrorReason.ConnectionFailed != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.ConnectionFailed, enum)
	}

	enum = ErrorReason.NameNotResolved
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"NameNotResolved"` != string(result) {
		t.Errorf("Expected '\"NameNotResolved\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"NameNotResolved"`), &enum)
	if ErrorReason.NameNotResolved != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.NameNotResolved, enum)
	}

	enum = ErrorReason.InternetDisconnected
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"InternetDisconnected"` != string(result) {
		t.Errorf("Expected '\"InternetDisconnected\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"InternetDisconnected"`), &enum)
	if ErrorReason.InternetDisconnected != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.InternetDisconnected, enum)
	}
}

func TestEnumErrorReason4(t *testing.T) {
	var enum ErrorReasonEnum
	var err error
	var result []byte

	enum = ErrorReason.AddressUnreachable
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"AddressUnreachable"` != string(result) {
		t.Errorf("Expected '\"AddressUnreachable\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"AddressUnreachable"`), &enum)
	if ErrorReason.AddressUnreachable != enum {
		t.Errorf("Expcected %d, got %d", ErrorReason.AddressUnreachable, enum)
	}
}
