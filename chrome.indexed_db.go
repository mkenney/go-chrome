package chrome

import (
	indexed_db "app/chrome/indexed_db"
	"app/chrome/protocol"
	"encoding/json"
)

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
) error {
	command := &protocol.Command{
		Method: "IndexedDB.clearObjectStore",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DeleteDatabase deletes a database.
*/
func (IndexedDB) DeleteDatabase(
	socket *Socket,
	params *indexed_db.DeleteDatabaseParams,
) error {
	command := &protocol.Command{
		Method: "IndexedDB.deleteDatabase",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables events from backend.
*/
func (IndexedDB) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "IndexedDB.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables events from backend.
*/
func (IndexedDB) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "IndexedDB.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestData requests data from object store or index.
*/
func (IndexedDB) RequestData(
	socket *Socket,
	params *indexed_db.RequestDataParams,
) (indexed_db.RequestDataResult, error) {
	command := &protocol.Command{
		Method: "IndexedDB.requestData",
		Params: params,
	}
	result := indexed_db.RequestDataResult{}
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
RequestDatabase requests database with given name in given frame.
*/
func (IndexedDB) RequestDatabase(
	socket *Socket,
	params *indexed_db.RequestDatabaseParams,
) (indexed_db.RequestDatabaseResult, error) {
	command := &protocol.Command{
		Method: "IndexedDB.requestDatabase",
		Params: params,
	}
	result := indexed_db.RequestDatabaseResult{}
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
RequestDatabaseNames requests database names for given security origin.
*/
func (IndexedDB) RequestDatabaseNames(
	socket *Socket,
	params *indexed_db.RequestDatabaseNamesParams,
) (indexed_db.RequestDatabaseNamesResult, error) {
	command := &protocol.Command{
		Method: "IndexedDB.requestDatabaseNames",
		Params: params,
	}
	result := indexed_db.RequestDatabaseNamesResult{}
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
