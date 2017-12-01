package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Console - https://chromedevtools.github.io/devtools-protocol/tot/Console/
DEPRECATED - use Runtime or Log instead.
*/
type Console struct{}

/*
Enable enables console domain, sends the messages collected so far to the client by means of the
messageAdded notification.
*/
func (Console) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Console.enable",
		params: nil,
	}
	command.method = "Console.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables console domain, prevents further console messages from being reported to the
client.
*/
func (Console) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Console.disable",
		params: nil,
	}
	command.method = "Console.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
ClearMessages does nothing.
*/
func (Console) ClearMessages(socket *Socket) error {
	command := &protocol.Command{
		method: "Console.clearMessages",
		params: nil,
	}
	command.method = "Console.clearMessages"
	socket.SendCommand(command)
	return command.Err
}

/*
OnMessageAdded adds a handler to the Console.messageAdded event. Console.messageAdded fires
whenever an active document stylesheet is removed.
*/
func (Console) OnMessageAdded(socket *Socket, callback func(event *console.MessageAddedEvent)) error {
	handler := protocol.NewEventHandler(
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
	return command.Err
}
