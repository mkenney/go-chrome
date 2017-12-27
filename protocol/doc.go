/*
Package protocol defines types and functions for communicating with the Chrome devtools websocket
protocol per https://chromedevtools.github.io/devtools-protocol/

The Chrome DevTools Protocol allows for tools to instrument, inspect, debug and profile Chromium,
Chrome and other Blink-based browsers. Many existing projects currently use the protocol. The Chrome
DevTools uses this protocol and the team maintains its API.

Instrumentation is divided into a number of domains (DOM, Debugger, Network etc.). Each domain
defines a number of commands it supports and events it generates. Both commands and events are
serialized JSON objects of a fixed structure. You can either debug over the wire using the raw
messages as they are described in the corresponding domain documentation, or use extension
JavaScript API.

Latest documentation

* https://chromedevtools.github.io/devtools-protocol/tot/

* https://chromium.googlesource.com/chromium/src/+log/master/third_party/WebKit/Source/core/inspector/browser_protocol.json

* https://chromium.googlesource.com/v8/v8/+/master/src/inspector/js_protocol.json
*/
package protocol
