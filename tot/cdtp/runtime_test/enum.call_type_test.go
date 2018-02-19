package runtime_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestEnumCallType(t *testing.T) {
	var enum runtime.CallTypeEnum
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

	enum = runtime.CallType.Log
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"log"` != string(result) {
		t.Errorf("Expected '\"log\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"log"`), &enum)
	if runtime.CallType.Log != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Log, enum)
	}

	enum = runtime.CallType.Debug
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"debug"` != string(result) {
		t.Errorf("Expected '\"debug\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"debug"`), &enum)
	if runtime.CallType.Debug != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Debug, enum)
	}

	enum = runtime.CallType.Info
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"info"` != string(result) {
		t.Errorf("Expected '\"info\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"info"`), &enum)
	if runtime.CallType.Info != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Info, enum)
	}
}

func TestEnumCallType2(t *testing.T) {
	var enum runtime.CallTypeEnum
	var err error
	var result []byte

	enum = runtime.CallType.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if runtime.CallType.Error != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Error, enum)
	}

	enum = runtime.CallType.Warning
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"warning"` != string(result) {
		t.Errorf("Expected '\"warning\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"warning"`), &enum)
	if runtime.CallType.Warning != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Warning, enum)
	}

	enum = runtime.CallType.Dir
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"dir"` != string(result) {
		t.Errorf("Expected '\"dir\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"dir"`), &enum)
	if runtime.CallType.Dir != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Dir, enum)
	}

	enum = runtime.CallType.Dirxml
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"dirxml"` != string(result) {
		t.Errorf("Expected '\"dirxml\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"dirxml"`), &enum)
	if runtime.CallType.Dirxml != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Dirxml, enum)
	}
}

func TestEnumCallType3(t *testing.T) {
	var enum runtime.CallTypeEnum
	var err error
	var result []byte

	enum = runtime.CallType.Table
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"table"` != string(result) {
		t.Errorf("Expected '\"table\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"table"`), &enum)
	if runtime.CallType.Table != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Table, enum)
	}

	enum = runtime.CallType.Trace
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"trace"` != string(result) {
		t.Errorf("Expected '\"trace\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"trace"`), &enum)
	if runtime.CallType.Trace != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Trace, enum)
	}

	enum = runtime.CallType.Clear
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"clear"` != string(result) {
		t.Errorf("Expected '\"clear\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"clear"`), &enum)
	if runtime.CallType.Clear != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Clear, enum)
	}

	enum = runtime.CallType.StartGroup
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"startGroup"` != string(result) {
		t.Errorf("Expected '\"startGroup\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"startGroup"`), &enum)
	if runtime.CallType.StartGroup != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.StartGroup, enum)
	}
}

func TestEnumCallType4(t *testing.T) {
	var enum runtime.CallTypeEnum
	var err error
	var result []byte

	enum = runtime.CallType.StartGroupCollapsed
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"startGroupCollapsed"` != string(result) {
		t.Errorf("Expected '\"startGroupCollapsed\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"startGroupCollapsed"`), &enum)
	if runtime.CallType.StartGroupCollapsed != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.StartGroupCollapsed, enum)
	}

	enum = runtime.CallType.EndGroup
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"endGroup"` != string(result) {
		t.Errorf("Expected '\"endGroup\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"endGroup"`), &enum)
	if runtime.CallType.EndGroup != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.EndGroup, enum)
	}

	enum = runtime.CallType.Assert
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"assert"` != string(result) {
		t.Errorf("Expected '\"assert\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"assert"`), &enum)
	if runtime.CallType.Assert != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Assert, enum)
	}

	enum = runtime.CallType.Profile
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"profile"` != string(result) {
		t.Errorf("Expected '\"profile\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"profile"`), &enum)
	if runtime.CallType.Profile != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Profile, enum)
	}
}

func TestEnumCallType5(t *testing.T) {
	var enum runtime.CallTypeEnum
	var err error
	var result []byte

	enum = runtime.CallType.ProfileEnd
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"profileEnd"` != string(result) {
		t.Errorf("Expected '\"profileEnd\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"profileEnd"`), &enum)
	if runtime.CallType.ProfileEnd != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.ProfileEnd, enum)
	}

	enum = runtime.CallType.Count
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"count"` != string(result) {
		t.Errorf("Expected '\"count\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"count"`), &enum)
	if runtime.CallType.Count != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.Count, enum)
	}

	enum = runtime.CallType.TimeEnd
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"timeEnd"` != string(result) {
		t.Errorf("Expected '\"timeEnd\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"timeEnd"`), &enum)
	if runtime.CallType.TimeEnd != enum {
		t.Errorf("Expcected %d, got %d", runtime.CallType.TimeEnd, enum)
	}
}
