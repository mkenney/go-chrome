package runtime

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestEnumObjectSubtype(t *testing.T) {
	var enum runtime.ObjectSubtypeEnum
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

	enum = runtime.ObjectSubtype.Array
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"array"` != string(result) {
		t.Errorf("Expected '\"array\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"array"`), &enum)
	if runtime.ObjectSubtype.Array != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Array, enum)
	}

	enum = runtime.ObjectSubtype.Null
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"null"` != string(result) {
		t.Errorf("Expected '\"null\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"null"`), &enum)
	if runtime.ObjectSubtype.Null != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Null, enum)
	}

	enum = runtime.ObjectSubtype.Node
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"node"` != string(result) {
		t.Errorf("Expected '\"node\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"node"`), &enum)
	if runtime.ObjectSubtype.Node != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Node, enum)
	}
}

func TestEnumObjectSubtype2(t *testing.T) {
	var enum runtime.ObjectSubtypeEnum
	var err error
	var result []byte

	enum = runtime.ObjectSubtype.Regexp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"regexp"` != string(result) {
		t.Errorf("Expected '\"regexp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"regexp"`), &enum)
	if runtime.ObjectSubtype.Regexp != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Regexp, enum)
	}

	enum = runtime.ObjectSubtype.Date
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"date"` != string(result) {
		t.Errorf("Expected '\"date\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"date"`), &enum)
	if runtime.ObjectSubtype.Date != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Date, enum)
	}

	enum = runtime.ObjectSubtype.Map
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"map"` != string(result) {
		t.Errorf("Expected '\"map\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"map"`), &enum)
	if runtime.ObjectSubtype.Map != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Map, enum)
	}

	enum = runtime.ObjectSubtype.Set
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"set"` != string(result) {
		t.Errorf("Expected '\"set\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"set"`), &enum)
	if runtime.ObjectSubtype.Set != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Set, enum)
	}
}

func TestEnumObjectSubtype3(t *testing.T) {
	var enum runtime.ObjectSubtypeEnum
	var err error
	var result []byte

	enum = runtime.ObjectSubtype.Weakmap
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"weakmap"` != string(result) {
		t.Errorf("Expected '\"weakmap\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"weakmap"`), &enum)
	if runtime.ObjectSubtype.Weakmap != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Weakmap, enum)
	}

	enum = runtime.ObjectSubtype.Weakset
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"weakset"` != string(result) {
		t.Errorf("Expected '\"weakset\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"weakset"`), &enum)
	if runtime.ObjectSubtype.Weakset != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Weakset, enum)
	}

	enum = runtime.ObjectSubtype.Iterator
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"iterator"` != string(result) {
		t.Errorf("Expected '\"iterator\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"iterator"`), &enum)
	if runtime.ObjectSubtype.Iterator != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Iterator, enum)
	}

	enum = runtime.ObjectSubtype.Generator
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"generator"` != string(result) {
		t.Errorf("Expected '\"generator\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"generator"`), &enum)
	if runtime.ObjectSubtype.Generator != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Generator, enum)
	}
}

func TestEnumObjectSubtype4(t *testing.T) {
	var enum runtime.ObjectSubtypeEnum
	var err error
	var result []byte

	enum = runtime.ObjectSubtype.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if runtime.ObjectSubtype.Error != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Error, enum)
	}

	enum = runtime.ObjectSubtype.Proxy
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"proxy"` != string(result) {
		t.Errorf("Expected '\"proxy\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"proxy"`), &enum)
	if runtime.ObjectSubtype.Proxy != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Proxy, enum)
	}

	enum = runtime.ObjectSubtype.Promise
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"promise"` != string(result) {
		t.Errorf("Expected '\"promise\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"promise"`), &enum)
	if runtime.ObjectSubtype.Promise != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Promise, enum)
	}

	enum = runtime.ObjectSubtype.Typedarray
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"typedarray"` != string(result) {
		t.Errorf("Expected '\"typedarray\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"typedarray"`), &enum)
	if runtime.ObjectSubtype.Typedarray != enum {
		t.Errorf("Expcected %d, got %d", runtime.ObjectSubtype.Typedarray, enum)
	}
}
