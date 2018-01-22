/*
Package storage provides type definitions for use with the Chrome Storage protocol

https://chromedevtools.github.io/devtools-protocol/tot/Storage/
*/
package storage

/*
UsageForType is usage for a storage type.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-UsageForType
*/
type UsageForType struct {
	// Name of storage type.
	Type TypeEnum `json:"storageType"`

	// Storage usage (bytes).
	Usage int `json:"usage"`
}
