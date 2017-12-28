package protocol

import (
	"encoding/json"

	chromeLog "github.com/mkenney/go-chrome/protocol/log"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Log is a struct that provides a namespace for the Chrome Log protocol methods. The Log protocol
provides access to log entries.

https://chromedevtools.github.io/devtools-protocol/tot/Log/
*/
var Log = _log{}

type _log struct{}

/*
Clear clears the log.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
*/
func (_log) Clear(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Log.clear", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables log domain, prevents further log entries from being reported to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
*/
func (_log) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Log.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables log domain, sends the entries collected so far to the client by means of the
`entryAdded` notification.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
*/
func (_log) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Log.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StartViolationsReport starts violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
func (_log) StartViolationsReport(
	socket sock.Socketer,
	params *chromeLog.StartViolationsReportParams,
) error {
	command := sock.NewCommand("Log.startViolationsReport", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopViolationsReport stops violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
*/
func (_log) StopViolationsReport(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Log.stopViolationsReport", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnEntryAdded adds a handler to the Log.entryAdded event. Log.entryAdded fires when a new message is
logged.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#event-entryAdded
*/
func (_log) OnEntryAdded(
	socket sock.Socketer,
	callback func(event *chromeLog.EntryAddedEvent),
) {
	handler := sock.NewEventHandler(
		"Log.entryAdded",
		func(response *sock.Response) {
			event := &chromeLog.EntryAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
