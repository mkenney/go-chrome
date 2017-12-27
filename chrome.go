/*
Package chrome aims to be a complete Chrome DevTools Protocol Viewer implementation
https://chromedevtools.github.io/devtools-protocol/

Work in progress

You should probably ignore this for now

From the official documentation

The Chrome DevTools Protocol allows for tools to instrument, inspect, debug and profile Chromium,
Chrome and other Blink-based browsers. Many existing projects currently use the protocol. The Chrome
DevTools uses this protocol and the team maintains its API.

Instrumentation is divided into a number of domains (DOM, Debugger, Network etc.). Each domain
defines a number of commands it supports and events it generates. Both commands and events are
serialized JSON objects of a fixed structure. You can either debug over the wire using the raw
messages as they are described in the corresponding domain documentation, or use extension
JavaScript API.

Protocol API Docs

The latest (tip-of-tree) protocol (tot)

It changes frequently and can break at any time. However it captures the full capabilities of the
Protocol, whereas the stable release is a subset. There is no backwards compatibility support
guaranteed for the capabilities it introduces.

Resources

	* chrome-debugging-protocol mailing list https://groups.google.com/d/forum/chrome-debugging-protocol
	* devtools-protocol repo issue tracker https://github.com/chromedevtools/devtools-protocol
		can also be used for concerns with the protocol. It also hosts the canonical copy of the
		json files.
	* Getting Started with Headless Chrome https://developers.google.com/web/updates/2017/04/headless-chrome
	* Headless Chromium readme https://chromium.googlesource.com/chromium/src/+/lkgr/headless/README.md
	* chrome-remote-interface node module https://github.com/cyrus-and/chrome-remote-interface/
		Also the wiki https://github.com/cyrus-and/chrome-remote-interface/wiki
		and issue tracker https://github.com/cyrus-and/chrome-remote-interface/issues?utf8=%E2%9C%93&q=is%3Aissue%20
	*  awesome-chrome-devtools page https://github.com/ChromeDevTools/awesome-chrome-devtools#chrome-devtools-protocol
		links to many of the tools in the protocol ecosystem, including protocol API libraries in
		JavaScript, TypeScript, Python, Java, and Go. Many applications and libraries already use
		the protocol.


Basics: Using DevTools as protocol client

The Developer Tools front-end can attach to a remotely running Chrome instance for debugging. For
this scenario to work, you should start your host Chrome instance with the remote-debugging-port
command line switch:

	chrome.exe --remote-debugging-port=9222

Then you can start a separate client Chrome instance, using a distinct user profile:

	chrome.exe --user-data-dir=<some directory>

Now you can navigate to the given port from your client and attach to any of the discovered tabs for
debugging: http://localhost:9222

You will find the Developer Tools interface identical to the embedded one and here is why:

	* When you navigate your client browser to the remote's Chrome port, Developer Tools front-end
	  is being served from the host Chrome as a Web Application from the Web Server.
	* It fetches HTML, JavaScript and CSS files over HTTP
	* Once loaded, Developer Tools establishes a Web Socket connection to its host and starts
	  exchanging JSON messages with it.

In this scenario, you can substitute Developer Tools front-end with your own implementation. Instead
of navigating to the HTML page at http://localhost:9222, your application can discover available
pages by requesting: http://localhost:9222/json and getting a JSON object with information about
inspectable pages along with the WebSocket addresses that you could use in order to start
instrumenting them.

Remote debugging is especially useful when debugging remote instances of the browser or attaching to
the embedded devices. Blink port owners are responsible for exposing debugging connections to the
external users.

Sniffing the protocol

This is especially handy to understand how the DevTools frontend makes use of the protocol. First,
run Chrome with the debugging port open:

	alias canary="/Applications/Google\ Chrome\ Canary.app/Contents/MacOS/Google\ Chrome\ Canary" canary --remote-debugging-port=9222 http://localhost:9222 http://chromium.org

Then, select the Chromium Projects item in the Inspectable Pages list. Now that DevTools is up and
fullscreen, open DevTools to inspect it. Cmd-R in the new inspector to make the first restart. Now
head to Network Panel, filter by Websocket, select the connection and click the Frames tab. Now you
can easily see the frames of WebSocket activity as you use the first instance of the DevTools.

DevTools protocol via Chrome extension

To allow chrome extensions to interact with the protocol, we introduced chrome.debugger extension
API that exposes this JSON message transport interface. As a result, you can not only attach to the
remotely running Chrome instance, but also instrument it from its own extension.

Chrome Debugger Extension API provides a higher level API where command domain, name and body are
provided explicitly in the `sendCommand`` call. This API hides request ids and handles binding of
the request with its response, hence allowing `sendCommand`` to report result in the callback
function call. One can also use this API in combination with the other Extension APIs.

If you are developing a Web-based IDE, you should implement an extension that exposes debugging
capabilities to your page and your IDE will be able to open pages with the target application, set
breakpoints there, evaluate expressions in console, live edit JavaScript and CSS, display live DOM,
network interaction and any other aspect that Developer Tools is instrumenting today.

Opening embedded Developer Tools will terminate the remote connection and thus detach the extension.
https://chromedevtools.github.io/devtools-protocol/#simultaneous

Frequently Asked Questions

How is the protocol defined

The canonical protocol definitions live in the Chromium source tree: (browser_protocol.json and
js_protocol.json). They are maintained manually by the DevTools engineering team. These files are
mirrored (hourly) on GitHub in the devtools-protocol repo.

The declarative protocol definitions are used across tools. Within Chromium, a binding layer is
created for the Chrome DevTools to interact with, and separately the protocol is used for Chrome
Headless’s C++ interface.

What’s the protocol_externs file

It’s created via generate_protocol_externs.py and useful for tools using closure compiler. The
TypeScript story is here.

Are the HTTP endpoints documented

Not yet. See bugger-daemon’s third-party docs. See also the endpoints implementation in Chromium.
/json/protocol was added in Chrome 60.

How do I access the browser target

The endpoint is exposed as webSocketDebuggerUrl in /json/version. Note the browser in the URL,
rather than page. If Chrome was launched with --remote-debugging-port=0 and chose an open port, the
browser endpoint is written to both stderr and the DevToolsActivePort file in browser profile
folder.

Does the protocol support multiple simultaneous clients

Yes, as of Chrome 63! See Multi-client remote debugging support.

Upon disconnnection, the outgoing client will receive a detached event. For example:

	{"method":"Inspector.detached","params":{"reason":"replaced_with_devtools"}}.

View the enum of possible reasons. (For reference: the original patch). After disconnection, some
apps have chosen to pause their state and offer a reconnect button.
*/
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
