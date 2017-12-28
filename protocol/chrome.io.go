package protocol

import (
	"encoding/json"

	io "github.com/mkenney/go-chrome/protocol/io"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
IO is a struct that provides a namespace for the Chrome IO protocol methods. The IO protocol
provides input/output operations for streams produced by DevTools.

https://chromedevtools.github.io/devtools-protocol/tot/IO/
*/
var IO = _io{}

type _io struct{}

/*
Close closes the stream and discards any temporary backing storage.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
*/
func (_io) Close(
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
func (_io) Read(
	socket sock.Socketer,
	params *io.ReadParams,
) (io.ReadResult, error) {
	command := sock.NewCommand("IO.read", params)
	result := io.ReadResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
*/
func (_io) ResolveBlob(
	socket sock.Socketer,
	params *io.ResolveBlobParams,
) (io.ResolveBlobResult, error) {
	command := sock.NewCommand("IO.resolveBlob", params)
	result := io.ResolveBlobResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}
