package protocol

import (
	"encoding/json"

	database "github.com/mkenney/go-chrome/protocol/database"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Database is a struct that provides a namespace for the Chrome Database protocol methods.

EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Database/
*/
type Database struct{}

/*
Disable disables database tracking, prevents database events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
*/
func (Database) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Database.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables database tracking, database events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-enable
*/
func (Database) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Database.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ExecuteSQL executes a SQL query.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
*/
func (Database) ExecuteSQL(
	socket sock.Socketer,
	params *database.ExecuteSQLParams,
) (database.ExecuteSQLResult, error) {
	command := sock.NewCommand("Database.executeSQL", params)
	result := database.ExecuteSQLResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetTableNames gets database table names.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
*/
func (Database) GetTableNames(
	socket sock.Socketer,
	params *database.GetTableNamesParams,
) error {
	command := sock.NewCommand("Database.getDatabaseTableNames", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnAdd adds a handler to the Database.addDatabase event. Database.addDatabase fires
whenever a database is added

https://chromedevtools.github.io/devtools-protocol/tot/Database/#event-addDatabase
*/
func (Database) OnAdd(
	socket sock.Socketer,
	callback func(event *database.AddEvent),
) {
	handler := sock.NewEventHandler(
		"Database.addDatabase",
		func(response *sock.Response) {
			event := &database.AddEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
