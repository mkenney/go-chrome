package chrome

import (
	"encoding/json"

	indexed_db "github.com/mkenney/go-chrome/indexed_db"
	"github.com/mkenney/go-chrome/protocol"
)

/*
IndexedDB - https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/
EXPERIMENTAL
*/
type IndexedDB struct{}

/*
ClearObjectStore clears all entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
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

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
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
DeleteObjectStoreEntries deletes a range of entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
func (IndexedDB) DeleteObjectStoreEntries(
	socket *Socket,
	params *indexed_db.DeleteObjectStoreEntriesParams,
) error {
	command := &protocol.Command{
		Method: "IndexedDB.deleteObjectStoreEntries",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
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

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
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

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
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

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
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

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
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
