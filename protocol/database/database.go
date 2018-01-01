/*
Package database provides type definitions for use with the Chrome Database protocol

https://chromedevtools.github.io/devtools-protocol/tot/Database/
*/
package database

/*
AddEvent represents Database.addDatabase event data.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#event-addDatabase
*/
type AddEvent struct {
	Database *Database `json:"database"`
}

/*
ID is a unique identifier of a database object.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-DatabaseId
*/
type ID string

/*
Database is a database object

https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-Database
*/
type Database struct {
	// Database ID.
	ID ID `json:"id"`

	// Database domain.
	Domain string `json:"domain"`

	// Database name.
	Name string `json:"name"`

	// Database version.
	Version string `json:"version"`
}

/*
Error is a database error.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-Error
*/
type Error struct {
	// Error message.
	Message string `json:"message"`

	// Error code.
	Code int `json:"code"`
}
