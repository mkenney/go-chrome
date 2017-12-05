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
ClearMessages does nothing.
*/
func (Console) ClearMessages(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Console.clearMessages",
	}
	command.method = "Console.clearMessages"
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Disable disables console domain, prevents further console messages from being reported to the
client.
*/
func (Console) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Console.disable",
	}
	command.method = "Console.disable"
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables console domain, sends the messages collected so far to the client by means of the
messageAdded notification.
*/
func (Console) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Console.enable",
	}
	command.method = "Console.enable"
	socket.SendCommand(command)
	return nil, command.Err
}

/*
OnMessageAdded adds a handler to the Console.messageAdded event. Console.messageAdded fires
whenever an active document stylesheet is removed.
*/
func (Console) OnMessageAdded(
	socket *Socket,
	callback func(event *console.MessageAddedEvent),
) {
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
}
