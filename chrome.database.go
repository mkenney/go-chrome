package chrome

import (
	"app/chrome/protocol"
	database "app/chrome/database"
)

/*
Database: https://chromedevtools.github.io/devtools-protocol/tot/Database/
EXPERIMENTAL
*/
type Database struct{}

/*
Enable enables database tracking, database events will now be delivered to the client.
*/
func (Database) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Database.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables database tracking, prevents database events from being sent to the client.
*/
func (Database) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Database.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
GetDatabaseTableNames gets database table names.
*/
func (Database) GetDatabaseTableNames(socket *Socket, params *database.GetDatabaseTableNamesParams) error {
	command := new(protocol.Command)
	command.method = "Database.getDatabaseTableNames"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ExecuteSQL executes a SQL query.
*/
func (Database) ExecuteSQL(socket *Socket, params *database.ExecuteSQLParams) error {
	command := new(protocol.Command)
	command.method = "Database.executeSQL"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
OnAddDatabase adds a handler to the Database.addDatabase event. Database.addDatabase fires
whenever a database is added
*/
func OnAddDatabase(socket *Socket, callback func(event *database.AddDatabaseEvent)) error {
	handler := protocol.NewEventHandler(
		"Database.addDatabase",
		func(name string, params []byte) {
			event := new(database.AddDatabaseEvent)
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		}
	)
	socket.AddEventHandler(handler)
	return command.Err
}
