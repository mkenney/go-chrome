package Database

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

/*
GetDatabaseTableNamesParams represents Database.getDatabaseTableNames parameters.
*/
type GetDatabaseTableNamesParams struct {
	DatabaseID DatabaseID `json:"databaseId"`
}

/*
ExecuteSQLParams represents Database.executeSQL parameters.
*/
type ExecuteSQLParams struct {
	DatabaseID DatabaseID `json:"databaseId"`
	Query      string     `json:"query"`
}

/*
AddDatabaseEvent represents Database.addDatabase event data.
*/
type AddDatabaseEvent struct {
	Database Database `json:"database"`
}
