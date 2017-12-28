package protocol

import (
	"encoding/json"

	console "github.com/mkenney/go-chrome/protocol/console"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Console is a struct that provides a namespace for the Chrome Console protocol methods. DEPRECATED -
use Runtime or Log instead.

https://chromedevtools.github.io/devtools-protocol/tot/Console/
*/
var Console = _console{}

type _console struct{}

/*
ClearMessages does nothing.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-clearMessages
*/
func (_console) ClearMessages(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Console.clearMessages", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables console domain, prevents further console messages from being reported to the
client.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-disable
*/
func (_console) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Console.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables console domain, sends the messages collected so far to the client by means of the
messageAdded notification.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-enable
*/
func (_console) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Console.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnMessageAdded adds a handler to the Console.messageAdded event. Console.messageAdded fires
whenever an active document stylesheet is removed.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#event-messageAdded
*/
func (_console) OnMessageAdded(
	socket sock.Socketer,
	callback func(event *console.MessageAddedEvent),
) {
	handler := sock.NewEventHandler(
		"Console.messageAdded",
		func(response *sock.Response) {
			event := &console.MessageAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
