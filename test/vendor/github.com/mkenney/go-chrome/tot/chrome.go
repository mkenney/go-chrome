/*
Package chrome aims to be a complete Chrome DevTools Protocol Viewer
implementation.

This version implements the Tip-of-Tree API. See
https://chromedevtools.github.io/devtools-protocol/tot/ for details.
*/
package chrome

import (
	"os"

	log "github.com/sirupsen/logrus"
)

/*
If a LOG_LEVEL environment variable exists set that value as the log level.
Useful during development.
*/
func init() {
	levelFlag := os.Getenv("LOG_LEVEL")
	if "" == levelFlag {
		levelFlag = "info"
	}
	level, err := log.ParseLevel(levelFlag)
	if nil == err {
		log.SetLevel(level)
	}
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
