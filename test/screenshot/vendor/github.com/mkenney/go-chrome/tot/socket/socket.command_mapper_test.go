package socket

import (
	"testing"
)

func TestSocketCommandMapperError(t *testing.T) {
	commandMap := NewCommandMap()
	_, err := commandMap.Get(1)
	if nil == err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}
}
