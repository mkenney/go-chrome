package chrome

import (
	io "app/chrome/io"
	"app/chrome/protocol"
)

/*
IO - https://chromedevtools.github.io/devtools-protocol/tot/IO/
Input/Output operations for streams produced by DevTools.
*/
type IO struct{}

/*
Close closes the stream and discards any temporary backing storage.
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
*/
func (IO) Read(
	socket *Socket,
	params *io.ReadParams,
) (io.ReadResult, error) {
	command := &protocol.Command{
		Method: "IO.read",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(io.ReadResult), command.Err
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.
*/
func (IO) ResolveBlob(
	socket *Socket,
	params *io.ResolveBlobParams,
) (io.ResolveBlobResult, error) {
	command := &protocol.Command{
		Method: "IO.resolveBlob",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(io.ResolveBlobResult), command.Err
}
