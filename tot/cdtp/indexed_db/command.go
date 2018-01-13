package indexedDB

/*
ClearObjectStoreParams represents IndexedDB.clearObjectStore parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
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
ClearObjectStoreResult represents the result of calls to IndexedDB.clearObjectStore.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
*/
type ClearObjectStoreResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DeleteDatabaseParams represents IndexedDB.deleteDatabase parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
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
DeleteDatabaseResult represents the result of calls to IndexedDB.deleteDatabase.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
*/
type DeleteDatabaseResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DeleteObjectStoreEntriesParams represents IndexedDB.deleteObjectStoreEntries
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
type DeleteObjectStoreEntriesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`
}

/*
DeleteObjectStoreEntriesResult represents the result of calls to
IndexedDB.deleteObjectStoreEntries.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
type DeleteObjectStoreEntriesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to IndexedDB.disable.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to IndexedDB.enable.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RequestDataParams represents IndexedDB.requestData parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
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
	KeyRange *KeyRange `json:"keyRange,omitempty"`
}

/*
RequestDataResult represents the result of calls to IndexedDB.requestData.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
*/
type RequestDataResult struct {
	// Array of object store data entries.
	ObjectStoreDataEntries []*DataEntry `json:"objectStoreDataEntries"`

	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RequestDatabaseParams represents IndexedDB.requestDatabase parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
*/
type RequestDatabaseParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`
}

/*
RequestDatabaseResult represents the result of calls to IndexedDB.requestDatabase.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
*/
type RequestDatabaseResult struct {
	// Database with an array of object stores.
	DatabaseWithObjectStores *DatabaseWithObjectStores `json:"databaseWithObjectStores"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RequestDatabaseNamesParams represents IndexedDB.requestDatabaseNames parameters.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
*/
type RequestDatabaseNamesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

/*
RequestDatabaseNamesResult represents the result of calls to IndexedDB.requestDatabaseNames.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
*/
type RequestDatabaseNamesResult struct {
	// Database names for origin.
	DatabaseNames []string `json:"databaseNames"`

	// Error information related to executing this method
	Err error `json:"-"`
}
