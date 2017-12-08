package chrome

import "app/chrome/protocol"

/*
IndexedDB - https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/
EXPERIMENTAL
*/
type IndexedDB struct{}

/*
ClearObjectStore clears all entries from an object store.
*/
func (IndexedDB) ClearObjectStore(
	socket *Socket,
	params *indexed_db.ClearObjectStoreParams,
) (indexed_db.ClearObjectStoreResult, error) {
	command := &protocol.Command{
		method: "IndexedDB.clearObjectStore",
		params: params,
	}
	socket.SendCommand(command)
	return indexed_db.ClearObjectStoreResult, command.Err
}

/*
DeleteDatabase deletes a database.
*/
func (IndexedDB) DeleteDatabase(
	socket *Socket,
	params *indexed_db.DeleteDatabaseParams,
) (nil, error) {
	command := &protocol.Command{
		method: "IndexedDB.deleteDatabase",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Disable disables events from backend.
*/
func (IndexedDB) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "IndexedDB.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables events from backend.
*/
func (IndexedDB) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "IndexedDB.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RequestData requests data from object store or index.
*/
func (IndexedDB) RequestData(
	socket *Socket,
	params *indexed_db.RequestDataParams,
) (indexed_db.RequestDataResult, error) {
	command := &protocol.Command{
		method: "IndexedDB.requestData",
		params: params,
	}
	socket.SendCommand(command)
	return indexed_db.RequestDataResult, command.Err
}

/*
RequestDatabase requests database with given name in given frame.
*/
func (IndexedDB) RequestDatabase(
	socket *Socket,
	params *indexed_db.RequestDatabaseParams,
) (indexed_db.RequestDatabaseResult, error) {
	command := &protocol.Command{
		method: "IndexedDB.requestDatabase",
		params: params,
	}
	socket.SendCommand(command)
	return indexed_db.RequestDatabaseResult, command.Err
}

/*
RequestDatabaseNames requests database names for given security origin.
*/
func (IndexedDB) RequestDatabaseNames(
	socket *Socket,
	params *indexed_db.RequestDatabaseNamesParams,
) (indexed_db.RequestDatabaseNamesResult, error) {
	command := &protocol.Command{
		method: "IndexedDB.requestDatabaseNames",
		params: params,
	}
	socket.SendCommand(command)
	return indexed_db.RequestDatabaseNamesResult, command.Err
}
