package runtime

import (
	"encoding/json"
	"testing"
)

func TestEnumObjectSubtype(t *testing.T) {
	var enum ObjectSubtypeEnum
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

	enum = ObjectSubtype.Array
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"array"` != string(result) {
		t.Errorf("Expected '\"array\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"array"`), &enum)
	if ObjectSubtype.Array != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Array, enum)
	}

	enum = ObjectSubtype.Null
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"null"` != string(result) {
		t.Errorf("Expected '\"null\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"null"`), &enum)
	if ObjectSubtype.Null != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Null, enum)
	}

	enum = ObjectSubtype.Node
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"node"` != string(result) {
		t.Errorf("Expected '\"node\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"node"`), &enum)
	if ObjectSubtype.Node != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Node, enum)
	}

	enum = ObjectSubtype.Regexp
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"regexp"` != string(result) {
		t.Errorf("Expected '\"regexp\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"regexp"`), &enum)
	if ObjectSubtype.Regexp != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Regexp, enum)
	}

	enum = ObjectSubtype.Date
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"date"` != string(result) {
		t.Errorf("Expected '\"date\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"date"`), &enum)
	if ObjectSubtype.Date != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Date, enum)
	}

	enum = ObjectSubtype.Map
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"map"` != string(result) {
		t.Errorf("Expected '\"map\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"map"`), &enum)
	if ObjectSubtype.Map != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Map, enum)
	}

	enum = ObjectSubtype.Set
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"set"` != string(result) {
		t.Errorf("Expected '\"set\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"set"`), &enum)
	if ObjectSubtype.Set != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Set, enum)
	}

	enum = ObjectSubtype.Weakmap
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"weakmap"` != string(result) {
		t.Errorf("Expected '\"weakmap\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"weakmap"`), &enum)
	if ObjectSubtype.Weakmap != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Weakmap, enum)
	}

	enum = ObjectSubtype.Weakset
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"weakset"` != string(result) {
		t.Errorf("Expected '\"weakset\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"weakset"`), &enum)
	if ObjectSubtype.Weakset != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Weakset, enum)
	}

	enum = ObjectSubtype.Iterator
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"iterator"` != string(result) {
		t.Errorf("Expected '\"iterator\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"iterator"`), &enum)
	if ObjectSubtype.Iterator != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Iterator, enum)
	}

	enum = ObjectSubtype.Generator
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"generator"` != string(result) {
		t.Errorf("Expected '\"generator\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"generator"`), &enum)
	if ObjectSubtype.Generator != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Generator, enum)
	}

	enum = ObjectSubtype.Error
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"error"` != string(result) {
		t.Errorf("Expected '\"error\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"error"`), &enum)
	if ObjectSubtype.Error != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Error, enum)
	}

	enum = ObjectSubtype.Proxy
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"proxy"` != string(result) {
		t.Errorf("Expected '\"proxy\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"proxy"`), &enum)
	if ObjectSubtype.Proxy != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Proxy, enum)
	}

	enum = ObjectSubtype.Promise
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"promise"` != string(result) {
		t.Errorf("Expected '\"promise\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"promise"`), &enum)
	if ObjectSubtype.Promise != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Promise, enum)
	}

	enum = ObjectSubtype.Typedarray
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"typedarray"` != string(result) {
		t.Errorf("Expected '\"typedarray\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"typedarray"`), &enum)
	if ObjectSubtype.Typedarray != enum {
		t.Errorf("Expcected %d, got %d", ObjectSubtype.Typedarray, enum)
	}
}
