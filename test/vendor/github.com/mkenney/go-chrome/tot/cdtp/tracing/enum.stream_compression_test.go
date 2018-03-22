package tracing

import (
	"encoding/json"
	"testing"
)

func TestEnumStreamCompression(t *testing.T) {
	var enum StreamCompressionEnum
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

	enum = StreamCompression.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if StreamCompression.None != enum {
		t.Errorf("Expcected %d, got %d", StreamCompression.None, enum)
	}

	enum = StreamCompression.Gzip
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"gzip"` != string(result) {
		t.Errorf("Expected '\"gzip\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"gzip"`), &enum)
	if StreamCompression.Gzip != enum {
		t.Errorf("Expcected %d, got %d", StreamCompression.Gzip, enum)
	}
}
