package socket

import (
	"encoding/json"

	chromeLog "github.com/mkenney/go-chrome/protocol/log"
	log "github.com/sirupsen/logrus"
)

/*
LogProtocol provides a namespace for the Chrome Log protocol methods. The Log protocol provides
access to log entries.

https://chromedevtools.github.io/devtools-protocol/tot/Log/
*/
type LogProtocol struct {
	Socket Socketer
}

/*
Clear clears the log.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
*/
func (protocol *LogProtocol) Clear() error {
	command := NewCommand("Log.clear", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables log domain, prevents further log entries from being reported to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
*/
func (protocol *LogProtocol) Disable() error {
	command := NewCommand("Log.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables log domain, sends the entries collected so far to the client by means of the
`entryAdded` notification.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
*/
func (protocol *LogProtocol) Enable() error {
	command := NewCommand("Log.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StartViolationsReport starts violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
func (protocol *LogProtocol) StartViolationsReport(
	params *chromeLog.StartViolationsReportParams,
) error {
	command := NewCommand("Log.startViolationsReport", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopViolationsReport stops violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
*/
func (protocol *LogProtocol) StopViolationsReport() error {
	command := NewCommand("Log.stopViolationsReport", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnEntryAdded adds a handler to the Log.entryAdded event. Log.entryAdded fires when a new message is
logged.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#event-entryAdded
*/
func (protocol *LogProtocol) OnEntryAdded(
	callback func(event *chromeLog.EntryAddedEvent),
) {
	handler := NewEventHandler(
		"Log.entryAdded",
		func(response *Response) {
			event := &chromeLog.EntryAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
