package chrome

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
New returns a pointer to a Chromium instance.
*/
func New(
	flags ChromiumFlags,
	binary string,
	workdir string,
	stdout string,
	stderr string,
) Chromium {
	return &Chrome{
		flags:   flags,
		binary:  binary,
		stderr:  stderr,
		stdout:  stdout,
		workdir: workdir,
	}
}

/*
Chrome implements Chromium.
*/
type Chrome struct {
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

Default value is 'localhost'
*/
func (chrome *Chrome) Address() string {
	if !chrome.Flags().Has("addr") {
		chrome.Flags().Set("addr", []interface{}{"localhost"})
	}
	value, _ := chrome.Flags().Get("addr")
	return value[0].(string)
}

/*
Flags implements Chromium.
*/
func (chrome *Chrome) Flags() ChromiumFlags {
	return chrome.flags
}

/*
Binary implements Chromium.

Default value is '/usr/bin/google-chrome' for use with the mkenney/chromium-headless
Docker image.
*/
func (chrome *Chrome) Binary() string {
	if "" == chrome.binary {
		chrome.binary = "/usr/bin/google-chrome"
	}
	return chrome.binary
}

/*
Close implements Chromium.
*/
func (chrome *Chrome) Close() error {
	if chrome.process != nil {
		if err := chrome.process.Signal(os.Interrupt); err != nil {
			return err
		}
		ps, err := chrome.process.Wait()
		if err != nil {
			return err
		}
		log.Infof("Chromium exited: %s", ps.String())
	}
	if chrome.stdOUTFile != nil {
		chrome.stdOUTFile.Close()
	}
	return nil
}

/*
DebuggingAddress implements Chromium.

Default value is '0.0.0.0'.
*/
func (chrome *Chrome) DebuggingAddress() string {
	if !chrome.Flags().Has("remote-debugging-address") {
		chrome.Flags().Set("remote-debugging-address", []interface{}{"0.0.0.0"})
	}
	value, _ := chrome.Flags().Get("remote-debugging-address")
	return value[0].(string)
}

/*
DebuggingPort implements Chromium.
*/
func (chrome *Chrome) DebuggingPort() int {
	if !chrome.Flags().Has("remote-debugging-port") {
		chrome.Flags().Set("remote-debugging-port", []interface{}{9222})
	}
	value, _ := chrome.Flags().Get("remote-debugging-port")
	return value[0].(int)
}

/*
GetTab implements Chromium.
*/
func (chrome *Chrome) GetTab(tabID string) (tab Tabber, err error) {
	for _, tab = range chrome.tabs {
		if tab.Data().ID == tabID {
			return tab, nil
		}
	}
	err = fmt.Errorf("Tab '%s' not found", tabID)
	return
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
func (chrome *Chrome) Launch() error {
	var err error

	// Default values for required parameters
	chrome.Address()
	chrome.DebuggingAddress()
	chrome.DebuggingPort()
	chrome.Port()
	if !chrome.Flags().Has("user-data-dir") {
		chrome.Flags().Set("user-data-dir", []interface{}{os.TempDir()})
	}

	if err := os.MkdirAll(chrome.Workdir(), 0700); err != nil {
		return fmt.Errorf("Cannot create working directory '%s'", chrome.Workdir())
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
			return fmt.Errorf("Cannot open error output file '%s'", chrome.STDERR())
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
			return fmt.Errorf("Cannot open standard output file '%s'", chrome.STDOUT())
		}
	}

	log.Infof("Starting process: %s %s", chrome.Binary(), chrome.Flags())
	var procAttributes os.ProcAttr
	procAttributes.Dir = chrome.Workdir()
	procAttributes.Files = []*os.File{nil, chrome.stdOUTFile, chrome.stdERRFile}
	chrome.process, err = os.StartProcess(
		chrome.Binary(),
		chrome.Flags().List(),
		&procAttributes,
	)
	if nil != err {
		chrome.stdOUTFile.Close()
		return err
	}

	// Wait up to 10 seconds for Chromium to start
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if _, err = chrome.Version(); nil == err {
			break
		}
	}
	if err != nil {
		log.Errorf("Chromium took too long to start")
		log.Debugf(err.Error())
		chrome.Close()
		return err
	}

	return nil
}

/*
Port implements Chromium.

Default value is 9222
*/
func (chrome *Chrome) Port() int {
	if !chrome.Flags().Has("port") {
		chrome.Flags().Set("port", []interface{}{9222})
	}
	value, _ := chrome.Flags().Get("port")
	return value[0].(int)
}

/*
Query implements Chromium.
*/
func (chrome *Chrome) Query(
	path string,
	params url.Values,
	msg interface{}, // Data receiver
) (interface{}, error) {
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	uri := fmt.Sprintf("http://%s:%d%s", chrome.Address(), chrome.Port(), path)
	log.Errorf("%s", uri)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Debugf("chrome:/%s %s", path, resp.Status)
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else if err := json.Unmarshal(content, &msg); err != nil {
		return content, nil
	}

	return msg, nil
}

/*
STDERR implements Chromium.
*/
func (chrome *Chrome) STDERR() string {
	return chrome.stderr
}

/*
STDOUT implements Chromium.
*/
func (chrome *Chrome) STDOUT() string {
	return chrome.stdout
}

/*
Tabs implements Chromium.
*/
func (chrome *Chrome) Tabs() []*Tab {
	return chrome.tabs
}

/*
Version implements Chromium.
*/
func (chrome *Chrome) Version() (*Version, error) {
	if nil == chrome.version {
		if _, err := chrome.Query(
			"/json/version",
			url.Values{},
			&chrome.version,
		); err != nil {
			return nil, err
		}
	}
	return chrome.version, nil
}

/*
Workdir implements Chromium.

Default value is /tmp/headless-chrome
*/
func (chrome *Chrome) Workdir() string {
	if "" == chrome.workdir {
		chrome.workdir = filepath.Join(os.TempDir(), "headless-chrome")
	}
	return chrome.workdir
}
