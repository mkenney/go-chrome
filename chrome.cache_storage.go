package chrome

import "app/chrome/protocol"

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
) (nil, error) {
	command := &protocol.Command{
		method: "CacheStorage.deleteCache",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
DeleteEntry deletes a cache entry.
*/
func (CacheStorage) DeleteEntry(
	socket *Socket,
	params *cache_storage.DeleteEntryParams,
) (nil, error) {
	command := &protocol.Command{
		method: "CacheStorage.deleteEntry",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RequestCacheNames requests cache names.
*/
func (CacheStorage) RequestCacheNames(
	socket *Socket,
	params *cache_storage.RequestCacheNamesParams,
) (cache_storage.RequestCacheNamesResult, error) {
	command := &protocol.Command{
		method: "CacheStorage.requestCacheNames",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(cache_storage.RequestCacheNamesResult), command.Err
}

/*
RequestCachedResponse fetches cache entry.
*/
func (CacheStorage) RequestCachedResponse(
	socket *Socket,
	params *cache_storage.RequestCachedResponseParams,
) (cache_storage.RequestCachedResponseResult, error) {
	command := &protocol.Command{
		method: "CacheStorage.requestCachedResponse",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(cache_storage.RequestCachedResponseResult), command.Err
}

/*
RequestEntries requests data from cache.
*/
func (CacheStorage) RequestEntries(
	socket *Socket,
	params *cache_storage.RequestEntriesParams,
) (cache_storage.RequestEntriesResult, error) {
	command := &protocol.Command{
		method: "CacheStorage.requestEntries",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(cache_storage.RequestEntriesResult), command.Err
}
