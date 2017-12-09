package IO

import (
	Runtime "app/chrome/runtime"
)

/*
CloseParams represents IO.close parameters.
*/
type CloseParams struct {
	// Handle of the stream to close.
	Handle StreamHandle `json:"handle"`
}

/*
ReadParams represents IO.read parameters.
*/
type ReadParams struct {
	// Handle of the stream to read.
	Handle StreamHandle `json:"handle"`

	// Optional. Seek to the specified offset before reading (if not specificed, proceed with offset
	// following the last read).
	Offset int `json:"offset,omitempty"`

	// Optional. Maximum number of bytes to read (left upon the agent discretion if not specified).
	Size int `json:"size,omitempty"`
}

/*
ReadResult represents the result of calls to IO.read.
*/
type ReadResult struct {
	// Set if the data is base64-encoded.
	Base64Encoded bool `json:"base64Encoded"`

	// Data that were read.
	Data string `json:"data"`

	// Set if the end-of-file condition occurred while reading.
	EOF bool `json:"eof"`
}

/*
ResolveBlobParams represents IO.resolveBlob parameters.
*/
type ResolveBlobParams struct {
	// Object ID of a Blob object wrapper.
	ObjectID Runtime.RemoteObjectID `json:"objectId"`
}

/*
ResolveBlobResult represents the result of calls to IO.resolveBlob.
*/
type ResolveBlobResult struct {
	// UUID of the specified Blob.
	UUID string `json:"uuid"`
}

/*
StreamHandle is either obtained from another method or specified as blob:<uuid> where <uuid> is an
UUID of a Blob.
*/
type StreamHandle string
