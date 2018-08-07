package chrome

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
)

/*
NewMock returns a pointer to a mock Chromium instance.
*/
func NewMock(
	flags ChromiumFlags,
	binary string,
	workdir string,
	stdout string,
	stderr string,
) *MockChrome {
	return &MockChrome{
		flags:   flags,
		binary:  binary,
		stderr:  stderr,
		stdout:  stdout,
		workdir: workdir,
	}
}

/*
MockChrome implements Chromium.
*/
type MockChrome struct {
	// flags stores CLI arguments for the Chromium binary.
	flags ChromiumFlags

	// Optional. binary is the path to the Chromium binary. Defaults to
	// '/usr/bin/google-chrome'.
	binary string

	// Optional. port is the port number the developer tools endpoints will
	// listen on. Defaults to 9222.
	//port int

	// tabs is a list of the currently open tabs.
	tabs []*Tab

	// version contains Chromium version information.
	version *Version

	// Optional. workdir is the path to the Chromium working directory. Defaults
	// to '/tmp/headless-chrome'.
	workdir string

	// Optional. stderr is a path to a file to be used to capture STDERR output.
	// Defaults to the system STDERR.
	stderr string

	// Optional. stdout is a path to a file to be used to capture STDOUT output.
	// Defaults to the system STDOUT.
	stdout string

	// stdERRFile is a pointer to a file handle to be used to capture STDERR
	// output.
	stdERRFile *os.File

	// stdOUTFile is a pointer to a file handle to be used to capture STDOUT
	// output.
	stdOUTFile *os.File

	// process is a pointer to the os.Process struct containing the process PID.
	process *os.Process
}

/*
Address implements Chromium.
*/
func (chrome *MockChrome) Address() string {
	return "localhost"
}

/*
Flags implements Chromium.
*/
func (chrome *MockChrome) Flags() ChromiumFlags {
	return chrome.flags
}

/*
Binary implements Chromium.
*/
func (chrome *MockChrome) Binary() string {
	if "" == chrome.binary {
		chrome.binary = "/usr/bin/google-chrome"
	}
	return chrome.binary
}

/*
Close implements Chromium.
*/
func (chrome *MockChrome) Close() error {
	if chrome.stdOUTFile != nil {
		chrome.stdOUTFile.Close()
	}
	return nil
}

/*
DebuggingAddress implements Chromium.

Default value is '0.0.0.0'.
*/
func (chrome *MockChrome) DebuggingAddress() string {
	if !chrome.Flags().Has("remote-debugging-address") {
		chrome.Flags().Set("remote-debugging-address", "0.0.0.0")
	}
	value, _ := chrome.Flags().Get("remote-debugging-address")
	return value.(string)
}

/*
DebuggingPort implements Chromium.
*/
func (chrome *MockChrome) DebuggingPort() int {
	if !chrome.Flags().Has("remote-debugging-port") {
		chrome.Flags().Set("remote-debugging-port", 9222)
	}
	value, _ := chrome.Flags().Get("remote-debugging-port")
	return value.(int)
}

/*
GetTab implements Chromium.
*/
func (chrome *MockChrome) GetTab(tabID string) (Tabber, error) {
	var tab Tabber
	var err error
	for _, tab = range chrome.tabs {
		if tab.Data().ID == tabID {
			return tab, nil
		}
	}
	err = errs.New(0, fmt.Sprintf("tab '%s' not found", tabID))
	return tab, err
}

/*
Launch implements Chromium.

This implementation makes it's best effort to set a few sane default values if
they aren't included in the Flags definition:

	addr = "localhost"
	remote-debugging-address = "0.0.0.0"
	remote-debugging-port = 9222
	port = 9222
	user-data-dir = os.TempDir() + chrome.Workdir()
	chrome.workdir = "headless-chrome"
	chrome.output = "/dev/stdout"
*/
func (chrome *MockChrome) Launch() error {
	var err error

	// Default values for required parameters
	chrome.Address()
	chrome.DebuggingAddress()
	chrome.DebuggingPort()
	chrome.Port()
	if !chrome.Flags().Has("user-data-dir") {
		chrome.Flags().Set("user-data-dir", os.TempDir())
	}

	if "" == chrome.STDERR() {
		chrome.stdERRFile = os.Stderr
	} else {
		chrome.stdERRFile, err = os.OpenFile(
			chrome.STDERR(),
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0600,
		)
		if err != nil {
			return errs.Wrap(err, 0, fmt.Sprintf("cannot open error output file '%s'", chrome.STDERR()))
		}
	}

	if "" == chrome.STDOUT() {
		chrome.stdOUTFile = os.Stdout
	} else {
		chrome.stdOUTFile, err = os.OpenFile(
			chrome.STDOUT(),
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0600,
		)
		if err != nil {
			return errs.Wrap(err, 0, fmt.Sprintf("cannot open standard output file '%s'", chrome.STDOUT()))
		}
	}

	log.Infof("Starting process: %s %s", chrome.Binary(), chrome.Flags())
	if nil != err {
		chrome.stdOUTFile.Close()
		chrome.stdERRFile.Close()
		return errs.Wrap(err, 0, "error starting chrome")
	}

	return nil
}

/*
Port implements Chromium.

Default value is 9222
*/
func (chrome *MockChrome) Port() int {
	if !chrome.Flags().Has("port") {
		chrome.Flags().Set("port", 9222)
	}
	value, _ := chrome.Flags().Get("port")
	return value.(int)
}

/*
Query implements Chromium.
*/
func (chrome *MockChrome) Query(
	path string,
	params url.Values,
	msg interface{}, // Data receiver
) (interface{}, error) {
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	log.Debugf("chrome:/%s %s", path)
	return msg, nil
}

/*
STDERR implements Chromium.
*/
func (chrome *MockChrome) STDERR() string {
	return chrome.stderr
}

/*
STDOUT implements Chromium.
*/
func (chrome *MockChrome) STDOUT() string {
	return chrome.stdout
}

/*
Tabs implements Chromium.
*/
func (chrome *MockChrome) Tabs() []*Tab {
	return chrome.tabs
}

/*
Version implements Chromium.
*/
func (chrome *MockChrome) Version() (*Version, error) {
	if nil == chrome.version {
		chrome.version = &Version{}
	}
	return chrome.version, nil
}

/*
Workdir implements Chromium.

Default value is /tmp/headless-chrome
*/
func (chrome *MockChrome) Workdir() string {
	if "" == chrome.workdir {
		chrome.workdir = filepath.Join(os.TempDir(), "headless-chrome")
	}
	return chrome.workdir
}

/*
NewTab spawns a new Tab and returns a reference to it
*/
func (chrome *MockChrome) NewTab(uri string) (*Tab, error) {
	var err error

	if "" == uri {
		uri = "about:blank"
	}
	targetURL, err := url.Parse(uri)
	if nil != err {
		return nil, errs.Wrap(err, 0, "invalid URL")
	}

	tab := &Tab{
		chrome: chrome,
		data: &TabData{
			Description:          "",
			DevtoolsFrontendURL:  "",
			ID:                   "",
			Title:                "",
			Type:                 "",
			URL:                  "",
			WebSocketDebuggerURL: "",
		},
		url: targetURL,
	}

	socket := NewMockSocket(targetURL)
	tab.socket = socket
	tab.protocol = socket
	chrome.tabs = append(chrome.tabs, tab)

	return tab, nil
}
