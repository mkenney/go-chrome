package Chrome

import (
	"encoding/json"

	io "github.com/mkenney/go-chrome/io"
	"github.com/mkenney/go-chrome/protocol"
)

/*
IO - https://chromedevtools.github.io/devtools-protocol/tot/IO/
Input/Output operations for streams produced by DevTools.
*/
type IO struct{}

/*
Close closes the stream and discards any temporary backing storage.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
*/
func (IO) Close(
	socket *Socket,
	params *io.CloseParams,
) error {
	command := &protocol.Command{
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
	socket *Socket,
	params *io.ReadParams,
) (io.ReadResult, error) {
	command := &protocol.Command{
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
	socket *Socket,
	params *io.ResolveBlobParams,
) (io.ResolveBlobResult, error) {
	command := &protocol.Command{
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
