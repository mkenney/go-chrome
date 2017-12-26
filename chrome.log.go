package Chrome

import (
	"encoding/json"

	chrome_log "github.com/mkenney/go-chrome/log"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/sirupsen/logrus"
)

/*
Log - https://chromedevtools.github.io/devtools-protocol/tot/Log/
Provides access to log entries.
*/
type Log struct{}

/*
Clear clears the log.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
*/
func (Log) Clear(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Log.clear",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables log domain, prevents further log entries from being reported to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
*/
func (Log) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Log.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables log domain, sends the entries collected so far to the client by means of the
`entryAdded` notification.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
*/
func (Log) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Log.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartViolationsReport starts violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
func (Log) StartViolationsReport(
	socket *Socket,
	params *chrome_log.StartViolationsReportParams,
) error {
	command := &protocol.Command{
		Method: "Log.startViolationsReport",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopViolationsReport stops violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
*/
func (Log) StopViolationsReport(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Log.stopViolationsReport",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnEntryAdded adds a handler to the Log.entryAdded event. Log.entryAdded fires when a new message is
logged.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#event-entryAdded
*/
func (Log) OnEntryAdded(
	socket *Socket,
	callback func(event *chrome_log.EntryAddedEvent),
) {
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
