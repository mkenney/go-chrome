package chrome

import (
	cache_storage "app/chrome/cache_storage"
	"app/chrome/protocol"
	"encoding/json"
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
	params *cache_storage.DeleteCacheParams,
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
	params *cache_storage.DeleteEntryParams,
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
	params *cache_storage.RequestCacheNamesParams,
) (cache_storage.RequestCacheNamesResult, error) {
	command := &protocol.Command{
		Method: "CacheStorage.requestCacheNames",
		Params: params,
	}
	result := cache_storage.RequestCacheNamesResult{}
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
	params *cache_storage.RequestCachedResponseParams,
) (cache_storage.RequestCachedResponseResult, error) {
	command := &protocol.Command{
		Method: "CacheStorage.requestCachedResponse",
		Params: params,
	}
	result := cache_storage.RequestCachedResponseResult{}
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
	params *cache_storage.RequestEntriesParams,
) (cache_storage.RequestEntriesResult, error) {
	command := &protocol.Command{
		Method: "CacheStorage.requestEntries",
		Params: params,
	}
	result := cache_storage.RequestEntriesResult{}
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
