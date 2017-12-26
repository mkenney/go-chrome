package protocol

import (
	"encoding/json"

	io "github.com/mkenney/go-chrome/protocol/io"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
IO is a struct that provides a namespace for the Chrome IO protocol methods.

The IO protocol provides input/output operations for streams produced by DevTools.

https://chromedevtools.github.io/devtools-protocol/tot/IO/
*/
type IO struct{}

/*
Close closes the stream and discards any temporary backing storage.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
*/
func (IO) Close(
	socket *sock.Socket,
	params *io.CloseParams,
) error {
	command := &sock.Command{
		Method: "IO.close",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Read reads a chunk of the stream.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
*/
func (IO) Read(
	socket *sock.Socket,
	params *io.ReadParams,
) (io.ReadResult, error) {
	command := &sock.Command{
		Method: "IO.read",
		Params: params,
	}
	result := io.ReadResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
*/
func (IO) ResolveBlob(
	socket *sock.Socket,
	params *io.ResolveBlobParams,
) (io.ResolveBlobResult, error) {
	command := &sock.Command{
		Method: "IO.resolveBlob",
		Params: params,
	}
	result := io.ResolveBlobResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}
