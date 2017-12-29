package protocol

import (
	"encoding/json"

	cacheStorage "github.com/mkenney/go-chrome/protocol/cache_storage"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
CacheStorage is a struct that provides a namespace for the Chrome CacheStorage protocol methods.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/
*/
var CacheStorage = _cachestorage{}

type _cachestorage struct{}

/*
DeleteCache deletes a cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
*/
func (_cachestorage) DeleteCache(
	socket sock.Socketer,
	params *cacheStorage.DeleteCacheParams,
) error {
	command := sock.NewCommand("CacheStorage.deleteCache", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DeleteEntry deletes a cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
*/
func (_cachestorage) DeleteEntry(
	socket sock.Socketer,
	params *cacheStorage.DeleteEntryParams,
) error {
	command := sock.NewCommand("CacheStorage.deleteEntry", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RequestCacheNames requests cache names.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
*/
func (_cachestorage) RequestCacheNames(
	socket sock.Socketer,
	params *cacheStorage.RequestCacheNamesParams,
) (*cacheStorage.RequestCacheNamesResult, error) {
	command := sock.NewCommand("CacheStorage.requestCacheNames", params)
	result := &cacheStorage.RequestCacheNamesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
RequestCachedResponse fetches cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
*/
func (_cachestorage) RequestCachedResponse(
	socket sock.Socketer,
	params *cacheStorage.RequestCachedResponseParams,
) (*cacheStorage.RequestCachedResponseResult, error) {
	command := sock.NewCommand("CacheStorage.requestCachedResponse", params)
	result := &cacheStorage.RequestCachedResponseResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
RequestEntries requests data from cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
*/
func (_cachestorage) RequestEntries(
	socket sock.Socketer,
	params *cacheStorage.RequestEntriesParams,
) (*cacheStorage.RequestEntriesResult, error) {
	command := sock.NewCommand("CacheStorage.requestEntries", params)
	result := &cacheStorage.RequestEntriesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
	}

	return result, command.Error()
}
