package chrome

import (
	cacheStorage "app/chrome/cache_storage"
	"app/chrome/protocol"
)

/*
CacheStorage: https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/
EXPERIMENTAL
*/
type CacheStorage struct{}

/*
RequestCacheNames requests cache names.
*/
func (CacheStorage) RequestCacheNames(socket *Socket, params *cacheStorage.RequestCacheNamesParams) error {
	command := new(protocol.Command)
	command.method = "CacheStorage.requestCacheNames"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RequestEntries requests data from cache.
*/
func (CacheStorage) RequestEntries(socket *Socket, params *cacheStorage.RequestEntriesParams) error {
	command := new(protocol.Command)
	command.method = "CacheStorage.requestEntries"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteCache deletes a cache.
*/
func (CacheStorage) DeleteCache(socket *Socket, params *cacheStorage.DeleteCacheParams) error {
	command := new(protocol.Command)
	command.method = "CacheStorage.deleteCache"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteEntry deletes a cache entry.
*/
func (CacheStorage) DeleteEntry(socket *Socket, params *cacheStorage.DeleteEntryParams) error {
	command := new(protocol.Command)
	command.method = "CacheStorage.deleteEntry"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RequestCachedResponse fetches cache entry.
*/
func (CacheStorage) RequestCachedResponse(socket *Socket, params *cacheStorage.RequestCachedResponseParams) error {
	command := new(protocol.Command)
	command.method = "CacheStorage.requestCachedResponse"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
