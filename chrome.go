/*
Package chrome aims to be a complete Chrome DevTools Protocol Viewer implementation

Work in progress

You should probably ignore this for now.

Official documentation

See https://chromedevtools.github.io/devtools-protocol/

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
		can also be used for concerns with the protocol. It also hosts the
		canonical copy of the json files.
	* Getting Started with Headless Chrome https://developers.google.com/web/updates/2017/04/headless-chrome
	* Headless Chromium readme https://chromium.googlesource.com/chromium/src/+/lkgr/headless/README.md
	* chrome-remote-interface node module https://github.com/cyrus-and/chrome-remote-interface/
		Wiki https://github.com/cyrus-and/chrome-remote-interface/wiki
		Issue tracker https://github.com/cyrus-and/chrome-remote-interface/issues?utf8=%E2%9C%93&q=is%3Aissue%20
	*  awesome-chrome-devtools page https://github.com/ChromeDevTools/awesome-chrome-devtools#chrome-devtools-protocol
		links to many of the tools in the protocol ecosystem, including protocol
		API libraries in JavaScript, TypeScript, Python, Java, and Go. Many
		applications and libraries already use the protocol.


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

	* When you navigate your client browser to the remote's Chrome port,
	Developer Tools front-end is being served from the host Chrome as a Web
	Application from the Web Server.
	* It fetches HTML, JavaScript and CSS files over HTTP
	* Once loaded, Developer Tools establishes a Web Socket connection to its
	host and starts exchanging JSON messages with it.

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
	"time"

	log "github.com/sirupsen/logrus"
)

/*
New returns a pointer to a Browser struct
*/
func New(
	args Commander,
	binary string,
	output string,
	workdir string,
) Chromium {
	return &Browser{
		args:    args,
		binary:  binary,
		output:  output,
		workdir: workdir,
	}
}

/*
Browser is a struct that manages the Chromium process
*/
type Browser struct {
	// Optional. address is the domain to use for accessing Chromium sockets (e.g. 'localhost')
	// Defaults to 'localhost'.
	//address string

	// args contains CLI arguments for the Chromium binary.
	args Commander

	// Optional. binary is the path to the Chromium binary. Defaults to
	// '/usr/bin/google-chrome'.
	binary string

	// Optional. debuggingAddress is the address number that the remote
	// debugging protocol will be available on. Defaults to '0.0.0.0'.
	//debuggingAddress string

	// Optional. debuggingPort is the port number that the remote debugging
	// protocol will be available on. Defaults to 9222.
	//debuggingPort int

	// Optional. output is a path to a file to be used to capture STDOUT and
	// STDERR output. Defaults to the system STDOUT.
	output string

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

	// outputFile is a pointer to a file handle to be used to capture STDOUT and
	// STDERR output.
	outputFile *os.File

	// process is a pointer to the os.Process struct containing the process PID.
	process *os.Process
}

/*
Address implements Chromium.
*/
func (browser *Browser) Address() string {
	values, err := browser.args.Get("address")
	if nil != err {
		return ""
	}
	return values[0].(string)
}

/*
Args implements Chromium.
*/
func (browser *Browser) Args() Commander {
	return browser.args
}

/*
Binary implements Chromium.
*/
func (browser *Browser) Binary() string {
	return browser.binary
}

/*
Close implements Chromium.
*/
func (browser *Browser) Close() error {
	if browser.process != nil {
		if err := browser.process.Signal(os.Interrupt); err != nil {
			return err
		}
		ps, err := browser.process.Wait()
		if err != nil {
			return err
		}
		log.Infof("Chromium exited: %s", ps.String())
	}
	if browser.outputFile != nil {
		browser.outputFile.Close()
	}
	return nil
}

/*
DebuggingAddress implements Chromium.
*/
func (browser *Browser) DebuggingAddress() string {
	values, err := browser.args.Get("remote-debugging-address")
	if nil != err {
		return ""
	}
	return values[0].(string)
}

/*
DebuggingPort implements Chromium.
*/
func (browser *Browser) DebuggingPort() int {
	values, err := browser.args.Get("remote-debugging-port")
	if nil != err {
		return 0
	}
	return values[0].(int)
}

/*
Launch implements Chromium.
*/
func (browser *Browser) Launch() error {
	var err error

	if "" == browser.Binary() {
		browser.binary = "/usr/bin/google-chrome"
	}

	if !browser.Args().Has("addr") {
		browser.Args().Set("addr", []interface{}{"localhost"})
	}
	if !browser.Args().Has("remote-debugging-address") {
		browser.Args().Set("remote-debugging-address", []interface{}{"0.0.0.0"})
	}
	if !browser.Args().Has("remote-debugging-port") {
		browser.Args().Set("remote-debugging-port", []interface{}{9222})
	}
	if !browser.Args().Has("port") {
		browser.Args().Set("port", []interface{}{9222})
	}
	if !browser.Args().Has("user-data-dir") {
		browser.Args().Set("user-data-dir", []interface{}{os.TempDir()})
	}

	if "" == browser.Workdir() {
		browser.workdir = filepath.Join(os.TempDir(), "headless-chrome")
	}
	if err := os.MkdirAll(browser.Workdir(), 0700); err != nil {
		return fmt.Errorf("Cannot create working directory '%s'", browser.Workdir())
	}

	if "" == browser.Output() {
		browser.outputFile = os.Stdout
	} else {
		browser.outputFile, err = os.OpenFile(browser.Output(), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return fmt.Errorf("Cannot open output file '%s'", browser.Output())
		}
	}

	log.Infof("Starting process: %s %s", browser.Binary(), browser.Args())
	var procAttributes os.ProcAttr
	procAttributes.Dir = browser.Workdir()
	procAttributes.Files = []*os.File{nil, browser.outputFile, browser.outputFile}
	browser.process, err = os.StartProcess(
		browser.Binary(),
		browser.Args().List(),
		&procAttributes,
	)
	if err != nil {
		browser.outputFile.Close()
		return err
	}

	// Wait up to 10 seconds for Chromium to start
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if err = browser.checkVersion(); nil == err {
			break
		}
	}
	if err != nil {
		log.Errorf("Chromium took too long to start: %s", err.Error())
		browser.Close()
		return err
	}

	return nil
}

/*
Output implements Chromium.
*/
func (browser *Browser) Output() string {
	return browser.output
}

/*
Port implements Chromium.
*/
func (browser *Browser) Port() int {
	values, err := browser.args.Get("port")
	if nil != err {
		return 0
	}
	return values[0].(int)
}

/*
Tabs implements Chromium.
*/
func (browser *Browser) Tabs() []*Tab {
	return browser.tabs
}

/*
Version implements Chromium.
*/
func (browser *Browser) Version() (*Version, error) {
	if "" == browser.version.Browser {
		if _, err := browser.Cmd("/json/version", url.Values{}, &browser.version); err != nil {
			return nil, err
		}
	}
	return browser.version, nil
}

/*
Workdir implements Chromium.
*/
func (browser *Browser) Workdir() string {
	return browser.workdir
}

/*
Version is a struct representing the Chromium version information.
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
GetTab returns an open tab instance.
*/
func (browser *Browser) GetTab(tabID string) (tab *Tab, err error) {
	for _, tab = range browser.tabs {
		if tab.Data().ID == tabID {
			return tab, nil
		}
	}
	err = fmt.Errorf("Tab '%s' not found", tabID)
	return
}

func (browser *Browser) checkVersion() error {
	if _, err := browser.Cmd("/json/version", url.Values{}, &browser.version); err != nil {
		return err
	}
	log.Infof("Chromium protocol version: %s", browser.version.ProtocolVersion)
	return nil
}

/*
Cmd queries the developer tools endpoints and returns JSON data in the provided struct.
*/
func (browser *Browser) Cmd(path string, params url.Values, msg interface{}) (interface{}, error) {
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	uri := fmt.Sprintf("http://%s:%d%s", browser.Address(), browser.Port(), path)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Infof("chrome:/%s %s", path, resp.Status)
	if 200 != resp.StatusCode {
		return resp.Status, nil
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else if err := json.Unmarshal(content, &msg); err != nil {
		return content, nil
	}

	return msg, nil
}
