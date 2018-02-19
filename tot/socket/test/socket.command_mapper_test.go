package test

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/socket"
)

func TestSocketCommandMapperError(t *testing.T) {
	commandMap := socket.NewCommandMap()
	_, err := commandMap.Get(1)
	if nil == err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}
}
