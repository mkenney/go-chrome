package storage

import (
	"encoding/json"
	"fmt"
)

type typeEnum struct {
	Appcache       TypeEnum
	Cookies        TypeEnum
	FileSystems    TypeEnum
	Indexeddb      TypeEnum
	LocalStorage   TypeEnum
	ShaderCache    TypeEnum
	Websql         TypeEnum
	ServiceWorkers TypeEnum
	CacheStorage   TypeEnum
	All            TypeEnum
	Other          TypeEnum
}

/*
Type provides named acces to the TypeEnum values.
*/
var Type = typeEnum{
	Appcache:       typeAppcache,
	Cookies:        typeCookies,
	FileSystems:    typeFileSystems,
	Indexeddb:      typeIndexeddb,
	LocalStorage:   typeLocalStorage,
	ShaderCache:    typeShaderCache,
	Websql:         typeWebsql,
	ServiceWorkers: typeServiceWorkers,
	CacheStorage:   typeCacheStorage,
	All:            typeAll,
	Other:          typeOther,
}

/*
TypeEnum represents possible storage types. Allowed values:
	- Type.Appcache       "appcache"
	- Type.Cookies        "cookies"
	- Type.FileSystems    "file_systems"
	- Type.Indexeddb      "indexeddb"
	- Type.LocalStorage   "local_storage"
	- Type.ShaderCache    "shader_cache"
	- Type.Websql         "websql"
	- Type.ServiceWorkers "service_workers"
	- Type.CacheStorage   "cache_storage"
	- Type.All            "all"
	- Type.Other          "other"

https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-StorageType
*/
type TypeEnum int

/*
String implements Stringer
*/
func (enum TypeEnum) String() string {
	return _typeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum TypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *TypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _typeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// typeAppcache represents the "appcache" value.
	typeAppcache TypeEnum = iota + 1
	// typeCookies represents the "cookies" value.
	typeCookies
	// typeFileSystems represents the "file_systems" value.
	typeFileSystems
	// typeIndexeddb represents the "indexeddb" value.
	typeIndexeddb
	// typeLocalStorage represents the "local_storage" value.
	typeLocalStorage
	// typeShaderCache represents the "shader_cache" value.
	typeShaderCache
	// typeWebsql represents the "websql" value.
	typeWebsql
	// typeServiceWorkers represents the "service_workers" value.
	typeServiceWorkers
	// typeCacheStorage represents the "cache_storage" value.
	typeCacheStorage
	// typeAll represents the "all" value.
	typeAll
	// typeOther represents the "other" value.
	typeOther
)

var _typeEnums = map[TypeEnum]string{
	typeAppcache:       "appcache",
	typeCookies:        "cookies",
	typeFileSystems:    "file_systems",
	typeIndexeddb:      "indexeddb",
	typeLocalStorage:   "local_storage",
	typeShaderCache:    "shader_cache",
	typeWebsql:         "websql",
	typeServiceWorkers: "service_workers",
	typeCacheStorage:   "cache_storage",
	typeAll:            "all",
	typeOther:          "other",
}
