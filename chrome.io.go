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
) (nil, error) {
	command := &protocol.Command{
		method: "IO.close",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Read reads a chunk of the stream.
*/
func (IO) Read(
	socket *Socket,
	params *io.ReadParams,
) (io.ReadResult, error) {
	command := &protocol.Command{
		method: "IO.read",
		params: params,
	}
	socket.SendCommand(command)
	return io.ReadResult, command.Err
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.
*/
func (IO) ResolveBlob(
	socket *Socket,
	params *io.ResolveBlobParams,
) (io.ResolveBlobResult, error) {
	command := &protocol.Command{
		method: "IO.resolveBlob",
		params: params,
	}
	socket.SendCommand(command)
	return io.ResolveBlobResult, command.Err
}
