package chrome

import (
	"net/url"

	"github.com/mkenney/go-chrome/socket"
)

/*
Tabber provides an interface for managing a Chromium tab
*/
type Tabber interface {
	// Browser returns the Chromium instance this tab is in
	Chromium() *Chrome

	// Close closes this chromium tab
	Close() (interface{}, error)

	// Data returns the tab metadata
	Data() *TabData

	// Protocol returns the socket.Protocoller interface for this tab
	Protocol() socket.Protocoller

	// Socket returns the socket.Socketer interface for this tab
	Socket() socket.Socketer

	// URL returns the URL of the websocket connection
	URL() *url.URL
}
