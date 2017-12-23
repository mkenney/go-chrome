package chrome

import (
	"encoding/json"

	cacheStorage "github.com/mkenney/go-chrome/cache_storage"
	"github.com/mkenney/go-chrome/protocol"
)

/*
CacheStorage - https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/
EXPERIMENTAL
*/
type CacheStorage struct{}

/*
DeleteCache deletes a cache.
*/
func (CacheStorage) DeleteCache(
	socket *Socket,
	params *cacheStorage.DeleteCacheParams,
) error {
	command := &protocol.Command{
		Method: "CacheStorage.deleteCache",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteEntry deletes a cache entry.
*/
func (CacheStorage) DeleteEntry(
	socket *Socket,
	params *cacheStorage.DeleteEntryParams,
) error {
	command := &protocol.Command{
		Method: "CacheStorage.deleteEntry",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestCacheNames requests cache names.
*/
func (CacheStorage) RequestCacheNames(
	socket *Socket,
	params *cacheStorage.RequestCacheNamesParams,
) (cacheStorage.RequestCacheNamesResult, error) {
	command := &protocol.Command{
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
*/
func (CacheStorage) RequestCachedResponse(
	socket *Socket,
	params *cacheStorage.RequestCachedResponseParams,
) (cacheStorage.RequestCachedResponseResult, error) {
	command := &protocol.Command{
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
*/
func (CacheStorage) RequestEntries(
	socket *Socket,
	params *cacheStorage.RequestEntriesParams,
) (cacheStorage.RequestEntriesResult, error) {
	command := &protocol.Command{
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
