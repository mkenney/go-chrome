/*
Package schema provides type definitions for use with the Chrome Schema protocol

https://chromedevtools.github.io/devtools-protocol/tot/Schema/
*/
package schema

/*
Domain is a description of the protocol domain.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/#type-Domain
*/
type Domain struct {
	// Domain name.
	Name string `json:"name"`

	// Domain version.
	Version string `json:"version"`
}
