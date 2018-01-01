package socket

import (
	"testing"
)

func TestConner(t *testing.T) {
	var err error
	socket, _ := NewMock("http://www.example.com/")

	err = socket.Connect()
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}

	err = socket.Disconnect()
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}
}
