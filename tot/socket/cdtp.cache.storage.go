package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cache/storage"
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
	params *storage.DeleteCacheParams,
) <-chan *storage.DeleteCacheResult {
	resultChan := make(chan *storage.DeleteCacheResult)
	command := NewCommand(protocol.Socket, "CacheStorage.deleteCache", params)
	result := &storage.DeleteCacheResult{}

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
	params *storage.DeleteEntryParams,
) <-chan *storage.DeleteEntryResult {
	resultChan := make(chan *storage.DeleteEntryResult)
	command := NewCommand(protocol.Socket, "CacheStorage.deleteEntry", params)
	result := &storage.DeleteEntryResult{}

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
	params *storage.RequestCacheNamesParams,
) <-chan *storage.RequestCacheNamesResult {
	resultChan := make(chan *storage.RequestCacheNamesResult)
	command := NewCommand(protocol.Socket, "CacheStorage.requestCacheNames", params)
	result := &storage.RequestCacheNamesResult{}

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
	params *storage.RequestCachedResponseParams,
) <-chan *storage.RequestCachedResponseResult {
	resultChan := make(chan *storage.RequestCachedResponseResult)
	command := NewCommand(protocol.Socket, "CacheStorage.requestCachedResponse", params)
	result := &storage.RequestCachedResponseResult{}

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
	params *storage.RequestEntriesParams,
) <-chan *storage.RequestEntriesResult {
	resultChan := make(chan *storage.RequestEntriesResult)
	command := NewCommand(protocol.Socket, "CacheStorage.requestEntries", params)
	result := &storage.RequestEntriesResult{}

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
