package chrome

import (
	io "app/chrome/io"
	"app/chrome/protocol"
)

/*
IO domain
*/
type IO struct{}

/*
Close closes the stream and discards any temporary backing storage.
*/
func (IO) Close(socket *Socket, params *io.CloseParams) error {
	command := new(protocol.Command)
	command.method = "IO.close"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
Read reads a chunk of the stream.
*/
func (IO) Read(socket *Socket, params *io.ReadParams) error {
	command := new(protocol.Command)
	command.method = "IO.read"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.
*/
func (IO) ResolveBlob(socket *Socket, params *io.ResolveBlobParams) error {
	command := new(protocol.Command)
	command.method = "IO.resolveBlob"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
