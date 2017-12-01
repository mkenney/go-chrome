package IO

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
ResolveBlobParams represents IO.resolveBlob parameters.
*/
type ResolveBlobParams struct {
	// Object ID of a Blob object wrapper.
	ObjectID Runtime.RemoteObjectID `json:"objectId"`
}

/*
StreamHandle is either obtained from another method or specifed as blob:<uuid> where <uuid> is an
UUID of a Blob.
*/
type StreamHandle string
