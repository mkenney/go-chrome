package protocol

import (
	"encoding/json"

	tethering "github.com/mkenney/go-chrome/protocol/tethering"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Tethering - https://chromedevtools.github.io/devtools-protocol/tot/Tethering/
Defines methods and events for browser port binding.
*/
type Tethering struct{}

/*
Bind requests browser port binding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-bind
*/
func (Tethering) Bind(
	socket *sock.Socket,
	params *tethering.BindParams,
) error {
	command := &sock.Command{
		Method: "Tethering.bind",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Unbind requests browser port unbinding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
*/
func (Tethering) Unbind(
	socket *sock.Socket,
	params *tethering.UnbindParams,
) error {
	command := &sock.Command{
		Method: "Tethering.unbind",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAccepted adds a handler to the Tethering.accepted event. Tethering.accepted fires when a port was
successfully bound and got a specified connection id.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#event-accepted
*/
func (Tethering) OnAccepted(
	socket *sock.Socket,
	callback func(event *tethering.AcceptedEvent),
) {
	handler := sock.NewEventHandler(
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
}
