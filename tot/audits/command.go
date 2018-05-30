package audits

import (
	"github.com/mkenney/go-chrome/tot/network"
)

/*
GetEncodedResponseParams represents Audits.getEncodedResponse parameters

https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
*/
type GetEncodedResponseParams struct {
	// Identifier of the network request to get content for.
	RequestID network.RequestID `json:"requestId"`

	// The encoding to use. Allowed values:
	//	- Encoding.Webp
	//	- Encoding.Jpeg
	//	- Encoding.Png
	Encoding EncodingEnum `json:"encoding"`

	// Optional. The quality of the encoding (0-1). (defaults to 1).
	Quality float64 `json:"quality,omitempty"`

	// Optional Whether to only return the size information (defaults to false).
	SizeOnly bool `json:"sizeOnly,omitempty"`
}

/*
GetEncodedResponseResult represents the result of calls to Audits.getEncodedResponse.

https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
*/
type GetEncodedResponseResult struct {
	// Optional. The encoded body as a base64 string. Omitted if sizeOnly is
	// true.
	Body string `json:"body,omitempty"`

	// Size before re-encoding.
	OriginalSize int64 `json:"originalSize"`

	// Size after re-encoding.
	EncodedSize int64 `json:"encodedSize"`

	// Error information related to executing this method
	Err error `json:"-"`
}
