package protocol

import (
	"encoding/json"

	indexedDB "github.com/mkenney/go-chrome/protocol/indexed_db"
	sock "github.com/mkenney/go-chrome/socket"
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
	socket *sock.Socket,
	params *indexedDB.ClearObjectStoreParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *indexedDB.DeleteDatabaseParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *indexedDB.DeleteObjectStoreEntriesParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *indexedDB.RequestDataParams,
) (indexedDB.RequestDataResult, error) {
	command := &sock.Command{
		Method: "IndexedDB.requestData",
		Params: params,
	}
	result := indexedDB.RequestDataResult{}
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
	socket *sock.Socket,
	params *indexedDB.RequestDatabaseParams,
) (indexedDB.RequestDatabaseResult, error) {
	command := &sock.Command{
		Method: "IndexedDB.requestDatabase",
		Params: params,
	}
	result := indexedDB.RequestDatabaseResult{}
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
	socket *sock.Socket,
	params *indexedDB.RequestDatabaseNamesParams,
) (indexedDB.RequestDatabaseNamesResult, error) {
	command := &sock.Command{
		Method: "IndexedDB.requestDatabaseNames",
		Params: params,
	}
	result := indexedDB.RequestDatabaseNamesResult{}
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
