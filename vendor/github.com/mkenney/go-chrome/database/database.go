package Database

/*
ExecuteSQLParams represents Database.executeSQL parameters.
*/
type ExecuteSQLParams struct {
	DatabaseID DatabaseID `json:"databaseId"`
	Query      string     `json:"query"`
}

/*
ExecuteSQLResult represents the result of calls to Database.executeSQL.
*/
type ExecuteSQLResult struct {
	// Column names.
	ColumnNames []string `json:"columnNames"`

	// Values.
	Values []interface{} `json:"values"`

	// Optional. Error, if any.
	SQLError Error `json:"sqlError,omitempty"`
}

/*
GetDatabaseTableNamesParams represents Database.getDatabaseTableNames parameters.
*/
type GetDatabaseTableNamesParams struct {
	DatabaseID DatabaseID `json:"databaseId"`
}

/*
GetDatabaseTableNamesResult represents the result of calls to Database.getDatabaseTableNames.
*/
type GetDatabaseTableNamesResult struct {
	// Table names.
	TableNames []string `json:"tableNames"`
}

/*
AddDatabaseEvent represents Database.addDatabase event data.
*/
type AddDatabaseEvent struct {
	Database Database `json:"database"`
}

/*
DatabaseID is a unique identifier of a database object.
*/
type DatabaseID string

/*
Database is a database object
*/
type Database struct {
	// Database ID.
	ID DatabaseID `json:"id"`

	// Database domain.
	Domain string `json:"domain"`

	// Database name.
	Name string `json:"name"`

	// Database version.
	Version string `json:"version"`
}

/*
Error is a database error.
*/
type Error struct {
	// Error message.
	Message string `json:"message"`

	// Error code.
	Code int `json:"code"`
}
