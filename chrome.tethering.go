package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Tethering - https://chromedevtools.github.io/devtools-protocol/tot/Tethering/
Defines methods and events for browser port binding.
*/
type Tethering struct{}

/*
Bind requests browser port binding.
*/
func (Tethering) Bind(socket *Socket, params *tethering.BindParams) error {
	command := &protocol.Command{
		method: "Tethering.bind",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Unbind requests browser port unbinding.
*/
func (Tethering) Unbind(socket *Socket, params *tethering.UnbindParams) error {
	command := &protocol.Command{
		method: "Tethering.unbind",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAccepted adds a handler to the Tethering.accepted event. Tethering.accepted fires when a port was
successfully bound and got a specified connection id.
*/
func (Tethering) OnAccepted(socket *Socket, callback func(event *tethering.AcceptedEvent)) error {
	handler := protocol.NewEventHandler(
		"Tethering.accepted",
		func(name string, params []byte) {
			event := &tethering.AcceptedEvent{}
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
