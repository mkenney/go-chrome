package socket

import (
	"encoding/json"

	indexedDB "github.com/mkenney/go-chrome/protocol/indexed_db"
)

/*
IndexedDBProtocol provides a namespace for the Chrome IndexedDB protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/ EXPERIMENTAL.
*/
type IndexedDBProtocol struct {
	Socket Socketer
}

/*
ClearObjectStore clears all entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
*/
func (protocol *IndexedDBProtocol) ClearObjectStore(
	params *indexedDB.ClearObjectStoreParams,
) error {
	command := NewCommand("IndexedDB.clearObjectStore", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DeleteDatabase deletes a database.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
*/
func (protocol *IndexedDBProtocol) DeleteDatabase(
	params *indexedDB.DeleteDatabaseParams,
) error {
	command := NewCommand("IndexedDB.deleteDatabase", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DeleteObjectStoreEntries deletes a range of entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
func (protocol *IndexedDBProtocol) DeleteObjectStoreEntries(
	params *indexedDB.DeleteObjectStoreEntriesParams,
) error {
	command := NewCommand("IndexedDB.deleteObjectStoreEntries", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
*/
func (protocol *IndexedDBProtocol) Disable() error {
	command := NewCommand("IndexedDB.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
*/
func (protocol *IndexedDBProtocol) Enable() error {
	command := NewCommand("IndexedDB.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RequestData requests data from object store or index.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
*/
func (protocol *IndexedDBProtocol) RequestData(
	params *indexedDB.RequestDataParams,
) (*indexedDB.RequestDataResult, error) {
	command := NewCommand("IndexedDB.requestData", params)
	result := &indexedDB.RequestDataResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
RequestDatabase requests database with given name in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
*/
func (protocol *IndexedDBProtocol) RequestDatabase(
	params *indexedDB.RequestDatabaseParams,
) (*indexedDB.RequestDatabaseResult, error) {
	command := NewCommand("IndexedDB.requestDatabase", params)
	result := &indexedDB.RequestDatabaseResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
RequestDatabaseNames requests database names for given security origin.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
*/
func (protocol *IndexedDBProtocol) RequestDatabaseNames(
	params *indexedDB.RequestDatabaseNamesParams,
) (*indexedDB.RequestDatabaseNamesResult, error) {
	command := NewCommand("IndexedDB.requestDatabaseNames", params)
	result := &indexedDB.RequestDatabaseNamesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}
