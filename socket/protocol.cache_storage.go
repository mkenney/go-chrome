package socket

import (
	"encoding/json"

	cacheStorage "github.com/mkenney/go-chrome/protocol/cache_storage"
)

/*
CacheStorageProtocol provides a namespace for the Chrome CacheStorage protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/ EXPERIMENTAL.
*/
type CacheStorageProtocol struct {
	Socket Socketer
}

/*
DeleteCache deletes a cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
*/
func (protocol *CacheStorageProtocol) DeleteCache(
	params *cacheStorage.DeleteCacheParams,
) error {
	command := NewCommand("CacheStorage.deleteCache", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DeleteEntry deletes a cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
*/
func (protocol *CacheStorageProtocol) DeleteEntry(
	params *cacheStorage.DeleteEntryParams,
) error {
	command := NewCommand("CacheStorage.deleteEntry", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RequestCacheNames requests cache names.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
*/
func (protocol *CacheStorageProtocol) RequestCacheNames(
	params *cacheStorage.RequestCacheNamesParams,
) (*cacheStorage.RequestCacheNamesResult, error) {
	command := NewCommand("CacheStorage.requestCacheNames", params)
	result := &cacheStorage.RequestCacheNamesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
RequestCachedResponse fetches cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
*/
func (protocol *CacheStorageProtocol) RequestCachedResponse(
	params *cacheStorage.RequestCachedResponseParams,
) (*cacheStorage.RequestCachedResponseResult, error) {
	command := NewCommand("CacheStorage.requestCachedResponse", params)
	result := &cacheStorage.RequestCachedResponseResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
RequestEntries requests data from cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
*/
func (protocol *CacheStorageProtocol) RequestEntries(
	params *cacheStorage.RequestEntriesParams,
) (*cacheStorage.RequestEntriesResult, error) {
	command := NewCommand("CacheStorage.requestEntries", params)
	result := &cacheStorage.RequestEntriesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}
