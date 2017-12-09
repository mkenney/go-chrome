package Audits

import (
	Network "app/chrome/network"
)

/*
GetEncodedResponseParams represents Audits.getEncodedResponse parameters
*/
type GetEncodedResponseParams struct {
	// Identifier of the network request to get content for.
	RequestID Network.RequestID `json:"requestId"`

	// The encoding to use. Allowed values: webp, jpeg, png.
	Encoding string `json:"encoding"`

	// Optional. The quality of the encoding (0-1). (defaults to 1).
	Quality float64 `json:"quality,omitempty"`

	// Optional Whether to only return the size information (defaults to false).
	SizeOnly bool `json:"sizeOnly,omitempty"`
}

/*
GetEncodedResponseResult represents the result of calls to Audits.getEncodedResponse.
*/
type GetEncodedResponseResult struct {
	// Optional. The encoded body as a base64 string. Omitted if sizeOnly is true.
	Body string `json:"body,omitempty"`

	// Size before re-encoding.
	OriginalSize int `json:"originalSize"`

	// Size after re-encoding.
	EncodedSize int `json:"encodedSize"`
}
