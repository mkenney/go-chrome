package tracing_test

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/tracing"
)

func TestEnumStreamCompression(t *testing.T) {
	var enum tracing.StreamCompressionEnum
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

	enum = tracing.StreamCompression.None
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"none"` != string(result) {
		t.Errorf("Expected '\"none\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"none"`), &enum)
	if tracing.StreamCompression.None != enum {
		t.Errorf("Expcected %d, got %d", tracing.StreamCompression.None, enum)
	}

	enum = tracing.StreamCompression.Gzip
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"gzip"` != string(result) {
		t.Errorf("Expected '\"gzip\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"gzip"`), &enum)
	if tracing.StreamCompression.Gzip != enum {
		t.Errorf("Expcected %d, got %d", tracing.StreamCompression.Gzip, enum)
	}
}
