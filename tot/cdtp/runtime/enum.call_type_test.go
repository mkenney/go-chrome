package runtime

import (
	"encoding/json"
	"testing"
)

func TestEnumCallType(t *testing.T) {
	var enum CallTypeEnum
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

	enum = CallType.Log
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"log"` != string(result) {
		t.Errorf("Expected '\"log\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"log"`), &enum)
	if CallType.Log != enum {
		t.Errorf("Expcected %d, got %d", CallType.Log, enum)
	}

	enum = CallType.Debug
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"debug"` != string(result) {
		t.Errorf("Expected '\"debug\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"debug"`), &enum)
	if CallType.Debug != enum {
		t.Errorf("Expcected %d, got %d", CallType.Debug, enum)
	}

	enum = CallType.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if CallType.Info != enum {
		t.Errorf("Expcected %d, got %d", CallType.Info, enum)
	}

	enum = CallType.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if CallType.Error != enum {
		t.Errorf("Expcected %d, got %d", CallType.Error, enum)
	}

	enum = CallType.Warning
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"warning"` != string(result) {
		t.Errorf("Expected '\"warning\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"warning"`), &enum)
	if CallType.Warning != enum {
		t.Errorf("Expcected %d, got %d", CallType.Warning, enum)
	}

	enum = CallType.Dir
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"dir"` != string(result) {
		t.Errorf("Expected '\"dir\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"dir"`), &enum)
	if CallType.Dir != enum {
		t.Errorf("Expcected %d, got %d", CallType.Dir, enum)
	}

	enum = CallType.Dirxml
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"dirxml"` != string(result) {
		t.Errorf("Expected '\"dirxml\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"dirxml"`), &enum)
	if CallType.Dirxml != enum {
		t.Errorf("Expcected %d, got %d", CallType.Dirxml, enum)
	}

	enum = CallType.Table
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"table"` != string(result) {
		t.Errorf("Expected '\"table\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"table"`), &enum)
	if CallType.Table != enum {
		t.Errorf("Expcected %d, got %d", CallType.Table, enum)
	}

	enum = CallType.Trace
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"trace"` != string(result) {
		t.Errorf("Expected '\"trace\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"trace"`), &enum)
	if CallType.Trace != enum {
		t.Errorf("Expcected %d, got %d", CallType.Trace, enum)
	}

	enum = CallType.Clear
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"clear"` != string(result) {
		t.Errorf("Expected '\"clear\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"clear"`), &enum)
	if CallType.Clear != enum {
		t.Errorf("Expcected %d, got %d", CallType.Clear, enum)
	}

	enum = CallType.StartGroup
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"startGroup"` != string(result) {
		t.Errorf("Expected '\"startGroup\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"startGroup"`), &enum)
	if CallType.StartGroup != enum {
		t.Errorf("Expcected %d, got %d", CallType.StartGroup, enum)
	}

	enum = CallType.StartGroupCollapsed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"startGroupCollapsed"` != string(result) {
		t.Errorf("Expected '\"startGroupCollapsed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"startGroupCollapsed"`), &enum)
	if CallType.StartGroupCollapsed != enum {
		t.Errorf("Expcected %d, got %d", CallType.StartGroupCollapsed, enum)
	}

	enum = CallType.EndGroup
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"endGroup"` != string(result) {
		t.Errorf("Expected '\"endGroup\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"endGroup"`), &enum)
	if CallType.EndGroup != enum {
		t.Errorf("Expcected %d, got %d", CallType.EndGroup, enum)
	}

	enum = CallType.Assert
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"assert"` != string(result) {
		t.Errorf("Expected '\"assert\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"assert"`), &enum)
	if CallType.Assert != enum {
		t.Errorf("Expcected %d, got %d", CallType.Assert, enum)
	}

	enum = CallType.Profile
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"profile"` != string(result) {
		t.Errorf("Expected '\"profile\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"profile"`), &enum)
	if CallType.Profile != enum {
		t.Errorf("Expcected %d, got %d", CallType.Profile, enum)
	}

	enum = CallType.ProfileEnd
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"profileEnd"` != string(result) {
		t.Errorf("Expected '\"profileEnd\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"profileEnd"`), &enum)
	if CallType.ProfileEnd != enum {
		t.Errorf("Expcected %d, got %d", CallType.ProfileEnd, enum)
	}

	enum = CallType.Count
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"count"` != string(result) {
		t.Errorf("Expected '\"count\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"count"`), &enum)
	if CallType.Count != enum {
		t.Errorf("Expcected %d, got %d", CallType.Count, enum)
	}

	enum = CallType.TimeEnd
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"timeEnd"` != string(result) {
		t.Errorf("Expected '\"timeEnd\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"timeEnd"`), &enum)
	if CallType.TimeEnd != enum {
		t.Errorf("Expcected %d, got %d", CallType.TimeEnd, enum)
	}
}
