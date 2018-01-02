package socket

import "sync"

/*
Commander defines the interface for websocket commands.
*/
type Commander interface {
	// Done manages the result of a command and unblockes the sync.WaitGroup.
	Done(result []byte, err error)

	// Error returns the most recent error, if any.
	Error() error

	// ID returns the command ID
	ID() int

	// Method returns the name of the Chrom DevTools Protocol method to be
	// called.
	Method() string

	// Params returns the command parameters.
	Params() interface{}

	// Result returns the result of the command returned by the websocket.
	Result() []byte

	// WaitGroup returns the sync.WaitGroup instance.
	WaitGroup() *sync.WaitGroup
}
