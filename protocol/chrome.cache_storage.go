package protocol

import (
	cacheStorage "github.com/mkenney/go-chrome/protocol/cache_storage"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
CacheStorage provides a namespace for the Chrome CacheStorage protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/ EXPERIMENTAL.
*/
var CacheStorage = CacheStorageProtocol{}

/*
CacheStorageProtocol defines the CacheStorage protocol methods.
*/
type CacheStorageProtocol struct{}

/*
DeleteCache deletes a cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
*/
func (CacheStorageProtocol) DeleteCache(
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
func (CacheStorageProtocol) DeleteEntry(
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
func (CacheStorageProtocol) RequestCacheNames(
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
func (CacheStorageProtocol) RequestCachedResponse(
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
func (CacheStorageProtocol) RequestEntries(
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
