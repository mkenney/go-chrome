package socket

import (
	"testing"
)

func TestConner(t *testing.T) {
	var err error

	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	err = soc.Connect()
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}

	if !soc.Connected() {
		t.Errorf("Expected true, got false")
	}

	soc.Disconnect()
}
