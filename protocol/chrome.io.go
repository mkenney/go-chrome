package protocol

import (
	io "github.com/mkenney/go-chrome/protocol/io"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
IO provides a namespace for the Chrome IO protocol methods. The IO protocol provides input/output
operations for streams produced by DevTools.

https://chromedevtools.github.io/devtools-protocol/tot/IO/
*/
var IO = IOProtocol{}

/*
IOProtocol defines the IO protocol methods.
*/
type IOProtocol struct{}

/*
Close closes the stream and discards any temporary backing storage.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
*/
func (IOProtocol) Close(
	socket sock.Socketer,
	params *io.CloseParams,
) error {
	command := sock.NewCommand("IO.close", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Read reads a chunk of the stream.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
*/
func (IOProtocol) Read(
	socket sock.Socketer,
	params *io.ReadParams,
) (*io.ReadResult, error) {
	command := sock.NewCommand("IO.read", params)
	result := &io.ReadResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
*/
func (IOProtocol) ResolveBlob(
	socket sock.Socketer,
	params *io.ResolveBlobParams,
) (*io.ResolveBlobResult, error) {
	command := sock.NewCommand("IO.resolveBlob", params)
	result := &io.ResolveBlobResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
