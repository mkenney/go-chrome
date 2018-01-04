package socket

import (
	"encoding/json"

	indexedDB "github.com/mkenney/go-chrome/cdtp/indexed_db"
)

/*
IndexedDBProtocol provides a namespace for the Chrome IndexedDB protocol
methods.

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
) chan *indexedDB.ClearObjectStoreResult {
	resultChan := make(chan *indexedDB.ClearObjectStoreResult)
	command := NewCommand(protocol.Socket, "IndexedDB.clearObjectStore", params)
	result := &indexedDB.ClearObjectStoreResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DeleteDatabase deletes a database.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
*/
func (protocol *IndexedDBProtocol) DeleteDatabase(
	params *indexedDB.DeleteDatabaseParams,
) chan *indexedDB.DeleteDatabaseResult {
	resultChan := make(chan *indexedDB.DeleteDatabaseResult)
	command := NewCommand(protocol.Socket, "IndexedDB.deleteDatabase", params)
	result := &indexedDB.DeleteDatabaseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DeleteObjectStoreEntries deletes a range of entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
func (protocol *IndexedDBProtocol) DeleteObjectStoreEntries(
	params *indexedDB.DeleteObjectStoreEntriesParams,
) chan *indexedDB.DeleteObjectStoreEntriesResult {
	resultChan := make(chan *indexedDB.DeleteObjectStoreEntriesResult)
	command := NewCommand(protocol.Socket, "IndexedDB.deleteObjectStoreEntries", params)
	result := &indexedDB.DeleteObjectStoreEntriesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Disable disables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
*/
func (protocol *IndexedDBProtocol) Disable() chan *indexedDB.DisableResult {
	resultChan := make(chan *indexedDB.DisableResult)
	command := NewCommand(protocol.Socket, "IndexedDB.disable", nil)
	result := &indexedDB.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
*/
func (protocol *IndexedDBProtocol) Enable() chan *indexedDB.EnableResult {
	resultChan := make(chan *indexedDB.EnableResult)
	command := NewCommand(protocol.Socket, "IndexedDB.enable", nil)
	result := &indexedDB.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RequestData requests data from object store or index.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
*/
func (protocol *IndexedDBProtocol) RequestData(
	params *indexedDB.RequestDataParams,
) chan *indexedDB.RequestDataResult {
	resultChan := make(chan *indexedDB.RequestDataResult)
	command := NewCommand(protocol.Socket, "IndexedDB.requestData", params)
	result := &indexedDB.RequestDataResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RequestDatabase requests database with given name in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
*/
func (protocol *IndexedDBProtocol) RequestDatabase(
	params *indexedDB.RequestDatabaseParams,
) chan *indexedDB.RequestDatabaseResult {
	resultChan := make(chan *indexedDB.RequestDatabaseResult)
	command := NewCommand(protocol.Socket, "IndexedDB.requestDatabase", params)
	result := &indexedDB.RequestDatabaseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RequestDatabaseNames requests database names for given security origin.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
*/
func (protocol *IndexedDBProtocol) RequestDatabaseNames(
	params *indexedDB.RequestDatabaseNamesParams,
) chan *indexedDB.RequestDatabaseNamesResult {
	resultChan := make(chan *indexedDB.RequestDatabaseNamesResult)
	command := NewCommand(protocol.Socket, "IndexedDB.requestDatabaseNames", params)
	result := &indexedDB.RequestDatabaseNamesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}
