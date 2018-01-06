package database

/*
DisableResult represents the result of calls to Database.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Database.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ExecuteSQLParams represents Database.executeSQL parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
*/
type ExecuteSQLParams struct {
	ID    ID     `json:"databaseId"`
	Query string `json:"query"`
}

/*
ExecuteSQLResult represents the result of calls to Database.executeSQL.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
*/
type ExecuteSQLResult struct {
	// Column names.
	ColumnNames []string `json:"columnNames"`

	// Values.
	Values []interface{} `json:"values"`

	// Optional. Error, if any.
	SQLError *Error `json:"sqlError,omitempty"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetTableNamesParams represents Database.getDatabaseTableNames parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
*/
type GetTableNamesParams struct {
	ID ID `json:"databaseId"`
}

/*
GetTableNamesResult represents the result of calls to Database.getDatabaseTableNames.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
*/
type GetTableNamesResult struct {
	// Table names.
	TableNames []string `json:"tableNames"`

	// Error information related to executing this method
	Err error `json:"-"`
}
