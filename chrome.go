package chrome

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mkenney/go-chrome/socket"
	log "github.com/sirupsen/logrus"
)

/*
New returns a pointer to a Chrome struct
*/
func New() *Chrome {
	return &Chrome{
		Address:          "localhost",
		Args:             Args{},
		Binary:           "/usr/bin/google-chrome",
		DebuggingAddress: "0.0.0.0",
		DebuggingPort:    9222,
		Output:           "/dev/stdout",
		Port:             9222,
		Tabs:             make([]*Tab, 0),
		Version:          Version{},
		Workdir:          "/tmp/headless-chrome",
	}
}

/*
Chrome is a struct that manages the Chrome process
*/
type Chrome struct {
	// Optional. Address is the domain to use for accessing Chrome sockets (e.g. 'localhost')
	// Defaults to 'localhost'.
	Address string `json:"address"`

	// Args contains CLI arguments for the Chrome binary.
	Args Args `json:"args"`

	// Optional. Binary is the path to the Chrome binary. Defaults to '/usr/bin/google-chrome'.
	Binary string `json:"binary"`

	// Optional. DebuggingAddress is the address number that the remote debugging protocol will be
	// available on. Defaults to '0.0.0.0'.
	DebuggingAddress string `json:"debugging_address"`

	// Optional. DebuggingPort is the port number that the remote debugging protocol will be
	// available on. Defaults to '9222'.
	DebuggingPort int `json:"debugging_port"`

	// Optional. Output is a path to a file to be used to capture STDOUT and STDERR output. Defaults
	// to '/dev/stdout'.
	Output string `json:"output"`

	// Optional. Port is the port number the developer tools endpoints will listen on. Defaults to
	// '9222'
	Port int `json:"port"`

	// Tabs is a list of the currently open tabs.
	Tabs []*Tab `json:"tabs"`

	// Version contains Chrome version information.
	Version Version `json:"version"`

	// Optional. Workdir is the path to the Chrome working directory. Defaults to
	// '/tmp/headless-chrome'.
	Workdir string `json:"workdir"`

	// output is a pointer to a file handle to be used to capture STDOUT and STDERR output.
	output *os.File

	// process is a pointer to the os.Process struct containing the process PID.
	process *os.Process
}

/*
Version is a struct representing the Chrome version information.
*/
type Version struct {
	Browser              string `json:"browser"`
	ProtocolVersion      string `json:"protocol-version"`
	UserAgent            string `json:"user-agent"`
	V8Version            string `json:"v8-version"`
	WebKitVersion        string `json:"webkit-version"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

/*
Args contains arguments to the Chrome executable
*/
type Args map[string]interface{}

/*
Contains checks to see if an argument is present.
*/
func (args Args) Contains(arg string) bool {
	_, ok := args[arg]
	return ok

	//r := regexp.MustCompile("(-?-?)([a-z]+[a-z\\-]+)([=]?)(.*)")
}

/*
Set sets a CLI argument.
*/
func (args Args) Set(arg string, value interface{}) {
	args[arg] = value
	if nil != value {
		switch value.(type) {
		case int:
			args[arg] = value.(int)
		case string:
			args[arg] = value.(string)
		default:
			panic(fmt.Sprintf("Invalid data type %q", value))
		}
	}
}

/*
List returns an array of formatted CLI parameters
*/
func (args Args) List() []string {
	list := make([]string, 0)
	for arg, val := range args {
		if nil == val {
			arg = fmt.Sprintf("--%s", arg)
		} else {
			switch val.(type) {
			case int:
				arg = fmt.Sprintf("--%s=%d", arg, val.(int))
			case string:
				arg = fmt.Sprintf("--%s=%s", arg, val.(string))
			default:
				panic(fmt.Sprintf("Invalid data type %q", val))
			}

		}
		list = append(list, arg)
	}
	return list
}

/*
String returns CLI parameters
*/
func (args Args) String() string {
	return strings.Join(args.List(), " ")
}

/*
Launch launches the Chrome process and returns the connected Chrome struct
*/
func (chrome *Chrome) Launch(args Args) error {
	var err error

	if "" == chrome.Binary {
		chrome.Binary = "/usr/bin/google-chrome"
	}

	args.Set("addr", chrome.Address)
	args.Set("port", chrome.Port)
	args.Set("remote-debugging-port", chrome.DebuggingPort)
	args.Set("remote-debugging-address", chrome.DebuggingAddress)

	if !args.Contains("user-data-dir") {
		args.Set("user-data-dir", "/tmp")
	}

	if "" == chrome.Workdir {
		chrome.Workdir = filepath.Join(os.TempDir(), "headless-chrome")
	}
	if err := os.MkdirAll(chrome.Workdir, 0700); err != nil {
		return fmt.Errorf("Cannot create working directory '%s': %s", chrome.Workdir, err.Error())
	}

	if "" == chrome.Output {
		chrome.Output = "/dev/stdout"
	}

	if nil == chrome.output {
		chrome.output, err = os.OpenFile(chrome.Output, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return fmt.Errorf("Cannot open output file '%s': %s", chrome.Output, err.Error())
		}
	}

	log.Infof("Starting process: %s %s", chrome.Binary, args)
	var procAttributes os.ProcAttr
	procAttributes.Dir = chrome.Workdir
	procAttributes.Files = []*os.File{nil, chrome.output, chrome.output}
	chrome.process, err = os.StartProcess(
		chrome.Binary,
		args.List(),
		&procAttributes,
	)
	if err != nil {
		chrome.output.Close()
		return err
	}

	// Wait up to 10 seconds for Chrome to start
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if err = chrome.checkVersion(); nil == err {
			break
		}
	}
	if err != nil {
		log.Errorf("Chrome took too long to start: %s", err.Error())
		chrome.Close()
		return err
	}

	return nil
}

/*
Close ends the Chrome process and cleans up.
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
		log.Infof("Chrome exited: %s", ps.String())
	}
	if chrome.output != nil {
		chrome.output.Close()
	}
	return nil
}

/*
NewSocket returns a new websocket connected to the Chrome instance for sending
commands through.
*/
func (chrome *Chrome) NewSocket() (*socket.Socket, error) {
	tabs, err := chrome.GetTabs()
	if nil != err {
		log.Fatal(err)
	}
	return socket.New(tabs[0].WebSocketDebuggerURL)
}

/*
GetTab returns a web socket connected to an existing tab in the chrome instance
for sending commands
*/
func (chrome *Chrome) GetTab(tabID string) (tab *Tab, err error) {
	for _, tab = range chrome.Tabs {
		if tab.ID == tabID {
			return tab, nil
		}
	}
	err = fmt.Errorf("Tab '%s' not found", tabID)
	return
}

/*
GetTabs returns an array of Tab structs representing open tabs in the Chrome
instance
*/
func (chrome *Chrome) GetTabs() ([]*Tab, error) {
	_, err := chrome.Cmd("/json/list", url.Values{}, &chrome.Tabs)
	return chrome.Tabs, err
}

func (chrome *Chrome) checkVersion() error {
	if _, err := chrome.Cmd("/json/version", url.Values{}, &chrome.Version); err != nil {
		return err
	}
	log.Infof("Chrome protocol version: %s", chrome.Version.ProtocolVersion)
	return nil
}

/*
Cmd queries the developer tools endpoints and returns JSON data in the
provided struct
*/
func (chrome *Chrome) Cmd(path string, params url.Values, msg interface{}) (interface{}, error) {
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	uri := fmt.Sprintf("http://%s:%d%s", chrome.Address, chrome.Port, path)

	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	log.Infof("chrome:/%s %s", path, resp.Status)
	if 200 != resp.StatusCode {
		return resp.Status, nil
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	} else if err := json.Unmarshal(content, &msg); err != nil {
		return content, nil
	}

	return msg, nil
}
