package chrome

import (
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func TestChromiumNew(t *testing.T) {
	chrome := New(
		&Flags{},
		"", //"path/to/chrome",
		"", //"path/to/stderr",
		"", //"path/to/stdout",
		"", //"path/to/workdir",
	)
	if "localhost" != chrome.Address() {
		t.Errorf("Expected 'localhost', received '%s'", chrome.Address())
	}
	if "/usr/bin/google-chrome" != chrome.Binary() {
		t.Errorf("Expected '/usr/bin/google-chrome', received '%s'", chrome.Binary())
	}
	if "0.0.0.0" != chrome.DebuggingAddress() {
		t.Errorf("Expected '0.0.0.0', received '%s'", chrome.DebuggingAddress())
	}
	if 9222 != chrome.DebuggingPort() {
		t.Errorf("Expected 9222, received '%d'", chrome.DebuggingPort())
	}
	if 9222 != chrome.Port() {
		t.Errorf("Expected 9222, received '%d'", chrome.Port())
	}
	if "" != chrome.STDERR() {
		t.Errorf("Expected empty string, received '%s'", chrome.STDERR())
	}
	if "" != chrome.STDOUT() {
		t.Errorf("Expected empty string, received '%s'", chrome.STDOUT())
	}
	if filepath.Join(os.TempDir(), "headless-chrome") != chrome.Workdir() {
		t.Errorf("Expected '%s', received '%s'", filepath.Join(os.TempDir()), chrome.Workdir())
	}
}

func TestChromiumClose(t *testing.T) {
	chrome := New(
		&Flags{},
		"", //"path/to/chrome",
		"", //"path/to/stderr",
		"", //"path/to/stdout",
		"", //"path/to/workdir",
	)
	chrome.Close()
}

func TestChromiumGetTab(t *testing.T) {
	chrome := New(
		&Flags{},
		"", //"path/to/chrome",
		"", //"path/to/stderr",
		"", //"path/to/stdout",
		"", //"path/to/workdir",
	)
	_, err := chrome.GetTab("some-tab")
	if nil == err {
		t.Errorf("Expected error, received nil")
	}
}

func TestChromiumLaunch(t *testing.T) {
	chrome := New(
		&Flags{},
		"", //"path/to/chrome",
		"", //"path/to/stderr",
		"", //"path/to/stdout",
		"", //"path/to/workdir",
	)
	err := chrome.Launch()
	if nil == err {
		t.Errorf("Expected error, received nil")
	}
}

func TestChromiumQuery(t *testing.T) {
	chrome := New(
		&Flags{
			"remote-debugging-port": 0,
			"port":                  0,
		},
		"",
		"",
		"",
		"",
	)
	data, err := chrome.Query("/json/version", url.Values{}, nil)
	if nil == err {
		t.Errorf("Expected error, received nil")
	}
	if nil != data {
		t.Errorf("Expected nil, received %v", data)
	}
}

func TestChromiumTabs(t *testing.T) {
	chrome := New(
		&Flags{
			"remote-debugging-port": 0,
			"port":                  0,
		},
		"", //"path/to/chrome",
		"", //"path/to/stderr",
		"", //"path/to/stdout",
		"", //"path/to/workdir",
	)
	tabs := chrome.Tabs()
	if nil != tabs {
		t.Errorf("Expected nil, received %v", tabs)
	}
}

func TestChromiumVersion(t *testing.T) {
	chrome := New(
		&Flags{
			"addr":                     "devnul",
			"remote-debugging-address": "devnul",
			"port":                     9222,
			"remote-debugging-port":    9222,
		},
		"", //"path/to/chrome",
		"", //"path/to/stderr",
		"", //"path/to/stdout",
		"", //"path/to/workdir",
	)
	version, err := chrome.Version()
	if nil == err {
		t.Errorf("Expected error, received nil")
	}
	if nil != version {
		t.Errorf("Expected nil, received %v", version)
	}
}
