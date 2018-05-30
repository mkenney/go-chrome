package io

import (
	"github.com/mkenney/go-chrome/tot/runtime"
)

/*
CloseParams represents IO.close parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
*/
type CloseParams struct {
	// Handle of the stream to close.
	Handle StreamHandle `json:"handle"`
}

/*
CloseResult represents the result of calls to IO.close.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
*/
type CloseResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ReadParams represents IO.read parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
*/
type ReadParams struct {
	// Handle of the stream to read.
	Handle StreamHandle `json:"handle"`

	// Optional. Seek to the specified offset before reading (if not specificed,
	// proceed with offset following the last read).
	Offset int `json:"offset,omitempty"`

	// Optional. Maximum number of bytes to read (left upon the agent discretion
	// if not specified).
	Size int `json:"size,omitempty"`
}

/*
ReadResult represents the result of calls to IO.read.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
*/
type ReadResult struct {
	// Set if the data is base64-encoded.
	Base64Encoded bool `json:"base64Encoded"`

	// Data that were read.
	Data string `json:"data"`

	// Set if the end-of-file condition occurred while reading.
	EOF bool `json:"eof"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ResolveBlobParams represents IO.resolveBlob parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
*/
type ResolveBlobParams struct {
	// Object ID of a Blob object wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
}

/*
ResolveBlobResult represents the result of calls to IO.resolveBlob.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
*/
type ResolveBlobResult struct {
	// UUID of the specified Blob.
	UUID string `json:"uuid"`

	// Error information related to executing this method
	Err error `json:"-"`
}
