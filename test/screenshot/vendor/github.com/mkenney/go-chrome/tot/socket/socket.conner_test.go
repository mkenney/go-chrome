package socket

import (
	"net/url"
	"testing"
)

func TestConner(t *testing.T) {
	var err error
	socketURL, _ := url.Parse("http://www.example.com/")
	socket := NewMock(socketURL)

	err = socket.Connect()
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}

	if !socket.Connected() {
		t.Errorf("Expected true, got false")
	}

	err = socket.Disconnect()
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}
}
