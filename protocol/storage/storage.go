package Storage

import (
	"fmt"
)

/*
StorageType is an enum of possible storage types.
*/
type StorageType int

const (
	_appcache StorageType = iota
	_cookies
	_fileSystems
	_indexeddb
	_localStorage
	_shaderCache
	_websql
	_serviceWorkers
	_cacheStorage
	_all
	_other
)

func (a StorageType) String() string {
	if a == 0 {
		return "appcache"
	}
	if a == 1 {
		return "cookies"
	}
	if a == 2 {
		return "file_systems"
	}
	if a == 3 {
		return "indexeddb"
	}
	if a == 4 {
		return "local_storage"
	}
	if a == 5 {
		return "shader_cache"
	}
	if a == 6 {
		return "websql"
	}
	if a == 7 {
		return "service_workers"
	}
	if a == 8 {
		return "cache_storage"
	}
	if a == 9 {
		return "all"
	}
	if a == 10 {
		return "other"
	}
	panic(fmt.Errorf("Invalid StorageType %d", a))
}

/*
UsageForType is usage for a storage type.
*/
type UsageForType struct {
	// Name of storage type.
	StorageType StorageType `json:"storageType"`

	// Storage usage (bytes).
	Usage int `json:"usage"`
}
