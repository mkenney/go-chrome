package IndexedDB

import (
	Runtime "app/chrome/runtime"
)

/*
ClearObjectStoreParams represents IndexedDB.clearObjectStore parameters.
*/
type ClearObjectStoreParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`

	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

/*
DeleteDatabaseParams represents IndexedDB.deleteDatabase parameters.
*/
type DeleteDatabaseParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`

	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

/*
RequestDataParams represents IndexedDB.requestData parameters.
*/
type RequestDataParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`

	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`

	// Index name, empty string for object store data requests.
	IndexName string `json:"indexName"`

	// Number of records to skip.
	SkipCount int `json:"skipCount"`

	// Number of records to fetch.
	PageSize int `json:"pageSize"`

	// Optional. Key range.
	//
	// This is an instance of KeyRange, but that doesn't omitempty correctly so it must be added
	// manually.
	KeyRange interface{} `json:"keyRange,omitempty"`
}

/*
RequestDataResult represents the result of calls to IndexedDB.requestData.
*/
type RequestDataResult struct {
	// Array of object store data entries.
	ObjectStoreDataEntries []DataEntry `json:"objectStoreDataEntries"`

	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

/*
RequestDatabaseParams represents IndexedDB.requestDatabase parameters.
*/
type RequestDatabaseParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`
}

/*
RequestDatabaseResult represents the result of calls to IndexedDB.requestDatabase.
*/
type RequestDatabaseResult struct {
	// Database with an array of object stores.
	DatabaseWithObjectStores DatabaseWithObjectStores `json:"databaseWithObjectStores"`
}

/*
RequestDatabaseNamesParams represents IndexedDB.requestDatabaseNames parameters.
*/
type RequestDatabaseNamesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

/*
RequestDatabaseNamesResult represents the result of calls to IndexedDB.requestDatabaseNames.
*/
type RequestDatabaseNamesResult struct {
	// Database names for origin.
	DatabaseNames []string `json:"databaseNames"`
}

/*
DatabaseWithObjectStores is a database with an array of object stores.
*/
type DatabaseWithObjectStores struct {
	// Database name.
	Name string `json:"name"`

	// Database version.
	Version int `json:"version"`

	// Object stores in this database.
	ObjectStores []*ObjectStore `json:"objectStores"`
}

/*
ObjectStore is a object store
*/
type ObjectStore struct {
	// Object store name.
	Name string `json:"name"`

	// Object store key path.
	KeyPath KeyPath `json:"keyPath"`

	// If true, object store has auto increment flag set.
	AutoIncrement bool `json:"autoIncrement"`

	// Indexes in this object store.
	Indexes []*ObjectStoreIndex `json:"indexes"`
}

/*
ObjectStoreIndex is an object store index.
*/
type ObjectStoreIndex struct {
	// Index name.
	Name string `json:"name"`

	// Index key path.
	KeyPath *KeyPath `json:"keyPath"`

	// If true, index is unique.
	Unique bool `json:"unique"`

	// If true, index allows multiple entries for a key.
	MultiEntry bool `json:"multiEntry"`
}

/*
Key is a key.
*/
type Key struct {
	// Key type. Allowed values: number, string, date, array.
	Type string `json:"type"`

	// Optional. Number value.
	Number float64 `json:"number,omitempty"`

	// Optional. String value.
	String string `json:"string,omitempty"`

	// Optional. Date value.
	Date float64 `json:"date,omitempty"`

	// Optional. Array value.
	Array []*Key `json:"array,omitempty"`
}

/*
KeyRange is a key range.
*/
type KeyRange struct {
	// Optional. Lower bound.
	//
	// This is an instance of Key, but that doesn't omitempty correctly so it must be added
	// manually.
	Lower interface{} `json:"lower,omitempty"`

	// Optional. Upper bound.
	//
	// This is an instance of Key, but that doesn't omitempty correctly so it must be added
	// manually.
	Upper interface{} `json:"upper,omitempty"`

	// If true lower bound is open.
	LowerOpen bool `json:"lowerOpen"`

	// If true upper bound is open.
	UpperOpen bool `json:"upperOpen"`
}

/*
DataEntry is a data entry.
*/
type DataEntry struct {
	// Key object.
	Key *Runtime.RemoteObject `json:"key"`

	// Primary key object.
	PrimaryKey *Runtime.RemoteObject `json:"primaryKey"`

	// Value object.
	Value *Runtime.RemoteObject `json:"value"`
}

/*
KeyPath is a key path.
*/
type KeyPath struct {
	// Key path type. Allowed values: null, string, array.
	Type string `json:"type"`

	// Optional. String value.
	String string `json:"string,omitempty"`

	// Optional. Array value.
	Array []string `json:"array,omitempty"`
}
