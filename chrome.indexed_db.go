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
func (IndexedDB) ClearObjectStore(socket *Socket, params *indexed_db.ClearObjectStoreParams) error {
	command := &protocol.Command{
		method: "IndexedDB.clearObjectStore",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteDatabase deletes a database.
*/
func (IndexedDB) DeleteDatabase(socket *Socket, params *indexed_db.DeleteDatabaseParams) error {
	command := &protocol.Command{
		method: "IndexedDB.deleteDatabase",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables events from backend.
*/
func (IndexedDB) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "IndexedDB.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables events from backend.
*/
func (IndexedDB) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "IndexedDB.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestData requests data from object store or index.
*/
func (IndexedDB) RequestData(socket *Socket, params *indexed_db.RequestDataParams) error {
	command := &protocol.Command{
		method: "IndexedDB.requestData",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestDatabase requests database with given name in given frame.
*/
func (IndexedDB) RequestDatabase(socket *Socket, params *indexed_db.RequestDatabaseParams) error {
	command := &protocol.Command{
		method: "IndexedDB.requestDatabase",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestDatabaseNames requests database names for given security origin.
*/
func (IndexedDB) RequestDatabaseNames(socket *Socket, params *indexed_db.RequestDatabaseNamesParams) error {
	command := &protocol.Command{
		method: "IndexedDB.requestDatabaseNames",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}
