package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/indexed/db"
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
	params *db.ClearObjectStoreParams,
) <-chan *db.ClearObjectStoreResult {
	resultChan := make(chan *db.ClearObjectStoreResult)
	command := NewCommand(protocol.Socket, "IndexedDB.clearObjectStore", params)
	result := &db.ClearObjectStoreResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
DeleteDatabase deletes a database.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
*/
func (protocol *IndexedDBProtocol) DeleteDatabase(
	params *db.DeleteDatabaseParams,
) <-chan *db.DeleteDatabaseResult {
	resultChan := make(chan *db.DeleteDatabaseResult)
	command := NewCommand(protocol.Socket, "IndexedDB.deleteDatabase", params)
	result := &db.DeleteDatabaseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
DeleteObjectStoreEntries deletes a range of entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
func (protocol *IndexedDBProtocol) DeleteObjectStoreEntries(
	params *db.DeleteObjectStoreEntriesParams,
) <-chan *db.DeleteObjectStoreEntriesResult {
	resultChan := make(chan *db.DeleteObjectStoreEntriesResult)
	command := NewCommand(protocol.Socket, "IndexedDB.deleteObjectStoreEntries", params)
	result := &db.DeleteObjectStoreEntriesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
*/
func (protocol *IndexedDBProtocol) Disable() <-chan *db.DisableResult {
	resultChan := make(chan *db.DisableResult)
	command := NewCommand(protocol.Socket, "IndexedDB.disable", nil)
	result := &db.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
*/
func (protocol *IndexedDBProtocol) Enable() <-chan *db.EnableResult {
	resultChan := make(chan *db.EnableResult)
	command := NewCommand(protocol.Socket, "IndexedDB.enable", nil)
	result := &db.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RequestData requests data from object store or index.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
*/
func (protocol *IndexedDBProtocol) RequestData(
	params *db.RequestDataParams,
) <-chan *db.RequestDataResult {
	resultChan := make(chan *db.RequestDataResult)
	command := NewCommand(protocol.Socket, "IndexedDB.requestData", params)
	result := &db.RequestDataResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RequestDatabase requests database with given name in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
*/
func (protocol *IndexedDBProtocol) RequestDatabase(
	params *db.RequestDatabaseParams,
) <-chan *db.RequestDatabaseResult {
	resultChan := make(chan *db.RequestDatabaseResult)
	command := NewCommand(protocol.Socket, "IndexedDB.requestDatabase", params)
	result := &db.RequestDatabaseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RequestDatabaseNames requests database names for given security origin.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
*/
func (protocol *IndexedDBProtocol) RequestDatabaseNames(
	params *db.RequestDatabaseNamesParams,
) <-chan *db.RequestDatabaseNamesResult {
	resultChan := make(chan *db.RequestDatabaseNamesResult)
	command := NewCommand(protocol.Socket, "IndexedDB.requestDatabaseNames", params)
	result := &db.RequestDatabaseNamesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}
