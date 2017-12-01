package chrome

import (
	indexedDB "app/chrome/indexed_db"
	"app/chrome/protocol"
)

/*
IndexedDB EXPERIMENTAL
*/
type IndexedDB struct{}

/*
ClearObjectStore clears all entries from an object store.
*/
func (IndexedDB) ClearObjectStore(socket *Socket, params *indexedDB.ClearObjectStoreParams) error {
	command := new(protocol.Command)
	command.method = "IndexedDB.clearObjectStore"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteDatabase deletes a database.
*/
func (IndexedDB) DeleteDatabase(socket *Socket, params *indexedDB.DeleteDatabaseParams) error {
	command := new(protocol.Command)
	command.method = "IndexedDB.deleteDatabase"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables events from backend.
*/
func (IndexedDB) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "IndexedDB.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables events from backend.
*/
func (IndexedDB) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "IndexedDB.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
RequestData requests data from object store or index.
*/
func (IndexedDB) RequestData(socket *Socket, params *indexedDB.RequestDataParams) error {
	command := new(protocol.Command)
	command.method = "IndexedDB.requestData"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RequestDatabase requests database with given name in given frame.
*/
func (IndexedDB) RequestDatabase(socket *Socket, params *indexedDB.RequestDatabaseParams) error {
	command := new(protocol.Command)
	command.method = "IndexedDB.requestDatabase"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
RequestDatabaseNames requests database names for given security origin.
*/
func (IndexedDB) RequestDatabaseNames(socket *Socket, params *indexedDB.RequestDatabaseNamesParams) error {
	command := new(protocol.Command)
	command.method = "IndexedDB.requestDatabaseNames"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}
