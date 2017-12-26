package protocol

import (
	"encoding/json"

	console "github.com/mkenney/go-chrome/protocol/console"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Console is a struct that provides a namespace for the Chrome Console protocol methods.

DEPRECATED - use Runtime or Log instead.

https://chromedevtools.github.io/devtools-protocol/tot/Console/
*/
type Console struct{}

/*
ClearMessages does nothing.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-clearMessages
*/
func (Console) ClearMessages(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Console.clearMessages",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables console domain, prevents further console messages from being reported to the
client.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-disable
*/
func (Console) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Console.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables console domain, sends the messages collected so far to the client by means of the
messageAdded notification.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-enable
*/
func (Console) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Console.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnMessageAdded adds a handler to the Console.messageAdded event. Console.messageAdded fires
whenever an active document stylesheet is removed.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#event-messageAdded
*/
func (Console) OnMessageAdded(
	socket *sock.Socket,
	callback func(event *console.MessageAddedEvent),
) {
	handler := sock.NewEventHandler(
		"Console.messageAdded",
		func(name string, params []byte) {
			event := &console.MessageAddedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
