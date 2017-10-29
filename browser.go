/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

/*
Browser is a struct that manages the Chrome process
*/
type Browser struct {
	output  *os.File
	process *os.Process
	address string
	port    int
	tabs    []*Tab
	version Version
}

/*
Tab is a struct representing an individual Chrome tab
*/
type Tab struct {
	Description          string `json:"description"`
	DevtoolsFrontendURL  string `json:"devtoolsFrontendUrl"`
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	URL                  string `json:"url"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
	Socket               *Socket
}

/*
Version is a struct representing the Chrome version
*/
type Version struct {
	Browser         string `json:"browser"`
	ProtocolVersion string `json:"protocol-version"`
	UserAgent       string `json:"user-agent"`
	WebKitVersion   string `json:"webkit-version"`
}

/*
New launches the Chrome process and returns the connected Browser struct
*/
func New(port int, addr, proxy, binary string) (*Browser, error) {
	args := []string{
		fmt.Sprintf("--port=%d", port),
		fmt.Sprintf("--addr=%s", addr),
	}
	if proxy != "" {
		args = append(args, "--proxy="+proxy)
	}

	workDir := filepath.Join(os.TempDir(), fmt.Sprintf("headless-chrome-%x", time.Now().UnixNano()))
	if err := os.MkdirAll(workDir, 0700); err != nil {
		return nil, fmt.Errorf("Cannot create working directory '%s': %v", workDir, err)
	}

	outputFile := filepath.Join(workDir, "output")
	output, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("Cannot create output file '%s': %v", outputFile, err)
	}

	log.Printf("Starting %s %v (working directory: %s) ...", binary, args, workDir)
	var procAttributes os.ProcAttr
	procAttributes.Dir = workDir
	procAttributes.Files = []*os.File{nil, output, output}
	fmt.Printf("DEBUG---------\n%v %v %v------------DEBUG\n", binary, args, &procAttributes)
	process, err := os.StartProcess(binary, args, &procAttributes)
	if err != nil {
		output.Close()
		return nil, err
	}

	browser := &Browser{
		output:  output,
		process: process,
		address: addr,
		port:    port,
	}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		if err = browser.checkVersion(); err == nil {
			break
		}
	}
	if err != nil {
		browser.Close()
		return nil, err
	}

	return browser, nil
}

/*
GetBrowser returns the current Chrome process
*/
func GetBrowser() (*Browser, error) {
	browser := &Browser{}
	if err := browser.checkVersion(); err != nil {
		return nil, err
	}
	return browser, nil
}

/*
Close ends the Chrome process and cleans up
*/
func (b *Browser) Close() error {
	if b.process != nil {
		if err := b.process.Signal(os.Interrupt); err != nil {
			return err
		}
		ps, err := b.process.Wait()
		if err != nil {
			return err
		}
		log.Printf("Headless Chromium exited: %s", ps.String())
	}
	if b.output != nil {
		b.output.Close()
	}
	return nil
}

/*
NewSocket returns a new websocket connected to the Chrome instance for sending
commands through
*/
func (b *Browser) NewSocket() (*Socket, error) {
	return newSocket(fmt.Sprintf("ws://%s:%d/devtools/browser", b.address, b.port))
}

/*
NewTab returns a web socket connected to a new tab in the chrome instance for
sending commands through
*/
func (b *Browser) NewTab(tabID string) (*Tab, error) {
	var tab *Tab
	var tabList []Tab

	for _, tab := range b.tabs {
		if tab.ID == tabID {
			return tab, fmt.Errorf("Tab '%s' already exists", tabID)
		}
	}

	socket, err := newSocket(fmt.Sprintf("ws://%s:%d/devtools/page/%s", b.address, b.port, tabID))
	if err != nil {
		return tab, err
	}

	err = b.getJSON("/json/list", tabList)
	for _, tmp := range tabList {
		if tmp.ID == tabID {
			tab = new(Tab)
			tab.Description = tmp.Description
			tab.DevtoolsFrontendURL = tmp.DevtoolsFrontendURL
			tab.ID = tmp.ID
			tab.Title = tmp.Title
			tab.Type = tmp.Type
			tab.URL = tmp.URL
			tab.WebSocketDebuggerURL = tmp.WebSocketDebuggerURL
			tab.Socket = socket
			b.tabs = append(b.tabs, tab)
			return tab, nil
		}
	}

	return tab, fmt.Errorf("Could not create tab '%s'", tabID)
}

/*
GetTab returns a web socket connected to an existing tab in the chrome instance
for sending commands
*/
func (b *Browser) GetTab(tabID string) (tab *Tab, err error) {
	for _, tab = range b.tabs {
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
	err := b.getJSON("/json/list", &b.tabs)
	return b.tabs, err
}

func (b *Browser) checkVersion() error {
	if err := b.getJSON("/json/version", &b.version); err != nil {
		return err
	}
	log.Printf("Browser protocol version: %v", b.version.ProtocolVersion)
	return nil
}

func (b *Browser) getJSON(path string, msg interface{}) error {
	uri := fmt.Sprintf("http://%s:%d%s", b.address, b.port, path)
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if content, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else if err := json.Unmarshal(content, msg); err != nil {
		return err
	}
	return nil
}
