package protocol

import (
	"encoding/json"

	indexedDB "github.com/mkenney/go-chrome/protocol/indexed_db"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
IndexedDB is a struct that provides a namespace for the Chrome IndexedDB protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/
*/
type IndexedDB struct{}

/*
ClearObjectStore clears all entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
*/
func (IndexedDB) ClearObjectStore(
	socket sock.Socketer,
	params *indexedDB.ClearObjectStoreParams,
) error {
	command := sock.NewCommand("IndexedDB.clearObjectStore", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DeleteDatabase deletes a database.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
*/
func (IndexedDB) DeleteDatabase(
	socket sock.Socketer,
	params *indexedDB.DeleteDatabaseParams,
) error {
	command := sock.NewCommand("IndexedDB.deleteDatabase", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DeleteObjectStoreEntries deletes a range of entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
func (IndexedDB) DeleteObjectStoreEntries(
	socket sock.Socketer,
	params *indexedDB.DeleteObjectStoreEntriesParams,
) error {
	command := sock.NewCommand("IndexedDB.deleteObjectStoreEntries", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
*/
func (IndexedDB) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("IndexedDB.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
*/
func (IndexedDB) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("IndexedDB.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
RequestData requests data from object store or index.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
*/
func (IndexedDB) RequestData(
	socket sock.Socketer,
	params *indexedDB.RequestDataParams,
) (indexedDB.RequestDataResult, error) {
	command := sock.NewCommand("IndexedDB.requestData", params)
	result := indexedDB.RequestDataResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
RequestDatabase requests database with given name in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
*/
func (IndexedDB) RequestDatabase(
	socket sock.Socketer,
	params *indexedDB.RequestDatabaseParams,
) (indexedDB.RequestDatabaseResult, error) {
	command := sock.NewCommand("IndexedDB.requestDatabase", params)
	result := indexedDB.RequestDatabaseResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
RequestDatabaseNames requests database names for given security origin.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
*/
func (IndexedDB) RequestDatabaseNames(
	socket sock.Socketer,
	params *indexedDB.RequestDatabaseNamesParams,
) (indexedDB.RequestDatabaseNamesResult, error) {
	command := sock.NewCommand("IndexedDB.requestDatabaseNames", params)
	result := indexedDB.RequestDatabaseNamesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}
