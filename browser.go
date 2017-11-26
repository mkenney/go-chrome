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

	log "github.com/Sirupsen/logrus"
)

/*
Browser is a struct that manages the Chrome process
*/
type Browser struct {
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

var browserInstance *Browser

/*
New launches the Chrome process and returns the connected Browser struct
*/
func Launch(port int, address, proxy, binary string) error {
	if 0 == port {
		port = 9222
	}
	if "" == address {
		address = "localhost"
	}
	if "" == binary {
		binary = "/usr/bin/google-chrome"
	}

	args := []string{
		fmt.Sprintf("--port=%d", port),
		fmt.Sprintf("--addr=%s", address),
		"--remote-debugging-port=9222",
		"--remote-debugging-address=0.0.0.0",
		"--disable-gpu",
		"--headless",
		"--hide-scrollbars",
		"--no-sandbox",
		"--no-first-run",
		"--disable-extensions",
		"--user-data-dir=/tmp",
	}
	if proxy != "" {
		args = append(args, "--proxy="+proxy)
	}

	workDir := filepath.Join(os.TempDir(), fmt.Sprintf("headless-chrome-%x", time.Now().UnixNano()))
	if err := os.MkdirAll(workDir, 0700); err != nil {
		return fmt.Errorf("Cannot create working directory '%s': %v", workDir, err)
	}

	outputFile := filepath.Join(workDir, "output")
	output, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("Cannot create output file '%s': %v", outputFile, err)
	}

	log.Infof("Starting %s %s", binary, strings.Join(args, " "))
	var procAttributes os.ProcAttr
	procAttributes.Dir = workDir
	procAttributes.Files = []*os.File{nil, output, output}
	process, err := os.StartProcess(binary, args, &procAttributes)
	if err != nil {
		output.Close()
		return err
	}

	browserInstance = new(Browser)
	browserInstance.Output = output
	browserInstance.Process = process
	browserInstance.Address = address
	browserInstance.Port = port
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
GetBrowser returns the current Chrome process
*/
func GetBrowser() (*Browser, error) {
	if nil == browserInstance {
		err := Launch(0, "", "", "")
		if nil != err {
			return nil, err
		}
	}
	return browserInstance, nil
}

/*
Close ends the Chrome process and cleans up
*/
func (b *Browser) Close() error {
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
func (b *Browser) NewSocket() (*Socket, error) {
	tabs, err := b.GetTabs()
	if nil != err {
		log.Fatal(err)
	}
	return NewSocket(tabs[0])
}

/*
GetTab returns a web socket connected to an existing tab in the chrome instance
for sending commands
*/
func (b *Browser) GetTab(tabID string) (tab *Tab, err error) {
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
func (b *Browser) GetTabs() ([]*Tab, error) {
	_, err := b.Cmd("/json/list", url.Values{}, &b.Tabs)
	return b.Tabs, err
}

func (b *Browser) checkVersion() error {
	if _, err := b.Cmd("/json/version", url.Values{}, &b.Version); err != nil {
		return err
	}
	log.Infof("Browser protocol version: %s", b.Version.ProtocolVersion)
	return nil
}

/*
Cmd queries the developer tools endpoints and returns JSON data in the
provided struct
*/
func (b *Browser) Cmd(path string, params url.Values, msg interface{}) (interface{}, error) {
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
