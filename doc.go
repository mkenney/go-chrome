/*
Package gochrome aims to be a complete Chrome DevTools Protocol Viewer
implementation.

Versions

Versioned packages are available. Curently the only version is `tot` or
Tip-of-Tree. Stable versions will be made available in the future.

Work in progress

You should probably ignore this for now.

Official documentation

See https://chromedevtools.github.io/devtools-protocol/

The Chrome DevTools Protocol allows for tools to instrument, inspect, debug and
profile Chromium, Chrome and other Blink-based browsers. Many existing projects
currently use the protocol. The Chrome DevTools uses this protocol and the team
maintains its API.

Instrumentation is divided into a number of domains (DOM, Debugger, Network
etc.). Each domain defines a number of commands it supports and events it
generates. Both commands and events are serialized JSON objects of a fixed
structure. You can either debug over the wire using the raw messages as they are
described in the corresponding domain documentation, or use extension JavaScript
API.

Protocol API Docs

The latest (tip-of-tree) protocol (tot)

It changes frequently and can break at any time. However it captures the full
capabilities of the Protocol, whereas the stable release is a subset. There is
no backwards compatibility support guaranteed for the capabilities it
introduces.

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

The Developer Tools front-end can attach to a remotely running Chrome instance
for debugging. For this scenario to work, you should start your host Chrome
instance with the remote-debugging-port command line switch:

	chrome.exe --remote-debugging-port=9222

Then you can start a separate client Chrome instance, using a distinct user
profile:

	chrome.exe --user-data-dir=<some directory>

Now you can navigate to the given port from your client and attach to any of the
discovered tabs for debugging: http://localhost:9222

You will find the Developer Tools interface identical to the embedded one and
here is why:

	* When you navigate your client browser to the remote's Chrome port,
	Developer Tools front-end is being served from the host Chrome as a Web
	Application from the Web Server.
	* It fetches HTML, JavaScript and CSS files over HTTP
	* Once loaded, Developer Tools establishes a Web Socket connection to its
	host and starts exchanging JSON messages with it.

In this scenario, you can substitute Developer Tools front-end with your own
implementation. Instead of navigating to the HTML page at http://localhost:9222,
your application can discover available pages by requesting: http://localhost:9222/json
and getting a JSON object with information about inspectable pages along with
the WebSocket addresses that you could use in order to start instrumenting them.

Remote debugging is especially useful when debugging remote instances of the
browser or attaching to the embedded devices. Blink port owners are responsible
for exposing debugging connections to the external users.

Sniffing the protocol

This is especially handy to understand how the DevTools frontend makes use of
the protocol. First, run Chrome with the debugging port open:

	alias canary="/Applications/Google\ Chrome\ Canary.app/Contents/MacOS/Google\ Chrome\ Canary" canary --remote-debugging-port=9222 http://localhost:9222 http://chromium.org

Then, select the Chromium Projects item in the Inspectable Pages list. Now that
DevTools is up and fullscreen, open DevTools to inspect it. Cmd-R in the new
inspector to make the first restart. Now head to Network Panel, filter by
Websocket, select the connection and click the Frames tab. Now you can easily
see the frames of WebSocket activity as you use the first instance of the
DevTools.

DevTools protocol via Chrome extension

To allow chrome extensions to interact with the protocol, we introduced chrome.debugger
extension API that exposes this JSON message transport interface. As a result,
you can not only attach to the remotely running Chrome instance, but also
instrument it from its own extension.

Chrome Debugger Extension API provides a higher level API where command domain,
name and body are provided explicitly in the `sendCommand` call. This API hides
request ids and handles binding of the request with its response, hence allowing
`sendCommand` to report result in the callback function call. One can also use
this API in combination with the other Extension APIs.

If you are developing a Web-based IDE, you should implement an extension that
exposes debugging capabilities to your page and your IDE will be able to open
pages with the target application, set breakpoints there, evaluate expressions
in console, live edit JavaScript and CSS, display live DOM, network interaction
and any other aspect that Developer Tools is instrumenting today.

Opening embedded Developer Tools will terminate the remote connection and thus
detach the extension. https://chromedevtools.github.io/devtools-protocol/#simultaneous

Frequently Asked Questions

How is the protocol defined

The canonical protocol definitions live in the Chromium source tree:
(browser_protocol.json and js_protocol.json). They are maintained manually by
the DevTools engineering team. These files are mirrored (hourly) on GitHub in
the devtools-protocol repo.

The declarative protocol definitions are used across tools. Within Chromium, a
binding layer is created for the Chrome DevTools to interact with, and
separately the protocol is used for Chrome Headless’s C++ interface.

What’s the protocol_externs file

It’s created via generate_protocol_externs.py and useful for tools using closure
compiler. The TypeScript story is here.

Are the HTTP endpoints documented

Not yet. See bugger-daemon’s third-party docs. See also the endpoints
implementation in Chromium. /json/protocol was added in Chrome 60.

How do I access the browser target

The endpoint is exposed as webSocketDebuggerUrl in /json/version. Note the
browser in the URL, rather than page. If Chrome was launched with
--remote-debugging-port=0 and chose an open port, the browser endpoint is
written to both stderr and the DevToolsActivePort file in browser profile
folder.

Does the protocol support multiple simultaneous clients

Yes, as of Chrome 63! See Multi-client remote debugging support.

Upon disconnnection, the outgoing client will receive a detached event. For
example:

	{"method":"Inspector.detached","params":{"reason":"replaced_with_devtools"}}.

View the enum of possible reasons. (For reference: the original patch). After
disconnection, some apps have chosen to pause their state and offer a reconnect
button.
*/
package gochrome
