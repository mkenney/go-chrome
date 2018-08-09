package chrome

import "net/url"

/*
Chromium defines an interface for interacting with Chromium based web browsers
*/
type Chromium interface {
	// Address returns the domain to use for accessing Chrome sockets (e.g.
	// 'localhost'). Should return a sane default value such as 'localhost'.
	Address() string

	// Binary returns the path to the Chromium binary. Should return a sane
	// default value such as '/usr/bin/google-chrome'.
	Binary() string

	// Close ends the Chromium process and cleans up.
	Close() error

	// GetTab returns an open Tabber instance, or an error if the requested tab
	// does not exist.
	GetTab(tabID string) (tab Tabber, err error)

	// RemoveTab removes tab reference from chrome tabs list.
	RemoveTab(tab *Tab)

	// DebuggingAddress returns the address that the remote debugging protocol
	// is available on. Should return a sane default value such as '0.0.0.0'.
	DebuggingAddress() string

	// DebuggingPort is the port number that the remote debugging protocol is
	// available on. Should return a sane default value such as 9222.
	DebuggingPort() int

	// Args returns a ChromiumFlags interface used to define and manage CLI
	// arguments to the Chromium binary. Only used when starting the chromium
	// system process.
	Flags() ChromiumFlags

	// Launch launches the Chromium process and returns the connected Chromium
	// struct.
	Launch() error

	// NewTab spawns a new tab and returns a reference to it.
	NewTab(url string) (*Tab, error)

	// Port returns the port number the developer tools endpoints will listen
	// on. Should return a sane default value such as 9222.
	Port() int

	// Query queries the developer tools endpoints and returns JSON data in the
	// provided struct.
	Query(path string, params url.Values, msg interface{}) (interface{}, error)

	// STDERR returns a string defining the location to write STDERR output.
	STDERR() string

	// STDOUT returns a string defining the location to write STDOUT output.
	STDOUT() string

	// Tabs returns the list of the currently open tabs.
	Tabs() []*Tab

	// Version returns Chromium version data.
	Version() (*Version, error)

	// Workdir returns the path of the Chromium working directory. Should return
	// a sane default value such as '/tmp/headless-chrome'.
	Workdir() string
}
