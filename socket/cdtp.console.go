package socket

import (
	"encoding/json"

	console "github.com/mkenney/go-chrome/cdtp/console"
	log "github.com/sirupsen/logrus"
)

/*
ConsoleProtocol provides a namespace for the Chrome Console protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Console/ DEPRECATED - use
Runtime or Log instead.
*/
type ConsoleProtocol struct {
	Socket Socketer
}

/*
ClearMessages does nothing.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-clearMessages
*/
func (protocol *ConsoleProtocol) ClearMessages() error {
	command := NewCommand("Console.clearMessages", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables console domain, prevents further console messages from being
reported to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-disable
*/
func (protocol *ConsoleProtocol) Disable() error {
	command := NewCommand("Console.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables console domain, sends the messages collected so far to the client
by means of the messageAdded notification.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-enable
*/
func (protocol *ConsoleProtocol) Enable() error {
	command := NewCommand("Console.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnMessageAdded adds a handler to the Console.messageAdded event.
Console.messageAdded fires whenever an active document stylesheet is removed.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#event-messageAdded
*/
func (protocol *ConsoleProtocol) OnMessageAdded(
	callback func(event *console.MessageAddedEvent),
) {
	handler := NewEventHandler(
		"Console.messageAdded",
		func(response *Response) {
			event := &console.MessageAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
