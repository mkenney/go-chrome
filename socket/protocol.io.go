package socket

import (
	"encoding/json"

	io "github.com/mkenney/go-chrome/protocol/io"
)

/*
IOProtocol provides a namespace for the Chrome IO protocol methods. The IO
protocol provides input/output operations for streams produced by DevTools.

https://chromedevtools.github.io/devtools-protocol/tot/IO/
*/
type IOProtocol struct {
	Socket Socketer
}

/*
Close closes the stream and discards any temporary backing storage.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
*/
func (protocol *IOProtocol) Close(
	params *io.CloseParams,
) error {
	command := NewCommand("IO.close", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Read reads a chunk of the stream.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
*/
func (protocol *IOProtocol) Read(
	params *io.ReadParams,
) (*io.ReadResult, error) {
	command := NewCommand("IO.read", params)
	result := &io.ReadResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
*/
func (protocol *IOProtocol) ResolveBlob(
	params *io.ResolveBlobParams,
) (*io.ResolveBlobResult, error) {
	command := NewCommand("IO.resolveBlob", params)
	result := &io.ResolveBlobResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}
