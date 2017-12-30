package socket

import (
	"encoding/json"

	database "github.com/mkenney/go-chrome/protocol/database"
	log "github.com/sirupsen/logrus"
)

/*
DatabaseProtocol provides a namespace for the Chrome Database protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Database/ EXPERIMENTAL.
*/
type DatabaseProtocol struct {
	Socket Socketer
}

/*
Disable disables database tracking, prevents database events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
*/
func (protocol *DatabaseProtocol) Disable() error {
	command := NewCommand("Database.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables database tracking, database events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-enable
*/
func (protocol *DatabaseProtocol) Enable() error {
	command := NewCommand("Database.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ExecuteSQL executes a SQL query.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
*/
func (protocol *DatabaseProtocol) ExecuteSQL(
	params *database.ExecuteSQLParams,
) (*database.ExecuteSQLResult, error) {
	command := NewCommand("Database.executeSQL", params)
	result := &database.ExecuteSQLResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetTableNames gets database table names.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
*/
func (protocol *DatabaseProtocol) GetTableNames(
	params *database.GetTableNamesParams,
) error {
	command := NewCommand("Database.getDatabaseTableNames", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnAdd adds a handler to the Database.addDatabase event. Database.addDatabase fires
whenever a database is added

https://chromedevtools.github.io/devtools-protocol/tot/Database/#event-addDatabase
*/
func (protocol *DatabaseProtocol) OnAdd(
	callback func(event *database.AddEvent),
) {
	handler := NewEventHandler(
		"Database.addDatabase",
		func(response *Response) {
			event := &database.AddEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
