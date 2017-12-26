package protocol

import (
	"encoding/json"

	cacheStorage "github.com/mkenney/go-chrome/protocol/cache_storage"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
CacheStorage is a struct that provides a namespace for the Chrome CacheStorage protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/
*/
type CacheStorage struct{}

/*
DeleteCache deletes a cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
*/
func (CacheStorage) DeleteCache(
	socket *sock.Socket,
	params *cacheStorage.DeleteCacheParams,
) error {
	command := &sock.Command{
		Method: "CacheStorage.deleteCache",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteEntry deletes a cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
*/
func (CacheStorage) DeleteEntry(
	socket *sock.Socket,
	params *cacheStorage.DeleteEntryParams,
) error {
	command := &sock.Command{
		Method: "CacheStorage.deleteEntry",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestCacheNames requests cache names.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
*/
func (CacheStorage) RequestCacheNames(
	socket *sock.Socket,
	params *cacheStorage.RequestCacheNamesParams,
) (cacheStorage.RequestCacheNamesResult, error) {
	command := &sock.Command{
		Method: "CacheStorage.requestCacheNames",
		Params: params,
	}
	result := cacheStorage.RequestCacheNamesResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
RequestCachedResponse fetches cache entry.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
*/
func (CacheStorage) RequestCachedResponse(
	socket *sock.Socket,
	params *cacheStorage.RequestCachedResponseParams,
) (cacheStorage.RequestCachedResponseResult, error) {
	command := &sock.Command{
		Method: "CacheStorage.requestCachedResponse",
		Params: params,
	}
	result := cacheStorage.RequestCachedResponseResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
RequestEntries requests data from cache.

https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
*/
func (CacheStorage) RequestEntries(
	socket *sock.Socket,
	params *cacheStorage.RequestEntriesParams,
) (cacheStorage.RequestEntriesResult, error) {
	command := &sock.Command{
		Method: "CacheStorage.requestEntries",
		Params: params,
	}
	result := cacheStorage.RequestEntriesResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}
