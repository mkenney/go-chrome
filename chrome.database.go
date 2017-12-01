package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Database - https://chromedevtools.github.io/devtools-protocol/tot/Database/
EXPERIMENTAL
*/
type Database struct{}

/*
Enable enables database tracking, database events will now be delivered to the client.
*/
func (Database) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Database.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables database tracking, prevents database events from being sent to the client.
*/
func (Database) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Database.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetDatabaseTableNames gets database table names.
*/
func (Database) GetDatabaseTableNames(socket *Socket, params *database.GetDatabaseTableNamesParams) error {
	command := &protocol.Command{
		method: "Database.getDatabaseTableNames",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ExecuteSQL executes a SQL query.
*/
func (Database) ExecuteSQL(socket *Socket, params *database.ExecuteSQLParams) error {
	command := &protocol.Command{
		method: "Database.executeSQL",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAddDatabase adds a handler to the Database.addDatabase event. Database.addDatabase fires
whenever a database is added
*/
func (Database) OnAddDatabase(socket *Socket, callback func(event *database.AddDatabaseEvent)) error {
	handler := protocol.NewEventHandler(
		"Database.addDatabase",
		func(name string, params []byte) {
			event := &database.AddDatabaseEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}
