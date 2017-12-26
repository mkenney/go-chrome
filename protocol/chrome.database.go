package protocol

import (
	"encoding/json"

	database "github.com/mkenney/go-chrome/protocol/database"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Database - https://chromedevtools.github.io/devtools-protocol/tot/Database/
EXPERIMENTAL
*/
type Database struct{}

/*
Disable disables database tracking, prevents database events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
*/
func (Database) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Database.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables database tracking, database events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-enable
*/
func (Database) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Database.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ExecuteSQL executes a SQL query.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
*/
func (Database) ExecuteSQL(
	socket *sock.Socket,
	params *database.ExecuteSQLParams,
) (database.ExecuteSQLResult, error) {
	command := &sock.Command{
		Method: "Database.executeSQL",
		Params: params,
	}
	result := database.ExecuteSQLResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetTableNames gets database table names.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
*/
func (Database) GetTableNames(
	socket *sock.Socket,
	params *database.GetTableNamesParams,
) error {
	command := &sock.Command{
		Method: "Database.getDatabaseTableNames",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAdd adds a handler to the Database.addDatabase event. Database.addDatabase fires
whenever a database is added

https://chromedevtools.github.io/devtools-protocol/tot/Database/#event-addDatabase
*/
func (Database) OnAdd(
	socket *sock.Socket,
	callback func(event *database.AddEvent),
) {
	handler := sock.NewEventHandler(
		"Database.addDatabase",
		func(name string, params []byte) {
			event := &database.AddEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
