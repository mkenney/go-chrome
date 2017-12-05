package chrome

import (
	chrome_log "app/chrome/log"
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Log - https://chromedevtools.github.io/devtools-protocol/tot/Log/
Provides access to log entries.
*/
type Log struct{}

/*
Clear clears the log.
*/
func (Log) Clear(socket *Socket) error {
	command := &protocol.Command{
		method: "Log.clear",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables log domain, prevents further log entries from being reported to the client.
*/
func (Log) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Log.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables log domain, sends the entries collected so far to the client by means of the
`entryAdded` notification.
*/
func (Log) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Log.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartViolationsReport starts violation reporting.
*/
func (Log) StartViolationsReport(socket *Socket, params *chrome_log.StartViolationsReportParams) error {
	command := &protocol.Command{
		method: "Log.startViolationsReport",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopViolationsReport stops violation reporting.
*/
func (Log) StopViolationsReport(socket *Socket) error {
	command := &protocol.Command{
		method: "Log.stopViolationsReport",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnEntryAdded adds a handler to the Log.entryAdded event. Log.entryAdded fires when a new message is
logged.
*/
func (Log) OnEntryAdded(socket *Socket, callback func(event *chrome_log.EntryAddedEvent)) {
	handler := protocol.NewEventHandler(
		"Log.entryAdded",
		func(name string, params []byte) {
			event := &chrome_log.EntryAddedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
