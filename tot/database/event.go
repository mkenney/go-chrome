package database

/*
AddEvent represents Database.addDatabase event data.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#event-addDatabase
*/
type AddEvent struct {
	// Database object.
	Database *Database

	// Error information related to this event
	Err error `json:"-"`
}
