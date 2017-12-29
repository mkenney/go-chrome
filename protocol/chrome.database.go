package protocol

import (
	"encoding/json"

	database "github.com/mkenney/go-chrome/protocol/database"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Database provides a namespace for the Chrome Database protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Database/ EXPERIMENTAL.
*/
var Database = DatabaseProtocol{}

/*
DatabaseProtocol defines the Database protocol methods.
*/
type DatabaseProtocol struct{}

/*
Disable disables database tracking, prevents database events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
*/
func (DatabaseProtocol) Disable(
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
func (DatabaseProtocol) Enable(
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
func (DatabaseProtocol) ExecuteSQL(
	socket sock.Socketer,
	params *database.ExecuteSQLParams,
) (*database.ExecuteSQLResult, error) {
	command := sock.NewCommand("Database.executeSQL", params)
	result := &database.ExecuteSQLResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetTableNames gets database table names.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
*/
func (DatabaseProtocol) GetTableNames(
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
func (DatabaseProtocol) OnAdd(
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
