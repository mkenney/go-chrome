/*
Package storage provides type definitions for use with the Chrome Storage protocol

https://chromedevtools.github.io/devtools-protocol/tot/Storage/
*/
package storage

/*
Type is an enum of possible storage types.

ALLOWED VALUES
	- appcache
	- cookies
	- file_systems
	- indexeddb
	- local_storage
	- shader_cache
	- websql
	- service_workers
	- cache_storage
	- all
	- other

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-StorageType
*/
type Type string

/*
UsageForType is usage for a storage type.

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-UsageForType
*/
type UsageForType struct {
	// Name of storage type.
	Type Type `json:"storageType"`

	// Storage usage (bytes).
	Usage int `json:"usage"`
}
