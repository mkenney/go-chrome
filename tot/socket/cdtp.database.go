package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/database"
)

/*
DatabaseProtocol provides a namespace for the Chrome Database protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Database/ EXPERIMENTAL.
*/
type DatabaseProtocol struct {
	Socket Socketer
}

/*
Disable disables database tracking, prevents database events from being sent to
the client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
*/
func (protocol *DatabaseProtocol) Disable() <-chan *database.DisableResult {
	resultChan := make(chan *database.DisableResult)
	command := NewCommand(protocol.Socket, "Database.disable", nil)
	result := &database.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables database tracking, database events will now be delivered to the
client.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-enable
*/
func (protocol *DatabaseProtocol) Enable() <-chan *database.EnableResult {
	resultChan := make(chan *database.EnableResult)
	command := NewCommand(protocol.Socket, "Database.enable", nil)
	result := &database.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
ExecuteSQL executes a SQL query.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
*/
func (protocol *DatabaseProtocol) ExecuteSQL(
	params *database.ExecuteSQLParams,
) <-chan *database.ExecuteSQLResult {
	resultChan := make(chan *database.ExecuteSQLResult)
	command := NewCommand(protocol.Socket, "Database.executeSQL", params)
	result := &database.ExecuteSQLResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetTableNames gets database table names.

https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
*/
func (protocol *DatabaseProtocol) GetTableNames(
	params *database.GetTableNamesParams,
) <-chan *database.GetTableNamesResult {
	resultChan := make(chan *database.GetTableNamesResult)
	command := NewCommand(protocol.Socket, "Database.executeSQL", params)
	result := &database.GetTableNamesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnAdd adds a handler to the Database.addDatabase event. Database.addDatabase
fires whenever a database is added

https://chromedevtools.github.io/devtools-protocol/tot/Database/#event-addDatabase
*/
func (protocol *DatabaseProtocol) OnAdd(
	callback func(event *database.AddEvent),
) {
	handler := NewEventHandler(
		"Database.addDatabase",
		func(response *Response) {
			event := &database.AddEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
