package Audits

import (
	Network "app/chrome/protocol/network"
)

/*
GetEncodedResponseParams represents Audits.getEncodedResponse parameters
*/
type GetEncodedResponseParams struct {
	// Identifier of the network request to get content for.
	RequestID Network.RequestID `json:"requestId"`

	// The encoding to use. Allowed values: webp, jpeg, png.
	Encoding string `json:"encoding"`

	// The quality of the encoding (0-1). (defaults to 1).
	Quality int `json:"quality"`

	// Whether to only return the size information (defaults to false).
	SizeOnly bool `json:"sizeOnly"`
}
