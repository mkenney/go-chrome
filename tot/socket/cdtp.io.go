package socket

import (
	"encoding/json"

	io "github.com/mkenney/go-chrome/tot/cdtp/io"
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
) chan *io.CloseResult {
	resultChan := make(chan *io.CloseResult)
	command := NewCommand(protocol.Socket, "IO.close", params)
	result := &io.CloseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Read reads a chunk of the stream.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
*/
func (protocol *IOProtocol) Read(
	params *io.ReadParams,
) chan *io.ReadResult {
	resultChan := make(chan *io.ReadResult)
	command := NewCommand(protocol.Socket, "IO.read", params)
	result := &io.ReadResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
ResolveBlob returns the UUID of Blob object specified by a remote object id.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
*/
func (protocol *IOProtocol) ResolveBlob(
	params *io.ResolveBlobParams,
) chan *io.ResolveBlobResult {
	resultChan := make(chan *io.ResolveBlobResult)
	command := NewCommand(protocol.Socket, "IO.resolveBlob", params)
	result := &io.ResolveBlobResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}
