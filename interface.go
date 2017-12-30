package chrome

import (
	"github.com/mkenney/go-chrome/socket"
)

/*
Chromium defines an interface for interacting with Chromium based web browsers
*/
type Chromium interface {
	// Address returns the domain to use for accessing Chrome sockets (e.g.
	// 'localhost'). Should return a sane default value such as 'localhost'.
	Address() string

	// Args returns a Commander interface used to define and manage CLI
	// arguments to the Chromium binary. Only used when launching the binary
	// from the API.
	Args() Commander

	// Binary returns the path to the Chromium binary. Should return a sane
	// default value such as '/usr/bin/google-chrome'.
	Binary() string

	// Close ends the Chromium process and cleans up.
	Close() error

	// DebuggingAddress returns the address that the remote debugging protocol
	// is available on. Should return a sane default value such as '0.0.0.0'.
	DebuggingAddress() string

	// DebuggingPort is the port number that the remote debugging protocol is
	// available on. Should return a sane default value such as 9222.
	DebuggingPort() int

	// Launch launches the Chromium process and returns the connected Chromium
	// struct.
	Launch() error

	// NewTab spawns a new tab and returns a reference to it.
	NewTab(url string) (*Tab, error)

	// Output returns the path to a file to be used to capture STDOUT and STDERR
	// output. Defaut value should be an empty string, in which case all output
	// is delivered to STDOUT.
	Output() string

	// Port returns the port number the developer tools endpoints will listen
	// on. Should return a sane default value such as 9222.
	Port() int

	// Tabs returns the list of the currently open tabs.
	Tabs() []*Tab

	// Version returns Chromium version data.
	Version() (*Version, error)

	// Workdir returns the path of the Chromium working directory. Should return
	// a sane default value such as '/tmp/headless-chrome'.
	Workdir() string
}

/*
Commander provides an interface for managing CLI arguments to the Chromium binary.
*/
type Commander interface {

	// Get returns the specified argument values
	Get(arg string) ([]interface{}, error)

	// Has checks to see if an argument is present.
	Has(arg string) bool

	// List returns an array of each argument for use in os.StartProcess
	List() []string

	// Set sets a CLI argument's values.
	Set(arg string, values []interface{}) error

	// String implments Stringer. It returns the set parameters formatted to be
	// passed to the command line.
	String() string
}

/*
Tabber provides an interface for managing a Chromium tab
*/
type Tabber interface {
	// Browser returns the Chromium instance this tab is in
	Browser() *Browser

	// Close closes this chromium tab
	Close() (interface{}, error)

	// Data returns the tab metadata
	Data() *TabData

	// Socket returns the socket.Socketer interface for this tab
	Socket() socket.Socketer
}
