package socket

import (
	"encoding/json"

	cacheStorage "github.com/mkenney/go-chrome/tot/cdtp/cache/storage"
)

/*
CacheStorageProtocol provides a namespace for the Chrome CacheStorage protocol
methods.

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
) <-chan *cacheStorage.DeleteCacheResult {
	resultChan := make(chan *cacheStorage.DeleteCacheResult)
	command := NewCommand(protocol.Socket, "CacheStorage.deleteCache", params)
	result := &cacheStorage.DeleteCacheResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
DeleteEntry deletes a cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
*/
func (protocol *CacheStorageProtocol) DeleteEntry(
	params *cacheStorage.DeleteEntryParams,
) <-chan *cacheStorage.DeleteEntryResult {
	resultChan := make(chan *cacheStorage.DeleteEntryResult)
	command := NewCommand(protocol.Socket, "CacheStorage.deleteEntry", params)
	result := &cacheStorage.DeleteEntryResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RequestCacheNames requests cache names.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
*/
func (protocol *CacheStorageProtocol) RequestCacheNames(
	params *cacheStorage.RequestCacheNamesParams,
) <-chan *cacheStorage.RequestCacheNamesResult {
	resultChan := make(chan *cacheStorage.RequestCacheNamesResult)
	command := NewCommand(protocol.Socket, "CacheStorage.requestCacheNames", params)
	result := &cacheStorage.RequestCacheNamesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RequestCachedResponse fetches cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
*/
func (protocol *CacheStorageProtocol) RequestCachedResponse(
	params *cacheStorage.RequestCachedResponseParams,
) <-chan *cacheStorage.RequestCachedResponseResult {
	resultChan := make(chan *cacheStorage.RequestCachedResponseResult)
	command := NewCommand(protocol.Socket, "CacheStorage.requestCachedResponse", params)
	result := &cacheStorage.RequestCachedResponseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RequestEntries requests data from cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
*/
func (protocol *CacheStorageProtocol) RequestEntries(
	params *cacheStorage.RequestEntriesParams,
) <-chan *cacheStorage.RequestEntriesResult {
	resultChan := make(chan *cacheStorage.RequestEntriesResult)
	command := NewCommand(protocol.Socket, "CacheStorage.requestEntries", params)
	result := &cacheStorage.RequestEntriesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}
