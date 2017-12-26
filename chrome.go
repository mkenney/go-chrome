package Chrome

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

func New() *Chrome {
	return &Chrome{}
}

/*
Chrome is a struct that manages the Chrome process
*/
type Chrome struct {
	Address string
	Output  *os.File
	Port    int
	Process *os.Process
	Tabs    []*Tab
	Version Version
}

/*
Version is a struct representing the Chrome version
*/
type Version struct {
	Browser              string `json:"browser"`
	ProtocolVersion      string `json:"protocol-version"`
	UserAgent            string `json:"user-agent"`
	V8Version            string `json:"v8-version"`
	WebKitVersion        string `json:"webkit-version"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

var browserInstance *Chrome

/*
Args contains arguments to the Chrome executable
*/
type Args map[string]interface{}

/*
Contains checks to see if an argument is present
*/
func (args Args) Contains(arg string) bool {
	_, ok := args[arg]
	return ok

	//r := regexp.MustCompile("(-?-?)([a-z]+[a-z\\-]+)([=]?)(.*)")
}

/*
Add adds an argument to the array
*/
func (args Args) Add(arg string, value interface{}) {
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
func Launch(binary string, args Args) error {
	if "" == binary {
		binary = "/usr/bin/google-chrome"
	}

	if !args.Contains("port") {
		args.Add("port", 9222)
	}

	if !args.Contains("addr") {
		args.Add("addr", "localhost")
	}

	if !args.Contains("remote-debugging-port") {
		args.Add("remote-debugging-port", "9222")
	}

	if !args.Contains("remote-debugging-address") {
		args.Add("remote-debugging-address", "0.0.0.0")
	}

	if !args.Contains("user-data-dir") {
		args.Add("user-data-dir", "/tmp")
	}

	workDir := filepath.Join(os.TempDir(), "headless-chrome")
	if err := os.MkdirAll(workDir, 0700); err != nil {
		return fmt.Errorf("Cannot create working directory '%s': %s", workDir, err.Error())
	}

	outputFile := filepath.Join(workDir, "output")
	output, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("Cannot open output file '%s': %s", outputFile, err.Error())
	}

	log.Infof("Starting process: %s %s", binary, args)
	var procAttributes os.ProcAttr
	procAttributes.Dir = workDir
	procAttributes.Files = []*os.File{nil, output, output}
	process, err := os.StartProcess(binary, args.List(), &procAttributes)
	if err != nil {
		output.Close()
		return err
	}

	browserInstance = &Chrome{
		Address: args["addr"].(string),
		Output:  output,
		Port:    args["port"].(int),
		Process: process,
	}
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		if err = browserInstance.checkVersion(); err == nil {
			break
		}
	}
	if err != nil {
		browserInstance.Close()
		return err
	}
	return nil
}

/*
GetChrome returns the current Chrome process
*/
func GetChrome() (*Chrome, error) {
	if nil == browserInstance {
		err := Launch("", Args{})
		if nil != err {
			return nil, err
		}
	}
	return browserInstance, nil
}

/*
Close ends the Chrome process and cleans up
*/
func (b *Chrome) Close() error {
	if b.Process != nil {
		if err := b.Process.Signal(os.Interrupt); err != nil {
			return err
		}
		ps, err := b.Process.Wait()
		if err != nil {
			return err
		}
		log.Infof("Chrome exited: %s", ps.String())
	}
	if b.Output != nil {
		b.Output.Close()
	}
	return nil
}

/*
NewSocket returns a new websocket connected to the Chrome instance for sending
commands through
*/
func (b *Chrome) NewSocket() (*socket.Socket, error) {
	tabs, err := b.GetTabs()
	if nil != err {
		log.Fatal(err)
	}
	return socket.New(tabs[0].WebSocketDebuggerURL)
}

/*
GetTab returns a web socket connected to an existing tab in the chrome instance
for sending commands
*/
func (b *Chrome) GetTab(tabID string) (tab *Tab, err error) {
	for _, tab = range b.Tabs {
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
func (b *Chrome) GetTabs() ([]*Tab, error) {
	_, err := b.Cmd("/json/list", url.Values{}, &b.Tabs)
	return b.Tabs, err
}

func (b *Chrome) checkVersion() error {
	if _, err := b.Cmd("/json/version", url.Values{}, &b.Version); err != nil {
		return err
	}
	log.Infof("Chrome protocol version: %s", b.Version.ProtocolVersion)
	return nil
}

/*
Cmd queries the developer tools endpoints and returns JSON data in the
provided struct
*/
func (b *Chrome) Cmd(path string, params url.Values, msg interface{}) (interface{}, error) {
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	uri := fmt.Sprintf("http://%s:%d%s", b.Address, b.Port, path)

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
