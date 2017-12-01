package chrome

import "app/chrome/protocol"

/*
CacheStorage - https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/
EXPERIMENTAL
*/
type CacheStorage struct{}

/*
RequestCacheNames requests cache names.
*/
func (CacheStorage) RequestCacheNames(socket *Socket, params *cache_storage.RequestCacheNamesParams) error {
	command := &protocol.Command{
		method: "CacheStorage.requestCacheNames",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestEntries requests data from cache.
*/
func (CacheStorage) RequestEntries(socket *Socket, params *cache_storage.RequestEntriesParams) error {
	command := &protocol.Command{
		method: "CacheStorage.requestEntries",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteCache deletes a cache.
*/
func (CacheStorage) DeleteCache(socket *Socket, params *cache_storage.DeleteCacheParams) error {
	command := &protocol.Command{
		method: "CacheStorage.deleteCache",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteEntry deletes a cache entry.
*/
func (CacheStorage) DeleteEntry(socket *Socket, params *cache_storage.DeleteEntryParams) error {
	command := &protocol.Command{
		method: "CacheStorage.deleteEntry",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestCachedResponse fetches cache entry.
*/
func (CacheStorage) RequestCachedResponse(socket *Socket, params *cache_storage.RequestCachedResponseParams) error {
	command := &protocol.Command{
		method: "CacheStorage.requestCachedResponse",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
