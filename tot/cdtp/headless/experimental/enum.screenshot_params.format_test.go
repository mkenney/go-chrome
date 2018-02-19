package experimental

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/cdtp/headless/experimental"
)

func TestEnumFormat(t *testing.T) {
	var enum experimental.FormatEnum
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

	enum = experimental.Format.Jpeg
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"jpeg"` != string(result) {
		t.Errorf("Expected '\"jpeg\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"jpeg"`), &enum)
	if experimental.Format.Jpeg != enum {
		t.Errorf("Expcected %d, got %d", experimental.Format.Jpeg, enum)
	}

	enum = experimental.Format.Png
	result, err = json.Marshal(enum)
	if nil != err {
		t.Errorf("Expected nil, got error")
	}
	if `"png"` != string(result) {
		t.Errorf("Expected '\"png\"', got '%s'", result)
	}
	json.Unmarshal([]byte(`"png"`), &enum)
	if experimental.Format.Png != enum {
		t.Errorf("Expcected %d, got %d", experimental.Format.Png, enum)
	}
}
